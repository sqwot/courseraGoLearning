package person

var (
	Public  = 1
	private = 1
)

type Person struct {
	Id     int
	Name   string
	secret string
}

func (p Person) UpdateSecret(secret string) {
	p.secret = secret
}
