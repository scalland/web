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

	"scalland/pkg/log"
	"scalland/pkg/utils"
)

var (
	cfgFile      string
	migrationsFS embed.FS
	webFS        embed.FS
)

// SetEmbedFS sets the embedded filesystems from main.go.
func SetEmbedFS(migrations, web embed.FS) {
	migrationsFS = migrations
	webFS = web
	// Also wire into the utils singleton so MigrateDB can access the files
	utils.GetUtils().MigrationsFS = migrations
}

var rootCmd = &cobra.Command{
	Use:   "",
	Short: "The website of Scalland Conultancy Services",
	Long:  `The website of Scalland Conultancy Services and its Enterprise Portal`,
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
	rootCmd.PersistentFlags().StringVarP(&cfgFile, "config", "c", "app.yml", "config file path")
}

func initConfig() {
	u := utils.GetUtils()
	logger := log.NewLogger()
	u.Logger = logger
	logger.Debugf("Starting initConfig...")

	if cfgFile != "" {
		absPath, err := filepath.Abs(cfgFile)
		if err == nil {
			viper.SetConfigFile(absPath)
		} else {
			viper.SetConfigFile(cfgFile)
		}
	} else {
		viper.AddConfigPath(".")
		logger.Debugf("Added config path: %s", ".")
		viper.SetConfigName("app")
		logger.Debugf("Set config name: %s", "app")
		viper.SetConfigType("yaml")
		logger.Debugf("Set config type: %s", "yaml")
	}

	viper.SetEnvPrefix(strings.ToUpper(""))
	logger.Debugf("Set env prefix: %s", strings.ToUpper(""))
	viper.AutomaticEnv()
	logger.Debugf("Set automatic env")
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	logger.Debugf("Set env key replacer: %s", "."+"_")

	// Set defaults
	viper.SetDefault("App.ServerPort", "8080")
	logger.Debugf("Set default server port: %s", "8080")
	viper.SetDefault("App.JWTSecret", "change-me-in-production")
	logger.Debugf("Set default JWT secret: %s", "change-me-in-production")
	viper.SetDefault("App.Timezone", "")
	logger.Debugf("Set default timezone: %s", "")

	if err := viper.ReadInConfig(); err == nil {
		logger.Infof("Using config: %s", viper.ConfigFileUsed())
		logger.Debugf("Using config: %s", viper.ConfigFileUsed())
	}

	if err := viper.Unmarshal(u.Config); err != nil {
		logger.Fatalf("Unable to decode config: %s", err.Error())
		logger.Debugf("Unable to decode config: %s", err.Error())
	}
	u.SetTimeZone(u.Config.App.Timezone)
	logger.Debugf("Set timezone: %s", u.Config.App.Timezone)

	viper.WatchConfig()
	logger.Debugf("Watching config")
	viper.OnConfigChange(func(e fsnotify.Event) {
		logger.Infof("Config changed: %s", e.Name)
		logger.Debugf("Config changed: %s", e.Name)
		if err := viper.Unmarshal(u.Config); err != nil {
			logger.Errorf("Config reload failed: %s", err.Error())
			logger.Debugf("Config reload failed: %s", err.Error())
		}
	})
}
