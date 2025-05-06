package invitation

import (
	"errors"
	"os"
)

func existingFile(files []string) (string, error) {
	for _, file := range files {
		_, err := os.Stat(file)
		if err != nil {
			if os.IsNotExist(err) {
				continue
			}

			continue
		}
		return file, nil
	}

	return "", errors.New("file not found")
}
