package log_level_test

import (
	"code42.com/level_log"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("LevelLogger", func() {
	var (
		logger  level_log.Logger
		level   level_log.Level
		printer *MockPrinter
	)

	BeforeEach(func() {
		printer = &MockPrinter{}
	})

	JustBeforeEach(func() {
		logger = level_log.New(level, printer)
	})

	Context("With TRACE Log Level", func() {
		BeforeEach(func() {
			level = level_log.TRACE
		})

		It("Logs Trace Calls", func() {
			logger.Trace("Trace Message %v", 1)
			Expect(printer.TimesCalled).To(Equal(1))
		})

		It("Logs Debug Calls", func() {
			logger.Debug("Debug Message %v %v", 1, 2)
			Expect(printer.TimesCalled).To(Equal(1))
		})

		It("Logs Info Calls", func() {
			logger.Info("Info Message %v %v %v", 1, 2, "foo")
			Expect(printer.TimesCalled).To(Equal(1))
		})

		It("Logs Warn Calls", func() {
			logger.Warn("Warn Message %v %v %v", 1, 2, "foo")
			Expect(printer.TimesCalled).To(Equal(1))
		})

		It("Logs Error Calls", func() {
			logger.Error("Error Message %v %v %v", 1, 2, "foo")
			Expect(printer.TimesCalled).To(Equal(1))
		})
	})

	Context("With DEBUG Log Level", func() {
		BeforeEach(func() {
			level = level_log.DEBUG
		})

		It("Does Not Log Trace Calls", func() {
			logger.Trace("Trace Message %v", 1)
			Expect(printer.TimesCalled).To(Equal(0))
		})

		It("Logs Debug Calls", func() {
			logger.Debug("Debug Message %v %v", 1, 2)
			Expect(printer.TimesCalled).To(Equal(1))
		})

		It("Logs Info Calls", func() {
			logger.Info("Info Message %v %v %v", 1, 2, "foo")
			Expect(printer.TimesCalled).To(Equal(1))
		})

		It("Logs Warn Calls", func() {
			logger.Warn("Warn Message %v %v %v", 1, 2, "foo")
			Expect(printer.TimesCalled).To(Equal(1))
		})

		It("Logs Error Calls", func() {
			logger.Error("Error Message %v %v %v", 1, 2, "foo")
			Expect(printer.TimesCalled).To(Equal(1))
		})
	})

	Context("With INFO Log Level", func() {
		BeforeEach(func() {
			level = level_log.INFO
		})

		It("Does Not Log Trace Calls", func() {
			logger.Trace("Trace Message %v", 1)
			Expect(printer.TimesCalled).To(Equal(0))
		})

		It("Does Not Log Debug Calls", func() {
			logger.Debug("Debug Message %v %v", 1, 2)
			Expect(printer.TimesCalled).To(Equal(0))
		})

		It("Logs Info Calls", func() {
			logger.Info("Info Message %v %v %v", 1, 2, "foo")
			Expect(printer.TimesCalled).To(Equal(1))
		})

		It("Logs Warn Calls", func() {
			logger.Warn("Warn Message %v %v %v", 1, 2, "foo")
			Expect(printer.TimesCalled).To(Equal(1))
		})

		It("Logs Error Calls", func() {
			logger.Error("Error Message %v %v %v", 1, 2, "foo")
			Expect(printer.TimesCalled).To(Equal(1))
		})
	})

	Context("With WARN Log Level", func() {
		BeforeEach(func() {
			level = level_log.WARN
		})

		It("Does Not Log Trace Calls", func() {
			logger.Trace("Trace Message %v", 1)
			Expect(printer.TimesCalled).To(Equal(0))
		})

		It("Does Not Log Debug Calls", func() {
			logger.Debug("Debug Message %v %v", 1, 2)
			Expect(printer.TimesCalled).To(Equal(0))
		})

		It("Does Not Log Info Calls", func() {
			logger.Info("Info Message %v %v %v", 1, 2, "foo")
			Expect(printer.TimesCalled).To(Equal(0))
		})

		It("Logs Warn Calls", func() {
			logger.Warn("Warn Message %v %v %v", 1, 2, "foo")
			Expect(printer.TimesCalled).To(Equal(1))
		})

		It("Logs Error Calls", func() {
			logger.Error("Error Message %v %v %v", 1, 2, "foo")
			Expect(printer.TimesCalled).To(Equal(1))
		})
	})

	Context("With ERROR Log Level", func() {
		BeforeEach(func() {
			level = level_log.ERROR
		})

		It("Does Not Log Trace Calls", func() {
			logger.Trace("Trace Message %v", 1)
			Expect(printer.TimesCalled).To(Equal(0))
		})

		It("Does Not Log Debug Calls", func() {
			logger.Debug("Debug Message %v %v", 1, 2)
			Expect(printer.TimesCalled).To(Equal(0))
		})

		It("Does Not Log Info Calls", func() {
			logger.Info("Info Message %v %v %v", 1, 2, "foo")
			Expect(printer.TimesCalled).To(Equal(0))
		})

		It("Does Not Log Warn Calls", func() {
			logger.Warn("Warn Message %v %v %v", 1, 2, "foo")
			Expect(printer.TimesCalled).To(Equal(0))
		})

		It("Logs Error Calls", func() {
			logger.Error("Error Message %v %v %v", 1, 2, "foo")
			Expect(printer.TimesCalled).To(Equal(1))
		})
	})

})

// Mocks

type MockPrinter struct {
	TimesCalled int
}

func (mp *MockPrinter) Printf(format string, v ...interface{}) {
	mp.TimesCalled += 1
	return
}
