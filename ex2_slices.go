package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {

    // initialize the image using make
    image := make([][]uint8, dy)

    // iterate over each pixel
    for x := 0; x < dy; x++ {
        image[x] = make([]uint8, dx)
        for y := 0; y < dx; y++ {

            // set a color value for pixels, try various func
            image[x][y] = uint8((x+y)/2)

        }
    }
    return image
}

func main() {
	pic.Show(Pic)
}
