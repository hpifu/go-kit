package main

import (
	"fmt"
	"github.com/hpifu/go-kit/hflag"
	"time"
)

func main() {
	hflag.Int("int", 123, "int flag")
	hflag.String("str", "", "str flag")
	hflag.IntSlice("int-slice", []int{1, 2, 3}, "int slice flag")
	hflag.IP("ip", nil, "ip flag")
	hflag.Time("time", time.Now(), "time flag")
	if err := hflag.Parse(); err != nil {
		panic(err)
	}

	fmt.Println("int =>", hflag.GetInt("int"))
	fmt.Println("str =>", hflag.GetString("str"))
	fmt.Println("int-slice =>", hflag.GetIntSlice("int-slice"))
	fmt.Println("ip =>", hflag.GetIP("ip"))
	fmt.Println("time =>", hflag.GetTime("time"))
}
