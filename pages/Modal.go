package pages

import "github.com/rivo/tview"

type ModalPage struct {
}

type modalPageProps struct {
	text     string
	onSubmit func()
}

func (p *ModalPage) Page(prop modalPageProps) tview.Primitive {
	modal := tview.NewModal().
		SetText(prop.text).
		AddButtons([]string{"Accept", "Cancel"}).
		SetDoneFunc(func(buttonIndex int, buttonLabel string) {
			if buttonLabel == "Cancel" {
			} else if buttonLabel == "Ok" {
				prop.onSubmit()
			}
		})

	return modal
}
