package main

import (
	"fmt"

	kvlib "github.com/Megavolv/okp4kviewlib"
	"github.com/sirupsen/logrus"
	flag "github.com/spf13/pflag"
)

var target int64
var count int64
var loglevel string
var path string

func init() {
	flag.Int64Var(&target, "target", 0, "Initial key number")
	flag.Int64Var(&count, "count", 1, "Number of keys requested")
	flag.StringVar(&loglevel, "level", "debug", "Log level (error|warn|info|debug)")
	flag.StringVar(&path, "path", "db/", "Path to keys")
	flag.Parse()
}

func main() {
	logger := logrus.New()

	switch loglevel {
	case "error":
		logger.SetLevel(logrus.ErrorLevel)
	case "warn":
		logger.SetLevel(logrus.WarnLevel)
	case "info":
		logger.SetLevel(logrus.InfoLevel)

	default:
		logger.SetLevel(logrus.DebugLevel)
	}

	list := kvlib.NewList(path, logger)
	defer list.CloseAll()
	data, _ := list.GetKeys(target, count)
	/*if err != nil {
		fmt.Println(err)
	}*/

	fmt.Println(data)
}
