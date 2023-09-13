package error_pkg

import "github.com/sirupsen/logrus"

func CheckErr(logger *logrus.Logger, err error) {
	if err != nil {
		logger.Error(err)
	}
}
