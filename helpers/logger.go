package helpers

import (
	"log"
	"os"
)

var (
	outfile, _ = os.OpenFile("./gin.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666) // update path for your needs
	logger     = log.New(outfile, "", 0)
)

func LogError(message string) {
	logger.Printf("Error: %s", message)
}

func LogDanger(message string) {
	logger.Printf("Danger: %s", message)
}

func LogInfo(message string) {
	logger.Printf("Info: %s", message)
}
