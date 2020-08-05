package main

import (
	"fmt"

	"fyne.io/fyne"
	"fyne.io/fyne/widget"
)

func NewObjectSetSpecWindow(dss string, objectSpec ObjectSpec) (*ObjectSpecWindowObject, error) {
	a := fyne.CurrentApp()
	win := a.NewWindow("Object Type")
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

//func makeFieldTab(fld FieldSpec) *widget.AccordionItem {
//	ai := widget.NewAccordionItem(fld.Name, createFieldEditorView(fld))
//	return ai
//}

//
//func createObjectSpecElement(objSpec ObjectSpec) fyne.CanvasObject {
//	nameLabel := widget.NewEntry()
//	nameLabel.Text = objSpec.Name
//
//	nameLabel.Resize(fyne.NewSize(200, 25))
//
//	editButton := widget.NewButton("Edit ...", func() {
//		//editObjectSpec(objSpec)
//		fmt.Printf("\nEdit Object Spec: %s ", objSpec.Name)
//	})
//
//	generateButton := widget.NewButton("Generate ...", func() {
//		//editObjectSpec(objSpec)
//		fmt.Printf("\nGenerate Object Spec: %s ", objSpec.Name)
//	})
//
//	generateCount := widget.NewEntry()
//	endLabel := widget.NewLabel("objects")
//
//	component := widget.NewHBox(nameLabel, editButton, generateButton, generateCount, endLabel)
//	component.Resize(fyne.NewSize(1000, 25))
//	return component
//}

func createFieldEditorView(fld FieldSpec) fyne.CanvasObject {

	//view := widget.NewVBox()
	//for _, objSpec := range ds.ObjectSpecList {
	//	view.Append(createObjectSpecElement(objSpec))
	//}
	//return view
	l := &widget.Label{Text: fld.Name + " - " + fld.Type}
	return l
}
