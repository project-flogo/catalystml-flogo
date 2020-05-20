package runcml

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"testing"
	"text/template"

	"github.com/stretchr/testify/assert"
)

var (
	files             = []string{"runcml.go", "samplecml.json", "metadata.go"}
	runCMLTest string = "runcml_test.go"
)

func TestRunCml(t *testing.T) {
	pwd, err := os.Getwd()
	assert.Nil(t, err)

	tempDir, err := ioutil.TempDir("", "runcml")
	assert.Nil(t, err)

	err = copyFiles(pwd, tempDir)
	assert.Nil(t, err)

	err = runTest(tempDir)
	assert.Nil(t, err)

	os.RemoveAll(tempDir)
}

var testFile string = `
package runcml

import (
	"testing"

	_ "github.com/project-flogo/catalystml-flogo/action"
	_ "github.com/project-flogo/catalystml-flogo/operations/common"
	_ "github.com/project-flogo/catalystml-flogo/operations/image_processing"
	_ "github.com/project-flogo/catalystml-flogo/operations/nlp"
	_ "github.com/project-flogo/catalystml-flogo/operations/string_processing"
	"github.com/project-flogo/catalystml-flogo/action/support/test"
	"github.com/stretchr/testify/assert"
)

func TestSample(t *testing.T) {
	inputs := make(map[string]interface{})
	inputs["data"] = map[string]interface{}{"paragraph": "Abc"}
	params := Params{CatalystMlURI: "file://samplecml.json"}
	optInitConext := test.NewOperationInitContext(params, nil)
	opt, err := New(optInitConext)
	assert.Nil(t, err)
	result, err := opt.Eval(inputs)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}
`

func copyFiles(pwd string, tempDir string) error {
	// Copy other files neccessary files.
	for _, val := range files {
		copyFile(filepath.Join(pwd, val), filepath.Join(tempDir, val))
	}

	// Copy test file.
	t, err := template.New("test").Parse(testFile)

	if err != nil {
		return err
	}
	f, err := os.Create(filepath.Join(tempDir, runCMLTest))
	if err != nil {
		return err
	}

	err = t.Execute(f, t)
	if err != nil {
		return err
	}

	return nil
}

func copyFile(srcFile, destFile string) error {
	input, err := ioutil.ReadFile(srcFile)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(destFile, input, 0644)
	if err != nil {
		return err
	}

	return nil
}

func runTest(dir string) error {

	err := execCmd(dir, exec.Command("go", "mod", "init", "runcml"))
	if err != nil {
		return err
	}

	return execCmd(dir, exec.Command("go", "test"))

}

func execCmd(dir string, cmd *exec.Cmd) error {

	cmd.Dir = dir

	out, err := cmd.CombinedOutput()
	fmt.Println(string(out))

	return err
}
