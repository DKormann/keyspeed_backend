package storage

import (
	"keyspeed/util"
	"os"
)

const filepath = "data/"

func Set(key string, data []byte) error {
	file, err := os.OpenFile(filepath+key, os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		return err
	}

	defer file.Close()

	_, err = file.Write(data)

	util.BLogln("wrote data: ", string(data))

	return err
}

func Append(key string, data []byte) error {
	file, err := os.OpenFile(filepath+key, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
	if err != nil {
		return err
	}

	defer file.Close()

	// add "\n\n" to separate the data

	_, err = file.Write(data)
	util.BLogln("wrote data: ", string(data))

	return err
}

func ListAll() (result string, err error) {

	files, err := os.ReadDir(filepath)

	if err != nil {
		return
	}

	for _, file := range files {
		result += file.Name() + "\n"
	}

	return
}

func Get(key string) ([]byte, error) {

	file, err := os.Open(filepath + key)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	stat, err := file.Stat()
	if err != nil {
		return nil, err
	}

	size := stat.Size()
	result := make([]byte, size)

	_, err = file.Read(result)
	if err != nil {
		return nil, err
	}

	util.GLogln("got data: ", string(result))

	return result, nil
}
