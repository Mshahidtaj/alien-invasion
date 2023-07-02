
Problem Statement: 




### Terminology
 1. Graph
        A graph is a data structure used to represent a collection of interconnected nodes, called `vertics`, where each `vertic` can be connected to one or more other vertics through `Edges`.
        Graphs are commonly used to model relationships & connections between various entities. e.g Social media use graphs to represent connections between users. Each user is a vertex, and the connections (friendships, follows, connections) between users are represented as edges.
### Data Structures
    Start by defining the necessary data structures to represent the 
            1. map, 
            2. cities
            3. roads.



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