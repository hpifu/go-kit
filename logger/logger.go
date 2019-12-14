package logger

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/lestrrat-go/file-rotatelogs"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"os"
	"path"
	"path/filepath"
	"runtime"
	"strings"
	"time"
)

var hostname string

func init() {
	var err error
	hostname, err = os.Hostname()
	if err != nil {
		hostname = "unknown"
	}
}

type CallerHook struct{}

func (hook *CallerHook) Levels() []logrus.Level {
	return logrus.AllLevels
}

func (hook *CallerHook) Fire(entry *logrus.Entry) error {
	for i := 5; i < 20; i++ {
		pc, file, line, ok := runtime.Caller(i)
		if !ok {
			break
		}
		if strings.Contains(file, "logrus") {
			continue
		}
		funcName := runtime.FuncForPC(pc).Name()
		entry.Data["@file"] = fmt.Sprintf("%s:%v:%s", path.Base(file), line, path.Base(funcName))
		break
	}

	return nil
}

type TextFormatter struct{}

func (f *TextFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	var buf bytes.Buffer
	for k, v := range entry.Data {
		if k[0] == '@' {
			continue
		}
		buf.WriteString(fmt.Sprintf("%v=%v,", k, v))
	}
	return []byte(fmt.Sprintf("[%v] [%v] [%v] [%v] [%v] %v\n", entry.Level, entry.Time.Format("2006-01-02 15:04:05"), entry.Data["@file"], entry.Data["@rid"], buf.String(), entry.Message)), nil
}

type JSONFormatter struct{}

func (f *JSONFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	entry.Data["@timestamp"] = time.Now().Format(time.RFC3339)
	entry.Data["@hostname"] = hostname
	buf, err := json.Marshal(entry.Data)
	if err != nil {
		return nil, err
	}
	return append(buf, '\n'), nil
}

type Options struct {
	Filename  string        `hflag:"usage: output filename"`
	MaxAge    time.Duration `hflag:"usage: keep time for retention"`
	Formatter string        `hflag:"usage: output format, json or text" hrule:"in json,text" hdef:"text"`
}

func NewLogger(options *Options) (*logrus.Logger, error) {
	log := logrus.New()
	if options.Formatter == "json" {
		log.Formatter = &JSONFormatter{}
	} else {
		log.Formatter = &TextFormatter{}
		log.AddHook(&CallerHook{})
	}
	if options.Filename == "" || options.Filename == "stdout" {
		return log, nil
	}

	abs, err := filepath.Abs(options.Filename)
	if err != nil {
		return nil, err
	}

	out, err := rotatelogs.New(
		abs+".%Y%m%d%H",
		rotatelogs.WithRotationTime(time.Hour),
		rotatelogs.WithLinkName(abs),
		rotatelogs.WithMaxAge(options.MaxAge),
	)
	if err != nil {
		return nil, err
	}
	log.Out = out

	return log, nil
}

func NewLoggerGroup(options []*Options) ([]*logrus.Logger, error) {
	var logs []*logrus.Logger
	for _, option := range options {
		log, err := NewLogger(option)
		if err != nil {
			return nil, err
		}
		logs = append(logs, log)
	}
	return logs, nil
}

func NewTextLogger(filename string, maxAge time.Duration) (*logrus.Logger, error) {
	return NewLogger(&Options{
		Filename:  filename,
		MaxAge:    maxAge,
		Formatter: "text",
	})
}

func NewJsonLogger(filename string, maxAge time.Duration) (*logrus.Logger, error) {
	return NewLogger(&Options{
		Filename:  filename,
		MaxAge:    maxAge,
		Formatter: "json",
	})
}

func NewTextLoggerWithViper(v *viper.Viper) (*logrus.Logger, error) {
	return NewTextLogger(v.GetString("filename"), v.GetDuration("maxAge"))
}

func NewJsonLoggerWithViper(v *viper.Viper) (*logrus.Logger, error) {
	return NewJsonLogger(v.GetString("filename"), v.GetDuration("maxAge"))
}

func NewLoggerGroupWithViper(v *viper.Viper) (*logrus.Logger, *logrus.Logger, *logrus.Logger, error) {
	infoLog, err := NewTextLoggerWithViper(v.Sub("infoLog"))
	if err != nil {
		return nil, nil, nil, err
	}
	warnLog, err := NewTextLoggerWithViper(v.Sub("warnLog"))
	if err != nil {
		return nil, nil, nil, err
	}
	accessLog, err := NewJsonLoggerWithViper(v.Sub("accessLog"))
	if err != nil {
		return nil, nil, nil, err
	}

	return infoLog, warnLog, accessLog, nil
}
