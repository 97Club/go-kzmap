package kreedz

import (
	"net"
)

func handleHttpError(err error) {
	switch err := err.(type) {
	case net.Error:
		if err.Timeout() {
			logger.Println(err.Error())
			return
		}
	default:
		panic(err)
	}
}