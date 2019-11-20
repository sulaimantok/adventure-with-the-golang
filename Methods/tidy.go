package main


//List package import
import(
	"fmt"
)

const(
	ConstExample = "Paten"
)

var (
	ExportedVar = 42
	nonExportedVar = "so say we all"
)

type User struct{
	FirstName, LastName string
	Location	*UserLocation
}

type UserLocation struct{
	City	string
	Country	string
}

func NewUser(firstname, lastname string) *User{
	return &User{FirstName : firstname,
			LastName : lastname,
			Location : &UserLocation{
				City: "Surabaya",
				Country : "Indonesia",
			},
		}
}

func (u *User) Greeting() string{
	return fmt.Sprintf("Hallo %s %s", u.FirstName,u.LastName)
}

func main(){
	us :=User { FirstName : "Ahmad",
	LastName : "Kamal",
	Location : &UserLocation{
		City: "Surabaya",
		Country : "Indonesia",
	}
	}
	fmt.Println(us.Greeting())
}