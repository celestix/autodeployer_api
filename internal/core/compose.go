package core

import "github.com/celestix/autodeployer_api/internal/db"

type DockerCompose struct {
	FileLocation string `json:"file_location"`
}

func (*DockerCompose) getType() db.ProjectType {
	return db.ProjectTypeCompose
}

func (*DockerCompose) Deploy(opts *DeployParams) error {
	return nil
}

func (s *DockerCompose) Stop(projectName string) error {
	return nil
}
