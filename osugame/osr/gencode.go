// +build ignore

package main

import (
	"bufio"
	"fmt"
	"github.com/hndada/rg-parser/internal/tools"
	"log"
	"os"
	"strings"
)

type fieldInfo struct {
	name      string
	fieldType string
}

// ScanStructs supposes gofmt was already proceeded at given file
// TODO: consider making it local library function (i think no)
func ScanStructs(path string) ([]string, map[string][]fieldInfo) {
	f, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	var structName string
	var infos []fieldInfo
	structs := make([]string, 0)
	m := make(map[string][]fieldInfo)

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		vs := strings.Fields(scanner.Text())
		switch {
		case len(vs) == 0 || vs[0] == "//":
			continue
		case vs[0] == "type" && len(vs) > 2 && vs[2] == "struct":
			structName = vs[1]
			structs = append(structs, structName)
			infos = make([]fieldInfo, 0)
		case structName != "" && len(vs) >= 2:
			info := fieldInfo{name: vs[0], fieldType: vs[1]}
			infos = append(infos, info)
		case vs[0] == "}":
			m[structName] = infos
			structName = ""
		}
	}
	return structs, m
}

func printSetFormatValue(localName, returnName string, f fieldInfo) {
	switch f.fieldType {
	case "string":
		fmt.Printf(`if %s.%s, err = readString(r); err != nil {
			return %s, err
		}
`, localName, f.name, returnName)
	case "[]Action":
		fmt.Printf(`if %s.%s, err = parseReplayData(r); err != nil {
			return %s, err
		}
`, localName, f.name, returnName)
	default:
		fmt.Printf(`if err = binary.Read(r, binary.LittleEndian, &%s.%s); err != nil {
		return %s, err
	}
`, localName, f.name, returnName)
	}
}

func main() {
	_, m := ScanStructs("format.go")
	{
		localName := tools.LocalName("Format")
		returnName := "&" + localName
		for _, f := range m["Format"] {
			printSetFormatValue(localName, returnName, f)
		}
	}
}
