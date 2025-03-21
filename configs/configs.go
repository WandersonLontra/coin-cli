package configs

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/viper"
)

var (
	BaseUrl 		string
	AccessKey 		string
	TTLCache		float64
	IsDevMode		bool
	CacheDir      	string
	CacheFile = "cache.json"
)

func init() {
	exePath, err := os.Executable()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading config file path: %s\n", err)
		os.Exit(1)
	}
	exeDir := filepath.Dir(exePath)

	viper.SetConfigName("configs")
	viper.SetConfigType("env")
	viper.AddConfigPath(".")
	viper.AddConfigPath(exeDir)

	viper.SetDefault("TTL_CACHE_IN_HOURS", float64(12))
	viper.SetDefault("DEV_MODE", false)

	err = viper.ReadInConfig()

	if err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			fmt.Fprintf(os.Stderr, "Error reading config file: %s\n", err)
			os.Exit(1)
		}
	}
	viper.AutomaticEnv();

	BaseUrl = viper.GetString("BASE_URL");
	AccessKey = viper.GetString("ACCESS_KEY");
	TTLCache = viper.GetFloat64("TTL_CACHE_IN_HOURS");
	IsDevMode = viper.GetBool("DEV_MODE");

	if BaseUrl == "" {
		fmt.Fprintln(os.Stderr, "Missing required BASE_URL environment variable")
		os.Exit(1)
	}

	if AccessKey == "" {
		fmt.Fprintln(os.Stderr, "Missing required ACCESS_KEY environment variable")
		os.Exit(1)
	}

	if IsDevMode {
		CacheDir = "./cache_file";
	} else {
		CacheDir = exeDir + "/cache_file";
	}
}