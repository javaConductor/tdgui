package main

import (
	"fmt"
	"fyne.io/fyne"
	"fyne.io/fyne/widget"
)

func makeDataSpecWindow(a fyne.App, dataSetSpecs []DataSetSpec) fyne.Window {

	win := a.NewWindow("Data Sets")
	newDataSetButton := widget.NewButton("New Data Set ... ", func() {
		fmt.Printf("NEW !!!")
	})
	ac := widget.NewAccordionContainer()
	for _, ds := range dataSetSpecs {
		ac.Append(makeDataSetTab(a, ds))
	}

	widget.NewVBox(newDataSetButton, ac)
	win.SetContent(widget.NewVBox(newDataSetButton, ac))
	return win

}

func makeDataSetTab(a fyne.App, ds DataSetSpec) *widget.AccordionItem {
	ai := widget.NewAccordionItem(ds.Name, createDataSetView(a, ds))
	return ai
}

func createDataSetView(a fyne.App, ds DataSetSpec) fyne.CanvasObject {

	view := widget.NewVBox()
	for _, objSpec := range ds.ObjectSpecList {
		nameLabel := widget.NewEntry()
		nameLabel.Text = objSpec.Name

		editButton := widget.NewButton("Edit ...", func() {
			//editObjectSpec(objSpec)
			fmt.Printf("Edit Object Spec: %s ", objSpec.Name)
		})

		generateButton := widget.NewButton("Generate ...", func() {
			//editObjectSpec(objSpec)
			fmt.Printf("Generate Object Spec: %s ", objSpec.Name)
		})

		generateCount := widget.NewEntry()
		endLabel := widget.NewLabel("objects")

		objSpecComponent := widget.NewHBox(nameLabel, editButton, generateButton, generateCount, endLabel)

		view.Append(objSpecComponent)

	}
	return view

}
