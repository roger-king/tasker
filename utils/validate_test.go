package utils

import (
	"testing"

	"github.com/stretchr/testify/require"
)

var expectedTestPayload = TestPayload{
	Name: "Steve Rogers",
}

type TestPayload struct {
	Name string `json:"name" validate:"required"`
}

func TestValidate_ShouldReturnValidStruct(t *testing.T) {
	var tp TestPayload

	args := map[string]interface{}{
		"name": "Steve Rogers",
	}

	err := Validate(args, &tp)

	require.NoError(t, err)
	require.Equal(t, expectedTestPayload, tp)
}

func TestValidate_ShouldReturnError(t *testing.T) {
	var tp TestPayload

	args := map[string]interface{}{
		"wrongKey": "Steve Rogers",
	}

	err := Validate(args, &tp)
	require.Error(t, err, "validation error in args and payload")
}
