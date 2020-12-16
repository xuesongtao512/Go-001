// @Desc:
// @CreateTime: 2020/10/11
package config

import (
    "encoding/json"
    "fmt"
    "log"

    "github.com/kylelemons/go-gypsy/yaml"
)

var (
    Configs *AllConf // 对外出口
)

// 保存所有配置
type AllConf struct {
    MysqlConf
    LogConf
}

// mysql 配置
type MysqlConf struct {
    MysqlHost string `json:"mysqlHost"`
    Username  string `json:"username"`
    Password  string `json:"password"`
    DbName    string `json:"dbname"`
}

// 日志配置
type LogConf struct {
    LogPath      string `json:"logPath"`
    LogSize      uint8  `json:"logSize,string"`
    MaxSaveFiles uint8  `json:"maxSaveFiles,string"`
    MaxSaveDays  uint8  `json:"maxSaveDays,string"`
    IsProduction bool   `json:"isProduction,string"`
}

// GetInfo 返回conf 详情
func InitConf(projectDir string) {
    configFile := fmt.Sprintf("%s\\config\\config.yaml", projectDir)
    config, err := yaml.ReadFile(configFile)
    if err != nil {
        panic(err)
    }
    // 将配置通过序列化来解析
    configByte, err := json.Marshal(config.Root)
    var (
        mysqlConf MysqlConf
        logConf LogConf
    )
    if err = json.Unmarshal(configByte, &mysqlConf); err != nil {
        log.Println("json.Unmarshal(configByte, &mysqlConf) is failed, err: ", err)
    }
    if err = json.Unmarshal(configByte, &logConf); err != nil {
        log.Println("json.Unmarshal(configByte, &logConf) is failed, err: ", err)
    }
    // mylog 的 path 还需要拼接
    logConf.LogPath = projectDir + logConf.LogPath
    log.Printf("mysqlConf: %v", mysqlConf)
    log.Printf("logConf: %v", logConf)
    Configs = &AllConf{
        mysqlConf,
        logConf,
    }
}
