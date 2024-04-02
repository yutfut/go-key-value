package conf

import (
	"encoding/json"
	"os"
)

type Conf struct {
	Redis struct {
		Host		string	`json:"Host"`
		Port		int		`json:"Port"`
		Password	string	`json:"Password"`
	} `json:"Redis"`
	Memcached struct {
		Host		string	`json:"Host"`
		Port		int		`json:"Port"`
		Username	string	`json:"Username"`
		Password	string	`json:"Password"`
	} `json:"Memcached"`
	Main struct {
		HTTPPort	int		`json:"HTTPPort"`
		GRPCPort	int		`json:"GRPCPort"`
		Host		string	`json:"Host"`
	} `json:"Main"`
}

func ReadConf(path *string) (*Conf, error) {
	data, err := os.ReadFile(*path)
	if err != nil {
		return &Conf{}, err
	}

	response := &Conf{}

	if err = json.Unmarshal(data, response); err != nil {
		return &Conf{}, err
	}

	return response, nil
}