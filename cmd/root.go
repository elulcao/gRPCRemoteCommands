package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"

	"github.com/elulcao/gRPCRemoteCommands/cmd/client"
	"github.com/elulcao/gRPCRemoteCommands/cmd/server"
)

var cfgFile string
var vpr = viper.New()

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:     "gRPCRemoteCommands",
	Short:   "CLI for executing remote commands over gRPC",
	Long:    `CLI for executing remote commands over gRPC`,
	Version: "0.1.0",
	Example: "gRPCRemoteCommands server",
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

// init is called before main()
func init() {
	cobra.OnInitialize(initConfig)
	// Cobra supports persistent flags, which, if defined here, will be global for your application.
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.gRPCRemoteCommands.yaml)")

	rootCmd.AddCommand(server.ServerCmd)
	rootCmd.AddCommand(client.ClientCmd)
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		vpr.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		// Search config in home directory with name ".gRPCRemoteCommands" (without extension).
		vpr.AddConfigPath(home)
		vpr.SetConfigType("yaml")
		vpr.SetConfigName(".gRPCRemoteCommands")
	}

	vpr.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := vpr.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "Using config file:", vpr.ConfigFileUsed())

		// Assign values from viper to cobra flags
		viperToCobraFlags(server.ServerCmd, vpr)
		viperToCobraFlags(client.ClientCmd, vpr)
	}
}

// viperToCobraFlags assigns vpr values to Cobra flags.
func viperToCobraFlags(cmd *cobra.Command, vpr *viper.Viper) {
	cmd.Flags().VisitAll(func(f *pflag.Flag) {
		if vpr.IsSet(f.Name) {
			_ = cmd.Flags().Set(f.Name, vpr.GetString(f.Name))
		}
	})
}
