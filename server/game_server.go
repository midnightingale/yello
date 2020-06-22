package main

import (
	"fmt"
	"strconv"
	"strings"
)

type direction string

const (
	startingPacLives int = 3

	left  direction = "L"
	right direction = "R"
	up    direction = "U"
	down  direction = "D"
)

type playerGameData struct { //TODO LOAD GAME DATA
	p                  *clientPlayer
	x                  int
	y                  int
	latestDirection    direction
	tileRepresentation tile
}

type game struct {
	p1, p2      *playerGameData
	active      bool
	maze        [][]tile
	pacLives    int
	pelletsLeft int
}

func initializeGameServer(p1, p2 *clientPlayer) {

	//handle maze
	maze, numRows := loadAndParseMazeFile(*mazeFile)

	if maze == nil {
		writeToAllPlayers(gameInstance, "PONG GAME-INVALID")
		disconnectAllPlayers(gameInstance)
		return //terminate game early!
	}

	var mazeMsg strings.Builder
	fmt.Fprintf(&mazeMsg, "%s-%s", strconv.Itoa(len(maze[0])), strconv.Itoa(numRows))
	for _, val := range maze {
		for _, el := range val {
			fmt.Fprintf(&mazeMsg, "-%s", string(el))
		}
	}

	gameInstance := &game{
		p1: &playerGameData{p: p1},
		p2: &playerGameData{p: p2},
		//p3:     &playerGameData{p: p3},
		//p4:     &playerGameData{p: p4},
		//p5:     &playerGameData{p: p5},
		active: true,
	}

	//now construct bit maze for us yes
	gameInstance.maze = constructBitMaze(maze)

	gameInstance.pacLives = startingPacLives

	//handle game initialization
	p1.writeChanneledMessage("PONG GAME-INIT " + "P1-" + p2.name) //Who needs JSON when you got -?
	p2.writeChanneledMessage("PONG GAME-INIT " + "P2-" + p1.name)

	//p1.writeChanneledMessage("PONG GAME-INIT " + "P1-" + p2.name + "-" + p3.name + "-" + p4.name + "-" + p5.name) //Who needs JSON when you got -?
	//p2.writeChanneledMessage("PONG GAME-INIT " + "P2-" + p1.name + "-" + p3.name + "-" + p4.name + "-" + p5.name)
	//p3.writeChanneledMessage("PONG GAME-INIT " + "P3-" + p1.name + "-" + p2.name + "-" + p4.name + "-" + p5.name)
	//p4.writeChanneledMessage("PONG GAME-INIT " + "P4-" + p1.name + "-" + p2.name + "-" + p3.name + "-" + p5.name)
	//p5.writeChanneledMessage("PONG GAME-INIT " + "P5-" + p1.name + "-" + p2.name + "-" + p3.name + "-" + p4.name)

	writeToAllPlayers(gameInstance, "PONG SET-LEVEL "+mazeMsg.String())

	p1.activeGame = gameInstance
	p2.activeGame = gameInstance
	//p3.activeGame = &gameInstance
	//p4.activeGame = &gameInstance
	//p5.activeGame = &gameInstance

	tendGame(gameInstance)
}

func tendGame(g *game) {

	//WRITE CODE HERE

	for g.active {
		//move players
		for _, player := range []*playerGameData{
			g.p1, g.p2} {
			projX, projY := player.x, player.y
			if player.latestDirection == "R" {
				projX += 1
			}
			if player.latestDirection == "L" {
				projX -= 1
			}
			if player.latestDirection == "U" {
				projY += 1
			}
			if player.latestDirection == "D" {
				projY -= 1
			}
		}
		//check collision with wall
		//check teleport
		//eat pellet
		//check collision
		////check super State

		//ze game loop!
	}
}

//INNER PROCESSES

func checkUpdateObjectStates() {

}

//TO ZE CLIENTSSSS

func updateObjectPosition(g *game, id string, v *posVector) {
	writeToAllPlayers(g, "PONG GAME-ENTITY-POS "+id+"-"+v.toString())
}

func sendPelletConsumed(g *game, v *posVector) {
	writeToAllPlayers(g, "PONG GAME-PELLET-HOM "+v.toString()) //client should get rid of pellet AND increase score for pac-man (if it is keeping track for visual purposes)
}

func sendScared(g *game) {
	writeToAllPlayers(g, "PONG GAME-SCARED 1")
}

func sendCeaseScared(g *game) {
	writeToAllPlayers(g, "PONG GAME-SCARED 0")
}

//func updateObjectState(g *game, o *gameObject, state objectState) {
//	writeToAllPlayers(g, "PONG GAME-OBJECT-STATE " + o.string_ID + "-" + string(objectState))
//}

//FROM ZE CLIENT

func playerUpdateDesiredDirection(req *playerRequest, argument string) {
	p := pickPlayerGameData(req.p.activeGame, req.p)
	switch argument {
	case "R":
		p.latestDirection = right
		break
	case "L":
		p.latestDirection = left
		break
	case "U":
		p.latestDirection = up
		break
	case "D":
		p.latestDirection = down
		break
	}
}

//FUNCTIONALITY

func moveEntity() {

}

//INNER UTILS

func writeToAllPlayers(g *game, msg string) {
	g.p1.p.writeChanneledMessage(msg)
	g.p2.p.writeChanneledMessage(msg)
}

func disconnectAllPlayers(g *game) {
	g.p1.p.disconnectChannel <- struct{}{}
	g.p2.p.disconnectChannel <- struct{}{}
}

func getNumberOfPellets(g *game) int {
	return 0
}

func constructBitMaze(sMaze [][]string) [][]tile {
	tileMaze := make([][]tile, 2) //TODO BETTER LEN HANDLING
	for row := 0; row < len(sMaze); row++ {
		for col := 0; col < len(sMaze[row]); col++ {
			tileMaze[row][col] = (tileToBit(sMaze[row][col]))
		}
	}
	return (tileMaze)
}

func pickPlayerGameData(g *game, p *clientPlayer) *playerGameData {
	switch p {
	case g.p1.p:
		return g.p1
	case g.p2.p:
		return g.p2
	}
	return nil
}
