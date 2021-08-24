package main

type Student struct{
	fname, lname string
}

func (s Student) Fullname() string {
	return fmt.Sprintf("%s %s", s.fname, s.lname)
}


func main() {
	var student []Student = []Student{
		Student{"Aziz", "Bahodirov"},
		
	}
}