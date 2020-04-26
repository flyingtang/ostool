package ostool

import (
	"golang.org/x/sys/windows/registry"
	"runtime"
	"strings"
)

func GetVersion() string {
	k, err := registry.OpenKey(registry.LOCAL_MACHINE, `SOFTWARE\Microsoft\Windows NT\CurrentVersion`, registry.QUERY_VALUE)
	if err != nil {
		return ""
	}
	defer k.Close()
	pn, _, err := k.GetStringValue("ProductName")
	if err != nil {
		return ""
	}
	return strings.Replace(pn+runtime.GOARCH, " ", "", -1)
}
