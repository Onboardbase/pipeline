package pipelines

func NewGithubConfig(pmutaPipeline *PipemutaPipeline) *GithubConfig {
	return &GithubConfig{
		PMutaPipeline: pmutaPipeline,
	}
}

type GithubConfig struct {
	PMutaPipeline *PipemutaPipeline
}

func (ghc *GithubConfig) GenerateSchema() []interface{} {
	ghpipes := make([]interface{}, 0)
	for wfi := range ghc.PMutaPipeline.Workflows {
		ghpipes = append(ghpipes, generateWorkflowSchema(wfi, ghc.PMutaPipeline))
	}
	return ghpipes
}

func generateWorkflowSchema(wfi int, pmuta *PipemutaPipeline) *GithubPipeline {
	wf := pmuta.Workflows[wfi]
	wfjobs := make(map[string]*GithubPipelineJobs, 0)
	for _, stage := range wf.Stages {
		wfjobs[stage] = pmuta.Jobs[stage].GithubJobConfig
	}

	return &GithubPipeline{
		Name:        pmuta.Name,
		RunName:     wf.Name,
		Concurrency: wf.GithubWorkflow.Concurrency,
		On:          wf.GithubWorkflow.On,
		Env:         wf.Env,
		Defaults: &GithubDefaults{
			Run: pmuta.Defaults.GithubDefaults.Run,
		},
		Permissions: wf.GithubWorkflow.Permissions,
		Jobs:        wfjobs,
	}
}
