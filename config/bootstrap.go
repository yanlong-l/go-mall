package config

import (
	"bytes"
	"embed"
	"fmt"
	"os"
	"slices"

	"github.com/spf13/viper"
)

//go:embed *.yaml
var configs embed.FS

func init() {
	// 读取env 环境变量
	env, exist := os.LookupEnv("env")
	fmt.Println("env", env)
	if !exist {
		panic("Please check your environment variable values")
	}
	ok := slices.Contains([]string{"dev", "prod", "test"}, env)
	if !ok {
		panic("Illegal value of environment variable Env")
	}
	vp := viper.New()
	// 根据环境变量的值，从configs中读取文件内容
	data, err := configs.ReadFile(fmt.Sprintf("application.%s.yaml", env))
	if err != nil {
		panic(err)
	}
	vp.SetConfigType("yaml")
	vp.ReadConfig(bytes.NewReader(data))

	vp.UnmarshalKey("app", &App)
	vp.UnmarshalKey("database", &Database)
}
