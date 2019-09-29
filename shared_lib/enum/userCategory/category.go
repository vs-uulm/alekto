package userCategory

type Category = string

const (
	Default        Category = "all"
	Unprivileged   Category = "unprivileged"
	Standard       Category = "standard"
	Privileged     Category = "privileged"
	HighPrivileged Category = "highprivileged"
)
