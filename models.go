package main

type author struct {
	id        int    `json:"id"`
	FirstName string `json:"FirstName" binding:"required" binding:"required"`
	LastName  string `json:"LastName" binding:"required" binding:"required"`
}

type publish struct {
	id          int    `json:"id"`
	NamePublish string `json:"NamePublish"`
	Address     string `json:"Address"`
	Site        string `json:"Site"`
}

type book struct {
	id          int    `json:"id"`
	Title       string `json:"Title" binding:"required"`
	Code        string `json:"Code" binding:"required"`
	YearPublish int    `json:"YearPublish" binding:"required"`
	CountPage   int    `json:"CountPage" binding:"required"`
	Price       int    `json:"Amount" binding:"required"`
	Hardcover   string `json:"Hardcover" binding:"required"`
	Abstract    string `json:"Abstract" binding:"required"`
	Status      string `json:"Status" binding:"required"`
	AuthorId    int    `json:"AuthorID"`
	PublishId   int    `json:"PublishId"`
	GenreId     int    `json:"GenreId"`
	SubjectId   int    `json:"SubjectId"`
}

type genre struct {
	id   int    `json:"id"`
	Name string `json:"Name" binding:"required"`
}

type subject struct {
	id   int    `json:"id"`
	Name string `json:"Name" binding:"required"`
}

type SP_InfoBook struct {
	FirstName   string `json:"FirstName"`
	LastName    string `json:"LastName"`
	Title       string `json:"Title"`
	PublishYear int    `json:"PublishYear"`
	Publish     string `json:"Publish"`
}
