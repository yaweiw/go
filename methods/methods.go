package methods

// Coordinate is (X,Y)
type Coordinate struct {
	X string
	Y string
}

// ICoordinate is interface of Coordinate
type ICoordinate interface {
	ToString() string
}

//ToString returns string representation of Coordinate
func (c Coordinate) ToString() string {
	return c.X + ":" + c.Y
}
