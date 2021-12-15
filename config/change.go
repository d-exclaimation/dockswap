//
//  change.go
//  config
//
//  Created by d-exclaimation on 5:42 PM.
//  Copyright © 2021 d-exclaimation. All rights reserved.
//

package config

import (
	"encoding/json"
	"errors"
	"github.com/d-exclaimation/dockswap/cli"
	"log"
	"path/filepath"
)

const dockerfile = "./Dockerfile"

func (c *Conf) Jump(next string) {
	path, ok := c.Swaps[next]
	if !ok {
		log.Fatalln(errors.New("dockswap: Cannot find swap file"))
	}
	c.Change(pair{next, path})
}

// Change the config json file
func (c *Conf) Change(next pair) {
	log.Printf("Swapping to %s\n", next)

	// Join path
	nextpath := filepath.Join("./.dockswap", next.value)

	nextValue, err := cli.Read(nextpath)
	if err != nil {
		log.Fatalln(err)
		return
	}

	if err := cli.Write(dockerfile, nextValue); err != nil {
		log.Fatalln(err)
		return
	}

	bytes, err := json.MarshalIndent(
		Conf{Current: next.key, Swaps: c.Swaps},
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
}
