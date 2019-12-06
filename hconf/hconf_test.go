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
		conf, err := NewHConfWithFile("test.json")
		So(err, ShouldBeNil)
		So(conf, ShouldNotBeNil)

		So(conf.GetDefaultInt("reporter.retry", 0), ShouldEqual, 3)
		So(conf.GetDefaultInt("reporter.retry"), ShouldEqual, 3)
		So(conf.GetDefaultInt("reporter.notexist"), ShouldEqual, 0)
		So(conf.GetDefaultDuration("collector[1].interval", time.Duration(0)), ShouldEqual, 30*time.Second)
		So(conf.GetDefaultDuration("logger.accessLog.maxAge"), ShouldEqual, 24*time.Hour)
		So(conf.GetDefaultFloat("scheduler.batch"), ShouldAlmostEqual, 5.0)
		So(conf.GetDefaultString("collector[1].class"), ShouldEqual, "MemoryCollector")
		DeleteFile()
	})
}

func TestHConfSet(t *testing.T) {
	Convey("test conf set", t, func() {
		CreateFile()
		conf, err := NewHConfWithFile("test.json")
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
		conf, err := NewHConfWithFile("test.json")
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

func TestHConfUnMarshal(t *testing.T) {
	type cinfo struct {
		Class    string `json:"class"`
		Table    string `json:"table"`
		Interval string `json:"interval"`
	}

	type rinfo struct {
		Address  string `json:"address"`
		Database string `json:"database"`
		Retry    int    `json:"retry"`
	}

	type linfo struct {
		Filename string        `json:"filename"`
		MaxAge   time.Duration `json:"maxAge"`
	}

	type logger struct {
		InfoLog linfo  `json:"infoLog"`
		WarnLog *linfo `json:"warnLog"`
	}

	Convey("test conf unmarshal", t, func() {
		CreateFile()
		conf, err := NewHConfWithFile("test.json")
		So(err, ShouldBeNil)
		So(conf, ShouldNotBeNil)

		{
			var infos []cinfo
			c, err := conf.Sub("collector")
			So(c, ShouldNotBeNil)
			So(err, ShouldBeNil)
			So(c.Unmarshal(&infos), ShouldBeNil)
			So(len(infos), ShouldEqual, 4)
			So(infos[0].Class, ShouldEqual, "CPUCollector")
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
			So(info.InfoLog.Filename, ShouldEqual, "log/monitor.info")
			So(info.WarnLog.Filename, ShouldEqual, "log/monitor.warn")
			So(info.InfoLog.MaxAge, ShouldEqual, 24*time.Hour)
		}
	})
}

func TestHConfParse(t *testing.T) {
	Convey("test conf parse", t, func() {
		conf := &HConf{separator: "."}
		Convey("pass case1", func() {
			infos, err := conf.parseKey("key1.key2[3][4].key3")
			So(err, ShouldBeNil)
			So(len(infos), ShouldEqual, 5)
			So(infos[0].key, ShouldEqual, "key1")
			So(infos[0].mod, ShouldEqual, MapMod)
			So(infos[1].key, ShouldEqual, "key2")
			So(infos[1].mod, ShouldEqual, MapMod)
			So(infos[2].idx, ShouldEqual, 3)
			So(infos[2].mod, ShouldEqual, ArrMod)
			So(infos[3].idx, ShouldEqual, 4)
			So(infos[3].mod, ShouldEqual, ArrMod)
			So(infos[4].key, ShouldEqual, "key3")
			So(infos[4].mod, ShouldEqual, MapMod)
		})

		Convey("pass case2", func() {
			infos, err := conf.parseKey("[1][2].key3[4].key5")
			So(err, ShouldBeNil)
			So(len(infos), ShouldEqual, 5)
			So(infos[0].idx, ShouldEqual, 1)
			So(infos[0].mod, ShouldEqual, ArrMod)
			So(infos[1].idx, ShouldEqual, 2)
			So(infos[1].mod, ShouldEqual, ArrMod)
			So(infos[2].key, ShouldEqual, "key3")
			So(infos[2].mod, ShouldEqual, MapMod)
			So(infos[3].idx, ShouldEqual, 4)
			So(infos[3].mod, ShouldEqual, ArrMod)
			So(infos[4].key, ShouldEqual, "key5")
			So(infos[4].mod, ShouldEqual, MapMod)
		})

		Convey("fail case1", func() {
			infos, err := conf.parseKey("[1][key2].key3[4].key5")
			So(err, ShouldNotBeNil)
			So(infos, ShouldBeNil)
		})
	})
}
