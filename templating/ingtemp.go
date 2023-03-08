package main

import (
	"html/template"
	"log"
	"os"
	//"gopkg.in/yaml.v3"
)

type Ingre struct {
	ServiceName string
	ServicePort int
}

func Ingtest() {

	i, err := os.Create("ingress.yaml")
	if err != nil {
		print("fail1")
	}

	f, err := os.OpenFile("deplay.yaml", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}

	tdi := Ingre{"yeet", 80}

	t, err := template.New("ingress.tmp").ParseFiles("ingress.tmp")
	if err != nil {
		print("fail2")
		panic(err)
	}
	err = t.Execute(i, tdi)
	if err != nil {
		print("fail3")
		panic(err)
	}

	err = t.Execute(f, tdi)
	if err != nil {
		print("fail3")
		panic(err)
	}

}
