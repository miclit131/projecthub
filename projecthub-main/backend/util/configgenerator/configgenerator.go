package configgenerator

import (
	"fmt"
	"io/ioutil"
	"os/exec"
	"path"
	"projecthub/config"
	"projecthub/util/filemanager"
	"gopkg.in/yaml.v3"
)

//TODO can be .yaml or .yml
//kompose-cli is written in go, future TODO?: import github.com/kubernetes/kompose
func CreateK8sConfigFromDockerCompose(projectPath string, targetPath string) error {
	cmd := exec.Command("kompose", "convert", "-f", path.Join(projectPath, "docker-compose.yaml"), "-o", targetPath, "--with-kompose-annotation=false")
	stdout, err := cmd.Output()

	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	fmt.Println(string(stdout))

	return nil
}

func InjectValuesIntoK8sConfig(objConfigPath string, imageTag string) error {
	// read obj config
	objBytes, err := filemanager.ReadFile(objConfigPath)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(objBytes))

	// unmarshal config
	objUnstructured := make(map[string]interface{})
	err = yaml.Unmarshal(objBytes, &objUnstructured)
	if err != nil {
		panic(err)
	}

	// iterate over items and inject sth for each kind
	for _, item := range objUnstructured["items"].([]interface{}) {

		var kind = item.(map[string]interface{})["kind"].(string)

		switch kind {

		case "Service":
			//TODO

		case "Deployment":
			//appsv1.Deployment{} //TODO?
			//fmt.Println(scheme.Scheme.AllKnownTypes())
			item.(map[string]interface{})["spec"].(map[string]interface{})["template"].(map[string]interface{})["spec"].(map[string]interface{})["containers"].([]interface{})[0].(map[string]interface{})["image"] = imageTag
		}
	}

	// inject ingress
	objUnstructured["items"] = append(objUnstructured["items"].([]interface{}), map[string]interface{}{
		"apiVersion": "networking.k8s.io/v1",
		"kind":       "Ingress",
		"metadata": map[string]interface{}{
			"name": "yeet3",
			"annotations": map[string]interface{}{
				"nginx.ingress.kubernetes.io/rewrite-target": "/",
			},
		},
		"spec": map[string]interface{}{
			"ingressClassName": "nginx",
			"rules": []map[string]interface{}{
				{
					"host": config.K8sIngressHost /*+ "/" + project.RepoName*/,
					"http": map[string]interface{}{
						"paths": []map[string]interface{}{
							{
								"pathType": "Prefix",
								"path":     "/",
								"backend": map[string]interface{}{
									"service": map[string]interface{}{
										"name": "webgl",
										"port": map[string]interface{}{
											"number": 80,
										},
									},
								},
							},
						},
					},
				},
			},
		},
	})

	// marshal config
	data, err := yaml.Marshal(objUnstructured)
	if err != nil {
		panic(err)
	}

	// write obj config
	err = ioutil.WriteFile(objConfigPath, data, 0)
	if err != nil {
		panic(err)
	}

	return nil
}
