package main

import (
	"io"
	"os/exec"
	"strings"
)

func getShortnames() []string {
	var shortnames = make([]string, len(emojiToCodepoints))
	var i int
	for key := range emojiToCodepoints {
		shortnames[i] = key
		i++
	}
	return shortnames
}

func dmenuChoose(options []string) string {
	dmenuList := strings.Join(options, "\n")
	cmd := exec.Command("dmenu", "-i")
	dmenuStdin, _ := cmd.StdinPipe()
	io.WriteString(dmenuStdin, dmenuList)
	dmenuStdin.Close()
	choiceRaw, _ := cmd.Output()
	choice := strings.TrimSuffix(string(choiceRaw), "\n")
	return choice
}

func typeCodepoints(codepoints []string) {
	exec.Command("xdotool", append([]string{"key"}, codepoints...)...).Run()
}

func main() {
	// choose an emoji using dmenu
	shortnames := getShortnames()
	shortname := dmenuChoose(shortnames)

	// get codepoints for chosen emoji
	codepoints := emojiToCodepoints[shortname]

	// type codepoints using xdotool
	typeCodepoints(codepoints)
}
