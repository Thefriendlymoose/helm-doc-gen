package parser

import (
	"fmt"
	"os"

	"github.com/goccy/go-yaml"
)

type ParsedFile struct {
	YamlContent          yaml.MapSlice
	Comments             yaml.CommentMap
	OrderedItems         []YamlItem
	OrderedTopLevelItems []string
	FilteredComments     yaml.CommentMap
	CommentPos           yaml.CommentPosition
}

type YamlItem struct {
	Path string
	Data any
}

func GetParsedFile(path string, commentPos yaml.CommentPosition) (*ParsedFile, error) {
	pf := ParsedFile{}
	err := pf.getDataAndComments(path)
	if err != nil {
		return nil, err
	}

	pf.OrderedItems = []YamlItem{}
	pf.OrderedTopLevelItems = []string{}
	pf.CommentPos = commentPos
	pf.traverseMapSlice(pf.YamlContent, "")
	pf.filterComments()
	return &pf, nil
}

func (pf *ParsedFile) getDataAndComments(path string) error {
	data, err := os.ReadFile(path)

	if err != nil {
		return fmt.Errorf("error reading file %v", err)
	}
	commentMap := yaml.CommentMap{}
	var result yaml.MapSlice
	err = yaml.UnmarshalWithOptions(data, &result, yaml.CommentToMap(commentMap), yaml.UseOrderedMap())
	if err != nil {
		return fmt.Errorf("error unmarshal file: %v", err)
	}

	pf.YamlContent = result
	pf.Comments = commentMap

	return nil
}

func (pf *ParsedFile) traverseMapSlice(data yaml.MapSlice, currentPath string) {
	for _, item := range data {
		key := fmt.Sprintf("%v", item.Key)
		absolutePath := key
		if currentPath != "" {
			absolutePath = currentPath + "." + key
		}

		switch v := item.Value.(type) {
		case yaml.MapSlice:
			if len(v) == 0 {
				pf.OrderedItems = append(pf.OrderedItems, YamlItem{Path: absolutePath, Data: "{}"})
			} else {
				if currentPath == "" {
					pf.OrderedTopLevelItems = append(pf.OrderedTopLevelItems, key)
				}
				pf.traverseMapSlice(v, absolutePath)
			}
		case []interface{}:
			pf.OrderedItems = append(pf.OrderedItems, YamlItem{Path: absolutePath, Data: v})
		default:
			pf.OrderedItems = append(pf.OrderedItems, YamlItem{Path: absolutePath, Data: v})
		}
	}
}

func (pf *ParsedFile) filterComments() {
	cm := pf.Comments
	cp := pf.CommentPos
	fc := make(yaml.CommentMap)
	for k, commentSlice := range cm {
		fs := []*yaml.Comment{}
		for _, comment := range commentSlice {
			if comment.Position == cp {
				fs = append(fs, comment)
			}
		}

		if len(fs) > 0 {
			fc[k] = fs
		}
	}

	pf.FilteredComments = fc
}

func (yi *YamlItem) getDerivedDefaultValue() string {
	return fmt.Sprint(yi.Data)
}

func (yi *YamlItem) getDerivedType() string {
	return fmt.Sprintf("%T", yi.Data)

}
