package internal

import (
	"fmt"
	"image"
	"image/color/palette"
	"image/draw"
	"image/gif"
	_ "image/png"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func parseFrameCount(stdout string) (int, error) {
	for line := range strings.SplitSeq(strings.TrimSpace(stdout), "\n") {
		if num, err := strconv.Atoi(strings.TrimSpace(line)); err == nil {
			return num, nil
		}
	}
	return 0, fmt.Errorf("output lacks a framecount number \"%s\"", stdout)
}

func SpritesheetToGIF(sheetPath, outputPath string, frameCount int) error {
	file, err := os.Open(filepath.Clean(sheetPath))
	if err != nil {
		return fmt.Errorf("failed to open sheet: %w", err)
	}
	defer func() { _ = file.Close() }()

	srcImg, _, err := image.Decode(file)
	if err != nil {
		return fmt.Errorf("failed to decode PNG: %w", err)
	}

	bounds := srcImg.Bounds()
	frameWidth := bounds.Dx() / frameCount
	frameHeight := bounds.Dy()

	var frames []*image.Paletted
	var delays []int

	for c := range frameCount {
		cropRect := image.Rect(c*frameWidth, 0, (c+1)*frameWidth, frameHeight)

		palettedFrame := image.NewPaletted(image.Rect(0, 0, frameWidth, frameHeight), palette.Plan9)

		draw.Draw(palettedFrame, palettedFrame.Bounds(), srcImg, cropRect.Min, draw.Over)

		frames = append(frames, palettedFrame)
		delays = append(delays, 10)
	}

	outFile, err := os.Create(filepath.Clean(outputPath))
	if err != nil {
		return fmt.Errorf("failed to create output file: %w", err)
	}
	defer func() { _ = outFile.Close() }()

	outGIF := &gif.GIF{
		Image:     frames,
		Delay:     delays,
		LoopCount: 0,
	}

	if err := gif.EncodeAll(outFile, outGIF); err != nil {
		return fmt.Errorf("failed to encode GIF: %w", err)
	}

	return nil
}
