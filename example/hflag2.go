package main

import (
	"fmt"
	"github.com/hpifu/go-kit/hflag"
	"time"
)

func main() {
	i := hflag.Int("int", 123, "int flag")
	s := hflag.String("str", "", "str flag")
	vi := hflag.IntSlice("int-slice", []int{1, 2, 3}, "int slice flag")
	ip := hflag.IP("ip", nil, "ip flag")
	t := hflag.Time("time", time.Now(), "time flag")
	if err := hflag.Parse(); err != nil {
		fmt.Println(hflag.Usage())
		panic(err)
	}

	fmt.Println("int =>", *i)
	fmt.Println("str =>", *s)
	fmt.Println("int-slice =>", *vi)
	fmt.Println("ip =>", *ip)
	fmt.Println("time =>", *t)
}
