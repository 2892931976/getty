/******************************************************
# DESC    : getty client/server options
# AUTHOR  : Alex Stocks
# VERSION : 1.0
# LICENCE : Apache License 2.0
# EMAIL   : alexstocks@foxmail.com
# MOD     : 2018-03-17 21:12
# FILE    : options.go
******************************************************/

package getty

/////////////////////////////////////////
// Server Options
/////////////////////////////////////////

type ServerOption func(*ServerOptions)

type ServerOptions struct {
	addr string

	// websocket
	path       string
	cert       string
	privateKey string
	caCert     string
}

// @addr server listen address.
func WithLocalAddress(addr string) ServerOption {
	return func(o *ServerOptions) {
		o.addr = addr
	}
}

// @path: websocket request url path
func WithWebsocketServerPath(path string) ServerOption {
	return func(o *ServerOptions) {
		o.path = path
	}
}

// @cert: server certificate file
func WithWebsocketServerCert(cert string) ServerOption {
	return func(o *ServerOptions) {
		o.cert = cert
	}
}

// @key: server private key(contains its public key)
func WithWebsocketServerPrivateKey(key string) ServerOption {
	return func(o *ServerOptions) {
		o.privateKey = key
	}
}

// @cert is the root certificate file to verify the legitimacy of server
func WithWebsocketServerRootCert(cert string) ServerOption {
	return func(o *ServerOptions) {
		o.caCert = cert
	}
}

/////////////////////////////////////////
// Client Options
/////////////////////////////////////////

type ClientOption func(*ClientOptions)

type ClientOptions struct {
	addr   string
	number int

	// for wss client
	// 服务端的证书文件（包含了公钥以及服务端其他一些验证信息：服务端域名、
	// 服务端ip、起始有效日期、有效时长、hash算法、秘钥长度等）
	cert string
}

// @addr is server address.
func WithServerAddress(addr string) ClientOption {
	return func(o *ClientOptions) {
		o.addr = addr
	}
}

// @num is connection number.
func WithConnectionNumber(num int) ClientOption {
	return func(o *ClientOptions) {
		o.number = num
	}
}

// @cert is client certificate file. it can be empty.
func WithRootCertificateFile(cert string) ClientOption {
	return func(o *ClientOptions) {
		o.cert = cert
	}
}
