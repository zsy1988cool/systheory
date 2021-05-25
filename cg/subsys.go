package cg

const CGROUP_SUBSYS_COUNT = 1

type cgroup_subsys_state struct {
	cgroup *cgroup
	flags int
}

type mem_cgroup struct {
	cgroup_subsys_state
	mem_info string
}

type cpu_cgroup struct {
	cgroup_subsys_state
	cpu_info string
}

type cgroup_subsys struct {
	create_func func(ss* cgroup_subsys, cgrp* cgroup) *cgroup_subsys_state
	destroy func(ss* cgroup_subsys, cgrp* cgroup)
	attach func(ss* cgroup_subsys, cgrp* cgroup, tsk* task_struct)
	populate func(ss* cgroup_subsys, cgrp* cgroup)
}

func (m cgroup_subsys_state) getInfo() string {
	return ""
}

func (m mem_cgroup) getInfo() string {
	return m.mem_info
}

func (m cpu_cgroup) getInfo() string {
	return m.cpu_info
}