package asciiart

import (
	"crypto/md5"
	"encoding/hex"
	"os"
)

const (
	hashST string = "b7e06e7f6a2d24d8da5d57d3cba6a2c7"
	hashSH string = "0ca33a970e2a1c5b53ecbcad43d60b40"
	hashTH string = "f7d527c38c0b2ea6df5c12dafb285fd1"
)

func GetMD5Hash(text string) string {
	hasher := md5.New()
	hasher.Write([]byte(text))
	return hex.EncodeToString(hasher.Sum(nil))
}

func CheckHash(s string) (string, bool) {
	if s == "st" {
		st, err := os.ReadFile("asciiart/standard.txt")
		if err != nil {
			return "", false
		}
		if GetMD5Hash(string(st)) != hashST {
			return "", false
		}
		return string(st), true
	} else if s == "sh" {
		sh, err := os.ReadFile("asciiart/shadow.txt")
		if err != nil {
			return "", false
		}
		if GetMD5Hash(string(sh)) != hashSH {
			return "", false
		}
		return string(sh), true
	} else if s == "th" {
		th, err := os.ReadFile("asciiart/thinkertoy.txt")
		if err != nil {
			return "", false
		}
		if GetMD5Hash(string(th)) != hashTH {
			return "", false
		}
		return string(th), true
	}

	return "", true
}
