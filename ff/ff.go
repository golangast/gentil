package gentil

import (
	"errors"
	"fmt"
	"os"
)

// make any folder
func Makefolder(p string) error {
	if err := os.MkdirAll(p, os.ModeSticky|os.ModePerm); err != nil {
		fmt.Println("~~~~could not create"+p, err)
		return err
	}
	return nil
}

// make any file
func Makefile(p string) (*os.File, error) {
	file, err := os.Create(p)
	if err != nil {
		return file, err
	}
	return file, nil
}

//make folder and file (sometimes needed to make sure go knows where they are or if they are generated yet)
func Filefolder(folder, file string) (*os.File, error) {
	var ct *os.File
	if _, err := os.Stat(folder + file); errors.Is(err, os.ErrNotExist) {
		Makefolder(folder + file)
		ct, err := Makefile(folder + file + "/create" + file + ".go")
		return ct, err
	}
	return ct, nil
}
