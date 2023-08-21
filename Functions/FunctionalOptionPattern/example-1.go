package main

import "fmt"

type Options struct {
	Host    string
	Port    int
	TLS     bool
	MaxConn int
}

type OptFunc func(*Options)

func defaultOptions() Options {
	return Options{
		Host:    "localhost",
		Port:    8080,
		TLS:     false,
		MaxConn: 10,
	}
}

type Server struct {
	Options
}

func WithHost(h string) OptFunc {
	return func(opt *Options) {
		opt.Host = h
	}
}

func WithPort(p int) OptFunc {
	return func(opt *Options) {
		opt.Port = p
	}
}

func WithMaxConn(maxConnection int) OptFunc {
	return func(opt *Options) {
		opt.MaxConn = maxConnection
	}
}

func WithTLS(opt *Options) {
	opt.TLS = true
}

func NewServer(optFuncs ...OptFunc) *Server {
	opts := defaultOptions()

	for _, o := range optFuncs {
		o(&opts)
	}

	return &Server{
		Options: opts,
	}
}

func (s *Server) WithTLS() *Server {
	s.TLS = true
	return s
}

func main() {
	s := NewServer(
		WithHost("127.0.0.1"),
		WithPort(3000),
		WithTLS,
		WithMaxConn(20),
	)
	fmt.Printf("%+v\n", s)

	//
	s2 := NewServer(WithHost("127.0.0.1")).WithTLS()
	fmt.Println(s2)
}
