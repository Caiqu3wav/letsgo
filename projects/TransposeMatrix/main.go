package main

import ("fmt")

func Transpose(matrix [][]int) [][]int {
		rows := len(matrix)
		cols := len(matrix[0])


       result := make([][]int, cols)
	   for i := range result {
		result[i] = make([]int, rows)
	   }

	   for i := 0; i < rows; i++ {
		   for j := 0; j < cols; j++ {
			result[j][i] = matrix[i][j]
		   }
	   }
	   for _, row := range result {fmt.Println(row)}
	   	return result
}


func main() {
	matrix := [][]int{
        {1, 2, 3},
        {4, 5, 6},
    }

	Transpose(matrix)

}