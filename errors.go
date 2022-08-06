package main

import "errors"

var (
	ErrPendingTransaction = errors.New("transaction is pending")
	ErrOutOfGas           = errors.New("Out of gas")
	ErrFailedTransaction  = errors.New("failed transaction")
)
