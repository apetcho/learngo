package level5

import (
	"bufio"
	"crypto/sha256"
	b64 "encoding/base64"
	"fmt"
	"io"
	"net"
	"net/url"
	"os"
	"path/filepath"
	str "strings"
)

// -*-
var println = fmt.Println
var print = fmt.Print
var printf = fmt.Printf

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

// -*---------------------*-
// -*- (04) ReadingFiles -*-
// -*---------------------*-
func ReadingFiles() {
	header("(04) Reading Files")
	check := func(e error) {
		if e != nil {
			panic(e)
		}
	}

	// -
	data, err := os.ReadFile("/tmp/data")
	check(err)
	print(string(data))

	fp, err := os.Open("/tmp/data")
	check(err)
	b1 := make([]byte, 5)
	n1, err := fp.Read(b1)
	check(err)
	printf("%d bytes: %s\n", n1, string(b1[:n1]))

	o2, err := fp.Seek(6, 0)
	check(err)
	b2 := make([]byte, 2)
	n2, err := fp.Read(b2)
	check(err)
	printf("%d bytes @ %d: ", n2, o2)
	printf("%v\n", string(b2[:n2]))

	o3, err := fp.Seek(6, 0)
	check(err)
	b3 := make([]byte, 2)
	n3, err := io.ReadAtLeast(fp, b3, 2)
	check(err)
	printf("%d bytes @ %d: %s\n", n3, o3, string(b3))
	_, err = fp.Seek(0, 0)
	check(err)

	r4 := bufio.NewReader(fp)
	b4, err := r4.Peek(5)
	check(err)
	printf("5 bytes: %s\n", string(b4))
	fp.Close()
}

// -*---------------------*-
// -*- (05) WritingFiles -*-
// -*---------------------*-
func WritingFiles() {
	header("(05) Writing Files")

	check := func(err error) {
		if err != nil {
			panic(err)
		}
	}

	// -
	d1 := []byte("hello\ngo\n")
	err := os.WriteFile("/tmp/data1", d1, 0644)
	check(err)

	fp, err := os.Create("/tmp/data2")
	check(err)
	defer fp.Close()

	d2 := []byte{115, 111, 109, 101, 10}
	n2, err := fp.Write(d2)
	check(err)
	printf("wrote %d bytes\n", n2)

	n3, err := fp.WriteString("writes\n")
	check(err)
	printf("wrote %d bytes\n", n3)
	fp.Sync()

	w := bufio.NewWriter(fp)
	n4, err := w.WriteString("buffered\n")
	check(err)
	printf("wrote %d bytes\n", n4)
	w.Flush()
}

// -*---------------------*-
// -*- (06) Line Filters -*-
// -*---------------------*-
func LineFilters() {
	header("(06) Line Filters")
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		ucl := str.ToUpper(scanner.Text())
		println(ucl)
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "error:", err)
		os.Exit(1)
	}
}

// -*------------------*-
// -*- (07) FilePaths -*-
// -*------------------*-
func FilePaths() {
	header("(07) File Paths")
	p := filepath.Join("dir1", "dir2", "filename")
	println("p:", p)

	println(filepath.Join("dir1//", "filename"))
	println(filepath.Join("dir1/../dir1", "filename"))

	println("Dir(p):", filepath.Dir(p))
	println("Base(p):", filepath.Base(p))

	println(filepath.IsAbs("dir/file"))
	println(filepath.IsAbs("/dir/file"))

	filename := "config.json"

	ext := filepath.Ext(filename)
	println(ext)
	println(str.TrimSuffix(filename, ext))

	rel, err := filepath.Rel("a/b", "a/b/t/file")
	if err != nil {
		panic(err)
	}
	println(rel)

	rel, err = filepath.Rel("a/b", "a/c/t/file")
	if err != nil {
		panic(err)
	}
	println(rel)
}
