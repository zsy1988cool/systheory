package cg

import "fmt"

var mem_cgroup_subsys = cgroup_subsys {
	create_func: mem_cgroup_create,
	destroy: mem_cgroup_destroy,
	attach: mem_cgroup_attach,
	populate: mem_cgroup_populate,
}

func mem_cgroup_subsys_id()  {
	fmt.Println("mem_cgroup_subsys_id")
}

func mem_cgroup_create(ss* cgroup_subsys, cgrp* cgroup) *cgroup_subsys_state {
	fmt.Println("mem_cgroup_create")
	return nil
}

func mem_cgroup_destroy(ss* cgroup_subsys, cgrp* cgroup)  {
	fmt.Println("mem_cgroup_destroy")
}

func mem_cgroup_populate(ss* cgroup_subsys, cgrp* cgroup)  {
	fmt.Println("mem_cgroup_populate")
}

func mem_cgroup_attach(ss* cgroup_subsys, cgrp* cgroup, tsk* task_struct)  {
	fmt.Println("mem_cgroup_attach")
}

