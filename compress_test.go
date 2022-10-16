package compressions

import (
	"strings"
	"testing"
)

type TestCase struct {
	Value    string
	Expected string
}

func TestCompress(t *testing.T) {
	large := new(strings.Builder)
	large.WriteString(strings.Repeat("b", 5000))
	large.WriteString(strings.Repeat("f", 3000000))
	large.WriteString(strings.Repeat("d", 3))
	large.WriteString("W")
	large.WriteString(strings.Repeat("b", 2000))

	testCases := []TestCase{
		{large.String(), "Wb7000d3f3000000"},
		{"", ""},
		{"eddffdd", "d4ef2"},
		{"aaabbbccccc", "a3b3c5"},
		{"aaaabbbcccccaaaa", "a8b3c5"},
		{"zzzzcccUUUuu", "U3c3u2z4"},
		{"ЯЯЯБББдд", "Б3Я3д2"},
		{"ЯЯЯБББдд", "Б3Я3д2"},
		{"шшшШШcchhт", "c2h2Ш2тш3"},
	}

	for _, test := range testCases {
		result := Compress(test.Value)

		if result != test.Expected {
			t.Errorf("Compress(%s) = %s; want %s", test.Value, result, test.Expected)
		}
	}
}
