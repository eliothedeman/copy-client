package main

import (
	"github.com/eliothedeman/copy-client/cc"
)

func main() {
	c := &cc.AuthConfig{}
	c.ConsumerKey = "K2fw6eqRv0x8RPIlysJWZBgW1Hqc1bzT"
	c.ConsumerSecret = "3TfSVXEFNeYrmoc3nTrH4Ea39Q0QgLULatwveuhFPvHjeQ2b"
	a := &cc.Auth{Config: c}
	a.Do()
}
