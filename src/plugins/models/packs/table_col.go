package packs

import (
	"github.com/NCUnix/cato/src/plugins/models"
)

type TableColTmplPack struct {
	MessageTypeName string
	Cols            []*models.Col
}
