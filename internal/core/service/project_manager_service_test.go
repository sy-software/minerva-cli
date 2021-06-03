package service

import (
	"testing"

	"github.com/sy-software/minerva-cli/utils"
)

type DummyFileUtils struct {
	t *testing.T
}

func (instance DummyFileUtils) CopyFile(src, dst string) (err error) {

	return
}
func (instance DummyFileUtils) CopyDir(src string, dst string, omit []string) (err error) {
	if dst != "./" {
		instance.t.Errorf("Destination must be current directory got %q", dst)
	}

	if len(omit) < 1 {
		instance.t.Errorf("At least .git dir must be omitted got none")
	} else if !utils.Contains(omit, ".git") {
		instance.t.Errorf(".git dir must be ignored got %v", omit)
	}

	return
}

func TestTemplateIsCopied(t *testing.T) {
	instance := ProjectManagerService{
		fileUtils: DummyFileUtils{
			t: t,
		},
	}

	err := instance.InitRepo("test")

	if err != nil {
		t.Errorf("Template should be copied: %s", err)
	}
}
