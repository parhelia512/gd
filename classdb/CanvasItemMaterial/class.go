// Package CanvasItemMaterial provides methods for working with CanvasItemMaterial object instances.
package CanvasItemMaterial

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
[CanvasItemMaterial]s provide a means of modifying the textures associated with a CanvasItem. They specialize in describing blend and lighting behaviors for textures. Use a [ShaderMaterial] to more fully customize a material's interactions with a [CanvasItem].
*/
type Instance [1]gdclass.CanvasItemMaterial

// Nil is a nil/null instance of the class. Equivalent to the zero value.
var Nil Instance

type Any interface {
	gd.IsClass
	AsCanvasItemMaterial() Instance
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.CanvasItemMaterial

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("CanvasItemMaterial"))
	casted := Instance{*(*gdclass.CanvasItemMaterial)(unsafe.Pointer(&object))}
	casted.AsRefCounted()[0].Reference()
	return casted
}

func (self Instance) BlendMode() gdclass.CanvasItemMaterialBlendMode {
	return gdclass.CanvasItemMaterialBlendMode(class(self).GetBlendMode())
}

func (self Instance) SetBlendMode(value gdclass.CanvasItemMaterialBlendMode) {
	class(self).SetBlendMode(value)
}

func (self Instance) LightMode() gdclass.CanvasItemMaterialLightMode {
	return gdclass.CanvasItemMaterialLightMode(class(self).GetLightMode())
}

func (self Instance) SetLightMode(value gdclass.CanvasItemMaterialLightMode) {
	class(self).SetLightMode(value)
}

func (self Instance) ParticlesAnimation() bool {
	return bool(class(self).GetParticlesAnimation())
}

func (self Instance) SetParticlesAnimation(value bool) {
	class(self).SetParticlesAnimation(value)
}

func (self Instance) ParticlesAnimHFrames() int {
	return int(int(class(self).GetParticlesAnimHFrames()))
}

func (self Instance) SetParticlesAnimHFrames(value int) {
	class(self).SetParticlesAnimHFrames(int64(value))
}

func (self Instance) ParticlesAnimVFrames() int {
	return int(int(class(self).GetParticlesAnimVFrames()))
}

func (self Instance) SetParticlesAnimVFrames(value int) {
	class(self).SetParticlesAnimVFrames(int64(value))
}

func (self Instance) ParticlesAnimLoop() bool {
	return bool(class(self).GetParticlesAnimLoop())
}

func (self Instance) SetParticlesAnimLoop(value bool) {
	class(self).SetParticlesAnimLoop(value)
}

//go:nosplit
func (self class) SetBlendMode(blend_mode gdclass.CanvasItemMaterialBlendMode) { //gd:CanvasItemMaterial.set_blend_mode
	var frame = callframe.New()
	callframe.Arg(frame, blend_mode)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CanvasItemMaterial.Bind_set_blend_mode, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetBlendMode() gdclass.CanvasItemMaterialBlendMode { //gd:CanvasItemMaterial.get_blend_mode
	var frame = callframe.New()
	var r_ret = callframe.Ret[gdclass.CanvasItemMaterialBlendMode](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CanvasItemMaterial.Bind_get_blend_mode, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetLightMode(light_mode gdclass.CanvasItemMaterialLightMode) { //gd:CanvasItemMaterial.set_light_mode
	var frame = callframe.New()
	callframe.Arg(frame, light_mode)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CanvasItemMaterial.Bind_set_light_mode, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetLightMode() gdclass.CanvasItemMaterialLightMode { //gd:CanvasItemMaterial.get_light_mode
	var frame = callframe.New()
	var r_ret = callframe.Ret[gdclass.CanvasItemMaterialLightMode](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CanvasItemMaterial.Bind_get_light_mode, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetParticlesAnimation(particles_anim bool) { //gd:CanvasItemMaterial.set_particles_animation
	var frame = callframe.New()
	callframe.Arg(frame, particles_anim)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CanvasItemMaterial.Bind_set_particles_animation, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetParticlesAnimation() bool { //gd:CanvasItemMaterial.get_particles_animation
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CanvasItemMaterial.Bind_get_particles_animation, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetParticlesAnimHFrames(frames int64) { //gd:CanvasItemMaterial.set_particles_anim_h_frames
	var frame = callframe.New()
	callframe.Arg(frame, frames)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CanvasItemMaterial.Bind_set_particles_anim_h_frames, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetParticlesAnimHFrames() int64 { //gd:CanvasItemMaterial.get_particles_anim_h_frames
	var frame = callframe.New()
	var r_ret = callframe.Ret[int64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CanvasItemMaterial.Bind_get_particles_anim_h_frames, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetParticlesAnimVFrames(frames int64) { //gd:CanvasItemMaterial.set_particles_anim_v_frames
	var frame = callframe.New()
	callframe.Arg(frame, frames)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CanvasItemMaterial.Bind_set_particles_anim_v_frames, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetParticlesAnimVFrames() int64 { //gd:CanvasItemMaterial.get_particles_anim_v_frames
	var frame = callframe.New()
	var r_ret = callframe.Ret[int64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CanvasItemMaterial.Bind_get_particles_anim_v_frames, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetParticlesAnimLoop(loop bool) { //gd:CanvasItemMaterial.set_particles_anim_loop
	var frame = callframe.New()
	callframe.Arg(frame, loop)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CanvasItemMaterial.Bind_set_particles_anim_loop, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetParticlesAnimLoop() bool { //gd:CanvasItemMaterial.get_particles_anim_loop
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CanvasItemMaterial.Bind_get_particles_anim_loop, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self class) AsCanvasItemMaterial() Advanced    { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsCanvasItemMaterial() Instance { return *((*Instance)(unsafe.Pointer(&self))) }
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
	gdclass.Register("CanvasItemMaterial", func(ptr gd.Object) any {
		return [1]gdclass.CanvasItemMaterial{*(*gdclass.CanvasItemMaterial)(unsafe.Pointer(&ptr))}
	})
}

type BlendMode = gdclass.CanvasItemMaterialBlendMode //gd:CanvasItemMaterial.BlendMode

const (
	/*Mix blending mode. Colors are assumed to be independent of the alpha (opacity) value.*/
	BlendModeMix BlendMode = 0
	/*Additive blending mode.*/
	BlendModeAdd BlendMode = 1
	/*Subtractive blending mode.*/
	BlendModeSub BlendMode = 2
	/*Multiplicative blending mode.*/
	BlendModeMul BlendMode = 3
	/*Mix blending mode. Colors are assumed to be premultiplied by the alpha (opacity) value.*/
	BlendModePremultAlpha BlendMode = 4
)

type LightMode = gdclass.CanvasItemMaterialLightMode //gd:CanvasItemMaterial.LightMode

const (
	/*Render the material using both light and non-light sensitive material properties.*/
	LightModeNormal LightMode = 0
	/*Render the material as if there were no light.*/
	LightModeUnshaded LightMode = 1
	/*Render the material as if there were only light.*/
	LightModeLightOnly LightMode = 2
)
