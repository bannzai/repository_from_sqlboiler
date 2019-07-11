package parser

type FileReader interface {
	Read(filePath string) string
}
