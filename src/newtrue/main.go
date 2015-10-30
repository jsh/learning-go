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

			cmd := exec.Command(outpath)
			content[byt] ^= (1 << bit)
			if err = cmd.Run(); err != nil {
				fmt.Printf("byte = %d, bit = %d\n", byt, bit)
			}
		}
	}
}
