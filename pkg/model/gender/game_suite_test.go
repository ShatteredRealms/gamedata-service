package gender_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestGender(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Gender Suite")
}
