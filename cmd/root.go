/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
    "os"

    "github.com/spf13/cobra"
    "github.com/adaviloper/aoc/config"
)

// Configuration struct and instance moved to package config

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "aoc",
	Short: "Quick scaffolding utility for setting up Advent of Code puzzles",
	Long: `aoc is a tiny CLI that scaffolds and runs Advent of Code puzzles. It creates a year/day folder structure, fetches your personal puzzle input, generates starter files, and can execute your solution script with Deno.

	What it does:
	- Creates <BASE>/<YEAR>/<DD>/ under a configured base directory
	- Downloads Advent of Code input using your AOC_SESSION cookie and writes data.ts
	- Generates a minimal main.<lang> and test.<lang> file for you to start coding
	- Runs the day’s script with Deno (intended for TypeScript templates)

	Defaults and date handling:
	- Without arguments, uses “today” for day and picks the AoC year (December => current year; otherwise previous year)
	- Intended for Dec 1–25; you can override the inferred date with arguments

	Requirements:
	- Deno installed and on PATH
	- Environment variable AOC_SESSION set to your session cookie value from adventofcode.com

	Configuration:
	- Reads $HOME/.config/aoc/config.json
  {
  "base_directory": "~/Code/advent-of-code",
  "template_language": "ts"
  }

	Files and layout created:
	- <BASE>/<YEAR>/<DD>/data.ts  // puzzle input as a string export
	- <BASE>/<YEAR>/<DD>/main.ts  // starter solution template
	- <BASE>/<YEAR>/<DD>/test.ts  // place sample input or tests
	Existing files are never overwritten.

	Commands:
	- aoc build [<year> <day>]   # scaffold files and fetch input
	- aoc run [<year> <day>]     # execute the day’s script via Deno

	Examples:
	- aoc build          # scaffold for today (AoC rules for year/day)
	- aoc build 5        # scaffold for Day 5 of the current year
	- aoc build 2024     # scaffold for 2024 for the current day
	- aoc build 2024 5   # scaffold for 2024 Day 5
	- aoc run            # run today’s script
	- aoc run 17         # run Day 17 of the current year 
	- aoc run 2023       # run 2023 for the current day  
	- aoc run 2023 17    # run 2023 Day 17                

	Tip: If you primarily solve in TypeScript with Deno, leave template_language as "ts".`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
    // Initialize global configuration
    config.Init()

    if err := rootCmd.Execute(); err != nil {
        os.Exit(1)
    }
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.aoc.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}


