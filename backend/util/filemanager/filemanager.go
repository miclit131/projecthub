package filemanager

import (
	"fmt"
	"io/ioutil"
	"net/url"
	"os"
	"path"
	"github.com/go-git/go-git/v5"
)

func CloneGitRepo(projectGitURL string) (projectRepoPath string, err error) {
	// Validate URL
	_, err = url.ParseRequestURI(projectGitURL)
	if err != nil {
		return "", err
	}

	// Create temp directory
	projectDir, err := ioutil.TempDir("", "project-")
	if err != nil {
		return "", err
	}

	// Clone Git Repository
	repo, err := git.PlainClone(projectDir, false, &git.CloneOptions{
		URL:      projectGitURL,
		Progress: os.Stdout,
	})
	if err != nil {
		return "", err
	}
	fmt.Println(repo)

	a, _ := repo.Config()
	fmt.Println(a)

	return projectDir, nil
}

func DeleteDirectoryRecursively(dirPath string) error {
	if err := os.RemoveAll(dirPath); err != nil {
		return err
	}
	return nil
}


func ReadFile(filePath string) ([]byte, error) {
	// Parse config
	file, err := os.Open(path.Join(filePath))
	if err != nil {
		return nil, err
	}
	// defer the closing of file so it can be parsed later on
	defer file.Close()

	byteValue, _ := ioutil.ReadAll(file)

	fmt.Println(file)
	return byteValue, nil
	}


/* func ScanForFiles(dirPAth string) (filePaths []string, err error) {
	var files []string

    root := path.Join(dirPAth, "k8s", "templates")
	err = filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
        files = append(files, path)
        return nil
    })
    if err != nil {
        panic(err)
    }
    for _, file := range files {
        fmt.Println(file)
    }

	return files, nil
} */

	// Scan files in project dir
	/* 	files, err := ioutil.ReadDir(projectDir)
	   	if err != nil {
	   		return "", err
	   	}
	   	// Specific file path example
	   	projectDirName := files[0].Name()
	   	projectPath = path.Join(projectDir, projectDirName)
	   	fmt.Println(projectPath) */

	// WalkDir?
	// https://zetcode.com/golang/find-file/