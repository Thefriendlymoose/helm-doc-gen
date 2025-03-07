package documenter

type DocumentCreator interface {
	GenerateDocumentTitle(title string)
	GenerateDocumentDescription(desc string)
	GenerateSectionTitle(title string)
	GenerateSectionDescription(desc string)
	GenerateTableHeader()
	GenerateTableRow(path string, typ string, def string, desc string)
	ToString() string
}
