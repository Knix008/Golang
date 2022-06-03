package main

import (
	"fmt"
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v3"
)

type Person struct {
	Name      string    `yaml:"name"`
	Job       string    `yaml:"job"`
	Skill     string    `yaml:"skill"`
	Employed  bool      `yaml:"employed"`
	Foods     []string  `yaml:"foods"`
	Languages Languages `yaml:"languages"`
}

type Languages struct {
	Perl   string `yaml:"perl"`
	Python string `yaml:"python"`
}

func getPerson(fileName string) (*Person, error) {
	buf, err := ioutil.ReadFile(fileName)
	if err != nil {
		return nil, err
	}
	p := &Person{}
	err = yaml.Unmarshal(buf, p)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}
	return p, nil
}

func main() {
	fmt.Println("Parsing YAML file...")
	p, err := getPerson("person.yaml")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%+v\n", p)
}
