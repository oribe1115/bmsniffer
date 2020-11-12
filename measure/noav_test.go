package measure

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNumberOfAccessedVariables(t *testing.T) {
	t.Parallel()

	dirPath := "noav/"

	tests := []struct {
		Name     string
		FileName string
		Expected int
	}{
		{
			Name:     "simple var",
			FileName: "a",
			Expected: 1,
		},
		{
			Name:     "simple recv",
			FileName: "b",
			Expected: 1,
		},
		{
			Name:     "simple pkgVar",
			FileName: "c",
			Expected: 2,
		},
		{
			Name:     "simple attribute",
			FileName: "d",
			Expected: 2,
		},
		{
			Name:     "multipule assign",
			FileName: "e",
			Expected: 4,
		},
		{
			Name:     "redecl with typeswitch",
			FileName: "f",
			Expected: 4,
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.Name, func(t *testing.T) {
			t.Parallel()

			funcDecl, info := getFuncDeclAndTypeInfo(t, dirPath+test.FileName)

			got := NumberOfAccessedVariables(funcDecl, info)
			assert.Equal(t, test.Expected, got)
		})
	}
}
