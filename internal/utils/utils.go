package utils

import (
	"fmt"
	"math"
	"os"
	"path"
)

type SizeValueSet struct {
	h1  int
	h2  int
	h3  int
	h4  int
	h5  int
	h6  int
	p   int
	smp int
	xsp int
}

func CreateWebSizeValues(baseSize int, ratio float64) SizeValueSet {
	highValues := make([]int, 0, 6)
	lastValue := baseSize
	for i := 0; i < 6; i++ {
		value := float64(lastValue) * ratio
		result := int(math.Round(value))
		highValues = append(highValues, result)
		lastValue = result
	}

	lowValues := make([]int, 0, 3)
	lastValue = baseSize
	for i := 0; i < 2; i++ {
		value := float64(lastValue) / ratio
		result := int(math.Round(value))
		lowValues = append(lowValues, result)
		lastValue = result
	}

	values := SizeValueSet{
		h1:  highValues[5],
		h2:  highValues[4],
		h3:  highValues[3],
		h4:  highValues[2],
		h5:  highValues[1],
		h6:  highValues[0],
		p:   baseSize,
		smp: lowValues[0],
		xsp: lowValues[1],
	}
	return values
}

func FormatWebSizes(values SizeValueSet) string {
	content := fmt.Sprintf("- %s: %dpx\n", " h1", values.h1) +
		fmt.Sprintf("- %s: %dpx\n", " h2", values.h2) +
		fmt.Sprintf("- %s: %dpx\n", " h3", values.h3) +
		fmt.Sprintf("- %s: %dpx\n", " h4", values.h4) +
		fmt.Sprintf("- %s: %dpx\n", " h5", values.h5) +
		fmt.Sprintf("- %s: %dpx\n", " h6", values.h6) +
		fmt.Sprintf("- %s: %dpx\n", "  p", values.p) +
		fmt.Sprintf("- %s: %dpx\n", "smp", values.smp) +
		fmt.Sprintf("- %s: %dpx\n", "xsp", values.xsp)
	return content
}

func SaveResultToFile(filename, content string) error {
	filename = filename + ".txt"
	file, err := os.Create(path.Join(".", filename))
	defer file.Close()
	if err != nil {
		return err
	}
	_, err = file.WriteString(content)
	if err != nil {
		return err
	}
	return nil
}
