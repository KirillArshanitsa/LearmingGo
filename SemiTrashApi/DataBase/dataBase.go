package DataBase

import "log"

type Book struct {
	Id            uint64 `json:"id"`
	Name          string `json:"name"`
	Author        Author `json:"author"`
	DateOfWriting string `json:"dateOfWriting"`
}

type Author struct {
	Name     string `json:"name"`
	Birthday string `json:"birthday"`
}

var (
	AllBooks []Book
	bookId   uint64 = 1
)

func init() {
	log.Println("Create data of books")
	Book1 := Book{
		Id:            bookId,
		Name:          "Book1",
		Author:        Author{Name: "Author1", Birthday: "26/06/1986"},
		DateOfWriting: "11/01/2009",
	}
	bookId++
	Book2 := Book{
		Id:            bookId,
		Name:          "Book2",
		Author:        Author{Name: "Author2", Birthday: "22/06/1996"},
		DateOfWriting: "11/01/2009",
	}

	AllBooks = append(AllBooks, Book1, Book2)
}

func GetBookById(bookId uint64) (*Book, bool) {
	for _, book := range AllBooks {
		if book.Id == bookId {
			return &book, true
		}
	}
	return &Book{}, false
}

func AddBook(book *Book) {
	bookId++
	book.Id = bookId
	AllBooks = append(AllBooks, *book)
}

func UpdateBook(book *Book) {
	for i := 0; i < len(AllBooks); i++ {
		if AllBooks[i].Id == book.Id {
			AllBooks[i] = *book
		}
	}
}

func DeleteBook(bookId uint64) bool {
	for i := 0; i < len(AllBooks); i++ {
		if AllBooks[i].Id == bookId {
			AllBooks = append(AllBooks[:i], AllBooks[i+1:]...)
			return true
		}
	}
	return false
}
