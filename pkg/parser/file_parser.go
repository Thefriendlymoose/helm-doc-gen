package parser

import (
	"fmt"
	"helm-doc-gen/pkg/documenter"
)

type tableRow struct {
	path         string
	typ          string
	defaultValue string
	description  string
}

type table struct {
	rows []*tableRow
}

type section struct {
	title       string
	description string
	table       *table
}

type Document struct {
	title       string
	description string
	sections    []*section
}

func GetDocumentation(pf *ParsedFile) *Document {
	document := Document{
		title:       "Dummy document header",
		description: "Dummy document description",
		sections: []*section{{
			title:       "Dummy section header",
			description: "Dummy section description",
			table:       &table{rows: []*tableRow{}},
		}},
	}

	for _, yi := range pf.OrderedItems {
		comments, ok := pf.FilteredComments[fmt.Sprintf("$.%s", yi.Path)]
		if ok {
			for _, comment := range comments {
				for _, t := range comment.Texts {
					if IsValidDocComment(t) {
						ct, err := GetComment(yi, t)
						if err != nil {
							continue
						}
						row := &tableRow{
							path:         ct.Path,
							typ:          ct.Typ,
							defaultValue: ct.DefaultValue,
							description:  ct.Descr,
						}
						document.sections[len(document.sections)-1].table.rows = append(
							document.sections[len(document.sections)-1].table.rows, row,
						)
					}
				}
			}
		}
	}

	return &document
}

func (d *Document) GenerateDocument(dc documenter.DocumentCreator) string {
	dc.GenerateDocumentTitle(d.title)
	dc.GenerateDocumentDescription(d.description)
	for _, section := range d.sections {
		section.generateSection(dc)
	}
	return dc.ToString()
}

func (s *section) generateSection(dc documenter.DocumentCreator) {
	dc.GenerateSectionTitle(s.title)
	dc.GenerateSectionDescription(s.description)
	s.table.generateTable(dc)
}

func (t *table) generateTable(dc documenter.DocumentCreator) {
	dc.GenerateTableHeader()
	for _, row := range t.rows {
		row.generateRow(dc)
	}
}

func (tr *tableRow) generateRow(dc documenter.DocumentCreator) {
	dc.GenerateTableRow(tr.path, tr.typ, tr.defaultValue, tr.description)
}
