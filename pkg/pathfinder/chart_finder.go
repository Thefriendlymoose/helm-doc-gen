package pathfinder

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
)

const (
	CHART_FILE  = "Chart.yaml"
	VALUES_FILE = "values.yaml"
	CHART_NAME  = "Chart"
	VALUES_NAME = "values"
)

type PathsToStuff struct {
	HelmDirectories map[string]map[string]string
}

func GetPathsToStuff(startingDir string) (*PathsToStuff, error) {
	pts := PathsToStuff{}
	err := pts.getHelmDirectories(startingDir)
	if err != nil {
		return nil, err
	}

	return &pts, nil
}

func (pts *PathsToStuff) getHelmDirectories(fromDir string) error {
	foundHelmDirs := make(map[string]map[string]string)

	err := filepath.Walk(fromDir, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.Name() == CHART_FILE {
			valuesAndChartPaths, err := getValueFilesPathFromHelmDir(filepath.Dir(path))
			if err != nil {
				return err
			}
			nameOfHelmDir := filepath.Base(filepath.Dir(path))
			foundHelmDirs[nameOfHelmDir] = valuesAndChartPaths
		}

		return nil
	})

	if err != nil {
		return fmt.Errorf("getHelmDirs: %s", err.Error())
	}
	pts.HelmDirectories = foundHelmDirs
	return nil
}

func getValueFilesPathFromHelmDir(path string) (map[string]string, error) {
	files, err := os.ReadDir(path)

	if err != nil {
		return nil, fmt.Errorf("getValueFilesPathFromHelmDir: %s", err.Error())
	}
	helmFiles := make(map[string]string)
	for _, file := range files {
		if !file.IsDir() {
			if file.Name() == CHART_FILE {
				helmFiles[CHART_NAME] = filepath.Join(path, file.Name())
			}

			if file.Name() == VALUES_FILE {
				helmFiles[VALUES_NAME] = filepath.Join(path, file.Name())
			}
		}
	}

	return helmFiles, nil
}
