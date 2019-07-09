package formatter

import (
	"os/exec"
)

type GoFormatter struct {
	FilePath string
}

func (formatter GoFormatter) GoFormat() {
	if err := exec.Command("gofmt", "-w", formatter.FilePath).Run(); err != nil {
		panic(err)
	}
}

func (formatter GoFormatter) GoImports() {
	if err := exec.Command("goimports", "-w", formatter.FilePath).Run(); err != nil {
		panic(err)
	}
}
