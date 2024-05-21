package main
import "fmt"
type Person struct {
	name string
	sex  string
	weight string
}

func main () {
	person := Person{"fake", "male", "140"}
	if person.name == "fake" {

		fmt.Println(person.sex, person.getName())
	}
}

func (l *Person) getName() string {
	return l.name
}