package go_celsium

import (
	"freon/pkg/api"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSprintfSuccessful(t *testing.T) {
	{ // Successful case where len(params) == len(IdentifierNamedList)
		text, err := Format(&api.Translation{
			Text:                "this {number} {name}",
			IdentifierNamedList: []string{"number", "name"},
		}, map[string]string{
			"number": "first",
			"name":   "test",
		})
		assert.Nil(t, err)
		assert.Equal(t, "this first test", text)
	}

	{ // Successful case where len(params) == len(IdentifierNamedList) and repeated named parameters
		text, err := Format(&api.Translation{
			Text:                "this {number} {number} {name}",
			IdentifierNamedList: []string{"number", "name"},
		}, map[string]string{
			"number": "first",
			"name":   "test",
		})
		assert.Nil(t, err)
		assert.Equal(t, "this first first test", text)
	}

	// Successful case where not using named parameters but choose in named list.
	// This bad situation but this case must be validated on backend and frontend not in client app.
	{
		text, err := Format(&api.Translation{
			Text:                "this",
			IdentifierNamedList: []string{"number", "name"},
		}, map[string]string{
			"number": "first",
			"name":   "test",
		})
		assert.Nil(t, err)
		assert.Equal(t, "this", text)
	}

	{ // Successful case with empty params
		text, err := Format(&api.Translation{
			Text: "just text",
		}, nil)
		assert.Nil(t, err)
		assert.Equal(t, "just text", text)
	}
}

func TestSprintfFailure(t *testing.T) {
	{
		_, err := Format(&api.Translation{
			Text:                "this {number} {name}",
			IdentifierNamedList: []string{"number"},
		}, map[string]string{
			"number": "first",
			"name":   "test",
		})
		assert.NotNil(t, err)
		assert.Equal(t, ErrNotValidParams, err)
	}

	{
		_, err := Format(&api.Translation{
			Text:                "this {number} {name}",
			IdentifierNamedList: []string{"number", "name"},
		}, map[string]string{
			"number": "first",
		})
		assert.NotNil(t, err)
		assert.Equal(t, ErrNotValidParams, err)
	}
}
