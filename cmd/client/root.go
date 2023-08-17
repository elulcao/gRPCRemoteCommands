package client

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// ClientCmd new release command
var ClientCmd = &cobra.Command{
	Use:     "client",
	Short:   "Start gRPC client",
	Long:    `Start gRPC client for executing a command in remote host, TLS = true`,
	Example: "gRPCRemoteCommands client --command hostname -h localhost -p 50051 -c /tmp/config.yaml",
	Run: func(cmd *cobra.Command, args []string) {
		_ = args
		err := clientMain(cmd)
		if err != nil {
			panic(err)
		}
	},
}

func init() {
	ClientCmd.Flags().String("command", "hostname", "Command to execute in remote host") // Command to execute in remote host
	ClientCmd.Flags().String("host", "localhost", "Host to connect to")                  // Host to connect to
	ClientCmd.Flags().String("port", "50051", "Port to connect to")                      // Port to connect to

	_ = viper.BindPFlag("command", ClientCmd.Flags().Lookup("command"))
	_ = viper.BindPFlag("host", ClientCmd.Flags().Lookup("host"))
	_ = viper.BindPFlag("port", ClientCmd.Flags().Lookup("port"))
}
