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
    Dates         bool `json:"defaultDates"`
    ReverseTags   bool `json:"defaultReverseTags"`
}

func getFile() (*os.File, error) {
    if runtime.GOOS == "linux" || runtime.GOOS == "darwin" {
        home, _ := os.UserHomeDir()
        return os.OpenFile(home+"/.lcl.config", os.O_CREATE|os.O_RDWR, 0644)
    }
    return os.OpenFile("~\\AppData\\Roaming\\lcl.config", os.O_CREATE|os.O_RDWR, 0644)
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
        data, _ = json.MarshalIndent(cfg, "", "  ")
        file.Write(data)
        Config = ConfigData{}
    }
    Config = cfg
}
