package main

import (
	"fmt"
	"systheory/cg"
)

func main() {
	fmt.Println("hello world")

	//mount -t cgroup -o memory memory /sys/fs/cgroup/memory
	cg.Mounted("cgroup", "memory", "/sys/fs/cgroup/memory")
}
