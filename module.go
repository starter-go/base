package base

import (
	"fmt"
	"strings"
)

const (
	theModuleName     = "github.com/starter-go/base"
	theModuleVersion  = "v0.9.12"
	theModuleRevision = 93
)

func GetModuleInfoText() string {

	b := new(strings.Builder)
	b.WriteString("[application.Module")

	b.WriteString(fmt.Sprintf(" Name:%s", theModuleName))
	b.WriteString(fmt.Sprintf(" Version:%s", theModuleVersion))
	b.WriteString(fmt.Sprintf(" Rev:%d", theModuleRevision))

	b.WriteRune(']')
	return b.String()

}
