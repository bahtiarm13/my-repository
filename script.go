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
