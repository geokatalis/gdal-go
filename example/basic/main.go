package main

import (
	"fmt"

	gdal "github.com/geokatalis/gdal-go"
)

func main() {
	gdal.GDALAllRegister()
	fmt.Println("GDAL version:", gdal.GDALVersionInfo("--version"))
	fmt.Println("Release name:", gdal.GDALVersionInfo("RELEASE_NAME"))
	fmt.Println("Driver count:", gdal.GDALGetDriverCount())
}
