package cmd

import (
	"fmt"
	"io"
	"os"

	"github.com/Onboardbase/pipemuta/pkg/utils"
	"github.com/a-h/generate"
	"github.com/spf13/cobra"
)

func generateSchema(cmd *cobra.Command, args []string) {
	file := cmd.Flag("file").Value.String()
	if file == "" {
		utils.ErrExit(fmt.Errorf("source configuration type is required"), 1)
	}

	pack_name := cmd.Flag("package").Value.String()
	if pack_name == "" {
		pack_name = "main"
	}
	dest := cmd.Flag("outfile").Value.String()
	if dest == "" {
		dest = "schema.go"
	}
	// generate the golang structs for both the source and destination pipelines
	schemas, err := generate.ReadInputFiles([]string{file}, true)
	if err != nil {
		fmt.Fprintf(os.Stderr, err.Error())
		os.Exit(1)
	}

	g := generate.New(schemas...)

	err = g.CreateTypes()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Failure generating structs: ", err)
		os.Exit(1)
	}

	var w io.Writer = os.Stdout

	if dest != "" {
		w, err = os.Create(dest)

		if err != nil {
			fmt.Fprintln(os.Stderr, "Error opening output file: ", err)
			return
		}
	}

	generate.Output(w, g, pack_name)
}

var generatorCmd = &cobra.Command{
	Use:   "generate",
	Short: "pipemuta generator -f gitlab-ci.yml -o schema.go",
	Long:  "pipemuta generator -f gitlab-ci.yml -o schema.go",
	Run:   generateSchema,
}

func init() {
	generatorCmd.Flags().StringP("outfile", "o", "", "the file name/path to output to e.g -o schema.go")
	generatorCmd.Flags().StringP("package", "p", "", "a go package name e.g -p package_name")
	generatorCmd.Flags().StringP("file", "f", "", "a valid JSON schema source of the schema to generate")

	rootCmd.AddCommand(generatorCmd)
}
