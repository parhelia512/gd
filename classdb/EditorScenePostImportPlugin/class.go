// Package EditorScenePostImportPlugin provides methods for working with EditorScenePostImportPlugin object instances.
package EditorScenePostImportPlugin

import "unsafe"
import "reflect"
import "slices"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/internal/gdclass"
import "graphics.gd/variant"
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
This plugin type exists to modify the process of importing scenes, allowing to change the content as well as add importer options at every stage of the process.

	See [Interface] for methods that can be overridden by a [Class] that extends it.

%!(EXTRA string=EditorScenePostImportPlugin)
*/
type Instance [1]gdclass.EditorScenePostImportPlugin
type Expanded [1]gdclass.EditorScenePostImportPlugin

// Nil is a nil/null instance of the class. Equivalent to the zero value.
var Nil Instance

type Any interface {
	gd.IsClass
	AsEditorScenePostImportPlugin() Instance
}
type Interface interface {
	//Override to add internal import options. These will appear in the 3D scene import dialog. Add options via [method add_import_option] and [method add_import_option_advanced].
	GetInternalImportOptions(category int)
	//Should return [code]true[/code] to show the given option, [code]false[/code] to hide the given option, or [code]null[/code] to ignore.
	GetInternalOptionVisibility(category int, for_animation bool, option string) any
	//Should return [code]true[/code] if the 3D view of the import dialog needs to update when changing the given option.
	GetInternalOptionUpdateViewRequired(category int, option string) any
	//Process a specific node or resource for a given category.
	InternalProcess(category int, base_node [1]gdclass.Node, node [1]gdclass.Node, resource [1]gdclass.Resource)
	//Override to add general import options. These will appear in the main import dock on the editor. Add options via [method add_import_option] and [method add_import_option_advanced].
	GetImportOptions(path string)
	//Should return [code]true[/code] to show the given option, [code]false[/code] to hide the given option, or [code]null[/code] to ignore.
	GetOptionVisibility(path string, for_animation bool, option string) any
	//Pre Process the scene. This function is called right after the scene format loader loaded the scene and no changes have been made.
	//Pre process may be used to adjust internal import options in the [code]"nodes"[/code], [code]"meshes"[/code], [code]"animations"[/code] or [code]"materials"[/code] keys inside [code]get_option_value("_subresources")[/code].
	PreProcess(scene [1]gdclass.Node)
	//Post process the scene. This function is called after the final scene has been configured.
	PostProcess(scene [1]gdclass.Node)
}

// Implementation implements [Interface] with empty methods.
type Implementation = implementation

type implementation struct{}

func (self implementation) GetInternalImportOptions(category int) { return }
func (self implementation) GetInternalOptionVisibility(category int, for_animation bool, option string) (_ any) {
	return
}
func (self implementation) GetInternalOptionUpdateViewRequired(category int, option string) (_ any) {
	return
}
func (self implementation) InternalProcess(category int, base_node [1]gdclass.Node, node [1]gdclass.Node, resource [1]gdclass.Resource) {
	return
}
func (self implementation) GetImportOptions(path string) { return }
func (self implementation) GetOptionVisibility(path string, for_animation bool, option string) (_ any) {
	return
}
func (self implementation) PreProcess(scene [1]gdclass.Node)  { return }
func (self implementation) PostProcess(scene [1]gdclass.Node) { return }

/*
Override to add internal import options. These will appear in the 3D scene import dialog. Add options via [method add_import_option] and [method add_import_option_advanced].
*/
func (Instance) _get_internal_import_options(impl func(ptr unsafe.Pointer, category int)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var category = gd.UnsafeGet[int64](p_args, 0)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, int(category))
	}
}

/*
Should return [code]true[/code] to show the given option, [code]false[/code] to hide the given option, or [code]null[/code] to ignore.
*/
func (Instance) _get_internal_option_visibility(impl func(ptr unsafe.Pointer, category int, for_animation bool, option string) any) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var category = gd.UnsafeGet[int64](p_args, 0)
		var for_animation = gd.UnsafeGet[bool](p_args, 1)
		var option = String.Via(gd.StringProxy{}, pointers.Pack(pointers.New[gd.String](gd.UnsafeGet[[1]gd.EnginePointer](p_args, 2))))
		defer pointers.End(gd.InternalString(option))
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, int(category), for_animation, option.String())
		ptr, ok := pointers.End(gd.InternalVariant(variant.New(ret)))

		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

/*
Should return [code]true[/code] if the 3D view of the import dialog needs to update when changing the given option.
*/
func (Instance) _get_internal_option_update_view_required(impl func(ptr unsafe.Pointer, category int, option string) any) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var category = gd.UnsafeGet[int64](p_args, 0)
		var option = String.Via(gd.StringProxy{}, pointers.Pack(pointers.New[gd.String](gd.UnsafeGet[[1]gd.EnginePointer](p_args, 1))))
		defer pointers.End(gd.InternalString(option))
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, int(category), option.String())
		ptr, ok := pointers.End(gd.InternalVariant(variant.New(ret)))

		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

/*
Process a specific node or resource for a given category.
*/
func (Instance) _internal_process(impl func(ptr unsafe.Pointer, category int, base_node [1]gdclass.Node, node [1]gdclass.Node, resource [1]gdclass.Resource)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var category = gd.UnsafeGet[int64](p_args, 0)
		var base_node = [1]gdclass.Node{pointers.New[gdclass.Node]([3]uint64{uint64(gd.UnsafeGet[gd.EnginePointer](p_args, 1))})}

		defer pointers.End(base_node[0])
		var node = [1]gdclass.Node{pointers.New[gdclass.Node]([3]uint64{uint64(gd.UnsafeGet[gd.EnginePointer](p_args, 2))})}

		defer pointers.End(node[0])
		var resource = [1]gdclass.Resource{pointers.New[gdclass.Resource]([3]uint64{uint64(gd.UnsafeGet[gd.EnginePointer](p_args, 3))})}

		defer pointers.End(resource[0])
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, int(category), base_node, node, resource)
	}
}

/*
Override to add general import options. These will appear in the main import dock on the editor. Add options via [method add_import_option] and [method add_import_option_advanced].
*/
func (Instance) _get_import_options(impl func(ptr unsafe.Pointer, path string)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var path = String.Via(gd.StringProxy{}, pointers.Pack(pointers.New[gd.String](gd.UnsafeGet[[1]gd.EnginePointer](p_args, 0))))
		defer pointers.End(gd.InternalString(path))
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, path.String())
	}
}

/*
Should return [code]true[/code] to show the given option, [code]false[/code] to hide the given option, or [code]null[/code] to ignore.
*/
func (Instance) _get_option_visibility(impl func(ptr unsafe.Pointer, path string, for_animation bool, option string) any) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var path = String.Via(gd.StringProxy{}, pointers.Pack(pointers.New[gd.String](gd.UnsafeGet[[1]gd.EnginePointer](p_args, 0))))
		defer pointers.End(gd.InternalString(path))
		var for_animation = gd.UnsafeGet[bool](p_args, 1)
		var option = String.Via(gd.StringProxy{}, pointers.Pack(pointers.New[gd.String](gd.UnsafeGet[[1]gd.EnginePointer](p_args, 2))))
		defer pointers.End(gd.InternalString(option))
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, path.String(), for_animation, option.String())
		ptr, ok := pointers.End(gd.InternalVariant(variant.New(ret)))

		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

/*
Pre Process the scene. This function is called right after the scene format loader loaded the scene and no changes have been made.
Pre process may be used to adjust internal import options in the [code]"nodes"[/code], [code]"meshes"[/code], [code]"animations"[/code] or [code]"materials"[/code] keys inside [code]get_option_value("_subresources")[/code].
*/
func (Instance) _pre_process(impl func(ptr unsafe.Pointer, scene [1]gdclass.Node)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var scene = [1]gdclass.Node{pointers.New[gdclass.Node]([3]uint64{uint64(gd.UnsafeGet[gd.EnginePointer](p_args, 0))})}

		defer pointers.End(scene[0])
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, scene)
	}
}

/*
Post process the scene. This function is called after the final scene has been configured.
*/
func (Instance) _post_process(impl func(ptr unsafe.Pointer, scene [1]gdclass.Node)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var scene = [1]gdclass.Node{pointers.New[gdclass.Node]([3]uint64{uint64(gd.UnsafeGet[gd.EnginePointer](p_args, 0))})}

		defer pointers.End(scene[0])
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, scene)
	}
}

/*
Query the value of an option. This function can only be called from those querying visibility, or processing.
*/
func (self Instance) GetOptionValue(name string) any { //gd:EditorScenePostImportPlugin.get_option_value
	return any(Advanced(self).GetOptionValue(String.Name(String.New(name))).Interface())
}

/*
Add a specific import option (name and default value only). This function can only be called from [method _get_import_options] and [method _get_internal_import_options].
*/
func (self Instance) AddImportOption(name string, value any) { //gd:EditorScenePostImportPlugin.add_import_option
	Advanced(self).AddImportOption(String.New(name), variant.New(value))
}

/*
Add a specific import option. This function can only be called from [method _get_import_options] and [method _get_internal_import_options].
*/
func (self Instance) AddImportOptionAdvanced(atype variant.Type, name string, default_value any) { //gd:EditorScenePostImportPlugin.add_import_option_advanced
	Advanced(self).AddImportOptionAdvanced(atype, String.New(name), variant.New(default_value), 0, String.New(""), int64(6))
}

/*
Add a specific import option. This function can only be called from [method _get_import_options] and [method _get_internal_import_options].
*/
func (self Expanded) AddImportOptionAdvanced(atype variant.Type, name string, default_value any, hint PropertyHint, hint_string string, usage_flags int) { //gd:EditorScenePostImportPlugin.add_import_option_advanced
	Advanced(self).AddImportOptionAdvanced(atype, String.New(name), variant.New(default_value), hint, String.New(hint_string), int64(usage_flags))
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.EditorScenePostImportPlugin

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("EditorScenePostImportPlugin"))
	casted := Instance{*(*gdclass.EditorScenePostImportPlugin)(unsafe.Pointer(&object))}
	casted.AsRefCounted()[0].Reference()
	return casted
}

/*
Override to add internal import options. These will appear in the 3D scene import dialog. Add options via [method add_import_option] and [method add_import_option_advanced].
*/
func (class) _get_internal_import_options(impl func(ptr unsafe.Pointer, category int64)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var category = gd.UnsafeGet[int64](p_args, 0)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, category)
	}
}

/*
Should return [code]true[/code] to show the given option, [code]false[/code] to hide the given option, or [code]null[/code] to ignore.
*/
func (class) _get_internal_option_visibility(impl func(ptr unsafe.Pointer, category int64, for_animation bool, option String.Readable) variant.Any) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var category = gd.UnsafeGet[int64](p_args, 0)
		var for_animation = gd.UnsafeGet[bool](p_args, 1)
		var option = String.Via(gd.StringProxy{}, pointers.Pack(pointers.New[gd.String](gd.UnsafeGet[[1]gd.EnginePointer](p_args, 2))))
		defer pointers.End(gd.InternalString(option))
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, category, for_animation, option)
		ptr, ok := pointers.End(gd.InternalVariant(ret))

		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

/*
Should return [code]true[/code] if the 3D view of the import dialog needs to update when changing the given option.
*/
func (class) _get_internal_option_update_view_required(impl func(ptr unsafe.Pointer, category int64, option String.Readable) variant.Any) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var category = gd.UnsafeGet[int64](p_args, 0)
		var option = String.Via(gd.StringProxy{}, pointers.Pack(pointers.New[gd.String](gd.UnsafeGet[[1]gd.EnginePointer](p_args, 1))))
		defer pointers.End(gd.InternalString(option))
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, category, option)
		ptr, ok := pointers.End(gd.InternalVariant(ret))

		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

/*
Process a specific node or resource for a given category.
*/
func (class) _internal_process(impl func(ptr unsafe.Pointer, category int64, base_node [1]gdclass.Node, node [1]gdclass.Node, resource [1]gdclass.Resource)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var category = gd.UnsafeGet[int64](p_args, 0)
		var base_node = [1]gdclass.Node{pointers.New[gdclass.Node]([3]uint64{uint64(gd.UnsafeGet[gd.EnginePointer](p_args, 1))})}

		defer pointers.End(base_node[0])
		var node = [1]gdclass.Node{pointers.New[gdclass.Node]([3]uint64{uint64(gd.UnsafeGet[gd.EnginePointer](p_args, 2))})}

		defer pointers.End(node[0])
		var resource = [1]gdclass.Resource{pointers.New[gdclass.Resource]([3]uint64{uint64(gd.UnsafeGet[gd.EnginePointer](p_args, 3))})}

		defer pointers.End(resource[0])
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, category, base_node, node, resource)
	}
}

/*
Override to add general import options. These will appear in the main import dock on the editor. Add options via [method add_import_option] and [method add_import_option_advanced].
*/
func (class) _get_import_options(impl func(ptr unsafe.Pointer, path String.Readable)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var path = String.Via(gd.StringProxy{}, pointers.Pack(pointers.New[gd.String](gd.UnsafeGet[[1]gd.EnginePointer](p_args, 0))))
		defer pointers.End(gd.InternalString(path))
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, path)
	}
}

/*
Should return [code]true[/code] to show the given option, [code]false[/code] to hide the given option, or [code]null[/code] to ignore.
*/
func (class) _get_option_visibility(impl func(ptr unsafe.Pointer, path String.Readable, for_animation bool, option String.Readable) variant.Any) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var path = String.Via(gd.StringProxy{}, pointers.Pack(pointers.New[gd.String](gd.UnsafeGet[[1]gd.EnginePointer](p_args, 0))))
		defer pointers.End(gd.InternalString(path))
		var for_animation = gd.UnsafeGet[bool](p_args, 1)
		var option = String.Via(gd.StringProxy{}, pointers.Pack(pointers.New[gd.String](gd.UnsafeGet[[1]gd.EnginePointer](p_args, 2))))
		defer pointers.End(gd.InternalString(option))
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, path, for_animation, option)
		ptr, ok := pointers.End(gd.InternalVariant(ret))

		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

/*
Pre Process the scene. This function is called right after the scene format loader loaded the scene and no changes have been made.
Pre process may be used to adjust internal import options in the [code]"nodes"[/code], [code]"meshes"[/code], [code]"animations"[/code] or [code]"materials"[/code] keys inside [code]get_option_value("_subresources")[/code].
*/
func (class) _pre_process(impl func(ptr unsafe.Pointer, scene [1]gdclass.Node)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var scene = [1]gdclass.Node{pointers.New[gdclass.Node]([3]uint64{uint64(gd.UnsafeGet[gd.EnginePointer](p_args, 0))})}

		defer pointers.End(scene[0])
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, scene)
	}
}

/*
Post process the scene. This function is called after the final scene has been configured.
*/
func (class) _post_process(impl func(ptr unsafe.Pointer, scene [1]gdclass.Node)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var scene = [1]gdclass.Node{pointers.New[gdclass.Node]([3]uint64{uint64(gd.UnsafeGet[gd.EnginePointer](p_args, 0))})}

		defer pointers.End(scene[0])
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, scene)
	}
}

/*
Query the value of an option. This function can only be called from those querying visibility, or processing.
*/
//go:nosplit
func (self class) GetOptionValue(name String.Name) variant.Any { //gd:EditorScenePostImportPlugin.get_option_value
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalStringName(name)))
	var r_ret = callframe.Ret[[3]uint64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.EditorScenePostImportPlugin.Bind_get_option_value, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = variant.Implementation(gd.VariantProxy{}, pointers.Pack(pointers.New[gd.Variant](r_ret.Get())))
	frame.Free()
	return ret
}

/*
Add a specific import option (name and default value only). This function can only be called from [method _get_import_options] and [method _get_internal_import_options].
*/
//go:nosplit
func (self class) AddImportOption(name String.Readable, value variant.Any) { //gd:EditorScenePostImportPlugin.add_import_option
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalString(name)))
	callframe.Arg(frame, pointers.Get(gd.InternalVariant(value)))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.EditorScenePostImportPlugin.Bind_add_import_option, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Add a specific import option. This function can only be called from [method _get_import_options] and [method _get_internal_import_options].
*/
//go:nosplit
func (self class) AddImportOptionAdvanced(atype variant.Type, name String.Readable, default_value variant.Any, hint PropertyHint, hint_string String.Readable, usage_flags int64) { //gd:EditorScenePostImportPlugin.add_import_option_advanced
	var frame = callframe.New()
	callframe.Arg(frame, atype)
	callframe.Arg(frame, pointers.Get(gd.InternalString(name)))
	callframe.Arg(frame, pointers.Get(gd.InternalVariant(default_value)))
	callframe.Arg(frame, hint)
	callframe.Arg(frame, pointers.Get(gd.InternalString(hint_string)))
	callframe.Arg(frame, usage_flags)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.EditorScenePostImportPlugin.Bind_add_import_option_advanced, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}
func (self class) AsEditorScenePostImportPlugin() Advanced {
	return *((*Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsEditorScenePostImportPlugin() Instance {
	return *((*Instance)(unsafe.Pointer(&self)))
}
func (self class) AsRefCounted() [1]gd.RefCounted {
	return *((*[1]gd.RefCounted)(unsafe.Pointer(&self)))
}
func (self Instance) AsRefCounted() [1]gd.RefCounted {
	return *((*[1]gd.RefCounted)(unsafe.Pointer(&self)))
}

func (self class) Virtual(name string) reflect.Value {
	switch name {
	case "_get_internal_import_options":
		return reflect.ValueOf(self._get_internal_import_options)
	case "_get_internal_option_visibility":
		return reflect.ValueOf(self._get_internal_option_visibility)
	case "_get_internal_option_update_view_required":
		return reflect.ValueOf(self._get_internal_option_update_view_required)
	case "_internal_process":
		return reflect.ValueOf(self._internal_process)
	case "_get_import_options":
		return reflect.ValueOf(self._get_import_options)
	case "_get_option_visibility":
		return reflect.ValueOf(self._get_option_visibility)
	case "_pre_process":
		return reflect.ValueOf(self._pre_process)
	case "_post_process":
		return reflect.ValueOf(self._post_process)
	default:
		return gd.VirtualByName(RefCounted.Advanced(self.AsRefCounted()), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	case "_get_internal_import_options":
		return reflect.ValueOf(self._get_internal_import_options)
	case "_get_internal_option_visibility":
		return reflect.ValueOf(self._get_internal_option_visibility)
	case "_get_internal_option_update_view_required":
		return reflect.ValueOf(self._get_internal_option_update_view_required)
	case "_internal_process":
		return reflect.ValueOf(self._internal_process)
	case "_get_import_options":
		return reflect.ValueOf(self._get_import_options)
	case "_get_option_visibility":
		return reflect.ValueOf(self._get_option_visibility)
	case "_pre_process":
		return reflect.ValueOf(self._pre_process)
	case "_post_process":
		return reflect.ValueOf(self._post_process)
	default:
		return gd.VirtualByName(RefCounted.Instance(self.AsRefCounted()), name)
	}
}
func init() {
	gdclass.Register("EditorScenePostImportPlugin", func(ptr gd.Object) any {
		return [1]gdclass.EditorScenePostImportPlugin{*(*gdclass.EditorScenePostImportPlugin)(unsafe.Pointer(&ptr))}
	})
}

type InternalImportCategory = gdclass.EditorScenePostImportPluginInternalImportCategory //gd:EditorScenePostImportPlugin.InternalImportCategory

const (
	InternalImportCategoryNode           InternalImportCategory = 0
	InternalImportCategoryMesh3dNode     InternalImportCategory = 1
	InternalImportCategoryMesh           InternalImportCategory = 2
	InternalImportCategoryMaterial       InternalImportCategory = 3
	InternalImportCategoryAnimation      InternalImportCategory = 4
	InternalImportCategoryAnimationNode  InternalImportCategory = 5
	InternalImportCategorySkeleton3dNode InternalImportCategory = 6
	InternalImportCategoryMax            InternalImportCategory = 7
)

type PropertyHint int

const (
	/*The property has no hint for the editor.*/
	PropertyHintNone PropertyHint = 0
	/*Hints that an [int] or [float] property should be within a range specified via the hint string [code]"min,max"[/code] or [code]"min,max,step"[/code]. The hint string can optionally include [code]"or_greater"[/code] and/or [code]"or_less"[/code] to allow manual input going respectively above the max or below the min values.
	  [b]Example:[/b] [code]"-360,360,1,or_greater,or_less"[/code].
	  Additionally, other keywords can be included: [code]"exp"[/code] for exponential range editing, [code]"radians_as_degrees"[/code] for editing radian angles in degrees (the range values are also in degrees), [code]"degrees"[/code] to hint at an angle and [code]"hide_slider"[/code] to hide the slider.*/
	PropertyHintRange PropertyHint = 1
	/*Hints that an [int] or [String] property is an enumerated value to pick in a list specified via a hint string.
	  The hint string is a comma separated list of names such as [code]"Hello,Something,Else"[/code]. Whitespaces are [b]not[/b] removed from either end of a name. For integer properties, the first name in the list has value 0, the next 1, and so on. Explicit values can also be specified by appending [code]:integer[/code] to the name, e.g. [code]"Zero,One,Three:3,Four,Six:6"[/code].*/
	PropertyHintEnum PropertyHint = 2
	/*Hints that a [String] property can be an enumerated value to pick in a list specified via a hint string such as [code]"Hello,Something,Else"[/code].
	  Unlike [constant PROPERTY_HINT_ENUM], a property with this hint still accepts arbitrary values and can be empty. The list of values serves to suggest possible values.*/
	PropertyHintEnumSuggestion PropertyHint = 3
	/*Hints that a [float] property should be edited via an exponential easing function. The hint string can include [code]"attenuation"[/code] to flip the curve horizontally and/or [code]"positive_only"[/code] to exclude in/out easing and limit values to be greater than or equal to zero.*/
	PropertyHintExpEasing PropertyHint = 4
	/*Hints that a vector property should allow its components to be linked. For example, this allows [member Vector2.x] and [member Vector2.y] to be edited together.*/
	PropertyHintLink PropertyHint = 5
	/*Hints that an [int] property is a bitmask with named bit flags.
	  The hint string is a comma separated list of names such as [code]"Bit0,Bit1,Bit2,Bit3"[/code]. Whitespaces are [b]not[/b] removed from either end of a name. The first name in the list has value 1, the next 2, then 4, 8, 16 and so on. Explicit values can also be specified by appending [code]:integer[/code] to the name, e.g. [code]"A:4,B:8,C:16"[/code]. You can also combine several flags ([code]"A:4,B:8,AB:12,C:16"[/code]).
	  [b]Note:[/b] A flag value must be at least [code]1[/code] and at most [code]2 ** 32 - 1[/code].
	  [b]Note:[/b] Unlike [constant PROPERTY_HINT_ENUM], the previous explicit value is not taken into account. For the hint [code]"A:16,B,C"[/code], A is 16, B is 2, C is 4.*/
	PropertyHintFlags PropertyHint = 6
	/*Hints that an [int] property is a bitmask using the optionally named 2D render layers.*/
	PropertyHintLayers2dRender PropertyHint = 7
	/*Hints that an [int] property is a bitmask using the optionally named 2D physics layers.*/
	PropertyHintLayers2dPhysics PropertyHint = 8
	/*Hints that an [int] property is a bitmask using the optionally named 2D navigation layers.*/
	PropertyHintLayers2dNavigation PropertyHint = 9
	/*Hints that an [int] property is a bitmask using the optionally named 3D render layers.*/
	PropertyHintLayers3dRender PropertyHint = 10
	/*Hints that an [int] property is a bitmask using the optionally named 3D physics layers.*/
	PropertyHintLayers3dPhysics PropertyHint = 11
	/*Hints that an [int] property is a bitmask using the optionally named 3D navigation layers.*/
	PropertyHintLayers3dNavigation PropertyHint = 12
	/*Hints that an integer property is a bitmask using the optionally named avoidance layers.*/
	PropertyHintLayersAvoidance PropertyHint = 37
	/*Hints that a [String] property is a path to a file. Editing it will show a file dialog for picking the path. The hint string can be a set of filters with wildcards like [code]"*.png,*.jpg"[/code].*/
	PropertyHintFile PropertyHint = 13
	/*Hints that a [String] property is a path to a directory. Editing it will show a file dialog for picking the path.*/
	PropertyHintDir PropertyHint = 14
	/*Hints that a [String] property is an absolute path to a file outside the project folder. Editing it will show a file dialog for picking the path. The hint string can be a set of filters with wildcards, like [code]"*.png,*.jpg"[/code].*/
	PropertyHintGlobalFile PropertyHint = 15
	/*Hints that a [String] property is an absolute path to a directory outside the project folder. Editing it will show a file dialog for picking the path.*/
	PropertyHintGlobalDir PropertyHint = 16
	/*Hints that a property is an instance of a [Resource]-derived type, optionally specified via the hint string (e.g. [code]"Texture2D"[/code]). Editing it will show a popup menu of valid resource types to instantiate.*/
	PropertyHintResourceType PropertyHint = 17
	/*Hints that a [String] property is text with line breaks. Editing it will show a text input field where line breaks can be typed.*/
	PropertyHintMultilineText PropertyHint = 18
	/*Hints that a [String] property is an [Expression].*/
	PropertyHintExpression PropertyHint = 19
	/*Hints that a [String] property should show a placeholder text on its input field, if empty. The hint string is the placeholder text to use.*/
	PropertyHintPlaceholderText PropertyHint = 20
	/*Hints that a [Color] property should be edited without affecting its transparency ([member Color.a] is not editable).*/
	PropertyHintColorNoAlpha PropertyHint = 21
	/*Hints that the property's value is an object encoded as object ID, with its type specified in the hint string. Used by the debugger.*/
	PropertyHintObjectId PropertyHint = 22
	/*If a property is [String], hints that the property represents a particular type (class). This allows to select a type from the create dialog. The property will store the selected type as a string.
	  If a property is [Array], hints the editor how to show elements. The [code]hint_string[/code] must encode nested types using [code]":"[/code] and [code]"/"[/code].
	  [codeblocks]
	  [gdscript]
	  # Array of elem_type.
	  hint_string = "%d:" % [elem_type]
	  hint_string = "%d/%d:%s" % [elem_type, elem_hint, elem_hint_string]
	  # Two-dimensional array of elem_type (array of arrays of elem_type).
	  hint_string = "%d:%d:" % [TYPE_ARRAY, elem_type]
	  hint_string = "%d:%d/%d:%s" % [TYPE_ARRAY, elem_type, elem_hint, elem_hint_string]
	  # Three-dimensional array of elem_type (array of arrays of arrays of elem_type).
	  hint_string = "%d:%d:%d:" % [TYPE_ARRAY, TYPE_ARRAY, elem_type]
	  hint_string = "%d:%d:%d/%d:%s" % [TYPE_ARRAY, TYPE_ARRAY, elem_type, elem_hint, elem_hint_string]
	  [/gdscript]
	  [csharp]
	  // Array of elemType.
	  hintString = $"{elemType:D}:";
	  hintString = $"{elemType:}/{elemHint:D}:{elemHintString}";
	  // Two-dimensional array of elemType (array of arrays of elemType).
	  hintString = $"{Variant.Type.Array:D}:{elemType:D}:";
	  hintString = $"{Variant.Type.Array:D}:{elemType:D}/{elemHint:D}:{elemHintString}";
	  // Three-dimensional array of elemType (array of arrays of arrays of elemType).
	  hintString = $"{Variant.Type.Array:D}:{Variant.Type.Array:D}:{elemType:D}:";
	  hintString = $"{Variant.Type.Array:D}:{Variant.Type.Array:D}:{elemType:D}/{elemHint:D}:{elemHintString}";
	  [/csharp]
	  [/codeblocks]
	  [b]Examples:[/b]
	  [codeblocks]
	  [gdscript]
	  hint_string = "%d:" % [TYPE_INT] # Array of integers.
	  hint_string = "%d/%d:1,10,1" % [TYPE_INT, PROPERTY_HINT_RANGE] # Array of integers (in range from 1 to 10).
	  hint_string = "%d/%d:Zero,One,Two" % [TYPE_INT, PROPERTY_HINT_ENUM] # Array of integers (an enum).
	  hint_string = "%d/%d:Zero,One,Three:3,Six:6" % [TYPE_INT, PROPERTY_HINT_ENUM] # Array of integers (an enum).
	  hint_string = "%d/%d:*.png" % [TYPE_STRING, PROPERTY_HINT_FILE] # Array of strings (file paths).
	  hint_string = "%d/%d:Texture2D" % [TYPE_OBJECT, PROPERTY_HINT_RESOURCE_TYPE] # Array of textures.

	  hint_string = "%d:%d:" % [TYPE_ARRAY, TYPE_FLOAT] # Two-dimensional array of floats.
	  hint_string = "%d:%d/%d:" % [TYPE_ARRAY, TYPE_STRING, PROPERTY_HINT_MULTILINE_TEXT] # Two-dimensional array of multiline strings.
	  hint_string = "%d:%d/%d:-1,1,0.1" % [TYPE_ARRAY, TYPE_FLOAT, PROPERTY_HINT_RANGE] # Two-dimensional array of floats (in range from -1 to 1).
	  hint_string = "%d:%d/%d:Texture2D" % [TYPE_ARRAY, TYPE_OBJECT, PROPERTY_HINT_RESOURCE_TYPE] # Two-dimensional array of textures.
	  [/gdscript]
	  [csharp]
	  hintString = $"{Variant.Type.Int:D}/{PropertyHint.Range:D}:1,10,1"; // Array of integers (in range from 1 to 10).
	  hintString = $"{Variant.Type.Int:D}/{PropertyHint.Enum:D}:Zero,One,Two"; // Array of integers (an enum).
	  hintString = $"{Variant.Type.Int:D}/{PropertyHint.Enum:D}:Zero,One,Three:3,Six:6"; // Array of integers (an enum).
	  hintString = $"{Variant.Type.String:D}/{PropertyHint.File:D}:*.png"; // Array of strings (file paths).
	  hintString = $"{Variant.Type.Object:D}/{PropertyHint.ResourceType:D}:Texture2D"; // Array of textures.

	  hintString = $"{Variant.Type.Array:D}:{Variant.Type.Float:D}:"; // Two-dimensional array of floats.
	  hintString = $"{Variant.Type.Array:D}:{Variant.Type.String:D}/{PropertyHint.MultilineText:D}:"; // Two-dimensional array of multiline strings.
	  hintString = $"{Variant.Type.Array:D}:{Variant.Type.Float:D}/{PropertyHint.Range:D}:-1,1,0.1"; // Two-dimensional array of floats (in range from -1 to 1).
	  hintString = $"{Variant.Type.Array:D}:{Variant.Type.Object:D}/{PropertyHint.ResourceType:D}:Texture2D"; // Two-dimensional array of textures.
	  [/csharp]
	  [/codeblocks]
	  [b]Note:[/b] The trailing colon is required for properly detecting built-in types.*/
	PropertyHintTypeString           PropertyHint = 23
	PropertyHintNodePathToEditedNode PropertyHint = 24
	/*Hints that an object is too big to be sent via the debugger.*/
	PropertyHintObjectTooBig PropertyHint = 25
	/*Hints that the hint string specifies valid node types for property of type [NodePath].*/
	PropertyHintNodePathValidTypes PropertyHint = 26
	/*Hints that a [String] property is a path to a file. Editing it will show a file dialog for picking the path for the file to be saved at. The dialog has access to the project's directory. The hint string can be a set of filters with wildcards like [code]"*.png,*.jpg"[/code]. See also [member FileDialog.filters].*/
	PropertyHintSaveFile PropertyHint = 27
	/*Hints that a [String] property is a path to a file. Editing it will show a file dialog for picking the path for the file to be saved at. The dialog has access to the entire filesystem. The hint string can be a set of filters with wildcards like [code]"*.png,*.jpg"[/code]. See also [member FileDialog.filters].*/
	PropertyHintGlobalSaveFile PropertyHint = 28
	PropertyHintIntIsObjectid  PropertyHint = 29
	/*Hints that an [int] property is a pointer. Used by GDExtension.*/
	PropertyHintIntIsPointer PropertyHint = 30
	/*Hints that a property is an [Array] with the stored type specified in the hint string.*/
	PropertyHintArrayType PropertyHint = 31
	/*Hints that a property is a [Dictionary] with the stored types specified in the hint string.*/
	PropertyHintDictionaryType PropertyHint = 38
	/*Hints that a string property is a locale code. Editing it will show a locale dialog for picking language and country.*/
	PropertyHintLocaleId PropertyHint = 32
	/*Hints that a dictionary property is string translation map. Dictionary keys are locale codes and, values are translated strings.*/
	PropertyHintLocalizableString PropertyHint = 33
	/*Hints that a property is an instance of a [Node]-derived type, optionally specified via the hint string (e.g. [code]"Node2D"[/code]). Editing it will show a dialog for picking a node from the scene.*/
	PropertyHintNodeType PropertyHint = 34
	/*Hints that a quaternion property should disable the temporary euler editor.*/
	PropertyHintHideQuaternionEdit PropertyHint = 35
	/*Hints that a string property is a password, and every character is replaced with the secret character.*/
	PropertyHintPassword PropertyHint = 36
	/*Hints that a [Callable] property should be displayed as a clickable button. When the button is pressed, the callable is called. The hint string specifies the button text and optionally an icon from the [code]"EditorIcons"[/code] theme type.
	  [codeblock lang=text]
	  "Click me!" - A button with the text "Click me!" and the default "Callable" icon.
	  "Click me!,ColorRect" - A button with the text "Click me!" and the "ColorRect" icon.
	  [/codeblock]
	  [b]Note:[/b] A [Callable] cannot be properly serialized and stored in a file, so it is recommended to use [constant PROPERTY_USAGE_EDITOR] instead of [constant PROPERTY_USAGE_DEFAULT].*/
	PropertyHintToolButton PropertyHint = 39
	/*Hints that a property will be changed on its own after setting, such as [member AudioStreamPlayer.playing] or [member GPUParticles3D.emitting].*/
	PropertyHintOneshot PropertyHint = 40
	/*Represents the size of the [enum PropertyHint] enum.*/
	PropertyHintMax PropertyHint = 42
)

type VariantType int

const (
	/*Variable is [code]null[/code].*/
	TypeNil VariantType = 0
	/*Variable is of type [bool].*/
	TypeBool VariantType = 1
	/*Variable is of type [int].*/
	TypeInt VariantType = 2
	/*Variable is of type [float].*/
	TypeFloat VariantType = 3
	/*Variable is of type [String].*/
	TypeString VariantType = 4
	/*Variable is of type [Vector2].*/
	TypeVector2 VariantType = 5
	/*Variable is of type [Vector2i].*/
	TypeVector2i VariantType = 6
	/*Variable is of type [Rect2].*/
	TypeRect2 VariantType = 7
	/*Variable is of type [Rect2i].*/
	TypeRect2i VariantType = 8
	/*Variable is of type [Vector3].*/
	TypeVector3 VariantType = 9
	/*Variable is of type [Vector3i].*/
	TypeVector3i VariantType = 10
	/*Variable is of type [Transform2D].*/
	TypeTransform2d VariantType = 11
	/*Variable is of type [Vector4].*/
	TypeVector4 VariantType = 12
	/*Variable is of type [Vector4i].*/
	TypeVector4i VariantType = 13
	/*Variable is of type [Plane].*/
	TypePlane VariantType = 14
	/*Variable is of type [Quaternion].*/
	TypeQuaternion VariantType = 15
	/*Variable is of type [AABB].*/
	TypeAabb VariantType = 16
	/*Variable is of type [Basis].*/
	TypeBasis VariantType = 17
	/*Variable is of type [Transform3D].*/
	TypeTransform3d VariantType = 18
	/*Variable is of type [Projection].*/
	TypeProjection VariantType = 19
	/*Variable is of type [Color].*/
	TypeColor VariantType = 20
	/*Variable is of type [StringName].*/
	TypeStringName VariantType = 21
	/*Variable is of type [NodePath].*/
	TypeNodePath VariantType = 22
	/*Variable is of type [RID].*/
	TypeRid VariantType = 23
	/*Variable is of type [Object].*/
	TypeObject VariantType = 24
	/*Variable is of type [Callable].*/
	TypeCallable VariantType = 25
	/*Variable is of type [Signal].*/
	TypeSignal VariantType = 26
	/*Variable is of type [Dictionary].*/
	TypeDictionary VariantType = 27
	/*Variable is of type [Array].*/
	TypeArray VariantType = 28
	/*Variable is of type [PackedByteArray].*/
	TypePackedByteArray VariantType = 29
	/*Variable is of type [PackedInt32Array].*/
	TypePackedInt32Array VariantType = 30
	/*Variable is of type [PackedInt64Array].*/
	TypePackedInt64Array VariantType = 31
	/*Variable is of type [PackedFloat32Array].*/
	TypePackedFloat32Array VariantType = 32
	/*Variable is of type [PackedFloat64Array].*/
	TypePackedFloat64Array VariantType = 33
	/*Variable is of type [PackedStringArray].*/
	TypePackedStringArray VariantType = 34
	/*Variable is of type [PackedVector2Array].*/
	TypePackedVector2Array VariantType = 35
	/*Variable is of type [PackedVector3Array].*/
	TypePackedVector3Array VariantType = 36
	/*Variable is of type [PackedColorArray].*/
	TypePackedColorArray VariantType = 37
	/*Variable is of type [PackedVector4Array].*/
	TypePackedVector4Array VariantType = 38
	/*Represents the size of the [enum Variant.Type] enum.*/
	TypeMax VariantType = 39
)
