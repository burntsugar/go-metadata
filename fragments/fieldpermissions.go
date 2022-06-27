package fragments

import (
	"strings"
)

type fieldPerms struct {
	editable  string
	fieldName string
	readable  string
}

func FieldPermissions() {
	mergeValues := getFieldPermissionMergeValues("fragments/merge-data/field-permissions.txt")

	template := readTemplate("fragments/templates/field-permissions-template.txt")

	placeholderNames := []string{"EDITABLE", "FIELDNAME", "READABLE"}

	fieldPermissionz := makeFieldPermissionFragments(template, placeholderNames, mergeValues)

	saveData(fieldPermissionz)
}

func makeFieldPermissionFragments(template string, placeholderNames []string, mergeValues []fieldPerms) string {
	start := 0
	stop := len(mergeValues) - 1
	var result string
	for i := start; i <= stop; i++ {
		n := mergeValues[i]
		res1 := merge(template, placeholderNames[0], n.editable)
		res2 := merge(res1, placeholderNames[1], n.fieldName)
		res3 := merge(res2, placeholderNames[2], n.readable)
		res3 += "\n"
		result += res3
	}
	return result
}

func getFieldPermissionMergeValues(mergeValuesFile string) []fieldPerms {
	mergeValueLines := readMergeValues(mergeValuesFile)
	result := []fieldPerms{}

	for _, s1 := range mergeValueLines {

		attrs := strings.Split(s1, ",")
		r := fieldPerms{
			readable:  strings.TrimSpace(attrs[0]),
			editable:  strings.TrimSpace(attrs[1]),
			fieldName: strings.TrimSpace(attrs[2]),
		}
		result = append(result, r)
	}
	return result
}
