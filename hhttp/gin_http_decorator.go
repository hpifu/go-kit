package hhttp

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/hpifu/go-kit/hrand"
	"github.com/sirupsen/logrus"
)

type GinHttpDecorator struct {
	InfoLog   *logrus.Logger
	WarnLog   *logrus.Logger
	AccessLog *logrus.Logger
}

func NewGinHttpDecorator(info, warn, access *logrus.Logger) *GinHttpDecorator {
	return &GinHttpDecorator{
		InfoLog:   info,
		WarnLog:   warn,
		AccessLog: access,
	}
}

type FileRes struct {
	Filename string
}

func (d *GinHttpDecorator) Decorate(inner func(string, *gin.Context) (interface{}, interface{}, int, error)) func(*gin.Context) {
	return func(c *gin.Context) {
		ts := time.Now()
		rid := c.DefaultQuery("rid", hrand.NewToken())
		req, res, status, err := inner(rid, c)
		if err != nil {
			c.String(status, err.Error())
			d.WarnLog.WithField("@rid", rid).WithField("err", err).Warn()
		} else if res == nil {
			c.Status(status)
		} else {
			switch res.(type) {
			case string:
				c.String(status, res.(string))
			case *FileRes:
				c.File(res.(*FileRes).Filename)
			default:
				c.JSON(status, res)
			}
		}

		errstr := ""
		if err != nil {
			errstr = err.Error()
		}
		d.AccessLog.WithFields(logrus.Fields{
			"client":    c.ClientIP(),
			"userAgent": c.GetHeader("User-Agent"),
			"host":      c.Request.Host,
			"method":    c.Request.Method,
			"url":       c.Request.URL.String(),
			"req":       req,
			"res":       res,
			"rid":       rid,
			"err":       errstr,
			"status":    status,
			"resTimeNs": time.Now().Sub(ts).Nanoseconds(),
		}).Info()
	}
}
