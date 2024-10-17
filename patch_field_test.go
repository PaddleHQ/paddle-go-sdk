package paddle_test

import (
	"encoding/json"
	"testing"

	paddle "github.com/PaddleHQ/paddle-go-sdk/v2"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type InnerStruct struct {
	Test string `json:"test"`
}

type TestStruct struct {
	StringExistsAndIsSet *paddle.PatchField[string]       `json:"string_exists_and_is_set,omitempty"`
	IntExistsAndIsNull   *paddle.PatchField[int]          `json:"int_exists_and_is_null,omitempty"`
	NullableString       *paddle.PatchField[*string]      `json:"nullable_string,omitempty"`
	FieldOmitted         *paddle.PatchField[*InnerStruct] `json:"field_omitted,omitempty"`
	InnerStructPresent   *paddle.PatchField[*InnerStruct] `json:"inner_struct_present,omitempty"`
}

func TestPatchField(t *testing.T) {
	s := TestStruct{
		StringExistsAndIsSet: paddle.NewPatchField("test"),
		IntExistsAndIsNull:   paddle.NewNullPatchField[int](),
		NullableString:       paddle.NewPtrPatchField("test"),
		FieldOmitted:         nil,
		InnerStructPresent:   paddle.NewPatchField(&InnerStruct{Test: "testing"}),
	}

	require.NotNil(t, s.StringExistsAndIsSet.Value())
	assert.Equal(t, "test", *s.StringExistsAndIsSet.Value())

	assert.Nil(t, s.IntExistsAndIsNull.Value())

	assert.Nil(t, s.FieldOmitted)

	require.NotNil(t, s.InnerStructPresent.Value())
	innerVal := *s.InnerStructPresent.Value()
	assert.Equal(t, "testing", innerVal.Test)

	jsonBytes, err := json.Marshal(s)
	require.NoError(t, err)

	assert.JSONEq(t, `{
		"string_exists_and_is_set": "test",
		"int_exists_and_is_null": null,
		"nullable_string": "test",
		"inner_struct_present": {
			"test": "testing"
		}
	}`, string(jsonBytes))
}
