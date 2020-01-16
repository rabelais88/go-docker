package myfile

import (
	"fmt"
	"io/ioutil"
)

type Str string

func (_str Str) SaveString(filename string) error { // returning error is nullable?
	bytes := []byte(_str)
	err := ioutil.WriteFile(filename, bytes, 0666)
	if err != nil {
		fmt.Println(`error!`, err)
	}
	return err
}
