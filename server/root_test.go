package server_test

import (
	"net/http"
	"net/http/httptest"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/suhlig/plaintweet/plaintweet"
	pt_server "github.com/suhlig/plaintweet/server"
)

var _ = Describe("root handler", func() {
	var (
		server   *pt_server.Server
		request  *http.Request
		response *httptest.ResponseRecorder
	)

	BeforeEach(func() {
		repository := plaintweet.NewMockRepository(42, "Hello World", "mock")
		server = pt_server.NewServer(repository).WithBlurb("Fiat lux")
		response = httptest.NewRecorder()
	})

	JustBeforeEach(func() {
		server.HandleRoot(response, request)
	})

	Context("GET /", func() {
		BeforeEach(func() {
			request = httptest.NewRequest(http.MethodGet, "/", nil)
		})

		It("succeeds", func() {
			Expect(response.Code).To(Equal(200))
		})

		It("responds with the blurb", func() {
			Expect(response.Body).To(ContainSubstring("Fiat lux"))
		})
	})

	Context("GET /42", func() {
		BeforeEach(func() {
			request = httptest.NewRequest(http.MethodGet, "/42", nil)
		})

		It("succeeds", func() {
			Expect(response.Code).To(Equal(200))
		})

		It("responds with the plain tweet", func() {
			Expect(response.Body).To(ContainSubstring(`"Hello World" -- @mock on #twitter`))
		})
	})

	Context("GET /11", func() {
		BeforeEach(func() {
			request = httptest.NewRequest(http.MethodGet, "/11", nil)
		})

		It("succeeds", func() {
			Expect(response.Code).To(Equal(404))
		})

		It("responds with the error", func() {
			Expect(response.Body).To(ContainSubstring("Error: no status found with that ID"))
		})
	})
})
