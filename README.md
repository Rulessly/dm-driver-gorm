# dm-driver-gorm
dm-driver-gorm

```go
package main

import (
	"fmt"
	"github.com/Rulessly/dm-driver-gorm"
	"gorm.io/gorm"
)

func main() {
	
	dsn := "dm://SYSDBA:PASSWORD@127.0.0.1:5236?ignoreCase=false&appName=wisdom&statEnable=false"
	db, err := gorm.Open(dm.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
	}

	type User struct {
		ID       string `gorm:"id"`
		USERNAME string `gorm:"username"`
		PASSWORD string `gorm:"password"`
	}

	var user User

	db.Raw("SELECT ID, USERNAME, PASSWORD FROM PERSON.PERSON WHERE ID = ?", "6").Scan(&user)

	fmt.Printf("%+v", user)
}

```
