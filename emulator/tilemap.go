package emulator

// TileMap is the structure that holds the tilemap data for the game
type TileMap [512][8][8]int

func blankTileMap() TileMap {
	tm := TileMap{}
	for i := 0; i < 512; i++ {
		for j := 0; j < 8; j++ {
			for k := 0; k < 8; k++ {
				tm[i][j][k] = 0
			}
		}
	}
	return tm
}
