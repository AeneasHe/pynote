package utils

func GetConfigPath() string {

	// 先检查开始目录下面是否有配置文件
	configPath := GetStartPath() + "/config.json"

	exist := PathExists(configPath)
	if exist {
		return configPath
	}

	// 再检查可执行文件所在目录下是否有配置文件
	exePath := GetExePath()

	configPath = exePath + "/config.json"

	return configPath
}
