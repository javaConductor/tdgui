package main

import (
	"fmt"
	"fyne.io/fyne/dialog"
	"fyne.io/fyne/theme"

	"fyne.io/fyne"
	"fyne.io/fyne/widget"
)

func NewDataSetSpecsWindow() (*DataSetSpecsWindowObject, error) {
	a := fyne.CurrentApp()
	win := a.NewWindow("Data Sets")
	dssWin := new(DataSetSpecsWindowObject)
	dssWin.window = win
	newDataSetButton := widget.NewButton("New Data Set ... ", func() {
		fmt.Printf("NEW !!!")
	})

	refreshButton := widget.NewButtonWithIcon("", theme.ViewRefreshIcon(), func() {
		dd := DataSetSpecsWindow(dssWin)
		err := dd.Refresh()
		if err != nil {
			fmt.Println("Error refreshing: ", err)
			//if err != nil {
			//	messageArea.SetText(err.Error())
			//	return
			//}
		}
		fmt.Println("Refresh DataSetSpecsWindow ")
	})

	topRow := widget.NewHBox(newDataSetButton, refreshButton)
	dssWin.listContainer = widget.NewAccordionContainer()

	//widget.NewVBox(newDataSetButton, ac)
	content := widget.NewVBox(topRow, dssWin.listContainer)
	content.Resize(fyne.NewSize(1000, 25))

	win.SetContent(content)
	DataSetSpecsWindow(dssWin).Refresh()

	return dssWin, nil

}

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
}

func (d *DataSetSpecsWindowObject) Window() fyne.Window {
	return d.window
}

func (d *DataSetSpecsWindowObject) RefreshWithData(dataSetSpecs []DataSetSpec) {
	populate(d.listContainer, dataSetSpecs)

}

func populate(ac *widget.AccordionContainer, dataSetSpecs []DataSetSpec) {
	ac.Items = ac.Items[:0]
	for _, ds := range dataSetSpecs {
		ac.Append(makeDataSetTab(ds))
	}
}

func (d *DataSetSpecsWindowObject) Refresh() error {
	dataSetSpecs, err := GetUserDataSets(GetUserInfo().Username)
	if err != nil {
		return err
	}
	d.RefreshWithData(dataSetSpecs)
	return nil
}

func (d *DataSetSpecsWindowObject) DataSetSpecs() []DataSetSpec {
	return d.dataSetSpecs
}

func makeDataSetTab(ds DataSetSpec) *widget.AccordionItem {
	ai := widget.NewAccordionItem(ds.Name, createDataSetView(ds))
	return ai
}

func createObjectSpecElement(dssName string, objSpec ObjectSpec) fyne.CanvasObject {
	nameLabel := widget.NewEntry()
	nameLabel.Text = objSpec.Name

	nameLabel.Resize(fyne.NewSize(200, 25))

	editButton := widget.NewButton("Edit ...", func() {
		//editObjectSpec(objSpec)
		fmt.Printf("\nEdit Object Spec: %s ", objSpec.Name)
		win, err := NewObjectSetSpecWindow("", objSpec)
		if err != nil {
			dialog.NewError(err, nil)
		}
		(*win.window).Show()
	})

	generateButton := widget.NewButton("Generate ...", func() {
		//editObjectSpec(objSpec)
		fmt.Printf("\nGenerate Object Spec: %s ", objSpec.Name)
	})

	generateCount := widget.NewEntry()
	endLabel := widget.NewLabel("objects")

	component := widget.NewHBox(nameLabel, editButton, generateButton, generateCount, endLabel)
	component.Resize(fyne.NewSize(1000, 25))
	return component
}

func createDataSetView(ds DataSetSpec) fyne.CanvasObject {

	view := widget.NewVBox()
	for _, objSpec := range ds.ObjectSpecList {
		view.Append(createObjectSpecElement(ds.Name, objSpec))
	}
	return view
}
