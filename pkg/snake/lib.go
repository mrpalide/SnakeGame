package snake

import (
	"errors"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

// ------- Main Functions -------
func StartGame() {
	start := false
	var game gameState
	var err error
	for !start {
		fmt.Print("Enter Board Size [eg. 10,10]: ")
		var hwString string
		fmt.Scanln(&hwString)
		game, err = initGame(hwString)
		if err != nil {
			fmt.Println(err)
		} else {
			start = true
		}
	}
	next(game)
}

func initGame(hw string) (gameState, error) {
	var game gameState
	var err error

	hwSlice := strings.Split(hw, ",")
	if len(hwSlice) != 2 {
		err = errors.New("Board Size Error!")
		return game, err
	}
	h, err := strconv.Atoi(hwSlice[0])
	if err != nil {
		err = errors.New("Board Size Error!")
		return game, err
	}
	w, err := strconv.Atoi(hwSlice[1])
	if err != nil {
		err = errors.New("Board Size Error!")
		return game, err
	}
	game.HW = [2]int{h, w}
	game.SnakeLenght = 2
	game.SnakeHead = [2]int{1, 0}
	game.SnakeTail = [2]int{0, 0}
	game.LastArrow = "D"
	game.SnakeBody = append(game.SnakeBody, game.SnakeTail, game.SnakeHead)
	game.Food = randomFoodPosition(game.HW, game.SnakeBody)
	return game, err
}

func next(game gameState) {
	var newSnakeHead point
	for {
		flushScreen()
		showScoreboard(game)
		showBoard(game)
		err, newSnakeHead := userInput(newSnakeHead, &game)
		if err != nil {
			continue
		}
		err = logicProcess(newSnakeHead, &game)
		if err != nil {
			flushScreen()
			fmt.Printf("Game Over!\n%s\n\n", err)
			showScoreboard(game)
			break
		}
	}
}

func logicProcess(newPoint point, game *gameState) error {
	var err error
	if existPoint(newPoint, game.SnakeBody) && newPoint != game.SnakeTail {
		err = errors.New("You Eat Yourself!")
	} else {
		game.SnakeHead = newPoint
		if outOfBoard(game.SnakeHead, game.HW) {
			err = errors.New("Out Of Board!")
		} else {
			if game.SnakeHead == game.Food {
				game.Score++
				game.SnakeLenght++
				game.Food = randomFoodPosition(game.HW, game.SnakeBody)
				game.SnakeBody = append(game.SnakeBody, game.SnakeHead, game.SnakeHead)
				game.SnakeBody = game.SnakeBody[1:]
			} else {
				game.SnakeBody = append(game.SnakeBody, game.SnakeHead)
				game.SnakeBody = game.SnakeBody[1:]
			}
			game.SnakeTail = game.SnakeBody[0]
			game.Round++
		}
	}
	return err
}

func userInput(newSnakeHead point, game *gameState) (error, point) {
	var err error
	fmt.Print("L[eft], U[p], R[ight], D[own] | RE[start] |E[xit] \n\n----> ")
	var arrow string
	fmt.Scanln(&arrow)
	switch strings.ToUpper(arrow) {
	case "L":
		if game.LastArrow == "R" {
			err = errors.New("Wrong Input!")
		}
		game.LastArrow = "L"
		newSnakeHead = point{game.SnakeHead[0], game.SnakeHead[1] - 1}
	case "U":
		if game.LastArrow == "D" {
			err = errors.New("Wrong Way!")
		}
		game.LastArrow = "U"
		newSnakeHead = point{game.SnakeHead[0] - 1, game.SnakeHead[1]}
	case "R":
		if game.LastArrow == "L" {
			err = errors.New("Wrong Way!")
		}
		game.LastArrow = "R"
		newSnakeHead = point{game.SnakeHead[0], game.SnakeHead[1] + 1}
	case "D":
		if game.LastArrow == "U" {
			err = errors.New("Wrong Way!")
		}
		game.LastArrow = "D"
		newSnakeHead = point{game.SnakeHead[0] + 1, game.SnakeHead[1]}
	case "E":
		flushScreen()
		fmt.Println("GoodBye!")
		os.Exit(0)
	case "RE":
		StartGame()
	default:
		err = errors.New("Wrong Input!")
	}
	return err, newSnakeHead
}

func outOfBoard(snakeHead, boardHW point) bool {
	if snakeHead[0] >= boardHW[0] || snakeHead[1] >= boardHW[0] || snakeHead[0] < 0 || snakeHead[1] < 0 {
		return true
	}
	return false
}

func showScoreboard(game gameState) {
	fmt.Printf("GAME INFO:\n")
	fmt.Printf(" ------------------- \n")
	fmt.Printf("|Round      |%.5d  |\n", game.Round)
	fmt.Printf("|-----------|-------|\n")
	fmt.Printf("|Lenght     |%.5d  |\n", game.SnakeLenght)
	fmt.Printf("|-----------|-------|\n")
	fmt.Printf("|Score      |%.5d  |\n", game.Score)
	fmt.Printf(" ------------------- \n\n")
}

func showBoard(game gameState) {
	var boardString string
	boardString += fmt.Sprintf(" %s \n", strings.Repeat("-", game.HW[1]*2))
	for i := 0; i < game.HW[0]; i++ {
		boardString += "| "
		for j := 0; j < game.HW[1]; j++ {
			p := point{i, j}
			if p == game.SnakeHead {
				boardString += "@ "
			} else if existPoint(p, game.SnakeBody) {
				boardString += "o "
			} else if p == game.Food {
				boardString += "* "
			} else {
				boardString += ". "
			}
		}
		boardString += "|\n"
	}
	boardString += fmt.Sprintf(" %s \n", strings.Repeat("-", game.HW[1]*2))
	fmt.Printf("GAME BOARD:\n")
	fmt.Print(boardString)
}

func randomFoodPosition(hw [2]int, snakeBody [][2]int) [2]int {
	for {
		position := [2]int{rand.Intn(hw[0]), rand.Intn(hw[1])}
		if !existPoint(position, snakeBody) {
			return position
		}
	}
}

// ------- Tools -------
func existPoint(point [2]int, snakeBody [][2]int) bool {
	for _, bodyPoint := range snakeBody {
		if bodyPoint == point {
			return true
		}
	}
	return false
}

func flushScreen() {
	fmt.Print("\033[H\033[2J")
}

// ------- Types -------
type gameState struct {
	HW          [2]int
	Score       int
	Food        [2]int
	Round       int
	SnakeLenght int
	SnakeHead   [2]int
	SnakeTail   [2]int
	SnakeBody   [][2]int
	LastArrow   string
}

type point [2]int
