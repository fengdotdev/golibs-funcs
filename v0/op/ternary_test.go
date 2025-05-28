package op_test

import (
	"testing"

	"github.com/fengdotdev/golibs-funcs/v0/op"
)

func TestTerary(t *testing.T) {
	tests := []struct {
		condition  bool
		trueValue  string
		falseValue string
		expected   string
	}{
		{true, "yes", "no", "yes"},
		{false, "yes", "no", "no"},
	}

	for _, test := range tests {
		result := op.Ternary(test.condition, test.trueValue, test.falseValue)
		if result != test.expected {
			t.Errorf("Ternary(%v, %q, %q) = %q; want %q", test.condition, test.trueValue, test.falseValue, result, test.expected)
		}
	}
}


func TestIf(t *testing.T) {
	tests := []struct {
		condition  bool
		trueValue  string
		falseValue string
		expected   string
	}{
		{true, "yes", "no", "yes"},
		{false, "yes", "no", "no"},
	}

	for _, test := range tests {
		result := op.If(test.condition, test.trueValue, test.falseValue)
		if result != test.expected {
			t.Errorf("If(%v, %q, %q) = %q; want %q", test.condition, test.trueValue, test.falseValue, result, test.expected)
		}
	}
}
