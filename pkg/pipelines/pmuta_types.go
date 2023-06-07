package pipelines

type PipemutaPipeline struct {
	Name     string `yaml:"name,omitempty"`
	Defaults struct {
		GithubDefaults *GithubDefaults        `yaml:"gh,omitempty"`
		GitlabDefaults *GitlabPipelineDefault `yaml:"gl,omitempty"`
		Env            map[string]interface{} `yaml:"env,omitempty"`
	} `yaml:"defaults,omitempty"`

	Workflows []PipemutaPipelineWorkflow `yaml:"workflows,omitempty"`

	Jobs map[string]struct {
		GithubJobConfig *GithubPipelineJobs `yaml:"gh,omitempty"`
		GitlabJobConfig *GitlabPipelineJob  `yaml:"gl,omitempty"`
	} `yaml:"jobs,omitempty"`
}

type PipemutaPipelineWorkflow struct {
	Name           string `yaml:"name,omitempty"`
	Description    string `yaml:"description,omitempty"`
	GitlabWorkflow struct {
		Except []string                      `yaml:"except,omitempty"`
		Only   interface{}                   `yaml:"only,omitempty"`
		Rules  []*GitlabPipelineWorkflowRule `yaml:"rules,omitempty"`
	} `yaml:"gl,omitempty"`
	GithubWorkflow struct {
		On          interface{}        `yaml:"on,omitempty"`
		Concurrency *GithubConcurrency `yaml:"concurrency,omitempty"`
		Permissions interface{}        `yaml:"permissions,omitempty"`
	} `yaml:"gh,omitempty"`
	Env    map[string]interface{} `yaml:"env,omitempty"`
	Stages []string               `yaml:"stages,omitempty"`
}
