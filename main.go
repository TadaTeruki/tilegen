package main

import (
	"fmt"
	"image"
	"log"
	"math"
	"os"
	"path/filepath"
	"strconv"
	"sync"

	"github.com/disintegration/imaging"
)

func main() {
	projectName := os.Getenv("PROJECT_NAME")
	if projectName == "" {
		projectName = "dem"
	}
	imageName := os.Getenv("IMAGE_NAME")
	if imageName == "" {
		imageName = projectName + ".png"
	}
	zoomRangeMinStr := os.Getenv("ZOOM_RANGE_MIN")
	if zoomRangeMinStr == "" {
		zoomRangeMinStr = "0"
	}
	zoomRangeMin, _ := strconv.Atoi(zoomRangeMinStr)
	zoomRangeMaxStr := os.Getenv("ZOOM_RANGE_MAX")
	if zoomRangeMaxStr == "" {
		zoomRangeMaxStr = "65535"
	}
	zoomRangeMax, _ := strconv.Atoi(zoomRangeMaxStr)
	tileSizeStr := os.Getenv("TILE_SIZE")
	if tileSizeStr == "" {
		tileSizeStr = "256"
	}
	tileSize, _ := strconv.Atoi(tileSizeStr)

	src, err := imaging.Open(imageName)
	if err != nil {
		log.Fatalf("failed to open image: %v", err)
	}

	width := src.Bounds().Dx()

	numTilesX := int(math.Ceil(float64(width) / float64(tileSize)))
	maxZoom := int(math.Log2(float64(numTilesX))) + 1

	fmt.Println("Generating tiles...")

	for z := zoomRangeMin; z <= zoomRangeMax && z <= maxZoom; z++ {
		zoomScale := math.Pow(2, float64(z))
		zoomTileSize := int(float64(width) / zoomScale)
		wg := sync.WaitGroup{}
		for x := 0; x < int(zoomScale); x++ {
			for y := 0; y < int(zoomScale); y++ {
				go func(x, y, z int, zoomScale float64, zoomTileSize int) {
					tileX := x * zoomTileSize
					tileY := y * zoomTileSize

					tile := imaging.Crop(src, image.Rect(tileX, tileY, tileX+zoomTileSize, tileY+zoomTileSize))
					tile = imaging.Resize(tile, tileSize, tileSize, imaging.Lanczos)

					tileName := fmt.Sprintf("%s/%d/%d/%d.png", projectName, z, x, y)
					if err := os.MkdirAll(filepath.Dir(tileName), 0755); err != nil {
						log.Fatalf("failed to create directory: %v", err)
					}
					if err := imaging.Save(tile, tileName); err != nil {
						log.Fatalf("failed to save tile: %v", err)
					}
					fmt.Printf("Generated tile: %s\n", tileName)
					wg.Done()
				}(x, y, z, zoomScale, zoomTileSize)
				wg.Add(1)
			}
		}
		wg.Wait()
	}

	fmt.Println("Tiles generated successfully for all zoom levels.")
}
