package documenter

import (
	"fmt"
	"strings"
)

type HTMLBuilder struct {
	sb strings.Builder
}

func GetHTMLBuilder() *HTMLBuilder {
	return &HTMLBuilder{sb: strings.Builder{}}
}

func (md *HTMLBuilder) GenerateDocumentTitle(title string) {
	md.sb.WriteString(fmt.Sprintf("# %s\n", title))
}

func (md *HTMLBuilder) GenerateDocumentDescription(description string) {
	md.sb.WriteString(fmt.Sprintf("%s\n", description))
}

func (md *HTMLBuilder) GenerateSectionTitle(title string) {
	md.sb.WriteString(fmt.Sprintf("## %s\n", title))
}

func (md *HTMLBuilder) GenerateSectionDescription(description string) {
	md.sb.WriteString(fmt.Sprintf("%s\n", description))
}

func (md *HTMLBuilder) GenerateTableHeader() {
	md.sb.WriteString("| path | type | default | description |\n")
	md.sb.WriteString("| ---- | ---- | ------- | ----------- |\n")
}

func (md *HTMLBuilder) GenerateTableRow(path string, typ string, def string, desc string) {
	md.sb.WriteString(fmt.Sprintf("| %s | %s | %s | %s |\n", path, typ, def, desc))
}

func (md *HTMLBuilder) ToString() string {
	return md.sb.String()
}
