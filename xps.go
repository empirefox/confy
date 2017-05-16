package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/empirefox/confy/xps"
)

var (
	decryptdir = flag.String("d", "", "Decrypt dir to place files, switch to encrypt mode if not set")
	password   = flag.String("k", "", "Password")
	configfile = flag.String("x", "xps-config.json", "Config file for xps")
)

func main() {
	flag.Parse()
	instance, err := xps.NewXps(*configfile, "json")
	if err != nil {
		panic(err)
	}

	if *decryptdir != "" {
		err = os.RemoveAll(*decryptdir)
		if err != nil {
			panic(err)
		}
		err = os.MkdirAll(*decryptdir, os.ModePerm)
		if err != nil {
			panic(err)
		}

		files, err := instance.DecryptXhexFile(*password)
		if err != nil {
			panic(err)
		}

		for name, content := range files {
			err = ioutil.WriteFile(filepath.Join(*decryptdir, name), content, os.ModePerm)
			if err != nil {
				panic(err)
			}
		}
		fmt.Printf("Decrypted ok to dir: %s\n", *decryptdir)
	} else {
		err = instance.EncryptXhexFile(*password)
		if err != nil {
			panic(err)
		}
		fmt.Printf("Encrypted ok to file: %s\n", instance.XpsFile)
	}
}
