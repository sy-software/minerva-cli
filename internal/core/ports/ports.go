package ports

// ProjectManager handles actions related to project lifecycles
type ProjectManager interface {
	// InitRepo should initialize a project repository
	InitRepo(projectName string) error
}

// A convenience tool to run native test with more options
type TestRunner interface {
	// RunTests executes go native test in the path with optional watch mode
	RunTests(path string, watch bool, args []string) error
	// Discovers available tests in the given path
	DiscoverTests(path string) error
}
