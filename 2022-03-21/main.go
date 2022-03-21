package march212022

import (
	"fmt"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

const fullDayInMinutes = 1440

var reTime = regexp.MustCompile(`^\d{2}:\d{2}$`)

func Tasks() {
	list := []string{"01:00", "08:15", "11:30", "13:45", "14:10", "20:05"}
	interval, err := smallestInterval(list)
	if err != nil {
		fmt.Printf("whoops, something bad happened: %s", err)
		return
	}

	fmt.Printf("Got result: %s\n", interval)
}

func smallestInterval(in []string) (string, error) {
	entries := make([]int, len(in))
	for i, t := range in {
		entry, err := entryToNumber(t)
		if err != nil {
			return "", fmt.Errorf("converting entry [%s] to number: %w", t, err)
		}
		entries[i] = entry
	}

	sort.Ints(entries)

	smallest := fullDayInMinutes
	for i, sorted := range entries[1:] {
		if sorted-entries[i] < smallest {
			smallest = sorted - entries[i]
		}
	}

	return numberToTimeString(smallest), nil
}

func numberToTimeString(in int) string {
	hours := in / 60
	minutes := in - hours*60
	if hours > 0 {
		return fmt.Sprintf("%d hours and %d minutes", hours, minutes)
	}

	return fmt.Sprintf("%d minutes", minutes)
}

func entryToNumber(in string) (int, error) {
	if !reTime.MatchString(in) {
		return 0, fmt.Errorf("incoming string does not match expected format. Got [%s]", in)
	}

	parts := strings.Split(in, ":")
	hours, err := strconv.Atoi(parts[0])
	if err != nil {
		return 0, fmt.Errorf("converting hour [%s] to int: %w", parts[0], err)
	}

	minutes, err := strconv.Atoi(parts[1])
	if err != nil {
		return 0, fmt.Errorf("converting minute [%s] to int: %w", parts[1], err)
	}

	return hours*60 + minutes, nil
}
