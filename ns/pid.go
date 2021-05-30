package ns

const PIDTYPE_MAX  = 3

type pid_type int

const (
	// 进程id
	PIDTYPE_PID pid_type = 0
	// 进程组id
	PIDTYPE_PGID pid_type = 1
	// 会话组id
	PIDTYPE_SID pid_type = 2
)

/**
  进程id对象,是内核对pid的内部表示
 */
type pid_t struct {

	// 引用计数器
	count int

	// 使用该pid的进程的列表
	tasks [PIDTYPE_MAX]*task_struct

	// 级别
	level int

	// 命名空间内pid标识
	numbers [2]upid
}

/**
  特定命名空间中可见的pid信息
 */
type upid struct {
	// pid数值
	nr int
	// 指向该id所属的命名空间
	ns* pid_namespace
	// 所有的upid链表
	// pid_chain
}

/**
  pid命名空间包含了进程id有关的信息
*/
type pid_namespace struct {

	// 命名空间的等级
	level int

	// 父命名空间
	parent* pid_namespace
}

/**
  用于进程连接pid的关联
 */
type pid_link struct {
	pid *pid_t
	node int
}

func attach_pid(tsk *task_struct, tpe pid_type, pid *pid_t)  {

}