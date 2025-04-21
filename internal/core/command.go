package core

import (
	"errors"
	"fmt"
	"os/exec"
	"strings"

	"github.com/celestix/autodeployer_api/internal/db"
)

type SingleCommand struct {
	Language string `json:"language"`
	Command  string `json:"command"`
}

func (*SingleCommand) getType() db.ProjectType {
	return db.ProjectTypeCommand
}

func (s *SingleCommand) Deploy(params *DeployParams) error {
	if params == nil {
		return errors.New("params cant be nil")
	}
	dir, err := params.getProjectDirectory()
	if err != nil {
		return err
	}
	fmt.Println(s.Command)
	cmdArgs := strings.Fields(s.Command)
	cmd := exec.Command(cmdArgs[0], cmdArgs[1:]...)
	cmd.Dir = dir
	cmd.Env = params.EnvVars
	globalProcessStore.Set(params.ProjectName, cmd.Process)
	return cmd.Start()
}

func (s *SingleCommand) Stop(projectName string) error {
	process := globalProcessStore.Get(projectName)
	if process == nil {
		return errors.New("unknown process")
	}
	return nil
}
