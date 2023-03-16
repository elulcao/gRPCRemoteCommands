package server

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"log"
	"net"
	"os/exec"
	"strings"

	ce "github.com/elulcao/gRPCRemoteCommands/certificates"
	pb "github.com/elulcao/gRPCRemoteCommands/proto"

	"github.com/spf13/cobra"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/reflection"
)

const (
// Define the path for the certificates if not embedded
// serverCertFile   = "scripts/certificates/server-cert.pem"
// serverKeyFile    = "scripts/certificates/server-key.pem"
// serverCACertFile = "scripts/certificates/ca-cert.pem"
// port = "50051"
)

// server is used to impleeent the server
type server struct {
	pb.UnimplementedRemoteCommandServer
}

// ExecuteCommand implements the ExecuteCommand RPC
func (s *server) ExecuteCommand(ctx context.Context, in *pb.CommandRequest) (*pb.CommandResponse, error) {
	cmd := strings.Fields(in.Cmd)[0]   // get the first word
	args := strings.Fields(in.Cmd)[1:] // get the rest of the words

	c := exec.Command(cmd, args...)
	stdout, err := c.CombinedOutput()

	log.Print("Execute command: ", in.Cmd)
	log.Print("Output  command: ", string(stdout))

	return &pb.CommandResponse{
		Out: string(stdout),
	}, err
}

// loadTLSConfig loads TLS configuration from the given parameters.
func loadTLSCredentials() (credentials.TransportCredentials, error) {
	serverCertFile := ce.GetCertByte("server-cert")
	serverKeyFile := ce.GetCertByte("server-key")
	serverCACertFile := ce.GetCertByte("ca-cert")

	// Load certificate of the CA who signed client's certificate
	//pemClientCA, err := ioutil.ReadFile(serverCACertFile)
	//if err != nil {
	//	return nil, err
	//}

	certPool := x509.NewCertPool()
	if !certPool.AppendCertsFromPEM(serverCACertFile) {
		return nil, fmt.Errorf("failed to add client CA's certificate")
	}

	// Load server's certificate and private key
	//serverCert, err := tls.LoadX509KeyPair(serverCertF.Name(), serverKeyF.Name())
	serverCert, err := tls.X509KeyPair(serverCertFile, serverKeyFile)
	if err != nil {
		return nil, err
	}

	// Create the credentials and return it
	config := &tls.Config{
		Certificates: []tls.Certificate{serverCert},
		ClientAuth:   tls.RequireAndVerifyClientCert,
		ClientCAs:    certPool,
	}

	return credentials.NewTLS(config), nil
}

// runGRPCServer runs the gRPC server on the given listener.
func runGRPCServer(listener net.Listener) error {
	tlsCredentials, err := loadTLSCredentials()
	if err != nil {
		return fmt.Errorf("cannot load TLS credentials: %w", err)
	}

	grpcServer := grpc.NewServer(grpc.Creds(tlsCredentials))
	pb.RegisterRemoteCommandServer(grpcServer, &server{})
	reflection.Register(grpcServer)

	log.Printf("Start GRPC server at %s, TLS = %t", listener.Addr().String(), true)

	return grpcServer.Serve(listener)
}

// serverMain is the entry point for the server.
func serverMain(cmd *cobra.Command) error {
	h, _ := cmd.Flags().GetString("host")
	p, _ := cmd.Flags().GetString("port")

	listener, err := net.Listen("tcp", fmt.Sprintf("%v:%v", h, p))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
		return err
	}

	err = runGRPCServer(listener)
	if err != nil {
		log.Fatal("cannot start server: ", err)
		return err
	}

	return nil
}
