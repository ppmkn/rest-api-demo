package main

type Book struct{
	ID     string  `json:"id"`	   // id книги
	Title  string  `json:"title"`  // название книги
	Author *Author `json:"author"` // указатель на структуру Author
}

type Author struct{
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}