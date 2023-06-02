package cmd

import (
	"fmt"

	"github.com/Onboardbase/pipemuta/pkg/pipemuta"
	"github.com/Onboardbase/pipemuta/pkg/utils"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "pipemuta -s gitlab -d github -f gitlab-ci.yml",
	Short: "move between multiple pipeline configurations",
	Long:  "Generate pipeline configurations from a base pipeline source.",
	Run:   initializePipeMuta,
}

func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func initializePipeMuta(cmd *cobra.Command, args []string) {
	source := cmd.Flag("source").Value.String()
	if source == "" {
		utils.ErrExit(fmt.Errorf("source configuration type is required"), 1)
	}

	destination := cmd.Flag("destination").Value.String()
	if destination == "" {
		utils.ErrExit(fmt.Errorf("destination configuration type is required"), 1)
	}

	file := cmd.Flag("file").Value.String()
	if file == "" {
		utils.ErrExit(fmt.Errorf("the source file path is required"), 1)
	}

	pmuta := pipemuta.NewPipeMuta(source, destination, file)
	pmuta.Transmute()
}

func init() {
	rootCmd.Flags().StringP("source", "s", "", "one of the supported pipeline configurations e.g -s github")
	rootCmd.Flags().StringP("destination", "d", "", "one of the suppported pipeline configurations e.g -d gitlab")
	rootCmd.Flags().StringP("file", "f", "", "the file to transmute from source configuration to destination configuration")
}
