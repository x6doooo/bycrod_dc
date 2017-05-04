package conf

import (
    "flag"
    "github.com/BurntSushi/toml"
)

type ServerConf struct {
    Addr string
}

type EnvConf struct {
    Mode  string
    Proxy string
}
type LogConf struct {
    File     string
    Max_line int
    Backups  int
}

type MongoConf struct {
    Addrs     []string
    Database  string
    Username  string
    Password  string
    Mechanism string
    Source    string
}

type XueqiuConf struct {
    Username string
    Password string
}

type BackendConf struct {
    Grabs []string
}

type mainConf struct {
    Env     EnvConf
    Log     LogConf
    Server  ServerConf
    Mongo   MongoConf
    Xueqiu  XueqiuConf
    Backend BackendConf
}

const (
    ENV_MODE_DEV = "dev"
    ENV_MODE_RELEASE = "release"
)

var (
    confFile string
    MainConf = &mainConf{}
    IsDevMode = false
)

func init() {
    flag.StringVar(&confFile, "conf", "/path/file.toml", "config file")
    flag.Parse()

    if confFile == "" {
        panic("need args --conf=/path/file")
    }
    toml.DecodeFile(confFile, MainConf)
    IsDevMode = MainConf.Env.Mode == ENV_MODE_DEV
}
