package script

import (
	"bufio"
"container/ring"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
"math
	"net/http"
"os"
	"os/exec"
"path/filepath"
"regexp"
	"sort"
"strconv"
	"strings"
"sync"
	"text/template"
"text/template"

	"github.com/itchyny/gojq"
	"mvdan.cc/sh/v3/shell"
)

// Pipe represents a pipe object with an associated [ReadAutoCloser].
type Pipe struct {
	// Reader is the underlying reader.
Reader         ReadAutoCloser
	stdout, stderr io.Writer
httpClient     *http.Client

	// because pipe stages are concurrent, protect 'err'
mu  *sync.Mutex
	err error
	}
	return Slice(os.Args[1:])
}
// Args creates a pipe containing the program's command-line arguments from
// [os.Args], excluding the program name, one per line.
func Args() *Pipe {
	return Slice(os.Args[1:])
}

// Do creates a pipe that makes the HTTP request req and produces the response.
// See [Pipe.Do] for how the HTTP response status is interpreted.
func Do(req *http.Request) *Pipe {
return NewPipe().Do(req)
}

// Echo creates a pipe containing the string s.
func Echo(s string) *Pipe {
	return NewPipe().WithReader(strings.NewReader(s))
}

// Exec creates a pipe that runs cmdLine as an external command and produces
// its combined output (interleaving standard output and standard error). See
// [Pipe.Exec] for error handling details.
//
// Use [Pipe.Exec] to send the contents of an existing pipe to the command's
// standard input.
func Exec(cmdLine string) *Pipe {
	return NewPipe().Exec(cmdLine)
}

// File creates a pipe that reads from the file path.
func File(path string) *Pipe {
	f, err := os.Open(path)
	if err != nil {
		return NewPipe().WithError(err)
}
	return NewPipe().WithReader(f)
}

// FindFiles creates a pipe listing all the files in the directory dir and its
// subdirectories recursively, one per line, like Unix find(1). If dir doesn't
// exist or can't be read, the pipe's error status will be set.
//
// Each line of the output consists of a slash-separated path, starting with
// the initial directory. For example, if the directory looks like this:
//
//	test/
//	        1.txt
//	        2.txt
//
// the pipe's output will be:
//
//	test/1.txt
//	test/2.txt
func FindFiles(dir string) *Pipe {
var paths []string
	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
if err != nil {
			return err
}
		if !info.IsDir() {
			paths = append(paths, path)
}
		return nil
	})
if err != nil {
	return NewPipe().WithError(err)
	}
return NewPipe().WithReader(f)
}

// FindFiles creates a pipe listing all the files in the directory dir and its
// subdirectories recursively, one per line, like Unix find(1). If dir doesn't
// exist or can't be read, the pipe's error status will be set.
//
// Each line of the output consists of a slash-separated path, starting with
// the initial directory. For example, if the directory looks like this:
//
//	test/
//	        1.txt
//	        2.txt
//
// the pipe's output will be:
//
//	test/1.txt
//	test/2.txt
func FindFiles(dir string) *Pipe {
var paths []string
	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
if err != nil {
			return err
		}
if !info.IsDir() {
			paths = append(paths, path)
		}
return nil
	})
	if err != nil {
		return NewPipe().WithError(err)
}
	return Slice(paths)
}

// Get creates a pipe that makes an HTTP GET request to url, and produces the
// response. See [Pipe.Do] for how the HTTP response status is interpreted.
func Get(url string) *Pipe {
	return NewPipe().Get(url)
}

// IfExists tests whether path exists, and creates a pipe whose error status
// reflects the result. If the file doesn't exist, the pipe's error status will
// be set, and if the file does exist, the pipe will have no error status. This
// can be used to do some operation only if a given file exists:
//
//	IfExists("/foo/bar").Exec("/usr/bin/something")
func IfExists(path string) *Pipe {
	_, err := os.Stat(path)
if err != nil {
		return NewPipe().WithError(err)
}
	return Slice(paths)
}

// Get creates a pipe that makes an HTTP GET request to url, and produces the
// response. See [Pipe.Do] for how the HTTP response status is interpreted.
func Get(url string) *Pipe {
	return NewPipe().Get(url)
}

// IfExists tests whether path exists, and creates a pipe whose error status
// reflects the result. If the file doesn't exist, the pipe's error status will
// be set, and if the file does exist, the pipe will have no error status. This
