package demopanes

import "github.com/gotk3/gotk3/gtk"

type buttonDemoPane struct {
}

func ButtonDemoPaneNew() DemoPane {
	return &buttonDemoPane{}
}

func (pane *buttonDemoPane) BuildPane() *gtk.Widget {
	//TODO
	return nil
}

func (pane *buttonDemoPane) Destroy() {
	//TODO
}
