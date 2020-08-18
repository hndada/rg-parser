package tools

import "strings"

// generate compressed name as local name; ex) TimingPoint -> tp
func LocalName(structName string) string {
	var name string
	for i, s := range strings.ToLower(structName) {
		if structName[i] != byte(s) {
			name += string(s)
		}
	}
	return name
}