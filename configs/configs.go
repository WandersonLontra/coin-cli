package configs

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
)

var (
	BaseUrl 		string
	AccessKey 		string
	CacheFile = "cache.json"
	CacheDir = "./cache_file/"
)

func init() {
	viper.SetConfigName("configs")
	viper.SetConfigType("env")
	viper.AddConfigPath(".")

	err := viper.ReadInConfig()

	if err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			fmt.Fprintf(os.Stderr, "Error reading config file: %s\n", err)
			os.Exit(1)
		}
	}
	viper.AutomaticEnv();

	BaseUrl = viper.GetString("BASE_URL");
	AccessKey = viper.GetString("ACCESS_KEY");

	if BaseUrl == "" {
		fmt.Fprintln(os.Stderr, "Missing required BASE_URL environment variable")
		os.Exit(1)
	}

	if AccessKey == "" {
		fmt.Fprintln(os.Stderr, "Missing required ACCESS_KEY environment variable")
		os.Exit(1)
	}
}