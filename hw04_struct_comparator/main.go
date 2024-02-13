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
	compareType Event
	id          string
	title       string
	author      string
	year        int
	size        int
	rate        float64
}

func NewBook(event Event, id, title, author string, year, size int, rate float64) *Book {
	return &Book{
		id:          id,
		title:       title,
		author:      author,
		year:        year,
		size:        size,
		rate:        rate,
		compareType: event,
	}
}

func (b *Book) IsComparison(other *Book) int {
	switch b.compareType {
	case Year:
		return compareYears(b, other)
	case Size:
		return compareSizes(b, other)
	case Rate:
		return compareRates(b, other)
	}
	return 0
}

func compareYears(b, other *Book) int {
	if b.year == other.year {
		return 0
	} else if b.year > other.year {
		return 1
	}
	return -1
}

func compareSizes(b, other *Book) int {
	if b.size == other.size {
		return 0
	} else if b.size > other.size {
		return 1
	}
	return -1
}

func compareRates(b, other *Book) int {
	if b.rate == other.rate {
		return 0
	} else if b.rate > other.rate {
		return 1
	}
	return -1
}

func main() {
	book := NewBook(Year, "0", "toxic", "DMA", 2000, 1200, 4.1)
	book1 := NewBook(Size, "1", "toxic", "DMA", 2000, 1200, 4.6)

	switch book.IsComparison(book1) {
	case 0:
		fmt.Println("Book equality")
	case 1:
		fmt.Println("book greater book1")
	case -1:
		fmt.Println("book smaller book1")
	}
}
