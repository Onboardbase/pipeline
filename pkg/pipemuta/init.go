package pipemuta

import (
	"fmt"
	"os"
	"strings"

	"github.com/Onboardbase/pipemuta/pkg/pipelines"
	"github.com/Onboardbase/pipemuta/pkg/pipelines/github"
	"github.com/Onboardbase/pipemuta/pkg/pipelines/gitlab"
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
	Source              SUPPORTED_PIPELINE
	Destination         SUPPORTED_PIPELINE
	FilePath            string
	SourcePipeline      pipelines.Pipeline
	DestinationPipeline pipelines.Pipeline
}

// NewPipeMuta generates a new instance of pipemuta
func NewPipeMuta(source string, dest string, filePath string) *PipeMuta {
	pmuta := &PipeMuta{
		Source:      strings.ToLower(source),
		Destination: strings.ToLower(dest),
		FilePath:    filePath,
	}
	pmuta.Validate()
	return pmuta
}

func validate_source(source SUPPORTED_PIPELINE) bool {
	return utils.Slice_contains(SUPPORTED_PIPELINES, source)
}

func validate_destination(dest SUPPORTED_PIPELINE) bool {
	return utils.Slice_contains(SUPPORTED_PIPELINES, dest)
}

func validate_file_path(fp string) bool {
	_, err := os.Stat(fp)
	return err == nil
}

func (pmuta *PipeMuta) Validate() {
	// Ensures that the source is valid
	if !validate_source(pmuta.Source) {
		utils.ErrExit(fmt.Errorf("source is not one of supported configurations"), 1)
	}

	// Ensures that the destination is correct
	if !validate_destination(pmuta.Destination) {
		utils.ErrExit(fmt.Errorf("destination is not one of supported configurations"), 1)
	}

	// ensures that both source can not be transformed to the same destination, e.g github to github - does not make sense at all
	if pmuta.Destination == pmuta.Source {
		utils.ErrExit(fmt.Errorf("destination and source can not be the same "), 1)
	}

	// ensures that the path is also valid
	if !validate_file_path(pmuta.FilePath) {
		utils.ErrExit(fmt.Errorf("the file containing %s CI configuration can not be found", pmuta.Source), 1)
	}
}

func initPipelineType(source SUPPORTED_PIPELINE) pipelines.Pipeline {
	switch source {
	case GITLAB:
		return &gitlab.GitlabPipeline{}
	case GITHUB:
		return &github.GithubPipeline{}
	}
	return nil
}

// Transmute this is the source file that initializes the process
func (pmuta *PipeMuta) Transmute() {
	// validate the source configuration with a preset yaml configuration

	source := initPipelineType(pmuta.Source)
	if source == nil {
		utils.ErrExit(fmt.Errorf("unable to initialize the configuration for %s source pipeline", pmuta.Source), 1)
	}
	destination := initPipelineType(pmuta.Destination)
	if destination == nil {
		utils.ErrExit(fmt.Errorf("unable to initialize the configuration for %s destination pipeline", pmuta.Source), 1)
	}
	pmuta.DestinationPipeline = destination
	pmuta.SourcePipeline = source
}
