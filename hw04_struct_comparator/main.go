package main

import (
	"fmt"
)

type Book struct {
	ID     int
	Title  string
	Author string
	Year   int
	Size   int
	Rate   float64
}

func (b *Book) SetID(id int) {
	b.ID = id
}

func (b *Book) GetID() int {
	return b.ID
}

func (b *Book) SetTitle(title string) {
	b.Title = title
}

func (b *Book) GetTitle() string {
	return b.Title
}

func (b *Book) SetAuthor(author string) {
	b.Author = author
}

func (b *Book) GetAuthor() string {
	return b.Author
}

func (b *Book) SetYear(year int) {
	b.Year = year
}

func (b *Book) GetYear() int {
	return b.Year
}

func (b *Book) SetSize(size int) {
	b.Size = size
}

func (b *Book) GetSize() int {
	return b.Size
}

func (b *Book) SetRate(rate float64) {
	b.Rate = rate
}

func (b *Book) GetRate() float64 {
	return b.Rate
}

type BookComparatorMode int

const (
	Year BookComparatorMode = iota
	Size
	Rate
)

type BookComparator struct {
	mode BookComparatorMode
}

func NewBookComparator(mode BookComparatorMode) BookComparator {
	return BookComparator{mode}
}

func (c *BookComparator) Compare(book1, book2 Book) int {
	switch c.mode {
	case Year:
		if book1.Year < book2.Year {
			return -1
		} else if book1.Year > book2.Year {
			return 1
		}
	case Size:
		if book1.Size < book2.Size {
			return -1
		} else if book1.Size > book2.Size {
			return 1
		}
	case Rate:
		if book1.Rate < book2.Rate {
			return -1
		} else if book1.Rate > book2.Rate {
			return 1
		}
	}
	return 0
}

func main() {
	book1 := Book{ID: 1, Title: "Book 1", Author: "Author 1", Year: 2020, Size: 200, Rate: 4.5}
	book2 := Book{ID: 2, Title: "Book 2", Author: "Author 2", Year: 2020, Size: 200, Rate: 4.5}

	comparatorByYear := NewBookComparator(Year)
	comparatorBySize := NewBookComparator(Size)
	comparatorByRate := NewBookComparator(Rate)

	switch comparatorByYear.Compare(book1, book2) {
	case -1:
		fmt.Println("Book 1 is older than Book 2")
	case 1:
		fmt.Println("Book 2 is older than Book 1")
	case 0:
		fmt.Println("Book 1 and Book 2 are identical")
	}
	switch comparatorBySize.Compare(book1, book2) {
	case -1:
		fmt.Println("Book 1 is smaller than Book 2")
	case 1:
		fmt.Println("Book 2 is smaller than Book 1")
	case 0:
		fmt.Println("Book 1 and Book 2 are identical")
	}

	switch comparatorByRate.Compare(book1, book2) {
	case -1:
		fmt.Println("Book 1 is cheaper than Book 2")
	case 1:
		fmt.Println("Book 2 is cheaper than Book 1")
	case 0:
		fmt.Println("Book 1 and Book 2 are identical")
	}
}
