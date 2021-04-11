package go_celsium //nolint:golint,stylecheck

import (
	"errors"
	"fmt"
	"strings"

	"freon/pkg/api"
)

var (
	ErrNotValidParams = errors.New("not valid params")
)

func Format(t *api.Translation, params map[string]string) (string, error) {
	return parseTranslation(t, params)
}

func parseTranslation(t *api.Translation, params map[string]string) (string, error) {
	if len(params) != len(t.IdentifierNamedList) {
		return "", ErrNotValidParams
	}

	orig := t.Text
	for namedParam, value := range params {
		orig = strings.ReplaceAll(orig, fmt.Sprintf("{%s}", namedParam), value)
	}

	return orig, nil
}
