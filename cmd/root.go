package cmd

import (
	"fmt"

	"github.com/Onboardbase/pipemuta/pkg/pipemuta"
	"github.com/Onboardbase/pipemuta/pkg/utils"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "pipemuta -i pipemuta.yaml -t github -o gitlab-ci.yml",
	Short: "Generate pipeline configurations from a base base config source",
	Long:  "Generate pipeline configurations from a base pipeline source.",
	Run:   initializePipeMuta,
}

func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func initializePipeMuta(cmd *cobra.Command, args []string) {
	outfile := cmd.Flag("output").Value.String()
	if outfile == "" {
		utils.ErrExit(fmt.Errorf("output file is required"), 1)
	}

	infile := cmd.Flag("input").Value.String()
	if infile == "" {
		utils.ErrExit(fmt.Errorf("inputfile is required"), 1)
	}

	configtype := cmd.Flag("type").Value.String()
	if configtype == "" {
		utils.ErrExit(fmt.Errorf("source configuration type is required"), 1)
	}

	muta := pipemuta.NewPipeMuta(infile, outfile, configtype)
	muta.Transmute()
}

func init() {
	rootCmd.Flags().StringP("output", "o", "out.yaml", "the file name/path to output to e.g -o schema.go")
	rootCmd.Flags().StringP("type", "t", "github", "a supported CI/CD service provider e.g -p gitlab")
	rootCmd.Flags().StringP("input", "i", "pipemuta.yml", "a valid Pipemuta configuration")
}
