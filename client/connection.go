package main

import (
	"crypto/tls"
	"flag"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	serverAddr   = flag.String("server", "localhost:50050", "gRPC Server address (host:port)")
	serverHost   = flag.String("server-host", "localhost", "Host name to which server IP should resolve")
	insecureFlag = flag.Bool("insecure", true, "Skip SSL validation? [false]")
	skipVerify   = flag.Bool("skip-verify", false, "Skip server hostname verification in SSL validation [false]")
)

func init() {
	flag.Parse()
}

// Connection creates a new gRPC connection to the server.
// host should be of the form domain:port, e.g., example.com:443
func Connection() (*grpc.ClientConn, error) {
	var opts []grpc.DialOption
	if *serverAddr == "" {
		log.Fatal("-server is empty")
	}
	if *serverHost != "" {
		opts = append(opts, grpc.WithAuthority(*serverHost))
	}
	if *insecureFlag {
		opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	} else {
		cred := credentials.NewTLS(&tls.Config{
			InsecureSkipVerify: *skipVerify,
		})
		opts = append(opts, grpc.WithTransportCredentials(cred))
	}
	return grpc.Dial(*serverAddr, opts...)
}
