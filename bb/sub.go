package bb

var Bb = "何か入力してください"

type Person struct {
	Name string
}

func (p *Person) Hello() string {
	return "Hello, " + p.Name
}
