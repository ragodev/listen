package main

import (
	"fmt"
	"os/exec"
	"strings"
)

var (
	wav = []byte{82, 73, 70, 70, 36, 125, 0, 0, 87, 65, 86, 69, 102, 109, 116, 32, 16, 0, 0, 0, 1, 0, 1, 0, 128, 62, 0, 0, 0, 125, 0, 0, 2, 0, 16, 0, 100, 97, 116, 97, 0, 125, 0, 0}
)

const (
	key = "okay computer"
)

func containsPhrase(b []int16) string {
	c := exec.Command("pocketsphinx_continuous",
		"-dict sphinx/2772.dic -lm sphinx/2772.lm -infile /dev/stdin")
	w, err := c.StdinPipe()
	if err != nil {
		panic(err)
	}
	w.Write(wav)
	for _, x := range b {
		w.Write([]byte{uint8(x & 0xff), uint8(x >> 8)})
	}
	out, err := c.Output()
	fmt.Println(string(out))
	if strings.Contains(string(out), key) {
		return key
	}
	return ""
}
