package cg

import "fmt"

type file_system_type struct {
	f_type string
}

type vfsmount struct {}


func cgroup_get_sb(fs_type* file_system_type, data *string, mnt *vfsmount) {

	// 创建空间
	root := kzalloc()

	// 绑定子系统
	rebind_subsystems(root, data)

	// 目录创建 cgroup 的管理文件
	cgroup_populate_dir(root.top_cgroup)
}

func kzalloc() *cgroupfs_root  {
	return &cgroupfs_root{}
}

func rebind_subsystems(root* cgroupfs_root, data *string)  {
	// 绑定超级块
	root.super_block = "dev"

	// 附加子系统
	root.subsys_list = append(root.subsys_list, mem_cgroup_subsys)
	root.number_of_cgroups = 1

	// 创建根cgroup对象,附加子系统统计信息
	mem_cgroup := mem_cgroup{
		cgroup_subsys_state: cgroup_subsys_state {
			cgroup: nil,
			flags: 0,
		},
		mem_info: "100M",
	}

	subsys := [CGROUP_SUBSYS_COUNT]*cgroup_subsys_state{ &mem_cgroup.cgroup_subsys_state }

	top_cgroup := cgroup {
		subsys: subsys,
		dentry: *data,
	}

	root.top_cgroup = &top_cgroup
}

func cgroup_populate_dir(cgroup *cgroup)  {
	fmt.Println("mouted dentry ", cgroup.dentry)
}

func find_task_by_vpid(pid int) *task_struct  {
	return &task_struct{
		task_id: pid,
	}
}

func find_css_set(cg *css_set, cgrp *cgroup) *css_set {
	task_cg := cg
	if cg == nil {
		task_cg = &css_set{}
	}

	// css_set用于收集不同cgroup上的子系统信息
	task_cg.subsys = cgrp.subsys
	return task_cg
}

func cgroup_attach_task(cgrp *cgroup, tsk * task_struct)  {

	// 根据新的cgroup查找css_set对象
	newcg := find_css_set(tsk.cgroups, cgrp)

	// 把进程的cgroups字段设置为新的css_set对象
	tsk.cgroups = newcg

	// 把进程添加到css_set对象的tasks列表中
	newcg.task = append(newcg.task, tsk)

	// 调用各个子系统的attach函数
	root := cgrp.root
	for i := 0; i < root.number_of_cgroups; i++ {
		subsys := root.subsys_list[i]
		subsys.attach(&subsys, nil, tsk)
	}

}
func attach_task_by_pid(cgrp *cgroup, pid int)  {
	tsk := find_task_by_vpid(pid)
	cgroup_attach_task(cgrp, tsk)
}

func Mounted(fs_type string, subsys string, dir string)  {
	cgroup_get_sb(&file_system_type{f_type: fs_type}, &dir, nil)
}