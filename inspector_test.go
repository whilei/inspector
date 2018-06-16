package inspector

import (
	"testing"
)

func TestParseHostsConfig(t *testing.T) {
	raw := []byte(`{"dev":["host1.local"],"product":["abc.com"]}`)
	var config map[string][]string
	config = parseHostsConfig(raw)
	if config == nil {
		t.Error("parse config file error")
	}
	if config["dev"] == nil {
		t.Error("failed")
	}
	devHosts := config["dev"]
	if len(devHosts) != 1 {
		t.Error("format error")
	}
	if devHosts[0] != "host1.local" {
		t.Error("error")
	}
	// Output: done
}

func TestDetectModel(t *testing.T) {
	ConfigDir = "config"
	HostsFile = "hosts.json"
	env := DetectModel()
	if env == "" {
		t.Error("detect error")
	}
	if env != "dev" {
		t.Error("wrong env")
	}
	//Output done
}

func TestDetectModel_2(t *testing.T) {
	ConfigDir = "config"
	HostsFile = "hosts1.json"
	env := DetectModel()
	if env == "" {
		t.Error("detect error")
	}
	if env != "product" {
		t.Error("wrong env")
	}
	//Output done
}
