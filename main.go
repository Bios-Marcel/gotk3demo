package main

import (
	"log"

	"github.com/Bios-Marcel/gotk3demo/demopanes"
	"github.com/gotk3/gotk3/glib"
	"github.com/gotk3/gotk3/gtk"
)

var currentDemoPane demopanes.DemoPane

func main() {
	app, gtkErr := gtk.ApplicationNew("gotk3.demo.app", glib.APPLICATION_FLAGS_NONE)
	if gtkErr != nil {
		log.Fatalf("Error creating new gtk3 application: %s.", gtkErr.Error())
	}

	app.Connect("activate", func() {
		buildApplicationWindow()
	})

	app.Run(nil)
}

func buildApplicationWindow() {
	window, gtkErr := gtk.ApplicationWindowNew(app)
	if gtkErr != nil {
		log.Fatalf("Error creating application windows: %s", gtkErr.Error())
	}

	featureList, gtkErr := gtk.ListBoxNew()
	if gtkErr != nil {
		log.Fatalf("Error creating feature selection list: %s", gtkErr.Error())
	}

	/*featureList.Connect("selected-rows-changed", func() {
		currentDemoPane.Destroy()
		layout.Remove(currentDemoPane)

		row := featureList.GetSelectedRow()

		TODO BUild new demo pane

		layout.PackEnd(newDemoPane, true, true, 0)
	})*/

	layout, gtkErr := gtk.BoxNew(gtk.ORIENTATION_HORIZONTAL, 5)
	if gtkErr != nil {
		log.Fatalf("Error creating layout: %s", gtkErr.Error())
	}

	layout.PackStart(featureList, false, false, 0)

	window.Add(layout)
	window.ShowAll()
}
