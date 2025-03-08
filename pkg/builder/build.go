package builder

import (
	"helm-doc-gen/pkg/documenter"
	"helm-doc-gen/pkg/parser"
	"helm-doc-gen/pkg/pathfinder"
	"helm-doc-gen/pkg/utils"
	"log"
	"os"
	"path/filepath"

	"github.com/goccy/go-yaml"
)

type BuildConfig struct {
	Path       *string
	WorkingDir *bool
	Git        *bool
	MarkDown   *bool
	HTML       *bool
	OutputDir  *string
}

func RunBuild(cmds *BuildConfig) {

	var wd string
	var err error

	if *cmds.Path != "" {
		if filepath.Base(*cmds.Path) == pathfinder.VALUES_FILE {
			wd = filepath.Dir(*cmds.Path)
		} else {
			log.Fatal(err)
			os.Exit(1)
		}
	} else if *cmds.WorkingDir {
		wd, err = utils.GetWorkingDir()

		if err != nil {
			log.Fatal(err)
			os.Exit(1)
		}
	} else if *cmds.Git {
		wd, err = utils.GetGitRoot(wd)

		if err != nil {
			log.Fatal(err)
			os.Exit(1)
		}
	}

	pts, err := pathfinder.GetPathsToStuff(wd)

	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	for name, filePaths := range pts.HelmDirectories {
		valuesPath := filePaths[pathfinder.VALUES_NAME]
		pf, err := parser.GetParsedFile(valuesPath, yaml.CommentHeadPosition)
		if err != nil {
			panic(err)
		}
		mdb := documenter.GetMarkdownBuilder()
		doc := parser.GetDocumentation(pf)
		documenter.GenerateFile(doc.GenerateDocument(mdb), name, documenter.MarkDown)
	}
}
