package measure

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaxNestingLevel(t *testing.T) {
	t.Parallel()

	dirPath := "maxnesting/"

	tests := []struct {
		Name     string
		FileName string
		Expected int
	}{
		{
			Name:     "blank function",
			FileName: "a",
			Expected: 0,
		},
		{
			Name:     "nesting with ifelse",
			FileName: "b",
			Expected: 3,
		},
		{
			Name:     "nesting with switch",
			FileName: "c",
			Expected: 3,
		},
		{
			Name:     "nesting with typeswitch",
			FileName: "d",
			Expected: 3,
		},
		{
			Name:     "nesting with select",
			FileName: "e",
			Expected: 2,
		},
		{
			Name:     "nesting with for",
			FileName: "f",
			Expected: 4,
		},
		{
			Name:     "nesting with funclit",
			FileName: "g",
			Expected: 4,
		},
		{
			Name:     "nesting with all",
			FileName: "z",
			Expected: 7,
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.Name, func(t *testing.T) {
			t.Parallel()

			_, funcDecl := getFsetAndFuncDecl(t, dirPath+test.FileName)

			got := MaxNestingLevel(funcDecl)
			assert.Equal(t, test.Expected, got)
		})
	}
}
