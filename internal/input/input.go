package input

import (
	"bufio"
	"context"
	"errors"
	"fmt"
	"os"
	"strings"
)

var (
	ErrBadHTTPResponse = errors.New("non-200 HTTP response")
)

// Scanner is a facade for a *bufio.Scanner object. It is useful for decoupling a
// solution from file IO.
type Scanner interface {
	// Scan advances the Scanner to the next token. It returns false if an error
	// occurs or if EOF has been reached.
	Scan() bool

	// Err returns the first non-EOF error returned by the Scanner.
	Err() error

	// Text returns the most recent token found by the Scanner as a string.
	Text() string

	// Close all resources allocated by this Scanner. This should be called to clean
	// up once the scanner is no longer needed.
	Close() error
}

type scanner struct {
	*bufio.Scanner

	close func() error
}

func (s *scanner) Close() error { return s.close() }

func NewFileScanner(ctx context.Context, path string) (Scanner, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("failed to open file (%w)", err)
	}

	scanner := &scanner{
		Scanner: bufio.NewScanner(file),
		close:   file.Close,
	}

	return scanner, nil
}

func NewStringScanner(ctx context.Context, input string) Scanner {
	scanner := &scanner{
		Scanner: bufio.NewScanner(strings.NewReader(input)),
		close:   func() error { return nil },
	}

	return scanner
}
