package pipelines

type PipelineConfig interface {
	// ParseInputFile(file string)
	GenerateSchema() []interface{}
}
