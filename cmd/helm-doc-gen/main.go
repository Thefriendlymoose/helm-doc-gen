package main

import (
	"fmt"
	"helm-doc-gen/pkg/documenter"
	"helm-doc-gen/pkg/parser"
	"helm-doc-gen/pkg/pathfinder"
	"helm-doc-gen/pkg/utils"
	"log"

	"github.com/goccy/go-yaml"
)

func main() {
	wd, err := utils.GetWorkingDir()

	if err != nil {
		log.Fatal(err)
	}

	pts, err := pathfinder.GetPathsToStuff(wd)

	if err != nil {
		log.Fatal(err)
	}

	for name, filePaths := range pts.HelmDirectories {
		valuesPath := filePaths[pathfinder.VALUES_NAME]
		pf, err := parser.GetParsedFile(valuesPath, yaml.CommentHeadPosition)
		if err != nil {
			panic(err)
		}
		mdb := documenter.GetMarkdownBuilder()
		mdb.GenerateTableHeader()

		for _, yi := range pf.OrderedItems {
			comments, ok := pf.FilteredComments[fmt.Sprintf("$.%s", yi.Path)]
			if ok {
				for _, comment := range comments {
					for _, t := range comment.Texts {
						if parser.IsValidDocComment(t) {
							ct, err := parser.GetComment(yi, t)
							if err != nil {
								continue
							}
							ct.GenerateRow(mdb)

						}
					}
				}
			}
		}
		documenter.GenerateFile(mdb.ToString(), name, documenter.MarkDown)
	}

}
