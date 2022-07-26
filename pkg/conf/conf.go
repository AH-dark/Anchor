package conf

import (
	"github.com/AH-dark/Anchor/pkg/utils"
	"gopkg.in/yaml.v3"
	"io"
	"os"
)

func Init(path string) {
	var err error

	// 文件不存在时自动生成
	if path == "" || !utils.Exists(path) {
		defaultConf, err := yaml.Marshal(&Config)

		// 创建初始配置文件
		confContent := utils.Replace(map[string]string{
			"{SessionSecret}": utils.RandStringRunes(64),
			"{HashIDSalt}":    utils.RandStringRunes(64),
		}, string(defaultConf))
		f, err := utils.CreatNestedFile(path)
		if err != nil {
			utils.Log().Panic("无法创建配置文件, %s", err)
		}

		// 写入配置文件
		_, err = f.WriteString(confContent)
		if err != nil {
			utils.Log().Panic("无法写入配置文件, %s", err)
		}

		f.Close()
	}

	// 打开配置文件
	f, err := os.Open(path)
	if err != nil {
		utils.Log().Panic("无法打开配置文件, %s", err)
	}

	// 读取配置文件
	var data []byte
	buf := make([]byte, 1024)
	for {
		n, err := f.Read(buf)
		if err != nil && err != io.EOF {
			utils.Log().Panic("无法读取配置文件, %s", err)
		}

		if n == 0 {
			break
		}

		data = append(data, buf[:n]...)
	}

	// 解析配置文件
	err = yaml.Unmarshal(data, &Config)
	if err != nil {
		utils.Log().Panic("无法解析配置文件, %s", err)
	}
	defer f.Close()

	utils.Log().Debug("%v", Config)

	// 重设log等级
	if !Config.System.Debug {
		utils.Level = utils.LevelInformational
		utils.GloablLogger = nil
		utils.Log()
	}
}
