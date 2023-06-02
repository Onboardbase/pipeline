package pipelines

type Pipeline interface {
	ParseInputFile(file string)
	GetSchema() string
}

type Schema interface{}
