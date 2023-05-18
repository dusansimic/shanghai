package main

import (
	"fmt"
	"os"

	"shanghai/shanghai"

	"github.com/akamensky/argparse"
)

var (
	Version = "dev"
)

func main() {
	parser := argparse.NewParser("shg", "Build hierarchical tree of container images")

	build := parser.NewCommand("build", "Build image")

	shgfile := build.String("f", "file", &argparse.Options{
		Required: false,
		Help:     "Shangaifile to use",
		Default:  "Shangaifile",
	})

	check := build.Flag("c", "check", &argparse.Options{
		Required: false,
		Help:     "Check Shangaifile for errors",
	})

	image := build.String("i", "image", &argparse.Options{
		Required: true,
		Help:     "Image to build",
	})

	err := parser.Parse(os.Args)
	if err != nil {
		fmt.Println(parser.Usage(err))
		os.Exit(1)
	}

	file, err := shanghai.ReadShanghaifile(*shgfile)
	if err != nil {
		fmt.Println(fmt.Errorf("failed to read ile: %w", err))
		os.Exit(1)
	}

	if err := shanghai.ValidateShanghaifile(file, *image); err != nil {
		fmt.Println(fmt.Errorf("failed to validate Shangaifile: %w", err))
		os.Exit(1)
	}

	if *check {
		os.Exit(0)
	}

	f, err := shanghai.SearchConfigFile()
	if err != nil {
		fmt.Println(fmt.Errorf("failed to find config file: %w", err))
		os.Exit(1)
	}

	cfg, err := shanghai.ReadConfig(f)
	if err != nil {
		fmt.Println(fmt.Errorf("failed to read config file: %w", err))
		os.Exit(1)
	}

	if err := shanghai.BuildImages(cfg, file, *image); err != nil {
		fmt.Println(fmt.Errorf("failed to build images: %w", err))
		os.Exit(1)
	}
}
