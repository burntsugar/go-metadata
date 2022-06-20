package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

var layoutnames string
var recordtypenames string

// Split on the comma
func crunchSplitFunc(data []byte, atEOF bool) (advance int, token []byte, err error) {

	if atEOF && len(data) == 0 {
		return 0, nil, nil
	}

	if i := strings.Index(string(data), ","); i >= 0 {
		return i + 1, data[0:i], nil
	}

	if atEOF {
		return len(data), data, nil
	}

	return
}

func LayoutNames() {
	b, err := ioutil.ReadFile("layoutnames.txt") // just pass the file name
	if err != nil {
		fmt.Print(err)
	}

	//fmt.Println(b) // print the content as 'bytes'

	layoutnames = string(b) // convert content to a 'string'

	//fmt.Println(str) // print the content as a 'string'

	scanner := bufio.NewScanner(strings.NewReader(layoutnames))
	scanner.Split(crunchSplitFunc)

	for scanner.Scan() {
		log.Print(scanner.Text())
	}

}

type names struct {
	layoutname     string
	recordtypename string
}

func splitting(filename string) []string {
	b, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Print(err)
	}

	filestring := string(b)
	return strings.Split(filestring, ",")
}

func ReadFile(names1 string, names2 string) {

	slice1 := splitting(names1)
	slice2 := splitting(names2)

	result := []names{}

	for i, s1 := range slice1 {
		r := names{
			layoutname:     s1,
			recordtypename: slice2[i],
		}
		result = append(result, r)
		fmt.Printf("%+v\n", r)
	}

	for _, rs := range result {
		fmt.Println(rs.layoutname)
	}

}

func main() {
	ReadFile("layoutnames.txt", "recordtypenames.txt")
}
