// dijkstra implement of dijkstra's algorithm
// which is an algorithm for finding the shortest paths between nodes in a graph.

package shortestpath

import (
	"fmt"
)

const (
	num = 6
)

var (
	element  [num][num]int
	distance [num]int
	book     [num]int
	infinity = 99999999

	vertex int
	side   int

	vertexone    int
	vertextwo    int
	lineDistance int
)

func dijkstra() {
	// Input numbers of vertex and side
	fmt.Scan(&vertex, &side)

	// initialize two demensions array
	for i := 0; i < vertex; i++ {
		for j := 0; j < vertex; j++ {
			if i == j {
				element[i][j] = 0
			} else {
				element[i][j] = infinity
			}
		}
	}

	//  Read sides
	for i := 0; i < side; i++ {
		fmt.Scan(&vertexone, &vertextwo, &lineDistance)
		element[vertexone][vertextwo] = lineDistance
	}

	// Initialize element array, from first vertex to other vertexs
	for i := 0; i < vertex; i++ {
		distance[i] = element[0][i]
	}

	for i := 0; i < vertex; i++ {
		book[i] = 0
	}

	// firsr vertex
	book[0] = 1

	var u, v int
	for i := 0; i < vertex; i++ {
		min := infinity

		// Search the min value in distance array
		for j := 0; j < vertex; j++ {
			if book[j] == 0 && distance[j] < min {
				min = distance[j]
				u = j
			}
		}

		book[u] = 1
		for v = 0; v < vertex; v++ {
			if element[u][v] < infinity {
				if distance[v] > distance[u]+element[u][v] {
					distance[v] = distance[u] + element[u][v]
				}
			}
		}
	}

	for i := 0; i < vertex; i++ {
		fmt.Println(distance[i])
	}
}

//    test datas
//    6 9
//    0 1 1
//    0 2 12
//    1 2 9
//    1 3 3
//    2 4 5
//    3 2 4
//    3 4 13
//    3 5 15
//    4 5 4
