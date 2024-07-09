package bootstrap

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"log"
)

// LoadConfig ...
func LoadConfig(path string) error {
	// 根据Flag的值设置Viper的配置文件路径
	resolveRealPath(path)

	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {             // Handle errors reading the config file
		panic(fmt.Errorf("fatal error config file: %w", err))
	}

	// 监控配置文件，并热加载
	watchConfig()

	return nil
}

// 如果未传递配置文件路径将使用约定的环境配置文件
func resolveRealPath(path string) {
	if path != "" {
		viper.SetConfigFile(path)
	} else {
		// 如果没有提供Flag，则使用默认的配置文件路径
		viper.SetConfigName("config")    // name of config file (without extension)
		viper.SetConfigType("yaml")      // REQUIRED if the config file does not have the extension in the name
		viper.AddConfigPath("./configs") // path to look for the config file in
	}
}

// 监控配置文件变动
// 注意：有些配置修改后，及时重新加载也要重新启动应用才行，比如端口
func watchConfig() {
	viper.WatchConfig()
	viper.OnConfigChange(func(in fsnotify.Event) {
		log.Printf("Configuration file changed: %s, reload it", in.Name)
		// 忽略错误
		err := LoadConfig(in.Name)
		if err != nil {
			return
		}
	})
}
