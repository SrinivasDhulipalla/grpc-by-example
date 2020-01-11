package main

import (
	"crypto/tls"
	"crypto/x509"
	"io/ioutil"

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
		Certificates: []tls.Certificate{certificate},
		RootCAs:      certPool,
	})

	cc, err := grpc.Dial("localhost:8443", grpc.WithTransportCredentials(transportCreds))
	if err != nil {
		panic(err.Error())
	}

	// client := pb.NewMyServiceClient(cc)
}
