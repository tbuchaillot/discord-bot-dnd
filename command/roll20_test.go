package command

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetDiceFaces(t *testing.T) {

	toTest := map[string]struct {
		msg           string
		expectedValid bool
		expectedFaces int
		expectedDices int
	}{
		"valid_20": {
			msg:           "!roll20",
			expectedValid: true,
			expectedFaces: 20,
			expectedDices: 1,
		},
		"valid_20_2": {
			msg:           "!roll20 2",
			expectedValid: true,
			expectedFaces: 20,
			expectedDices: 2,
		},
		"valid_20_2_rand": {
			msg:           "!roll20 2 asdasd",
			expectedValid: true,
			expectedFaces: 20,
			expectedDices: 2,
		},
		"invalid_with_string": {
			msg:           "!rolltom",
			expectedValid: false,
			expectedFaces: 0,
			expectedDices: 1,
		},
		"invalid_empty": {
			msg:           "!roll",
			expectedValid: false,
			expectedFaces: 0,
			expectedDices: 1,
		},
	}

	for test, values := range toTest {
		t.Run(test, func(t *testing.T) {
			valid, dices, faces := getDiceInfo(values.msg)

			assert.Equal(t, values.expectedValid, valid)
			assert.Equal(t, values.expectedDices, dices)
			assert.Equal(t, values.expectedFaces, faces)
		})
	}
}
