package main

import (
	_ "embed"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"io/fs"
	"io/ioutil"
)
//go:embed "icon.png"
var applicationIcon []byte
const applicationId = "com.lagraff.scribal"
const filePath = "scribal.txt"

func main() {
	application := app.NewWithID(applicationId)

	window := application.NewWindow("Scribal")
	window.SetIcon(fyne.NewStaticResource("icon", applicationIcon))
	window.Resize(fyne.NewSize(320, 569))


	text := binding.NewString()
	entry := widget.NewMultiLineEntry()
	entry.Bind(text)
	entry.Validator = nil // Turn off validation
	entry.SetPlaceHolder("Your thoughts here")
	entry.OnChanged = func(s string) {
		ioutil.WriteFile(filePath, []byte(s), fs.ModePerm)
	}

	fileBytes, err := ioutil.ReadFile(filePath)
	if err == nil {
		text.Set(string(fileBytes))
	}

	menu := fyne.NewMenu("File", fyne.NewMenuItem("Save", func() {}))
	mainMenu := fyne.NewMainMenu(menu)
	window.SetMainMenu(mainMenu)
	content := container.New(layout.NewMaxLayout(), entry)

	window.SetContent(content)
	window.ShowAndRun()
}

