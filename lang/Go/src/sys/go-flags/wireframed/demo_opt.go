////////////////////////////////////////////////////////////////////////////
// Program: redo
// Purpose: global option redo
// Authors: Myself <me@mine.org> (c) 2022, All rights reserved
////////////////////////////////////////////////////////////////////////////

package main

////////////////////////////////////////////////////////////////////////////
// Constant and data type/structure definitions

// The OptsT type defines all the configurable options from cli.
type OptsT struct {
	Host    string `short:"H" long:"host" env:"REDO_HOST" description:"host address" default:"localhost"`
	Port    int    `short:"p" long:"port" env:"REDO_PORT" description:"listening port" default:"80"`
	Force   bool   `short:"f" long:"force" env:"REDO_FORCE" description:"force start"`
	Verbflg func() `short:"v" long:"verbose" description:"Verbose mode (Multiple -v options increase the verbosity)"`
	Verbose int
}
