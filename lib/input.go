package lib

import (
	"io/ioutil"
	"strconv"
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

// ReadFileAsIntSlice does what you'd think
func ReadFileAsIntSlice(filename string) []int {
	s := ReadFile(filename)
	memo := []int{}
	for _, c := range strings.Split(s, ",") {
		i, _ := strconv.Atoi(c)
		memo = append(memo, i)
	}
	return memo
}
