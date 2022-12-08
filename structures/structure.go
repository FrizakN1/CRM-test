package structures

type Setting struct {
	Address string
	Port    string
	DbHost  string
	DbPort  string
	DbUser  string
	DbPass  string
	DbName  string
}

type User struct {
	ID       int
	Login    string
	Password string
	Name     string
	Surname  string
	Role     int
}
