package documenter

import (
	"fmt"
	"strings"
)

type MarkdownBuilder struct {
	sb strings.Builder
}

func GetMarkdownBuilder() *MarkdownBuilder {
	return &MarkdownBuilder{sb: strings.Builder{}}
}

func (md *MarkdownBuilder) GenerateSection(section string) {
	md.sb.WriteString(fmt.Sprintf("# %s\n", section))
}

func (md *MarkdownBuilder) GenerateTableHeader() {
	md.sb.WriteString("| path | type | default | description |\n")
	md.sb.WriteString("| ---- | ---- | ------- | ----------- |\n")
}

func (md *MarkdownBuilder) GenerateTableRow(path string, typ string, def string, desc string) {
	md.sb.WriteString(fmt.Sprintf("| %s | %s | %s | %s |\n", path, typ, def, desc))
}

func (md *MarkdownBuilder) ToString() string {
	return md.sb.String()
}
