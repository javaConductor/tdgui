package main

import (
	"fmt"
	"sort"

	"fyne.io/fyne/layout"

	"fyne.io/fyne"
	"fyne.io/fyne/widget"
)

// FieldAndConstraints ...
type FieldAndConstraints struct {
	fieldSpec   FieldSpec
	constraints map[string]Constraint
}

// FieldSpecEditorObject ...
type FieldSpecEditorObject struct {
	FieldSpec
	chUpdate           chan<- FieldAndConstraints
	constraintsEditor  *fyne.CanvasObject
	constraints        map[string]*Constraint
	constraintsChannel chan map[string]*Constraint
	component          *fyne.CanvasObject
	window             *fyne.Window
}

// FieldSpecEditor ...
type FieldSpecEditor interface {
	Component() *fyne.CanvasObject
	createView() error
	refreshConstraints(map[string]*Constraint)
}

// FieldTypeConstraints ...
type FieldTypeConstraints struct {
	Name        string
	DisplayName string
	Required    bool
}

func constraintsForType(typename string) (map[string]FieldTypeConstraints, error) {

	metadata, err := GetMetadata()
	if err != nil {
		return map[string]FieldTypeConstraints{}, err
	}

	typeConstraints := metadata.TypeInfo.TypeConstraints[typename]
	typeConstraintDisplayNames := metadata.TypeInfo.TypeConstraintDisplayNames[typename]

	var constraints map[string]FieldTypeConstraints
	for constraintName, required := range typeConstraints {
		ftc := FieldTypeConstraints{
			Name:        constraintName,
			DisplayName: typeConstraintDisplayNames[constraintName],
			Required:    required}
		constraints[constraintName] = ftc
	}
	return constraints, nil
}

// NewFieldSpecEditor ...
func NewFieldSpecEditor(
	objectSpecName string,
	fielsSpec FieldSpec,
	chUpdate chan<- FieldAndConstraints) (*FieldSpecEditor, error) {
	constrantsChannel := make(chan map[string]*Constraint)

	fieldSpecEditor := FieldSpecEditorObject{
		chUpdate:           chUpdate,
		constraintsChannel: constrantsChannel,
	}
	fieldSpecEditor.createView()
	var f FieldSpecEditor = FieldSpecEditor(&fieldSpecEditor)
	return &f, nil
}

// Component ...
func (fse *FieldSpecEditorObject) Component() *fyne.CanvasObject {
	return fse.component
}
func (fse *FieldSpecEditorObject) createView() error {

	a := fyne.CurrentApp()
	window := a.NewWindow("Data Sets")
	content, err := fse.createContent()
	if err != nil {
		return err
	}
	window.SetContent(*content)
	fse.window = &window

	go func() {
		select {
		case constraint := <-fse.constraintsChannel:
			for name, v := range constraint {
				fse.constraints[name] = v
			}
		}
	}()

	return nil
}

func (fse *FieldSpecEditorObject) createContent() (*fyne.CanvasObject, error) {
	typeChanged := func(s string) {
		///
		fse.FieldSpec.Type = s
		var w *fyne.Window = fse.window

		newContent, err := fse.createContent()
		if err != nil {
			//TODO log something
			fmt.Println("ERROR: Field", fse.FieldSpec.Name, s, err)
			return
		}
		(*w).SetContent(*newContent)
		fmt.Println("selected type ", s)
	}

	fieldName := widget.NewEntry()
	possibleConstraints, err := constraintsForType(fse.FieldSpec.Type)
	if err != nil {
		return nil, err
	}

	// get the constraint names
	constraintNames := make([]string, 0, len(possibleConstraints))
	for k := range possibleConstraints {
		constraintNames = append(constraintNames, k)
	}
	// sort them
	sort.Strings(constraintNames)

	typeSelect := widget.NewSelect(constraintNames, typeChanged)

	topRow := fyne.NewContainerWithLayout(
		layout.NewBorderLayout(nil, nil, fieldName, typeSelect),
		fieldName, typeSelect,
	)

	constraintsEditor := fse.createConstraintsEditor(possibleConstraints)
	content := widget.NewVBox(topRow, *constraintsEditor)
	fc := fyne.CanvasObject(content)
	return &fc, nil
}

func (fse *FieldSpecEditorObject) refreshConstraints(fldConstraints map[string]*Constraint) {

}

func (fse *FieldSpecEditorObject) createConstraintsEditor(possibleConstraints map[string]FieldTypeConstraints,
) *fyne.CanvasObject {

	/// create the check boxes
	checkboxes := fse.createConstraintCheckBoxes(possibleConstraints, fse.constraints)

	//then each constraint editor to the right of the checkboxes
	constraintEditors := make([]fyne.CanvasObject, len(possibleConstraints))
	constraintEditors = append(constraintEditors, checkboxes)

	for name := range possibleConstraints {

		constraintEditor := createConstraintEditor(name, fse.constraints[name])
		constraintEditors = append(constraintEditors, constraintEditor)
	}
	constraintEditorBox := fyne.CanvasObject(widget.NewVBox(constraintEditors...))
	return &constraintEditorBox

}
func createConstraintEditor(name string, constraint *Constraint) fyne.CanvasObject {
	return widget.NewLabel(name)
}

func (fse *FieldSpecEditorObject) createConstraintCheckBoxes(
	possibleConstraints map[string]FieldTypeConstraints,
	fieldConstraints map[string]*Constraint) fyne.CanvasObject {
	var checkBoxes []fyne.CanvasObject
	for cn, ftc := range possibleConstraints {
		check := widget.NewCheck(ftc.DisplayName, func(on bool) {
			fmt.Println(ftc.DisplayName, on)
		})
		check.SetChecked(fieldConstraints[cn] != nil || ftc.Required)
		checkBoxes = append(checkBoxes, check)
	}

	return fyne.NewContainerWithLayout(layout.NewVBoxLayout(),
		checkBoxes...)
}
