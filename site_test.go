package libhastie_test

import (
	"os"

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
		site = NewSite("test")
	})
	AfterEach(func() {
		pubklic := "test/publi"
		os.RemoveAll(pubklic)
	})
	It("Should load all the files and directories", func() {
		err = site.Load("test")
		Expect(err).ShouldNot(HaveOccurred())
	})
	It("SHould load pages", func() {
		site.LoadPages()
		Expect(len(site.Pages)).ShouldNot(Equal(0))
	})
	It("Builds the site", func() {
		err = site.Build()
		Expect(err).ShouldNot(HaveOccurred())
	})
	It("Loads files in a specific directory", func() {
		s := LoadFiles("test/posts")
		Expect(s).ShouldNot(BeEmpty())
	})
	Measure("How fast to load by walk", func(b Benchmarker) {
		_ = b.Time("Load", func() {
			s := LoadFiles("test")
			Expect(s).ShouldNot(BeEmpty())
		})
	}, 10)
	Measure("How fast to load without walk", func(b Benchmarker) {
		_ = b.Time("Load", func() {
			site.Load("test")
			Expect(site.Files).ShouldNot(BeEmpty())
		})
	}, 10)
	Measure("How fast can libhastie build", func(b Benchmarker) {
		_ = b.Time("Build", func() {
			err := site.Build()
			Expect(err).ShouldNot(HaveOccurred())
		})
	}, 10)
})
