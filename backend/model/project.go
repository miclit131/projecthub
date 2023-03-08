package model

import (
	"errors"
	"path"
	"projecthub/config"
	"strconv"
	"github.com/oriser/regroup"
)

type CreateProjectDTO struct {
	StageID int    `json:"stageId"`
	Name    string `json:"name"`
	GitURL  string `json:"gitUrl"`
}

type Project struct {
	StageID       int    `json:"stageId"`
	StageURL      string `json:"stageUrl"`
	Name          string `json:"name"`
	EmbeddedURL	  string `json:"embeddedUrl"`
	GitURL        string `json:"-"`
	RepoName      string `regroup:"name" json:"-"`
	RepoOwner     string `regroup:"owner" json:"-"`
	ImageTag      string `json:"-"`
	K8sConfigPath string `json:"-"`
}

// Project constructor
func NewProject(input CreateProjectDTO) (*Project, error) {
	// Check if input is valid
	if input.StageID < 1 { //TODO check if already present in DB
		return nil, errors.New("INVALID STAGE ID")
	}
	if input.Name == "" {
		return nil, errors.New("INVALID NAME")
	}
	if input.GitURL == "" {
		return nil, errors.New("INVALID URL")
	}

	// Init project object
	project := &Project{
		StageID:  input.StageID,
		StageURL: "https://www.hdm-stuttgart.de/stage/projekt_detail/projekt_details?projekt_ID=" + strconv.Itoa(input.StageID),
		Name:     input.Name,
		GitURL:   input.GitURL,
	}

	// Parse repo name and owner from URL
	r := regroup.MustCompile(`(?P<host>(git@|https://)([\w\.@\-]+)(/|:))(?P<owner>[\w,\-,\_]+)/(?P<name>[\w,\-,\_]+)(.git){0,1}((/){0,1})`)
	if err := r.MatchToTarget(project.GitURL, project); err != nil {
		return nil, err
	}

	// Compose image tag and k8s config path
	project.ImageTag = config.RegistryAddress + "/" + config.RegistryUsername + "/" + project.RepoName + "_" + strconv.Itoa(project.StageID)
	project.K8sConfigPath = path.Join(config.K8sProjectObjPath, project.RepoName+"_"+strconv.Itoa(project.StageID)+".yaml")
	project.EmbeddedURL = "http://" + config.K8sIngressHost /*+ "/" + project.RepoName*/ //TODO

	return project, nil
}
