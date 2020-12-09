package day8

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

var (
	ErrInvalidInstructionMsg = errors.New("invalid instruction message")
	ErrUnknownOp             = errors.New("unknown operation in instruction message")
)

type operation int

const (
	acc operation = iota
	jmp
	nop
)

type instruction struct {
	operation operation
	argument  int

	hit bool // if true, indicates this instruction has already been executed
}

func (i *instruction) Unmarshal(msg string) error {
	pair := strings.SplitN(msg, " ", 2)
	if len(pair) != 2 {
		return fmt.Errorf("%w (got %d components)", ErrInvalidInstructionMsg, len(pair))
	}

	var op operation
	switch pair[0] {
	case "acc":
		op = acc
	case "jmp":
		op = jmp
	case "nop":
		op = nop
	default:
		return fmt.Errorf("%w (%s)", ErrUnknownOp, pair[0])
	}

	arg, err := strconv.Atoi(strings.TrimPrefix(pair[1], "+"))
	if err != nil {
		return fmt.Errorf("invalid argument: %w", err)
	}

	i.operation = op
	i.argument = arg

	return nil
}
