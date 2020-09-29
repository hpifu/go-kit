package hashx

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestHasher32(t *testing.T) {
	Convey("test hasher32", t, func() {
		for _, info := range []struct {
			algorithm HashAlgorithm32
			value     uint32
		}{
			{FNV32, 1418570095},
			{FNV32A, 3582672807},
			{CRC32IEEE, 222957957},
			{CRC32CASTAGNOLI, 3381945770},
			{CRC32KOOPMAN, 3744939324},
			{ADLER32, 436929629},
			{MURMUR32, 1586663183},
		} {
			So(NewHasher32(info.algorithm).Hash32S("hello world"), ShouldEqual, info.value)
		}
	})
}

func TestHasher64(t *testing.T) {
	Convey("test hasher64", t, func() {
		for _, info := range []struct {
			algorithm HashAlgorithm64
			value     uint64
		}{
			{FNV64, 9065573210506989167},
			{FNV64A, 8618312879776256743},
			{CRC64ISO, 13388989860809387070},
			{CRC64ECMA, 5981764153023615706},
			{MURMUR64, 5998619086395760910},
		} {
			So(NewHasher64(info.algorithm).Hash64S("hello world"), ShouldEqual, info.value)
		}
	})
}

func TestHash128(t *testing.T) {
	Convey("test hasher128", t, func() {
		for _, info := range []struct {
			algorithm HashAlgorithm128
			high      uint64
			low       uint64
		}{
			{FNV128, 16262890844614405877, 6225721494403853343},
			{FNV128A, 7788227449506557636, 13336604906580551351},
			{MURMUR128, 5998619086395760910, 12364428806279881649},
		} {
			high, low := NewHasher128(info.algorithm).Hash128S("hello world")
			So(high, ShouldEqual, info.high)
			So(low, ShouldEqual, info.low)
		}
	})
}

func TestMd5Hash(t *testing.T) {
	s := "17744581949"
	sMd5 := Md5Hash(s)
	if sMd5 == "f72da137dae79425eb9c060234185a55" {
		t.Log(sMd5)
	} else {
		t.Errorf("the %s md5 %s may be not right.", s, sMd5)
	}
}

func TestBKDRHash(t *testing.T) {
	s := "f72da137dae79425eb9c060234185a55"
	ih := BKDRHash(s)
	if ih == 721271279 {
		t.Log(ih)
	} else {
		t.Errorf("the %s bkdr %v may be not right.", s, ih)
	}
}

func TestDJBHash(t *testing.T) {
	s := "f72da137dae79425eb9c060234185a55"
	ih := DJBHash(s)
	if ih == 1400130340 {
		t.Log(ih)
	} else {
		t.Errorf("the %s bkdr %v may be not right.", s, ih)
	}
}

func TestELFHash(t *testing.T) {
	s := "f72da137dae79425eb9c060234185a55"
	ih := ELFHash(s)
	if ih == 1733014661 {
		t.Log(ih)
	} else {
		t.Errorf("the %s bkdr %v may be not right.", s, ih)
	}
}

func TestFNVHash(t *testing.T) {
	s := "f72da137dae79425eb9c060234185a55"
	ih := FNVHash(s)
	if ih == 933970133 {
		t.Log(ih)
	} else {
		t.Errorf("the %s bkdr %v may be not right.", s, ih)
	}
}
