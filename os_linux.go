package ostool

import (
	"fmt"
	"io/ioutil"
	"os/exec"
	"strings"
)

func GetVersion() string {
	if v := getByLSB(); v != "" {
		return v
	} else if v := getByOSRelease(); v != "" {
		return v
	}
	return ""
}

func getByLSB() string {
	proc := exec.Command("lsb_release", "-irc")
	raw, err := proc.Output()
	if err != nil {
		return ""
	}
	if len(raw) == 0 {
		return ""
	}

	result := string(raw)
	infoMap := map[string]string{}
	for _, line := range strings.Split(result, "\n") {
		if line == "" {
			continue
		}
		ts := strings.Split(line, ":")
		if len(ts) != 2 {
			continue
		}
		key := strings.Trim(ts[0], " \t")
		value := strings.Trim(ts[1], " \t")
		switch key {
		case "Distributor ID":
			infoMap["system"] = value
		case "Release":
			infoMap["version"] = value
		}
	}
	if len(infoMap) == 2 {
		return fmt.Sprintf("%v%v", infoMap["system"], infoMap["version"])
	}
	return ""
}

func getByOSRelease() string {
	bys, err := ioutil.ReadFile("/etc/os-release")
	if err != nil {
		return ""
	}

	infoMap := map[string]string{}
	for _, line := range strings.Split(string(bys), "\n") {
		if line == "" {
			continue
		}

		ts := strings.Split(line, "=")
		key := strings.Trim(ts[0], "\"")
		value := strings.Trim(ts[1], "\"")
		switch key {
		case "NAME", "DISTRIB_ID", "ID":
			infoMap["system"] = value
		case "VERSION_ID", "DISTRIB_RELEASE", "VERSION":
			infoMap["version"] = value
		}
	}

	if len(infoMap) == 2 {
		return fmt.Sprintf("%v%v", infoMap["system"], infoMap["version"])
	}
	return ""
}
