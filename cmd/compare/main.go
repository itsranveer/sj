package main

import (
	"bufio"
	"context"
	"os"
	"strings"

	"github.com/itsranveer/sj/worker"
	"github.com/sirupsen/logrus"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	logrus.Info("Enter Version 1: ")
	version1, _ := reader.ReadString('\n')
	version1 = strings.Replace(version1, "\n", "", -1)

	logrus.Info("Enter Version 2: ")
	version2, _ := reader.ReadString('\n')
	version2 = strings.Replace(version2, "\n", "", -1)

	versionApp := worker.NewVersion()
	result := versionApp.CompareVersion(context.Background(), version1, version2)

	switch result {
	case 1:
		logrus.Infof("Version %s is bigger than Version %s", version1, version2)
	case 0:
		logrus.Infof("Version %s is same as Version %s", version1, version2)
	case -1:
		logrus.Infof("Version %s is smaller than Version %s", version1, version2)
	}

}
