/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/adaviloper/advent-of-code/aoc/internal"
	"github.com/spf13/cobra"
)

// buildCmd represents the build command
var buildCmd = &cobra.Command{
    Use:   "build",
    Short: "Create today's AoC directory and data.ts",
    Long:  "Creates <YYYY>/<DD>/ and writes data.ts with Advent of Code input using AOC_SESSION.",
    Args: cobra.MaximumNArgs(2),
    RunE: func(cmd *cobra.Command, args []string) error {
        year, day, err := internal.GetDateForPuzzle(args)
        if err != nil {
        	return err
        }

        // Resolve base directory for assumed year
        yearDir, err := resolveYearRootDirectory(cfg.BaseDirectory, year)
        if err != nil {
            return err
        }

        // Create day directory (01-25)
        dayDir := filepath.Join(yearDir, fmt.Sprintf("%02d", day))
        if err := os.MkdirAll(dayDir, 0o755); err != nil {
            return fmt.Errorf("failed to create day directory %s: %w", dayDir, err)
        }

        dataFilePath := filepath.Join(dayDir, fmt.Sprintf("data.%s", cfg.TemplateLang))
        if internal.FileExists(dataFilePath) {
            // If it exists, do not overwrite to avoid losing edits
            fmt.Printf("data.%s already exists at %s, skipping.\n", cfg.TemplateLang, dataFilePath)
        } else {
            // Fetch input from Advent of Code
            session := os.Getenv("AOC_SESSION")
            if strings.TrimSpace(session) == "" {
                return errors.New("AOC_SESSION environment variable is not set. Get your session cookie from adventofcode.com and set AOC_SESSION")
            }

            input, err := fetchAocInput(year, day, session)
            if err != nil {
                return err
            }

            if err := writeTSDataFile(dataFilePath, input); err != nil {
                return err
            }
            fmt.Printf("Wrote %s\n", dataFilePath)
        }

        // Also create a test file with an empty string if it doesn't exist
        testFilePath := filepath.Join(dayDir, fmt.Sprintf("test.%s", cfg.TemplateLang))
        if internal.FileExists(testFilePath) {
            fmt.Printf("test.%s already exists at %s, skipping.\n", cfg.TemplateLang, testFilePath)
        } else {
            if err := writeTSDataFile(testFilePath, "update me"); err != nil {
                return err
            }
            fmt.Printf("Wrote %s\n", testFilePath)
        }

        createMainPuzzleFile(year, day)

        return nil
    },
}

func init() {
	rootCmd.AddCommand(buildCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// buildCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// buildCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func createMainPuzzleFile(year int, day int) {
    // Also create a test.ts file with an empty string if it doesn't exist
    // puzzleFilePath := filepath.Join(dayDir, fmt.Sprintf("main.%s", cfg.TemplateLang))
    fmt.Sprintf("hit %s", year, day)
    puzzleFilePath := fmt.Sprintf("%s/%d/%02d/main.%s", cfg.BaseDirectory, year, day, cfg.TemplateLang)
    if internal.FileExists(puzzleFilePath) {
        fmt.Printf("main.%s already exists at %s, skipping.\n", cfg.TemplateLang, puzzleFilePath)
    } else {
        if err := writeEmptyPuzzleFile(puzzleFilePath, year, day); err != nil {
            return
        }
        fmt.Printf("Wrote %s\n", puzzleFilePath)
    }
}

// resolveYearRootDirectory tries to locate the year directory (e.g., 2024) either
// in the current working directory or in the parent if running from the aoc CLI subfolder.
// If it doesn't exist, it will be created in the most reasonable place.
func resolveYearRootDirectory(baseDir string, year int) (string, error) {
    yearName := fmt.Sprintf("%d", year)

    // Prefer ./<year> if it exists
    dir := filepath.Join(baseDir, yearName)
    if internal.DirExists(dir) {
        return dir, nil
    }

    // Otherwise, create ./<year>
    if err := os.MkdirAll(dir, 0o755); err != nil {
        return "", fmt.Errorf("failed to create year directory %s: %w", dir, err)
    }
    return dir, nil
}

// fetchAocInput downloads the Advent of Code input for the given year and day
// using the provided session cookie value.
func fetchAocInput(year int, day int, session string) (string, error) {
    // url := fmt.Sprintf("https://adventofcode.com/%d/day/%d/input", year, day)
    url := fmt.Sprintf("https://adventofcode.com/%d/day/%d/input", year, day)

    req, err := http.NewRequest(http.MethodGet, url, nil)
    if err != nil {
        return "", fmt.Errorf("failed to create request: %w", err)
    }
    req.Header.Set("Cookie", "session="+session)
    req.Header.Set("User-Agent", "aoc-cli (+https://adventofcode.com)")

    client := &http.Client{Timeout: 30 * time.Second}
    resp, err := client.Do(req)
    if err != nil {
        return "", fmt.Errorf("failed to fetch input: %w", err)
    }
    defer resp.Body.Close()

    if resp.StatusCode != http.StatusOK {
        // Read at most some bytes to include in error for debugging
        body, _ := io.ReadAll(io.LimitReader(resp.Body, 2048))
        return "", fmt.Errorf("unexpected status %d from AoC. Body: %s", resp.StatusCode, strings.TrimSpace(string(body)))
    }

    bytes, err := io.ReadAll(resp.Body)
    if err != nil {
        return "", fmt.Errorf("failed to read input body: %w", err)
    }
    return string(bytes), nil
}

// writeTSDataFile writes a TypeScript module exporting the input as a template literal.
func writeTSDataFile(path string, input string) error {
    // Escape backticks and prevent template interpolation
    escaped := strings.ReplaceAll(input, "`", "\\`")
    escaped = strings.ReplaceAll(escaped, "${", "\\${")

    content := fmt.Sprintf("export const data = `\n%s`\n", escaped)

    if err := os.WriteFile(path, []byte(content), 0o644); err != nil {
        return fmt.Errorf("failed to write %s: %w", path, err)
    }
    return nil
}

func writeEmptyPuzzleFile(path string, year int, day int) error {
	content := fmt.Sprintf(`/**
 * https://adventofcode.com/%d/day/%02d
 */

import { data as realData } from './data.%s';
import { data as testData } from './test.%s';
import * as helpers from '../../../utils/helpers.%s';
import * as utils from '../../../utils/types.%s';

const input = Deno.args.includes('--real') ? realData : testData;

const p1 = (input) => {
  console.log('p1: ', input);
  return 'get to work';
};

const p2 = (input) => {
  console.log('p2: ', input);
  return 'get to work';
};

console.log(p1(input));
console.log(p2(input));
`, year, day, cfg.TemplateLang, cfg.TemplateLang, cfg.TemplateLang, cfg.TemplateLang)

    if err := os.WriteFile(path, []byte(content), 0o644); err != nil {
        return fmt.Errorf("failed to write %s: %w", path, err)
    }
    return nil
}
