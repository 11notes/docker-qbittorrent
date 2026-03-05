package main

import (
	"os"
	"syscall"
	"crypto/sha512"
  "encoding/base64"

  "github.com/11notes/go-eleven"
	"golang.org/x/crypto/pbkdf2"
)

const APP_BIN = "/usr/local/bin/qbittorrent"
const APP_CONFIG_ENV = "QBITTORRENT_CONFIG"
const APP_CONFIG string = "/qbittorrent/etc/qBittorrent.conf"
const APP_DEFAULT_PASSWORD string = "@ByteArray(188J/h/wfAYQ9H+mTl/7lA==:j/+e2SwJUi9g+IPiEG2+Pix9W0IOv2c20QjrmBUhr4TBUXO3fcMv6leeU6qK8834xiq8fngh8ShwYDfYO0w6lg==)"

func main(){
	// write env to file if set
	eleven.Container.EnvToFile(APP_CONFIG_ENV, APP_CONFIG)

	// check if using default config, if yes, replace default password with random one and print to log
  if ok, _ := eleven.Util.FileContains(APP_CONFIG, APP_DEFAULT_PASSWORD); ok {
		password := eleven.Util.Password()
		salt := eleven.Util.Password()
  	hash := pbkdf2.Key([]byte(password), []byte(salt), 100000, 64, sha512.New)
		hashedPassword := "@ByteArray(" + base64.StdEncoding.EncodeToString([]byte(salt)) + ":" + base64.StdEncoding.EncodeToString([]byte(hash)) + ")"
		replaced, err := eleven.Util.FileReplaceStrings(APP_CONFIG, map[string]any{APP_DEFAULT_PASSWORD:hashedPassword})
		if err != nil {
			eleven.LogFatal("could not set a new default password: %s", err)
		}
		if replaced {
			eleven.Log("INF", "password for account admin: %s", password)
		}else{
			eleven.LogFatal("could not set a new default password, because it could not be found, please check your config %s", APP_CONFIG)
		}
	}

	// start qbittorrent and replace process with it
	if err := syscall.Exec(APP_BIN, []string{"qbittorrent", "--profile=/opt"}, os.Environ()); err != nil {
		os.Exit(1)
	}
}