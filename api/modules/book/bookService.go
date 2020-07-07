package book

import (
	"GoStudy/WebFrameworks/Echo/api/model"
	"errors"

	uuid "github.com/satori/go.uuid"
)

func readBookService(id string) (model.Book, error){
	booksData, err := getAllBooksDAO()
	if err != nil {
		return model.Book{}, errors.New("DB_READ_ERROR")
	}
	for _, bookObj := range booksData {
		if bookObj.ID == id{
			return bookObj, nil
		}
	}
	return model.Book{}, errors.New("NOT_FOUND")
}



func createBookService(bookObj model.Book)(string, error){
	bookObj.ID = uuid.Must(uuid.NewV4()).String()

	err := createBookDAO(bookObj)
	if err != nil {
		return "", errors.New("DB_SAVE_ERROR")
	}

	return bookObj.ID, nil
}

func updateBookService(bookObj model.Book)(string, error){
	booksData, err := getAllBooksDAO()
	if err != nil {
		return "", errors.New("DB_READ_ERROR")
	}
	for idx, dbObj := range booksData {
		if bookObj.ID == dbObj.ID{
			bookObj.ID = dbObj.ID
			booksData[idx] = bookObj
			err = updateBookDAO(booksData)
			if err != nil {
				return "", errors.New("DB_SAVE_ERROR")
			}
			return bookObj.ID, nil
		}
	}

	return bookObj.ID, errors.New("NOT_FOUND")
}


func deleteBookService(id string) error{
	booksData, err := getAllBooksDAO()
	if err != nil {
		return  errors.New("DB_READ_ERROR")
	}
	for idx, bookObj := range booksData {
		if bookObj.ID == id{
			booksData = append(booksData[:idx], booksData[idx+1:]...)
			err = updateBookDAO(booksData)
			if err != nil {
				return  errors.New("DB_SAVE_ERROR")
			}
			return nil
		}
	}
	return errors.New("NOT_FOUND")
}