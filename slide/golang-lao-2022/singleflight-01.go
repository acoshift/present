package main

import (
	"fmt"
	"strconv"
	"time"

	"golang.org/x/sync/singleflight"
)

// START 1 OMIT
func main() {
	go getUserCompanyID(1)
	go getUserCompanyID(1)
	go getUserCompanyID(1)
	go getUserCompanyID(2)
	go getUserCompanyID(3)
	go getUserCompanyID(1)
	time.Sleep(time.Second)
	fmt.Println("sleep")
	go getUserCompanyID(1)
	time.Sleep(time.Second)
}

// END 1 OMIT

// START 2 OMIT
var userCompanySingleflight = &singleflight.Group{}

func getUserCompanyID(userID int64) (int64, error) {
	key := strconv.FormatInt(userID, 10)
	v, err, _ := userCompanySingleflight.Do(key, func() (interface{}, error) {
		return getUserCompanyIDFromDB(userID)
	})
	return v.(int64), err
}

func getUserCompanyIDFromDB(userID int64) (int64, error) {
	fmt.Printf("db: get user=%d\n", userID)
	time.Sleep(20 * time.Millisecond)
	return map[int64]int64{
		1: 1,
		2: 1,
		3: 1,
		4: 2,
		5: 2,
	}[userID], nil
}

// END 2 OMIT
