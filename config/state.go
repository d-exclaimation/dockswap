//
//  state.go
//  config
//
//  Created by d-exclaimation on 4:14 PM.
//  Copyright © 2021 d-exclaimation. All rights reserved.
//

package config

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

type Conf struct {
	Current string           `json:"current,omitempty"`
	Swaps   map[string]string `json:"swaps,omitempty"`
}

func defaultConf() *Conf {
	curr := "build"
	return &Conf{
		Current: curr,
		Swaps: map[string]string{
			"build": "./build.swap",
		},
	}
}

func Get() *Conf {
	jsonFile, err := os.Open("./.dockswap/conf.json")
	if err != nil {
		return defaultConf()
	}
	defer func() {
		_ = jsonFile.Close()
	}()

	bytes, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		return defaultConf()
	}

	var conf Conf

	if err := json.Unmarshal(bytes, &conf); err != nil {
		return defaultConf()
	}

	return &conf
}

