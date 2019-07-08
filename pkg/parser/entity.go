package parser

type Entity struct {
	FileReader
}

func (parser Entity) Parse() {
	content := parser.FileReader.Read()
}
