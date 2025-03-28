// Package OpenXRAnalogThresholdModifier provides methods for working with OpenXRAnalogThresholdModifier object instances.
package OpenXRAnalogThresholdModifier

import "unsafe"
import "reflect"
import "slices"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/internal/gdclass"
import "graphics.gd/variant"
import "graphics.gd/classdb/OpenXRActionBindingModifier"
import "graphics.gd/classdb/OpenXRBindingModifier"
import "graphics.gd/classdb/Resource"
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
The analog threshold binding modifier can modify a float input to a boolean input with specified thresholds.
See [url=https://registry.khronos.org/OpenXR/specs/1.1/html/xrspec.html#XR_VALVE_analog_threshold]XR_VALVE_analog_threshold[/url] for in-depth details.
*/
type Instance [1]gdclass.OpenXRAnalogThresholdModifier

// Nil is a nil/null instance of the class. Equivalent to the zero value.
var Nil Instance

type Any interface {
	gd.IsClass
	AsOpenXRAnalogThresholdModifier() Instance
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.OpenXRAnalogThresholdModifier

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("OpenXRAnalogThresholdModifier"))
	casted := Instance{*(*gdclass.OpenXRAnalogThresholdModifier)(unsafe.Pointer(&object))}
	casted.AsRefCounted()[0].Reference()
	return casted
}

func (self Instance) OnThreshold() Float.X {
	return Float.X(Float.X(class(self).GetOnThreshold()))
}

func (self Instance) SetOnThreshold(value Float.X) {
	class(self).SetOnThreshold(float64(value))
}

func (self Instance) OffThreshold() Float.X {
	return Float.X(Float.X(class(self).GetOffThreshold()))
}

func (self Instance) SetOffThreshold(value Float.X) {
	class(self).SetOffThreshold(float64(value))
}

func (self Instance) OnHaptic() [1]gdclass.OpenXRHapticBase {
	return [1]gdclass.OpenXRHapticBase(class(self).GetOnHaptic())
}

func (self Instance) SetOnHaptic(value [1]gdclass.OpenXRHapticBase) {
	class(self).SetOnHaptic(value)
}

func (self Instance) OffHaptic() [1]gdclass.OpenXRHapticBase {
	return [1]gdclass.OpenXRHapticBase(class(self).GetOffHaptic())
}

func (self Instance) SetOffHaptic(value [1]gdclass.OpenXRHapticBase) {
	class(self).SetOffHaptic(value)
}

//go:nosplit
func (self class) SetOnThreshold(on_threshold float64) { //gd:OpenXRAnalogThresholdModifier.set_on_threshold
	var frame = callframe.New()
	callframe.Arg(frame, on_threshold)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.OpenXRAnalogThresholdModifier.Bind_set_on_threshold, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetOnThreshold() float64 { //gd:OpenXRAnalogThresholdModifier.get_on_threshold
	var frame = callframe.New()
	var r_ret = callframe.Ret[float64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.OpenXRAnalogThresholdModifier.Bind_get_on_threshold, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetOffThreshold(off_threshold float64) { //gd:OpenXRAnalogThresholdModifier.set_off_threshold
	var frame = callframe.New()
	callframe.Arg(frame, off_threshold)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.OpenXRAnalogThresholdModifier.Bind_set_off_threshold, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetOffThreshold() float64 { //gd:OpenXRAnalogThresholdModifier.get_off_threshold
	var frame = callframe.New()
	var r_ret = callframe.Ret[float64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.OpenXRAnalogThresholdModifier.Bind_get_off_threshold, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetOnHaptic(haptic [1]gdclass.OpenXRHapticBase) { //gd:OpenXRAnalogThresholdModifier.set_on_haptic
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(haptic[0])[0])
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.OpenXRAnalogThresholdModifier.Bind_set_on_haptic, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetOnHaptic() [1]gdclass.OpenXRHapticBase { //gd:OpenXRAnalogThresholdModifier.get_on_haptic
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.OpenXRAnalogThresholdModifier.Bind_get_on_haptic, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = [1]gdclass.OpenXRHapticBase{gd.PointerWithOwnershipTransferredToGo[gdclass.OpenXRHapticBase](r_ret.Get())}
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetOffHaptic(haptic [1]gdclass.OpenXRHapticBase) { //gd:OpenXRAnalogThresholdModifier.set_off_haptic
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(haptic[0])[0])
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.OpenXRAnalogThresholdModifier.Bind_set_off_haptic, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetOffHaptic() [1]gdclass.OpenXRHapticBase { //gd:OpenXRAnalogThresholdModifier.get_off_haptic
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.OpenXRAnalogThresholdModifier.Bind_get_off_haptic, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = [1]gdclass.OpenXRHapticBase{gd.PointerWithOwnershipTransferredToGo[gdclass.OpenXRHapticBase](r_ret.Get())}
	frame.Free()
	return ret
}
func (self class) AsOpenXRAnalogThresholdModifier() Advanced {
	return *((*Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsOpenXRAnalogThresholdModifier() Instance {
	return *((*Instance)(unsafe.Pointer(&self)))
}
func (self class) AsOpenXRActionBindingModifier() OpenXRActionBindingModifier.Advanced {
	return *((*OpenXRActionBindingModifier.Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsOpenXRActionBindingModifier() OpenXRActionBindingModifier.Instance {
	return *((*OpenXRActionBindingModifier.Instance)(unsafe.Pointer(&self)))
}
func (self class) AsOpenXRBindingModifier() OpenXRBindingModifier.Advanced {
	return *((*OpenXRBindingModifier.Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsOpenXRBindingModifier() OpenXRBindingModifier.Instance {
	return *((*OpenXRBindingModifier.Instance)(unsafe.Pointer(&self)))
}
func (self class) AsResource() Resource.Advanced {
	return *((*Resource.Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsResource() Resource.Instance {
	return *((*Resource.Instance)(unsafe.Pointer(&self)))
}
func (self class) AsRefCounted() [1]gd.RefCounted {
	return *((*[1]gd.RefCounted)(unsafe.Pointer(&self)))
}
func (self Instance) AsRefCounted() [1]gd.RefCounted {
	return *((*[1]gd.RefCounted)(unsafe.Pointer(&self)))
}

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(OpenXRActionBindingModifier.Advanced(self.AsOpenXRActionBindingModifier()), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(OpenXRActionBindingModifier.Instance(self.AsOpenXRActionBindingModifier()), name)
	}
}
func init() {
	gdclass.Register("OpenXRAnalogThresholdModifier", func(ptr gd.Object) any {
		return [1]gdclass.OpenXRAnalogThresholdModifier{*(*gdclass.OpenXRAnalogThresholdModifier)(unsafe.Pointer(&ptr))}
	})
}
