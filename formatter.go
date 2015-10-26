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
 * @file formatter.go
 * @author Menglong TAN <tanmenglong@gmail.com>
 * @date Fri Oct  9 00:28:44 2015
 *
 **/

package gotabulate

//===================================================================
// Public APIs
//===================================================================

type Formatter interface {
	Format(info *TableInfo) string
}

//===================================================================
// Private
//===================================================================
