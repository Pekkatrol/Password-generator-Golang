package main

import (
	"fmt"
	"os"
)

func main() {
    argv := os.Args
    var settings setting
	var password string

	if (len(argv) != 5) {
	    os.Exit(2)
	}
    settings = get_setting(argv)
	check_settings(settings)
	password = generate_password(settings)
	fmt.Printf("The generated password is: %s\n", password)
}