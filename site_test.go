package libhastie_test

import (
	. "github.com/gernest/libhastie"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Site", func() {
	var (
		site *SiteStruct
		err  error
	)
	BeforeEach(func() {
		site = &SiteStruct{}
	})
	It("Loads", func() {
		err = site.Load("test")
		Expect(err).ShouldNot(HaveOccurred())
	})
	It("Works", func() {
		_ = site.Load("test")
		Expect(len(site.Directories)).Should(Equal(3))
		Expect(len(site.Files)).Should(Equal(14))
	})
})
