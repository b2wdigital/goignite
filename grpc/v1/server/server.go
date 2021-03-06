package gigrpc

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"
	"net"

	giconfig "github.com/b2wdigital/goignite/v2/config"
	gilog "github.com/b2wdigital/goignite/v2/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/channelz/service"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/encoding/gzip"
	"google.golang.org/grpc/reflection"
)

type Ext func(ctx context.Context) []grpc.ServerOption

type Server struct {
	server           *grpc.Server
	serviceRegistrar grpc.ServiceRegistrar
	options          *Options
}

func NewDefault(ctx context.Context, exts ...Ext) *Server {
	opt, err := DefaultOptions()
	if err != nil {
		panic(err)
	}
	return New(ctx, opt, exts...)
}

func New(ctx context.Context, opt *Options, exts ...Ext) *Server {

	logger := gilog.FromContext(ctx)

	err := gzip.SetLevel(9)
	if err != nil {
		logger.Fatalf("could not set level: %s", err.Error())
	}

	var s *grpc.Server

	var serverOptions []grpc.ServerOption

	if giconfig.Bool(tlsEnabled) {

		// Load the certificates from disk
		certificate, err := tls.LoadX509KeyPair(opt.TLS.CertFile, opt.TLS.KeyFile)
		if err != nil {
			logger.Fatalf("could not load server key pair: %s", err.Error())
		}

		// Create a certificate pool from the certificate authority
		certPool := x509.NewCertPool()
		ca, err := ioutil.ReadFile(opt.TLS.CAFile)
		if err != nil {
			logger.Fatalf("could not read ca certificate: %s", err.Error())
		}

		// Append the client certificates from the CA
		if ok := certPool.AppendCertsFromPEM(ca); !ok {
			logger.Fatalf("failed to append client certs")
		}

		// Create the TLS credentials
		creds := credentials.NewTLS(&tls.Config{
			ClientAuth:   tls.RequireAndVerifyClientCert,
			Certificates: []tls.Certificate{certificate},
			ClientCAs:    certPool,
		})

		serverOptions = append(serverOptions, grpc.Creds(creds))
	}

	for _, ext := range exts {
		serverOptions = append(serverOptions, ext(ctx)...)
	}

	serverOptions = append(serverOptions, grpc.MaxConcurrentStreams(uint32(opt.MaxConcurrentStreams)))

	s = grpc.NewServer(serverOptions...)

	// grpc.InitialConnWindowSize(100),
	// grpc.InitialWindowSize(100),

	return &Server{
		server:  s,
		options: opt,
	}
}

func (s *Server) Server() *grpc.Server {
	return s.server
}

func (s *Server) ServiceRegistrar() grpc.ServiceRegistrar {
	return s.server
}

func (s *Server) Serve(ctx context.Context) {

	logger := gilog.FromContext(ctx)

	service.RegisterChannelzServiceToServer(s.server)

	// Register reflection service on gRPC server.
	reflection.Register(s.server)

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", s.options.Port))
	if err != nil {
		logger.Fatalf("failed to listen: %v", err.Error())
	}

	logger.Infof("grpc server started on port %v", s.options.Port)

	if err := s.server.Serve(lis); err != nil {
		logger.Fatalf("failed to serve: %v", err.Error())
	}

}
