package gitlab

type GitlabPipeline struct {
	Schema string
}

func (glab *GitlabPipeline) ParseInputFile(filepath string) {}

func (glab *GitlabPipeline) GetSchema() string {
	return get_schema()
}
