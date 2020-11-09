package measure

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNumberOfAccessedVariables(t *testing.T) {
	t.Parallel()

	dirPath := "testdata/noav/"

	tests := []struct {
		Name     string
		FileName string
		Expected int
	}{
		{
			Name:     "simple var",
			FileName: "a/a.go",
			Expected: 1,
		},
		{
			Name:     "simple recv",
			FileName: "b/b.go",
			Expected: 1,
		},
		{
			Name:     "simple pkgVar",
			FileName: "c/c.go",
			Expected: 2,
		},
		{
			Name:     "simple attribute",
			FileName: "d/d.go",
			Expected: 2,
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.Name, func(t *testing.T) {
			t.Parallel()

			_, funcDecl, info, err := getFsetAndFuncDeclAndInfo(t, dirPath+test.FileName)
			assert.NoError(t, err)

			got := NumberOfAccessedVariables(funcDecl, info)
			assert.Equal(t, test.Expected, got)
		})
	}
}
