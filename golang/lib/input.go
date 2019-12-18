package lib

import (
	"io/ioutil"
	"strings"

	"github.com/sirupsen/logrus"
)

// ReadFile reads a file
func ReadFile(filename string) string {
	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		logrus.WithError(err).WithField("path", filename).Fatal("unable to read file")
	}
	return strings.TrimSpace(string(bytes))
}
