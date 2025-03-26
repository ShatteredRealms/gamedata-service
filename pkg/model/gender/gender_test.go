package gender_test

import (
	"github.com/go-faker/faker/v4"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/ShatteredRealms/gamedata-service/pkg/model/gender"
)

var _ = Describe("pkg/model/gender.Gender", func() {
	Describe("IsValidGender", func() {
		It("should pass for valid genders", func() {
			Expect(gender.IsValid(gender.Male)).To(BeTrue())
			Expect(gender.IsValid(gender.Female)).To(BeTrue())
		})

		It("should fail for invalid genders", func() {
			Expect(gender.IsValid("invalid")).To(BeFalse())
			Expect(gender.IsValid(gender.Gender(faker.Username()))).To(BeFalse())
		})
	})
})
