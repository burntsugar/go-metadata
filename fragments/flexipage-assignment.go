package fragments

import (
	"strings"
)

type flexipageAttributes struct {
	flexipageName  string
	objectApiName  string
	recordtypename string
	profileName    string
}

func FlexipageAssignments() {
	mergeValues := getFlexipageMergeValues("fragments/merge-data/flexipage-assignment.txt")
	template := readTemplate("fragments/templates/flexipage-assignment-template.txt")
	placeholderNames := []string{"FLEXIPAGENAME", "OBJECTAPINAME", "RECORDTYPENAME", "PROFILENAME"}
	flexipageAssignments := makeFlexipageAssignmentFragments(template, placeholderNames, mergeValues)
	//fmt.Print(flexipageAssignments)

	saveData(flexipageAssignments)
}

func makeFlexipageAssignmentFragments(template string, placeholderNames []string, mergeValues []flexipageAttributes) string {
	start := 0
	stop := len(mergeValues) - 1
	var result string
	for i := start; i <= stop; i++ {
		n := mergeValues[i]
		res1 := merge(template, placeholderNames[0], n.flexipageName)
		res2 := merge(res1, placeholderNames[1], n.objectApiName)
		res3 := merge(res2, placeholderNames[2], n.recordtypename)
		res4 := merge(res3, placeholderNames[3], n.profileName)
		res4 += "\n"
		result += res4
	}
	return result
}

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
