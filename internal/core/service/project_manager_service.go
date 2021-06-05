package service

import "github.com/sy-software/minerva-cli/utils"

// Base template address for a service repository
const SERVICE_TEMPLATE = "https://github.com/sy-software/minerva-go-template.git"

type ProjectManagerService struct {
	fileUtils utils.FileUtilsInterface
}

func NewProjectManagerService() ProjectManagerService {
	return ProjectManagerService{
		fileUtils: utils.FileUtils{},
	}
}

func (manager ProjectManagerService) InitRepo(projectName string) (err error) {
	utils.Info("Initializing new service project")
	templatePath, err := utils.CloneTemplate(SERVICE_TEMPLATE)

	if err != nil {
		return
	}

	err = manager.fileUtils.CopyDir(templatePath, "./", []string{".git"})
	return
}
