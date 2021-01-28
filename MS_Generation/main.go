package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"strings"
	"time"

	pb "github.com/iavealokin/microservices/MS_Generation/user"
	"google.golang.org/grpc"
)

//User struct ...
type User struct {
	Login    string `json:"login"`
	Username string `json:"username"`
	Surname  string `json:"surname"`
	Birthday string `json:"birthday"`
	Password string `json:"password"`
}

var (
	lowerCharSet   = "abcdedfghijklmnopqrst"
	upperCharSet   = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	specialCharSet = "!@#$%&*"
	numberSet      = "0123456789"
	allCharSet     = lowerCharSet + upperCharSet + specialCharSet + numberSet
)

func main() {
	NewUser := getUser()
	outUser, err := json.Marshal(NewUser)
	if err != nil {
		fmt.Println(err)
		return
	}
	sendUser(outUser)
}

func getUser() User {
	var usernames = map[int]string{
		0:  "John",
		1:  "David",
		2:  "Steve",
		3:  "Peter",
		4:  "Elena",
		5:  "Caroline",
		6:  "Vanda",
		7:  "William",
		8:  "Harper",
		9:  "Evelyn",
		10: "Jackson",
		11: "Jack",
		12: "James",
		13: "Mason",
		14: "Ella",
		15: "Avery",
	}
	var surnames = map[int]string{
		0:  "Davidson",
		1:  "Maksimoff",
		2:  "Rodgers",
		3:  "Stark",
		4:  "Quill",
		5:  "Doe",
		6:  "Jobs",
		7:  "Aaron",
		8:  "Abbot",
		9:  "Lang",
		10: "Longbottom",
		11: "Potter",
		12: "Snape",
		13: "Adams",
		14: "Willson",
		15: "Daniels",
	}
	rand.Seed(time.Now().UnixNano())
	name := usernames[rand.Intn(16)]
	surname := surnames[rand.Intn(16)]
	birthday := randate().Format("2006-01-02")
	login := surname[0:rand.Intn(len(surname))] + name[0:rand.Intn(len(name))]
	if len(login) < 3 {
		login = login + surname[0:rand.Intn(len(surname))] + name[0:rand.Intn(len(name))]
	}
	rand.Seed(time.Now().Unix())
	minSpecialChar := 1
	minNum := 1
	minUpperCase := 1
	passwordLength := 8
	password := generatePassword(passwordLength, minSpecialChar, minNum, minUpperCase)

	minSpecialChar = 2
	minNum = 2
	minUpperCase = 2
	passwordLength = 20
	password = generatePassword(passwordLength, minSpecialChar, minNum, minUpperCase)
	user := User{
		login,
		name,
		surname,
		birthday,
		password,
	}
	return user
}

func randate() time.Time {
	min := time.Date(1940, 1, 0, 0, 0, 0, 0, time.UTC).Unix()
	max := time.Date(2018, 1, 0, 0, 0, 0, 0, time.UTC).Unix()
	delta := max - min

	sec := rand.Int63n(delta) + min
	return time.Unix(sec, 0)
}

func generatePassword(passwordLength, minSpecialChar, minNum, minUpperCase int) string {
	var password strings.Builder

	//Set special character
	for i := 0; i < minSpecialChar; i++ {
		random := rand.Intn(len(specialCharSet))
		password.WriteString(string(specialCharSet[random]))
	}

	//Set numeric
	for i := 0; i < minNum; i++ {
		random := rand.Intn(len(numberSet))
		password.WriteString(string(numberSet[random]))
	}

	//Set uppercase
	for i := 0; i < minUpperCase; i++ {
		random := rand.Intn(len(upperCharSet))
		password.WriteString(string(upperCharSet[random]))
	}

	remainingLength := passwordLength - minSpecialChar - minNum - minUpperCase
	for i := 0; i < remainingLength; i++ {
		random := rand.Intn(len(allCharSet))
		password.WriteString(string(allCharSet[random]))
	}
	inRune := []rune(password.String())
	rand.Shuffle(len(inRune), func(i, j int) {
		inRune[i], inRune[j] = inRune[j], inRune[i]
	})
	return string(inRune)
}

func sendUser(user []byte) {
	conn, err := grpc.Dial("localhost:20100", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	c := pb.NewUserClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	rply, err := c.SendPass(ctx, &pb.MsgRequest{Message: user})
	if err != nil {
		log.Println("something went wrong", err)
	}
	log.Println(rply.Sent)

}
