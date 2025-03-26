package profession

import (
	"errors"
	"strings"
)

type Profession string

const (
	Necromancer Profession = "Necromancer"
)

var (
	ErrorInvalid = errors.New("invalid profession")
)

func FromString(s string) (Profession, error) {
	s = strings.ToLower(s)
	switch s {
	case "necromancer":
		return Necromancer, nil
	default:
		return "", ErrorInvalid
	}
}
