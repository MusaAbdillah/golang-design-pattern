package main

type Visitor interface {
	visitForCircle(*Circle)
	visitForSquare(*Square)
	visitForRectangle(*Rectangle)
}
