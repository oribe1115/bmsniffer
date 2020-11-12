package measure

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLineOfCode(t *testing.T) {
	t.Parallel()

	dirPath := "loc/"

	tests := []struct {
		Name     string
		FileName string
		Expected int
	}{
		{
			Name:     "blank function",
			FileName: "a",
			Expected: 1,
		},
		{
			Name:     "simple function",
			FileName: "b",
			Expected: 9,
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.Name, func(t *testing.T) {
			t.Parallel()

			fset, funcDecl := getFsetAndFuncDecl(t, dirPath+test.FileName)

			got := LineOfCode(fset, funcDecl)
			assert.Equal(t, test.Expected, got)
		})
	}

}
