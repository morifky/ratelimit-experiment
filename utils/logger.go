package utils

import (
	"log"
	"os"

	"go.uber.org/zap"
)

const (
	logPath    = "./logs"
	outputFile = logPath + "/go.log"
)

func InitLogger() {
	if _, err := os.Stat(logPath); os.IsNotExist(err) {
		_ = os.Mkdir(logPath, os.ModePerm)
	}
	_, err := os.OpenFile(outputFile, os.O_RDONLY|os.O_CREATE, 0666)
	if err != nil {
		log.Fatal(err)
	}
	c := zap.NewProductionConfig()
	c.OutputPaths = []string{"stdout", outputFile}
	l, err := c.Build()
	if err != nil {
		log.Fatal(err)
	}
	zap.ReplaceGlobals(l)
}
