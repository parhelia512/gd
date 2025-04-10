// Package ConfirmationDialog provides methods for working with ConfirmationDialog object instances.
package ConfirmationDialog

import "unsafe"
import "reflect"
import "slices"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/internal/gdclass"
import "graphics.gd/variant"
import "graphics.gd/classdb/AcceptDialog"
import "graphics.gd/classdb/Node"
import "graphics.gd/classdb/Viewport"
import "graphics.gd/classdb/Window"
import "graphics.gd/variant/Array"
import "graphics.gd/variant/Callable"
import "graphics.gd/variant/Dictionary"
import "graphics.gd/variant/Error"
import "graphics.gd/variant/Float"
import "graphics.gd/variant/Object"
import "graphics.gd/variant/Packed"
import "graphics.gd/variant/Path"
import "graphics.gd/variant/RID"
import "graphics.gd/variant/RefCounted"
import "graphics.gd/variant/String"

var _ Object.ID
var _ RefCounted.Instance
var _ unsafe.Pointer
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Cycle
var _ = Array.Nil
var _ variant.Any
var _ Callable.Function
var _ Dictionary.Any
var _ RID.Any
var _ String.Readable
var _ Path.ToNode
var _ Packed.Bytes
var _ Error.Code
var _ Float.X
var _ = slices.Delete[[]struct{}, struct{}]

/*
A dialog used for confirmation of actions. This window is similar to [AcceptDialog], but pressing its Cancel button can have a different outcome from pressing the OK button. The order of the two buttons varies depending on the host OS.
To get cancel action, you can use:
[codeblocks]
[gdscript]
get_cancel_button().pressed.connect(_on_canceled)
[/gdscript]
[csharp]
GetCancelButton().Pressed += OnCanceled;
[/csharp]
[/codeblocks]
*/
type Instance [1]gdclass.ConfirmationDialog

// Nil is a nil/null instance of the class. Equivalent to the zero value.
var Nil Instance

type Any interface {
	gd.IsClass
	AsConfirmationDialog() Instance
}

/*
Returns the cancel button.
[b]Warning:[/b] This is a required internal node, removing and freeing it may cause a crash. If you wish to hide it or any of its children, use their [member CanvasItem.visible] property.
*/
func (self Instance) GetCancelButton() [1]gdclass.Button { //gd:ConfirmationDialog.get_cancel_button
	return [1]gdclass.Button(Advanced(self).GetCancelButton())
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.ConfirmationDialog

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("ConfirmationDialog"))
	casted := Instance{*(*gdclass.ConfirmationDialog)(unsafe.Pointer(&object))}
	return casted
}

func (self Instance) CancelButtonText() string {
	return string(class(self).GetCancelButtonText().String())
}

func (self Instance) SetCancelButtonText(value string) {
	class(self).SetCancelButtonText(String.New(value))
}

/*
Returns the cancel button.
[b]Warning:[/b] This is a required internal node, removing and freeing it may cause a crash. If you wish to hide it or any of its children, use their [member CanvasItem.visible] property.
*/
//go:nosplit
func (self class) GetCancelButton() [1]gdclass.Button { //gd:ConfirmationDialog.get_cancel_button
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ConfirmationDialog.Bind_get_cancel_button, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = [1]gdclass.Button{gd.PointerLifetimeBoundTo[gdclass.Button](self.AsObject(), r_ret.Get())}
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetCancelButtonText(text String.Readable) { //gd:ConfirmationDialog.set_cancel_button_text
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalString(text)))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ConfirmationDialog.Bind_set_cancel_button_text, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetCancelButtonText() String.Readable { //gd:ConfirmationDialog.get_cancel_button_text
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ConfirmationDialog.Bind_get_cancel_button_text, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = String.Via(gd.StringProxy{}, pointers.Pack(pointers.New[gd.String](r_ret.Get())))
	frame.Free()
	return ret
}
func (self class) AsConfirmationDialog() Advanced    { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsConfirmationDialog() Instance { return *((*Instance)(unsafe.Pointer(&self))) }
func (self class) AsAcceptDialog() AcceptDialog.Advanced {
	return *((*AcceptDialog.Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsAcceptDialog() AcceptDialog.Instance {
	return *((*AcceptDialog.Instance)(unsafe.Pointer(&self)))
}
func (self class) AsWindow() Window.Advanced    { return *((*Window.Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsWindow() Window.Instance { return *((*Window.Instance)(unsafe.Pointer(&self))) }
func (self class) AsViewport() Viewport.Advanced {
	return *((*Viewport.Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsViewport() Viewport.Instance {
	return *((*Viewport.Instance)(unsafe.Pointer(&self)))
}
func (self class) AsNode() Node.Advanced    { return *((*Node.Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsNode() Node.Instance { return *((*Node.Instance)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(AcceptDialog.Advanced(self.AsAcceptDialog()), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(AcceptDialog.Instance(self.AsAcceptDialog()), name)
	}
}
func init() {
	gdclass.Register("ConfirmationDialog", func(ptr gd.Object) any {
		return [1]gdclass.ConfirmationDialog{*(*gdclass.ConfirmationDialog)(unsafe.Pointer(&ptr))}
	})
}
