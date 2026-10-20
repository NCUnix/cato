package packs

import "github.com/NCUnix/cato/src/plugins/models"

type FieldPack struct {
	*models.Field
	Tags     string
	Comments string
}
