package server

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// NewSubCmdServer new release command
func NewSubCmdServer() *cobra.Command {
	serverCmd := &cobra.Command{
		Use:     "server",
		Short:   "Start gRPC server",
		Long:    "Start gRPC server, TLS = true",
		Example: "gRPCRemoteCommands server --host localhost -p 50051 -c /tmp/config.yaml",
		Run: func(cmd *cobra.Command, args []string) {
			err := serverMain(cmd)
			if err != nil {
				panic(err)
			}
		},
	}

	serverCmd.Flags().String("host", "localhost", "Host address") // Host address
	viper.BindPFlag("host", serverCmd.Flags().Lookup("host"))
	serverCmd.Flags().String("port", "50051", "Port to listen on") // Port to listen on
	viper.BindPFlag("port", serverCmd.Flags().Lookup("port"))

	return serverCmd
}
