package main

import "testing"

func TestNormal(t *testing.T) {
	type test struct {
		input          string
		expectedOutput string
	}
	tests := []test{
		{input: "1234567890", expectedOutput: "1234567890"},
		{input: "123 456 7891", expectedOutput: "1234567891"},
		{input: "(123) 456 7892", expectedOutput: "1234567892"},
		{input: "(123) 456-7893", expectedOutput: "1234567893"},
		{input: "123-456-78943", expectedOutput: "1234567894"},
		{input: "123-456-7890", expectedOutput: "1234567890"},
		{input: "1234567892", expectedOutput: "1234567892"},
		{input: "(123)456-7892", expectedOutput: "1234567892"},
	}

	for _, c := range tests {
		result := normalize(c.input)
		if result != c.expectedOutput {
			t.Errorf("Error: want: %s got: %s", c.expectedOutput, result)
		}
	}
}
