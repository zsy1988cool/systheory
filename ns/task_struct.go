package ns

/**
  进程对象
 */
type task_struct struct {

	// 关联到进程的命名空间
	ns_proxy *ns_proxy

	// 关联的pid对象
	pids [PIDTYPE_MAX]pid_link
}
