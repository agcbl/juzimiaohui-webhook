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

type Juzihudong struct {
	Endpoint string
	Token string
}

type Keyword struct {
	Tick int
}

type Config struct {
	Database *Database
	Lark *LarkBotPath
	Juzihudong *Juzihudong
	Keyword *Keyword
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
