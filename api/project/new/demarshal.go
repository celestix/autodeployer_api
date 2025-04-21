package new

import (
	"encoding/json"
	"errors"

	"github.com/celestix/autodeployer_api/internal/core"
	"github.com/celestix/autodeployer_api/internal/db"
)

func (n *NewProjectRequest) UnmarshalJSON(data []byte) error {
	var t struct {
		Name                 string         `json:"name"`
		Branch               string         `json:"branch"`
		Type                 db.ProjectType `json:"type"`
		RepoUrl              string         `json:"repo_url"`
		EnvironmentVariables [][]string     `json:"environment_variables"`
	}
	err := json.Unmarshal(data, &t)
	if err != nil {
		return err
	}
	var deploymentInfo core.ProjectDeploymentInfo
	switch t.Type {
	case db.ProjectTypeCommand:
		var d struct {
			DeploymentInfo *core.SingleCommand `json:"deployment_info"`
		}
		if err = json.Unmarshal(data, &d); err != nil {
			return err
		}
		deploymentInfo = d.DeploymentInfo
	case db.ProjectTypeCompose:
		var d struct {
			DeploymentInfo *core.DockerCompose `json:"deployment_info"`
		}
		if err = json.Unmarshal(data, &d); err != nil {
			return err
		}
		deploymentInfo = d.DeploymentInfo
	case db.ProjectTypeShell:
		var d struct {
			DeploymentInfo *core.ShellScript `json:"deployment_info"`
		}
		if err = json.Unmarshal(data, &d); err != nil {
			return err
		}
		deploymentInfo = d.DeploymentInfo
	case db.ProjectTypeDockerfile:
		var d struct {
			DeploymentInfo *core.Dockerfile `json:"deployment_info"`
		}
		if err = json.Unmarshal(data, &d); err != nil {
			return err
		}
		deploymentInfo = d.DeploymentInfo
	default:
		return errors.New("unknown project type")
	}
	*n = NewProjectRequest{
		Name:                 t.Name,
		Branch:               t.Branch,
		Type:                 t.Type,
		RepoUrl:              t.RepoUrl,
		EnvironmentVariables: t.EnvironmentVariables,
		DeploymentInfo:       deploymentInfo,
	}
	return nil
}
