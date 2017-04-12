package main

import (
	"flag"

	"github.com/eparis/glog-test/path1"
	"github.com/eparis/glog-test/path2"
)

func main() {
	flag.Parse()
	path1.MyFunc()
	path2.MyFunc()
}
