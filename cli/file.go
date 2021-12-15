//
//  file.go
//  cli
//
//  Created by d-exclaimation on 5:02 PM.
//  Copyright © 2021 d-exclaimation. All rights reserved.
//

package cli

import (
	"io/ioutil"
	"os"
)

func Read(name string) (string, error) {
	jsonFile, err := os.Open(name)
	if err != nil {
		return "", err
	}
	defer func() {
		_ = jsonFile.Close()
	}()

	bytes, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

func Write(name string, content string) error {
	f, err := os.OpenFile(name, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755)
	if err != nil {
		return err
	}
	defer func() {
		_ = f.Close()
	}()
	_, err = f.Write([]byte(content))
	return err
}
