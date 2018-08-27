package main

import (
  "fmt"
  "bufio"
  "os"
  "strconv"
  "math/rand"
)

func UserInput(prompt string) string {
  fmt.Print(prompt)
  input := bufio.NewScanner(os.Stdin)
  input.Scan()
  return input.Text()
}

func IterateGrid(length int, function func(i,j int)) {
  for i := 0; i < length; i++ {
    for j := 0; j < length; j++ {
      function(i,j)
    }
  }
}

func IndicesInbounds(length,i,j int) bool {
  i_in_bounds := i >= 0 && i < length
  j_in_bounds := j >= 0 && j < length

  return i_in_bounds && j_in_bounds
}

func BuildGrid(length int) [][]int {
  tic := make([][]int, length)
  for i := 0; i < length; i++ {
    tic[i] = make([]int, length)
  }
  return tic
}

func SeedGrid(grid [][]int) {
  IterateGrid(len(grid), func(i,j int) {
    grid[i][j] = rand.Intn(2)
  })
}

func CountAliveNeighbors(grid [][]int,i,j int) int {
  alive_neighbors := 0
  neighbor_offsets := [8][2]int{{0,1},{0,-1},{1,0},{-1,0},{1,1},{1,-1},{-1,1},{-1,-1},}

  for k := 0; k < len(neighbor_offsets); k++ {
    i_offset := i + neighbor_offsets[k][0]
    j_offset := j + neighbor_offsets[k][1]

    if IndicesInbounds(len(grid),i_offset,j_offset) {
      alive_neighbors += grid[i_offset][j_offset]
    }
  }
  return alive_neighbors
}

func AliveNextTic(value,alive_neighbors int) bool {
  return alive_neighbors == 3 || (value == 1 && alive_neighbors == 2)
}

func CreateNextTic(tic,next_tic [][]int) [][]int {
  IterateGrid(len(tic), func(i,j int) {
    if AliveNextTic(tic[i][j],CountAliveNeighbors(tic,i,j)) {
      next_tic[i][j] = 1
    } else {
      next_tic[i][j] = 0
    }
  })
  return next_tic
}

func main() {
  length, err := strconv.Atoi(UserInput("Enter value for dimensions of Conway's game: "))
  if err != nil {
    length = 0
  }

  tic, next_tic := BuildGrid(length), BuildGrid(length)
  SeedGrid(tic)

  tick_count := 0
  for {
    if UserInput("\nContinue to next tic (y/n): ") == "n" {
      break
    } else {
      tick_count++
      fmt.Print("\nTIC #", tick_count, "\n")
      for i := 0; i < length; i++ {
				fmt.Println(tic[i])
      }
      tic = CreateNextTic(tic,next_tic)
    }
  }
}
