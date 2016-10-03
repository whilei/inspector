package inspector

import (
	"encoding/json"
	"os"
	"path/filepath"
	"io/ioutil"
)

func hostname() string {
	hostname, err := os.Hostname()
	if err != nil {
		hostname = ""
	}
	return hostname
}

func DetectModel() string {
	appPath, err := filepath.Abs("../inspector/resource")
	if err != nil {
		panic(err)
	}
	hostsConfigFile := filepath.Join(appPath, ConfigDir, HostsFile)
	conf, err := ioutil.ReadFile(hostsConfigFile);
	if err != nil {
		panic("hosts config not found : " + hostsConfigFile)
	}
	config := parseHostsConfig(conf)
	hostname := hostname()
	for i, k := range config {
		for _, h := range k {
			if h == hostname {
				return i
			}
		}
	}
	return ""
}

func parseHostsConfig(content []byte) map[string][]string {
	var config map[string][]string
	err := json.Unmarshal(content, &config)
	if err != nil {
		panic("xxxxxxxxx")
	}
	return config
}
