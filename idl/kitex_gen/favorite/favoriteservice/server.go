// Code generated by Kitex v0.10.1. DO NOT EDIT.
package favoriteservice

import (
	server "github.com/cloudwego/kitex/server"
	favorite "tiktok-simple/idl/kitex_gen/favorite"
)

// NewServer creates a server.Server with the given handler and options.
func NewServer(handler favorite.FavoriteService, opts ...server.Option) server.Server {
	var options []server.Option

	options = append(options, opts...)

	svr := server.NewServer(options...)
	if err := svr.RegisterService(serviceInfo(), handler); err != nil {
		panic(err)
	}
	return svr
}

func RegisterService(svr server.Server, handler favorite.FavoriteService, opts ...server.RegisterOption) error {
	return svr.RegisterService(serviceInfo(), handler, opts...)
}
