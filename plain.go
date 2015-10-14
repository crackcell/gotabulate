/***************************************************************
 *
 * Copyright (c) 2015, Wenbin Xiao <xwb1989@gmail.com>
 *
 * This program is free software; you can redistribute it
 * and/or modify it under the terms of the BSD licence
 *
 **************************************************************/

/**
 *
 *
 * @file plain.go
 * @author Wenbin Xiao <xwb1989gmail.com>
 * @date Wed Oct 14 2015
 *
 **/

package gotabulate

import (
	"fmt"
	"strings"
)

//===================================================================
// Public APIs
//===================================================================
type PlainFormatter struct{}

func NewPlainFormatter() *PlainFormatter {
	return &PlainFormatter{}
}

func (this *PlainFormatter) Format(table *Table) string {
	str := ""
	var header []string
	if table.FirstRowHeader && len(table.Data) > 0 {
		header = table.Data[0]
	} else {
		header = table.Headers
	}
	str += this.formatLine(header, table)
	rowStart := 0
	if table.FirstRowHeader {
		rowStart = 1
	}
	for i := rowStart; i < len(table.Data); i++ {
		str += this.formatLine(table.Data[i], table)
	}
	return str
}

//===================================================================
// Private
//===================================================================

func (this *PlainFormatter) formatLine(row []string, table *Table) string {
	str := ""
	l := len(row)
	str += fmt.Sprintf("%s%s   ", row[0], strings.Repeat(" ", table.CellWidth[0]-len(row[0])))
	for i := 1; i < table.ColumnSize; i++ {
		v := ""
		if i < l {
			v = row[i]
		}
		str += fmt.Sprintf("%s%s   ", strings.Repeat(" ", table.CellWidth[i]-len(v)), v)
	}
	str += "\n"
	return str
}
