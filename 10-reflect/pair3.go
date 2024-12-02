package main

type Reader interface {
	ReadBook()
}

type Writer interface {
	WriteBook()
}

type Book struct {
}

func (b *Book) ReadBook() {
	println("Reading book")
}

func (b *Book) WriteBook() {
	println("Writing book")
}

func main() {
	b := &Book{}

	var r Reader = b

	r.ReadBook()

	var w Writer
	w = r.(Writer)

	w.WriteBook()
}
