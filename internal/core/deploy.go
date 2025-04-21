package core

import (
	"fmt"

	"github.com/celestix/autodeployer_api/config"
)

type DeployParams struct {
	ProjectName string
	RepoOwner   string
	RepoName    string
	EnvVars     []string
}

func (d *DeployParams) getProjectDirectory() (string, error) {
	if d.RepoName == "" || d.RepoOwner == "" {
		return "", fmt.Errorf("RepoName and RepoOwner must be set")
	}
	return fmt.Sprintf("%s/%s", config.Data.DataDirectory, d.ProjectName), nil
}

func DeployProject(dInfo ProjectDeploymentInfo, opts *DeployParams) error {
	return dInfo.Deploy(opts)
}
