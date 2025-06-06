package datum_test

import (
	"testing"

	"github.com/fengdotdev/golibs-funcs/v0/datum"
	"github.com/fengdotdev/golibs-testing/assert"
)

const (
	foo     = "foo"
	fooHash = "2c26b46b68ffc68ff99b453c1d30413413422d706483bfa0f98a5e886266e7ae"

	bar     = "bar"
	barHash = "fcde2b2edba56bf408601fb721fe9b5c338d10ee429ea04fae5511b68fbf8fb9"
)

func TestValidateSHA256(t *testing.T) {
	assert.Nil(t, datum.ValidateSHA256(foo, fooHash))
	assert.Nil(t, datum.ValidateSHA256(bar, barHash))
	assert.NotNil(t, datum.ValidateSHA256(foo, barHash))
	assert.NotNil(t, datum.ValidateSHA256(bar, fooHash))
}

func TestValidateSHA256Bytes(t *testing.T) {
	assert.Nil(t, datum.ValidateSHA256Bytes([]byte(foo), fooHash))
	assert.Nil(t, datum.ValidateSHA256Bytes([]byte(bar), barHash))
	assert.NotNil(t, datum.ValidateSHA256Bytes([]byte(foo), barHash))
	assert.NotNil(t, datum.ValidateSHA256Bytes([]byte(bar), fooHash))
}

func TestGetSHA256(t *testing.T) {
	assert.Equal(t, fooHash, datum.GetSHA256(foo))
	assert.Equal(t, barHash, datum.GetSHA256(bar))
}

func TestGetSHA256Bytes(t *testing.T) {
	assert.Equal(t, fooHash, datum.GetSHA256Bytes([]byte(foo)))
	assert.Equal(t, barHash, datum.GetSHA256Bytes([]byte(bar)))
}
