package main

import (
	"flag"
	"fmt"
	"os"
)

//需要赋值的变量
var version = ""

//通过flag包设置-version参数
var printVersion bool

func init() {
	flag.BoolVar(&printVersion, "version", false, "print program build version")
	flag.Parse()
}

func main() {
	if printVersion {
		println(version)
		os.Exit(0)
	}
	fmt.Printf("example for print version")
}
