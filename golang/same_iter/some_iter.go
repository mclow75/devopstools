package main

/*
Итераторы появились в Go v1.23 и нужны для разделения логики работы
с последовательностью(массивом) и логикой (циклом) обработки.
https://youtu.be/wL2c-OemZeA
*/
import (
	"fmt"
	"iter"
)

type Student struct {
	LastName    string
	FirstName   string
	MiddleName  string
	YearOfStudy int
}

var students = []Student{
	Student{LastName: "Иванов", FirstName: "Альберт", MiddleName: "Петрович", YearOfStudy: 3},
	Student{LastName: "Смирнов", FirstName: "Денис", MiddleName: "Олегович", YearOfStudy: 2},
	Student{LastName: "Татаринова", FirstName: "Зульфия", MiddleName: "Агамнова", YearOfStudy: 3},
	Student{LastName: "Волкова", FirstName: "Елена", MiddleName: "Петровна", YearOfStudy: 3},
	Student{LastName: "Морозова", FirstName: "Виктория", MiddleName: "Николаевна", YearOfStudy: 1},
	Student{LastName: "Фёдорова", FirstName: "Полина", MiddleName: "Максимовна", YearOfStudy: 1},
}

func main() {
	it iter.Seq[Student] := slices.Values(s: students)
	for student Student := range it {
		fmt.Println(student)
	}
}
