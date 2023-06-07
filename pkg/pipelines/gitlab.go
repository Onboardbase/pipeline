package pipelines

import (
	"github.com/Onboardbase/pipemuta/pkg/utils"
)

func NewGitlabConfig(pmutaPipeline *PipemutaPipeline) *GitlabConfig {
	return &GitlabConfig{
		PMutaPipeline: pmutaPipeline,
	}
}

type GitlabConfig struct {
	PMutaPipeline *PipemutaPipeline
}

func (ghc *GitlabConfig) GenerateSchema() []interface{} {
	ghpipes := make([]interface{}, 0)
	pmuta := ghc.PMutaPipeline
	workflow := map[string]interface{}{}
	stages := []string{}
	for wfi, wf := range ghc.PMutaPipeline.Workflows {
		for jobkey, job := range pmuta.Jobs {
			if utils.Slice_contains(wf.Stages, jobkey) {
				job.GitlabJobConfig.Stage = wf.Name
			}
			workflow[jobkey] = job.GitlabJobConfig
		}
		workflow["default"] = generateGitlabWorkflowSchema(wfi, pmuta)
		stages = append(stages, wf.Name)
		workflow["workflow"] = &GitlabPipelineWorkflow{
			Name:  pmuta.Name,
			Rules: wf.GitlabWorkflow.Rules,
		}
	}
	workflow["stages"] = stages
	workflow["includes"] = pmuta.Defaults.GitlabDefaults.Include

	ghpipes = append(ghpipes, workflow)
	return ghpipes
}

func generateGitlabWorkflowSchema(wfi int, pmuta *PipemutaPipeline) *GitlabPipeline {
	return &GitlabPipeline{
		BeforeScript:  pmuta.Defaults.GitlabDefaults.BeforeScript,
		AfterScript:   pmuta.Defaults.GitlabDefaults.AfterScript,
		Artifacts:     pmuta.Defaults.GitlabDefaults.Artifacts,
		Cache:         pmuta.Defaults.GitlabDefaults.Cache,
		Hooks:         pmuta.Defaults.GitlabDefaults.Hooks,
		Image:         pmuta.Defaults.GitlabDefaults.Image,
		Timeout:       pmuta.Defaults.GitlabDefaults.Timeout,
		Interruptible: pmuta.Defaults.GitlabDefaults.Interruptible,
		Retry:         pmuta.Defaults.GitlabDefaults.Retry,
		Services:      pmuta.Defaults.GitlabDefaults.Services,
		Tags:          pmuta.Defaults.GitlabDefaults.Tags,
	}
}
