package cmd

import (
	"embed"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"/pkg/log"
	"/pkg/utils"
)

var (
	cfgFile   string
	configsFS embed.FS
	migrationsFS embed.FS
	webFS     embed.FS
)

// SetEmbedFS sets the embedded filesystems from main.go.
func SetEmbedFS(configs, migrations, web embed.FS) {
	configsFS = configs
	migrationsFS = migrations
	webFS = web
}

var rootCmd = &cobra.Command{
	Use:   "",
	Short: "The website of Scalland Conultancy Services",
	Long:  ``,
}

// Execute runs the root command.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file path")
}

func initConfig() {
	u := utils.GetUtils()
	logger := log.NewLogger()

	if cfgFile != "" {
		absPath, err := filepath.Abs(cfgFile)
		if err == nil {
			viper.SetConfigFile(absPath)
		} else {
			viper.SetConfigFile(cfgFile)
		}
	} else {
		viper.AddConfigPath("./configs")
		viper.AddConfigPath(".")
		viper.SetConfigName("app")
		viper.SetConfigType("yaml")
	}

	viper.SetEnvPrefix(strings.ToUpper(""))
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	// Set defaults
	viper.SetDefault("App.ServerPort", "0")
	viper.SetDefault("App.JWTSecret", "change-me-in-production")
	viper.SetDefault("App.Timezone", "")

	if err := viper.ReadInConfig(); err == nil {
		logger.Infof("Using config: %s", viper.ConfigFileUsed())
	}

	if err := viper.Unmarshal(u.Config); err != nil {
		logger.Fatalf("Unable to decode config: %s", err.Error())
	}

	u.Logger = logger
	u.SetTimeZone(u.Config.App.Timezone)

	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		logger.Infof("Config changed: %s", e.Name)
		if err := viper.Unmarshal(u.Config); err != nil {
			logger.Errorf("Config reload failed: %s", err.Error())
		}
	})
}
