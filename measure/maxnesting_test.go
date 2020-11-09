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
