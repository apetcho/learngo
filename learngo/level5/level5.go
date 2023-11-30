package level5

import (
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
