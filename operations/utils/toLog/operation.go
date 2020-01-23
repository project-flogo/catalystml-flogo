package toLog

import (
	"fmt"
	"os"
	"strings"

	"github.com/project-flogo/catalystml-flogo/action/operation"
	"github.com/project-flogo/core/data/metadata"
	"github.com/project-flogo/core/support/log"
)

type Operation struct {
	sortByKey bool
	filename  string
	toFile    bool
	logger    log.Logger
}

func New(ctx operation.InitContext) (operation.Operation, error) {
	p := &Params{}

	err := metadata.MapToStruct(ctx.Params(), p, true)

	if err != nil {
		return nil, err
	}

	logger := ctx.Logger()

	toFile := false
	filename := p.ToFilePath
	if "" != p.ToFilePath {
		err = prepareForOutputFile(filename, p.ClearWhileStart, logger)
		if err != nil {
			return nil, err
		}
		toFile = true
	}

	return &Operation{
		filename: filename,
		toFile:   toFile,
		logger:   logger}, nil
}

func (operation *Operation) Eval(inputs map[string]interface{}) (interface{}, error) {

	var err error
	in := &Input{}

	err = in.FromMap(inputs)
	if err != nil {
		return nil, err
	}

	msg := in.Data
	if operation.toFile {
		operation.logger.Debug("File name : ", operation.filename)
		f, err := os.OpenFile(operation.filename, os.O_APPEND|os.O_WRONLY, 0600)
		if err != nil {
			panic(err)
		}
		f.WriteString(fmt.Sprintln(msg))

		defer f.Close()

	} else {
		operation.logger.Info(msg)
	}

	return nil, err
}

func prepareForOutputFile(
	filename string,
	clearWhileStart bool,
	logger log.Logger) error {

	outputFolderPath, _ := splitFilename(filename)

	err := os.MkdirAll(outputFolderPath, os.ModePerm)
	if nil != err {
		logger.Error("Unable to create folder ...")
		return err
	}

	fileExist := true
	_, err = os.Stat(filename)
	if nil != err {
		if os.IsNotExist(err) {
			fileExist = false
		}
	}

	if clearWhileStart {
		if fileExist {
			err = os.Remove(filename)
			if err != nil {
				logger.Error(err)
			}
		}

		_, err = os.Create(filename)
		if nil != err {
			logger.Error("Unable to create file ...")
			return err
		}
	}

	return nil
}

func splitFilename(filename string) (string, string) {
	if "" != filename {
		indexSlash := strings.LastIndex(filename, "/")
		indexBackslash := strings.LastIndex(filename, "\\")
		var index int
		if indexSlash > indexBackslash {
			index = indexSlash
		} else {
			index = indexBackslash
		}
		return filename[:index], filename[index+1:]
	}
	return "", ""
}
