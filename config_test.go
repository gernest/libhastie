package libhastie_test

import (
	. "github.com/gernest/libhastie"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Config", func() {
	var (
		err    error
		source string
		config *Config
	)
	BeforeEach(func() {
		source = "test/hastie.json"
		config = new(Config)
	})
	It("SHould parse the configuration file", func() {
		err = config.Load(source)
		Expect(err).ShouldNot(HaveOccurred())
	})
	It("SHould store the values in the config object", func() {
		_ = config.Load(source)
		Expect(config.SourceDir).Should(Equal("posts"))
		Expect(config.PublishDir).Should(Equal("public"))
		Expect(config.LayoutDir).Should(Equal("layouts"))
	})
	It("Should resolve into default config.json file", func() {
		err = config.Load("test")
		Expect(err).ShouldNot(HaveOccurred())
	})
	It("Should load default config file", func() {
		_ = config.Load("test")
		Expect(config.PublishDir).Should(Equal("public"))
	})
})
