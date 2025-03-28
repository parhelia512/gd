// Package ProceduralSkyMaterial provides methods for working with ProceduralSkyMaterial object instances.
package ProceduralSkyMaterial

import "unsafe"
import "reflect"
import "slices"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/internal/gdclass"
import "graphics.gd/variant"
import "graphics.gd/classdb/Material"
import "graphics.gd/classdb/Resource"
import "graphics.gd/variant/Array"
import "graphics.gd/variant/Callable"
import "graphics.gd/variant/Color"
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
[ProceduralSkyMaterial] provides a way to create an effective background quickly by defining procedural parameters for the sun, the sky and the ground. The sky and ground are defined by a main color, a color at the horizon, and an easing curve to interpolate between them. Suns are described by a position in the sky, a color, and a max angle from the sun at which the easing curve ends. The max angle therefore defines the size of the sun in the sky.
[ProceduralSkyMaterial] supports up to 4 suns, using the color, and energy, direction, and angular distance of the first four [DirectionalLight3D] nodes in the scene. This means that the suns are defined individually by the properties of their corresponding [DirectionalLight3D]s and globally by [member sun_angle_max] and [member sun_curve].
[ProceduralSkyMaterial] uses a lightweight shader to draw the sky and is therefore suited for real-time updates. This makes it a great option for a sky that is simple and computationally cheap, but unrealistic. If you need a more realistic procedural option, use [PhysicalSkyMaterial].
*/
type Instance [1]gdclass.ProceduralSkyMaterial

// Nil is a nil/null instance of the class. Equivalent to the zero value.
var Nil Instance

type Any interface {
	gd.IsClass
	AsProceduralSkyMaterial() Instance
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.ProceduralSkyMaterial

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("ProceduralSkyMaterial"))
	casted := Instance{*(*gdclass.ProceduralSkyMaterial)(unsafe.Pointer(&object))}
	casted.AsRefCounted()[0].Reference()
	return casted
}

func (self Instance) SkyTopColor() Color.RGBA {
	return Color.RGBA(class(self).GetSkyTopColor())
}

func (self Instance) SetSkyTopColor(value Color.RGBA) {
	class(self).SetSkyTopColor(Color.RGBA(value))
}

func (self Instance) SkyHorizonColor() Color.RGBA {
	return Color.RGBA(class(self).GetSkyHorizonColor())
}

func (self Instance) SetSkyHorizonColor(value Color.RGBA) {
	class(self).SetSkyHorizonColor(Color.RGBA(value))
}

func (self Instance) SkyCurve() Float.X {
	return Float.X(Float.X(class(self).GetSkyCurve()))
}

func (self Instance) SetSkyCurve(value Float.X) {
	class(self).SetSkyCurve(float64(value))
}

func (self Instance) SkyEnergyMultiplier() Float.X {
	return Float.X(Float.X(class(self).GetSkyEnergyMultiplier()))
}

func (self Instance) SetSkyEnergyMultiplier(value Float.X) {
	class(self).SetSkyEnergyMultiplier(float64(value))
}

func (self Instance) SkyCover() [1]gdclass.Texture2D {
	return [1]gdclass.Texture2D(class(self).GetSkyCover())
}

func (self Instance) SetSkyCover(value [1]gdclass.Texture2D) {
	class(self).SetSkyCover(value)
}

func (self Instance) SkyCoverModulate() Color.RGBA {
	return Color.RGBA(class(self).GetSkyCoverModulate())
}

func (self Instance) SetSkyCoverModulate(value Color.RGBA) {
	class(self).SetSkyCoverModulate(Color.RGBA(value))
}

func (self Instance) GroundBottomColor() Color.RGBA {
	return Color.RGBA(class(self).GetGroundBottomColor())
}

func (self Instance) SetGroundBottomColor(value Color.RGBA) {
	class(self).SetGroundBottomColor(Color.RGBA(value))
}

func (self Instance) GroundHorizonColor() Color.RGBA {
	return Color.RGBA(class(self).GetGroundHorizonColor())
}

func (self Instance) SetGroundHorizonColor(value Color.RGBA) {
	class(self).SetGroundHorizonColor(Color.RGBA(value))
}

func (self Instance) GroundCurve() Float.X {
	return Float.X(Float.X(class(self).GetGroundCurve()))
}

func (self Instance) SetGroundCurve(value Float.X) {
	class(self).SetGroundCurve(float64(value))
}

func (self Instance) GroundEnergyMultiplier() Float.X {
	return Float.X(Float.X(class(self).GetGroundEnergyMultiplier()))
}

func (self Instance) SetGroundEnergyMultiplier(value Float.X) {
	class(self).SetGroundEnergyMultiplier(float64(value))
}

func (self Instance) SunAngleMax() Float.X {
	return Float.X(Float.X(class(self).GetSunAngleMax()))
}

func (self Instance) SetSunAngleMax(value Float.X) {
	class(self).SetSunAngleMax(float64(value))
}

func (self Instance) SunCurve() Float.X {
	return Float.X(Float.X(class(self).GetSunCurve()))
}

func (self Instance) SetSunCurve(value Float.X) {
	class(self).SetSunCurve(float64(value))
}

func (self Instance) UseDebanding() bool {
	return bool(class(self).GetUseDebanding())
}

func (self Instance) SetUseDebanding(value bool) {
	class(self).SetUseDebanding(value)
}

func (self Instance) EnergyMultiplier() Float.X {
	return Float.X(Float.X(class(self).GetEnergyMultiplier()))
}

func (self Instance) SetEnergyMultiplier(value Float.X) {
	class(self).SetEnergyMultiplier(float64(value))
}

//go:nosplit
func (self class) SetSkyTopColor(color Color.RGBA) { //gd:ProceduralSkyMaterial.set_sky_top_color
	var frame = callframe.New()
	callframe.Arg(frame, color)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ProceduralSkyMaterial.Bind_set_sky_top_color, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetSkyTopColor() Color.RGBA { //gd:ProceduralSkyMaterial.get_sky_top_color
	var frame = callframe.New()
	var r_ret = callframe.Ret[Color.RGBA](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ProceduralSkyMaterial.Bind_get_sky_top_color, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetSkyHorizonColor(color Color.RGBA) { //gd:ProceduralSkyMaterial.set_sky_horizon_color
	var frame = callframe.New()
	callframe.Arg(frame, color)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ProceduralSkyMaterial.Bind_set_sky_horizon_color, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetSkyHorizonColor() Color.RGBA { //gd:ProceduralSkyMaterial.get_sky_horizon_color
	var frame = callframe.New()
	var r_ret = callframe.Ret[Color.RGBA](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ProceduralSkyMaterial.Bind_get_sky_horizon_color, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetSkyCurve(curve float64) { //gd:ProceduralSkyMaterial.set_sky_curve
	var frame = callframe.New()
	callframe.Arg(frame, curve)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ProceduralSkyMaterial.Bind_set_sky_curve, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetSkyCurve() float64 { //gd:ProceduralSkyMaterial.get_sky_curve
	var frame = callframe.New()
	var r_ret = callframe.Ret[float64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ProceduralSkyMaterial.Bind_get_sky_curve, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetSkyEnergyMultiplier(multiplier float64) { //gd:ProceduralSkyMaterial.set_sky_energy_multiplier
	var frame = callframe.New()
	callframe.Arg(frame, multiplier)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ProceduralSkyMaterial.Bind_set_sky_energy_multiplier, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetSkyEnergyMultiplier() float64 { //gd:ProceduralSkyMaterial.get_sky_energy_multiplier
	var frame = callframe.New()
	var r_ret = callframe.Ret[float64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ProceduralSkyMaterial.Bind_get_sky_energy_multiplier, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetSkyCover(sky_cover [1]gdclass.Texture2D) { //gd:ProceduralSkyMaterial.set_sky_cover
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(sky_cover[0])[0])
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ProceduralSkyMaterial.Bind_set_sky_cover, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetSkyCover() [1]gdclass.Texture2D { //gd:ProceduralSkyMaterial.get_sky_cover
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ProceduralSkyMaterial.Bind_get_sky_cover, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = [1]gdclass.Texture2D{gd.PointerWithOwnershipTransferredToGo[gdclass.Texture2D](r_ret.Get())}
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetSkyCoverModulate(color Color.RGBA) { //gd:ProceduralSkyMaterial.set_sky_cover_modulate
	var frame = callframe.New()
	callframe.Arg(frame, color)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ProceduralSkyMaterial.Bind_set_sky_cover_modulate, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetSkyCoverModulate() Color.RGBA { //gd:ProceduralSkyMaterial.get_sky_cover_modulate
	var frame = callframe.New()
	var r_ret = callframe.Ret[Color.RGBA](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ProceduralSkyMaterial.Bind_get_sky_cover_modulate, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetGroundBottomColor(color Color.RGBA) { //gd:ProceduralSkyMaterial.set_ground_bottom_color
	var frame = callframe.New()
	callframe.Arg(frame, color)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ProceduralSkyMaterial.Bind_set_ground_bottom_color, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetGroundBottomColor() Color.RGBA { //gd:ProceduralSkyMaterial.get_ground_bottom_color
	var frame = callframe.New()
	var r_ret = callframe.Ret[Color.RGBA](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ProceduralSkyMaterial.Bind_get_ground_bottom_color, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetGroundHorizonColor(color Color.RGBA) { //gd:ProceduralSkyMaterial.set_ground_horizon_color
	var frame = callframe.New()
	callframe.Arg(frame, color)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ProceduralSkyMaterial.Bind_set_ground_horizon_color, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetGroundHorizonColor() Color.RGBA { //gd:ProceduralSkyMaterial.get_ground_horizon_color
	var frame = callframe.New()
	var r_ret = callframe.Ret[Color.RGBA](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ProceduralSkyMaterial.Bind_get_ground_horizon_color, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetGroundCurve(curve float64) { //gd:ProceduralSkyMaterial.set_ground_curve
	var frame = callframe.New()
	callframe.Arg(frame, curve)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ProceduralSkyMaterial.Bind_set_ground_curve, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetGroundCurve() float64 { //gd:ProceduralSkyMaterial.get_ground_curve
	var frame = callframe.New()
	var r_ret = callframe.Ret[float64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ProceduralSkyMaterial.Bind_get_ground_curve, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetGroundEnergyMultiplier(energy float64) { //gd:ProceduralSkyMaterial.set_ground_energy_multiplier
	var frame = callframe.New()
	callframe.Arg(frame, energy)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ProceduralSkyMaterial.Bind_set_ground_energy_multiplier, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetGroundEnergyMultiplier() float64 { //gd:ProceduralSkyMaterial.get_ground_energy_multiplier
	var frame = callframe.New()
	var r_ret = callframe.Ret[float64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ProceduralSkyMaterial.Bind_get_ground_energy_multiplier, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetSunAngleMax(degrees float64) { //gd:ProceduralSkyMaterial.set_sun_angle_max
	var frame = callframe.New()
	callframe.Arg(frame, degrees)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ProceduralSkyMaterial.Bind_set_sun_angle_max, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetSunAngleMax() float64 { //gd:ProceduralSkyMaterial.get_sun_angle_max
	var frame = callframe.New()
	var r_ret = callframe.Ret[float64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ProceduralSkyMaterial.Bind_get_sun_angle_max, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetSunCurve(curve float64) { //gd:ProceduralSkyMaterial.set_sun_curve
	var frame = callframe.New()
	callframe.Arg(frame, curve)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ProceduralSkyMaterial.Bind_set_sun_curve, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetSunCurve() float64 { //gd:ProceduralSkyMaterial.get_sun_curve
	var frame = callframe.New()
	var r_ret = callframe.Ret[float64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ProceduralSkyMaterial.Bind_get_sun_curve, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetUseDebanding(use_debanding bool) { //gd:ProceduralSkyMaterial.set_use_debanding
	var frame = callframe.New()
	callframe.Arg(frame, use_debanding)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ProceduralSkyMaterial.Bind_set_use_debanding, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetUseDebanding() bool { //gd:ProceduralSkyMaterial.get_use_debanding
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ProceduralSkyMaterial.Bind_get_use_debanding, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetEnergyMultiplier(multiplier float64) { //gd:ProceduralSkyMaterial.set_energy_multiplier
	var frame = callframe.New()
	callframe.Arg(frame, multiplier)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ProceduralSkyMaterial.Bind_set_energy_multiplier, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetEnergyMultiplier() float64 { //gd:ProceduralSkyMaterial.get_energy_multiplier
	var frame = callframe.New()
	var r_ret = callframe.Ret[float64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ProceduralSkyMaterial.Bind_get_energy_multiplier, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self class) AsProceduralSkyMaterial() Advanced    { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsProceduralSkyMaterial() Instance { return *((*Instance)(unsafe.Pointer(&self))) }
func (self class) AsMaterial() Material.Advanced {
	return *((*Material.Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsMaterial() Material.Instance {
	return *((*Material.Instance)(unsafe.Pointer(&self)))
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
		return gd.VirtualByName(Material.Advanced(self.AsMaterial()), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(Material.Instance(self.AsMaterial()), name)
	}
}
func init() {
	gdclass.Register("ProceduralSkyMaterial", func(ptr gd.Object) any {
		return [1]gdclass.ProceduralSkyMaterial{*(*gdclass.ProceduralSkyMaterial)(unsafe.Pointer(&ptr))}
	})
}
