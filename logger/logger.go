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

func NewTextLogger(filename string, maxAge time.Duration) (*logrus.Logger, error) {
	log := logrus.New()
	log.Formatter = &TextFormatter{}
	log.AddHook(&CallerHook{})
	if filename == "" || filename == "stdout" {
		return log, nil
	}

	abs, err := filepath.Abs(filename)
	if err != nil {
		return nil, err
	}

	out, err := rotatelogs.New(
		abs+".%Y%m%d%H",
		rotatelogs.WithRotationTime(time.Hour),
		rotatelogs.WithLinkName(abs),
		rotatelogs.WithMaxAge(maxAge),
	)
	if err != nil {
		return nil, err
	}
	log.Out = out

	return log, nil
}

func NewJsonLogger(filename string, maxAge time.Duration) (*logrus.Logger, error) {
	log := logrus.New()
	log.Formatter = &JSONFormatter{}
	if filename == "" || filename == "stdout" {
		return log, nil
	}

	abs, err := filepath.Abs(filename)
	if err != nil {
		return nil, err
	}

	out, err := rotatelogs.New(
		abs+".%Y%m%d%H",
		rotatelogs.WithRotationTime(time.Hour),
		rotatelogs.WithLinkName(abs),
		rotatelogs.WithMaxAge(maxAge),
	)
	if err != nil {
		return nil, err
	}
	log.Out = out

	return log, nil
}

func NewTextLoggerWithViper(v *viper.Viper) (*logrus.Logger, error) {
	return NewTextLogger(v.GetString("filename"), v.GetDuration("maxAge"))
}

func NewJsonLoggerWithViper(v *viper.Viper) (*logrus.Logger, error) {
	return NewJsonLogger(v.GetString("filename"), v.GetDuration("maxAge"))
}
