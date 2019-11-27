package main

import (
	"fmt"
	"github.com/hpifu/go-kit/hflag"
)

func main() {
	hflag.AddFlag("int", "int flag", hflag.Required(), hflag.Shorthand("i"), hflag.Type("int"), hflag.DefaultValue("123"))
	hflag.AddFlag("str", "str flag", hflag.Shorthand("s"), hflag.Required())
	hflag.AddFlag("int-slice", "int slice flag", hflag.Type("[]int"), hflag.DefaultValue("1,2,3"))
	hflag.AddFlag("ip", "ip flag", hflag.Type("ip"))
	hflag.AddFlag("time", "time flag", hflag.Type("time"), hflag.DefaultValue("2019-11-27"))
	if err := hflag.Parse(); err != nil {
		panic(err)
	}

	fmt.Println("int =>", hflag.GetInt("i"))
	fmt.Println("str =>", hflag.GetString("s"))
	fmt.Println("int-slice =>", hflag.GetIntSlice("int-slice"))
	fmt.Println("ip =>", hflag.GetIP("ip"))
	fmt.Println("time =>", hflag.GetTime("time"))
}
