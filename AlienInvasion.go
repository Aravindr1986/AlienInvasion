package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strings"
)

type city struct {
	name      string
	neigbours map[string]*city
	destroyed bool
}
type alien struct {
	active     bool
	location   string
	totalmoves int
}

var cities map[string]city
var aliens []alien

func move() string {

	mv := rand.Int() % 10
	if mv == 1 {
		return "north"
	}
	if mv == 2 {
		return "south"
	}
	if mv == 3 {
		return "east"
	}
	if mv == 4 {
		return "west"
	}
	return "Invalid"
}
func generateOrGetCity(st string) city {

	gtcity := cities[st]
	fmt.Println("gtcity :" + gtcity.name)
	fmt.Println("st :", st)
	if gtcity.name != st {
		gtcity = city{st, nil, false}
		gtcity.neigbours = make(map[string]*city)
		cities[st] = gtcity
	}
	return (gtcity)
}
func generateCityMap(fileName string) {
	file, err := os.Open(fileName)

	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	var line string
	cities = make(map[string]city)
	for {

		line, err = reader.ReadString('\n')
		line = strings.Replace(line, "\n", "", -1)
		fmt.Println("line :", line)
		str := strings.Split(line, " ")
		generateOrGetCity(str[0])
		for i := 1; i < len(str); i++ {
			st := strings.Split(str[i], "=")
			st[1] = strings.TrimRight(st[1], "\r\n")
			fmt.Print("st[1]:", st[1])
			fmt.Print("lenght:", len(st[1]))
			neig := generateOrGetCity(st[1])
			fmt.Print("neig name :", neig.name)
			cities[str[0]].neigbours[st[0]] = &neig
		}

		if err != nil {
			break
		}
	}
	fmt.Println("The last printer!!")
	for k, v := range cities {
		fmt.Println("key: ", k, " value: ", v)
		//fmt.Println(v.name)
	}
}
func main() {
	//reader := bufio.NewReader(os.Stdin)
	//fmt.Println("Enter the file  : ")
	//filename, _ := reader.ReadString('\n')
	//filename = strings.Replace(filename, "\n", "", -1)
	//fmt.Print("reading ", filename)
	//fmt.Print("hello")
	generateCityMap("file.txt")

}
