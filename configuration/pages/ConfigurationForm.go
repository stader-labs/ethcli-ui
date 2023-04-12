package pages

import (
	"fmt"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
	"github.com/stader-labs/ethcli-ui/configuration/components"
	"github.com/stader-labs/ethcli-ui/configuration/config"
	"github.com/stader-labs/ethcli-ui/configuration/state"
)

type ConfigurationForm struct {
	*PageType
	form            *tview.Form
	descriptionView *tview.TextView
	header          *tview.TextView
}

func (p *ConfigurationForm) Page() tview.Primitive {
	p.form = components.Form()
	p.descriptionView = tview.NewTextView()

	descriptionContent := tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(nil, 2, 1, false).
		AddItem(p.descriptionView, 0, 1, false).
		AddItem(nil, 2, 1, false)

	body := tview.NewFlex().
		SetDirection(tview.FlexColumn).
		AddItem(nil, 1, 1, false).
		AddItem(p.form, 100, 1, false).
		AddItem(nil, 1, 1, false).
		AddItem(components.VerticalLine(tcell.ColorGray), 1, 1, false).
		AddItem(nil, 2, 1, false).
		AddItem(descriptionContent, 60, 1, false).
		AddItem(nil, 0, 1, false)

	p.header = components.Header(
		fmt.Sprintf("Configuration Summary > %s", state.Categories.Option.Selected),
	)

	root := tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(p.header, 3, 1, false).
		AddItem(body, 0, 1, false).
		AddItem(
			components.Footer(config.PageID.ConfigurationForm, p.App, p.GoBack),
			5, 1, false,
		)

	return root
}

func (p *ConfigurationForm) addFormItemEvents(fr fieldReturnType) {
	switch fr.info.Type {
	case "select":
		f := fr.field.(*tview.DropDown).SetSelectedFunc(func(text string, index int) {
			log.Infof("f.SetSelectedFunc on field: %s, text: %s, index: %d", fr.info.Label, text, index)
			state.Configuration[fr.info.Key] = text
			p.updateDescription(fr.info)
			if fr.info.Children != nil && len(fr.info.Children) > 0 {
				p.updateForm()
			}
		})

		f.GetList().SetChangedFunc(func(index int, mainText string, secondaryText string, shortcut rune) {
			log.Infof("f.List.SetChangedFunc on field: %s, index: %d, mainText: %s, secondaryText: %s, shortcut: %d", fr.info.Label, index, mainText, secondaryText, int32(shortcut))
			if f.IsOpen() {
				if fr.info.OptionDescriptions != nil {
					optionDesc := fr.info.OptionDescriptions[mainText]
					p.descriptionView.SetText(optionDesc)
				} else {
					p.descriptionView.SetText("")
				}
			}
		})

		f.SetFocusFunc(func() {
			log.Infof("f.SetFocusFunc on field: %s", fr.info.Label)
			p.updateDescription(fr.info)
		})
	case "text", "int":
		fr.field.(*tview.InputField).
			SetChangedFunc(func(text string) {
				state.Configuration[fr.info.Key] = text
			}).
			SetFocusFunc(func() {
				log.Infof("f.SetFocusFunc on field: %s", fr.info.Label)
				p.updateDescription(fr.info)
			})
	case "checkbox":
		fr.field.(*tview.Checkbox).
			SetChangedFunc(func(checked bool) {
				log.Infof("f.SetChangedFunc on field: %s, checked: %t", fr.info.Label, checked)
				state.Configuration[fr.info.Key] = checked

				if fr.info.Children != nil && len(fr.info.Children) > 0 {
					p.updateForm()
				}
			}).
			SetFocusFunc(func() {
				log.Infof("f.SetFocusFunc on field: %s", fr.info.Label)
				p.updateDescription(fr.info)
			})
	}
}

func (p *ConfigurationForm) updateDescription(fr config.FormFieldType) {
	desc := fr.Description
	p.descriptionView.SetText(desc)
}

func (p *ConfigurationForm) updateForm() {
	p.form.Clear(true)
	log.Infof("Updating form for [%s]", state.Categories.Option.Selected)
	pageFields, ok := config.ConfigurationFields[state.Categories.Option.Selected]
	log.Infof("total pageFields: %d for category [%s]", len(pageFields), state.Categories.Option.Selected)
	if ok {
		fs := getFields(pageFields)
		log.Infof("Found [%d] fields for [%s]", len(fs), state.Categories.Option.Selected)

		for _, fr := range fs {
			p.addFormItemEvents(fr)
			p.form.AddFormItem(fr.field)
		}
	} else {
		log.Errorf("No fields found for [%s]", state.Categories.Option.Selected)
	}
	p.App.SetFocus(p.form)
}

func (p *ConfigurationForm) HandleEvents(event *tcell.EventKey) *tcell.EventKey {
	return event
}

func (p *ConfigurationForm) OnResume() {
	p.descriptionView.SetText("")
	log.Infof("OnResume: [%s]", config.PageID.ConfigurationForm)
	p.header.SetText(fmt.Sprintf("\nConfiguration Summary > %s", state.Categories.Option.Selected))
	p.updateForm()
}

func (p *ConfigurationForm) GoBack() {
	log.Infof("Going back to [%s]", config.PageID.Categories)
	ChangePage(config.PageID.Categories, p.App)
}

func (p *ConfigurationForm) GetFirstElement() tview.Primitive {
	return p.form
}
