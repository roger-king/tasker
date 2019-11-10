package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var expectedTestPayload = TestPayload{
	Name: "Steve Rogers",
}

type TestPayload struct {
	Name string `json:"name"`
}

func TestValidate_ShouldReturnValidStruct(t *testing.T) {
	var tp TestPayload

	assert := assert.New(t)
	args := map[string]interface{}{
		"name": "Steve Rogers",
	}

	err := Validate(args, &tp)

	assert.NoError(err)

	assert.Equal(expectedTestPayload, tp)
}
