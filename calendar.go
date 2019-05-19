package calendar

import (
	"strings"
	"time"
)

func max(i1, i2 int) int {
	if i1 > i2 {
		return i1
	}
	return i2
}

func center(s string, width int) string {
	length := len(s)
	if (width == length){
		return s
	}
	padding := (width - length) / 2
	var builder strings.Builder
	if (width - length) % 2 == 0 {
		empty := strings.Repeat(" ", padding)
		builder.WriteString(empty)
		builder.WriteString(s)
		builder.WriteString(empty)
	} else {
		empty1 := strings.Repeat(" ", padding)
		empty2 := strings.Repeat(" ", padding + 1)
		builder.WriteString(empty1)
		builder.WriteString(s)
		builder.WriteString(empty2)
	}
	return builder.String()
}

var lengths [12]int = [12]int{31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}
var weekDays [7]string = [7]string{"Su", "Mo", "Tu", "We", "Th", "Fr", "Sa"}

type Calendar struct {
	Month time.Month
	Year  int
}

type CellDrawFunc func(day int, Month int, year int) string

func (c Calendar) FirstDay() time.Weekday {
	date := time.Date(c.Year, c.Month, 1, 0, 0, 0, 0, time.UTC)
	return date.Weekday()
}

func (c Calendar) Draw(f CellDrawFunc) string {
	length := lengths[c.Month-1]
	var results []string
	maxSize := 2
	for i := 1; i <= length; i++ {
		cellContent := f(i, int(c.Month), c.Year)
		results = append(results, cellContent)
		maxSize = max(maxSize, len(cellContent))
	}

	var buffer []string
	for _, day := range weekDays {
		buffer = append(buffer, center(day, maxSize), " ")
	}
	buffer = append(buffer, "\n")

	emptyDays := int(c.FirstDay())
	for i := 0; i < emptyDays; i++ {
		buffer = append(buffer, center(" ", maxSize), " ")
	}

	var i int
	for i = 0; i < length; i++ {
		buffer = append(buffer, center(results[i], maxSize))
		if (i+emptyDays) % 7 == 6 {
			buffer = append(buffer, "\n")
		} else {
			buffer = append(buffer, " ")
		}
	}

	for remain := i + emptyDays; remain % 7 != 0; remain++ {
		buffer = append(buffer, center(" ", maxSize))
		if remain % 7 != 6 {
			buffer = append(buffer, " ")
		}
	}

	return strings.Join(buffer, "")
}
