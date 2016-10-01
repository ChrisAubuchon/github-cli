package commands

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func testSelector(v *string) Selector {
	return newSelector(v, "default", []string{"default", "one", "two", "three"})
}

func Test_Selector(t *testing.T) {
	var value string
	s := testSelector(&value)

	assert := assert.New(t)

	// Test that the default is set properly
	assert.Equal("default", value)

	err := s.Set("one")
	assert.Nil(err)
	assert.Equal("one", value)

	err = s.Set("three")
	assert.Nil(err)
	assert.Equal("three", value)

	// Test that setting the same value works
	err = s.Set("three")
	assert.Nil(err)
	assert.Equal("three", value)

	err = s.Set("four")
	assert.NotNil(err)

	err = s.Set("")
	assert.Nil(err)
	assert.Equal("default", value)
}
