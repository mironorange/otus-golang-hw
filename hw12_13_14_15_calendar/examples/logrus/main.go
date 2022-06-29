package main

import "github.com/sirupsen/logrus"

func main() {
	logrus.SetLevel(logrus.InfoLevel)
	logrus.Info("Hello!")
}
