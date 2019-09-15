package conf

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"path/filepath"
)

var conf *Conf

func init() {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		panic(err)
	}
	conf = new(Conf)
	yamlFile, err := ioutil.ReadFile(dir + "/conf.yaml")
	if err != nil {
		panic(err)
	}
	err = yaml.Unmarshal(yamlFile, conf)
	if err != nil {
		panic(err)
	}
}

type Conf struct {
	App   AppConf   `yaml:"app"`
	MySQL MySQLConf `yaml:"mysql"`
}

type AppConf struct {
	Debug   bool   `yaml:"debug"`
	ApiPort string `yaml:"api_port"`
}

type MySQLConf struct {
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	Db       string `yaml:"db"`
}

func GetConfig() *Conf {
	return conf
}

func (m MySQLConf) GetConnectionStr() string {
	// "root:0314@/smart_car?charset=utf8&parseTime=True&loc=Local"
	return m.User + ":" + m.Password + "@/" + m.Db + "?" + "charset=utf8&parseTime=True&loc=Local"
}
