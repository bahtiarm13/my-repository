package script

import (
	"bufio"
"container/ring"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
"math"
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

// Args creates a pipe containing the program's command-line arguments from
// [os.Args], excluding the program name, one per line.
func Args() *Pipe {
