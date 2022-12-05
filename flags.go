package whatever

import (
	"flag"
	"fmt"
	"io"
)

const helpMessage = `
This is a protoc plugin that is used to generate files from template using protocol buffer descriptors as template inputs.

EXAMPLE: Generate HTML docs
protoc --whatever_out=. --whatever_opt=template.tpl,README.md protos/*.proto
`

type Flags struct {
	appName     string
	flagSet     *flag.FlagSet
	err         error
	showHelp    bool
	showVersion bool
	writer      io.Writer
}

// Code returns the status code to exit with after handling the supplied flags
func (f *Flags) Code() int {
	if f.err != nil {
		return 1
	}

	return 0
}

// HasMatch returns whether or not the supplied args are matches. For example, passing `--help` will match, or some
// unknown parameter, but passing nothing will not.
func (f *Flags) HasMatch() bool {
	return f.ShowHelp() || f.ShowVersion()
}

// ShowHelp determines whether or not to show the help message
func (f *Flags) ShowHelp() bool {
	return f.err != nil || f.showHelp
}

// ShowVersion determines whether or not to show the version message
func (f *Flags) ShowVersion() bool {
	return f.showVersion
}

// PrintHelp prints the usage string including all flags to the `io.Writer` that was supplied to the `Flags` object.
func (f *Flags) PrintHelp() {
	fmt.Fprintf(f.writer, "Usage of %s:\n", f.appName)
	fmt.Fprintf(f.writer, "%s\n", helpMessage)
	fmt.Fprintf(f.writer, "FLAGS\n")
	f.flagSet.PrintDefaults()
}

// PrintVersion prints the version string to the `io.Writer` that was supplied to the `Flags` object.
func (f *Flags) PrintVersion() {
	fmt.Fprintf(f.writer, "%s version %s\n", f.appName, VERSION)
}

// ParseFlags parses the supplied options are returns a `Flags` object to the caller.
//
// Parameters:
//   - `w` - the `io.Writer` to use for printing messages (help, version, etc.)
//   - `args` - the set of args the program was invoked with (typically `os.Args`)
func ParseFlags(w io.Writer, args []string) *Flags {
	f := Flags{appName: args[0], writer: w}

	f.flagSet = flag.NewFlagSet(args[0], flag.ContinueOnError)
	f.flagSet.BoolVar(&f.showHelp, "help", false, "Show this help message")
	f.flagSet.BoolVar(&f.showVersion, "version", false, fmt.Sprintf("Print the current version (%v)", VERSION))
	f.flagSet.SetOutput(w)

	// prevent showing help on parse error
	f.flagSet.Usage = func() {}

	f.err = f.flagSet.Parse(args[1:])
	return &f
}
