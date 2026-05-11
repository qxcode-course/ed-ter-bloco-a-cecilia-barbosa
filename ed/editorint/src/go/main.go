package main

import (
	"fmt"

	"github.com/gdamore/tcell/v2"
)

type Editor struct {
	lines  *List[*List[rune]]
	line   *Node[*List[rune]]
	cursor *Node[rune]
	screen tcell.Screen
	style  tcell.Style
}

func (e *Editor) InsertChar(r rune) {
	e.cursor = e.line.Value.Insert(e.cursor, r)
	e.cursor = e.cursor.Next() // inseriu, vai pra proxima
}

func (e *Editor) KeyLeft() {
	if e.cursor != e.line.Value.Front() { // Se o cursor não está no início da linha
		e.cursor = e.cursor.Prev() // Move o cursor para a esquerda
		return
	}
	// Estamos no início da linha
	if e.line != e.lines.Front() { // Se não está na primeira linha
		e.line = e.line.Prev()        // Move pa ra a linha anterior
		e.cursor = e.line.Value.End() // Move o cursor para o final da linha
	}
}

func (e *Editor) KeyEnter() {
	if e.cursor == e.line.Value.End() {
		novaLinha := NewList[rune]() // nova lista de rune
		e.lines.Insert(e.line.Next(), novaLinha)
		e.line = e.line.Next()
		e.cursor = e.line.Value.Front()
		return
	}

	novaLinha := NewList[rune]()
	for i := e.cursor; i != e.line.Value.End(); {
		novaLinha.PushBack(i.Value)
		i = e.line.Value.Erase(i)
	}
	e.lines.Insert(e.line.Next(), novaLinha)
	e.line = e.line.Next()
	e.cursor = e.line.Value.Front()
}

func (e *Editor) KeyRight() {
	if e.cursor != e.line.Value.Back() { // Se o cursor não está no início da linha
		e.cursor = e.cursor.Next() // Move o cursor para a direita
		return
	}
	// Estamos no início da linha
	if e.line != e.lines.Back() { // Se não está na primeira linha
		e.line = e.line.Next()          // Move para a linha depois
		e.cursor = e.line.Value.Front() // Move o cursor para o final da linha
	}
}

func (e *Editor) KeyUp() {
	if e.line == e.lines.Front() {
		return
	}
	posicao := e.line.Value.IndexOf(e.cursor)
	e.line = e.line.Prev()

	temp := e.line.Value.Front()
	for i := 0; i < posicao && temp != e.line.Value.End(); i++ {
		temp = temp.Next()
	}

	e.cursor = temp
}

func (e *Editor) KeyDown() {
	if e.line == e.lines.Back() {
		return
	}
	posicao := e.line.Value.IndexOf(e.cursor)
	e.line = e.line.Next()

	temp := e.line.Value.Front()
	for i := 0; i < posicao && temp != e.line.Value.End(); i++ {
		temp = temp.Next()
	}

	e.cursor = temp
}

func (e *Editor) KeyBackspace() {
	if e.cursor == e.line.Value.Front() {
		linhaAnterior := e.line.Prev()
		for i := e.line.Value.Front(); i != e.line.Value.End(); i = i.Next() {
			linhaAnterior.Value.PushBack(i.Value)
		}
		e.lines.Erase(e.line)
		e.line = linhaAnterior
		e.cursor = e.line.Value.Back()
		return
	}

	e.cursor = e.line.Value.Erase(e.cursor)
	e.cursor = e.cursor.Prev()

}

func (e *Editor) KeyDelete() {
	if e.cursor != e.line.Value.End() {
		e.cursor = e.line.Value.Erase(e.cursor)
		return
	}

	prox := e.line.Next()
	for i := prox.Value.Front(); i != prox.Value.End(); i = i.Next() {
		e.line.Value.PushBack(i.Value)
	}

	e.lines.Erase(prox)
	e.cursor = e.line.Value.End()
}

func main() {
	// Texto inicial e posição do cursor
	editor := NewEditor()
	editor.Draw()
	editor.MainLoop()
	defer editor.screen.Fini() // Encerra a tela ao sair
}

func (e *Editor) MainLoop() {
	for {
		ev := e.screen.PollEvent()
		switch ev := ev.(type) {
		case *tcell.EventKey:
			switch ev.Key() {
			case tcell.KeyEsc, tcell.KeyCtrlC:
				return
			case tcell.KeyEnter:
				e.KeyEnter()
			case tcell.KeyLeft:
				e.KeyLeft()
			case tcell.KeyRight:
				e.KeyRight()
			case tcell.KeyUp:
				e.KeyUp()
			case tcell.KeyDown:
				e.KeyDown()
			case tcell.KeyBackspace, tcell.KeyBackspace2:
				e.KeyBackspace()
			case tcell.KeyDelete:
				e.KeyDelete()
			default:
				if ev.Rune() != 0 {
					e.InsertChar(ev.Rune())
				}
			}
			e.Draw()
		case *tcell.EventResize:
			e.screen.Sync()
			e.Draw()
		}
	}
}

func NewEditor() *Editor {
	e := &Editor{}
	// Inicializa a tela
	screen, err := tcell.NewScreen()
	if err != nil {
		fmt.Printf("erro ao criar a tela: %v", err)
	}
	if err := screen.Init(); err != nil {
		fmt.Printf("erro ao iniciar a tela: %v", err)
	}
	e.screen = screen
	e.lines = NewList[*List[rune]]()
	e.lines.PushBack(NewList[rune]())
	e.line = e.lines.Front()
	e.cursor = e.line.Value.Back()
	// Define o estilo do texto (branco com fundo preto)
	e.style = tcell.StyleDefault.Foreground(tcell.ColorWhite).Background(tcell.ColorBlack)

	// Limpa a tela e define o estilo base
	e.screen.SetStyle(e.style)
	e.screen.Clear()
	return e
}

func (e *Editor) Draw() {
	e.screen.Clear()
	x := 0
	y := 0
	for line := e.lines.Front(); line != e.lines.End(); line = line.Next() {
		for char := line.Value.Front(); ; char = char.Next() {
			data := char.Value
			if char == line.Value.End() {
				data = '⤶'
			}
			if data == ' ' {
				data = '·'
			}
			if char == e.cursor {
				e.screen.SetContent(x, y, data, nil, e.style.Reverse(true))
			} else {
				e.screen.SetContent(x, y, data, nil, e.style)
			}
			x++
			if char == line.Value.End() {
				break
			}
		}
		y++
		x = 0
	}
	e.screen.Show()
}
