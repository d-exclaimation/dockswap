//
//  args.go
//  cli
//
//  Created by d-exclaimation on 4:02 PM.
//  Copyright © 2021 d-exclaimation. All rights reserved.
//

package cli

import (
	"errors"
	"os"
)

// Args get the argument passed when running this executable
func Args() []string {
	return os.Args[1:]
}

type Operations string

const (
	Initialize Operations = "init"
	Next       Operations = "next"
	Jump       Operations = "jump"
	Error      Operations = "error"
)

func Operation() Operations {
	if len(Args()) == 0 {
		if _, err := os.Stat("./.dockswap/conf.json"); errors.Is(err, os.ErrNotExist) {
			return Initialize
		}
		return Next
	}

	op := Args()[0]

	switch op {
	case "init", "initialize":
		return Initialize
	case "next", "swap":
		return Next
	case "jump", "skip":
		return Jump
	default:
		return Error
	}
}
