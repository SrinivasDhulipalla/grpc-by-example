package main

import (
	"crypto/tls"
	"crypto/x509"
	"io/ioutil"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

func main() {
	tlsKey := "/path/to/tls.key"
    tlsCert := "/path/to/tls.crt"
	tlsCA := "/path/to/ca.crt"
	
	certificate, err := tls.LoadX509KeyPair(tlsCert, tlsKey)
	if err != nil {
		panic(err.Error())
	}

	certPool := x509.NewCertPool()
	bs, err := ioutil.ReadFile(tlsCA)
	if err != nil {
		panic(err.Error())
	}

	ok := certPool.AppendCertsFromPEM(bs)
	if !ok {
		panic("failed to append certs")
	}

	transportCreds := credentials.NewTLS(&tls.Config{
		ClientAuth:   tls.RequireAndVerifyClientCert,
		Certificates: []tls.Certificate{certificate},
		ClientCAs:    certPool,
	})

	server := grpc.NewServer(grpc.Creds(transportCreds))

	//pb.RegisterMyService(server, implementation)

	listener, err := net.Listen("tcp", "0.0.0.0:8443")
	if err != nil {
		panic(err.Error())
	}

	err = server.Serve(listener)
	if err != nil {
		panic(err.Error())
	}
}