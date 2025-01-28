package lib

type Point struct{ X, Y int }

func (p Point) Add(q Point) Point { return Point{p.X + q.X, p.Y + q.Y} }

func (p Point) Sub(q Point) Point { return Point{p.X - q.X, p.Y - q.Y} }

func (p Point) Scale(c int) Point { return Point{p.X * c, p.Y * c} }
