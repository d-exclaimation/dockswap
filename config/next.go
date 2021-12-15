//
//  next.go
//  config
//
//  Created by d-exclaimation on 4:39 PM.
//  Copyright © 2021 d-exclaimation. All rights reserved.
//

package config

import (
	"errors"
	"log"
	"sort"
)

// Next swap current with the next in alphabetical order swap file
func (c *Conf) Next() {
	if c == nil {
		log.Fatalln(errors.New("dockswap: No conf found"))
		return
	}
	choices := c.BesideCurrent()
	if len(choices) == 0 {
		log.Fatalln(errors.New("dockswap: No choices left"))
		return
	}

	next := choices[0]
	c.Change(next)
}

// Valid check whether current is a valid location
func (c *Conf) Valid() bool {
	_, ok := c.Swaps[c.Current]
	return ok
}

type pair struct {
	key   string
	value string
}

// BesideCurrent find all swap files beside current
func (c *Conf) BesideCurrent() []pair {
	reduce := func() int {
		if c.Valid() {
			return 1
		}
		return 0
	}()

	// create slice and store keys
	keys := make([]string, 0, len(c.Swaps)-reduce)
	for k := range c.Swaps {
		if k == c.Current {
			continue
		}
		keys = append(keys, k)
	}

	sort.Strings(keys)

	res := make([]pair, len(keys))
	for i, key := range keys {
		res[i] = pair{key: key, value: c.Swaps[key]}
	}

	return res
}
