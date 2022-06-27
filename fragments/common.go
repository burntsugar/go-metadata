package fragments

import (
	"fmt"
	"io/ioutil"
	"log"
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

func merge(template string, placeholderName string, mergeValue string) string {
	s1 := strings.Replace(template, placeholderName, mergeValue, -1)
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

func dateTimeString() (dateTimeString string) {
	t := time.Now()
	const l = "2 Jan 2006 15:04:05"
	return t.Format(l)
}

func openFile() (f *os.File) {
	fileName := "fragments/output/" + dateTimeString() + "-data.txt"

	f, err := os.OpenFile(fileName,
		os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}

	return f
}

func saveData(dataString string) {
	t := time.Now()
	const layout = "2 Jan 2006 15:04:05"
	fileName := "fragments/output/" + t.Format(layout) + "-data.txt"

	f, err := os.OpenFile(fileName,
		os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	sectionDate := t.Format(layout)
	if _, err := f.WriteString(sectionDate + "\n" + dataString + "\n"); err != nil {
		fmt.Println(err)
	}

	fmt.Println("Fragment created!")
}
