package ports

// ProjectManager handles actions related to project lifecycles
type ProjectManager interface {
	// InitRepo should initialize a project repository
	InitRepo(projectName string) error
}
