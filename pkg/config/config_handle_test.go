package config_test

import (
	"testing"

	"github.com/Netsocs-Team/driver.sdk_go/pkg/config"
	"github.com/stretchr/testify/assert"
)

type TestStruct struct {
	Field1 string `json:"field1"`
	Field2 int    `json:"field2"`
}

func TestUnmarshalAny(t *testing.T) {
	// Test case 1: Unmarshal to a struct
	bytes := []byte(`{"field1": "test", "field2": 123}`)
	var result *TestStruct
	var err error
	result, err = config.UnmarshalAny[TestStruct](bytes)
	assert.Nil(t, err)
	assert.Equal(t, "test", result.Field1)
	assert.Equal(t, 123, result.Field2)

	// Test case 2: Unmarshal to an int
	bytes = []byte(`123`)
	var resultInt *int
	resultInt, err = config.UnmarshalAny[int](bytes)
	assert.Nil(t, err)
	assert.Equal(t, 123, *resultInt)

	// Test case 3: Unmarshal to a string
	bytes = []byte(`"test"`)
	var resultString *string
	resultString, err = config.UnmarshalAny[string](bytes)
	assert.Nil(t, err)
	assert.Equal(t, "test", *resultString)

	// Test case 4: Unmarshal invalid JSON
	bytes = []byte(`{"field1": "test", "field2":}`)
	result, err = config.UnmarshalAny[TestStruct](bytes)
	assert.NotNil(t, err)
	assert.Nil(t, result)
}
