package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

var cfg Config

func main() {
	defer func() {
		fmt.Print("\npress enter to exit...")
		fmt.Scanln()
	}()

	created := GetConfig()
	if created {
		fmt.Println("created config.json, please configure before running the program again")
		return
	}

	var images []Image

	if err := filepath.Walk(cfg.InputDir, func(path string, f os.FileInfo, err error) error {
		if ValidFile(path) {
			image, err := ParseImage(path)
			if err != nil {
				fmt.Printf("%s: couldn't parse image\n", path)
			}

			images = append(images, image)
		}
		return err
	}); err != nil {
		fmt.Println("couldn't find input directory")
		return
	}

	sort.Slice(images, func(i, j int) bool {
		return images[i].TakenAt.Before(images[j].TakenAt)
	})

	fmt.Printf("found %d images\n", len(images))

	if CreateIfNotExists(cfg.OutputDir) {
		IncrementNextImage(1)
		cfg.NextImage = 1
		fmt.Println("reset next image to 1")
	}

	count := 0
	unknowns := 0

	for i, img := range images {
		count = i
		var dayDir string
		var name string

		if img.TakenAt.Year() == 1 {
			dayDir = fmt.Sprintf("%s/unknown", cfg.OutputDir)
			if unknowns == 0 {
				CreateIfNotExists(dayDir)
			}
			unknowns++
			name = fmt.Sprintf("%sU_%04d", cfg.NamePrefix, cfg.NextImage-1+unknowns)
		} else {
			yearDir := fmt.Sprintf("%s/%02d", cfg.OutputDir, img.TakenAt.Year())
			monthDir := fmt.Sprintf("%s/%02d", yearDir, img.TakenAt.Month())
			dayDir = fmt.Sprintf("%s/%02d", monthDir, img.TakenAt.Day())
			CreateIfNotExists(yearDir, monthDir, dayDir)
			name = fmt.Sprintf("%s_%04d", cfg.NamePrefix, cfg.NextImage+i-unknowns)
		}

		dat, err := os.ReadFile(img.Path)
		if err != nil {
			log.Printf("%s: couldn't read", img.Path)
			continue
		}

		err = os.WriteFile(fmt.Sprintf("%s/%s%s", dayDir, name, strings.ToLower(filepath.Ext(img.Path))), dat, os.ModePerm)
		if err != nil {
			log.Println(err)
		}
	}
	IncrementNextImage(cfg.NextImage + count - unknowns)
}
