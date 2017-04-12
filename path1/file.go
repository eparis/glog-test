package path1

import (
	"github.com/golang/glog"
)

func MyFunc() {
	glog.V(1).Infof("V(1) in package 1")
	glog.V(2).Infof("V(2) in package 1")
}
