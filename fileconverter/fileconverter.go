package fileconverter

import (
	"ServerProject/data"
	"encoding/json"
	"io/ioutil"
)

type FileConverter struct {
	filename string
}

func NewFileConverter(fn string) FileConverter {
	return FileConverter{
		filename: fn,
	}
}

func (f FileConverter) ConvertToGo() (data.OuterData, error) {
	d := data.OuterData{}
	jsonBytes, err := ioutil.ReadFile(f.filename)
	if err != nil {
		return d, err
	}

	err = json.Unmarshal(jsonBytes, &d)
	if err != nil {
		return d, err
	}
	return d, nil
}

func (f FileConverter) ConvertToFile(outerData data.OuterData) error {
	jsonBytes, err := json.Marshal(outerData)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(f.filename, jsonBytes, 0644)
	if err != nil {
		return err
	}
	return nil
}
