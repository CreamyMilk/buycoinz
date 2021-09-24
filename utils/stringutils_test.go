package utils

import "testing"

func TestToSentenceCase(t *testing.T) {
	//Test Table
	tt := []struct {
		name   string
		input  string
		output string
	}{
		{"Empty String", "", ""},
		{"Single Letter LowerCase", "a", "A"},
		{"Single Letter Upper Case", "B", "B"},
		{"Single Word", "milk", "Milk"},
		{"Long sentence Word", "this sentence is super long", "This Sentence Is Super Long"},
		{"Special Characters", "*.<><><", "*.<><><"},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			result := ToSentenceCase(tc.input)
			if result != tc.output {
				t.Errorf("Expected %s to be [%s] but got %s", tc.input, result, tc.output)
			}
		})
	}
}
