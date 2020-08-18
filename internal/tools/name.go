package tools

import "strings"

// LocalName generates compressed name as local name
// For example, LocalName returns 'tp' when 'TimingPoint' is input.
func LocalName(structName string) string {
	var name string
	for i, s := range strings.ToLower(structName) {
		if structName[i] != byte(s) {
			name += string(s)
		}
	}
	return name
}