package cli

import (
	"flag"
	"fmt"
	"helm-doc-gen/pkg/builder"
	"os"
)

// Run parses and executes the correct subcommand
func Run() {
	if len(os.Args) < 2 {
		fmt.Println("Expected 'build'")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "build":
		handleBuild(os.Args[2:])
	default:
		fmt.Println("Unknown command:", os.Args[1])
		os.Exit(1)
	}
}

func handleBuild(args []string) {

	buildCommand := flag.NewFlagSet("build", flag.ExitOnError)
	buildPath := buildCommand.String("path", "", "set -path [AbsolutePathToFile] path to file values.yaml file to build docs")
	buildWorkingDir := buildCommand.Bool("runInWd", false, "set -runInWd if you want to run the program from the working directory, will then search and produce docs for all helm charts")
	buildGit := buildCommand.Bool("runGit", false, "set -runGit if you want from the working directory find the git root and find all charts")
	buildMd := buildCommand.Bool("md", false, "set -md if you want to produce docs in markdown")
	buildHTML := buildCommand.Bool("html", false, "set -html if you want to produce docs in html")
	buildOutputDir := buildCommand.String("output", "", "set -output [AbsolutePathToDir] where to place the output, will generate directories if missing")

	if len(args) == 0 {
		buildCommand.Usage()
		os.Exit(0)
	}

	buildCommand.Parse(args)

	buildConfig := &builder.BuildConfig{
		Path:       buildPath,
		WorkingDir: buildWorkingDir,
		Git:        buildGit,
		MarkDown:   buildMd,
		HTML:       buildHTML,
		OutputDir:  buildOutputDir,
	}

	builder.RunBuild(buildConfig)
}
