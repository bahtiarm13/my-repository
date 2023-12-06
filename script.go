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
// can be used to do some operation only if a given file exists:
//
//	IfExists("/foo/bar").Exec("/usr/bin/something")
func IfExists(path string) *Pipe {
	_, err := os.Stat(path)
if err != nil {
		return NewPipe().WithError(err)
	}
	return NewPipe()
}

// ListFiles creates a pipe containing the files or directories specified by
// path, one per line. path can be a glob expression, as for [filepath.Match].
// For example:
//
//	ListFiles("/data/*").Stdout()
//
// ListFiles does not recurse into subdirectories; use [FindFiles] instead.
func ListFiles(path string) *Pipe {
	if strings.ContainsAny(path, "[]^*?\\{}!") {
	fileNames, err := filepath.Glob(path)
		if err != nil {
			return NewPipe().WithError(err)
		}
return Slice(fileNames)
	}
	entries, err := os.ReadDir(path)
if err != nil {
		// Check for the case where the path matches exactly one file
		s, err := os.Stat(path)
if err != nil {
			return NewPipe().WithError(err)
		}
		if !s.IsDir() {
			return Echo(path)
}
		return NewPipe().WithError(err)
	}
	matches := make([]string, len(entries))
for i, e := range entries {
		matches[i] = filepath.Join(path, e.Name())
}
	return Slice(matches)
}

// NewPipe creates a new pipe with an empty reader (use [Pipe.WithReader] to
// attach another reader to it).
func NewPipe() *Pipe {
return &Pipe{
		Reader:     ReadAutoCloser{},
	mu:         new(sync.Mutex),
	stdout:     os.Stdout,
		httpClient: http.DefaultClient,
	}
}

// Post creates a pipe that makes an HTTP POST request to url, with an empty
// body, and produces the response. See [Pipe.Do] for how the HTTP response
// status is interpreted.
func Post(url string) *Pipe {
return NewPipe().Post(url)
}

// Slice creates a pipe containing each element of s, one per line.
func Slice(s []string) *Pipe {
return Echo(strings.Join(s, "\n") + "\n")
}

// Stdin creates a pipe that reads from [os.Stdin].
func Stdin() *Pipe {
	return NewPipe().WithReader(os.Stdin)
}

// AppendFile appends the contents of the pipe to the file path, creating it if
// necessary, and returns the number of bytes successfully written, or an
// error.
func (p *Pipe) AppendFile(path string) (int64, error) {
	return p.writeOrAppendFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY)
}

// Basename reads paths from the pipe, one per line, and removes any leading
// directory components from each. So, for example, /usr/local/bin/foo would
// become just foo. This is the complementary operation to [Pipe.Dirname].
//
// If any line is empty, Basename will transform it to a single dot. Trailing
// slashes are removed. The behaviour of Basename is the same as
// [filepath.Base] (not by coincidence).
func (p *Pipe) Basename() *Pipe {
	return p.FilterLine(filepath.Base)
}

// Bytes returns the contents of the pipe as a []byte, or an error.
func (p *Pipe) Bytes() ([]byte, error) {
	if p.Error() != nil {
	return nil, p.Error()
	}
	data, err := io.ReadAll(p)
if err != nil {
		p.SetError(err)
}
	return data, p.Error()
}

// Close closes the pipe's associated reader. This is a no-op if the reader is
// not an [io.Closer].
func (p *Pipe) Close() error {
	return p.Reader.Close()
}

// Column produces column col of each line of input, where the first column is
// column 1, and columns are delimited by Unicode whitespace. Lines containing
// fewer than col columns will be skipped.
func (p *Pipe) Column(col int) *Pipe {
return p.FilterScan(func(line string, w io.Writer) {
		columns := strings.Fields(line)
	if col > 0 && col <= len(columns) {
	fmt.Fprintln(w, columns[col-1])
		}
	})
}

// Concat reads paths from the pipe, one per line, and produces the contents of
// all the corresponding files in sequence. If there are any errors (for
// example, non-existent files), these will be ignored, execution will
// continue, and the pipe's error status will not be set.
//
// This makes it convenient to write programs that take a list of paths on the
// command line. For example:
//
//	script.Args().Concat().Stdout()
//
// The list of paths could also come from a file:
//
//	script.File("filelist.txt").Concat()
//
// Or from the output of a command:
//
//	script.Exec("ls /var/app/config/").Concat().Stdout()
//
// Each input file will be closed once it has been fully read. If any of the
// files can't be opened or read, Concat will simply skip these and carry on,
// without setting the pipe's error status. This mimics the behaviour of Unix
// cat(1).
func (p *Pipe) Concat() *Pipe {
var readers []io.Reader
	p.FilterScan(func(line string, w io.Writer) {
columns := strings.Fields(line)
if col > 0 && col <= len(columns) {
			fmt.Fprintln(w, columns[col-1])
		}
	})
}

// Concat reads paths from the pipe, one per line, and produces the contents of
// all the corresponding files in sequence. If there are any errors (for
// example, non-existent files), these will be ignored, execution will
// continue, and the pipe's error status will not be set.
//
// This makes it convenient to write programs that take a list of paths on the
// command line. For example:
//
//	script.Args().Concat().Stdout()
