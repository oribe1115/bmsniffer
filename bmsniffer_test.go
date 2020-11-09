package bmsniffer_test

import (
	"testing"

	"github.com/oribe1115/bmsniffer"
	"golang.org/x/tools/go/analysis/analysistest"
)

// TestAnalyzer is a test for Analyzer.
func TestAnalyzer(t *testing.T) {
	testdata := analysistest.TestData()
	analysistest.Run(t, testdata, bmsniffer.Analyzer, "a")
}

