package main

import (
	"fmt"
	"systheory/cg"
)

func main() {
	fmt.Println("演示cgroup")

	//mount -t cgroup -o memory memory /sys/fs/cgroup/memory
	cg.Mounted("cgroup", "memory", "/sys/fs/cgroup/memory")

	// echo 123012 > /sys/fs/cgroup/memory/tasks
	cg.AttchTask(123, "/sys/fs/cgroup/memory/tasks")
}
