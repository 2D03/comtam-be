package conf

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"os"
)

type config struct {
	MainDBName string
	LogDBName  string
	MainAuthDB string
	LogAuthDB  string
}

var Config *config

func init() {
	env := os.Getenv("env")
	switch env {
	case "local":
		Config = &config{
			MainDBName: "comtam_local_api",
			LogDBName:  "comtam_local_api_log",
			MainAuthDB: "admin",
			LogAuthDB:  "admin",
		}
	case "dev":
		Config = &config{
			MainDBName: "comtam_dev_api",
			LogDBName:  "comtam_dev_api_log",
			MainAuthDB: "admin",
			LogAuthDB:  "admin",
		}
	case "prd":
		Config = &config{
			MainDBName: "comtam_prd_api",
			LogDBName:  "comtam_prd_api_log",
			MainAuthDB: "admin",
			LogAuthDB:  "admin",
		}
	}
}

func GetConfigDBMap() (map[string]string, error) {
	var configMap map[string]string
	configStr := os.Getenv("config")
	decoded, err := base64.URLEncoding.DecodeString(configStr)
	if err != nil {
		fmt.Println("[Parse config] Convert B64 config string error: " + err.Error())
		return nil, err
	}
	err = json.Unmarshal(decoded, &configMap)
	if err != nil {
		fmt.Println("[Parse config] Parse JSON with config string error: " + err.Error())
		return nil, err
	}

	return configMap, err
}
