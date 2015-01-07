package log

import (
	"github.com/Sirupsen/logrus"
	"github.com/Unknwon/goconfig"
	"strings"
)

var Logger = logrus.New()

func init() {

	c, err := goconfig.LoadConfigFile("conf/conf.ini")
	if err != nil {
		panic(err)
	}

	level, err := c.GetValue("Server", "log_level")
	if err != nil {
		level = "Info"
	}

	levelMap := make(map[string]logrus.Level)
	levelMap["Debug"] = logrus.DebugLevel
	levelMap["Info"] = logrus.InfoLevel
	levelMap["Warn"] = logrus.WarnLevel
	levelMap["Error"] = logrus.ErrorLevel
	levelMap["Fatal"] = logrus.FatalLevel
	levelMap["Panic"] = logrus.PanicLevel

	logLevel, ok := levelMap[strings.Title(level)]
	if !ok {
		logLevel = levelMap["Info"]
	}

	Logger.Level = logLevel
}
