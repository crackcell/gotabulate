/***************************************************************
 *
 * Copyright (c) 2015, Menglong TAN <tanmenglong@gmail.com>
 *
 * This program is free software; you can redistribute it
 * and/or modify it under the terms of the BSD licence
 *
 **************************************************************/

/**
 *
 *
 * @file table.go
 * @author Menglong TAN <tanmenglong@gmail.com>
 * @date Thu Oct  8 23:58:11 2015
 *
 **/

package gotable

import (
	"fmt"
)

//===================================================================
// Public APIs
//===================================================================

type Table struct {
	header []string
	col    int
}

func New(row, col int) *Table {
	return &Table{}
}

func (this *Table) AddHeader(header []string) {
	this.header = header
	this.col = len(header)
}

//===================================================================
// Private
//===================================================================
