package main

import (
	"html/template"
	"os"
	//"gopkg.in/yaml.v3"
)

type WebGl struct {
	InternalPort int
	ExternalPort int
	Index        string
	ConName      string
}

func main() {

	// befüllen der WebGl structur
	td := WebGl{80, 8085, "index.html", "fields"}

	//aufrufen der Funktion aus ingresstemp.go
	Ingtest()

	//leere Dateien erstellen um die später mit befüllten templates zu beschreiben
	j, err := os.Create("docker-compose.yaml")
	if err != nil {
		print("fail")
	}

	k, err := os.Create("webgl.conf")
	if err != nil {
		print("fail")
	}

	//templates mit WebGl struct befüllen
	t, err := template.New("webgl.tmp").ParseFiles("webgl.tmp")
	if err != nil {
		panic(err)
	}
	//template t mit td befüllen und in k reinschreiben
	err = t.Execute(k, td)
	if err != nil {
		panic(err)
	}

	d, err := template.New("docker-compose.tmp").ParseFiles("docker-compose.tmp")
	if err != nil {
		panic(err)
	}

	err = d.Execute(j, td)
	if err != nil {
		panic(err)
	}
}
