package main

import (
	_ "embed"
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"io/fs"
	"io/ioutil"
	"time"
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

	fileBytes, err := ioutil.ReadFile(filePath)
	if err == nil {
		text.Set(string(fileBytes))
	}

	content := container.New(layout.NewMaxLayout(), entry)

	window.SetContent(content)

	done := make(chan bool)
	go autoSave(text, done)
	window.ShowAndRun()
	done <- true
}

func autoSave(text binding.String, cancel <-chan bool) {
	ticker := time.NewTicker(500 * time.Millisecond)
	for {
		select {
		case <-cancel:
			return
		case t := <-ticker.C:
			fmt.Println("Tick at", t)
			content, err := text.Get()
			if err != nil {
				panic(err)
			}
			ioutil.WriteFile(filePath, []byte(content), fs.ModePerm)
		}
	}
}