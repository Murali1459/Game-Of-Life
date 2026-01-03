package main

import (
	"fmt"
	"math/rand/v2"
	"time"

	tea "github.com/charmbracelet/bubbletea"
)

const FPS = 60
const FRAME_DURATION = time.Second / FPS

var AgeToAscii = map[int]rune{
	0: ' ',
	1: '.',
	2: '░',
	3: '▒',
	4: '█',
}

type Gof struct {
	generation      int
	gameState       string
	age             [][]int
	width           int
	height          int
	userInteraction bool
	x               int
	y               int
}

func (g Gof) getNeighbors(x, y int) (alive int) {
	for yy := -1; yy <= 1; yy++ {
		for xx := -1; xx <= 1; xx++ {
			nX := x + xx
			nY := y + yy
			if nY < 0 || nY >= g.height || nX < 0 || nX >= g.width || (xx == 0 && yy == 0) {
				continue
			}
			if g.age[nY][nX] > 0 {
				alive++
			}
		}
	}
	return alive
}

func (g Gof) Init() tea.Cmd {
	return nil
}

func (g Gof) getNextState() tea.Model {
	nextState := make([][]int, g.height)
	for y := range nextState {
		nextState[y] = make([]int, g.width)
		copy(nextState[y], g.age[y])
	}
	for y := g.height - 1; y >= 0; y-- {
		for x := g.width - 1; x >= 0; x-- {
			alive := g.getNeighbors(x, y)
			aliveCell := g.age[y][x] > 0
			if aliveCell {
				nextState[y][x]++
				if alive < 2 || alive > 3 {
					nextState[y][x] = 0
				}
			} else {
				if alive == 3 {
					nextState[y][x] = 1
				}
			}
		}
	}
	g.generation++
	g.age = nextState
	return g
}

type Update string

func doTick() tea.Cmd {
	return tea.Tick(FRAME_DURATION, func(t time.Time) tea.Msg {
		return Update(t.String())
	})
}

func (g Gof) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.MouseMsg:
		g.x = msg.X
		g.y = msg.Y
		if msg.Action == tea.MouseActionPress {
			g.userInteraction = true
		}
		if g.userInteraction {
			g.age[msg.Y][msg.X] = 1
		}
		if msg.Action == tea.MouseActionRelease {
			g.userInteraction = false
		}
		return g, nil
	case tea.KeyMsg:
		return g, tea.Quit
	case tea.WindowSizeMsg:
		if g.gameState == "Ready" {
			return g, nil
		}
		g.width = msg.Width
		g.height = msg.Height - 1
		g.gameState = "Ready"
		age := [][]int{}
		for range g.height {
			lage := make([]int, g.width)
			for x := range g.width {
				lage[x] = 0
				if rand.Float32() > 0.5 {
					lage[x] = 1
				}
			}
			age = append(age, lage)
		}
		g.age = age
		return g, doTick()
	case Update:
		return g.getNextState(), doTick()
	default:
		return g, nil
	}
}

func (g Gof) View() string {
	if g.gameState != "Ready" {
		return "Initializing...!"
	}
	s := fmt.Sprintf("Generation: %d", g.generation)
	mouse := fmt.Sprintf("Location: (%d,%d)", g.x, g.y)
	for range g.width - (len(s) + len(mouse)) {
		s += " "
	}
	s += mouse
	for _, y := range g.age {
		ls := ""
		for _, x := range y {
			age := min(len(AgeToAscii), x)
			ls += fmt.Sprintf("%s", string(AgeToAscii[age]))
		}
		s += ls + "\n"
	}
	return s
}
