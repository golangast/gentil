package gentil

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"runtime"
	"strings"
	"text/template"

	"github.com/Masterminds/sprig"
)

//f is for file, o is for old text, n is for new text
func UpdateText(f string, o string, n string) error {
	input, err := ioutil.ReadFile(f)
	if err != nil {
		fmt.Println(err)
	}

	output := bytes.Replace(input, []byte(o), []byte(n), -1)

	if err = ioutil.WriteFile(f, output, 0666); err != nil {
		fmt.Println(err)
	}

	return nil
}
func FindTextNReturn(p, str string) string {
	// Open file for reading.
	var file, err = os.OpenFile(p, os.O_RDWR, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	toplevel := TrimDot(str)
	property := TrimDotright(str)
	strs := strings.Replace(property, ".", " ", 1)
	// fmt.Println(str)
	// Read file, line by line
	var text = make([]byte, 1024)
	for {
		_, err = file.Read(text)

		if strings.Contains(string(text), toplevel) {
			//is the dot string and split it
			if strings.Contains(string(text), strs) {
				return string(text)
			}
		}
		// Break if finally arrived at end of file
		if err == io.EOF {
			break
		}

		// Break if error occured
		if err != nil && err != io.EOF {
			fmt.Println(err)

		}
	}

	// fmt.Println("Reading from file.")
	fmt.Println(string(text))

	return ""
}
func TrimDot(s string) string {
	if idx := strings.Index(s, "."); idx != -1 {
		return s[:idx]
	}
	return s
}
func TrimDotright(s string) string {
	if idx := strings.Index(s, "."); idx != -1 {
		return s[idx:]
	}
	return s
}

//used for terminal
func Shellout(command string) (string, string, error) {
	var stdout bytes.Buffer
	var stderr bytes.Buffer

	var cmd *exec.Cmd
	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/C", command)
	} else {
		cmd = exec.Command("bash", "-c", command)
	}

	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	err := cmd.Run()
	return stdout.String(), stderr.String(), err
}

// kill a process
func killProcessByName(procname string) int {
	kill := exec.Command("taskkill", "/im", procname, "/T", "/F")
	err := kill.Run()
	if err != nil {
		return -1
	}
	return 0
}
func Runapplinuxmac() (string, string, error) {
	err, out, errout := Shellout(`go run *.go`)
	return out, err, errout
}
func Runwindows() (string, string, error) {
	err, out, errout := Shellout(`go run .`)
	return out, err, errout
}
func Reload() (string, string, error) {
	err, out, errout := Shellout("pwd && cd app && go mod tidy && go mod vendor && go install && go build")
	return out, err, errout

}

// write any template to file
func Writetemplate(temp string, f *os.File, d map[string]string) error {
	functionMap := sprig.TxtFuncMap()
	dbmb := template.Must(template.New("queue").Funcs(functionMap).Parse(temp))
	err := dbmb.Execute(f, d)
	if err != nil {
		return err
	}
	return nil
}

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
