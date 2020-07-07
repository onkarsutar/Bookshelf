package file

import "io/ioutil"

// ReadFile reads contents from provided file path
func ReadFile(filePath string) ([]byte, error) {
	return ioutil.ReadFile(filePath)

}

// WriteFile writes provided  bytes to file
func WriteFile(filePath string, data []byte) error {
	return ioutil.WriteFile(filePath, data, 0644)
}