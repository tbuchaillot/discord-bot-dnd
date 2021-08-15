package command

import "testing"

func TestGetDiceFaces(t *testing.T) {

	toTest := map[string]struct {
		msg      string
		expected int
	}{
		"valid_20": {
			msg:      "!roll20",
			expected: 20,
		},
		"invalid_with_string": {
			msg:      "!rolltom",
			expected: 0,
		},
		"invalid_empty": {
			msg:      "!roll",
			expected: 0,
		},
	}

	for test, values := range toTest {
		t.Run(test, func(t *testing.T) {
			val := getDiceFaces(values.msg)
			if val != values.expected {
				t.Fail()
			}
		})
	}
}
