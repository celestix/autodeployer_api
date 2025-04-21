package core

import (
	"errors"
	"os/exec"

	"github.com/celestix/autodeployer_api/internal/db"
)

type ShellScript struct {
	FileLocation string `json:"file_location"`
}

func (*ShellScript) getType() db.ProjectType {
	return db.ProjectTypeShell
}

func (s *ShellScript) Deploy(params *DeployParams) error {
	if params == nil {
		return errors.New("params cant be nil")
	}
	dir, err := params.getProjectDirectory()
	if err != nil {
		return err
	}
	cmd := exec.Command(s.FileLocation)
	cmd.Env = params.EnvVars
	cmd.Dir = dir
	globalProcessStore.Set(params.ProjectName, cmd.Process)
	return cmd.Start()
}

func (s *ShellScript) Stop(projectName string) error {
	process := globalProcessStore.Get(projectName)
	if process == nil {
		return errors.New("unknown process")
	}
	return nil
}
