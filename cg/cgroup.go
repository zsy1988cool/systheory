package cg

type task_struct struct {
	task_id int
	cgroups *css_set
}
type cgroupfs_root struct {
	super_block string

	number_of_cgroups int
	top_cgroup        *cgroup
	subsys_list []cgroup_subsys

}

type css_set struct {
	list []css_set
	task []*task_struct
	subsys [CGROUP_SUBSYS_COUNT]*cgroup_subsys_state
}

type cgroup struct {
	count int

	parent   *cgroup
	sibling  *cgroup
	children []*cgroup
	top_cgroup *cgroup

	subsys [CGROUP_SUBSYS_COUNT]*cgroup_subsys_state

	root *cgroupfs_root

	css_sets []*css_set

	dentry string
}
