package documenter

type DocumentCreator interface {
	GenerateSection(section string)
	GenerateTableHeader()
	GenerateTableRow(path string, typ string, def string, desc string)
	ToString() string
}
