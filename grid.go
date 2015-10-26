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
 * @file psql.go
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
type GridFormatter struct{}

func NewGridFormatter() *GridFormatter {
	return &GridFormatter{}
}

func (this *GridFormatter) Format(info *TableInfo) string {
	str := ""
	var header []string
	if info.FirstRowHeader && len(info.Data) > 0 {
		header = info.Data[0]
	} else {
		header = info.Headers
	}
	str += this.formatLineSep(info, "-")
	str += this.formatLine(header, info)
	str += this.formatLineSep(info, "=")
	rowStart := 0
	if info.FirstRowHeader {
		rowStart = 1
	}
	for i := rowStart; i < len(info.Data); i++ {
		str += this.formatLine(info.Data[i], info)
		str += this.formatLineSep(info, "-")
	}
	return str
}

//===================================================================
// Private
//===================================================================

func (this *GridFormatter) formatLine(row []string, info *TableInfo) string {
	str := ""
	l := len(row)
	str += fmt.Sprintf("| %s%s ", row[0], strings.Repeat(" ", info.CellWidth[0]-len(row[0])))
	for i := 1; i < info.ColumnSize; i++ {
		v := ""
		if i < l {
			v = row[i]
		}
		str += fmt.Sprintf("| %s%s ", strings.Repeat(" ", info.CellWidth[i]-len(v)), v)
	}
	str += "|\n"
	return str
}

func (this *GridFormatter) formatLineSep(info *TableInfo, sym string) string {
	str := fmt.Sprintf("+%s", strings.Repeat(sym, info.CellWidth[0]+2))
	for i := 1; i < info.ColumnSize; i++ {
		str += fmt.Sprintf("+%s", strings.Repeat(sym, info.CellWidth[i]+2))
	}
	str += "+\n"
	return str
}
