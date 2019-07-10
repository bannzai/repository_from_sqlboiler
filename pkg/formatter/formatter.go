package formatter

import (
	"fmt"
	"os/exec"
)

type GoFormatter struct {
	FilePath string
}

func (formatter GoFormatter) GoFormat() {
	if err := exec.Command("gofmt", "-w", formatter.FilePath).Run(); err != nil {
		panic(fmt.Errorf("[ERROR] gofmt -w %v", formatter.FilePath))
	}
}

func (formatter GoFormatter) GoImports() {
	if err := exec.Command("goimports", "-w", formatter.FilePath).Run(); err != nil {
		panic(fmt.Errorf("[ERROR] goimports -w %v", formatter.FilePath))
	}
}
