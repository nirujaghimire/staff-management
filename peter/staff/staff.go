package staff

type Staff struct {
	Username  string  `json:"username"`
	Password  string  `json:"password"`
	FirstName string  `json:"firstName"`
	LastName  string  `json:"lastName"`
	Email     string  `json:"email"`
	Phone     string  `json:"phone"`
	Address   string  `json:"address"`
	Position  string  `json:"position"`
	Salary    float32 `json:"salary"`
}

var staffDb = map[string]Staff{}

func Add(staff Staff) {
	staffDb[staff.Username] = staff
}

func Get(username string) Staff {
	return staffDb[username]
}

func Edit(username string, staff Staff) {
	staffDb[username] = staff
}

func Delete(username string) {
	delete(staffDb, username)
}

func GetAll() []Staff {
	var values []Staff
	for _, value := range staffDb {
		values = append(values, value)
	}
	return values
}
