package ns

/**
  汇集命名空间
 */
type ns_proxy struct {
	count int

	uts_ns *uts_namespace
	pid_ns *pid_namespace
}
