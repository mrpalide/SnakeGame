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
	var game GameState
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

func initGame(hw string) (GameState, error) {
	var game GameState
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
	if h <= 1 || w <= 1 {
		err = errors.New("Board Size Error!")
		return game, err
	}
	game.HW = [2]int{h, w}
	game.SnakeLenght = 2
	game.SnakeHead = [2]int{1, 0}
	game.SnakeTail = [2]int{0, 0}
	game.LastDirection = "D"
	game.SnakeBody = append(game.SnakeBody, game.SnakeTail, game.SnakeHead)
	game.Food, _ = randomFoodPosition(game.HW, game.SnakeBody)
	return game, err
}

func next(game GameState) {
	var newSnakeHead point
	for {
		flushScreen()
		showScoreboard(&game)
		showBoard(&game)
		newSnakeHead, err := userInput(newSnakeHead, &game)
		if err != nil {
			continue
		}
		err = logicProcess(newSnakeHead, &game)
		if err != nil {
			flushScreen()
			fmt.Printf("Game Over!\n%s\n\n", err)
			showScoreboard(&game)
			break
		}
	}
}

func logicProcess(newPoint point, game *GameState) error {
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
				game.SnakeBody = append(game.SnakeBody, game.SnakeHead, game.SnakeHead)
				game.SnakeBody = game.SnakeBody[1:]
				game.Food, err = randomFoodPosition(game.HW, game.SnakeBody)
				if err != nil {
					flushScreen()
					showScoreboard(game)
					showBoard(game)
					fmt.Println("You WIN!")
					os.Exit(0)
				}
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

func userInput(newSnakeHead point, game *GameState) (point, error) {
	var err error
	fmt.Print("L[eft], U[p], R[ight], D[own] | RE[start] |E[xit] \n\n----> ")
	var input string
	fmt.Scanln(&input)
	switch strings.ToUpper(input) {
	case "L":
		if game.LastDirection == "R" {
			err = errors.New("Wrong Input!")
		}
		game.LastDirection = "L"
		newSnakeHead = point{game.SnakeHead[0], game.SnakeHead[1] - 1}
	case "U":
		if game.LastDirection == "D" {
			err = errors.New("Wrong Way!")
		}
		game.LastDirection = "U"
		newSnakeHead = point{game.SnakeHead[0] - 1, game.SnakeHead[1]}
	case "R":
		if game.LastDirection == "L" {
			err = errors.New("Wrong Way!")
		}
		game.LastDirection = "R"
		newSnakeHead = point{game.SnakeHead[0], game.SnakeHead[1] + 1}
	case "D":
		if game.LastDirection == "U" {
			err = errors.New("Wrong Way!")
		}
		game.LastDirection = "D"
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
	return newSnakeHead, err
}

func outOfBoard(snakeHead, boardHW point) bool {
	if snakeHead[0] >= boardHW[0] || snakeHead[1] >= boardHW[0] || snakeHead[0] < 0 || snakeHead[1] < 0 {
		return true
	}
	return false
}

func showScoreboard(game *GameState) {
	fmt.Printf("GAME INFO:\n")
	fmt.Printf(" ------------------- \n")
	fmt.Printf("|Round      |%.5d  |\n", game.Round)
	fmt.Printf("|-----------|-------|\n")
	fmt.Printf("|Lenght     |%.5d  |\n", game.SnakeLenght)
	fmt.Printf("|-----------|-------|\n")
	fmt.Printf("|Score      |%.5d  |\n", game.Score)
	fmt.Printf(" ------------------- \n\n")
}

func showBoard(game *GameState) {
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

func randomFoodPosition(hw [2]int, snakeBody [][2]int) ([2]int, error) {
	var err error
	var position point

	for {
		if len(snakeBody) == hw[0]*hw[1] {
			err = errors.New("You WIN!")
			return position, err
		}
		position = [2]int{rand.Intn(hw[0]), rand.Intn(hw[1])}
		if !existPoint(position, snakeBody) {
			return position, err
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
type GameState struct {
	HW            [2]int
	Score         int
	Food          [2]int
	Round         int
	SnakeLenght   int
	SnakeHead     [2]int
	SnakeTail     [2]int
	SnakeBody     [][2]int
	LastDirection string
}

type point [2]int
