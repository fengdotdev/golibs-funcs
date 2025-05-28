package savers

import (
	"encoding/gob"
	"os"
)

// LoadGob loads data of any type from a file in gob format.
// It returns the data and an error if there is an issue during loading.
func LoadGob[T any](path string) (T, error) {
	var zero T

	file, err := os.Open(path)
	if err != nil {
		return zero, err
	}
	defer file.Close()

	var data T
	decoder := gob.NewDecoder(file)

	if err := decoder.Decode(&data); err != nil {
		return zero, err
	}

	return data, nil
}

// LoadGobTo loads data from a file in gob format into the provided variable.
// The data parameter should be a pointer to the variable you want to populate.
func LoadGobTo[T any](path string, data *T) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	decoder := gob.NewDecoder(file)

	if err := decoder.Decode(data); err != nil {
		return err
	}

	return nil
}
