package main

import "fmt"
import "errors"

func checkUserName(userName string) error {
	if userName == "alierbey" {
		return errors.New("bu kullanıcı adı  zaten alınmıs")
	}
	fmt.Println(userName)
	return nil
}

func main() {
	userNames := []string{"ahmet", "alierbey", "ayşe"}
	for _, i := range userNames {
		if err := checkUserName(i); err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(i, "kullanıcı adı uygun")
		}
	}
}
