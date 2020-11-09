package measure

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLineOfCode(t *testing.T) {
	t.Parallel()

	dirPath := "testdata/loc/"

	tests := []struct {
		Name     string
		FileName string
		Expected int
	}{
		{
			Name:     "blank function",
			FileName: "a/a.go",
			Expected: 1,
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.Name, func(t *testing.T) {
			t.Parallel()

			fset, funcDecl, err := getFsetAndFuncDecl(t, dirPath+test.FileName)
			assert.NoError(t, err)

			got := LineOfCode(fset, funcDecl)
			assert.Equal(t, test.Expected, got)
		})
	}

}
