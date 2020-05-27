package repositories

import (
	"database/sql"
	"log"

	"github.com/zainokta/bookstore_fiber/models"
)

type BookRepo struct {
	db *sql.DB
}

func NewBookRepo(db *sql.DB) *BookRepo {
	return &BookRepo{
		db: db,
	}
}

func (b *BookRepo) All() ([]models.Book, error) {
	rows, err := b.db.Query("SELECT * FROM books")
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer rows.Close()

	result := []models.Book{}

	for rows.Next() {
		book := models.Book{}
		err := rows.Scan(&book.ID, &book.Title, &book.Author, &book.Price, &book.CreatedAt, &book.UpdatedAt)
		if err != nil {
			log.Println(err)
			return nil, err
		}

		result = append(result, book)
	}

	return result, nil
}

func (b *BookRepo) Find(id int64) (models.Book, error) {
	rows, err := b.db.Query("SELECT * FROM books WHERE id=$1", id)
	if err != nil {
		log.Println(err)
		return models.Book{}, err
	}
	defer rows.Close()

	book := models.Book{}
	var hasResult bool = false
	for rows.Next() {
		hasResult = true
		err = rows.Scan(&book.ID, &book.Title, &book.Author, &book.Price, &book.CreatedAt, &book.UpdatedAt)
		if err != nil {
			log.Println(err)
			return models.Book{}, err
		}
	}
	if !hasResult {
		return models.Book{}, nil
	}

	return book, nil
}

func (b *BookRepo) Create(book *models.Book) error {
	stmt, err := b.db.Prepare("INSERT INTO books(title,author,price) VALUES($1,$2,$3)")
	if err != nil {
		log.Println(err)
		return err
	}

	_, err = stmt.Exec(&book.Title, &book.Author, &book.Price)
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}

func (b *BookRepo) Update(id int64, book *models.Book) error {
	stmt, err := b.db.Prepare("UPDATE books SET title=$1, author=$2, price=$3 WHERE id=$4")
	if err != nil {
		log.Println(err)
	}

	_, err = stmt.Exec(&book.Title, &book.Author, &book.Price, id)
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}

func (b *BookRepo) Delete(id int64) error {
	stmt, err := b.db.Prepare("DELETE FROM books WHERE id=$1")
	if err != nil {
		log.Println(err)
	}

	_, err = stmt.Exec(id)
	if err != nil {
		log.Println(err)
	}

	return nil
}
