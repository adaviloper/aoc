module aoc

go 1.24.6

require github.com/spf13/cobra v1.10.1

require (
	github.com/inconshreveable/mousetrap v1.1.0 // indirect
	github.com/spf13/pflag v1.0.9 // indirect
)

require (
	github.com/adaviloper/aoc v0.0.0-20251012233414-8eeafcbaf3d3
	internal/utils v1.0.0
)

replace internal/utils => ./internal/utils
