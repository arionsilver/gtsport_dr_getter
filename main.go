package main

import (
	"fmt"
	"github.com/arionsilver/gtsport_dr_getter/client"
	"log"
	"os"
	"strconv"
)

func driverLetter(point string) string {
	num, err := strconv.ParseInt(point, 10, 32)
	if err != nil {
		return "?"
	}

	if num < 4000 {
		return "D"
	} else if num < 10000 {
		return "C"
	} else if num < 30000 {
		return "B"
	} else if num < 50000 {
		return "A"
	} else {
		return "A+"
	}
}

func mannerLetter(point string) string {
	num, err := strconv.ParseInt(point, 10, 32)
	if err != nil {
		return "?"
	}

	if num < 10 {
		return "E"
	} else if num < 20 {
		return "D"
	} else if num < 40 {
		return "C"
	} else if num < 65 {
		return "B"
	} else if num < 80 {
		return "A"
	} else {
		return "S"
	}
}

func main() {
	var userNo int64
	var err error
	if len(os.Args) > 1 {
		userNo, err = strconv.ParseInt(os.Args[1], 10, 32)

		if err != nil {
			log.Fatal(err)
		}
	}

	if userNo == 0 {
		return
	}

	name, err := client.GetUserName(int(userNo))
	if err != nil {
		log.Fatal(err)
	}

	profile, err := client.GetUserProfile(int(userNo))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s's DR is %s(%s) and SR is %s(%s)",
		name,
		profile.DriverPoint,
		driverLetter(profile.DriverPoint),
		profile.MannerPoint,
		mannerLetter(profile.MannerPoint))
}
