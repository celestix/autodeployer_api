package db

type ProjectType int

const (
	_ ProjectType = iota
	ProjectTypeDockerfile
	ProjectTypeCompose
	ProjectTypeShell
	ProjectTypeCommand
)

type Project struct {
	Id         uint            `gorm:"primary_key;autoIncrement:true" json:"id"`
	EnvVars    []ProjectEnvVar `gorm:"foreignKey:ProjectId" json:"env"`
	Name       string          `json:"name"`
	Type       ProjectType     `json:"type"`
	RepoOwner  string          `json:"repo_owner"`
	RepoName   string          `json:"repo_name"`
	Branch     string          `json:"branch"`
	Gho        string          `json:"-"`
	Deployment []byte          `json:"-"`
}

type ProjectEnvVar struct {
	ProjectId uint   `gorm:"primary_key" json:"-"`
	Key       string `gorm:"primary_key" json:"key"`
	Value     string `json:"value"`
}

func AddProject(name string, ptype ProjectType, repoOwner, repoName, branch, gho string, deployment []byte, envVars []ProjectEnvVar) error {
	var w = Project{
		Name:       name,
		Type:       ptype,
		Branch:     branch,
		Deployment: deployment,
		EnvVars:    envVars,
		RepoOwner:  repoOwner,
		RepoName:   repoName,
		Gho:        gho,
	}
	tx := SESSION.Begin()
	tx.Create(&w)
	tx.Commit()
	return tx.Error
}

func GetProject(id uint) (*Project, error) {
	var w Project
	tx := SESSION.Begin()
	tx.First(&w, id)
	tx.Commit()
	return &w, tx.Error
}

func GetProjectByName(repoOwner, repoName string) (*Project, error) {
	var w Project
	tx := SESSION.Begin()
	tx.Where("repo_owner = ? AND repo_name = ?", repoOwner, repoName).First(&w)
	tx.Commit()
	return &w, tx.Error
}

func ListProjects() ([]Project, error) {
	var ws []Project
	tx := SESSION.Begin()
	tx.Find(&ws)
	tx.Commit()
	return ws, tx.Error
}
