package gender

import "errors"

type Gender string

var (
	Genders = map[Gender]struct{}{}
	Male    = register("male")
	Female  = register("female")

	ErrorInvalid = errors.New("invalid gender")
)

func IsValid(gender Gender) bool {
	_, found := Genders[gender]
	return found
}

func register(gender string) Gender {
	Genders[Gender(gender)] = struct{}{}
	return Gender(gender)
}
