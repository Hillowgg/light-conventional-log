package config

import (
    "encoding/json"
    "io"
    "os"
    "runtime"
)

var Config ConfigData

type ConfigData struct {
    IncludeScopes bool `json:"defaultIncludeScopes"`
    Interactive   bool `json:"defaultInteractive"`
}

func getFile() (*os.File, error) {
    if runtime.GOOS == "linux" || runtime.GOOS == "darwin" {
        return os.OpenFile("~/.lcl.config", os.O_CREATE|os.O_RDONLY, 0666)
    }
    return os.OpenFile("~\\AppData\\Roaming\\lcl.config", os.O_CREATE|os.O_RDONLY, 0666)
}

func LoadConfig() {
    file, err := getFile()
    if err != nil {
        panic(err)
    }
    defer file.Close()
    cfg := ConfigData{}

    data, err := io.ReadAll(file)
    if err != nil {
        panic(err)
    }
    err = json.Unmarshal(data, &cfg)
    if err != nil {
        data, _ = json.Marshal(cfg)
        file.Write(data)
        Config = ConfigData{}
    }
    Config = cfg
}
