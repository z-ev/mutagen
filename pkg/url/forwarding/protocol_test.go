package forwarding

import (
	"testing"
)

// TestIsValidProtocol tests that the isValidProtocol function behaves as
// expected for a variety of test cases.
func TestIsValidProtocol(t *testing.T) {
	// Set up test cases.
	testCases := []struct {
		protocol string
		expected bool
	}{
		{"", false},
		{"a", false},
		{"invalid", false},
		{"tcp", true},
		{"tcp4", true},
		{"tcp6", true},
		{"unix", true},
	}

	// Process test cases.
	for _, testCase := range testCases {
		if valid := isValidProtocol(testCase.protocol); valid != testCase.expected {
			t.Error("protocol validity does not match expected:", valid, "!=", testCase.expected)
		}
	}
}
