package client

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// NewSubCmdClient new release command
func NewSubCmdClient() *cobra.Command {
	clientCmd := &cobra.Command{
		Use:     "client",
		Short:   "Start gRPC client",
		Long:    `Start gRPC client for executing a command in remote host, TLS = true`,
		Example: "gRPCRemoteCommands client --command hostname -h localhost -p 50051 -c /tmp/config.yaml",
		Run: func(cmd *cobra.Command, args []string) {
			err := clientMain(cmd)
			if err != nil {
				panic(err)
			}
		},
	}

	clientCmd.Flags().String("command", "hostname", "Command to execute in remote host") // Command to execute in remote host
	viper.BindPFlag("command", clientCmd.Flags().Lookup("command"))
	clientCmd.Flags().String("host", "localhost", "Host to connect to") // Host to connect to
	viper.BindPFlag("host", clientCmd.Flags().Lookup("host"))
	clientCmd.Flags().String("port", "50051", "Port to connect to") // Port to connect to
	viper.BindPFlag("port", clientCmd.Flags().Lookup("port"))

	return clientCmd
}
