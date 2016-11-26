package app

import (
	"github.com/murlokswarm/log"
	"github.com/murlokswarm/markup"
)

var (
	// OnLaunch is a handler which (if set) is called when the app is
	// initialized and ready.
	// The main window should be created here.
	OnLaunch func()

	// OnFocus is a handler which (if set) is called when the app became
	// focused.
	OnFocus func()

	// OnBlur is a handler which (if set) is called when the app lost the
	// focus.
	OnBlur func()

	// OnReopen is a handler which (if set) is called when the app is reopened.
	// Eg. when the dock icon is clicked.
	OnReopen func(hasVisibleWindow bool)

	// OnFileOpen is a handler which (if set) is called when a file is targeted
	// to be opened with the app.
	OnFileOpen func(filename string)

	// OnTerminate is a handler which (if set) is called when the app is
	// requested to terminates. Return false cancels the termination request.
	OnTerminate func() bool

	// OnFinalize is a handler which (if set) is called when the app is about
	// to be terminated.
	// It should be used to perform any final cleanup before the application
	// terminates.
	OnFinalize func()
)

// Run runs the app.
func Run() {
	driver.Run()
}

// Render renders a component.
func Render(c Componer) {
	var ctx Contexter
	var elems []*markup.Element
	var err error

	if ctx, err = Context(c); err != nil {
		log.Panic(err)
	}

	if elems, err = markup.Sync(c); err != nil {
		log.Error(err)
		return
	}

	for _, elem := range elems {
		ctx.Render(elem)
	}
}

// Menu returns the app menu context.
func Menu() Contexter {
	return driver.AppMenu()
}

// Dock returns the dock context.
func Dock() Contexter {
	return driver.Dock()
}
