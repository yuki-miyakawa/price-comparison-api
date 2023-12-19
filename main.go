package main

import (
	"github.com/y-miyakaw/price-comparison-api/cmd/login"
)

func main() {
	user := login.CreateJWT()
	login.ParseJWT(user)

}
