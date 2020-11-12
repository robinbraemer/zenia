package testdata

import (
	"fmt"
	"github.com/robinbraemer/zenia/internal/acl"
	"gopkg.in/yaml.v3"
	"io"
	"os"
	"path/filepath"
)

func LoadNamespaceConfigs(file string) (a []acl.Namespace, err error) {
	var r io.Reader
	return a, filepath.Walk(file, func(path string, info os.FileInfo, err error) error {
		if info == nil || info.IsDir() {
			return nil // skip
		}
		ext := filepath.Ext(info.Name())
		if !(ext == ".yaml" || ext == ".yml") {
			return nil
		}
		r, err = os.Open(path)
		if err != nil {
			return err
		}
		var ns acl.Namespace
		err = yaml.NewDecoder(r).Decode(&ns)
		if err != nil {
			return fmt.Errorf("error decoding %s: %w",
				info.Name(), err)
		}
		a = append(a, ns)
		return nil
	})
}
