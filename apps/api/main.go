package main

import (
	"flag"
	"github.com/sirupsen/logrus"
	"github.com/ypapax/comment/api"
	"github.com/ypapax/comment/config"
	"github.com/ypapax/logrus_conf"
	"os"
)

func main() {
	if err := logrus_conf.Files("comment", logrus.TraceLevel); err != nil {
		logrus.Error(err)
		os.Exit(1)
	}
	var confPath string
	flag.StringVar(&confPath, "conf", "conf.yaml", "path to config file")
	flag.Parse()
	f, err := os.Open(confPath)
	if err != nil {
		logrus.Error(err)
		os.Exit(1)
	}
	c, err := config.Parse(f)
	if err != nil {
		logrus.Error(err)
		os.Exit(1)
	}
	if err := api.Serve(*c); err != nil {
		logrus.Error(err)
		os.Exit(1)
	}
}
