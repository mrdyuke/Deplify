package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
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
		readJSON, err := os.ReadFile("package.json")
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
		resp, err := http.Get("https://cold-meadow-d455.mrdyuke.workers.dev/")
		if err != nil {
			return nil, nil, err
		}
		defer resp.Body.Close()
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return nil, nil, err
		}
		var alternativesJSON AlternativesJSON
		err = json.Unmarshal(body, &alternativesJSON)
		if err != nil {
			return nil, nil, err
		}
		return nil, &alternativesJSON, nil

	default:
		return nil, nil, fmt.Errorf("unsupported JSON type: %s", jsonType)
	}
}

func main() {
	packageData, _, err := NewJSON("package")
	if err != nil {
		fmt.Println(err)
		return
	}

	_, alternativeData, err := NewJSON("alternatives")
	if err != nil {
		fmt.Println(err)
		return
	}

	packageDataArr := []string{}

	for packageNames := range packageData.Dependencies {
		packageDataArr = append(packageDataArr, packageNames)
	}

	for packageNames := range packageData.DevDependencies {
		packageDataArr = append(packageDataArr, packageNames)
	}

	for _, packageNames := range packageDataArr {
		alts, exists := alternativeData.Alternatives[packageNames]
		if !exists || len(alts) == 0 {
			continue
		}

		fmt.Printf("\n %s\n", color.GreenString(packageNames))
		for alternativeNames, description := range alts {
			fmt.Printf("  %s: %s\n", color.BlueString(alternativeNames), color.YellowString(description))

		}
		/* In my opinion, it's better to read the text with a slight delay,
		so I'll leave this as a comment in case you feel that way too. */

		// time.Sleep(200 * time.Millisecond)
	}
	fmt.Println("")
	fmt.Scanln()

}
