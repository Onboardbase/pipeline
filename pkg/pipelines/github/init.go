package github

type GithubPipeline struct {
	Schema string
}

func (glab *GithubPipeline) ParseInputFile(filepath string) {}

func (glab *GithubPipeline) GetSchema() string {
	return get_schema()
}
