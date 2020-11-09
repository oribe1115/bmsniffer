package measure

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaxNestingLevel(t *testing.T) {
	t.Parallel()

	dirPath := "testdata/maxnesting/"

	tests := []struct {
		Name     string
		FileName string
		Expected int
	}{
		{
			Name:     "blank function",
			FileName: "a/a.go",
			Expected: 0,
		},
		{
			Name:     "nesting with ifelse",
			FileName: "b/b.go",
			Expected: 3,
		},
		{
			Name:     "nesting with switch",
			FileName: "c/c.go",
			Expected: 3,
		},
		{
			Name:     "nesting with typeswitch",
			FileName: "d/d.go",
			Expected: 3,
		},
		{
			Name:     "nesting with select",
			FileName: "e/e.go",
			Expected: 2,
		},
		{
			Name:     "nesting with for",
			FileName: "f/f.go",
			Expected: 4,
		},
		{
			Name:     "nesting with funclit",
			FileName: "g/g.go",
			Expected: 4,
		},
		{
			Name:     "nesting with all",
			FileName: "z/z.go",
			Expected: 7,
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.Name, func(t *testing.T) {
			t.Parallel()

			_, funcDecl, err := getFsetAndFuncDecl(t, dirPath+test.FileName)
			assert.NoError(t, err)

			got := MaxNestingLevel(funcDecl)
			assert.Equal(t, test.Expected, got)
		})
	}
}
