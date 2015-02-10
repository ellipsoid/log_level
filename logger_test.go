package log_level_test

import (
	"github.com/ellipsoid/log_level"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Logger", func() {
	var (
		logger  log_level.Logger
		level   log_level.Level
		printer *MockPrinter
	)

	BeforeEach(func() {
		printer = &MockPrinter{}
	})

	JustBeforeEach(func() {
		logger = log_level.New(level, printer)
	})

	Context("With TRACE Log Level", func() {
		BeforeEach(func() {
			level = log_level.TRACE
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
			level = log_level.DEBUG
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
			level = log_level.INFO
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
			level = log_level.WARN
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
			level = log_level.ERROR
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
