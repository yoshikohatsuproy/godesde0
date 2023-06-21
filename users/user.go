package users

import (
	"fmt"
	"time"

	"github.com/yoshio/godesde0/modelos"
)

func AltaUsuario() {
	u := new(modelos.User)
	u.AddUser(10, "Jonathan", time.Now(), true)
	fmt.Println(u)
}
