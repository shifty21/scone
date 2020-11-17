package command

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"sort"
	"strings"
	"text/tabwriter"

	"github.com/mitchellh/cli"
	"github.com/shifty21/scone/config"
	"github.com/shifty21/scone/vaultinterface"
)

var commonCommands = []string{
	"grpc",
	"scone",
	"auto",
}

//RunOptions customizes output
type RunOptions struct {
	Stdout  io.Writer
	Stderr  io.Writer
	config  *config.Configuration
	options []vaultinterface.Option
	Vault   *vaultinterface.Vault
}

//Run runs a given command
func Run(args []string) int {
	var runOpts *RunOptions
	if runOpts == nil {
		runOpts = &RunOptions{}
	}

	if runOpts.Stdout == nil {
		runOpts.Stdout = os.Stdout
	}
	if runOpts.Stderr == nil {
		runOpts.Stderr = os.Stderr
	}
	runOpts.config = config.LoadConfig()
	v, err := vaultinterface.NewVaultInterface()
	if err != nil {
		fmt.Printf("Couldnt initialize vaultinterface %v", err)
		return 1
	}
	var options []vaultinterface.Option
	runOpts.options = options
	runOpts.Vault = v
	// Commands is the mapping of all the available commands.
	var Commands map[string]cli.CommandFactory
	Commands = map[string]cli.CommandFactory{
		"cas": func() (cli.Command, error) {
			return &CAS{
				RunOptions: runOpts,
			}, nil
		},
		"grpc": func() (cli.Command, error) {
			return &GRPC{
				RunOptions: runOpts,
			}, nil
		},
		"auto": func() (cli.Command, error) {
			return &Auto{
				RunOptions: runOpts,
			}, nil
		},
		"scone": func() (cli.Command, error) {
			return &Scone{
				RunOptions: runOpts,
			}, nil
		},
		"version": func() (cli.Command, error) {
			return &Version{
				RunOptions: runOpts,
			}, nil
		},
	}
	cli := &cli.CLI{
		Name:     "vault_init",
		Args:     args,
		Commands: Commands,
		HelpFunc: groupedHelpFunc(
			cli.BasicHelpFunc("vault_init"),
		),
		HelpWriter:                 runOpts.Stderr,
		HiddenCommands:             []string{"version"},
		Autocomplete:               true,
		AutocompleteNoDefaultFlags: true,
	}

	exitCode, err := cli.Run()
	if err != nil {
		fmt.Fprintf(runOpts.Stderr, "Error while executing command %v", err)
	}
	return exitCode

}

func groupedHelpFunc(f cli.HelpFunc) cli.HelpFunc {
	return func(commands map[string]cli.CommandFactory) string {
		var b bytes.Buffer
		tw := tabwriter.NewWriter(&b, 0, 2, 6, ' ', 0)

		fmt.Fprintf(tw, "Usage: vault_init <command> [args]\n\n")
		fmt.Fprintf(tw, "Common commands:\n")
		for _, v := range commonCommands {
			printCommand(tw, v, commands[v])
		}

		otherCommands := make([]string, 0, len(commands))
		for k := range commands {
			found := false
			for _, v := range commonCommands {
				if k == v {
					found = true
					break
				}
			}

			if !found {
				otherCommands = append(otherCommands, k)
			}
		}
		sort.Strings(otherCommands)

		fmt.Fprintf(tw, "\n")
		fmt.Fprintf(tw, "Other commands:\n")
		for _, v := range otherCommands {
			printCommand(tw, v, commands[v])
		}

		tw.Flush()

		return strings.TrimSpace(b.String())
	}
}

func printCommand(w io.Writer, name string, cmdFn cli.CommandFactory) {
	cmd, err := cmdFn()
	if err != nil {
		panic(fmt.Sprintf("failed to load %q command: %s", name, err))
	}
	fmt.Fprintf(w, "    %s\t%s\n", name, cmd.Synopsis())
}
