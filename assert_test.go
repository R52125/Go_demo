package Go_demo

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSomething1(t *testing.T) {
	// assert equality
	assert.Equal(t, 123, 123, "they should be equal")

	// assert inequality
	assert.NotEqual(t, 123, 456, "they should not be equal")

	// assert for nil (good for errors)
	assert.Nil(t, nil)

	type User struct {
		userName string
		password string
	}
	str := &User{
		userName: "R52125",
		password: "Hey, boy!!!",
	}
	// assert for not nil (good when you expect something)
	if assert.NotNil(t, str) {

		// now we know that object isn't nil, we are safe to make
		// further assertions without causing any errors
		assert.Equal(t, "R52125", str.userName)
	}
}

// if you assert many times, use the below:
func TestSomething2(t *testing.T) {
	assert := assert.New(t)
	// assert equality
	assert.Equal(123, 123, "they should be equal")

	// assert inequality
	assert.NotEqual(123, 456, "they should not be equal")

	// assert for nil (good for errors)
	assert.Nil(nil)

	type User struct {
		userName string
		password string
	}
	str := &User{
		userName: "R52125",
		password: "Hey, boy!!!",
	}
	// assert for not nil (good when you expect something)
	if assert.NotNil(str) {

		// now we know that object isn't nil, we are safe to make
		// further assertions without causing any errors
		assert.Equal("R52125", str.userName)
	}
}
