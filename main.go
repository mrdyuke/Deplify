package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/fatih/color"
)

// Structures
type PkgJSON struct {
	Name            string            `json:"name"`
	Version         string            `json:"version"`
	Dependencies    map[string]string `json:"dependencies"`
	DevDependencies map[string]string `json:"devDependencies"`
}

type AltJSON struct {
	Alternatives map[string]map[string]string `json:"alternatives"`
}

// Constructors
func NewPkgJSON(p string) (*PkgJSON, error) {
	read, err := os.ReadFile(p)
	if err != nil {
		return nil, err
	}

	var pkg PkgJSON
	err = json.Unmarshal(read, &pkg)
	if err != nil {
		return nil, err
	}

	return &pkg, nil
}

func NewAltJSON(p string) (*AltJSON, error) {
	resp, err := http.Get(p)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var alt AltJSON
	err = json.Unmarshal(body, &alt)
	if err != nil {
		return nil, err
	}

	return &alt, nil
}

// Methods
func (pkg *PkgJSON) CombineDeps() []string {
	pkgNames := make([]string, 0, len(pkg.Dependencies)+len(pkg.DevDependencies))
	for name := range pkg.Dependencies {
		pkgNames = append(pkgNames, name)
	}
	for name := range pkg.DevDependencies {
		pkgNames = append(pkgNames, name)
	}
	return pkgNames
}

func main() {
	pkgJSON, err := NewPkgJSON("package.json")
	if err != nil {
		fmt.Printf("\n %s\n\n", color.RedString("%v", err))
		return
	}

	altJSON, err := NewAltJSON("https://cold-meadow-d455.mrdyuke.workers.dev/")
	if err != nil {
		fmt.Printf("\n %s\n\n", color.RedString("%v", err))
		return
	}

	fmt.Printf("\n %s%s\n", color.BlueString("Project: "), color.GreenString(pkgJSON.Name))
	fmt.Printf(" %s%s\n\n", color.BlueString("Version: "), color.GreenString(pkgJSON.Version))
	fmt.Printf(" %s\n", strings.Repeat("━", len(pkgJSON.Name)+len(pkgJSON.Version)+10))

	pkgNames := pkgJSON.CombineDeps()

	for _, pkgName := range pkgNames {
		altMap, ok := altJSON.Alternatives[pkgName]
		if !ok || len(altMap) == 0 {
			continue
		}
		fmt.Printf("\n %s%s\n", color.BlueString("Package: "), color.WhiteString(pkgName))

		for altName, altDescr := range altMap {
			fmt.Printf(" %s%s\n", color.GreenString(" ⚬ %s: ", altName), color.YellowString(altDescr))
		}
	}

	fmt.Println()

}
