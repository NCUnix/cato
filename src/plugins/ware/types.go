package ware

import (
	"github.com/NCUnix/cato/src/plugins/common"
	"github.com/NCUnix/cato/src/plugins/models"
)

type fileGenerator func(ctx *common.GenContext) ([]*models.GenerateFileDesc, error)
