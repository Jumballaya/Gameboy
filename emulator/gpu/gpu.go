package gpu

// Memory Management Unit struct
type GPU struct {
	vram          [8192]int
	oam           [160]int
	reg           []int
	tilemap       TileMap
	objData       []int
	objDataSorted []int
	palette       Palette
	scanRow       []int

	screen ImgData

	curLine    int
	curScan    int
	lineMode   int
	modeClocks int

	yScroll int
	xScroll int
	raster  int
	ints    int

	lcdON int
	bgON  int
	objON int
	winON int

	objSize int

	bgTileBase  int
	bgMapBase   int
	winTileBase int
}

// New creates a new GPU
func New() *GPU {
	gpu := &GPU{}
	gpu.Reset()
	return gpu
}

// Reset sets all the values back to their starting positions
func (gpu *GPU) Reset() {
	// Set VRAM
	for i := 0; i < 8192; i++ {
		gpu.vram[i] = 0
	}

	// Set OAM
	for i := 0; i < 160; i++ {
		gpu.oam[i] = 0
	}

	gpu.tilemap = blankTileMap()
	gpu.palette = blankPalette

	// @TODO: Set up renderer here
	gpu.screen = gameboyBLACK
}

// ReadByte reads the memory at the given byte address
func (gpu *GPU) ReadByte(addr int) int {
	return gpu.vram[addr%len(gpu.vram)]
}

// WriteByte is the opposite of ReadByte and writes the val at the given address
func (gpu *GPU) WriteByte(addr, val int) {
	gpu.vram[addr%len(gpu.vram)] = val
}
