package asserty

import "reflect"

// True checks if the value is true. If not, it panics.
func True(value bool) {
	if !value {
		panic("Assertion failed")
	}
}

// False checks if the value is false. If not, it panics.
func False(value bool) {
	if value {
		panic("Assertion failed")
	}
}

// Equal checks if the value is equal to the expected value. If not, it panics.
func Equal(value interface{}, expected interface{}) {
	if !reflect.DeepEqual(value, expected) {
		panic("Assertion failed")
	}
}

// Nil checks if the value is nil. If not, it panics.
func Nil(value interface{}) {
	if value != nil {
		panic("Assertion failed")
	}
}

// Err expects an error. If the error is nil, it panics.
func Err(err error) {
	if err == nil {
		panic("Assertion failed")
	}
}

// NoError checks if the error is nil. If not, it panics with the error message.
func NoError(err error) {
	if err != nil {
		panic("Assertion NoError failed" + err.Error())
	}
}
