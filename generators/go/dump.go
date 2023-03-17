package generators

import (
	"encoding/csv"
	"fmt"
	"io"
	"shared"
)

func DumpCSV(w io.Writer, points []*shared.DenseVec) {
	writer := csv.NewWriter(w)
	for _, point := range points {
		var row []string
		for i := 0; i < point.Len(); i++ {
			row = append(row, fmt.Sprintf("%f", point.AtVec(i)))
		}
		if err := writer.Write(row); err != nil {
			panic(err)
		}
	}
	writer.Flush()
	if err := writer.Error(); err != nil {
		panic(err)
	}
}
