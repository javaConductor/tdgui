package main

import (
	"fmt"
	"fyne.io/fyne/theme"
	"sort"

	"fyne.io/fyne/layout"

	"fyne.io/fyne"
	"fyne.io/fyne/widget"
)

// FieldAndConstraints ...
type FieldAndConstraints struct {
	fieldSpec   FieldSpec
	constraints map[string]*Constraint
}

// FieldSpecEditorObject ...
type FieldSpecEditorObject struct {
	FieldSpec
	chUpdate           chan<- FieldAndConstraints
	constraints        map[string]*Constraint
	constraintsChannel chan *Constraint
	window             *fyne.Window
	widget             *fyne.CanvasObject
	key                string
}

// FieldSpecEditor ...
type FieldSpecEditor interface {
	Widget() *fyne.CanvasObject
	createView() error
	Key() string
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

	var constraints = make(map[string]FieldTypeConstraints, len(typeConstraints))
	for constraintName, required := range typeConstraints {
		ftc := FieldTypeConstraints{
			Name:        constraintName,
			DisplayName: typeConstraintDisplayNames[constraintName],
			Required:    required}
		constraints[constraintName] = ftc
	}
	return constraints, nil
}

// Widget ...
func (fse *FieldSpecEditorObject) Widget() *fyne.CanvasObject {
	return fse.widget
}

func (fse *FieldSpecEditorObject) Key() string {
	return fse.key
}

func (fse *FieldSpecEditorObject) createView() error {

	content, err := fse.createContent()
	if err != nil {
		return err
	}
	fse.widget = content

	go func() {
		select {

		/// Constraint Editors send updates here to be added to field
		case constraint := <-fse.constraintsChannel:
			fse.constraints[constraint.Name] = constraint
		}
	}()

	return nil
}

func (fse *FieldSpecEditorObject) createContent() (*fyne.CanvasObject, error) {
	typeChanged := func(s string) {
		///
		fse.FieldSpec.Type = s

		newContent, err := fse.createContent()
		if err != nil {
			//TODO log something
			fmt.Println("ERROR: Field", fse.FieldSpec.Name, s, err)
			return
		}

		(*fse.window).SetContent(*newContent)
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

	saveButton := widget.NewButtonWithIcon("", theme.DocumentSaveIcon(), func() {
		fmt.Printf("Saving !!!")
		fse.chUpdate <- FieldAndConstraints{
			fieldSpec:   fse.FieldSpec,
			constraints: fse.constraints,
		}
	})

	topRow := fyne.NewContainerWithLayout(
		layout.NewBorderLayout(nil, nil, fieldName, typeSelect),
		fieldName, typeSelect, saveButton,
	)

	constraintsEditor := fse.createConstraintsEditor(possibleConstraints)
	content := widget.NewVBox(topRow, *constraintsEditor)
	fc := fyne.CanvasObject(content)
	return &fc, nil
}

func (fse *FieldSpecEditorObject) createConstraintsEditor(possibleConstraints map[string]FieldTypeConstraints,
) *fyne.CanvasObject {

	/// create the check boxes
	checkboxes := fse.createConstraintCheckBoxes(possibleConstraints)

	//then each constraint editor to the right of the checkboxes
	constraintEditors := make([]fyne.CanvasObject, 0, len(possibleConstraints))
	constraintEditors = append(constraintEditors, checkboxes)
	for name := range possibleConstraints {
		constraintEditor := createConstraintEditor(name)
		constraintEditors = append(constraintEditors, constraintEditor)
	}
	constraintEditorBox := fyne.CanvasObject(widget.NewVBox(constraintEditors...))
	return &constraintEditorBox

}

func (fse *FieldSpecEditorObject) createConstraintCheckBoxes(possibleConstraints map[string]FieldTypeConstraints) fyne.CanvasObject {
	var checkBoxes []fyne.CanvasObject
	for cn, ftc := range possibleConstraints {
		check := widget.NewCheck(ftc.DisplayName, func(on bool) {
			fmt.Println(ftc.DisplayName, on)
		})
		check.SetChecked(fse.constraints[cn] != nil || ftc.Required)
		checkBoxes = append(checkBoxes, check)
	}

	return fyne.NewContainerWithLayout(layout.NewVBoxLayout(),
		checkBoxes...)
}

func createConstraintEditor(name string) fyne.CanvasObject {
	return widget.NewLabel(name)
}

// NewFieldSpecEditor ...
func NewFieldSpecEditor(
	objectSpecName string,
	fieldSpec FieldSpec,
	chUpdate chan<- FieldAndConstraints) (*FieldSpecEditor, error) {
	constraintsChannel := make(chan *Constraint)

	fieldSpecEditor := FieldSpecEditorObject{
		chUpdate:           chUpdate,
		constraintsChannel: constraintsChannel,
		FieldSpec:          fieldSpec,
	}
	err := fieldSpecEditor.createView()
	if err != nil {
		return nil, err
	}
	var f = FieldSpecEditor(&fieldSpecEditor)
	return &f, nil
}
