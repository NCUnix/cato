package utils

import (
	"strings"

	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"

	"github.com/NCUnix/cato/generated"
)

func GetGoPackageName(importPath string) string {
	patterns := strings.Split(GetGoFilePath(importPath), "/")
	if len(patterns) == 0 || patterns[0] == "." {
		return ""
	}
	return patterns[len(patterns)-1]
}

func GetGoFilePath(importPath string) string {
	return strings.Trim(importPath, "\"")
}

func GetTagKey(tagRaw string) string {
	patterns := strings.Split(tagRaw, ":")
	// invalid tag format
	if len(patterns) < 2 {
		return ""
	}
	return patterns[0]
}

func GetCatoPackageFromFile(filedesc protoreflect.FileDescriptor) (string, bool) {
	if !proto.HasExtension(filedesc.Options(), generated.E_CatoOpt) {
		return "", false
	}
	catoOptions := proto.GetExtension(filedesc.Options(), generated.E_CatoOpt).(*generated.CatoOptions)
	return catoOptions.CatoPackage, catoOptions.CatoPackage != ""
}
