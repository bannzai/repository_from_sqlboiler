package generator

type GoFormatter interface {
	GoFormat()
	GoImports()
}
