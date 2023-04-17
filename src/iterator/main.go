package iterator

import (
	"errors"
	"fmt"
)

type Iterator[T any] interface {
	HasNext() bool
	Next() (T, error)
}

type Aggregator[T any] interface {
	Iterator() Iterator[T]
}

type Book struct {
	name string
}

func NewBook(name string) Book {
	return Book{name}
}

func (b *Book) Name() string {
	return b.name
}

type BookShelf struct {
	books []*Book
	last  int
}

func NewBookShelf() BookShelf {
	return BookShelf{[]*Book{}, 0}
}

func (bs BookShelf) BookAt(index int) (*Book, error) {
	if len := len(bs.books); index > len {
		errorMsg := fmt.Sprintf("index %d is out of range, must be less than %d", index, len)
		return nil, errors.New(errorMsg)
	}
	return bs.books[index], nil
}

func (bs *BookShelf) AppendBook(book *Book) {
	bs.books = append(bs.books, book)
	bs.last++
}

func (bs *BookShelf) Length() int {
	return bs.last
}

func (bs *BookShelf) Iterator() Iterator[*Book] {
	iterator := NewBookShelfIterator(bs)
	return &iterator
}

type BookShelfIterator struct {
	bookShelf *BookShelf
	index     int
}

func NewBookShelfIterator(bookShelf *BookShelf) BookShelfIterator {
	return BookShelfIterator{bookShelf, 0}
}

func (bsi *BookShelfIterator) HasNext() bool {
	return bsi.index < bsi.bookShelf.Length()
}

func (bsi *BookShelfIterator) Next() (*Book, error) {
	book, err := bsi.bookShelf.BookAt(bsi.index)
	if err != nil {
		return nil, err
	}
	bsi.index++
	return book, nil
}

func main() {
	var ag Aggregator[*Book]
	bs := NewBookShelf()
	ag = &bs

	titles := [...]string{"A", "B", "C", "D"}
	for _, title := range titles {
		book := NewBook(title)
		bs.AppendBook(&book)
	}

	it := ag.Iterator()
	for it.HasNext() {
		book, err := it.Next()
		if err != nil {
			fmt.Printf("err: %v\n", err)
			break
		}
		fmt.Printf("%s\n", book.Name())
	}
}

// codes for main package
type Executer struct{}

func NewExecuter() Executer {
	return Executer{}
}

func (e Executer) Label() string {
	return "iterator pattern"
}

func (e Executer) Do() {
	main()
}
