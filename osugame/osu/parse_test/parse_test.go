package parse_test

import (
	"fmt"
	"github.com/hndada/gosu/rg-parser/osugame/osu"
	"log"
	"testing"
)

var tests = []string{"testo.osu", "testt.osu", "testc.osu", "testm.osu"}

func TestParse(t *testing.T) {
	for _, s := range tests {
		o, err := osu.Parse(s)
		if err != nil {
			log.Fatal(err)
		}
		// fmt.Printf("%+v", o)
		fmt.Printf("%+v\n%+v\n", o.General, o.Metadata)
	}
}

func BenchmarkParse(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, err := osu.Parse(tests[0])
		if err != nil {
			log.Fatal(err)
		}
	}
}

// BenchmarkParallel

