package reader

import (
	"fmt"
	"io/ioutil"
)

type File struct{}

func (File) Read(filePath string) string {
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		panic(fmt.Errorf("Can not read %v, and got error :%v", filePath, err))
	}
	return string(data)
}
