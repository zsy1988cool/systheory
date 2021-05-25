package cg

import "fmt"

var cpu_cgroup_subsys = cgroup_subsys {
	create_func: cpu_cgroup_create,
	destroy: cpu_cgroup_destroy,
	attach: cpu_cgroup_attach,
	populate: cpu_cgroup_populate,
}

func cpu_cgroup_subsys_id()  {
	fmt.Println("cpu_cgroup_subsys_id")
}

func cpu_cgroup_create(ss* cgroup_subsys, cgrp* cgroup) *cgroup_subsys_state {
	fmt.Println("cpu_cgroup_create")
	return nil
}


func cpu_cgroup_destroy(ss* cgroup_subsys, cgrp* cgroup)  {
	fmt.Println("cpu_cgroup_destroy")
}

func cpu_cgroup_populate(ss* cgroup_subsys, cgrp* cgroup)  {
	fmt.Println("cpu_cgroup_populate")
}

func cpu_cgroup_attach(ss* cgroup_subsys, cgrp* cgroup, tsk* task_struct)  {
	fmt.Println("cpu_cgroup_attach")
}

