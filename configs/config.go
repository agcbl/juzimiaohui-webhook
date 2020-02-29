package configs

import (
	"github.com/BurntSushi/toml"
	"io/ioutil"
)

type Database struct {
	Name string
	User string
	Password string
	Host string
}

type LarkBotPath struct {
	Path string
}

type Group struct {
	Groups []string
}

type Word struct {
	Words []string
}

type Config struct {
	Database *Database
	Lark *LarkBotPath
	Group *Group
	Word *Word
}

var DefaultConfig Config

func NewConfig(path string) {
	var data []byte
	var err error
	if data, err = ioutil.ReadFile(path); err != nil {
		panic(err)
	}

	if _, err = toml.Decode(string(data), &DefaultConfig); err != nil {
		panic(err)
	}
}
