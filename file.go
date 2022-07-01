package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/rwcarlsen/goexif/exif"
)

func ValidFile(name string) bool {
	ext := filepath.Ext(name)
	for _, t := range cfg.SupportedFormats {
		if "."+t == strings.ToLower(ext) {
			return true
		}
	}

	return false
}

func CreateIfNotExists(dirs ...string) bool {
	changes := false
	for _, dir := range dirs {
		if _, err := os.Stat(dir); os.IsNotExist(err) {
			os.Mkdir(dir, os.ModePerm)
			changes = true
		}
	}
	return changes
}

func ParseImage(path string) (Image, error) {
	i := Image{
		Path: path,
	}
	file, err := os.Open(path)
	if err != nil {
		return i, err
	}

	x, err := exif.Decode(file)
	if err != nil {
		fmt.Printf("%s: couldn't decode exif data\n", path)
	} else {
		i.TakenAt, err = x.DateTime()
		if err != nil {
			return i, err
		}
	}

	return i, nil
}

type Image struct {
	Path    string
	TakenAt time.Time
}
