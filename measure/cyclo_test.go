package measure

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCyclomaticComplexity(t *testing.T) {
	t.Parallel()

	dirPath := "cyclo/"

	tests := []struct {
		Name     string
		FileName string
		Expected int
	}{
		{
			Name:     "simple func",
			FileName: "a",
			Expected: 1,
		},
		{
			Name:     "branch with ifelse",
			FileName: "b",
			Expected: 2,
		},
		{
			Name:     "branch with simple for",
			FileName: "c",
			Expected: 2,
		},
		{
			Name:     "branch with switch",
			FileName: "d",
			Expected: 3,
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.Name, func(t *testing.T) {
			t.Parallel()

			ssaFunc := getSSAFunc(t, dirPath+test.FileName)

			got := CyclomaticComplexity(ssaFunc)
			assert.Equal(t, test.Expected, got)
		})
	}
}

func TestCountFlowGraphValues(t *testing.T) {
	t.Parallel()

	dirPath := "cyclo/"

	tests := []struct {
		Name              string
		FileName          string
		ExpectedNodeCount int
		ExpectedEdgeCount int
	}{
		{
			Name:              "simple func",
			FileName:          "a",
			ExpectedNodeCount: 1,
			ExpectedEdgeCount: 0,
		},
		{
			Name:              "branch with ifelse",
			FileName:          "b",
			ExpectedNodeCount: 4,
			ExpectedEdgeCount: 4,
		},
		{
			Name:              "branch with simple for",
			FileName:          "c",
			ExpectedNodeCount: 3,
			ExpectedEdgeCount: 3,
		},
		{
			Name:              "branch with switch",
			FileName:          "d",
			ExpectedNodeCount: 6,
			ExpectedEdgeCount: 7,
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.Name, func(t *testing.T) {
			t.Parallel()

			ssaFunc := getSSAFunc(t, dirPath+test.FileName)

			gotNodeCount, gotEdgeCount := countFlowGraphValues(ssaFunc)
			assert.Equal(t, test.ExpectedNodeCount, gotNodeCount)
			assert.Equal(t, test.ExpectedEdgeCount, gotEdgeCount)
		})
	}
}

func TestCyclomaticComplexityNew(t *testing.T) {
	t.Parallel()

	dirPath := "cyclo/"

	tests := []struct {
		Name     string
		FileName string
		Expected int
	}{
		{
			Name:     "simple func",
			FileName: "a",
			Expected: 1,
		},
		{
			Name:     "branch with ifelse",
			FileName: "b",
			Expected: 2,
		},
		{
			Name:     "branch with simple for",
			FileName: "c",
			Expected: 2,
		},
		{
			Name:     "branch with switch",
			FileName: "d",
			Expected: 3,
		},
		{
			Name:     "branch with closure",
			FileName: "e",
			Expected: 2,
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.Name, func(t *testing.T) {
			t.Parallel()

			funcName, ssaData := getSSADataAndFuncName(t, dirPath+test.FileName)

			got := CyclomaticComplexityNew(funcName, ssaData)
			assert.Equal(t, test.Expected, got)
		})
	}
}

func TestCountFlowGraphValuesNew(t *testing.T) {
	t.Parallel()

	dirPath := "cyclo/"

	tests := []struct {
		Name              string
		FileName          string
		ExpectedNodeCount int
		ExpectedEdgeCount int
	}{
		{
			Name:              "simple func",
			FileName:          "a",
			ExpectedNodeCount: 1,
			ExpectedEdgeCount: 0,
		},
		{
			Name:              "branch with ifelse",
			FileName:          "b",
			ExpectedNodeCount: 4,
			ExpectedEdgeCount: 4,
		},
		{
			Name:              "branch with simple for",
			FileName:          "c",
			ExpectedNodeCount: 3,
			ExpectedEdgeCount: 3,
		},
		{
			Name:              "branch with switch",
			FileName:          "d",
			ExpectedNodeCount: 6,
			ExpectedEdgeCount: 7,
		},
		{
			Name:              "branch with closure",
			FileName:          "e",
			ExpectedNodeCount: 6,
			ExpectedEdgeCount: 6,
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.Name, func(t *testing.T) {
			t.Parallel()

			funcName, ssaData := getSSADataAndFuncName(t, dirPath+test.FileName)

			gotNodeCount, gotEdgeCount := ssaData.countFlowGraphValues(ssaData.getSSAFunc(funcName))
			assert.Equal(t, test.ExpectedNodeCount, gotNodeCount)
			assert.Equal(t, test.ExpectedEdgeCount, gotEdgeCount)
		})
	}
}
