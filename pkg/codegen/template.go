package codegen

import (
	"github.com/goradd/orm/pkg/model"
	"io"
)

type Template interface {
	Overwrite() bool
}

type DatabaseGenerator interface {
	Template
	FileName(string) string
	GenerateDatabase(*model.Database, io.Writer) error
}

type TableGenerator interface {
	Template
	FileName(*model.Table) string
	GenerateTable(*model.Table, io.Writer) error
}

type EnumGenerator interface {
	Template
	FileName(*model.EnumTable) string
	GenerateEnum(*model.EnumTable, io.Writer) error
}

var templates []Template

func RegisterTemplate(t Template) {
	templates = append(templates, t)
}
