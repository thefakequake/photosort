package main

import (
	"encoding/json"
	"os"
)

type Config struct {
	InputDir         string   `json:"inputDir"`
	OutputDir        string   `json:"outputDir"`
	SupportedFormats []string `json:"supportedFormats"`
	NamePrefix       string   `json:"namePrefix"`
	NextImage        int      `json:"nextImage"`
}

const defaultConfig = `{
	"inputDir": ".",
	"outputDir": "out",
	"supportedFormats": [
		"jpeg",
		"jpg",
		"png"
	],
	"namePrefix": "IMG",
	"nextImage": 1
}`

func GetConfig() bool {
	var c Config

	if _, err := os.Stat("config.json"); os.IsNotExist(err) {
		os.WriteFile("config.json", []byte(defaultConfig), os.ModePerm)
		return true
	}

	f, _ := os.Open("config.json")
	json.NewDecoder(f).Decode(&c)

	cfg = c

	return false
}

func IncrementNextImage(num int) {
	var c Config

	f, _ := os.Open("config.json")
	json.NewDecoder(f).Decode(&c)

	c.NextImage = num

	dat, _ := json.MarshalIndent(c, "", "\t")

	os.WriteFile("config.json", dat, os.ModePerm)
}
