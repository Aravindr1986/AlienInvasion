# Alien Invasion

### Introduction
The program is to develop a simulation for an Alien invasion. A map of cities and its neighbors are provided,
The aliens can appear in random at any city. The aliens then randomly move to any of the neighboring cities. 
A city gets destroyed when two aliens are in the same city at the same time. This results in both the aliens 
and the city being destroyed. 

### Assumptions
The following assumptions were made in the program
1) An alien becomes inactive(destroyed) when it has reached the end of the steps or two are found in the same city.
2) The number of aliens should be less than the number of cities.
3) In the intial step no two aliens appear in the same city.
4) Even if an alien is trapped, It countinues until 10000 steps are reached. 
5) The input file format does not change and the last line is a city entry.
6) The system running the program as go installed in them. 

### Approach 
The solution program genrates the a map of the cities based on the input file. It then randomly assigns the aliens based on the number of 
aliens and cities. Then it loops through each alien and moves the alien one step at a time. During each round the program checks to see if 
city to which it is moving to is already occupied by another alien. If it is, then that city along with the two aliens are deactivated.
The map of the cities and aliens are done using map data structure in go along with struct and pointers. 

### Run the program
To run the program, copy the AlienInvasion.go file and the input file to a folder an issue the below command

go run AlienInvasion.go <pathtofile/cityfilename> aliencount


