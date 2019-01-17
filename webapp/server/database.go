package server

import (
	"io"
	"io/ioutil"
	"os"

	"github.com/jmoiron/sqlx"
	"gopkg.in/yaml.v2"
)

type DBConfigs map[string]*DBConfig

type DBConfig struct {
	Datasource string `yaml:"datasource"`
}

func (c *DBConfig) Open() (*sqlx.DB, error) {
	return sqlx.Open("mysql", c.Datasource)
}

func NewDBConfigsFromFile(path string) (DBConfigs, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	return NewDBConfigs(f)
}

func NewDBConfigs(r io.Reader) (DBConfigs, error) {
	b, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, err
	}
	var configs DBConfigs
	if err = yaml.Unmarshal(b, &configs); err != nil {
		return nil, err
	}
	return configs, nil
}
