package pages

import (
	"github.com/rivo/tview"
	"github.com/stader-labs/ethcli-ui/configuration/components"
	"github.com/stader-labs/ethcli-ui/configuration/config"
	"github.com/stader-labs/ethcli-ui/configuration/state"
	"github.com/stader-labs/ethcli-ui/configuration/utils"
)

func ChangePage(nextPage string, app *tview.Application) {
	currentPageName, _ := Pages.GetFrontPage()
	log.Infof("ChangePage: from [%s] to [%s]", currentPageName, nextPage)
	Pages.SwitchToPage(nextPage)
	pageInstance := All[nextPage]
	firstElement := pageInstance.GetFirstElement()
	app.SetFocus(firstElement)
	pageInstance.OnResume()
}

type fieldReturnType struct {
	info  config.FormFieldType
	field tview.FormItem
}

func getSelectField(field config.FormFieldType) fieldReturnType {
	selectField := components.Dropdown().
		SetLabel(field.Label).
		SetOptions(field.Options, nil)

	// SetSelectedFunc(func(text string, index int) {
	// 	log.Infof("SetSelectedFunc on field: %s, text: %s, index: %d", field.Label, text, index)
	// }).
	// SetDoneFunc(func(key tcell.Key) {
	// 	log.Infof("SetDoneFunc on field: %s, key: %s", field.Label, key)
	// })

	value, ok := state.Configuration[field.Key]
	if ok {
		valueString, err := utils.InterfaceToString(value)
		if err != nil {
			log.Errorf("Unable to convert interface to string. Error: %s, key: %s", err, field.Key)
		} else {
			valueIndex := utils.IndexOf(field.Options, valueString)
			selectField.SetCurrentOption(valueIndex)
		}
	}

	return fieldReturnType{field, selectField}
}

func getIntField(field config.FormFieldType) fieldReturnType {
	inputField := tview.NewInputField().
		SetLabel(field.Label).
		SetFieldWidth(0)

	inputField.SetAcceptanceFunc(func(textToCheck string, lastChar rune) bool {
		return utils.IsInt(textToCheck)
	})

	value, ok := state.Configuration[field.Key]
	if ok {
		valueString, err := utils.InterfaceToString(value)
		if err != nil {
			log.Errorf("Unable to convert interface to string. Error: %s, key: %s", err, field.Key)
		} else {
			inputField.SetText(valueString)
		}
	}

	return fieldReturnType{field, inputField}
}

func getTextField(field config.FormFieldType) fieldReturnType {
	inputField := tview.NewInputField().
		SetLabel(field.Label).
		SetFieldWidth(0)

	value, ok := state.Configuration[field.Key]
	if ok {
		valueString, err := utils.InterfaceToString(value)
		if err != nil {
			log.Errorf("Unable to convert interface to string. Error: %s, key: %s", err, field.Key)
		} else {
			inputField.SetText(valueString)
		}
	}

	return fieldReturnType{field, inputField}
}

func getCheckboxField(field config.FormFieldType) fieldReturnType {
	checkboxField := components.Checkbox().SetLabel(field.Label)

	if value, ok := state.Configuration[field.Key]; ok {
		if v, okb := value.(bool); okb {
			checkboxField.SetChecked(v)
		} else {
			log.Errorf("Unable to convert interface to bool. key: %s", field.Key)
		}
	}

	return fieldReturnType{field, checkboxField}
}

func getFields(configFields []config.FormFieldType) []fieldReturnType {
	log.Infof("getFields: %d", len(configFields))

	items := []fieldReturnType{}
	for _, field := range configFields {
		switch field.Type {
		case "select":
			items = append(items, getSelectField(field))
		case "text":
			items = append(items, getTextField(field))
		case "int":
			items = append(items, getIntField(field))
		case "checkbox":
			items = append(items, getCheckboxField(field))
		}

		if field.Children != nil {
			fieldValue, ok := state.Configuration[field.Key]
			if ok {
				childrenKey, err := utils.InterfaceToString(fieldValue)
				if err != nil {
					log.Errorf("Unable to convert interface to string. Error: %s, key: %s", err, field.Key)
					continue
				}

				children, ok := field.Children[childrenKey]
				if ok {
					items = append(items, getFields(children)...)
				}
			}
		}
	}
	return items
}
