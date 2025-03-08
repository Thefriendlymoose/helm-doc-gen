package cli

import (
	"flag"
	"fmt"
	"helm-doc-gen/pkg/builder"
	"os"
)

var commands = map[string]string{
	"build": "Builds documentation",
	"help":  "Displays this help message",
}

// Run parses and executes the correct subcommand
func Run() {
	if len(os.Args) < 2 {
		fmt.Println("Expected 'build'")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "build":
		handleBuild(os.Args[2:])
	case "help":
		printUsage()
	default:
		fmt.Println("Unknown command:", os.Args[1])
		os.Exit(1)
	}
}

func handleBuild(args []string) {

	buildCommand := flag.NewFlagSet("build", flag.ExitOnError)
	buildPath := buildCommand.String("path", "", "set -path='path/to/file' path to file values.yaml file to build docs")
	buildWorkingDir := buildCommand.Bool("runInWd", false, "set -runInWd if you want to run the program from the working directory, will then search and produce docs for all helm charts")
	buildGit := buildCommand.Bool("runGit", false, "set -runGit if you want from the working directory find the git root and find all charts")
	buildMd := buildCommand.Bool("md", false, "set -md if you want to produce docs in markdown")
	buildHTML := buildCommand.Bool("html", false, "set -html if you want to produce docs in html")
	buildOutputDir := buildCommand.String("output", "", "set -output='path/to/dir' where to place the output, will generate directories if missing, if not specified will generate in workingdir in directory ./helm-docs-output")

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

func printUsage() {
	fmt.Println("Usage: project <command> [options]")
	fmt.Println("\nAvailable commands:")
	for cmd, desc := range commands {
		fmt.Printf("  %-10s %s\n", cmd, desc)
	}
	fmt.Println("\nUse 'project <command> -h' for more details.")
}
