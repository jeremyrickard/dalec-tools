package buildinfo

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/Azure/dalec"
	"github.com/spf13/cobra"
)

func NewCmd() *cobra.Command {
	bc := buildInfoCmd{}
	cmd := &cobra.Command{
		Use:   "build-info",
		Short: "Get the VERSION and COMMIT from a DALEC build def",
		Long:  "Get the VERSION and COMMIT from a DALEC build def",
		RunE:  bc.run,
	}

	f := cmd.Flags()
	f.StringVar(&bc.def, "definition", "", "path to the definition file")

	return cmd
}

type buildInfoCmd struct {
	def string
}

type buildinfo struct {
	Version string `json:"version"`
	Commit  string `json:"commit"`
}

func (bic *buildInfoCmd) run(cmd *cobra.Command, _ []string) error {
	b, err := os.ReadFile(bic.def)
	if err != nil {
		return err
	}
	spec, err := dalec.LoadSpec(b)
	if err != nil {
		return err
	}
	args := spec.Args

	bi := buildinfo{}
	bi.Version = args["VERSION"]
	bi.Commit = args["COMMIT"]

	b, err = json.Marshal(bi)
	if err != nil {
		return err
	}
	fmt.Fprint(cmd.OutOrStdout(), string(b))
	return nil
}
