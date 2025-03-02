package parser

import (
	"fmt"
	"helm-doc-gen/pkg/documenter"
	"regexp"
	"strings"
)

type CommentType int

const (
	PARAM CommentType = iota
	SECTION
	NOTVALID
)

var DocCommentType = map[CommentType]string{
	PARAM:    "@param",
	SECTION:  "@section",
	NOTVALID: "NOTVALID",
}

func (rt CommentType) String() string {
	return DocCommentType[rt]
}

type Comment struct {
	CommentType  CommentType
	Typ          string
	DerivedTyp   string
	Descr        string
	Path         string
	DefaultValue string
}

func GetComment(yi YamlItem, comment string) (*Comment, error) {
	ct := getCommentType(comment)
	if ct == NOTVALID {
		return nil, fmt.Errorf("%s, is not a valid comment type", comment)
	}
	typ, descr := parseParams(comment)
	if typ == "" {
		typ = yi.getDerivedType()
	}
	c := Comment{
		CommentType:  ct,
		Typ:          typ,
		Descr:        descr,
		Path:         yi.Path,
		DefaultValue: yi.getDerivedDefaultValue(),
	}

	return &c, nil
}

func (ct *Comment) GenerateRow(dc documenter.DocumentCreator) {
	switch ct.CommentType {
	case PARAM:
		dc.GenerateTableRow(ct.Path, ct.Typ, ct.DefaultValue, ct.Descr)
	case SECTION:
		// TODO
		dc.GenerateSection("TEMP")
	}
}

func IsValidDocComment(comment string) bool {
	if strings.Contains(comment, PARAM.String()) || strings.Contains(comment, SECTION.String()) {
		return true
	}
	return false
}

func getCommentType(comment string) CommentType {
	if strings.Contains(comment, PARAM.String()) {
		return PARAM
	}
	if strings.Contains(comment, SECTION.String()) {
		return SECTION
	}

	return NOTVALID
}

func parseParams(param string) (string, string) {

	typeRegex := regexp.MustCompile(`@type (\S+)`)
	descrRegex := regexp.MustCompile(`@descr (.+)`)

	typeMatch := typeRegex.FindStringSubmatch(param)
	descrMatch := descrRegex.FindStringSubmatch(param)

	var paramType, description string
	if len(typeMatch) > 1 {
		paramType = typeMatch[1]
	} else {
		paramType = ""
	}
	if len(descrMatch) > 1 {
		description = descrMatch[1]
	} else {
		description = ""
	}

	return paramType, description
}
