package server_test

import (
	"net/http"
	"net/http/httptest"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	plaintweet "github.com/suhlig/plaintweet/server"
)

var _ = Describe("/", func() {
	var (
		server   *plaintweet.Server
		request  *http.Request
		response *httptest.ResponseRecorder
	)

	BeforeEach(func() {
		server = plaintweet.NewServer().WithBlurb("Fiat lux")
		request = httptest.NewRequest(http.MethodGet, "/", nil)
		response = httptest.NewRecorder()
	})

	JustBeforeEach(func() {
		server.HandleRoot(response, request)
	})

	It("succeeds", func() {
		Expect(response.Code).To(Equal(200))
	})

	It("responds with the blurb", func() {
		Expect(response.Body).To(ContainSubstring("Fiat lux"))
	})
})
