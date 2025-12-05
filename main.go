package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/fatih/color"
)

type PackageJSON struct {
	Info            string            `json:"info"`
	ProjectName     string            `json:"name"`
	ProjectVersion  string            `json:"version"`
	Dependencies    map[string]string `json:"dependencies"`
	DevDependencies map[string]string `json:"devDependencies"`
}

type Alternative struct {
	Alternatives map[string]map[string]string `json:"alternatives"`
}

func main() {
	packageData, err := os.ReadFile("package.mock.json")
	if err != nil {
		color.Red("%s", err)
		return
	}

	var packageJSON PackageJSON
	err = json.Unmarshal(packageData, &packageJSON)
	if err != nil {
		color.Red("%s", err)
		return
	}

	alternativeData, err := os.ReadFile("alternatives.json")
	if err != nil {
		color.Red("%s", err)
		return
	}

	var alternativesJSON Alternative
	err = json.Unmarshal(alternativeData, &alternativesJSON)
	if err != nil {
		color.Red("%s", err)
		return
	}

	// Удалить напоминание когда основной функционал будет готов
	fmt.Printf("\n %s \n", color.RedString(packageJSON.Info))

	fmt.Printf("\n Project name: %s", color.GreenString(packageJSON.ProjectName))
	fmt.Printf("\n Version: %s \n", color.GreenString(packageJSON.ProjectVersion))
	fmt.Printf("\n %s \n", strings.Repeat("=", 30))

	// реализовать поиск совпадений ключей в alternativesJSON и packageJSON
	// если ключи совпадают, то вывести ключ и значение из packageJSON

}
