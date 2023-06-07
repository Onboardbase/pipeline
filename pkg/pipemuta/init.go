package pipemuta

import (
	"fmt"
	"os"
	"path"
	"strings"

	"github.com/Onboardbase/pipemuta/pkg/pipelines"
	"github.com/Onboardbase/pipemuta/pkg/utils"
)

type SUPPORTED_PIPELINE = string

const GITHUB = "github"
const GITLAB = "gitlab"

var SUPPORTED_PIPELINES = []SUPPORTED_PIPELINE{
	GITHUB,
	GITLAB,
}

type PipeMuta struct {
	Configtype SUPPORTED_PIPELINE
	ConfigPath string
	OutputPath string
}

// NewPipeMuta generates a new instance of pipemuta
func NewPipeMuta(inputPath string, outputPath string, configtype string) *PipeMuta {
	pmuta := &PipeMuta{
		// Source:      strings.ToLower(source),
		Configtype: strings.ToLower(configtype),
		ConfigPath: inputPath,
		OutputPath: outputPath,
	}
	pmuta.Validate()
	return pmuta
}

func validate_destination(dest SUPPORTED_PIPELINE) bool {
	return utils.Slice_contains(SUPPORTED_PIPELINES, dest)
}

func validate_file_path(fp string) bool {
	_, err := os.Stat(fp)
	return err == nil
}

func (pmuta *PipeMuta) Validate() {
	// Ensures that the destination is correct
	if !validate_destination(pmuta.Configtype) {
		utils.ErrExit(fmt.Errorf("destination is not one of supported configurations"), 1)
	}

	// ensures that the path is also valid
	if !validate_file_path(pmuta.ConfigPath) {
		utils.ErrExit(fmt.Errorf("pipemuta configuration %s can not be found", pmuta.ConfigPath), 1)
	}
}

func initPipelineType(source SUPPORTED_PIPELINE, pconf *pipelines.PipemutaPipeline) pipelines.PipelineConfig {
	switch source {
	case GITLAB:
		return pipelines.NewGitlabConfig(pconf)
	case GITHUB:
		return pipelines.NewGithubConfig(pconf)
	}
	return nil
}

// Transmute this is the source file that initializes the process
func (pmuta *PipeMuta) Transmute() {
	// convert the input to struct
	pconf, err := pmuta.UnMarshalInputfile()
	if err != nil {
		utils.ErrExit(fmt.Errorf("unable to parse input configuration: %s", err.Error()), 1)
	}

	destConfig := initPipelineType(pmuta.Configtype, pconf)
	if destConfig == nil {
		utils.ErrExit(fmt.Errorf("unable to generate the configuration for %s destination CI/CD", pmuta.Configtype), 1)
	}

	for _, fschema := range destConfig.GenerateSchema() {
		// extract the required struct to the destination config
		outbyte, err := pmuta.MarshalOutput(fschema)
		if err != nil {
			utils.ErrExit(fmt.Errorf("unable to generate the configuration for %s destination CI/CD", pmuta.Configtype), 1)
		}
		err = createOutputPath(pmuta.OutputPath)
		if err != nil {
			utils.ErrExit(fmt.Errorf("unable to create output path: %s", pmuta.OutputPath), 1)
		}
		err = os.WriteFile(pmuta.OutputPath, outbyte, 0644)
		if err != nil {
			utils.ErrExit(fmt.Errorf("can not create file to output path %s", pmuta.OutputPath), 1)
		}
	}
}

func createOutputPath(outputPath string) (err error) {
	outfilename := path.Base(outputPath)
	pathtofile := strings.Replace(outputPath, outfilename, "", -1)
	if !validate_file_path(pathtofile) {
		err = os.MkdirAll(pathtofile, os.ModePerm)
	}
	return err
}
