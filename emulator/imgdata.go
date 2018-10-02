package emulator

// ImgData holds the information for the screen
// It is an integer array that is L x W x 4. Each pixel is split into 4 numbers
// representing (Alpha, Red, Green, Blue)
type ImgData struct {
	length int
	width  int
	data   [89600]int // 160 * 144 * 4
}

func gameboySolidColor(color int) ImgData {
	length := 160 // Gameboy screen length
	width := 144  // Gameboy screen width
	data := [89600]int{}

	for i := 0; i < 89600; i++ {
		data[i] = color
	}

	return ImgData{
		length: length,
		width:  width,
		data:   data,
	}
}

var gameboyBLACK = gameboySolidColor(0)
var gameboyWHITE = gameboySolidColor(255)
var gameboyGREY = gameboySolidColor(100)
