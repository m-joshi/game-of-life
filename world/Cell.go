package world

type Cell interface {
	Resurrect()
	Kill()
	GetState() bool
	NextGen()
}

type cell struct {
	RowIndex    int
	ColumnIndex int
	State       bool
}

func NewCell(rowIndex, columnIndex int) Cell {
	return &cell{
		RowIndex:    rowIndex,
		ColumnIndex: columnIndex,
		State:       false,
	}
}

func (c *cell) Resurrect() {
	c.State = true
}

func (c *cell) Kill() {
	c.State = false
}

func (c *cell) GetState() bool {
	return c.State
}

func (c *cell) NextGen() {
	row := c.RowIndex
	column := c.ColumnIndex

	isAlive := World[row][column].GetState()
	aliveNeighborsCount := c.countAliveNeighbors()

	if isAlive && aliveNeighborsCount < 2 {
		World[row][column].Kill()
	}

	if isAlive && aliveNeighborsCount > 3 {
		World[row][column].Kill()
	}

	if !isAlive && aliveNeighborsCount == 3 {
		World[row][column].Resurrect()
	}

}

func (c *cell)countAliveNeighbors() int {
	count := 0
	row := c.RowIndex
	column := c.ColumnIndex

	if World[row-1][column-1].GetState() {
		count++
	}
	if World[row-1][column].GetState() {
		count++
	}
	if World[row-1][column+1].GetState() {
		count++
	}
	if World[row][column-1].GetState() {
		count++
	}
	if World[row][column+1].GetState() {
		count++
	}
	if World[row+1][column-1].GetState() {
		count++
	}
	if World[row+1][column].GetState() {
		count++
	}
	if World[row+1][column+1].GetState() {
		count++
	}

	return count
}