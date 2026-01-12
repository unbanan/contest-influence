package protobuf

import (
	simulation "contest-influence/proto/simulation"
	"math/rand"
)

type MapFactory struct{}

func (f *MapFactory) CreateSmallMap() *simulation.Map {
	return &simulation.Map{
		Nrows: 5,
		Ncols: 5,
		FieldMask: []bool{
			true,  true,  true,  true,  true,
			true,  true,  false, true,  true,
			true,  false, true,  false, true,
			true,  true,  false, true,  true,
			true,  true,  true,  true,  true,
		},
		BigCells: []*simulation.Position{
			{Row: 2, Col: 2},
		},
		StartPositions: []*simulation.Position{
			{Row: 0, Col: 0},
			{Row: 4, Col: 4},
		},
	}
}

func (f *MapFactory) CreateMediumMap() *simulation.Map {
	fieldMask := make([]bool, 64) 
	for i := range fieldMask {
		fieldMask[i] = (i%5 != 0)
	}
	
	return &simulation.Map{
		Nrows: 8,
		Ncols: 8,
		FieldMask: fieldMask,
		BigCells: []*simulation.Position{
			{Row: 2, Col: 2},
			{Row: 2, Col: 5},
			{Row: 5, Col: 2},
			{Row: 5, Col: 5},
		},
		StartPositions: []*simulation.Position{
			{Row: 0, Col: 0},
			{Row: 0, Col: 7},
			{Row: 7, Col: 0},
			{Row: 7, Col: 7},
		},
	}
}

func (f *MapFactory) CreateLargeMap() *simulation.Map {
	fieldMask := make([]bool, 144)
	
	for i := 0; i < 12; i++ {
		for j := 0; j < 12; j++ {
			idx := i*12 + j
			fieldMask[idx] = (i == 0 || i == 11 || j == 0 || j == 11 || 
				i == j || i+j == 11 ||
				(i >= 4 && i <= 7 && j >= 4 && j <= 7))
		}
	}
	
	return &simulation.Map{
		Nrows: 12,
		Ncols: 12,
		FieldMask: fieldMask,
		BigCells: []*simulation.Position{
			{Row: 3, Col: 3},
			{Row: 3, Col: 8},
			{Row: 8, Col: 3},
			{Row: 8, Col: 8},
			{Row: 5, Col: 5},
			{Row: 5, Col: 6},
			{Row: 6, Col: 5},
			{Row: 6, Col: 6},
		},
		StartPositions: []*simulation.Position{
			{Row: 0, Col: 0},
			{Row: 0, Col: 11},
			{Row: 11, Col: 0},
			{Row: 11, Col: 11},
			{Row: 0, Col: 5},
			{Row: 5, Col: 0},
			{Row: 11, Col: 5},
			{Row: 5, Col: 11},
		},
	}
}

func (f *MapFactory) CreateIslandMap() *simulation.Map {
	fieldMask := make([]bool, 100)
	
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			idx := i*10 + j
			fieldMask[idx] = (i >= 2 && i <= 7 && j >= 2 && j <= 7)
		}
	}
	
	return &simulation.Map{
		Nrows: 10,
		Ncols: 10,
		FieldMask: fieldMask,
		BigCells: []*simulation.Position{
			{Row: 4, Col: 4},
			{Row: 4, Col: 5},
			{Row: 5, Col: 4},
			{Row: 5, Col: 5},
		},
		StartPositions: []*simulation.Position{
			{Row: 2, Col: 2},
			{Row: 2, Col: 7},
			{Row: 7, Col: 2},
			{Row: 7, Col: 7},
		},
	}
}

func (f *MapFactory) CreateMazeMap() *simulation.Map {
	fieldMask := []bool{
		true,  true,  true,  false, true,  true,  true,
		true,  false, true,  false, true,  false, true,
		true,  false, true,  true,  true,  false, true,
		false, false, false, true,  false, false, false,
		true,  false, true,  true,  true,  false, true,
		true,  false, true,  false, true,  false, true,
		true,  true,  true,  false, true,  true,  true,
	}
	
	return &simulation.Map{
		Nrows: 7,
		Ncols: 7,
		FieldMask: fieldMask,
		BigCells: []*simulation.Position{
			{Row: 3, Col: 3},
		},
		StartPositions: []*simulation.Position{
			{Row: 0, Col: 0},
			{Row: 6, Col: 6},
		},
	}
}

func (f *MapFactory) CreateRandomMap(rows, cols int32, bigCellsCount, startPositionsCount int) *simulation.Map {
	totalCells := rows * cols
	fieldMask := make([]bool, totalCells)
	
	for i := range fieldMask {
		fieldMask[i] = rand.Intn(100) < 70
	}
	
	fieldMask[0] = true
	fieldMask[int(cols-1)] = true
	fieldMask[int((rows-1)*cols)] = true
	fieldMask[int(totalCells-1)] = true
	
	bigCells := make([]*simulation.Position, 0, bigCellsCount)
	for i := 0; i < bigCellsCount; i++ {
		row := uint32(rand.Intn(int(rows)))
		col := uint32(rand.Intn(int(cols)))
		bigCells = append(bigCells, &simulation.Position{Row: row, Col: col})
	}
	
	startPositions := make([]*simulation.Position, 0, startPositionsCount)
	cornerPositions := []*simulation.Position{
		{Row: 0, Col: 0},
		{Row: 0, Col: uint32(cols - 1)},
		{Row: uint32(rows - 1), Col: 0},
		{Row: uint32(rows - 1), Col: uint32(cols - 1)},
	}
	
	for i := 0; i < startPositionsCount && i < len(cornerPositions); i++ {
		startPositions = append(startPositions, cornerPositions[i])
	}
	
	for i := len(startPositions); i < startPositionsCount; i++ {
		row := uint32(rand.Intn(int(rows)))
		col := uint32(rand.Intn(int(cols)))
		startPositions = append(startPositions, &simulation.Position{Row: row, Col: col})
	}
	
	return &simulation.Map{
		Nrows:          rows,
		Ncols:          cols,
		FieldMask:      fieldMask,
		BigCells:       bigCells,
		StartPositions: startPositions,
	}
}

func (f *MapFactory) GetAllMaps() map[string]*simulation.Map {
	return map[string]*simulation.Map{
		"small":   f.CreateSmallMap(),
		"medium":  f.CreateMediumMap(),
		"large":   f.CreateLargeMap(),
		"island":  f.CreateIslandMap(),
		"maze":    f.CreateMazeMap(),
		"random8": f.CreateRandomMap(8, 8, 4, 4),
	}
}