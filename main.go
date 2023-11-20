package logruswrapper

import (
	"strings"
	"syscall"

	formatter "github.com/fabienm/go-logrus-formatters"
	"github.com/sirupsen/logrus"
)

var Logger = logrus.New()

func init() {
	hostname, ok := syscall.Getenv("HOSTNAME")
	if !ok {
		hostname = "DefaultHostName"
	}
	gelfFmt := formatter.NewGelf(hostname)
	Logger.SetFormatter(gelfFmt)

	loglevel, _ := syscall.Getenv("LOGLEVEL")
	loglevel = strings.ToUpper(loglevel)
	switch loglevel {
	case "PANIC":
		Logger.SetLevel(logrus.PanicLevel)
	case "FATAL":
		Logger.SetLevel(logrus.FatalLevel)
	case "ERROR":
		Logger.SetLevel(logrus.ErrorLevel)
	case "WARN":
		Logger.SetLevel(logrus.WarnLevel)
	case "INFO":
		Logger.SetLevel(logrus.InfoLevel)
	case "DEBUG":
		Logger.SetLevel(logrus.DebugLevel)
	case "TRACE":
		Logger.SetLevel(logrus.TraceLevel)

	default:
		Logger.SetLevel(logrus.WarnLevel)
	}

}
