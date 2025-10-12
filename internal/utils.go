/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package internal

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

// GetDateForPuzzle returns today's day of month limited to the AoC range [1, 25].
func GetDateForPuzzle(args []string) (int, int, error) {
    now := time.Now()
    year, month, day := now.Year(), int(now.Month()), now.Day()
    if day > 25 { day = 25 }

    if _, err := validateDay(day); err != nil {
        return 0, 0, err
    }

    if month < 12 {
        year -= 1
    }

	if len(args) == 0 {
		return year, day, nil
	}

    firstParam, errA := strconv.Atoi(args[0])
  	if errA != nil {
    	return 0, 0, fmt.Errorf("invalid day %d: %w", firstParam, errA)
  	}

	// override if positional arg provided
    if len(args) == 1 {
        if firstParam > 25 {
            year = firstParam
        } else {
            day = firstParam
        }
    }

    if len(args) == 2 {
        secondParam, errB := strconv.Atoi(args[1])
  	    if errB != nil {
    	    return 0, 0, fmt.Errorf("invalid day %d: %w", secondParam, errB)
  	    }
  	    year = firstParam
        day = secondParam
    }

    if _, err := validateYear(year); err != nil {
        return 0, 0, err
    }

    return year, day, nil
}

func validateDay(day int) (int, error) {
  	if day < 1 || day > 25 {
    	return 0, fmt.Errorf("day must be between 1 and 25, received %d", day)
  	}

  	return day, nil
}

func validateYear(year int) (int, error) {
	minYear := 2015

  	if year < minYear {
    	return 0, fmt.Errorf("year must be greater than %d, received %d", minYear, year)
  	}

  	return year, nil
}

func DirExists(path string) bool {
    info, err := os.Stat(path)
    if err != nil {
        return false
    }
    return info.IsDir()
}

func FileExists(path string) bool {
    info, err := os.Stat(path)
    if err != nil {
        return false
    }
    return !info.IsDir()
}

