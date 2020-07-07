package book

import (
	"GoStudy/WebFrameworks/Echo/api/helper/file"
	"GoStudy/WebFrameworks/Echo/api/model"

	"github.com/pquerna/ffjson/ffjson"
)

func getAllBooksDAO() ([]model.Book, error){
	books := [] model.Book{}
	bookDBFilepath := model.GetBooksDBFilePath()

	byteArray, err := file.ReadFile(bookDBFilepath)
	if err != nil {
		return books, err
	}

	err = ffjson.Unmarshal(byteArray, &books)
	if err != nil {
		return books, err
	}

	return books, nil
}

func createBookDAO(bookObj model.Book) (error){
	books := [] model.Book{}
	bookDBFilepath := model.GetBooksDBFilePath()

	byteArray, err := file.ReadFile(bookDBFilepath)
	if err != nil {
		return err
	}

	err = ffjson.Unmarshal(byteArray, &books)
	if err != nil {
		return  err
	}

	books= append(books, bookObj)

	byteArray,err = ffjson.Marshal(&books)
	if err != nil {
		return  err
	}
	err= file.WriteFile(bookDBFilepath,byteArray)
	if err != nil {
		return  err
	}
	return  nil
}

func updateBookDAO(books []model.Book) error{
	bookDBFilepath := model.GetBooksDBFilePath()

	byteArray,err := ffjson.Marshal( &books)
	if err != nil {
		return  err
	}
	err= file.WriteFile(bookDBFilepath,byteArray)
	if err != nil {
		return  err
	}
	return  nil
}