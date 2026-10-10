package jarvis

import (
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	cfgFile string
	rootCmd = &cobra.Command{
		Use:   "jarvis",
		Short: "An AI harness",
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			return initializeConfig(cmd)
		},
	}
)

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func init() {
	// Add our --config flag
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "$HOME/config.yaml")
}

func initializeConfig(cmd *cobra.Command) error {
	viper.SetEnvPrefix("JARVIS")

	// Allow to use nested keys (e.g. `JARVIS_LOG_LEVEL`)
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "*", "-", "*"))
	viper.AutomaticEnv()

	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		// Search for a config file in default locations
		home, err := os.UserHomeDir()
		// Only panic if we can't get the /home dir
		cobra.CheckErr(err)

		// Search for a config file with the name "config"
		viper.AddConfigPath(".")
		viper.AddConfigPath(home + "/.jarvis")
		viper.SetConfigName("config")
		viper.SetConfigType("yaml")
	}

	// If a config file is found, read it in. We use a robust error check
	// to ignore "file not found" errors, but panic on any other error
	if err := viper.ReadInConfig(); err != nil {
		// We don't care if the config file doesn't exist
		if _, ok := errors.AsType[viper.ConfigFileNotFoundError](err); !ok {
			return err
		}
	}

	err := viper.BindPFlags(cmd.Flags())
	if err != nil {
		return err
	}

	// Use it while debugging
	// TODO: make it available only when DEBUG level activated
	fmt.Println("Configuration initialized. Using config file: ", viper.ConfigFileUsed())
	return nil
}
