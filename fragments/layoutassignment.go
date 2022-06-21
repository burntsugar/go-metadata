package fragments

import "fmt"

type layoutAttributes struct {
	layoutname     string
	recordtypename string
}

func LayoutAssignments() {
	mergeValues := getLayoutMergeValues("fragments/merge-data/layout-names.txt", "fragments/merge-data/recordtype-names.txt")
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

func getLayoutMergeValues(layoutNamesFile string, recordTypeNamesFile string) []layoutAttributes {
	delimiter := ","
	slice1 := readMergeValues(layoutNamesFile, delimiter)
	slice2 := readMergeValues(recordTypeNamesFile, delimiter)
	result := []layoutAttributes{}
	for i, s1 := range slice1 {
		r := layoutAttributes{
			layoutname:     s1,
			recordtypename: slice2[i],
		}
		result = append(result, r)
	}
	return result
}
