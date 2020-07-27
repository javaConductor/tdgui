package main

import (
	"fmt"
	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/widget"
)

func main() {
	t := Pc()
	fmt.Println("ExpressionConstraint(main):", t)
	a := app.New()
	w := a.NewWindow("Test Data Generator")

	hello := widget.NewLabel("Welcome to TDG!!")
	w.SetContent(widget.NewVBox(
		hello,
		widget.NewButton("Hi!", func() {
			hello.SetText("Welcome :)")
			showLogin(a)
		}),
	))

	w.Resize(fyne.Size{Height: 200, Width: 500})
	w.ShowAndRun()
}

func showLogin(a fyne.App) {

	win := a.NewWindow("Login")

	userNameWidget := widget.NewEntry()
	passwordWidget := widget.NewPasswordEntry()
	messageArea := widget.NewLabel("")
	loginButton := widget.NewButton("Login!", func() {
		token, err := Authenticate(userNameWidget.Text, passwordWidget.Text)
		if err != nil {
			messageArea.SetText( err.Error())
			return
		}


		messageArea.SetText(  fmt.Sprintf("Recieved token: %s", token))
	})

	updateButton := func(s string) {
		if userNameWidget.Text == "" || passwordWidget.Text == "" {
			loginButton.Disable()
			return
		}
		loginButton.Enable()
	}

	userNameWidget.OnChanged = updateButton
	passwordWidget.OnChanged = updateButton

	box := widget.NewVBox(
		widget.NewLabel("Username"),
		userNameWidget,
		widget.NewLabel("Password"),
		passwordWidget,
		loginButton,
		messageArea,
	)
	win.SetContent(box)
	updateButton("")
	win.Resize(fyne.NewSize(200, 200))
	win.Show()
}

