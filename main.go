package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"

	"github.com/jung-kurt/gofpdf"
)

const (
	PageSize    = 210.0 // PDF page size in millimeters (A4)
	CellSize    = 10.0  // Cell size in millimeters
	FontSize    = 8.0   // Font size in points
	WordsPerRow = 10    // Number of words to display per row
)

var (
	Letters = []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ")
)

func main() {
	words := []string{"HELLO", "WORLD", "GO", "PROGRAMMING", "WORDSEARCH", "ACTIVITY", "TAVITO", "RUST", "JAVASCRIPT"}

	// Generate the word search grid
	grid := generateGrid(words)

	// Generate PDF
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetFont("Arial", "", FontSize)

	for row := 0; row < len(grid); row++ {
		for col := 0; col < len(grid[row]); col++ {
			x := float64(col) * CellSize
			y := float64(row) * CellSize

			pdf.SetXY(x, y)
			pdf.Cell(CellSize, CellSize, string(grid[row][col]))
		}
	}

	// Save PDF to file
	err := pdf.OutputFileAndClose("wordsearch.pdf")
	if err != nil {
		fmt.Println("Error generating PDF:", err)
		return
	}

	fmt.Println("PDF generated successfully!")
}

func generateGrid(words []string) [][]rune {
	rand.Seed(time.Now().UnixNano())

	// Create an empty grid
	grid := make([][]rune, int(PageSize/CellSize))
	for i := range grid {
		grid[i] = make([]rune, int(PageSize/CellSize))
	}

	// Place words in the grid
	for _, word := range words {
		word = strings.ToUpper(word)
		length := len(word)
		direction := rand.Intn(4) // 0: horizontal, 1: vertical, 2: diagonal-up, 3: diagonal-down

		for {
			// Randomly choose a starting position for the word
			row := rand.Intn(len(grid))
			col := rand.Intn(len(grid[row]))

			// Check if the word fits in the chosen direction
			if fits(grid, word, length, direction, row, col) {
				// Place the word in the grid
				placeWord(grid, word, length, direction, row, col)
				break
			}
		}
	}

	// Fill empty cells with random letters
	fillEmptyCells(grid)

	return grid
}

func fits(grid [][]rune, word string, length, direction, row, col int) bool {
	// Check if the word fits horizontally
	if direction == 0 && col+length <= len(grid[row]) {
		for i := 0; i < length; i++ {
			if grid[row][col+i] != 0 && grid[row][col+i] != rune(word[i]) {
				return false
			}
		}
		return true
	}

	// Check if the word fits vertically
	if direction == 1 && row+length <= len(grid) {
		for i := 0; i < length; i++ {
			if grid[row+i][col] != 0 && grid[row+i][col] != rune(word[i]) {
				return false
			}
		}
		return true
	}

	// Check if the word fits diagonally (upward)
	if direction == 2 && row+1 >= length && col+length <= len(grid[row]) {
		for i := 0; i < length; i++ {
			if grid[row-i][col+i] != 0 && grid[row-i][col+i] != rune(word[i]) {
				return false
			}
		}
		return true
	}

	// Check if the word fits diagonally (downward)
	if direction == 3 && row+length <= len(grid) && col+length <= len(grid[row]) {
		for i := 0; i < length; i++ {
			if grid[row+i][col+i] != 0 && grid[row+i][col+i] != rune(word[i]) {
				return false
			}
		}
		return true
	}

	return false
}

func placeWord(grid [][]rune, word string, length, direction, row, col int) {
	for i := 0; i < length; i++ {
		switch direction {
		case 0: // Horizontal
			grid[row][col+i] = rune(word[i])
		case 1: // Vertical
			grid[row+i][col] = rune(word[i])
		case 2: // Diagonal (upward)
			grid[row-i][col+i] = rune(word[i])
		case 3: // Diagonal (downward)
			grid[row+i][col+i] = rune(word[i])
		}
	}
}

func fillEmptyCells(grid [][]rune) {
	for row := 0; row < len(grid); row++ {
		for col := 0; col < len(grid[row]); col++ {
			if grid[row][col] == 0 {
				grid[row][col] = getRandomLetter()
			}
		}
	}
}

func getRandomLetter() rune {
	return Letters[rand.Intn(len(Letters))]
}
