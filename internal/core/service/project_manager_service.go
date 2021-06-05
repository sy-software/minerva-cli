package service

import "github.com/sy-software/minerva-cli/utils"

// Base template address for a service repository
const SERVICE_TEMPLATE = "https://github.com/sy-software/minerva-go-template.git"

// ProjectManagerService handles actions related to project lifecycles
type ProjectManagerService struct {
	fileUtils utils.FileUtilsInterface
}

// NewProjectManagerService creates a ProjectManagerService instance with default properties
func NewProjectManagerService() ProjectManagerService {
	return ProjectManagerService{
		fileUtils: utils.FileUtils{},
	}
}

// InitRepo initializes a project repository by cloning a template
func (manager ProjectManagerService) InitRepo(projectName string) (err error) {
	utils.Info("Initializing new service project")
	templatePath, err := utils.CloneTemplate(SERVICE_TEMPLATE)

	if err != nil {
		return
	}

	err = manager.fileUtils.CopyDir(templatePath, "./", []string{".git"})
	return
}
