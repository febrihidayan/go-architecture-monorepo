package valueobject

import "github.com/febrihidayan/go-architecture-monorepo/pkg/lang"

/**
 * Reference mime types
 *
 * https://stackoverflow.com/a/4212908
 * https://codingislove.com/list-mime-types-2016/
 */
const (
	XLSX string = "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet"
	XLSM string = "application/vnd.ms-excel.sheet.macroEnabled.12"
	XLS  string = "application/vnd.ms-excel"
	CSV  string = "application/csv"
	PDF  string = "application/pdf"
	DOC  string = "application/msword"
	DOCX string = "application/vnd.openxmlformats-officedocument.wordprocessingml.document"
	RTF  string = "application/rtf"

	JPG   string = "image/jpg"
	JPEG  string = "image/jpeg"
	PNG   string = "image/png"
	JFIF  string = "image/jfif"
	PJPEG string = "image/pjeg"
	PJP   string = "image/pjp"
	GIF   string = "image/gif"
	SVG   string = "image/svg"
)

type FileType struct {
	value string
}

func NewFileType(val string) (FileType, error) {
	file := FileType{value: val}
	if val != "" {
		if !file.IsAllowedExtension() {
			return FileType{}, lang.ErrUnsupportFile
		}
	}
	return file, nil
}

func (f FileType) IsAllowedExtension() bool {
	allowedExts := []string{
		XLSX, XLSM, XLS, CSV, PDF, DOC, DOCX, RTF,
		JPG, JPEG, PNG, JFIF, PJPEG, PJP, GIF, SVG,
	}

	for _, allowedExt := range allowedExts {
		if allowedExt == f.value {
			return true
		}
	}
	return false
}

func (f FileType) Type() string {
	return f.value
}

func (f FileType) String() string {
	mime := ""
	switch f.value {
	case XLSX:
		mime = "xlsx"
	case XLSM:
		mime = "xlsm"
	case XLS:
		mime = "xls"
	case CSV:
		mime = "csv"
	case PDF:
		mime = "pdf"
	case DOC:
		mime = "doc"
	case DOCX:
		mime = "docx"
	case RTF:
		mime = "rtf"
	case JPG:
		mime = "jpg"
	case JPEG:
		mime = "jpeg"
	case PNG:
		mime = "png"
	case JFIF:
		mime = "jfif"
	case PJPEG:
		mime = "pjpeg"
	case PJP:
		mime = "pjp"
	case GIF:
		mime = "gif"
	case SVG:
		mime = "svg"
	}

	return mime
}
