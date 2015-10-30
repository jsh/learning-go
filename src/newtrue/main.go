package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
)

const (
	inpath  = "./true.in"
	outpath = "./true.out"
)

func main() {

	hPtr := flag.Bool("help", false, "Help")
	vPtr := flag.Bool("version", false, "Version")

	flag.Parse()

	args := make([]string, 2)

	if hPtr == true {
		args = append(args, "--help")
	}
	if vPtr == true {
		args = append(args, "--version")
	}

	content, err := ioutil.ReadFile(inpath)
	if err != nil {
		os.Exit(1)
	}


	for byt := range content {
		for bit := uint32(0); bit < 8; bit++ {
			content[byt] ^= (1 << bit)

			if err = ioutil.WriteFile(outpath, content, 0777); err != nil {
				os.Exit(1)
			}

			cmd := exec.Command(outpath, args...)
			content[byt] ^= (1 << bit)
			if err = cmd.Run(); err != nil {
				fmt.Printf("byte = %d, bit = %d\n", byt, bit)
			}
		}
	}
}
