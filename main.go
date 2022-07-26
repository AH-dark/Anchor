package main

import (
	"flag"
	"github.com/AH-dark/Anchor/bootstrap"
	"github.com/AH-dark/Anchor/pkg/conf"
	"github.com/AH-dark/Anchor/pkg/utils"
	"github.com/AH-dark/Anchor/routers"
)

var (
	confPath string
)

func init() {
	flag.StringVar(&confPath, "c", utils.RelativePath("config.yaml"), "配置文件路径")
	flag.Parse()

	bootstrap.Init(confPath)
}

func main() {
	r := routers.InitRouter()

	utils.Log().Info("准备监听 %s", conf.Config.System.Listen)
	err := r.Run(conf.Config.System.Listen)
	if err != nil {
		utils.Log().Panic("监听 %s 时失败，%s", conf.Config.System.Listen, err)
		return
	}
}
