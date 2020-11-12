package measure

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// func TestCyclomaticComplexity(t *testing.T) {
// 	t.Parallel()

// 	dirPath := "cyclo/"

// 	tests := []struct {
// 		Name     string
// 		FileName string
// 		Expected int
// 	}{
// 		{
// 			Name:     "simple func",
// 			FileName: "a",
// 			Expected: 1,
// 		},
// 	}

// 	for _, test := range tests {
// 		test := test
// 		t.Run(test.Name, func(t *testing.T) {
// 			t.Parallel()

// 			_, ssaFunc := getSSAFunc(t, dirPath+test.FileName)

// 			CyclomaticComplexity(ssaFunc)
// 		})
// 	}
// }

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
	}

	for _, test := range tests {
		test := test
		t.Run(test.Name, func(t *testing.T) {
			t.Parallel()

			_, ssaFunc := getSSAFunc(t, dirPath+test.FileName)

			gotNodeCount, gotEdgeCount := countFlowGraphValues(ssaFunc)
			assert.Equal(t, test.ExpectedNodeCount, gotNodeCount)
			assert.Equal(t, test.ExpectedEdgeCount, gotEdgeCount)
		})
	}
}
