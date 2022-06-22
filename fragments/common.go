package fragments

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func readTemplate(templateFile string) string {
	b, err := ioutil.ReadFile(templateFile)
	if err != nil {
		fmt.Print(err)
	}
	return string(b)
}

func merge(template string, placeholderName string, mergename string) string {
	s1 := strings.Replace(template, placeholderName, mergename, -1)
	return s1
}

func readMergeValues(filename string) []string {
	b, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Print(err)
	}
	filestring := string(b)
	lines := strings.Split(filestring, "\n")

	return lines
}
