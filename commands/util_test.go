package commands 

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_stringInSlice(t *testing.T) {
	assert := assert.New(t)

	rval := stringInSlice("test", []string{"test", "foo", "bar", "baaz"})
	assert.Equal(rval, true)
	rval = stringInSlice("baaz", []string{"test", "foo", "bar", "baaz"})
	assert.Equal(rval, true)
	rval = stringInSlice("foo", []string{"test", "foo", "bar", "baaz"})
	assert.Equal(rval, true)
	rval = stringInSlice("quux", []string{"test", "foo", "bar", "baaz"})
	assert.Equal(rval, false)

}
