package services

import (
	"fmt"
	"os"
)

func createFolder(name string) (string, error) {
	f, err := os.Stat(name)
	if err != nil {
		if os.IsNotExist(err) {
			err = os.Mkdir(name, 0755)
			if err != nil {
				return "", err
			}
			return name, nil
		}
		return "", err
	}

	if f.IsDir() {
		return "", fmt.Errorf("%s already exists", name)
	}
	return "", fmt.Errorf("not a folder")
}
