package handlers

import (
	"bytes"
	"encoding/base64"
	"image"
	_ "image/jpeg" // Register JPEG decoder
	_ "image/png"  // Register PNG decoder
	"os"
	"strings"

	webp "github.com/chai2010/webp"
)

// DecodeBase64ToWebP converts a Base64-encoded image to WebP format
func DecodeBase64ToWebP(base64Data, outputPath string) error {
	// Clean data URL prefix if present
	if idx := strings.Index(base64Data, ","); idx != -1 {
		base64Data = base64Data[idx+1:]
	}

	// Decode Base64 data
	decoded, err := base64.StdEncoding.DecodeString(base64Data)
	if err != nil {
		return err
	}

	// Decode image using bytes.Reader for efficiency
	img, _, err := image.Decode(bytes.NewReader(decoded))
	if err != nil {
		return err
	}

	// Create output file
	outFile, err := os.Create(outputPath+".webp")
	if err != nil {
		return err
	}
	defer outFile.Close()

	// Configure WebP encoding options
	options := &webp.Options{
		Lossless: true,   // Lossy compression
		//Quality:  quality, // Range: 0.0 (lowest) to 100.0 (highest)
	}

	// Encode to WebP with specified quality
	if err := webp.Encode(outFile, img, options); err != nil {
		return err
	}

	return nil
}

// Example usage:
// func main() {
// 	err := DecodeBase64ToWebP(
// 		"data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAAAEAAAABCAYAAAAfFcSJAAAADUlEQVR42mP8z8BQDwAEhQGAhKmMIQAAAABJRU5ErkJggg==",
// 		"output.webp",
// 		85.0, // Quality setting
// 	)
// 	if err != nil {
// 		panic(err)
// 	}
// }