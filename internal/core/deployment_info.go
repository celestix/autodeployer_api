package core

import (
	"encoding/json"
	"errors"

	"github.com/celestix/autodeployer_api/internal/db"
)

type ProjectDeploymentInfo interface {
	getType() db.ProjectType
	Deploy(*DeployParams) error
	Stop(string) error
}

func GetProjectDeploymentInfo(projectType db.ProjectType, buf []byte) (ProjectDeploymentInfo, error) {
	switch projectType {
	case db.ProjectTypeDockerfile:
		var v Dockerfile
		err := json.Unmarshal(buf, &v)
		if err != nil {
			return nil, err
		}
		return &v, nil
	case db.ProjectTypeCompose:
		var v DockerCompose
		err := json.Unmarshal(buf, &v)
		if err != nil {
			return nil, err
		}
		return &v, nil
	case db.ProjectTypeShell:
		var v ShellScript
		err := json.Unmarshal(buf, &v)
		if err != nil {
			return nil, err
		}
		return &v, nil
	case db.ProjectTypeCommand:
		var v SingleCommand
		err := json.Unmarshal(buf, &v)
		if err != nil {
			return nil, err
		}
		return &v, nil
	}
	return nil, errors.New("unknown project type")
}

type DeserializeFunc func([]byte) (DeployFunc, error)
type DeployFunc func() error

var ProjectDeployment = map[db.ProjectType]DeserializeFunc{
	db.ProjectTypeDockerfile: func(b []byte) (DeployFunc, error) {
		var v Dockerfile
		err := json.Unmarshal(b, &v)
		if err != nil {
			return nil, err
		}
		return func() error {
			return nil
		}, nil
	},
	db.ProjectTypeCompose: func(b []byte) (DeployFunc, error) {
		var v DockerCompose
		err := json.Unmarshal(b, &v)
		if err != nil {
			return nil, err
		}
		return func() error {
			return nil
		}, nil
	},
	db.ProjectTypeShell: func(b []byte) (DeployFunc, error) {
		var v ShellScript
		err := json.Unmarshal(b, &v)
		if err != nil {
			return nil, err
		}
		return func() error {
			return nil
		}, nil
	},
	db.ProjectTypeCommand: func(b []byte) (DeployFunc, error) {
		var v SingleCommand
		err := json.Unmarshal(b, &v)
		if err != nil {
			return nil, err
		}
		return func() error {
			return nil
		}, nil
	},
}
