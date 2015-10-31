package main

import (
	"bytes"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
)

const (
	bmk     = "./bmk"
	inpath  = "./true.in"
	outpath = "./true.out"
)

func main() {

	hPtr := flag.Bool("help", false, "Help")
	vPtr := flag.Bool("version", false, "Version")

	flag.Parse()

	args := make([]string, 2)

	if *hPtr == true {
		args = append(args, "--help")
	}
	if *vPtr == true {
		args = append(args, "--version")
	}

	content, err := ioutil.ReadFile(inpath)
	if err != nil {
		os.Exit(1)
	}

	expect, err := ioutil.ReadFile(bmk)
	if err != nil {
		os.Exit(1)
	}

	for nbyte := range content {
		for nbit := uint32(0); nbit < 8; nbit++ {
			content[nbyte] ^= (1 << nbit)

			if err = ioutil.WriteFile(outpath, content, 0777); err != nil {
				os.Exit(1)
			}

			cmd := exec.Command(outpath, args...)
			content[nbyte] ^= (1 << nbit)
			out, err := cmd.CombinedOutput()
			switch {
			case (bytes.Compare(expect, out) != 0) && (err != nil):
				{
					fmt.Printf("bad-out-bad-exit: %d.%d\n", nbyte, nbit)
				}
			case (bytes.Compare(expect, out) != 0):
				{
					fmt.Printf("bad-out-good-exit: %d.%d\n", nbyte, nbit)
				}
			case err != nil:
				{
					fmt.Printf("good-out-bad-exit: %d.%d\n", nbyte, nbit)
				}
			}
		}
	}
}
