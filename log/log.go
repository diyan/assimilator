package log

import (
	"github.com/Sirupsen/logrus"
	"github.com/diyan/assimilator/conf"
	logrusfmt "github.com/x-cray/logrus-prefixed-formatter"
)

// Init setups logging subsystem according to provided configuration
func Init(config conf.Config) {
	// TODO ForceColors only if codegangsta/gin detected
	logrus.SetFormatter(&logrusfmt.TextFormatter{ShortTimestamp: true, ForceColors: true})
}
