package main

import (
	"fmt"

	"fyne.io/fyne"
	"fyne.io/fyne/widget"
)

// ObjectSpecEditor ...
type ObjectSpecEditor interface {
	GetObjectSpec() *ObjectSpec
	Close()
	populateAll(ac *widget.AccordionContainer, objectSpecName string, fieldSpecs []FieldSpec) error
	//populateFieldConstraints(ac *widget.AccordionContainer, fieldSpecs []FieldSpec)
	Window() *fyne.Window
	createView() error
	Key() string
}

// ObjectSpecEditorObject ...
type ObjectSpecEditorObject struct {
	ObjectSpec
	key              string //TODO maybe get rid of this and just use the Func
	chUpdate         <-chan ObjectSpec
	chUpdateField    chan FieldAndConstraints
	dataSetName      string
	listContainer    *widget.AccordionContainer
	window           *fyne.Window
	fieldSpecEditors map[string]*FieldSpecEditor
}

func (osEditor *ObjectSpecEditorObject) createView() error {
	a := fyne.CurrentApp()
	win := a.NewWindow(osEditor.Name)
	osEditor.window = &win

	content, err := osEditor.createContent()
	if err != nil {
		return err
	}
	win.SetContent(*content)

	go func() {
		select {
		case fieldSpecAndConstraints := <-osEditor.chUpdateField:

			for n, fld := range osEditor.Fields {
				if fld.Name == fieldSpecAndConstraints.fieldSpec.Name {
					osEditor.Fields[n] = fieldSpecAndConstraints.fieldSpec
					osEditor.ObjectSpec.Constraints[fld.Name] = fieldSpecAndConstraints.constraints
					_ = osEditor.populateAll(osEditor.listContainer, osEditor.Name, osEditor.Fields)
				}
			}
		}
	}()
	return nil
}

func (osEditor *ObjectSpecEditorObject) createContent() (*fyne.CanvasObject, error) {
	newFieldButton := widget.NewButton("New Field ... ", func() {
		fmt.Printf("NEW FIELD!!!")
	})

	topRow := widget.NewHBox(newFieldButton)
	osEditor.listContainer = widget.NewAccordionContainer()

	err := osEditor.populateAll(osEditor.listContainer, osEditor.ObjectSpec.Name, osEditor.ObjectSpec.Fields)
	if err != nil {
		return nil, err
	}
	content := widget.NewVBox(topRow, osEditor.listContainer)
	fc := fyne.CanvasObject(content)
	return &fc, nil
}

func (osEditor *ObjectSpecEditorObject) populateAll(ac *widget.AccordionContainer, objectSpecName string, fieldSpecs []FieldSpec) error {
	ac.Items = ac.Items[:0]
	for _, fld := range fieldSpecs {
		ai, err := osEditor.createAccordionItemForField(objectSpecName, fld)
		if err != nil {
			return err
		}
		ac.Append(ai)
	}
	return nil
}

func (osEditor *ObjectSpecEditorObject) createAccordionItemForField(objectSpecName string, fld FieldSpec) (*widget.AccordionItem, error) {
	detail, err := osEditor.createFieldEditorView(objectSpecName, fld)
	if err != nil {
		return nil, err
	}

	w := (*detail).Widget()
	item := widget.AccordionItem{Title: fld.Name, Detail: *w}
	return &item, nil
}

func (osEditor *ObjectSpecEditorObject) createFieldEditorView(objectSpecName string, fld FieldSpec) (*FieldSpecEditor, error) {
	edt, err := NewFieldSpecEditor(objectSpecName, fld, osEditor.chUpdateField)
	if err != nil {
		return nil, err
	}
	osEditor.fieldSpecEditors[fld.Name] = edt
	return edt, nil
}

func (osEditor *ObjectSpecEditorObject) GetObjectSpec() *ObjectSpec {
	spec := (*osEditor).ObjectSpec
	return &spec
}

func (osEditor *ObjectSpecEditorObject) Close() {
	(*osEditor.window).Hide()
}

func (osEditor *ObjectSpecEditorObject) Window() *fyne.Window {
	return osEditor.window
}

func (osEditor *ObjectSpecEditorObject) Key() string {
	return osEditor.key
}

func NewObjectSetSpecEditor(dssName string, objectSpec ObjectSpec, chUpdate chan<- *ObjectSpec) (*ObjectSpecEditor, error) {
	osEditor := ObjectSpecEditorObject{
		ObjectSpec:       objectSpec,
		chUpdate:         make(<-chan ObjectSpec),
		dataSetName:      dssName,
		key:              dssName + "." + objectSpec.Name,
		fieldSpecEditors: map[string]*FieldSpecEditor{},
	}

	err := osEditor.createView()
	if err != nil {
		return nil, err
	}
	var f = ObjectSpecEditor(&osEditor)
	return &f, nil
}
