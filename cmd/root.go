package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/cedi/kkpctl/pkg/config"
	"github.com/mitchellh/go-homedir"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

const (
	envPrefix = "KKPCTL"
)

var (
	// Config holds the global configuration for kkpctl
	Config config.Config

	// ConfigPath is the path to our configuration file on disk
	ConfigPath string

	apiToken   string
	baseURL    string
	outputType string
	sortBy     string
	snail      *viper.Viper
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "kkpctl [(-o|--output=)text|json|yaml] [(--sort=name|date)]",
	Short: "A CLI for interacting with Kubermatic Kubernetes Platform.",
	Long:  `This is a CLI for interacting with the REST API of Kubermatic Kubernetes Platform.`,
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		// You can bind cobra and viper in a few locations, but PersistencePreRunE on the root command works well
		return initConfig(cmd)
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {

	var err error
	Config, err = config.Read(ConfigPath)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	err = rootCmd.Execute()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	// Find home directory.
	home, err := homedir.Dir()
	if err != nil {
		fmt.Println("Failed to find home directory: " + err.Error())
		os.Exit(1)
	}
	rootCmd.PersistentFlags().StringVar(&ConfigPath, "config", home+"/.config/kkpctl/config.v2.yaml", "The Path to the configuration file")

	rootCmd.PersistentFlags().StringVarP(&apiToken, "bearer", "t", "", "API token for authenticating with Kubermatic API.")
	viper.BindPFlag("bearer", rootCmd.PersistentFlags().Lookup("bearer"))

	rootCmd.PersistentFlags().StringVarP(&baseURL, "url", "u", "", "The KKP URL to use")
	viper.BindPFlag("url", rootCmd.PersistentFlags().Lookup("url"))

	rootCmd.PersistentFlags().StringVarP(&outputType, "output", "o", "text", "The output type to use")
	viper.BindPFlag("output", rootCmd.PersistentFlags().Lookup("output"))
	rootCmd.RegisterFlagCompletionFunc("output", func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		return []string{"text", "json", "yaml"}, cobra.ShellCompDirectiveDefault
	})

	rootCmd.PersistentFlags().StringVar(&sortBy, "sort", "name", "Sort text output by which attribute (\"name\" or \"date\")")
	viper.BindPFlag("sort", rootCmd.PersistentFlags().Lookup("sort"))
	rootCmd.RegisterFlagCompletionFunc("sort", func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		return []string{"name", "date"}, cobra.ShellCompDirectiveDefault
	})
}

// initConfig reads in config file and ENV variables if set.
func initConfig(cmd *cobra.Command) error {
	snail = viper.New()

	snail.SetConfigType("yaml")

	// Find home directory.
	home, err := homedir.Dir()
	if err != nil {
		return errors.Wrap(err, "Failed to find home directory")
	}

	snail.AddConfigPath(home + "/.config/kkpctl")
	snail.SetConfigName("config.yaml")

	// If a config file is found, read it in.
	err = snail.ReadInConfig()
	if err != nil {
		// It's okay if there isn't a config file
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			return errors.Wrap(err, "Error reading config file")
		}
		return errors.Wrap(err, "No config file found")
	}

	// When we bind flags to environment variables expect that the
	// environment variables are prefixed, e.g. a flag like --url
	// binds to an environment variable KKP_URL. This helps
	// avoid conflicts.
	snail.SetEnvPrefix(envPrefix)

	// Bind to environment variables
	// Works great for simple config names, but needs help for names
	// like --favorite-color which we fix in the bindFlags function
	snail.AutomaticEnv()

	// Bind the current command's flags to viper
	bindFlags(cmd, snail)

	return nil
}

func bindFlags(cmd *cobra.Command, v *viper.Viper) {
	cmd.Flags().VisitAll(func(f *pflag.Flag) {
		// Environment variables can't have dashes in them, so bind them to their equivalent
		// keys with underscores, e.g. --favorite-color to STING_FAVORITE_COLOR
		if strings.Contains(f.Name, "-") {
			envVarSuffix := strings.ToUpper(strings.ReplaceAll(f.Name, "-", "_"))
			v.BindEnv(f.Name, fmt.Sprintf("%s_%s", envPrefix, envVarSuffix))
		}

		// Apply the viper config value to the flag when the flag is not set and viper has a value
		if !f.Changed && v.IsSet(f.Name) {
			val := v.Get(f.Name)
			cmd.Flags().Set(f.Name, fmt.Sprintf("%v", val))
		}
	})
}
