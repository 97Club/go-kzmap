package kreedz

import (
	"log"
	"os"
)

var logger *log.Logger
func init() {
	fl, err := os.OpenFile("./data/error", os.O_APPEND|os.O_CREATE, 0755)
	if err != nil {
		panic(err)
	}
	logger = log.New(fl, "kz", log.Llongfile)
}
