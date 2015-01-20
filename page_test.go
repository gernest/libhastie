package libhastie_test

import (
	. "github.com/gernest/libhastie"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Page", func() {
	var (
		source string
		page   *Page
	)
	BeforeEach(func() {
		source = "test/posts/index.md"
	})
	It("Creates a new page", func() {
		page = NewPage(source)

		Expect(page.Title).Should(Equal("Home"))
		Expect(page.Layout).Should(Equal("indexpage"))
	})

})
