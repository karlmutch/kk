package k

import (
	"github.com/jroimartin/gocui"
)

type UIReducer struct {
	Table        *TableView
	Modal        *ModalView
	State        *StateView
	Debug        *DebugView
	ActiveScreen Screen
}

type DebugView struct {
	Cursor int
}

type StateView struct {
	Cursor int
}

type TableView struct {
	Cursor   int
	Selected string
	Kind     Kind
	Lines    [][]string
	Headers  []string
}

type ModalView struct {
	Title    string
	Kind     Kind
	Cursor   int
	Selected string
	Lines    []string
}

func (ur *UIReducer) CursorMove(g1 *gocui.Gui, delta int) {
	g1.Update(
		func(g *gocui.Gui) error {
			switch ur.ActiveScreen {
			case ScreenDebug:
				ur.Debug.Cursor = ur.Debug.Cursor + delta
				if ur.Debug.Cursor < 0 {
					ur.Debug.Cursor = 0
				}
			case ScreenModal:
				ur.Modal.Cursor = ur.Modal.Cursor + delta
				if ur.Modal.Cursor < 0 {
					ur.Modal.Cursor = 0
				}
				if ur.Modal.Cursor > len(ur.Modal.Lines)-1 {
					ur.Modal.Cursor = len(ur.Modal.Lines) - 1
				}
				if len(ur.Modal.Lines) > 0 {
					ur.Modal.Selected = ur.Modal.Lines[ur.Modal.Cursor]
				}
			case ScreenTable:
				// originalPosition := ur.Table.Cursor
				if len(ur.Table.Lines) < 2 {
					return nil
				}
				ur.Table.Cursor = ur.Table.Cursor + delta
				switch {
				case ur.Table.Cursor < 1:
					ur.Table.Cursor = 1
				case ur.Table.Cursor > len(ur.Table.Lines):
					ur.Table.Cursor = len(ur.Table.Lines)
				}
				ur.Table.Selected = ur.Table.Lines[ur.Table.Cursor-1][0]
			case ScreenState:
				ur.State.Cursor = ur.State.Cursor + delta
				if ur.State.Cursor < 0 {
					ur.State.Cursor = 0
				}
			}
			return nil
		},
	)
}

func (tv *TableView) SetHeaders(g1 *gocui.Gui, vals []string) {
	g1.Update(
		func(g *gocui.Gui) error {
			tv.Headers = vals
			return nil
		},
	)
}

func (ur *UIReducer) SetTableActive(g1 *gocui.Gui) {
	g1.Update(
		func(g *gocui.Gui) error {
			ur.ActiveScreen = ScreenTable
			return nil
		},
	)
}

func (ur *UIReducer) SetStateActive(g1 *gocui.Gui) {
	g1.Update(
		func(g *gocui.Gui) error {
			ur.ActiveScreen = ScreenState
			return nil
		},
	)

}
func (ur *UIReducer) SetDebugActive(g1 *gocui.Gui) {
	g1.Update(
		func(g *gocui.Gui) error {
			ur.ActiveScreen = ScreenDebug
			return nil
		},
	)
}
func (ur *UIReducer) SetModalActive(g1 *gocui.Gui) {
	g1.Update(
		func(g *gocui.Gui) error {
			ur.ActiveScreen = ScreenModal
			return nil
		},
	)
}

func (p *ModalView) SetModalKind(g1 *gocui.Gui, kind Kind) {
	g1.Update(
		func(g *gocui.Gui) error {
			p.Kind = kind
			return nil
		},
	)
}

func (p *TableView) SetKind(g1 *gocui.Gui, kind Kind) {
	g1.Update(
		func(g *gocui.Gui) error {
			p.Kind = kind
			return nil
		},
	)
}

func (p *TableView) SetLines(g1 *gocui.Gui, lines [][]string) {
	g1.Update(
		func(g *gocui.Gui) error {
			p.Lines = lines
			return nil
		},
	)
}

func (p *ModalView) SetLines(g1 *gocui.Gui, lines []string) {
	g1.Update(
		func(g *gocui.Gui) error {
			p.Lines = lines
			return nil
		},
	)
}

func (p *ModalView) SetTitle(g1 *gocui.Gui, title string) {
	g1.Update(
		func(g *gocui.Gui) error {
			p.Title = title
			return nil
		},
	)
}