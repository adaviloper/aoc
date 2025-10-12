/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os/exec"

	"github.com/adaviloper/advent-of-code/aoc/internal"
	"github.com/spf13/cobra"
)

// runCmd represents the run command
var runCmd = &cobra.Command{
	Use:   "run",
    Short: "Run a day's Advent of Code solution with Deno",
    Long: `Runs the solution script for a given Advent of Code day using Deno.

Defaults and date handling:
- With no arguments, uses today's date and infers the AoC year (December => current year; otherwise previous year)
- You can override by providing positional arguments: aoc run [<year> <day>] or aoc run [<day>]

Execution details:
- Executes <BASE>/<YEAR>/<DD>/main.<lang>, where <BASE> and <lang> come from ~/.config/aoc/config.json
- Uses Deno with read permissions and quiet output
- Ensure Deno is installed and on your PATH

Examples:
- aoc run            # run today's script
- aoc run 2024 05    # run 2024 Day 05
- aoc run 17         # run Day 17 of the inferred AoC year`,
    RunE: func(cmd *cobra.Command, args []string) error {
        year, day, err := internal.GetDateForPuzzle(args)
        if err != nil {
            return err
        }

        return runPuzzleSection(year, day, useReal)
    },
}

func init() {
	rootCmd.AddCommand(runCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// runCmd.PersistentFlags().String("foo", "", "A help for foo")

    // Cobra supports local flags which will only run when this command
    // is called directly.
    runCmd.Flags().BoolVarP(&useReal, "real", "R", false, "Use real input instead of test input")
}

var useReal bool

func runPuzzleSection(year int, day int, useReal bool) error {
    filePath := fmt.Sprintf("%s/%d/%02d/main.%s", cfg.BaseDirectory, year, day, cfg.TemplateLang)
	if !internal.FileExists(filePath) {
		fmt.Printf("[%s] does not exist\n", filePath)
		return fmt.Errorf("[%s] does not exist\n", filePath)
	}

    args := []string{"run", "--allow-read", "--quiet", filePath}
    if useReal {
        args = append(args, "--real")
    }

    command := exec.Command("deno", args...)
    output, err := command.CombinedOutput()
    if err != nil {
        return fmt.Errorf("deno failed: %w\n%s", err, string(output))
    }

	fmt.Printf("Results for December %d, %d:\n", day, year)
    fmt.Printf("%s\n", string(output))

	return nil
}
