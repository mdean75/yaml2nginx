package main

import (
	"fmt"
	"github.com/aluttik/go-crossplane"
	"gopkg.in/yaml.v3"
	"os"
)

type server struct {
	Config struct {
		Directive string `yaml:"directive"`
		Block     []struct {
			Directive string   `yaml:"directive"`
			Args      []string `yaml:"args"`
		} `yaml:"block"`
	} `yaml:"config"`
}

func main() {
	b, err := os.ReadFile("server.yaml")
	if err != nil {
		fmt.Println(err)
		return
	}

	var data server
	err = yaml.Unmarshal(b, &data)
	if err != nil {
		fmt.Println(err)
		return
	}

	b, err = yaml.Marshal(data)
	if err != nil {
		fmt.Println(err)
		return
	}

	// base config object with non-nil Parsed slice
	cp := crossplane.Config{
		Parsed: []crossplane.Directive{},
	}

	// dataBlockList represents the slice of directives inside the server block
	var dataBlockList []crossplane.Directive
	for _, item := range data.Config.Block {
		dataBlockList = append(dataBlockList, crossplane.Directive{Directive: item.Directive, Args: item.Args})
	}

	// parsedList represents the base config object's Parsed slice of directives which is the server directive and slice of block directives
	parsedList := crossplane.Directive{Directive: data.Config.Directive, Block: &[]crossplane.Directive{}}

	// add the list of data directives to the server block
	*parsedList.Block = append(*parsedList.Block, dataBlockList...)

	// finally add all the directives to the root server ParsedList ... and BOOM!! didn't have to play with indexes
	cp.Parsed = append(cp.Parsed, parsedList)

	f, err := os.Create("server_new3.conf")
	if err != nil {
		fmt.Println(err)
		return
	}

	defer f.Close()

	err = crossplane.Build(f, cp, &crossplane.BuildOptions{})
	if err != nil {
		fmt.Println(err)
		return
	}

	f, err = os.Create("server_new2.conf")
	if err != nil {
		fmt.Println(err)
		return
	}
	f.Close()
	f.Write(nil)
}
