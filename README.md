
### Overview: 
 To design & implement Mad aliens invastion to the imaginary world.

### Design & assumptions summary
 1. Create a world map using provided file.
 2. Create `N` number of  aliens based on input
 3. We assume that there can be no more than twice the number of aliens as there are cities
 4. Similulate the movement of aliens, where the aliens randomly choose a neighboring city to move to. If an alien encounters another alien in the new city, `they fight and destroy the city`
 5. When a city is destroyed, it is removed from the map, along with any roads leading to or from it
 6. The program prints the movement of each alien and any encounters or city destructions that occur.
 7. After a certain number of iterations, the program stops, and the final state of the map is printed.

<img src="images/alien-invasion.png" />

### Implementation Approach 
After understanding the requirement, i tried to split it into possible module to start development & my seqeunece apparoach was 
   1. Defining Data Structure like cities, Alien etc
   2. Methods to operate on these data structure like `readWorldMap`,`CreateAliens` etc
   3. Start implementing logic for `SimulateAlienMovement` operation & other utility methods to support the process.

### Data Structures
    Started by defining the necessary data structures, understanding them is important as whole logic revolves around them, some main data structure are mentioned as below
        1. City Type `contains City Name, its Neighboring cities and their directions as map & list of Aliens assign to that city`
        2. Alien Type 'its object to represent key info about Alien e.g Alien ID, Alien state & most importently current city
        3. WorldMap `Its a map of `City Type` 



Summary:
The program represents a world map with various cities. Each city is connected to its neighboring cities by roads.
1. There are aliens moving around the cities randomly.
2. The program starts by reading the world map from a file and printing it on the screen.
3. It then creates a specified number of aliens and assigns them random starting cities.
4. The initial positions of the aliens are printed on the screen.
5. The program simulates the movement of aliens over a number of iterations.
    In each iteration, the aliens randomly choose a neighboring city to move to.
    If an alien encounters another alien in the new city, they fight and destroy the city.
    When a city is destroyed, it is removed from the map, along with any roads leading to or from it.
    The program prints the movement of each alien and any encounters or city destructions that occur.
    After a certain number of iterations, the program stops, and the final state of the map is printed.

### Assumptions 
1. The number of aliens should not greater then number of cities.
2. 