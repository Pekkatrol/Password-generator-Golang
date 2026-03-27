package main

import (
	"fmt"
	"os"
	"strconv"
	"math/rand"
)

type setting struct {
	width   int
	letters bool
	digits  bool
	symbols bool
}

func print_error(text string) {
	fmt.Fprint(os.Stderr, text)
	os.Exit(2)
}

func check_settings(settings setting) {
	if settings.digits == false && settings.letters == false && settings.symbols == false {
		print_error("You need at least of [letters | digits | symbols]\n")
	}
	if settings.width < 1 {
		print_error("Invalid password width\n")
	}
}

func get_setting(argv []string) (settings setting) {
	settings = setting{10, false, false, false}
	const nb_arg int = 4
	var ok error

	settings.width, ok = strconv.Atoi(argv[1])
	if ok != nil || settings.width < 1 {
		print_error("Invalid password width\n")
	}
	settings.letters, ok = strconv.ParseBool(argv[2])
	if ok != nil {
		print_error("Invalid parameter for the letters\n")
	}
	settings.digits, ok = strconv.ParseBool(argv[3])
	if ok != nil {
		print_error("Invalid parameter for the digits\n")
	}
	settings.symbols, ok = strconv.ParseBool(argv[4])
	if ok != nil {
		print_error("Invalid parameter for the symbols\n")
	}
	return
}

func random_character(settings setting) (character string) {
	var pool string
	var index int

	if (settings.letters == true) {
	    pool += "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	}
	if (settings.digits == true) {
	    pool += "0123456789"
    }
	if (settings.symbols == true) {
		pool += "@#!?%$€&-_=+-"
	}
    index = rand.Intn(len(pool))
	character = string(pool[index])
	return
}

func generate_password(settings setting) string {
	password := make([]byte, settings.width)

	for index := 0; index < settings.width; index++ {
		c := random_character(settings)
		password[index] = c[0]
	}
	return string(password)
}
