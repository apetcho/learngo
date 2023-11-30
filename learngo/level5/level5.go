package level5

import (
	"crypto/sha256"
	b64 "encoding/base64"
	"fmt"
	"net"
	"net/url"
	str "strings"
)

// -*-
var println = fmt.Println

// -*- header(text) -*-
func header(text string) {
	num := len(text) + 2
	line := str.Join([]string{"-*--", str.Repeat("-", num), "--*-"}, "")
	fmt.Println()
	fmt.Println(line)
	fmt.Println("-*- ", text, " -*-")
	fmt.Println(line)
}

// -*-------------------*-
// -*- (01) URLParsing -*-
// -*-------------------*-
func URLParsing() {
	header("(01) URL Parsing")

	addr := "postgres://user:pass@host.com:5432/path?k=v#f"

	u, err := url.Parse(addr)
	if err != nil {
		panic(err)
	}
	println(u.Scheme)
	println(u.User)
	println(u.User.Username())
	password, _ := u.User.Password()
	println(password)

	println(u.Host)
	host, port, _ := net.SplitHostPort(u.Host)
	println(host)
	println(port)
	println(u.Path)
	println(u.Fragment)
	println(u.RawQuery)
	m, _ := url.ParseQuery(u.RawQuery)
	println(m)
	println(m["k"][0])
}

// -*---------------------*-
// -*- (02) SHA256Hashes -*-
// -*---------------------*-
func SHA256Hashes() {
	header("(02) SHA256 Hashes")
	text := "sha256 this string"
	h := sha256.New()
	h.Write([]byte(text))
	bs := h.Sum(nil)
	println(text)
	fmt.Printf("%x\n", bs)
}

// -*----------------------*-
// -*- (03) Bas64Encoding -*-
// -*----------------------*-
func Bas64Encoding() {
	header("(03) Base64 Encoding")
	data := "abc123!?$*()'-=@~"

	sEnc := b64.StdEncoding.EncodeToString([]byte(data))
	println(sEnc)

	sDec, _ := b64.StdEncoding.DecodeString(sEnc)
	println(string(sDec))
	println()

	uEnc := b64.URLEncoding.EncodeToString([]byte(data))
	println(uEnc)
	uDec, _ := b64.URLEncoding.DecodeString(uEnc)
	println(string(uDec))
}
