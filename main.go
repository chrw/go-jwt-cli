package main;

import (
	"os"
	"github.com/k0kubun/pp"
	"code.google.com/p/getopt"
	"github.com/dgrijalva/jwt-go"
)

func main () {
	key := getopt.StringLong("secret", 's', "", "JWT validation secret to validate the signature")
	help:= getopt.BoolLong("help",  'h', "", "show this help")
	getopt.SetParameters("[Encoded JWT String]")
	getopt.Parse()
	
	if *help || getopt.NArgs() < 1 {
		getopt.Usage()
		os.Exit(0)
	}

	keyfunc := func (id *jwt.Token) (interface{}, error) {
		k := []byte(*key)
		return k, nil;
	}

	token, _ := jwt.Parse(getopt.Args()[0],
		keyfunc);
	pp.Print(token);
}

