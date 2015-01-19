package libhastie_test

import (
	. "github.com/gernest/libhastie"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Frontmatter", func() {
	var (
		file string
		err  error
	)
	BeforeEach(func() {
		file = "test/posts/index.md"
	})
	It("SHould open the file and parse the contents", func() {
		_, _, err = FrontMatter(file)
		Expect(err).ShouldNot(HaveOccurred())
	})
	It("Should Separate front matter from the rest of the file content", func() {
		front, _, _ := FrontMatter(file)
		m := make(map[string]string)
		m["title"] = "Home"
		m["layout"] = "indexpage"
		Expect(front).Should(Equal(m))
	})
	It("Should return the body of the file content without frontmatter", func() {
		_, body, _ := FrontMatter(file)
		Expect(body).ShouldNot(BeEmpty())
	})
})
