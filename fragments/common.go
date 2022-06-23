package fragments

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"time"
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

func saveData(dataString string) {
	t := time.Now()
	const layout = "2 Jan 2006 15:04:05"
	fileName := "fragments/output/" + t.Format(layout) + "-data.txt"

	f, err := os.OpenFile(fileName,
		os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()

	sectionDate := t.Format(layout)
	if _, err := f.WriteString(sectionDate + "\n" + dataString + "\n"); err != nil {
		fmt.Println(err)
	}

	fmt.Println("Fragment created!")
}
