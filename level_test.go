package log_level_test

import (
	"github.com/ellipsoid/log_level"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("GetLevel", func() {
	var (
		levelName string
		err       error
		level     log_level.Level
	)

	JustBeforeEach(func() {
		level, err = log_level.GetLevel(levelName)
	})

	Context("With 'TRACE'", func() {
		BeforeEach(func() {
			levelName = "TRACE"
		})

		It("Returns the TRACE Level object", func() {
			Expect(level).To(Equal(log_level.TRACE))
			Expect(err).To(BeNil())
		})
	})

	Context("With 'DEBUG'", func() {
		BeforeEach(func() {
			levelName = "DEBUG"
		})

		It("Returns the DEBUG Level object", func() {
			Expect(level).To(Equal(log_level.DEBUG))
			Expect(err).To(BeNil())
		})
	})

	Context("With 'INFO'", func() {
		BeforeEach(func() {
			levelName = "INFO"
		})

		It("Returns the INFO Level object", func() {
			Expect(level).To(Equal(log_level.INFO))
			Expect(err).To(BeNil())
		})
	})

	Context("With 'WARN'", func() {
		BeforeEach(func() {
			levelName = "WARN"
		})

		It("Returns the WARN Level object", func() {
			Expect(level).To(Equal(log_level.WARN))
			Expect(err).To(BeNil())
		})
	})

	Context("With 'ERROR'", func() {
		BeforeEach(func() {
			levelName = "ERROR"
		})

		It("Returns the ERROR Level object", func() {
			Expect(level).To(Equal(log_level.ERROR))
			Expect(err).To(BeNil())
		})
	})

	Context("With an invalid Level name", func() {
		BeforeEach(func() {
			levelName = "NOT A REAL LEVEL"
		})

		It("Returns an error", func() {
			Expect(level).To(Equal(log_level.INVALID))
			Expect(err).ToNot(BeNil())
		})
	})
})
