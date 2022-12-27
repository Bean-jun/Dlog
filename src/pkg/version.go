package pkg

import (
	"errors"
	"strconv"
	"strings"
)

const (
	Version = "2022.11.04.6417e50"
)

func ParseVersion() (int, error) {
	versionSlice := strings.Split(Version, ".")
	if len(versionSlice) < 1 {
		return -1, errors.New("parse version failed")
	}
	versionStr := strings.Join(versionSlice[:len(versionSlice)-1], "")
	version, err := strconv.Atoi(versionStr)
	if err != nil {
		return -1, err
	}
	return version, nil
}
