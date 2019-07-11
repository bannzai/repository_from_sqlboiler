package writer

import (
	"fmt"
	"io/ioutil"
	"os"
)

type File struct {
	FilePath string
}

type FileWriter interface {
	Write(content string)
}

func (f File) Write(content string) {
	if err := ioutil.WriteFile(f.FilePath, []byte(content), 0644); err != nil {
		panic(err)
	}
	fmt.Fprintf(os.Stdout, "Generated %s. \n", f.FilePath)
}
