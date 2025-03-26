package srv_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("GameDataContext", func() {
	Describe("Close", func() {
		It("should exist", func() {
			Expect(func() { srvCtx.Close() }).NotTo(Panic())
		})
	})
})
