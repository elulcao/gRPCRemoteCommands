package client

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"log"
	"time"

	"github.com/spf13/cobra"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"

	ce "github.com/elulcao/gRPCRemoteCommands/certificates"
	pb "github.com/elulcao/gRPCRemoteCommands/proto"
)

const (
// Define the path for the certificates if not embedded
//clientCertFile   = "scripts/certificates/client-cert.pem"
//clientKeyFile    = "scripts/certificates/client-key.pem"
//clientCACertFile = "scripts/certificates/ca-cert.pem"
//port = "50051"
)

// loadTLSCredentials loads TLS credentials from files.
func loadTLSCredentials() (credentials.TransportCredentials, error) {
	clientCertFile := ce.GetCertByte("client-cert")
	clientKeyFile := ce.GetCertByte("client-key")
	clientCACertFile := ce.GetCertByte("ca-cert")

	// Load certificate of the CA who signed client's certificate
	//pemServerCA, err := ioutil.ReadFile(clientCACertFile)
	//if err != nil {
	//	return nil, err
	//}

	certPool := x509.NewCertPool()
	if !certPool.AppendCertsFromPEM(clientCACertFile) {
		return nil, fmt.Errorf("failed to add client CA's certificate")
	}

	// Load client's certificate and private key
	//clientCert, err := tls.LoadX509KeyPair(clientCertF.Name(), clientKeyF.Name())
	clientCert, err := tls.X509KeyPair(clientCertFile, clientKeyFile)
	if err != nil {
		return nil, err
	}

	// Create the credentials and return it
	config := &tls.Config{
		Certificates:       []tls.Certificate{clientCert},
		RootCAs:            certPool,
		InsecureSkipVerify: true,
	}

	return credentials.NewTLS(config), nil
}

// clientMain is the entry point for the client.
func clientMain(cmd *cobra.Command) error {
	h, _ := cmd.Flags().GetString("host")
	c, _ := cmd.Flags().GetString("command")
	p, _ := cmd.Flags().GetString("port")

	log.Printf("dial server %s, TLS = %t", h, true)

	tlsCredentials, err := loadTLSCredentials()
	if err != nil {
		log.Fatal("cannot load TLS credentials: ", err)
		return err
	}

	transportOption := grpc.WithTransportCredentials(tlsCredentials)

	conn, err := grpc.Dial(
		fmt.Sprintf("%v:%v", h, p),
		transportOption,
	)
	if err != nil {
		log.Fatal("cannot dial server: ", err)
		return err
	}
	defer conn.Close()

	client := pb.NewRemoteCommandClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	r, err := client.ExecuteCommand(ctx, &pb.CommandRequest{
		Cmd: c,
	})
	if err != nil {
		log.Fatalf("could not execute command: %v", err)
		return err
	}
	log.Printf("Output: \n %s", r.Out)

	return nil
}
