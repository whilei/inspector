package inspector

import (
	"encoding/json"
	"os"
	"path/filepath"
	"io/ioutil"
	"errors"
	"github.com/astaxie/beego"
)

func hostname() string {
	hostname, err := os.Hostname()
	if err != nil {
		hostname = ""
	}
	return hostname
}

func ConfigFile() (string, error) {
	env := DetectModel()
	if env == "" {
		return "", errors.New("cant detect run mode")
	}
	configFile := filepath.Join(beego.AppPath, ConfigDir, "app-" + env)
	return configFile, nil
}

func Inspect(formate string) {
	if file, err := ConfigFile(); err != nil {
		panic(err)
	} else {
		//fmt.Println(file)
		beego.LoadAppConfig(formate, file + "." + formate)
	}
}

func DetectModel() string {
	hostsConfigFile := filepath.Join(ConfigDir, HostsFile)
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
