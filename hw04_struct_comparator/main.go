package main

import (
	"fmt"
)

type Event float64

const (
	Year Event = iota
	Size
	Rate
)

type Book struct {
	id     string
	title  string
	author string
	year   int
	size   int
	rate   float64
}

func NewBook(event Event, id, title, author string, year, size int, rate float64) *Book {
	return &Book{
		id:     id,
		title:  title,
		author: author,
		year:   year,
		size:   size,
		rate:   rate,
	}
}

func (b *Book) IsComparison(other *Book, event Event) bool {
	switch event {
	case Year:
		return b.year > other.year
	case Size:
		return b.size > other.size
	case Rate:
		return b.rate > other.rate
	}
	return false

}

func main() {
	book := NewBook(Year, "0", "toxic", "DMA", 1999, 1200, 4.6)
	book1 := NewBook(Size, "1", "toxic", "DMA", 2000, 1000, 4.0)

	if book.IsComparison(book1, Year) {
		fmt.Println("book is bigger than book1")
	} else {
		fmt.Println("book is smaller than book1")
	}

	if book.IsComparison(book1, Size) {
		fmt.Println("book is bigger than book1")
	} else {
		fmt.Println("book is smaller than book1")
	}

	if book.IsComparison(book1, Rate) {
		fmt.Println("book is bigger than book1")
	} else {
		fmt.Println("book is smaller than book1")
	}
}
