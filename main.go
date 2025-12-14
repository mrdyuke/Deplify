package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/fatih/color"
)

type PackageJSON struct {
	Dependencies    map[string]string `json:"dependencies"`
	DevDependencies map[string]string `json:"devDependencies"`
}

type AlternativesJSON struct {
	Alternatives map[string]map[string]string `json:"alternatives"`
}

func NewJSON(jsonType string) (*PackageJSON, *AlternativesJSON, error) {
	switch jsonType {

	case "package":
		readJSON, err := os.ReadFile("package.mock.json")
		if err != nil {
			return nil, nil, err
		}
		var packageJSON PackageJSON
		err = json.Unmarshal(readJSON, &packageJSON)
		if err != nil {
			return nil, nil, err
		}
		return &packageJSON, nil, nil

	case "alternatives":
		readJSON, err := os.ReadFile("alternatives.json")
		if err != nil {
			return nil, nil, err
		}
		var alternativesJSON AlternativesJSON
		err = json.Unmarshal(readJSON, &alternativesJSON)
		if err != nil {
			return nil, nil, err
		}
		return nil, &alternativesJSON, nil

	default:
		return nil, nil, fmt.Errorf("unsupported JSON type: %s", jsonType)
	}
}

func main() {
	packagesData, _, err := NewJSON("package")
	if err != nil {
		fmt.Println(err)
		return
	}

	/* 	_, alternativesData, err := NewJSON("alternatives")
	   	if err != nil {
	   		fmt.Println(err)
	   		return
	   	} */

	packageDataArr := []string{}

	for packageNames := range packagesData.Dependencies {
		packageDataArr = append(packageDataArr, packageNames)
	}

	for packageNames := range packagesData.DevDependencies {
		packageDataArr = append(packageDataArr, packageNames)
	}

	for _, packagesData := range packageDataArr {
		fmt.Printf("\n %s \n", color.GreenString(packagesData))
	}
}
