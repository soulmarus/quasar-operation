package message

import (
	"testing"
)

func TestGetMessage(t *testing.T) {
	var tests = []struct {
		messages [][]string
		w        string
	}{
		{[][]string{{"", "este", "es", "un", "mensaje"},
			{"este", "", "un", "mensaje"},
			{"", "este", "es", "", ""}}, "este es un mensaje"},
		{[][]string{{"", "este", "es", "un", "mensaje"},
			{"este", "", "un", "mensaje"},
			{"", "este", "es", "", "secreto"}}, "este es un mensaje secreto"},
		{[][]string{{"", "es", "", "un", "mensaje"},
			{"este", "", "", "", ""},
			{"", "", "", "", "secreto"}}, "este es un mensaje secreto"},
		{[][]string{{"este", "", "", "mensaje", ""},
			{"", "es", "", "", "secreto"},
			{"este", "", "un", "", ""}}, "este es un mensaje secreto"},
		{[][]string{{"este", "mensaje"},
			{"es", ""},
			{"un"}}, "este es un mensaje"},
		{[][]string{}, ""},
	}

	for _, test := range tests {
		got := GetMessage(test.messages...)

		if got != test.w {
			t.Errorf("Got \"%s\" wanted \"%s\" \n", got, test.w)
		}
	}
}
