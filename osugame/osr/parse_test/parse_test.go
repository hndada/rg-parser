package parse_test

import (
	"fmt"
	"github.com/hndada/rg-parser/osugame/osr"
	"log"
	"testing"
)

func TestParse(t *testing.T) {
	r, err := osr.Parse("test.osr")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s's replay. The score is %d\n", r.PlayerName, r.Score)
	var time int64
	for _, rd := range r.ReplayData {
		time += rd.W
		fmt.Printf("%d: %+v\n", time, rd)
	}
}
