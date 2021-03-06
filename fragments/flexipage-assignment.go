package fragments

import (
	"fmt"
	"strings"
)

type flexipageAttributes struct {
	flexipageName  string
	objectApiName  string
	recordtypename string
	profileName    string
}

func FlexipageAssignments() {

	// newplaceholderNames := []string{}
	// if hasRecordType {
	// 	newplaceholderNames := []string{"FLEXIPAGENAME", "OBJECTAPINAME", "RECORDTYPENAME", "PROFILENAME"}
	// }

	mergeValues := getFlexipageMergeValues("fragments/merge-data/flexipage-assignment.txt")
	template := readTemplate("fragments/templates/flexipage-assignment-template.txt")
	placeholderNames := []string{"FLEXIPAGENAME", "OBJECTAPINAME", "RECORDTYPENAME", "PROFILENAME"}
	makeFlexipageAssignmentFragments(template, placeholderNames, mergeValues)
}

func makeFlexipageAssignmentFragments(template string, placeholderNames []string, mergeValues []flexipageAttributes) {
	// f := openFile("fragments/output/flexipage-output-data.txt")
	// defer f.Close()

	rtElementText := "    <recordType>RECORDTYPENAME</recordType>"
	fragment := []string{}
	start := 0
	stop := len(mergeValues) - 1
	for i := start; i <= stop; i++ {
		n := mergeValues[i]
		merge1 := merge(template, placeholderNames[0], n.flexipageName)
		merge2 := merge(merge1, placeholderNames[1], n.objectApiName)

		mergeText := n.recordtypename
		placeholderfragment := placeholderNames[2]
		if n.recordtypename == "NONE" {
			mergeText = ""
			placeholderfragment = rtElementText
		}
		//fmt.Println(mergeText + placeholderfragment)
		// sSize := len(merge(merge2, placeholderfragment, mergeText))
		// fmt.Println(sSize)
		// if sSize == 0 {
		// 	fmt.Println(sSize)
		// }
		merge3 := merge(merge2, placeholderfragment, mergeText)
		sSize := len(merge3)
		fmt.Println(sSize)
		merge4 := merge(merge3, placeholderNames[3], n.profileName)

		// if _, err := fmt.Fprintf(f, "%s\n", merge4); err != nil {
		// 	log.Fatal(err)
		// }
		fragment = append(fragment, merge4)

	}
	//removeEmptyLines("fragments/output/flexipage-output-data.txt")
	//finalFrag := removeEmptyLinesFromSlice(fragment)
	// saveData(strings.Join(finalFrag, ""), "fragments/output/flexipage-output-data.txt")
	saveData(strings.Join(fragment, ""), "fragments/output/flexipage-output-data.txt")
	// fmt.Println(fragment)
	// fmt.Println("Fragment created!")
}

// func makeFlexipageAssignmentFragments(template string, placeholderNames []string, mergeValues []flexipageAttributes) {
// 	f := openFile("fragments/output/flexipage-output-data.txt")
// 	defer f.Close()

// 	rtElementText := "    <recordType>RECORDTYPENAME</recordType>"

// 	start := 0
// 	stop := len(mergeValues) - 1
// 	for i := start; i <= stop; i++ {
// 		n := mergeValues[i]
// 		merge1 := merge(template, placeholderNames[0], n.flexipageName)
// 		merge2 := merge(merge1, placeholderNames[1], n.objectApiName)

// 		mergeText := n.recordtypename
// 		placeholderfragment := placeholderNames[2]
// 		if n.recordtypename == "NONE" {
// 			mergeText = ""
// 			placeholderfragment = rtElementText
// 		}
// 		merge3 := merge(merge2, placeholderfragment, mergeText)
// 		merge4 := merge(merge3, placeholderNames[3], n.profileName)

// 		if _, err := fmt.Fprintf(f, "%s\n", merge4); err != nil {
// 			log.Fatal(err)
// 		}

// 	}
// 	removeEmptyLines("fragments/output/flexipage-output-data.txt")
// 	fmt.Println("Fragment created!")
// }

// func makeFlexipageAssignmentFragments(template string, placeholderNames []string, mergeValues []flexipageAttributes) {
// 	f := openFile()
// 	defer f.Close()

// 	rtElementText := "    <recordType>RECORDTYPENAME</recordType>"

// 	start := 0
// 	stop := len(mergeValues) - 1
// 	for i := start; i <= stop; i++ {
// 		n := mergeValues[i]
// 		merge1 := merge(template, placeholderNames[0], n.flexipageName)
// 		merge2 := merge(merge1, placeholderNames[1], n.objectApiName)
// 		merge3 := merge(merge2, placeholderNames[2], n.recordtypename)
// 		merge4 := merge(merge3, placeholderNames[3], n.profileName)

// 		// if _, err := fmt.Fprintf(f, "%s\n%s\n", dateTimeString(), merge4); err != nil {
// 		// 	log.Fatal(err)
// 		// }

// 		if _, err := fmt.Fprintf(f, "%s\n", merge4); err != nil {
// 			log.Fatal(err)
// 		}
// 	}
// 	fmt.Println("Fragment created!")
// }

func getFlexipageMergeValues(mergeValuesFile string) []flexipageAttributes {
	mergeValueLines := readMergeValues(mergeValuesFile)
	result := []flexipageAttributes{}
	for _, s1 := range mergeValueLines {
		attrs := strings.Split(s1, ",")
		r := flexipageAttributes{
			flexipageName:  strings.TrimSpace(attrs[0]),
			objectApiName:  strings.TrimSpace(attrs[1]),
			recordtypename: strings.TrimSpace(attrs[2]),
			profileName:    strings.TrimSpace(attrs[3]),
		}
		result = append(result, r)
	}
	return result
}
