package config

import (
	"github.com/trustwallet/blockatlas/pkg/errors"
	"github.com/trustwallet/blockatlas/pkg/logger"
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
	logger.Info("try to replace vars inside js files", logger.Params{"token": Configuration.Dashboard.Token})
	for _, f := range files {
		if !strings.HasSuffix(f, ".js") {
			continue
		}
		err := replaceVars(f, varName, Configuration.Dashboard.Token)
		if err != nil {
			return errors.E(err, "Failed to replace token var from file", errors.Params{"file": f, "path": rootPath})
		}
	}
	logger.Info("files replaces", logger.Params{"files": len(files)})
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
	logger.Info("replacing vars inside js file", logger.Params{"path": path})
	read, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}
	newContents := strings.Replace(string(read), old, new, -1)
	return ioutil.WriteFile(path, []byte(newContents), 0)
}
