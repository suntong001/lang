////////////////////////////////////////////////////////////////////////////
// Program: redo
// Purpose: global option redo
// Authors: Myself <me@mine.org> (c) 2022, All rights reserved
////////////////////////////////////////////////////////////////////////////

package main

import (
//  	"fmt"
//  	"os"

//	"github.com/go-easygen/go-flags"
)

////////////////////////////////////////////////////////////////////////////
// Constant and data type/structure definitions

// The OptsT type defines all the configurable options from cli.
//  type OptsT struct {
//  	Host	string	`short:"H" long:"host" env:"REDO_HOST" description:"host address" default:"localhost"`
//  	Port	int	`short:"p" long:"port" env:"REDO_PORT" description:"listening port" default:"80"`
//  	Force	bool	`short:"f" long:"force" env:"REDO_FORCE" description:"force start"`
//  	Verbflg func()  `short:"v" long:"verbose" description:"Verbose mode (Multiple -v options increase the verbosity)"`
//  	Verbose uint

//  }

// Template for main starts here

////////////////////////////////////////////////////////////////////////////
// Global variables definitions

//  var (
//          progname  = "redo"
//          version   = "0.1.0"
//          date = "2022-01-17"

//  	// Opts store all the configurable options
//  	Opts OptsT
//  )
//
//  var parser = flags.NewParser(&Opts, flags.Default)

////////////////////////////////////////////////////////////////////////////
// Function definitions

// Function main
//  func main() {
//  	Opts.Verbflg = func() {
//  		Opts.Verbose++
//  	}
//
//  	if _, err := parser.Parse(); err != nil {
//  		switch flagsErr := err.(type) {
//  		case flags.ErrorType:
//  			if flagsErr == flags.ErrHelp {
//  				os.Exit(0)
//  			}
//  			os.Exit(1)
//  		default:
//  			fmt.Println()
//  			parser.WriteHelp(os.Stdout)
//  			os.Exit(1)
//  		}
//  	}
//  	fmt.Println("")
//  }
// Template for main ends here

// Template for "build" CLI handling starts here
////////////////////////////////////////////////////////////////////////////
// Program: redo
// Purpose: global option redo
// Authors: Myself <me@mine.org> (c) 2022, All rights reserved
////////////////////////////////////////////////////////////////////////////

//  package main

//  import (
//	"fmt"
//	"os"
//
//	"github.com/go-easygen/go-flags/clis"
//  )

// *** Sub-command: build ***

////////////////////////////////////////////////////////////////////////////
// Constant and data type/structure definitions

// The BuildCommand type defines all the configurable options from cli.
//  type BuildCommand struct {
//  	Dir	string	`long:"dir" description:"source code root dir" default:"./"`
//  }

//
//  var buildCommand BuildCommand
//
//  func init() {
//  	parser.AddCommand("build",
//  		"Build the network application",
//  		"Usage:\n  redo build [Options] Arch(i386|amd64)",
//  		&buildCommand)
//  }
//
//  func (x *BuildCommand) Execute(args []string) error {
//  	return x.Exec(args)
//  }
//
// Exec implements the business logic of command `build`
// func (x *BuildCommand) Exec(args []string) error {
// 	clis.Setup(fmt.Sprintf("%s::%s", progname, "build"), Opts.Verbose)
// 	fmt.Fprintf(os.Stderr, "Build the network application\n")
// 	// fmt.Fprintf(os.Stderr, "Copyright (C) 2022, Myself <me@mine.org>\n\n")
// 	// clis.Verbose(1, "Doing Build, with %+v, %+v\n", Opts, args)
// 	// fmt.Println(x.Dir)
// 	// // err := ...
// 	// // clis.WarnOn("Doing Build", err)
// 	// // or,
// 	// // clis.AbortOn("Doing Build", err)
// 	return nil
// }
// Template for "build" CLI handling ends here

// Template for "install" CLI handling starts here
////////////////////////////////////////////////////////////////////////////
// Program: redo
// Purpose: global option redo
// Authors: Myself <me@mine.org> (c) 2022, All rights reserved
////////////////////////////////////////////////////////////////////////////

//  package main

//  import (
//	"fmt"
//	"os"
//
//	"github.com/go-easygen/go-flags/clis"
//  )

// *** Sub-command: install ***

////////////////////////////////////////////////////////////////////////////
// Constant and data type/structure definitions

// The InstallCommand type defines all the configurable options from cli.
//  type InstallCommand struct {
//  	Dir	string	`short:"d" description:"source code root dir" default:"./"`
//  	Suffix	string	`long:"suffix" description:"source file suffix" default:".go,.c,.s"`
//  }

//
//  var installCommand InstallCommand
//
//  func init() {
//  	parser.AddCommand("install",
//  		"Install the network application",
//  		"The add command adds a file to the repository. Use -a to add all files",
//  		&installCommand)
//  }
//
//  func (x *InstallCommand) Execute(args []string) error {
//  	return x.Exec(args)
//  }
//
// Exec implements the business logic of command `install`
// func (x *InstallCommand) Exec(args []string) error {
// 	clis.Setup(fmt.Sprintf("%s::%s", progname, "install"), Opts.Verbose)
// 	fmt.Fprintf(os.Stderr, "Install the network application\n")
// 	// fmt.Fprintf(os.Stderr, "Copyright (C) 2022, Myself <me@mine.org>\n\n")
// 	// clis.Verbose(1, "Doing Install, with %+v, %+v\n", Opts, args)
// 	// fmt.Println(x.Dir, x.Suffix)
// 	// // err := ...
// 	// // clis.WarnOn("Doing Install", err)
// 	// // or,
// 	// // clis.AbortOn("Doing Install", err)
// 	return nil
// }
// Template for "install" CLI handling ends here

// Template for "publish" CLI handling starts here
////////////////////////////////////////////////////////////////////////////
// Program: redo
// Purpose: global option redo
// Authors: Myself <me@mine.org> (c) 2022, All rights reserved
////////////////////////////////////////////////////////////////////////////

//  package main

//  import (
//	"fmt"
//	"os"
//
//	"github.com/go-easygen/go-flags/clis"
//  )

// *** Sub-command: publish ***

////////////////////////////////////////////////////////////////////////////
// Constant and data type/structure definitions

// The PublishCommand type defines all the configurable options from cli.
//  type PublishCommand struct {
//  	Dir	string	`short:"d" long:"dir" description:"publish dir" required:"true"`
//  	Suffix	[]string	`short:"s" long:"suffix" description:"source file suffix for publish" choice:".go" choice:".c" choice:".h"`
//  	Out	string	`short:"o" long:"out" description:"output filename"`
//
//  	// Example of positional arguments
//  Args struct {
//    ID   string
//    Num  int
//    Rest []string
//  } `positional-args:"yes" required:"yes"`

//  }

//
//  var publishCommand PublishCommand
//
//  func init() {
//  	parser.AddCommand("publish",
//  		"Publish the network application",
//  		"Publish the built network application to central repo",
//  		&publishCommand)
//  }
//
//  func (x *PublishCommand) Execute(args []string) error {
//  	return x.Exec(args)
//  }
//
// Exec implements the business logic of command `publish`
// func (x *PublishCommand) Exec(args []string) error {
// 	clis.Setup(fmt.Sprintf("%s::%s", progname, "publish"), Opts.Verbose)
// 	fmt.Fprintf(os.Stderr, "Publish the network application\n")
// 	// fmt.Fprintf(os.Stderr, "Copyright (C) 2022, Myself <me@mine.org>\n\n")
// 	// clis.Verbose(1, "Doing Publish, with %+v, %+v\n", Opts, args)
// 	// fmt.Println(x.Dir, x.Suffix, x.Out, x.Args)
// 	// // err := ...
// 	// // clis.WarnOn("Doing Publish", err)
// 	// // or,
// 	// // clis.AbortOn("Doing Publish", err)
// 	return nil
// }
// Template for "publish" CLI handling ends here
