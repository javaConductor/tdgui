package main

import (
	"fmt"
	"fyne.io/fyne"
	"fyne.io/fyne/dialog"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/theme"
	"fyne.io/fyne/widget"
)

// DataSetSpecsWindow ...
type DataSetSpecsWindow interface {
	RefreshWithData(dataSetSpecs []DataSetSpec)
	Refresh() error
	DataSetSpecs() []DataSetSpec
	Window() fyne.Window
}

// DataSetSpecsWindowObject ...
type DataSetSpecsWindowObject struct {
	dataSetSpecs  []DataSetSpec
	listContainer *widget.AccordionContainer
	window        fyne.Window

	objectSpecEditors map[string]*ObjectSpecEditor
}

func (d *DataSetSpecsWindowObject) Window() fyne.Window {
	return d.window
}

func (d *DataSetSpecsWindowObject) RefreshWithData(dataSetSpecs []DataSetSpec) {
	d.populate(d.listContainer, dataSetSpecs)
}

func (d *DataSetSpecsWindowObject) populate(ac *widget.AccordionContainer, dataSetSpecs []DataSetSpec) {
	//ac.Items = ac.Items[:0]

	for _, ds := range dataSetSpecs {
		ac.Append(d.makeDataSetTab(ds))
	}
}

func (d *DataSetSpecsWindowObject) Refresh() error {
	dataSetSpecs, err := GetUserDataSets(GetUserInfo().Username)
	if err != nil {
		return err
	}
	d.dataSetSpecs = dataSetSpecs
	d.RefreshWithData(dataSetSpecs)
	return nil
}

func (d *DataSetSpecsWindowObject) DataSetSpecs() []DataSetSpec {
	return d.dataSetSpecs
}

func (d *DataSetSpecsWindowObject) makeDataSetTab(ds DataSetSpec) *widget.AccordionItem {
	ai := widget.NewAccordionItem(ds.Name, d.createDataSetView(ds))
	return ai
}

func (d *DataSetSpecsWindowObject) createObjectSpecElement(dssName string, objSpec ObjectSpec) fyne.CanvasObject {
	nameLabel := widget.NewEntry()
	nameLabel.Text = objSpec.Name
	nameLabel.Resize(fyne.NewSize(200, 25))

	editButton := widget.NewButton("Edit ...", func() {
		//editObjectSpec(objSpec)
		fmt.Printf("\nEdit Object Spec: %s ", objSpec.Name)

		win := d.objectSpecEditors[dssName+"."+objSpec.Name]
		if win != nil {
			(*((*win).Window())).Show()
			return
		}
		// NewObjectSetSpecEditor(dssName string, objectSpec ObjectSpec, chUpdate chan<- *ObjectSpec) (*ObjectSpecEditor, error)
		win, err := NewObjectSetSpecEditor(dssName, objSpec, nil)
		if err != nil {
			dialog.NewError(err, nil)
		}
		// add to the list
		d.objectSpecEditors[dssName+"."+objSpec.Name] = win
		(*((*win).Window())).Show()
	})

	generateButton := widget.NewButton("Generate ...", func() {
		//editObjectSpec(objSpec)
		fmt.Printf("\nGenerate Object Spec: %s ", objSpec.Name)
	})

	generateCount := widget.NewEntry()
	endLabel := widget.NewLabel("objects")

	/// put left 3 components in box leave first one
	left := widget.NewHBox(editButton, layout.NewSpacer(), generateButton, generateCount, endLabel)
	/// create border layout with nameLabel: West and leftBox: East

	container := fyne.NewContainerWithLayout(
		layout.NewBorderLayout(nil, nil, nameLabel, left),
		nameLabel, left,
	)

	return container
}

func (d *DataSetSpecsWindowObject) createDataSetView(ds DataSetSpec) fyne.CanvasObject {
	view := widget.NewVBox()
	for _, objSpec := range ds.ObjectSpecList {
		view.Append(d.createObjectSpecElement(ds.Name, objSpec))
	}
	return view
}

func (dssWin *DataSetSpecsWindowObject) createView() error {

	newDataSetButton := widget.NewButtonWithIcon("New Data Set ... ", theme.DocumentCreateIcon(), func() {
		fmt.Printf("NEW !!!")
	})

	refreshButton := widget.NewButtonWithIcon("", theme.ViewRefreshIcon(), func() {
		dd := DataSetSpecsWindow(dssWin)
		err := dd.Refresh()
		if err != nil {
			fmt.Println("Error refreshing: ", err)
		}
		fmt.Println("Refresh DataSetSpecsWindow ")
	})

	topRow := widget.NewHBox(newDataSetButton, refreshButton)
	dssWin.listContainer = widget.NewAccordionContainer()

	//widget.NewVBox(newDataSetButton, ac)
	content := widget.NewVBox(topRow, dssWin.listContainer)
	content.Resize(fyne.NewSize(1000, 25))

	dssWin.window.SetContent(content)
	err := dssWin.Refresh()

	return err
}

func NewDataSetSpecsWindow() (*DataSetSpecsWindowObject, error) {
	a := fyne.CurrentApp()
	win := a.NewWindow("Data Sets")
	dssWin := new(DataSetSpecsWindowObject)
	dssWin.window = win
	dssWin.objectSpecEditors = map[string]*ObjectSpecEditor{}

	err := dssWin.createView()
	if err != nil {
		return nil, err
	}

	err = dssWin.Refresh()
	win.Show()
	if err != nil {
		return nil, err
	}
	return dssWin, nil
}
