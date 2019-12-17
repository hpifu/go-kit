package hconf

import (
	. "github.com/smartystreets/goconvey/convey"
	"os"
	"testing"
	"time"
)

func CreateFile() {
	fp, _ := os.Create("test.json")
	_, _ = fp.WriteString(`{
  "scheduler": {
    "batch": 5
  },
  "reporter": {
    "address": "http://influxdb:8086",
    "database": "monitor",
    "retry": 3
  },
  "collector": [
    {
     "class": "CPUCollector",
     "table": "cpu",
     "interval": "30s"
    },
    {
      "class": "MemoryCollector",
      "table": "mem",
      "interval": "30s"
    },
    {
      "class": "LoadAvgCollector",
      "table": "load",
      "interval": "30s"
    },
    {
      "class": "DiskCollector",
      "table": "disk",
      "interval": "30s"
    }
  ],
  "logger": {
    "infoLog": {
      "filename": "log/monitor.info",
      "maxAge": "24h"
    },
    "warnLog": {
      "filename": "log/monitor.warn",
      "maxAge": "24h"
    },
    "accessLog": {
      "filename": "log/monitor.access",
      "maxAge": "24h"
    }
  }
}`)
	fp.Close()
}

func DeleteFile() {
	_ = os.Remove("test.json")
}

func TestHConfGet(t *testing.T) {
	Convey("test conf", t, func() {
		CreateFile()
		conf, err := New("json", "local", "test.json")
		So(err, ShouldBeNil)
		So(conf, ShouldNotBeNil)

		So(conf.GetDefaultInt("reporter.retry", 0), ShouldEqual, 3)
		So(conf.GetDefaultInt("reporter.retry"), ShouldEqual, 3)
		So(conf.GetDefaultInt("reporter.notexist"), ShouldEqual, 0)
		So(conf.GetDefaultDuration("collector[1].interval", time.Duration(0)), ShouldEqual, 30*time.Second)
		So(conf.GetDefaultDuration("logger.accessLog.maxAge"), ShouldEqual, 24*time.Hour)
		So(conf.GetDefaultFloat64("scheduler.batch"), ShouldAlmostEqual, 5.0)
		So(conf.GetDefaultString("collector[1].class"), ShouldEqual, "MemoryCollector")
		DeleteFile()
	})
}

func TestHConfSet(t *testing.T) {
	Convey("test conf set", t, func() {
		CreateFile()
		conf, err := New("json", "local", "test.json")
		So(err, ShouldBeNil)
		So(conf, ShouldNotBeNil)
		So(conf.Set("key1.key2", "val3"), ShouldBeNil)
		So(conf.GetDefaultString("key1.key2"), ShouldEqual, "val3")
		So(conf.Set("collector[1].interval", "20s"), ShouldBeNil)
		So(conf.GetDefaultString("collector[1].interval"), ShouldEqual, "20s")
		DeleteFile()
	})
}

func TestHConfBindEnv(t *testing.T) {
	Convey("test conf bind env", t, func() {
		CreateFile()
		conf, err := New("json", "local", "test.json")
		So(err, ShouldBeNil)
		So(conf, ShouldNotBeNil)

		conf.SetEnvPrefix("TEST")
		_ = os.Setenv("TEST_COLLECTOR_1_INTERVAL", "100s")
		So(conf.BindEnv("collector[1].interval"), ShouldBeNil)
		So(conf.GetDefaultString("collector[1].interval"), ShouldEqual, "100s")

		_ = os.Setenv("TEST_SCHEDULER_BATCH", "10")
		So(conf.BindEnv("scheduler.batch"), ShouldBeNil)
		So(conf.GetDefaultInt("scheduler.batch"), ShouldEqual, 10)
		DeleteFile()
	})
}

func TestHConfUnmarshal(t *testing.T) {
	type cinfo struct {
		Type     string `hconf:"class"`
		Table    string
		Interval string
	}

	type rinfo struct {
		Address  string
		Database string
		Retry    int
	}

	type linfo struct {
		Filename string
		MaxAge   time.Duration
		Format   string
	}

	type logger struct {
		Info   linfo  `hconf:"infoLog"`
		Warn   *linfo `hconf:"warnLog"`
		ErrLog linfo  `hconf:"-"`
	}

	Convey("test conf unmarshal", t, func() {
		CreateFile()
		conf, err := New("json", "local", "test.json")
		So(err, ShouldBeNil)
		So(conf, ShouldNotBeNil)

		{
			var infos []cinfo
			c, err := conf.Sub("collector")
			So(c, ShouldNotBeNil)
			So(err, ShouldBeNil)
			So(c.Unmarshal(&infos), ShouldBeNil)
			So(len(infos), ShouldEqual, 4)
			So(infos[0].Type, ShouldEqual, "CPUCollector")
			So(infos[0].Table, ShouldEqual, "cpu")
			So(infos[0].Interval, ShouldEqual, "30s")
		}

		{
			var info rinfo
			c, err := conf.Sub("reporter")
			So(c, ShouldNotBeNil)
			So(err, ShouldBeNil)
			So(c.Unmarshal(&info), ShouldBeNil)
			So(info.Retry, ShouldEqual, 3)
			So(info.Address, ShouldEqual, "http://influxdb:8086")
			So(info.Database, ShouldEqual, "monitor")
		}

		{
			var info logger
			c, err := conf.Sub("logger")
			So(c, ShouldNotBeNil)
			So(err, ShouldBeNil)
			So(c.Unmarshal(&info), ShouldBeNil)
			So(info.Info.Filename, ShouldEqual, "log/monitor.info")
			So(info.Warn.Filename, ShouldEqual, "log/monitor.warn")
			So(info.Info.MaxAge, ShouldEqual, 24*time.Hour)
		}
	})
}
