package word

type ExportWordFile = wordFile

func NewExportWordFile(path string) *ExportWordFile {
	return &ExportWordFile{path:path}
}

func (w *ExportWordFile) ExportWords() []string {
	return w.words
}
