package startup

import (
	"graphics.gd/classdb"
	NodeClass "graphics.gd/classdb/Node"
	SceneTreeClass "graphics.gd/classdb/SceneTree"
	gd "graphics.gd/internal"
	"graphics.gd/internal/pointers"
	"graphics.gd/variant/Callable"
	"graphics.gd/variant/Float"
)

// goRuntime is injected into the scene tree so that the process function can process
// the frame-based garbage collection routine.
type goRuntime struct {
	classdb.Extension[goRuntime, NodeClass.Instance] `gd:"GoRuntime"`
}

func (goRuntime) Ready() {

}

func (gr goRuntime) AsNode() NodeClass.Instance { return gr.Super().AsNode() }

func (goRuntime) Process(delta Float.X) {
	gd.NewCallable(func() {
		Callable.Cycle()
		pointers.Cycle()
	}).CallDeferred()
}

func init() {
	gd.StartupFunctions = append(gd.StartupFunctions, func() {
		classdb.Register[goRuntime]()
	})
	gd.PostStartupFunctions = append(gd.PostStartupFunctions, func() {
		gd.NewCallable(func() {
			SceneTreeClass.Add(new(goRuntime))
		}).CallDeferred()
	})
}
