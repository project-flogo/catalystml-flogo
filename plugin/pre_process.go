package plugin

import (
	"fmt"

	"github.com/project-flogo/cli/common"
	"github.com/project-flogo/cli/util"
)

func init() {
	// Register Build Pre-Processor to the flogo cli.
	common.RegisterBuildPreProcessor(&PreProcess{})
}

const (
	pkgPath = "github.com/project-flogo/catalystml-flogo/operations"
)

// Operations to be downloaded.
var operations []string = []string{"cleaning", "common", "math", "nlp", "restructuring", "string_processing"}

type PreProcess struct{}

func (p *PreProcess) DoPreProcessing(project common.AppProject, options common.BuildOptions) error {
	var importPaths []string

	currImports, err := common.CurrentProject().DepManager().GetAllImports()
	if err != nil {
		return err
	}

	// Build the Import path from the package name and the operations package name
	for _, val := range operations {

		importPath := pkgPath + "/" + val
		if _, ok := currImports[importPath]; !ok {
			importPaths = append(importPaths, importPath)
		}

	}

	// Parse the imports.
	imports, err := util.ParseImports(importPaths)
	if err != nil {
		return err
	}

	fmt.Println("Installing Operations.")
	// Install the imports to the app.
	return common.CurrentProject().AddImports(false, true, imports...)

}
