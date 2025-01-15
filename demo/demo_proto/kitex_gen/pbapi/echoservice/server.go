// Code generated by Kitex v0.9.1. DO NOT EDIT.
package echoservice

import (
	server "github.com/cloudwego/kitex/server"
	pbapi "github.com/py/biz-demo/gomall/demo/demo_proto/kitex_gen/pbapi"
)

// NewServer creates a server.Server with the given handler and options.
func NewServer(handler pbapi.EchoService, opts ...server.Option) server.Server {
	var options []server.Option

	options = append(options, opts...)

	svr := server.NewServer(options...)
	if err := svr.RegisterService(serviceInfo(), handler); err != nil {
		panic(err)
	}
	return svr
}

func RegisterService(svr server.Server, handler pbapi.EchoService, opts ...server.RegisterOption) error {
	return svr.RegisterService(serviceInfo(), handler, opts...)
}
