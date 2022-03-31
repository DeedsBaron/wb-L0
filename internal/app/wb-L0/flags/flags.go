package flags

import (
	"flag"
	"github.com/sirupsen/logrus"
	"os"
)

var (
	ConfigPath string
)

func init() {
	ParseFlags()
	logrus.Info("Flags parsed")
}

func ParseFlags() {
	path, err := os.Getwd()
	if err != nil {
		logrus.Fatal("Can't parse config: ", err.Error())
	}
	flag.StringVar(&ConfigPath, "config-path", path+"/config/wb-L0.toml", "path to config file")
	flag.Parse()
	if len(flag.Args()) != 0 {
		logrus.Fatal("Wrong binary parameters, try -help")
	}
}
