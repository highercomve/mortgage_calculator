package main

import (
	"fmt"
	"log"
	"os"
	"path"

	"github.com/highercomve/mortgage_calculator/server"
	"github.com/spf13/viper"
)

func init() {
	viper.SetDefault("host", "")
	viper.SetDefault("port", "9090")
	viper.SetDefault("debug", false)

	// Find home directory.
	home, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}

	p, err := os.Executable()
	if err != nil {
		log.Fatal(err)
	}

	// Search config in home directory with name ".config"
	viper.AddConfigPath(home)
	viper.AddConfigPath(path.Dir(p))
	viper.SetConfigType("yaml")
	viper.SetConfigName(".config")

	viper.AutomaticEnv()
	if err := viper.ReadInConfig(); err != nil {
		log.Println(err)
	}
}

// @title Mortgage Calculator API
// @version 1.0
// @description Mortgage Calculator API.
// @termsOfService http://swagger.io/terms/

// @contact.name Sergio Marin
// @contact.url https://highercomve.github.io
// @contact.email @highercomve

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @BasePath /
func main() {
	serverAddres := fmt.Sprintf("%s:%s", viper.GetString("host"), viper.GetString("port"))
	server.Start(serverAddres)
}
