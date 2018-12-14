package demopanes

import "github.com/gotk3/gotk3/gtk"

//DemoPane is a thing that is able to build a demo widget and destroy it
//after the user is done viewing the demo
type DemoPane interface {
	BuildPane() *gtk.Widget
	Destroy()
}
