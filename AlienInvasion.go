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
var aliens []alien                   //alien array
var citytonub = make(map[int]string) //for randomly selecting alien locations
var iterator int
var citytoalien = make(map[string]int) //maps alien to a city

func move() string {

	mv := rand.Intn(3)
	if mv == 0 {
		return "north"
	}
	if mv == 1 {
		return "south"
	}
	if mv == 2 {
		return "east"
	}
	if mv == 3 {
		return "west"
	}
	return "Invalid"
}
func generateAlienOnMap(count int) {
	aliens = make([]alien, count)
	for i := 0; i < count; i++ {
		x := rand.Intn(iterator)
		fmt.Println("x=", x)
		aliens[i] = alien{true, citytonub[x], 0}
		citytoalien[citytonub[x]] = i
		if i == 0 {
			citytoalien[citytonub[x]] = -1
		} //handling zeroth alien differently
		fmt.Println("Alien ", i, " at :", citytonub[x], " cit map:", citytoalien[citytonub[x]])
	}
}
func generateOrGetCity(st string) city {

	gtcity := cities[st]
	fmt.Println("gtcity :" + gtcity.name)
	fmt.Println("st :", st)
	if gtcity.name != st { //new city detected
		gtcity = city{st, nil, false}
		gtcity.neigbours = make(map[string]*city)
		cities[st] = gtcity
		citytonub[iterator] = st
		iterator++
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
		str := strings.Split(line, " ")
		generateOrGetCity(str[0])
		for i := 1; i < len(str); i++ {
			st := strings.Split(str[i], "=")
			st[1] = strings.TrimRight(st[1], "\r\n")
			neig := generateOrGetCity(st[1])
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
func checkEnd() bool {

	for i := 0; i < len(aliens); i++ {
		if aliens[i].active == true {
			return true
		}
	}
	return false
}
func moveAliens() {

	for i := 0; i < len(aliens); i++ {
		if aliens[i].active {
			aliens[i].totalmoves++
			dir := move()
			if cities[aliens[i].location].neigbours[dir] != nil && !cities[aliens[i].location].neigbours[dir].destroyed {
				nam := cities[aliens[i].location].neigbours[dir].name
				if citytoalien[nam] != 0 { //two aliens in the same
					fmt.Println("Citi ", nam, " Destroyed!!")
					cities[aliens[i].location].neigbours[dir].destroyed = true
					aliens[i].active = false
					if citytoalien[nam] == -1 {
						aliens[0].active = false
					} else {
						aliens[citytoalien[nam]].active = false
					}

				}
			}
			if aliens[i].totalmoves == 10000 {
				aliens[i].active = false
			}

		}
	}

}
func moveTillEnd() {

	for checkEnd() != false {
		moveAliens()
	}
}
func main() {
	//reader := bufio.NewReader(os.Stdin)
	//fmt.Println("Enter the file  : ")
	//filename, _ := reader.ReadString('\n')
	//filename = strings.Replace(filename, "\n", "", -1)
	generateCityMap("file.txt")
	generateAlienOnMap(2)
	//fmt.Println("value : ", cities["v"].neigbours["east"])
	moveTillEnd()
}
