package conversion

import (
	"errors"
	"strconv"
)

func ConvertStringsToFloats(stringValues []string) ([]float64, error) {
	// Convert the price strings to float64
	floatValues := make([]float64, len((stringValues)))
	for index, line := range stringValues {
		floatPrice, err := strconv.ParseFloat(line, 64)

		if err != nil {
			return nil, errors.New("Strings to floats conversion failed")
		}

		floatValues[index] = floatPrice
	}

	return floatValues, nil
}
