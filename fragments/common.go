package fragments

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
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
	// fileName := "fragments/output/" + dateTimeString() + "-data.txt"
	fileName := "fragments/output/output-data.txt"

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
	// fileName := "fragments/output/" + t.Format(layout) + "-data.txt"

	fileName := "fragments/output/output-data.txt"

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

	//removeEmptyLines()
}

func removeEmptyLines() {
	fmt.Println("REMOVE EMPTY LINES!")
	f := openFile()
	defer f.Close()

	//newLines := []string{}
	lines := readMergeValues("fragments/output/output-data.txt")
	for _, line := range lines {

		// str := strings.TrimSpace(line)
		lineLen := len(line)
		fmt.Println(line + " " + strconv.Itoa(lineLen))
		if len(line) > 1 {
			//newLines = append(newLines, line)
			if _, err := fmt.Fprintf(f, "%s\n", line); err != nil {
				log.Fatal(err)
			}
		}
		//s = strings.Replace(lines, "\n\n", "\n", 1)
	}

}

// func removeEmptyLines() {
// 	fmt.Println("REMOVE EMPTY LINES!")
// 	f := openFile()
// 	defer f.Close()

// 	//newLines := []string{}
// 	lines := readMergeValues("fragments/output/output-data.txt")
// 	for _, line := range lines {

// 		str := strings.TrimSpace(line)
// 		fmt.Println(len(str))
// 		if len(str) > 0 {
// 			//newLines = append(newLines, line)
// 			if _, err := fmt.Fprintf(f, "%s\n", str); err != nil {
// 				log.Fatal(err)
// 			}
// 		}
// 		//s = strings.Replace(lines, "\n\n", "\n", 1)
// 	}

// }

// func removeEmptyLines() {
// 	fmt.Println("REMOVE EMPTY LINES!")
// 	f := openFile()
// 	defer f.Close()

// 	//newLines := []string{}
// 	lines := readMergeValues("fragments/output/output-data.txt")
// 	for _, line := range lines {
// 		if len(line) > 0 {
// 			str := strings.TrimSpace(line)
// 			fmt.Println(len(str))
// 			//newLines = append(newLines, line)
// 			if _, err := fmt.Fprintf(f, "%s\n", str); err != nil {
// 				log.Fatal(err)
// 			}
// 		}
// 		//s = strings.Replace(lines, "\n\n", "\n", 1)
// 	}

// }

// func removeEmptyLines() {
// 	newLines := []string{}
// 	lines := readMergeValues("fragments/output/output-data.txt")
// 	for _, line := range lines {
// 		if len(line) > 0 {
// 			newLines = append(newLines, line)
// 		}
// 		//s = strings.Replace(lines, "\n\n", "\n", 1)
// 	}

// 	f := openFile()
// 	defer f.Close()

// 	if _, err := fmt.Fprintf(f, "%s\n", merge4); err != nil {
// 		log.Fatal(err)
// 	}
// }
