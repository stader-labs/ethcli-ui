package pages

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type PageInterface interface {
	// These will be called during initialisation
	// in pages/init.go
	SetPageID(id string)
	SetApp(app *tview.Application)

	// These will be implemented by each page
	Page() tview.Primitive
	HandleEvents(event *tcell.EventKey) *tcell.EventKey
	GetFirstElement() tview.Primitive
	OnResume()
	GoBack()
}

type PageType struct {
	ID  string
	App *tview.Application
}

func (p *PageType) Page() tview.Primitive {
	log.Warnf("Implement Page: %s")

	p.App.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		return p.HandleEvents(event)
	})

	return nil
}

func (p *PageType) HandleEvents(event *tcell.EventKey) *tcell.EventKey {
	key := event.Key()
	log.Warnf("Implement HandleEvents [key]: %s", key)

	if key == tcell.KeyEscape {
		p.GoBack()
	}

	return event
}

func (p *PageType) OnResume() {
	log.Warnf("Implement OnResume: %s", p.ID)
}

// GetFirstElement returns the first element in the page,
// so that the app can set focus on it
func (p *PageType) GetFirstElement() tview.Primitive {
	log.Warnf("Implement GetFirstElement: %s", p.ID)
	return nil
}

func (p *PageType) GoBack() {
	log.Warnf("Implement GoBack: %s", p.ID)
}

func (p *PageType) SetPageID(id string) {
	p.ID = id
}

func (p *PageType) SetApp(app *tview.Application) {
	p.App = app
}
