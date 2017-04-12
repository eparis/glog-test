example usage:
./glog-test --logtostderr --v=1 --vmodule="file=2"
This prints BOTH path1 and path2 V(2)

I had hoped that I could do
./glog-test --logtostderr --v=1 --vmodule="path1/file=2"
But it didn't do anything...
