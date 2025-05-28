package savers

import (
	"encoding/gob"
	"os"
)

// SaveGob saves data of any type to a file in gob format.
// It returns an error if the file already exists or if there is an issue during saving.
// It is a generic function that can handle any type T were T is a type that can be encoded with gob. struct must have exported/public fields.
func SaveGob[T any](path string, data T) error {

	// Check if the file already exists
	if thereIsFile(path) {
		return os.ErrExist
	}

	return ForcedSaveGob(path, data)
}

// ForcedSaveGob saves data of any type to a file in gob format.
// Careful overwrites the file if it already exists.
func ForcedSaveGob[T any](path string, data T) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := gob.NewEncoder(file)

	if err := encoder.Encode(data); err != nil {
		return err
	}

	return nil
}

// thereIsFile checks if a file exists at the given path.
func thereIsFile(path string) bool {
	_, err := os.Stat(path)
	return !os.IsNotExist(err)
}
