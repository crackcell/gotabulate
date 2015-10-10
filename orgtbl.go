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
 * @file orgtbl.go
 * @author Menglong TAN <tanmenglong@gmail.com>
 * @date Sat Oct 10 16:35:28 2015
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

type OrgtblFormatter struct{}

func NewOrgtblFormatter() *OrgtblFormatter {
	return &OrgtblFormatter{}
}

func (this *OrgtblFormatter) Format(table *Table) string {
	str := ""
	var header []string
	if table.FirstRowHeader && len(table.Data) > 0 {
		header = table.Data[0]
	} else {
		header = table.Headers
	}
	str += this.formatLine(header, table)
	str += this.formatHeaderSep(table)
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

func (this *OrgtblFormatter) formatLine(row []string, table *Table) string {
	str := ""
	l := len(row)
	str += fmt.Sprintf("| %s%s ", row[0], strings.Repeat(" ", table.CellWidth[0]-len(row[0])))
	for i := 1; i < table.ColumnSize; i++ {
		v := ""
		if i < l {
			v = row[i]
		}
		str += fmt.Sprintf("| %s%s ", strings.Repeat(" ", table.CellWidth[i]-len(v)), v)
	}
	str += "|\n"
	return str
}

func (this *OrgtblFormatter) formatHeaderSep(table *Table) string {
	str := fmt.Sprintf("|%s", strings.Repeat("-", table.CellWidth[0]+2))
	for i := 1; i < table.ColumnSize; i++ {
		str += fmt.Sprintf("+%s", strings.Repeat("-", table.CellWidth[i]+2))
	}
	str += "|\n"
	return str
}
