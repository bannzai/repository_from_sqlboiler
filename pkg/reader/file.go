package reader

import (
	"fmt"
	"io/ioutil"
)

type File struct {
	FilePath string
}

func (f File) Read() string {
	data, err := ioutil.ReadFile(f.FilePath)
	if err != nil {
		panic(fmt.Errorf("Can not read %v, and got error :%v", f.FilePath, err))
	}
	return string(data)
}
