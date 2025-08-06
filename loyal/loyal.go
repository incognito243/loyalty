package loyal

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

// ParseLog reads log lines in the format "Timestamp, PageId, CustomerId" from an io.Reader and returns a map
// mapping each CustomerId to a set of unique PageIds they have visited.
func ParseLog(r io.Reader) (map[string]map[string]struct{}, error) {
	scanner := bufio.NewScanner(r)
	visits := make(map[string]map[string]struct{})
	lineNo := 0

	for scanner.Scan() {
		lineNo++
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}
		parts := strings.Split(line, ",")
		if len(parts) != 3 {
			return nil, fmt.Errorf("invalid record on line %d: %q", lineNo, line)
		}
		pageID := strings.TrimSpace(parts[1])
		custID := strings.TrimSpace(parts[2])
		if pageID == "" || custID == "" {
			return nil, fmt.Errorf("empty field on line %d: %q", lineNo, line)
		}
		if visits[custID] == nil {
			visits[custID] = make(map[string]struct{})
		}
		visits[custID][pageID] = struct{}{}
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return visits, nil
}

// FindLoyalCustomers returns all CustomerIds who:
//  1. appear in both logs (day1, day2)
//  2. have at least 2 distinct PageIDs across both days
func FindLoyalCustomers(day1, day2 io.Reader) ([]string, error) {
	v1, err := ParseLog(day1)
	if err != nil {
		return nil, fmt.Errorf("day1 parse: %w", err)
	}
	v2, err := ParseLog(day2)
	if err != nil {
		return nil, fmt.Errorf("day2 parse: %w", err)
	}

	var loyal []string
	for cust, pages1 := range v1 {
		pages2, ok := v2[cust]
		if !ok {
			continue
		}
		uniq := make(map[string]struct{}, len(pages1)+len(pages2))
		for p := range pages1 {
			uniq[p] = struct{}{}
		}
		for p := range pages2 {
			uniq[p] = struct{}{}
		}
		if len(uniq) >= 2 {
			loyal = append(loyal, cust)
		}
	}
	return loyal, nil
}
