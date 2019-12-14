package main

import (
	"fmt"
	"net"
	"time"

	"github.com/hpifu/go-kit/hflag"
)

type MySubFlags struct {
	F1 int    `hflag:"--f1; default:20; usage:f1 flag"`
	F2 string `hflag:"--f2; default:hatlonely; usage:f2 flag"`
}

type MyFlags struct {
	I        int        `hflag:"--int, -i; required; default: 123; usage: int flag"`
	S        string     `hflag:"--str, -s; required; usage: str flag"`
	IntSlice []int      `hflag:"--int-slice; default: 1,2,3; usage: int slice flag"`
	IP       net.IP     `hflag:"--ip; usage: ip flag"`
	Time     time.Time  `hflag:"--time; usage: time flag; default: 2019-11-27"`
	Pos      string     `hflag:"pos; usage: pos flag"`
	Sub      MySubFlags `hflag:"sub"`
}

func main() {
	mf := &MyFlags{}
	if err := hflag.Bind(mf); err != nil {
		panic(err)
	}
	if err := hflag.Parse(); err != nil {
		fmt.Println(hflag.Usage())
		panic(err)
	}

	fmt.Printf("%#v\n", mf)
	fmt.Println("int =>", mf.I)
	fmt.Println("str =>", mf.S)
	fmt.Println("int-slice =>", mf.IntSlice)
	fmt.Println("ip =>", mf.IP)
	fmt.Println("time =>", mf.Time)
	fmt.Println("sub.f1 =>", mf.Sub.F1)
	fmt.Println("sub.f2 =>", mf.Sub.F2)
}
