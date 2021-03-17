# SnakeGame  
### About 
This is a simple version of SnakeGame that write for SkyCoin Go developer position test task.  
### Libraries
It's write in pure Go without any 3rd-party libraries.  
Two simple printing fucntion show GAME INFO and GAME BOARD in your cli.  
No key-binding library used, so for new direction you should use **L,U,R,D** (or in lowercase) as input, and for *exit* and *restart* game you can use **E** and **RE**.  
### Game RULES!
- You cannot take an oposite direction of your last move. Your back (as snake) will break! :]
- You cannot move over your body! If you try, it means eating yourself, and the beginning of the end of your snakey life! ;)
- You cannot get out from the board! Where are you going dude? :D  
### How to Play?
Just clone and run!
```
git clone https://github.com/mrpalide/SnakeGame.git
cd SnakeGame
go run ./cmd/snake
```
After starting the game, you should choose the board size in **x,y** format.  
```
Enter Board Size [eg. 10,10]: x,y
```
Finaly, enjoy your snakey life! :)  
### How Long? <font size="1">[SkyCoin Question]</font>
```
Q: How long it took you to finish this test task?  
A: Near 5 hours!
```