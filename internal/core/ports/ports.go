package ports

type ProjectManager interface {
	InitRepo(projectName string) error
}
