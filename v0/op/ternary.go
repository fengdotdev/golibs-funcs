package op

// Ternary returns trueValue if condition is true, otherwise returns falseValue.
// It is a generic function that works with any type T.
// This is useful for simplifying conditional logic in your code.
func Ternary[T any](condition bool, trueValue, falseValue T) T {
	if condition {
		return trueValue
	}
	return falseValue
}

// If is a convenience function that provides a more readable way to use Ternary.
// It takes a condition and two values, returning the first value if the condition is true,
// and the second value if the condition is false.
// This function is particularly useful for inline conditional expressions.
func If(condition bool, trueValue, falseValue any) any {
	// If is an alias for Ternary for better readability.
	return Ternary(condition, trueValue, falseValue)
}
