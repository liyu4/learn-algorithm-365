// dijkstra implement of dijkstra's algorithm
// which is an algorithm for finding the shortest paths between nodes in a graph.

package dijkstra

import (
	"fmt"
)

const (
	num = 10
)

var (
	element  [num][num]int
	distance [num]int
	book     [num]int
	infnity  = 99999999

	vertex int
	side   int

	sideone   int
	sidetwo   int
	sidethree int
)

func dijkastra() {
	// Input numbers of vertex and side
	fmt.Scan(&vertex, &side)

	// initialize two demensions array
	for i := 0; i < vertex; i++ {
		for j := 0; j < side; j++ {
			if i == j {
				element[i][j] = 0
			} else {
				element[i][j] = infnity
			}
		}
	}

	// Read sides
	for i := 0; i < side; i++ {
		fmt.Scan(&sideone, &sidetwo, &sidethree)
	}

	// Initialize element array, from first vertex to other vertexs
	for i := 0; i < vertex; i++ {
		distance[i] = element[1][i]
	}

	for i := 0; i < vertex; i++ {
		book[i] = 0
	}

	book[1] = 1

	var u, v int
	for i := 1; i < vertex-1; i++ {
		min := infnity

		for j := 1; j <= vertex; j++ {
			if book[j] == 0 && distance[j] < min {
				min = distance[j]
				u = j
			}
		}

		book[u] = 1
		for v = 1; v <= vertex; v++ {
			if element[u][v] < infnity {
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

// test datas

// 6 9
// 1 2 1
// 1 3 12
// 2 3 9
// 2 4 3
// 3 5 5
// 4 3 4
// 4 5 13
// 4 6 15
// 5 6 4
