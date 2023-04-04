package server

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// ServerCmd new release command
var ServerCmd = &cobra.Command{
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

func init() {
	ServerCmd.Flags().String("host", "localhost", "Host address")  // Host address
	ServerCmd.Flags().String("port", "50051", "Port to listen on") // Port to listen on

	viper.BindPFlag("host", ServerCmd.Flags().Lookup("host"))
	viper.BindPFlag("port", ServerCmd.Flags().Lookup("port"))
}
