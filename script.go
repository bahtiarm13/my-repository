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
