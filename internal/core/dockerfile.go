package core

import (
	"errors"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"

	"github.com/celestix/autodeployer_api/internal/db"
)

type Dockerfile struct {
	FileLocation   string     `json:"file_location"`
	HostPorts      string     `json:"host_ports"`
	ContainerPorts string     `json:"container_ports"`
	Volumes        [][]string `json:"volumes"`
	NetworkMode    int        `json:"network_mode"`
	NetworkName    string     `json:"network_name"`
}

func (*Dockerfile) getType() db.ProjectType {
	return db.ProjectTypeDockerfile
}

func (s *Dockerfile) build(projectName, dir string) error {
	fmt.Println("dockerfile path:", s.FileLocation)
	fmt.Println("dockerfile dir:", dir)
	fmt.Println("dockerfile project name:", projectName)
	cmd := exec.Command("docker", "build", "-t", projectName, "-f", s.FileLocation, ".")
	cmd.Dir = dir
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func (s *Dockerfile) Deploy(params *DeployParams) error {
	if params == nil {
		return errors.New("params cant be nil")
	}
	fmt.Println("dockerfile brooooooooo")
	dir, err := params.getProjectDirectory()
	if err != nil {
		return err
	}
	if err := s.build(params.ProjectName, dir); err != nil {
		return err
	}
	fmt.Println("ran dockerfile")
	args := []string{"run", "-d", "--name", params.ProjectName}
	if s.NetworkMode == 1 {
		args = append(args, "--network", "host")
	}
	hostPorts := strings.Split(s.HostPorts, ",")
	containerPorts := strings.Split(s.ContainerPorts, ",")
	if len(hostPorts) != len(containerPorts) {
		return errors.New("invalid number of ports")
	}
	for i := 0; i < len(hostPorts); i++ {
		hostPort := strings.TrimSpace(hostPorts[i])
		containerPort := strings.TrimSpace(containerPorts[i])
		if hostPort == "" || containerPort == "" {
			continue
		}
		args = append(args, "-p", hostPort+":"+containerPort)
	}
	for _, volume := range s.Volumes {
		if volume[0] == "" || volume[1] == "" {
			continue
		}
		// args = append(args, "-v", strings.TrimSpace(volume[0])+":"+strings.TrimSpace(volume[1]))
	}
	cmd := exec.Command("docker", "rm", "-f", params.ProjectName)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Run()
	args = append(args, params.ProjectName)
	log.Println("Running docker command with args:", args)
	cmd = exec.Command("docker", args...)
	cmd.Env = params.EnvVars
	cmd.Dir = dir
	return cmd.Run()
}

func (s *Dockerfile) Stop(projectName string) error {
	cmd := exec.Command("docker", "stop", projectName)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		return err
	}
	cmd = exec.Command("docker", "rm", "-f", projectName)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Run()
	return nil
}
