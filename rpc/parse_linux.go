//go:build !windows
// +build !windows

package rpc

func RpcParse(resp []byte) error {
	return nil
}
