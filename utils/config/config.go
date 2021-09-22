package config

import (
	"github.com/tomazk/envcfg"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"path/filepath"
)

// declare a type that will hold your env variables
type Cfg struct {
	LOGLEVEL        string `yaml:"LOGLEVEL"`
	FAH_CONN_STRING string `yaml:"FAH_CONN_STRING"`
	PORT            string `yaml:"PORT"`
	PARALLEL_CNT    int    `yaml:"PARALLEL_CNT"`
	FIBER_CNT       int    `yaml:"FIBER_CNT"`
}

func InitConfig(conf *Cfg) error {

	err := envcfg.Unmarshal(conf)
	if err != nil {
		log.Panic("не удалось прочитать переменные окружения:", err)
	}

	//TODO выставить значения по умолчанию
	//conf.FAH_CONN_STRING = "apps/apps@172.31.203.249:1529/fah"
	//conf.PORT = ":80"
	//conf.LOGLEVEL = "DEBUG"
	//conf.PARALLEL_CNT = 10
	//conf.LOG_FILE = "/Users/vitaliy/go/src/dev/otr/fahLoadTest/fahTest/fahTest.log"
	filename, _ := filepath.Abs("fahTest/config.yml")
	yamlFile, err := ioutil.ReadFile(filename)
	err = yaml.Unmarshal(yamlFile, conf)
	//TODO Логика проверки переменных окружения
	return nil
}