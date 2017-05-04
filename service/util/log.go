package util

import (
    "bycrod_dc/conf"
    "github.com/x6doooo/smog"
)

var (
    Logger smog.LoggerInterface
)

func init() {
    logConf := conf.MainConf.Log;
    Logger = smog.NewLogger(logConf.File, logConf.Max_line, logConf.Backups, conf.IsDevMode)
}