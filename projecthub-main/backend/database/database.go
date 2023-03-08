package database

import (
	"context"
	"errors"
	"fmt"
	"projecthub/config"
	"projecthub/model"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var projectsMap = map[int]model.Project{}

func GetAllProjects() []model.Project {
	projects := make([]model.Project, 0, len(projectsMap))
	for _, v := range projectsMap {
		projects = append(projects, v)
	}

	return projects
}

func SaveProject(project model.Project) {
	projectsMap[project.StageID] = project
}

func GetProjectById(projectId int) (*model.Project, error) {
	if project, found := projectsMap[projectId]; found {
		return &project, nil
	}
	return nil, errors.New("project not found")
}

func GetExampleProjects() []model.Project {
	var projects = []model.Project{}

	projects = append(projects, model.Project{
		StageID:     11111,
		StageURL:    "https://www.hdm-stuttgart.de/stage/projekt_detail/projekt_details?projekt_ID=3877",
		Name:        "CoolApp1",
		EmbeddedURL: "http://127.0.0.1:5500/index.html",
	})

	projects = append(projects, model.Project{
		StageID:     22222,
		StageURL:    "https://www.hdm-stuttgart.de/stage/projekt_detail/projekt_details?projekt_ID=3886",
		Name:        "CoolApp2",
		EmbeddedURL: "http://127.0.0.1:5500/index.html",
	})

	projects = append(projects, model.Project{
		StageID:     33333,
		StageURL:    "https://www.hdm-stuttgart.de/stage/projekt_detail/projekt_details?projekt_ID=3885",
		Name:        "CoolApp3",
		EmbeddedURL: "http://127.0.0.1:5500/index.html",
	})

	projects = append(projects, model.Project{
		StageID:     44444,
		StageURL:    "https://www.hdm-stuttgart.de/stage/projekt_detail/projekt_details?projekt_ID=3884",
		Name:        "CoolApp4",
		EmbeddedURL: "http://127.0.0.1:5500/index.html",
	})

	projects = append(projects, model.Project{
		StageID:     55555,
		StageURL:    "https://www.hdm-stuttgart.de/stage/projekt_detail/projekt_details?projekt_ID=3877",
		Name:        "CoolApp1",
		EmbeddedURL: "http://127.0.0.1:5500/index.html",
	})

	projects = append(projects, model.Project{
		StageID:     66666,
		StageURL:    "https://www.hdm-stuttgart.de/stage/projekt_detail/projekt_details?projekt_ID=3882",
		Name:        "CoolApp6",
		EmbeddedURL: "http://127.0.0.1:5500/index.html",
	})

	return projects
}

func ConnectToDb() {
	// Create a new client and connect to the server
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(config.DatabaseUri))
	if err != nil {
		panic(err)
	}
	defer func() {
		if err = client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()
	// Ping the primary
	if err := client.Ping(context.TODO(), readpref.Primary()); err != nil {
		panic(err)
	}
	fmt.Println("Successfully connected and pinged.")
}
