package bigsum

import (
	"testing"
)

func TestSum(t *testing.T) {
	bigNumber := BigNumber{}
	tests := []struct {
		x, y, expected string
	}{
		{"123", "456", "579"},
		{"999", "1", "1000"},
		{"0", "0", "0"},
		{"100", "900", "1000"},
		{"1234", "897", "2131"},
	}

	for _, tt := range tests {
		result := bigNumber.Sum(tt.x, tt.y)
		if result != tt.expected {
			t.Errorf("sum(%s, %s) = %s; expected %s", tt.x, tt.y, result, tt.expected)
		}
	}
}
