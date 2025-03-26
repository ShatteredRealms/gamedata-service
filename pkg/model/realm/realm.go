package realm

import "errors"

type Realm string

const ()

var (
	Realms = map[Realm]struct{}{}
	Human  = register("human")
	Cyborg = register("cyborg")

	ErrorInvalidRealm = errors.New("invalid realm")
)

func IsValid(realm Realm) bool {
	_, found := Realms[realm]
	return found
}

func register(realm string) Realm {
	Realms[Realm(realm)] = struct{}{}
	return Realm(realm)
}
