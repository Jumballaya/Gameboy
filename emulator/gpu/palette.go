package gpu

// Palette holds the GPUs palette data
type Palette struct {
	background []int
	obj0       []int
	obj1       []int
}

var blankPalette = Palette{
	background: []int{255, 255, 255, 255},
	obj0:       []int{255, 255, 255, 255},
	obj1:       []int{255, 255, 255, 255},
}
