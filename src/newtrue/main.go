package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
)

const (
	inpath = "./true.in"
	outpath = "./true.out"
)

func main() {
	var content []byte
	var err error

	if content, err = ioutil.ReadFile(inpath); err != nil {
		os.Exit(1)
	}

	if err = ioutil.WriteFile(outpath, content, 0777); err != nil {
		os.Exit(1)
	}

	cmd := exec.Command(outpath)
	if err = cmd.Run(); err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
}
