package worker

import (
	"context"
	"strconv"
	"strings"

	"github.com/sirupsen/logrus"
)

type Version struct{}

func NewVersion() *Version {
	return &Version{}
}

func (v *Version) CompareVersion(ctx context.Context, version1 string, version2 string) int {
	v1Array := strings.Split(version1, ".")
	v2Array := strings.Split(version2, ".")

	for i := 0; i < max(len(v1Array), len(v2Array)); i++ {
		v1 := 0
		v2 := 0
		var err error

		if i < len(v1Array) {
			v1, err = strconv.Atoi(v1Array[i])
			if err != nil {
				logrus.WithError(err).Errorf("Error during v1 conversion, %v", err)

				return 0
			}
		}

		if i < len(v2Array) {
			v2, err = strconv.Atoi(v2Array[i])
			if err != nil {
				logrus.WithError(err).Errorf("Error during v2 conversion, %v", err)

				return 0
			}
		}

		if v1 > v2 {
			return 1
		} else if v1 < v2 {
			return -1
		}
	}

	return 0
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
