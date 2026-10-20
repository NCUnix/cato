package migrate

import (
	"log"

	"google.golang.org/protobuf/reflect/protoreflect"

	"github.com/NCUnix/cato/generated"
	"github.com/NCUnix/cato/src/plugins/common"
	"github.com/NCUnix/cato/src/plugins/sprinkles"
)

func init() {
	sprinkles.Register(func() sprinkles.Sprinkle {
		return new(MysqlSprinkle)
	})
}

type MysqlSprinkle struct {
	ddlOpt       *generated.DdlOptions
	mysqlImplOpt *generated.MysqlImplDdlOpt
}

func (m *MysqlSprinkle) FromExtType() protoreflect.ExtensionType {
	return generated.E_DdlOpt
}

func (m *MysqlSprinkle) WorkOn(desc protoreflect.Descriptor) bool {
	_, ok := desc.(protoreflect.MessageDescriptor)
	return ok
}

func (m *MysqlSprinkle) Init(value interface{}) {
	exValue, ok := value.(*generated.DdlOptions)
	if !ok {
		log.Fatalln("[-]cato MysqlSprinkle expect DdlOptions")
	}
	m.ddlOpt = exValue
	mysqlImpl := exValue.MysqlImplOpt
	if mysqlImpl == nil {
		return
	}
	m.mysqlImplOpt = mysqlImpl
}

func (m *MysqlSprinkle) Register(ctx *common.GenContext) error {
	if m.ddlOpt == nil || m.mysqlImplOpt == nil {
		return nil
	}
	return nil
}
