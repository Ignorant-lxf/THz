package THz

import (
	"net"
)

func (thz *THz) init() *THz {
	thz.route.Sort()
	thz.IRouter = nil // not allowed updates routing
	return thz
}

func (thz *THz) Stop() error               { return thz.init().srv.Shutdown() }
func (thz *THz) Run(ln net.Listener) error { return thz.init().srv.Serve(ln) }
func (thz *THz) RunTLS(ln net.Listener, cert, key string) error {
	return thz.init().srv.ServeTLS(ln, cert, key)
}
func (thz *THz) ListenAndServe(addr string) error {
	return thz.init().srv.ListenAndServe(addr)
}
