package fragments

import (
	"strings"
)

type packageContent struct {
	packageTypeMemb []string
	typeName        string
}

func PackageManifest() {
	mergeValues := getPackageMergeValues("fragments/merge-data/package-metadata.txt")
	template := readTemplate("fragments/templates/package-template.txt")
	typesTemplate := readTemplate("fragments/templates/package-content-template.txt")
	memberTemplate := readTemplate("fragments/templates/package-type-members-template.txt")

	typeNodes := []string{}
	for _, pc := range mergeValues {
		members := []string{}

		for _, m := range pc.packageTypeMemb {
			memb := merge(memberTemplate, "MEMBERNAME", m)
			members = append(members, memb)
		}
		joinMembers := strings.TrimSpace(strings.Join(members, "\n\t"))
		typeNode1 := merge(typesTemplate, "MEMBERS", joinMembers)
		typeNode2 := merge(typeNode1, "NAME", pc.typeName)
		typeNodeEnd := typeNode2 + "\n"
		typeNodes = append(typeNodes, typeNodeEnd)
	}
	fragment := merge(template, "TYPES", strings.TrimSpace(strings.Join(typeNodes, "")))
	saveData(fragment, "fragments/output/package-output-data.txt")
}

func getPackageMergeValues(mergeValuesFile string) []packageContent {

	mergeValueLines := readMergeValues(mergeValuesFile)
	m := make(map[string][]string)

	for _, s1 := range mergeValueLines {
		attrs := strings.Split(s1, ",")
		trimmedAttr0 := strings.TrimSpace(attrs[0])
		trimmedAttr1 := strings.TrimSpace(attrs[1])
		sl, exists := m[trimmedAttr1]
		if !exists {
			sl = []string{}
		}
		sl = append(sl, trimmedAttr0)
		m[trimmedAttr1] = sl
	}

	content := []packageContent{}
	for tn, tName := range m {
		pc := packageContent{
			typeName:        tn,
			packageTypeMemb: tName,
		}
		content = append(content, pc)
	}
	return content
}
