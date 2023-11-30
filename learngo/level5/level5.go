package level5

import (
	"bufio"
	"crypto/sha256"
	"embed"
	b64 "encoding/base64"
	"flag"
	"fmt"
	"io"
	"io/fs"
	"net"
	"net/url"
	"os"
	"path/filepath"
	str "strings"
)

// -*-
var println = fmt.Println

// var print = fmt.Print
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

// -*----------------------*-
// -*- (08) GoDirectories -*-
// -*----------------------*-
func GoDirectories() {
	header("(08) Directories")
	check := func(err error) {
		if err != nil {
			panic(err)
		}
	}

	// -
	err := os.Mkdir("subdir", 0755)
	check(err)

	defer os.RemoveAll("subdir")

	createEmptyfile := func(name string) {
		data := []byte("")
		check(os.WriteFile(name, data, 0644))
	}
	createEmptyfile("subdir/file1")

	err = os.MkdirAll("subdir/parent/child", 0755)
	check(err)

	createEmptyfile("subdir/parent/file2")
	createEmptyfile("subdir/parent/file3")
	createEmptyfile("subdir/parent/child/file4")

	c, err := os.ReadDir("subdir/parent")
	check(err)
	println("Listing subdir/parent")
	for _, entry := range c {
		println(" ", entry.Name(), entry.IsDir())
	}

	err = os.Chdir("subdir/parent/child")
	check(err)

	println("Listing subdir/parent/child")
	for _, entry := range c {
		println(" ", entry.Name(), entry.IsDir())
	}

	err = os.Chdir("../../..")
	check(err)

	visit := func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		println(" ", path, d.IsDir())
		return nil
	}

	println("Visiting subdir")
	_ = filepath.WalkDir("subdir", visit)
}

// -*-------------------------------------*-
// -*- (09) TemporaryFilesAndDirectories -*-
// -*-------------------------------------*-
func TemporaryFilesAndDirectories() {
	header("(09) Temporary Files and Directories")

	check := func(err error) {
		if err != nil {
			panic(err)
		}
	}

	// -
	fp, err := os.CreateTemp("", "sample")
	check(err)
	println("Temp file name:", fp.Name())

	defer os.Remove(fp.Name())

	_, err = fp.Write([]byte{1, 2, 3, 4})
	check(err)

	dname, err := os.MkdirTemp("", "sampledir")
	check(err)
	println("Temp dir name:", dname)

	defer os.RemoveAll(dname)

	fname := filepath.Join(dname, "file1")
	err = os.WriteFile(fname, []byte{1, 2}, 0666)
	check(err)
}

// -*-----------------------*-
// -*- (10) EmbedDirective -*-
// -*-----------------------*-

//go:embed folder/single_file.txt
var fileString string

//go:embed folder/single_file.txt
var fileByte []byte

//go:embed folder/single_file.txt
//go:ember folder/*.hash
var folder embed.FS

func EmbedDirective() {
	header("(10) Embed Directive")
	print(fileString)
	print(string(fileByte))

	content1, _ := folder.ReadFile("folder/file1.hash")
	print(string(content1))

	content2, _ := folder.ReadFile("folder/file2.hash")
	print(string(content2))
}

// -*------------------------*-
// -*- (12) CommandLineArgs -*-
// -*------------------------*-
func CommandLineArgs() {
	header("(12) Command-Line Args")

	argsWithProg := os.Args
	argsWithoutProg := os.Args[1:]

	arg := os.Args[3]
	println(argsWithProg)
	println(argsWithoutProg)
	println(arg)
}

// -*---------------------------*-
// -*- (13) CommandLineFlags() -*-
// -*---------------------------*-
func CommandLineFlags() {
	header("(13) Command-Line Flags")
	wordPtr := flag.String("word", "foo", "a string")
	numPtr := flag.Int("numb", 42, "an int")
	forkPtr := flag.Bool("fork", false, "a bool")

	var svar string
	flag.StringVar(&svar, "svar", "bar", "a string var")
	flag.Parse()

	println("word:", *wordPtr)
	println("numb:", *numPtr)
	println("fork:", *forkPtr)
	println("svar:", svar)
	println("tail:", flag.Args())
}

// -*-------------------------------*-
// -*- (14) CommandLineSubcommands -*-
// -*-------------------------------*-
func CommandLineSubcommands() {
	header("(14) Command-Line Subcommands")
	fooCmd := flag.NewFlagSet("foo", flag.ExitOnError)
	fooEnable := fooCmd.Bool("enable", false, "enable")
	fooName := fooCmd.String("name", "", "name")

	barCmd := flag.NewFlagSet("bar", flag.ExitOnError)
	barLevel := barCmd.Int("level", 0, "level")

	if len(os.Args) < 2 {
		println("expected 'foo' or 'bar' subcommands")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "foo":
		fooCmd.Parse(os.Args[2:])
		println("subcommand 'foo'")
		println("  enable:", *fooEnable)
		println("  name:", *fooName)
		println("  tail:", fooCmd.Args())
	case "bar":
		barCmd.Parse(os.Args[2:])
		println("subcommand 'bar'")
		println("  level:", *barLevel)
		println("  tail:", barCmd.Args())
	default:
		println("expected 'foo' or 'bar' subcommands")
		os.Exit(1)
	}
}

// -*-----------------------------*-
// -*- (15) EnvironmentVariables -*-
// -*-----------------------------*-
func EnvironmentVariables() {
	header("(15) Environment Variables")
	os.Setenv("FOO", "1")
	println("FOO:", os.Getenv("FOO"))
	println("BAR:", os.Getenv("BAR"))
	println()
	for _, err := range os.Environ() {
		pair := str.SplitN(err, "=", 2)
		println(pair[0])
	}
}
