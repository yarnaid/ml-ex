package main

import (
	"generators"
	"os"
	"shared"

	"github.com/spf13/cobra"
)

var generateCmd = &cobra.Command{
	Use: "generate",
	Run: func(cmd *cobra.Command, args []string) {
		f, err := os.Create(cfg.SampleOutFileName)
		if err != nil {
			panic(err)
		}
		W := shared.NewDenseVec(1)
		W.SetVec(0, 0.5)
		generator := generators.LinRegGenerator{
			Number: cfg.SamplesNumber,
			Dim:    1,
			W:      W,
			B:      1.,
			Sigma:  1.,
			Area: shared.NDArea{
				Ranges: []shared.LinRange{
					{Start: -10., End: 10.},
				},
			},
		}
		data := generator.Generate()
		generators.DumpCSV(f, data)
	},
}

func init() {
	rootCmd.AddCommand(generateCmd)
}
