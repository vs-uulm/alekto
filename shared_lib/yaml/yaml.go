package yaml

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"sync"
)

var mu sync.Mutex

/**
 * parse data into given struct
 */
func ParseToInterface(data []byte, v interface{}) {

	err := yaml.Unmarshal([]byte(data), v)

	if err != nil {
		fmt.Printf("ParseToInterfaceError data: %v \n", string(data))
		fmt.Printf("error: %v", err)
	}
}

func LoadStructFromFile(filename string, v interface{}) {

	var file []byte

	file, _ = ioutil.ReadFile(filename)

	ParseToInterface(file, v)
}

func SaveStructToFile(filename string, v interface{}) error {

	mu.Lock()
	defer mu.Unlock()

	bytes, err := yaml.Marshal(v)
	if err != nil {
		return err
	}

	return ioutil.WriteFile(filename, bytes, 0644)
}

func AppendStructToFile(filename string, v interface{}) error {

	bytes, err := yaml.Marshal(v)
	if err != nil {
		return err
	}

	mu.Lock()
	defer mu.Unlock()

	file, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	defer file.Close()

	if err != nil {
		panic(err)
	}

	_, err = file.Write(bytes)

	return err
}

func LoadStructOfFile(filename string) (v interface{}, err error) {

	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		return v, err
	}

	err = yaml.Unmarshal(bytes, &v)
	if err != nil {
		return v, err
	}

	return v, nil
}
