package flags

import (
	"flag"
	"log"
	"os"
)

var (
	ConfigPath string
)

func init() {
	ParseFlags()
}

func ParseFlags() {
	path, err := os.Getwd()
	if err != nil {
		log.Fatal("Can't parse config: ", err.Error())
	}
	flag.StringVar(&ConfigPath, "config-path", path+"/config/wb-L0.toml", "path to config file")
	flag.Parse()
	if len(flag.Args()) != 0 {
		log.Fatal("Wrong binary parameters, try -help")
	}
}
