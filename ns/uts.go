package ns

/**
uts命名空间包含了内核有关的信息，
*/
type uts_namespace struct {
	// 引用计数器
	kref int
	// 系统信息
	name utsname
}

type utsname struct {
	// 系统名称
	sysname string
	// 内核版本
	release string
	// 机器名称
	machine string
	// nis域名
	domainname string
}

const UTS_SYSNAME = "zhouseyi"
const UTS_RELEASE = "3.10.0-693.e17"
const UTS_MACHINE = "zsy machine"
const UTS_DOMAINNAME = "localdomain"

/**
  全局uts命名空间
 */
var init_uts_ns = uts_namespace {
	kref: 0,
	name: utsname{
		sysname: UTS_SYSNAME,
		release: UTS_RELEASE,
		machine: UTS_MACHINE,
		domainname: UTS_DOMAINNAME,
	},
}