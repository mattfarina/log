package log_test

import (
	"fmt"

	"github.com/Masterminds/log-go"
	"github.com/Masterminds/log-go/impl/logrus"
)

type Foo struct {
	Logger log.Logger
}

func (f *Foo) DoSomething() {
	f.Logger.Info("Hello Logging")
}

func Example() {

	// Using the default Logger
	log.Info("Hello")
	log.Error("World")

	// Create a logrus logger with default configuration that uses the log
	// interface. Note, logrus can be setup with default settings or setup with
	// custom settings using a second constructor.
	lgrs := logrus.NewStandard()

	// Set logrus as the global logger
	log.Current = lgrs

	// Logrus is now used globally for logging
	log.Warn("Warning through logrus")

	f1 := Foo{
		Logger: lgrs,
	}

	// Logging in DoSomething will use the set logger which is logrus
	f1.DoSomething()

	f2 := Foo{
		// The log package uses the global logger from the standard library log
		// package. A custom standard library logger can be used with the
		// github.com/Masterminds/log-go/impl/std package.
		Logger: log.NewStandard(),
	}

	// Logging in DoSomething will the logger from the standard library
	f2.DoSomething()

	// Need to detect the logger being used? You can check for the type.
	switch log.Current.(type) {
	case *log.StdLogger:
		fmt.Println("The default logger")
	case *logrus.Logger:
		fmt.Printf("Logrus is used for logging")
	default:
		fmt.Printf("Something else that implements the interface")
	}
}
