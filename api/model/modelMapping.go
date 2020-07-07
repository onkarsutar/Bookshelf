package model

var (
	DBROOTPATH = "D:/data/"
)

func GetBooksDBFilePath() string {
	return DBROOTPATH + "books.json"
}