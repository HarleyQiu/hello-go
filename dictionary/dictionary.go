package dictionary

import (
	"errors"
)

type Dictionary map[string]string

var ErrNotFound = errors.New("找不到您要找的词")

func (d Dictionary) Search(word string) (string, error) {
	definition, ok := d[word]
	if !ok {
		return "", ErrNotFound
	}
	return definition, nil
}
