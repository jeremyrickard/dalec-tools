package cmd

import (
	"github.com/jeremyrickard/dalec-tools/cmd/buildinfo"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

const (
	rootName             = "dt"
	rootShortDescription = "A collection of tools for interacting with DALEC build defs"
	rootLongDescription  = "A collection of tools for interacting with DALEC build defs"
)

var (
	debug bool
	trace bool
)

// NewRootCmd returns the root command for dt
func NewRootCmd() *cobra.Command {
	rootCmd := &cobra.Command{
		Use:   rootName,
		Short: rootShortDescription,
		Long:  rootLongDescription,
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			if debug {
				log.SetLevel(log.DebugLevel)
			}
			if trace {
				log.SetLevel(log.TraceLevel)
			}
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return cmd.Usage()
		},
	}

	p := rootCmd.PersistentFlags()
	p.BoolVar(&debug, "debug", false, "enable debug level logging")
	p.BoolVar(&trace, "trace", false, "enable trace level logging")

	rootCmd.AddCommand(buildinfo.NewCmd())

	return rootCmd
}
