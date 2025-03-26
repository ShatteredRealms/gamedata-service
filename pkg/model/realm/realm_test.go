package realm_test

import (
	"github.com/go-faker/faker/v4"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/ShatteredRealms/gamedata-service/pkg/model/realm"
)

var _ = Describe("pkg/model/realm.Realm", func() {
	Describe("IsValidRealm", func() {
		It("should pass for valid realms", func() {
			Expect(realm.IsValid(realm.Human)).To(BeTrue())
			Expect(realm.IsValid(realm.Cyborg)).To(BeTrue())
		})

		It("should fail for invalid realms", func() {
			Expect(realm.IsValid("invalid")).To(BeFalse())
			Expect(realm.IsValid(realm.Realm(faker.Username()))).To(BeFalse())
		})
	})
})
