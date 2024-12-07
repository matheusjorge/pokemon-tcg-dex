package utils

import (
	"github.com/k0kubun/go-ansi"
	"github.com/schollz/progressbar/v3"
)

func CreateProgressBar(length int, msg string) *progressbar.ProgressBar {
	bar := progressbar.NewOptions(
		length,
		progressbar.OptionSetWriter(ansi.NewAnsiStdout()),
		progressbar.OptionEnableColorCodes(true),
		// progressbar.OptionShowBytes(true),
		progressbar.OptionSetWidth(15),
		progressbar.OptionSetDescription(msg),
		progressbar.OptionSetTheme(progressbar.Theme{
			Saucer:        "[green]=[reset]",
			SaucerHead:    "[green]>[reset]",
			SaucerPadding: " ",
			BarStart:      "[",
			BarEnd:        "]",
		}),
		progressbar.OptionShowElapsedTimeOnFinish(),
		progressbar.OptionClearOnFinish(),
	)

	return bar
}
