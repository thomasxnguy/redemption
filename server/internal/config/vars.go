package config

import (
	"github.com/trustwallet/blockatlas/pkg/errors"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

const (
	rootPath = "./../client/build/static/js/"
	varName  = "{{.API_AUTH_TOKEN}}"
)

func ReplaceAuthToken() error {
	files, err := filesFromFolder(rootPath)
	if err != nil {
		return errors.E(err, "Failed to get files from folder", errors.Params{"path": rootPath})
	}
	for _, f := range files {
		if !strings.HasSuffix(f, ".js") {
			continue
		}
		err := replaceVars(f, varName, Configuration.Api.Auth_Token)
		if err != nil {
			return errors.E(err, "Failed to replace token var from file", errors.Params{"file": f, "path": rootPath})
		}
	}
	return nil
}

func filesFromFolder(root string) (files []string, err error) {
	err = filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if path != root {
			files = append(files, path)
		}
		return err
	})
	return
}

func replaceVars(path, old, new string) error {
	read, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}
	newContents := strings.Replace(string(read), old, new, -1)
	return ioutil.WriteFile(path, []byte(newContents), 0)
}
