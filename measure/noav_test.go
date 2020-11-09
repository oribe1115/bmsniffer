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
		PkgPath  string
		Expected int
	}{
		{
			Name:     "simple func",
			FileName: "a/a.go",
			PkgPath:  "a",
			Expected: 1,
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.Name, func(t *testing.T) {
			t.Parallel()

			_, funcDecl, info, err := getFsetAndFuncDeclAndInfo(t, dirPath+test.FileName, test.PkgPath)
			assert.NoError(t, err)

			got := NumberOfAccessedVariables(funcDecl, info)
			assert.Equal(t, test.Expected, got)
		})
	}
}
