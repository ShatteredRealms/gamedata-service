package srv_test

import (
	"github.com/ShatteredRealms/gamedata-service/pkg/srv"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Grpc Server", func() {
	Describe("NewServer", func() {
		It("should exist", func() {
			Expect(func() { srv.NewGameDataServiceServer(nil, nil) }).NotTo(Panic())
		})
	})
})
