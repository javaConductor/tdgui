package main

import (
	"fmt"
	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/theme"
	"fyne.io/fyne/widget"
)

func main() {
	a := app.New()

	w := createMainScreen(a)
	w.Resize(fyne.Size{Height: 200, Width: 500})

	w.ShowAndRun()
}

var mainWindow fyne.Window

func createMainScreen(a fyne.App) fyne.Window {

	mainWindow = createAuthWindow(a)

	return mainWindow
}

func createAuthWindow(a fyne.App) fyne.Window {
	var w fyne.Window
	if UserAuthenticated() {
		w = createLogoutWindow(a)

	} else {
		w = createLoginWindow(a)
	}

	return w
}
func refresh(a fyne.App) {
	mw := mainWindow

	mainWindow = createMainScreen(a)
	mainWindow.Show()
	mw.Close()
}

func createLoginWindow(a fyne.App) fyne.Window {

	win := a.NewWindow("Login")

	userNameWidget := widget.NewEntry()
	userNameWidget.SetPlaceHolder("USERNAME")
	passwordWidget := widget.NewPasswordEntry()
	passwordWidget.SetPlaceHolder("PASSWORD")
	messageArea := widget.NewLabel("")
	loginButton := widget.NewButton("Login!", func() {
		token, err := Authenticate(userNameWidget.Text, passwordWidget.Text)
		if err != nil {
			messageArea.SetText(err.Error())
			return
		}

		SetUserInfo(UserInfo{userNameWidget.Text, token})
		messageArea.SetText(fmt.Sprintf("Recieved token: %s", token))
		dataSetSpecs, err := GetUserDataSets(userNameWidget.Text)
		if err != nil {
			messageArea.SetText(err.Error())
			return
		}

		CreateDataSetSpecsWindow(a, dataSetSpecs).Show()

		refresh(a)
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

	box := widget.NewHBox(
		createToolbar(),
		//widget.NewLabel("Username"),
		userNameWidget,
		//widget.NewLabel("Password"),
		passwordWidget,
		loginButton,
		messageArea,
	)
	win.SetContent(box)
	//win.Resize(fyne.NewSize(200, 200))
	updateButton("")
	return win
}

func createLogoutWindow(a fyne.App) fyne.Window {

	win := a.NewWindow("Logout")

	userLabel := widget.NewLabel("User: " + _userInfo.Username)
	messageArea := widget.NewLabel("")
	logoutButton := widget.NewButton("Logout!", func() {

		SetUserInfo(UserInfo{})
		refresh(a)
	})

	box := widget.NewVBox(
		userLabel,
		logoutButton,
		messageArea,
	)
	win.SetContent(box)
	return win
}

func createToolbar() fyne.CanvasObject {
	return widget.NewToolbar(widget.NewToolbarAction(theme.MailComposeIcon(), func() { fmt.Println("New") }),
		widget.NewToolbarSeparator(),
		widget.NewToolbarSpacer(),
		widget.NewToolbarAction(theme.ContentCutIcon(), func() { fmt.Println("Cut") }),
		widget.NewToolbarAction(theme.ContentCopyIcon(), func() { fmt.Println("Copy") }),
		widget.NewToolbarAction(theme.ContentPasteIcon(), func() { fmt.Println("Paste") }),
	)
}
