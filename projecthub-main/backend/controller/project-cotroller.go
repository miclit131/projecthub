package controller

import (
	"fmt"
	"net/http"
	"projecthub/database"
	"projecthub/model"
	"projecthub/service/docker"
	"projecthub/service/kubernetes"
	"projecthub/util/configgenerator"
	"projecthub/util/filemanager"
	"strconv"
	"github.com/gin-gonic/gin"
)

// Create new project from POST /project
func CreateProject(c *gin.Context) {
	// Validate user input
	var input model.CreateProjectDTO
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create Project from input
	project, err := model.NewProject(input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Download project
	projectRepoPath, err := filemanager.CloneGitRepo(project.GitURL)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Build image(s) //TODO return paths if multiple images
	if err := docker.BuildImage(project.ImageTag, projectRepoPath); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Push image(s) to registry
	if err := docker.PushImageToRegistry(project.ImageTag); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create Kubernetes config
	if err := configgenerator.CreateK8sConfigFromDockerCompose(projectRepoPath, project.K8sConfigPath); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Inject data into Kubernetes config
	if err := configgenerator.InjectValuesIntoK8sConfig(project.K8sConfigPath, project.ImageTag); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Deploy on Kubernetes simple
	if err := kubernetes.ApplySimple(project.K8sConfigPath); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Deploy on Kubernetes with client-go
	/* 		if err := kubernetes.Apply2(project.K8sConfigPath); err != nil {
	   		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	   		return
	   	} */

	fmt.Println(project)

	// Save to DB
	database.SaveProject(*project)

	// Delete image(s) from Docker
	if err := docker.DeleteImage(project.ImageTag); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Delete local project repository
	if err := filemanager.DeleteDirectoryRecursively(projectRepoPath); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, project)
}

// Returns all saved Projects GET /projects
func GetAllProjects(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"projects": database.GetAllProjects(),
	})
}

// Return Project for given ID GET /project/:projectid
func GetProjectById(c *gin.Context) {
	projectIdString := c.Param("projectid")
	projectId, err := strconv.Atoi(projectIdString)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	project, err := database.GetProjectById(projectId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, project)
}

func GetExampleProjects(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"projects": database.GetExampleProjects(),
	})
}
