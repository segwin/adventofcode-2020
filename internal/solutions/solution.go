package solutions

import "context"

// Solution is the interface implemented by all solutions for any given day.
type Solution interface {
	Run(ctx context.Context, inputFile string)
}
