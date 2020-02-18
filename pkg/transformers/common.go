package transformers

import (
	"encoding/csv"
	"log"
	"os"
)

func parseCsv(fileName string) ([][]string, error) {

	f, err := os.Open(fileName)
	if err != nil {
		log.Fatal("Unable to read input file "+fileName, err)
	}
	defer f.Close()

	csvReader := csv.NewReader(f)
	records, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal("Unable to parse file as CSV for "+fileName, err)
	}

	return records, err
}
