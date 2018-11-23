package main

import (
	"encoding/json"
	"io"
	"os"
	"os/exec"
	"strings"
)

func chooseShortname(list []string) string {
	dmenuList := strings.Join(list, "\n")
	c := exec.Command("dmenu", "-i")
	stdin, _ := c.StdinPipe()
	io.WriteString(stdin, dmenuList)
	stdin.Close()
	b, err := c.Output()
	if err != nil {
		os.Exit(0)
	}
	c.Run()
	shortname := strings.TrimSpace(string(b))
	return shortname
}

func typeCodepoints(codepoints string) {
	cps := strings.Split(strings.Replace(codepoints, "+", "", -1), " ")
	for _, cp := range cps {
		exec.Command("xdotool", "key", cp).Run()
	}
}

func main() {
	// get emoji shortnames and their respective codepoints
	var emojiToCodepoints map[string]string
	json.Unmarshal([]byte(emojis), &emojiToCodepoints)

	// choose an emoji using dmenu
	var shortnames []string
	for k := range emojiToCodepoints {
		shortnames = append(shortnames, k)
	}
	shortname := chooseShortname(shortnames)

	// get codepoints for chosen emoji
	codepoints := emojiToCodepoints[shortname]

	// type codepoints using xdotool
	typeCodepoints(codepoints)
}
