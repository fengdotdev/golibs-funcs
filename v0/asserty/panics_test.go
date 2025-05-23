package asserty_test

import (
	"errors"
	"fmt"
	"testing"

	"github.com/fengdotdev/golibs-funcs/v0/asserty"
	"github.com/fengdotdev/golibs-testing/assert"
)

// testing objects
type obj struct {
	Name string
}

type obj2 struct {
	name string
}

func NewOnobj2(name string) *obj2 {
	return &obj2{
		name: name,
	}
}

func alwaysError() error {
	return errors.New("always error")
}
func alwaysNilErr() error {
	return nil
}

func TestTrue(t *testing.T) {

	t.Run("asserty.True against false", func(t *testing.T) {
		didPanic := false

		func() {

			defer func() {
				if r := recover(); r != nil {
					didPanic = true
				}
			}()

			asserty.True(false) // should panic
		}()

		assert.TrueWithMessage(t, didPanic, "should panic")
	})

	t.Run("asserty.True aganins True", func(t *testing.T) {
		didPanic := false

		func() {
			defer func() {
				if r := recover(); r != nil {
					didPanic = true
				}
			}()

			asserty.True(true) // should not panic
		}()

		assert.FalseWithMessage(t, didPanic, "should not panic")
	})

}

func TestFalse(t *testing.T) {

	t.Run("asserty.False against true", func(t *testing.T) {
		didPanic := false

		func() {
			defer func() {
				if r := recover(); r != nil {
					didPanic = true
				}
			}()

			asserty.False(true) // should panic
		}()

		assert.TrueWithMessage(t, didPanic, "should panic")
	})

	t.Run("asserty.False aganins False", func(t *testing.T) {
		didPanic := false

		func() {
			defer func() {
				if r := recover(); r != nil {
					didPanic = true
				}
			}()

			asserty.False(false) // should not panic
		}()

		assert.FalseWithMessage(t, didPanic, "should not panic")
	})
}

func TestEqual(t *testing.T) {

	t.Run("asserty.Equal against diferent 1", func(t *testing.T) {
		didPanic := false

		func() {
			defer func() {
				if r := recover(); r != nil {
					didPanic = true
				}
			}()

			asserty.Equal(obj{Name: "test"}, obj{Name: "test2"}) // should panic
		}()

		assert.TrueWithMessage(t, didPanic, "Should panic when values are different")
	})

	t.Run("asserty.Equal against diferent 2", func(t *testing.T) {
		didPanic := false

		func() {
			defer func() {
				if r := recover(); r != nil {
					didPanic = true
				}
			}()

			asserty.Equal(2, 3) // should panic
		}()

		assert.TrueWithMessage(t, didPanic, "Should panic when values are different")
	})

	t.Run("asserty.Equal against equals 1", func(t *testing.T) {
		didPanic := false

		func() {
			defer func() {
				if r := recover(); r != nil {
					didPanic = true
				}
			}()

			asserty.Equal(obj{Name: "test"}, obj{Name: "test"}) // should not panic
		}()

		assert.FalseWithMessage(t, didPanic, "Should not panic when values are equal")
	})

	t.Run("asserty.Equal against equals 2", func(t *testing.T) {
		didPanic := false

		func() {
			defer func() {
				if r := recover(); r != nil {
					didPanic = true
				}
			}()

			asserty.Equal(1, 1) // should not panic
		}()

		assert.FalseWithMessage(t, didPanic, "Should not panic when values are equal")
	})

}

func TestNil(t *testing.T) {
	t.Run("asserty.Nil against not nil", func(t *testing.T) {
		didPanic := false

		func() {
			defer func() {
				if r := recover(); r != nil {
					didPanic = true
				}
			}()

			asserty.Nil(obj{Name: "test"}) // should panic
		}()

		assert.TrueWithMessage(t, didPanic, "Should panic when value is not nil")
	})

	t.Run("asserty.Nil against nil", func(t *testing.T) {
		didPanic := false

		func() {
			defer func() {
				if r := recover(); r != nil {
					didPanic = true
				}
			}()

			asserty.Nil(nil) // should not panic
		}()

		assert.FalseWithMessage(t, didPanic, "Should not panic when value is nil")
	})

}

func TestErr(t *testing.T) {
	t.Run("asserty.Err against nil", func(t *testing.T) {
		didPanic := false

		func() {
			defer func() {
				if r := recover(); r != nil {
					didPanic = true
				}
			}()

			asserty.Err(alwaysNilErr()) // should panic
		}()

		assert.TrueWithMessage(t, didPanic, "Should panic when value is nil")
	})

	t.Run("asserty.Err against some error", func(t *testing.T) {
		didPanic := false

		func() {
			defer func() {
				if r := recover(); r != nil {
					didPanic = true
				}
			}()

			asserty.Err(alwaysError()) // should not panic
		}()

		assert.FalseWithMessage(t, didPanic, "Should not panic when value is not nil")
	})

	t.Run("asserty.Err against fmt err", func(t *testing.T) {

		didPanic := false

		func() {
			defer func() {
				if r := recover(); r != nil {
					didPanic = true
				}
			}()

			asserty.Err(fmt.Errorf("some err")) // should not panic
		}()

		assert.FalseWithMessage(t, didPanic, "Should not panic when value is not nil")
	})
}

func TestNoError(t *testing.T) {

	t.Run("asserty.NoError against some error", func(t *testing.T) {
		didPanic := false

		func() {
			defer func() {
				if r := recover(); r != nil {
					didPanic = true
				}
			}()

			asserty.NoError(alwaysError()) // should panic
		}()

		assert.TrueWithMessage(t, didPanic, "Should panic when value is not nil")
	})

	t.Run("asserty.NoError against nil", func(t *testing.T) {
		didPanic := false

		func() {
			defer func() {
				if r := recover(); r != nil {
					didPanic = true
				}
			}()

			asserty.NoError(alwaysNilErr()) // should not panic
		}()

		assert.FalseWithMessage(t, didPanic, "Should not panic when value is nil")
	})

	t.Run("asserty.NoError against fmt err", func(t *testing.T) {
		didPanic := false

		func() {
			defer func() {
				if r := recover(); r != nil {
					didPanic = true
				}
			}()

			asserty.NoError(fmt.Errorf("some err")) // should panic
		}()

		assert.TrueWithMessage(t, didPanic, "Should panic when value is not nil")
	})

}
