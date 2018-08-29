package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"
)

//each city is numbered from 1 to M cites
type city struct {
	id        int
	name      string
	neighbors map[int]*city //north = 1 south =2 East = 3 West = 4
	destroyed bool
}
type alien struct {
	active     bool
	location   int
	totalmoves int
}

var cities map[int]*city            //map of cities
var aliens []alien                  //alien array
var iterator int                    //track of the number of cities
var citytoalien = make(map[int]int) //maps alien to a city

func move() int {

	mv := rand.Intn(4) //detemining the direction of the next move
	if mv == 0 {
		mv++
	}
	return mv
}
func translateDirection(direc string) int { //transalator for random alien directions
	if direc == "north" {
		return 1
	}
	if direc == "south" {
		return 2
	}
	if direc == "east" {
		return 3
	}
	if direc == "west" {
		return 4
	}
	return 0
}

func generateAlienOnMap(count int) { //generating aliens and assigning to random city
	aliens = make([]alien, count+1) //count of alien + 1 to avoid zero conflict
	x := 0
	for i := 1; i <= count; i++ {
		for citytoalien[x] != 0 { //ensuring that two aliens are not assigned to the same citi intially
			x = rand.Intn(iterator) //randomly select the aliens location out of available cities
		}
		if x == 0 {
			x++
		}
		aliens[i] = alien{true, x, 0}
		citytoalien[x] = i //mapping alien to city

		fmt.Println("Alien ", i, " at ", cities[x].name)
	}
}
func generateOrGetCity(cityname string) *city { //generates new city object

	var gtcity city
	iterator++
	gtcity = city{iterator, cityname, nil, false}
	gtcity.neighbors = make(map[int]*city)
	cities[iterator] = &gtcity
	return (&gtcity)
}
func generateCityMap(fileName string) {

	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	var line string
	cities = make(map[int]*city)
	var citiesread = make(map[string]int)

	for {

		line, err = reader.ReadString('\n')
		line = strings.Replace(line, "\n", "", -1)
		str := strings.Split(line, " ")
		if citiesread[str[0]] == 0 { //check if the city has been encountered before
			generateOrGetCity(str[0]) //generate if not
			citiesread[str[0]] = iterator
		}
		//handling neighbors

		for i := 1; i < len(str); i++ {
			st := strings.Split(str[i], "=")
			st[1] = strings.TrimRight(st[1], "\r\n") //represents city
			var neig *city
			if citiesread[st[1]] == 0 {
				neig = generateOrGetCity(st[1])
				citiesread[st[1]] = iterator
			} else {
				neig = cities[citiesread[st[1]]]
			}
			cities[citiesread[str[0]]].neighbors[translateDirection(st[0])] = neig
		}
		if err != nil {
			break
		}
	}
	fmt.Println("Generated Map!!")
	for k, v := range cities {
		fmt.Println("key: ", k, " value: ", v)
		//fmt.Println(v.name)
	}
}

func checkEnd() bool { //checks if every alien is active or not

	for i := 0; i < len(aliens); i++ {
		if aliens[i].active == true {
			return true
		}
	}
	return false
}

func moveAliens() {

	//var ctemp city
	for i := 0; i < len(aliens); i++ {
		if aliens[i].active {
			aliens[i].totalmoves++
			mvdir := move()
			//randomly selects drections to move

			if cities[aliens[i].location].neighbors[mvdir] != nil && !cities[aliens[i].location].neighbors[mvdir].destroyed { //check if the city neighbor exist and that it is not destroyed
				cid := cities[aliens[i].location].neighbors[mvdir].id //name of the neigbouring city
				if citytoalien[cid] != 0 {                            //two aliens in the same
					cities[aliens[i].location].neighbors[mvdir].destroyed = true
					//	delete(cities, cid)
					//cities[cid] = ctemp //destroy the neigbouring city
					//fmt.Println("-----------------------------------------------------------------------------------")
					fmt.Println("Citi ", cities[aliens[i].location].neighbors[mvdir].name, " Destroyed!! ", cities[aliens[i].location].neighbors[mvdir].destroyed)
					//fmt.Println("-----------------------------------------------------------------------------------")
					aliens[i].active = false                //destroy the alien
					aliens[citytoalien[cid]].active = false //destroy the alien already present in the city
					citytoalien[aliens[i].location] = 0     //reset the number of aliens in the current city
					//fmt.Println("Generated Map!!")
					/*for k, v := range cities {
						fmt.Println("key: ", k, " value: ", v)
						//fmt.Println(v.name)
					}*/
				} else { //if there are no conflicts move the alien to the new location and reset the old locations
					citytoalien[aliens[i].location] = 0
					aliens[i].location = cities[cid].id
					citytoalien[cid] = i
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
	rand.Seed(time.Now().UTC().UnixNano())
	generateCityMap("file2.txt")
	generateAlienOnMap(25)
	//fmt.Println("value : ", cities["v"].neighbors["east"])
	moveTillEnd()
}
