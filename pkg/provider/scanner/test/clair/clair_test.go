package clair

import (
	"flag"
	"github.com/q8s-io/heimdall/pkg/provider/scanner"
	"log"
	"testing"
)

func TestClairResult(t *testing.T) {
	flag.Parse()
	imageFullName := "known:0.9.9"
	result, _ := scanner.ClairScan(imageFullName, 5)
	log.Print(result)
}
