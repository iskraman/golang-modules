package fslib

import (
	"io/fs"
	"io/ioutil"

	"github.com/iskraman/golang-modules/syslog"
)

func FileReader(filename string) ([]byte, error) {
	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		syslog.WAR("%s", err)
		return nil, err
	}

	return bytes, err
}

func FileWriter(filename string, w []byte, perm fs.FileMode) error {
	if perm == 0 {
		perm = 0644
	}

	err := ioutil.WriteFile(filename, w, perm)
	if err != nil {
		syslog.WAR("%s", err)
		return err
	}
	return err
}
