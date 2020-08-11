package main

import (
	"fmt"
	"fyne.io/fyne/layout"

	"fyne.io/fyne"
	"fyne.io/fyne/widget"
)

func NewObjectSetSpecWindow(dssName string, objectSpec ObjectSpec) (*ObjectSpecWindowObject, error) {
	a := fyne.CurrentApp()
	win := a.NewWindow(dssName)
	osWin := new(ObjectSpecWindowObject)
	osWin.window = &win

	newFieldButton := widget.NewButton("New Field ... ", func() {
		fmt.Printf("NEW FIELD!!!")
	})

	topRow := widget.NewHBox(newFieldButton)
	osWin.listContainer = widget.NewAccordionContainer()

	content := widget.NewVBox(topRow, osWin.listContainer)

	win.SetContent(content)
	ObjectSpecWindow(osWin).populate(osWin.listContainer, objectSpec.Fields)

	return osWin, nil

}

// ObjectSpecWindow ...
type ObjectSpecWindow interface {
	GetObjectSpec() ObjectSpec
	Window() *fyne.Window
	Close()
	populate(ac *widget.AccordionContainer, fieldSpecs []FieldSpec)
	//populateFieldConstraints(ac *widget.AccordionContainer, fieldSpecs []FieldSpec)

}

// ObjectSpecWindowObject ...
type ObjectSpecWindowObject struct {
	name          string
	fieldList     []FieldSpec
	window        *fyne.Window
	listContainer *widget.AccordionContainer
	dataSetName   string
}

func (osWin *ObjectSpecWindowObject) GetObjectSpec() ObjectSpec {
	return ObjectSpec{Name: osWin.name, Fields: osWin.fieldList}
}

func (osWin *ObjectSpecWindowObject) Close() {
	(*osWin.window).Close()
}

func (osWin *ObjectSpecWindowObject) Window() *fyne.Window {
	return osWin.window
}

func (osWin *ObjectSpecWindowObject) populate(ac *widget.AccordionContainer, fieldSpecs []FieldSpec) {
	ac.Items = ac.Items[:0]
	for _, fld := range fieldSpecs {
		ai := &widget.AccordionItem{Title: fld.Name,
			Detail: createFieldEditorView(fld),
		}
		ac.Append(ai)
	}
}

func createFieldEditorView(fld FieldSpec) fyne.CanvasObject {

	//view := widget.NewVBox()
	//for _, objSpec := range ds.ObjectSpecList {
	//	view.Append(createObjectSpecElement(objSpec))
	//}
	//return view

	l := &widget.Label{Text: fld.Name + " - " + fld.Type}
	return l
}

type ConstraintName struct {
	name        string
	displayName string
}

func createConstraintCheckBoxes(possibleConstraints []ConstraintName, fieldConstraints map[string]*Constraint) fyne.CanvasObject {
	var checkBoxes []fyne.CanvasObject
	for _, cn := range possibleConstraints {
		check := widget.NewCheck(cn.displayName, func(on bool) {
			fmt.Println("checked", on)
		})
		check.SetChecked(fieldConstraints[cn.name] != nil)
		checkBoxes = append(checkBoxes, check)
	}

	return fyne.NewContainerWithLayout(layout.NewVBoxLayout(),
		checkBoxes...)
}
