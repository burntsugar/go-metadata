package fragments

import (
	"strings"
)

type rtVisibilityAttributes struct {
	isdefault      string
	recordtypename string
	isVisible      string
}

func RecordTypeVisibility() {
	mergeValues := getRecordTypeVisibilityMergeValues("fragments/merge-data/recordtype-visibility.txt")
	template := readTemplate("fragments/templates/recordtype-visibilities-template.txt")
	placeholderNames := []string{"DEFAULT", "RECORDTYPENAME", "VISIBLE"}
	recordTypeVisibilities := makeRecordTypeVisibilityFragments(template, placeholderNames, mergeValues)
	saveData(recordTypeVisibilities)
}

func makeRecordTypeVisibilityFragments(template string, placeholderNames []string, mergeValues []rtVisibilityAttributes) string {
	start := 0
	stop := len(mergeValues) - 1
	var result string
	for i := start; i <= stop; i++ {
		n := mergeValues[i]
		res1 := merge(template, placeholderNames[0], n.isdefault)
		res2 := merge(res1, placeholderNames[1], n.recordtypename)
		res3 := merge(res2, placeholderNames[2], n.isVisible)
		res3 += "\n"
		result += res3
	}
	return result
}

func getRecordTypeVisibilityMergeValues(mergeValuesFile string) []rtVisibilityAttributes {
	mergeValueLines := readMergeValues(mergeValuesFile)
	result := []rtVisibilityAttributes{}
	for _, s1 := range mergeValueLines {

		attrs := strings.Split(s1, ",")
		r := rtVisibilityAttributes{
			isdefault:      attrs[0],
			recordtypename: attrs[1],
			isVisible:      attrs[2],
		}
		result = append(result, r)
	}
	return result
}
