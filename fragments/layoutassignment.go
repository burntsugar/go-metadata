package fragments

import (
	"fmt"
	"strings"
)

type layoutAttributes struct {
	layoutname     string
	recordtypename string
}

func LayoutAssignments() {
	mergeValues := getLayoutMergeValues("fragments/merge-data/layout-assignment.txt")
	template := readTemplate("fragments/templates/layout-assignment-template.txt")
	placeholderNames := []string{"LAYOUTNAME", "RECORDTYPENAME"}
	layoutAssignments := makeLayoutAssignmentFragments(template, placeholderNames, mergeValues)
	fmt.Print(layoutAssignments)
}

func makeLayoutAssignmentFragments(template string, placeholderNames []string, mergeValues []layoutAttributes) string {
	start := 0
	stop := len(mergeValues) - 1
	var result string
	for i := start; i <= stop; i++ {
		n := mergeValues[i]
		res1 := merge(template, placeholderNames[0], n.layoutname)
		res2 := merge(res1, placeholderNames[1], n.recordtypename)
		res2 += "\n"
		result += res2
	}
	return result
}

func getLayoutMergeValues(mergeValuesFile string) []layoutAttributes {
	mergeValueLines := readMergeValues(mergeValuesFile)
	result := []layoutAttributes{}
	for _, s1 := range mergeValueLines {
		attrs := strings.Split(s1, ",")
		r := layoutAttributes{
			layoutname:     attrs[0],
			recordtypename: attrs[1],
		}
		result = append(result, r)
	}
	return result
}
