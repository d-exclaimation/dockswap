//
//  init.go
//  config
//
//  Created by d-exclaimation on 6:17 PM.
//  Copyright © 2021 d-exclaimation. All rights reserved.
//

package config

import (
	"encoding/json"
	"github.com/d-exclaimation/dockswap/cli"
	"log"
	"os"
)

func (c *Conf) Init() {
	err := os.Mkdir("./.dockswap", 0755)
	if err != nil {
		log.Fatal(err)
	}

	bytes, err := json.MarshalIndent(
		Conf{Current: c.Current, Swaps: c.Swaps},
		"",
		"  ",
	)

	if err != nil {
		log.Fatalln(err)
		return
	}

	if err := cli.Write("./.dockswap/conf.json", string(bytes)); err != nil {
		log.Fatalln(err)
		return
	}

	build := "# Edit this file"

	if err := cli.Write("./.dockswap/build.swap", build); err != nil {
		log.Fatalln(err)
		return
	}
}
