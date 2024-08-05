package main

import (
	"fmt"
)

type Book struct {
	id     int
	title  string
	author string
	year   int
	size   int
	rate   float64
}

func NewBook(id int, title, author string, year, size int, rate float64) *Book {
	return &Book{
		id:     id,
		title:  title,
		author: author,
		year:   year,
		size:   size,
		rate:   rate,
	}
}

func (b *Book) GetID() int {
	return b.id
}

func (b *Book) GetTitle() string {
	return b.title
}

func (b *Book) GetAuthor() string {
	return b.author
}

func (b *Book) GetYear() int {
	return b.year
}

func (b *Book) GetSize() int {
	return b.size
}

func (b *Book) GetRate() float64 {
	return b.rate
}

func (b *Book) SetID(id int) {
	b.id = id
}

func (b *Book) SetTitle(title string) {
	b.title = title
}

func (b *Book) SetAuthor(author string) {
	b.author = author
}

func (b *Book) SetYear(year int) {
	b.year = year
}

func (b *Book) SetSize(size int) {
	b.size = size
}

func (b *Book) SetRate(rate float64) {
	b.rate = rate
}

type CompareMode int

const (
	ByYear CompareMode = iota
	BySize
	ByRate
)

type BookComparer struct {
	mode CompareMode
}

func NewBookComparer(mode CompareMode) *BookComparer {
	return &BookComparer{mode: mode}
}

func (bc *BookComparer) Compare(b1, b2 *Book) (bool, error) {
	if b1 == nil || b2 == nil {
		return false, fmt.Errorf("one or both books are nil")
	}

	switch bc.mode {
	case ByYear:
		return b1.year > b2.year, nil
	case BySize:
		return b1.size > b2.size, nil
	case ByRate:
		return b1.rate > b2.rate, nil
	default:
		return false, fmt.Errorf("invalid compare mode")
	}
}

func main() {
	book1 := NewBook(1, "Book One", "Author A", 2020, 300, 4.5)
	book2 := NewBook(2, "Book Two", "Author B", 2021, 250, 4.7)

	comparer := NewBookComparer(ByYear)
	result, err := comparer.Compare(book1, book2)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Book1 > Book2 by Year:", result)
	}

	comparer = NewBookComparer(BySize)
	result, err = comparer.Compare(book1, book2)

	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Book1 > Book2 by Size:", result)
	}

	comparer = NewBookComparer(ByRate)
	result, err = comparer.Compare(book1, book2)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Book1 > Book2 by Rate:", result)
	}
}

//hw04-test
