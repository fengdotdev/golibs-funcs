package savers_test

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"os"
	"testing"

	"github.com/fengdotdev/golibs-funcs/v0/savers"
	"github.com/fengdotdev/golibs-testing/assert"
)

// get the hash of the test name
func hashTName(t *testing.T) string {
	t.Helper()

	// get t name
	name := t.Name()
	b := []byte(name)
	hash := sha256.New()
	hash.Write(b)
	return hex.EncodeToString(hash.Sum(nil))

}

// pathTName returns a unique path for the test based on its name
func pathTName(t *testing.T) string {
	name := hashTName(t)
	return fmt.Sprintf("%s/%s.gob", os.TempDir(), name)
}

const (
	temp_path = "test_load_gob"
)

func TestSaveLoadGob(t *testing.T) {
	// Create a temporary file to save the gob data
	tmpFile, err := os.CreateTemp("", temp_path)
	if err != nil {
		t.Fatalf("Failed to create temp file: %v", err)
	}
	defer os.Remove(tmpFile.Name())

	t.Run("gob map string string", func(t *testing.T) {

		path := pathTName(t)

		// Create a map to save
		data := map[string]string{
			"key1": "value1",
			"key2": "value2",
			"key3": "value3",
		}
		err := savers.SaveGob(path, data)
		t.Log(err)
		assert.NotErrorWithMessage(t, err, "SaveGob should not return an error")

		loadedData, err := savers.LoadGob[map[string]string](path)
		assert.NotErrorWithMessage(t, err, "LoadGob should not return an error")
		assert.EqualWithMessage(t, data, loadedData, "Loaded data should match saved data")

		defer func() {
			// Clean up the temporary file
			err := os.Remove(path)
			assert.NotErrorWithMessage(t, err, "Failed to remove temp file")
		}()

	})

	t.Run("gob map string int", func(t *testing.T) {
		path := pathTName(t)

		// Create a map to save
		data := map[string]int{
			"key1": 1,
			"key2": 2,
			"key3": 3,
		}
		err := savers.SaveGob(path, data)
		assert.NotErrorWithMessage(t, err, "SaveGob should not return an error")

		loadedData, err := savers.LoadGob[map[string]int](path)
		assert.NotErrorWithMessage(t, err, "LoadGob should not return an error")
		assert.EqualWithMessage(t, data, loadedData, "Loaded data should match saved data")

		defer func() {
			// Clean up the temporary file
			err := os.Remove(path)
			assert.NotErrorWithMessage(t, err, "Failed to remove temp file")
		}()
	})

	t.Run("gob map string struct", func(t *testing.T) {
		path := pathTName(t)

		type MyStruct struct {
			Field1 string
			Field2 int
		}

		// Create a map to save
		data := map[string]MyStruct{
			"key1": {Field1: "value1", Field2: 1},
			"key2": {Field1: "value2", Field2: 2},
			"key3": {Field1: "value3", Field2: 3},
		}
		err := savers.SaveGob(path, data)
		assert.NotErrorWithMessage(t, err, "SaveGob should not return an error")

		loadedData, err := savers.LoadGob[map[string]MyStruct](path)
		assert.NotErrorWithMessage(t, err, "LoadGob should not return an error")
		assert.EqualWithMessage(t, data, loadedData, "Loaded data should match saved data")

		defer func() {
			// Clean up the temporary file
			err := os.Remove(path)
			assert.NotErrorWithMessage(t, err, "Failed to remove temp file")
		}()
	})

	t.Run("gob map string slice", func(t *testing.T) {
		path := pathTName(t)

		// Create a map to save
		data := map[string][]string{
			"key1": {"value1a", "value1b"},
			"key2": {"value2a", "value2b"},
			"key3": {"value3a", "value3b"},
		}
		err := savers.SaveGob(path, data)
		assert.NotErrorWithMessage(t, err, "SaveGob should not return an error")

		loadedData, err := savers.LoadGob[map[string][]string](path)
		assert.NotErrorWithMessage(t, err, "LoadGob should not return an error")
		assert.EqualWithMessage(t, data, loadedData, "Loaded data should match saved data")

		defer func() {
			// Clean up the temporary file
			err := os.Remove(path)
			assert.NotErrorWithMessage(t, err, "Failed to remove temp file")
		}()
	})

	t.Run("Complex struct with nested map", func(t *testing.T) {
		path := pathTName(t)

		type InnerStruct struct {
			FieldA string
			FieldB int
		}

		type OuterStruct struct {
			Name   string
			Value  int
			Nested map[string]InnerStruct
		}

		data := OuterStruct{
			Name:  "Test",
			Value: 42,
			Nested: map[string]InnerStruct{
				"inner1": {FieldA: "A", FieldB: 1},
				"inner2": {FieldA: "B", FieldB: 2},
			},
		}

		err := savers.SaveGob(path, data)
		assert.NotErrorWithMessage(t, err, "SaveGob should not return an error")

		loadedData, err := savers.LoadGob[OuterStruct](path)
		assert.NotErrorWithMessage(t, err, "LoadGob should not return an error")
		assert.EqualWithMessage(t, data, loadedData, "Loaded data should match saved data")

		// Check if the nested map is correctly loaded
		assert.NotNilWithMessage(t, loadedData.Nested, "Nested map should not be nil")
		assert.EqualWithMessage(t, 2, len(loadedData.Nested), "Nested map should have 2 elements")
		assert.EqualWithMessage(t, "A", loadedData.Nested["inner1"].FieldA, "Nested fieldA should match")
		assert.EqualWithMessage(t, 1, loadedData.Nested["inner1"].FieldB, "Nested fieldB should match")
		assert.EqualWithMessage(t, "B", loadedData.Nested["inner2"].FieldA, "Nested fieldA should match")
		assert.EqualWithMessage(t, 2, loadedData.Nested["inner2"].FieldB, "Nested fieldB should match")

		defer func() {
			// Clean up the temporary file
			err := os.Remove(path)
			assert.NotErrorWithMessage(t, err, "Failed to remove temp file")
		}()
	})
}
