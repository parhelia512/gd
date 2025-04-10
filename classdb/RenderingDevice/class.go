// Package RenderingDevice provides methods for working with RenderingDevice object instances.
package RenderingDevice

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
import "graphics.gd/variant/Color"
import "graphics.gd/variant/Dictionary"
import "graphics.gd/variant/Error"
import "graphics.gd/variant/Float"
import "graphics.gd/variant/Object"
import "graphics.gd/variant/Packed"
import "graphics.gd/variant/Path"
import "graphics.gd/variant/RID"
import "graphics.gd/variant/Rect2"
import "graphics.gd/variant/RefCounted"
import "graphics.gd/variant/String"
import "graphics.gd/variant/Vector2i"
import "graphics.gd/variant/Vector3"

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
[RenderingDevice] is an abstraction for working with modern low-level graphics APIs such as Vulkan. Compared to [RenderingServer] (which works with Godot's own rendering subsystems), [RenderingDevice] is much lower-level and allows working more directly with the underlying graphics APIs. [RenderingDevice] is used in Godot to provide support for several modern low-level graphics APIs while reducing the amount of code duplication required. [RenderingDevice] can also be used in your own projects to perform things that are not exposed by [RenderingServer] or high-level nodes, such as using compute shaders.
On startup, Godot creates a global [RenderingDevice] which can be retrieved using [method RenderingServer.get_rendering_device]. This global [RenderingDevice] performs drawing to the screen.
[b]Local RenderingDevices:[/b] Using [method RenderingServer.create_local_rendering_device], you can create "secondary" rendering devices to perform drawing and GPU compute operations on separate threads.
[b]Note:[/b] [RenderingDevice] assumes intermediate knowledge of modern graphics APIs such as Vulkan, Direct3D 12, Metal or WebGPU. These graphics APIs are lower-level than OpenGL or Direct3D 11, requiring you to perform what was previously done by the graphics driver itself. If you have difficulty understanding the concepts used in this class, follow the [url=https://vulkan-tutorial.com/]Vulkan Tutorial[/url] or [url=https://vkguide.dev/]Vulkan Guide[/url]. It's recommended to have existing modern OpenGL or Direct3D 11 knowledge before attempting to learn a low-level graphics API.
[b]Note:[/b] [RenderingDevice] is not available when running in headless mode or when using the Compatibility rendering method.
*/
type Instance [1]gdclass.RenderingDevice
type Expanded [1]gdclass.RenderingDevice

// Nil is a nil/null instance of the class. Equivalent to the zero value.
var Nil Instance

type Any interface {
	gd.IsClass
	AsRenderingDevice() Instance
}

/*
Creates a new texture. It can be accessed with the RID that is returned.
Once finished with your RID, you will want to free the RID using the RenderingDevice's [method free_rid] method.
[b]Note:[/b] Not to be confused with [method RenderingServer.texture_2d_create], which creates the Godot-specific [Texture2D] resource as opposed to the graphics API's own texture type.
*/
func (self Instance) TextureCreate(format [1]gdclass.RDTextureFormat, view [1]gdclass.RDTextureView) RID.Texture { //gd:RenderingDevice.texture_create
	return RID.Texture(Advanced(self).TextureCreate(format, view, gd.ArrayFromSlice[Array.Contains[Packed.Bytes]]([1][][]byte{}[0])))
}

/*
Creates a new texture. It can be accessed with the RID that is returned.
Once finished with your RID, you will want to free the RID using the RenderingDevice's [method free_rid] method.
[b]Note:[/b] Not to be confused with [method RenderingServer.texture_2d_create], which creates the Godot-specific [Texture2D] resource as opposed to the graphics API's own texture type.
*/
func (self Expanded) TextureCreate(format [1]gdclass.RDTextureFormat, view [1]gdclass.RDTextureView, data [][]byte) RID.Texture { //gd:RenderingDevice.texture_create
	return RID.Texture(Advanced(self).TextureCreate(format, view, gd.ArrayFromSlice[Array.Contains[Packed.Bytes]](data)))
}

/*
Creates a shared texture using the specified [param view] and the texture information from [param with_texture].
*/
func (self Instance) TextureCreateShared(view [1]gdclass.RDTextureView, with_texture RID.Texture) RID.Texture { //gd:RenderingDevice.texture_create_shared
	return RID.Texture(Advanced(self).TextureCreateShared(view, RID.Any(with_texture)))
}

/*
Creates a shared texture using the specified [param view] and the texture information from [param with_texture]'s [param layer] and [param mipmap]. The number of included mipmaps from the original texture can be controlled using the [param mipmaps] parameter. Only relevant for textures with multiple layers, such as 3D textures, texture arrays and cubemaps. For single-layer textures, use [method texture_create_shared].
For 2D textures (which only have one layer), [param layer] must be [code]0[/code].
[b]Note:[/b] Layer slicing is only supported for 2D texture arrays, not 3D textures or cubemaps.
*/
func (self Instance) TextureCreateSharedFromSlice(view [1]gdclass.RDTextureView, with_texture RID.Texture, layer int, mipmap int) RID.Texture { //gd:RenderingDevice.texture_create_shared_from_slice
	return RID.Texture(Advanced(self).TextureCreateSharedFromSlice(view, RID.Any(with_texture), int64(layer), int64(mipmap), int64(1), 0))
}

/*
Creates a shared texture using the specified [param view] and the texture information from [param with_texture]'s [param layer] and [param mipmap]. The number of included mipmaps from the original texture can be controlled using the [param mipmaps] parameter. Only relevant for textures with multiple layers, such as 3D textures, texture arrays and cubemaps. For single-layer textures, use [method texture_create_shared].
For 2D textures (which only have one layer), [param layer] must be [code]0[/code].
[b]Note:[/b] Layer slicing is only supported for 2D texture arrays, not 3D textures or cubemaps.
*/
func (self Expanded) TextureCreateSharedFromSlice(view [1]gdclass.RDTextureView, with_texture RID.Texture, layer int, mipmap int, mipmaps int, slice_type gdclass.RenderingDeviceTextureSliceType) RID.Texture { //gd:RenderingDevice.texture_create_shared_from_slice
	return RID.Texture(Advanced(self).TextureCreateSharedFromSlice(view, RID.Any(with_texture), int64(layer), int64(mipmap), int64(mipmaps), slice_type))
}

/*
Returns an RID for an existing [param image] ([code]VkImage[/code]) with the given [param type], [param format], [param samples], [param usage_flags], [param width], [param height], [param depth], and [param layers]. This can be used to allow Godot to render onto foreign images.
*/
func (self Instance) TextureCreateFromExtension(atype gdclass.RenderingDeviceTextureType, format gdclass.RenderingDeviceDataFormat, samples gdclass.RenderingDeviceTextureSamples, usage_flags gdclass.RenderingDeviceTextureUsageBits, image int, width int, height int, depth int, layers int) RID.Texture { //gd:RenderingDevice.texture_create_from_extension
	return RID.Texture(Advanced(self).TextureCreateFromExtension(atype, format, samples, usage_flags, int64(image), int64(width), int64(height), int64(depth), int64(layers)))
}

/*
Updates texture data with new data, replacing the previous data in place. The updated texture data must have the same dimensions and format. For 2D textures (which only have one layer), [param layer] must be [code]0[/code]. Returns [constant @GlobalScope.OK] if the update was successful, [constant @GlobalScope.ERR_INVALID_PARAMETER] otherwise.
[b]Note:[/b] Updating textures is forbidden during creation of a draw or compute list.
[b]Note:[/b] The existing [param texture] can't be updated while a draw list that uses it as part of a framebuffer is being created. Ensure the draw list is finalized (and that the color/depth texture using it is not set to [constant FINAL_ACTION_CONTINUE]) to update this texture.
[b]Note:[/b] The existing [param texture] requires the [constant TEXTURE_USAGE_CAN_UPDATE_BIT] to be updatable.
*/
func (self Instance) TextureUpdate(texture RID.Texture, layer int, data []byte) error { //gd:RenderingDevice.texture_update
	return error(gd.ToError(Advanced(self).TextureUpdate(RID.Any(texture), int64(layer), Packed.Bytes(Packed.New(data...)))))
}

/*
Returns the [param texture] data for the specified [param layer] as raw binary data. For 2D textures (which only have one layer), [param layer] must be [code]0[/code].
[b]Note:[/b] [param texture] can't be retrieved while a draw list that uses it as part of a framebuffer is being created. Ensure the draw list is finalized (and that the color/depth texture using it is not set to [constant FINAL_ACTION_CONTINUE]) to retrieve this texture. Otherwise, an error is printed and a empty [PackedByteArray] is returned.
[b]Note:[/b] [param texture] requires the [constant TEXTURE_USAGE_CAN_COPY_FROM_BIT] to be retrieved. Otherwise, an error is printed and a empty [PackedByteArray] is returned.
[b]Note:[/b] This method will block the GPU from working until the data is retrieved. Refer to [method texture_get_data_async] for an alternative that returns the data in more performant way.
*/
func (self Instance) TextureGetData(texture RID.Texture, layer int) []byte { //gd:RenderingDevice.texture_get_data
	return []byte(Advanced(self).TextureGetData(RID.Any(texture), int64(layer)).Bytes())
}

/*
Asynchronous version of [method texture_get_data]. RenderingDevice will call [param callback] in a certain amount of frames with the data the texture had at the time of the request.
[b]Note:[/b] At the moment, the delay corresponds to the amount of frames specified by [member ProjectSettings.rendering/rendering_device/vsync/frame_queue_size].
[b]Note:[/b] Downloading large textures can have a prohibitive cost for real-time even when using the asynchronous method due to hardware bandwidth limitations. When dealing with large resources, you can adjust settings such as [member ProjectSettings.rendering/rendering_device/staging_buffer/texture_download_region_size_px] and [member ProjectSettings.rendering/rendering_device/staging_buffer/block_size_kb] to improve the transfer speed at the cost of extra memory.
[codeblock]
func _texture_get_data_callback(array):

	value = array.decode_u32(0)

...

rd.texture_get_data_async(texture, 0, _texture_get_data_callback)
[/codeblock]
*/
func (self Instance) TextureGetDataAsync(texture RID.Texture, layer int, callback func(data []byte)) error { //gd:RenderingDevice.texture_get_data_async
	return error(gd.ToError(Advanced(self).TextureGetDataAsync(RID.Any(texture), int64(layer), Callable.New(callback))))
}

/*
Returns [code]true[/code] if the specified [param format] is supported for the given [param usage_flags], [code]false[/code] otherwise.
*/
func (self Instance) TextureIsFormatSupportedForUsage(format gdclass.RenderingDeviceDataFormat, usage_flags gdclass.RenderingDeviceTextureUsageBits) bool { //gd:RenderingDevice.texture_is_format_supported_for_usage
	return bool(Advanced(self).TextureIsFormatSupportedForUsage(format, usage_flags))
}

/*
Returns [code]true[/code] if the [param texture] is shared, [code]false[/code] otherwise. See [RDTextureView].
*/
func (self Instance) TextureIsShared(texture RID.Texture) bool { //gd:RenderingDevice.texture_is_shared
	return bool(Advanced(self).TextureIsShared(RID.Any(texture)))
}

/*
Returns [code]true[/code] if the [param texture] is valid, [code]false[/code] otherwise.
*/
func (self Instance) TextureIsValid(texture RID.Texture) bool { //gd:RenderingDevice.texture_is_valid
	return bool(Advanced(self).TextureIsValid(RID.Any(texture)))
}

/*
Updates the discardable property of [param texture].
If a texture is discardable, its contents do not need to be preserved between frames. This flag is only relevant when the texture is used as target in a draw list.
This information is used by [RenderingDevice] to figure out if a texture's contents can be discarded, eliminating unnecessary writes to memory and boosting performance.
*/
func (self Instance) TextureSetDiscardable(texture RID.Texture, discardable bool) { //gd:RenderingDevice.texture_set_discardable
	Advanced(self).TextureSetDiscardable(RID.Any(texture), discardable)
}

/*
Returns [code]true[/code] if the [param texture] is discardable, [code]false[/code] otherwise. See [RDTextureFormat] or [method texture_set_discardable].
*/
func (self Instance) TextureIsDiscardable(texture RID.Texture) bool { //gd:RenderingDevice.texture_is_discardable
	return bool(Advanced(self).TextureIsDiscardable(RID.Any(texture)))
}

/*
Copies the [param from_texture] to [param to_texture] with the specified [param from_pos], [param to_pos] and [param size] coordinates. The Z axis of the [param from_pos], [param to_pos] and [param size] must be [code]0[/code] for 2-dimensional textures. Source and destination mipmaps/layers must also be specified, with these parameters being [code]0[/code] for textures without mipmaps or single-layer textures. Returns [constant @GlobalScope.OK] if the texture copy was successful or [constant @GlobalScope.ERR_INVALID_PARAMETER] otherwise.
[b]Note:[/b] [param from_texture] texture can't be copied while a draw list that uses it as part of a framebuffer is being created. Ensure the draw list is finalized (and that the color/depth texture using it is not set to [constant FINAL_ACTION_CONTINUE]) to copy this texture.
[b]Note:[/b] [param from_texture] texture requires the [constant TEXTURE_USAGE_CAN_COPY_FROM_BIT] to be retrieved.
[b]Note:[/b] [param to_texture] can't be copied while a draw list that uses it as part of a framebuffer is being created. Ensure the draw list is finalized (and that the color/depth texture using it is not set to [constant FINAL_ACTION_CONTINUE]) to copy this texture.
[b]Note:[/b] [param to_texture] requires the [constant TEXTURE_USAGE_CAN_COPY_TO_BIT] to be retrieved.
[b]Note:[/b] [param from_texture] and [param to_texture] must be of the same type (color or depth).
*/
func (self Instance) TextureCopy(from_texture RID.Texture, to_texture RID.Texture, from_pos Vector3.XYZ, to_pos Vector3.XYZ, size Vector3.XYZ, src_mipmap int, dst_mipmap int, src_layer int, dst_layer int) error { //gd:RenderingDevice.texture_copy
	return error(gd.ToError(Advanced(self).TextureCopy(RID.Any(from_texture), RID.Any(to_texture), Vector3.XYZ(from_pos), Vector3.XYZ(to_pos), Vector3.XYZ(size), int64(src_mipmap), int64(dst_mipmap), int64(src_layer), int64(dst_layer))))
}

/*
Clears the specified [param texture] by replacing all of its pixels with the specified [param color]. [param base_mipmap] and [param mipmap_count] determine which mipmaps of the texture are affected by this clear operation, while [param base_layer] and [param layer_count] determine which layers of a 3D texture (or texture array) are affected by this clear operation. For 2D textures (which only have one layer by design), [param base_layer] must be [code]0[/code] and [param layer_count] must be [code]1[/code].
[b]Note:[/b] [param texture] can't be cleared while a draw list that uses it as part of a framebuffer is being created. Ensure the draw list is finalized (and that the color/depth texture using it is not set to [constant FINAL_ACTION_CONTINUE]) to clear this texture.
*/
func (self Instance) TextureClear(texture RID.Texture, color Color.RGBA, base_mipmap int, mipmap_count int, base_layer int, layer_count int) error { //gd:RenderingDevice.texture_clear
	return error(gd.ToError(Advanced(self).TextureClear(RID.Any(texture), Color.RGBA(color), int64(base_mipmap), int64(mipmap_count), int64(base_layer), int64(layer_count))))
}

/*
Resolves the [param from_texture] texture onto [param to_texture] with multisample antialiasing enabled. This must be used when rendering a framebuffer for MSAA to work. Returns [constant @GlobalScope.OK] if successful, [constant @GlobalScope.ERR_INVALID_PARAMETER] otherwise.
[b]Note:[/b] [param from_texture] and [param to_texture] textures must have the same dimension, format and type (color or depth).
[b]Note:[/b] [param from_texture] can't be copied while a draw list that uses it as part of a framebuffer is being created. Ensure the draw list is finalized (and that the color/depth texture using it is not set to [constant FINAL_ACTION_CONTINUE]) to resolve this texture.
[b]Note:[/b] [param from_texture] requires the [constant TEXTURE_USAGE_CAN_COPY_FROM_BIT] to be retrieved.
[b]Note:[/b] [param from_texture] must be multisampled and must also be 2D (or a slice of a 3D/cubemap texture).
[b]Note:[/b] [param to_texture] can't be copied while a draw list that uses it as part of a framebuffer is being created. Ensure the draw list is finalized (and that the color/depth texture using it is not set to [constant FINAL_ACTION_CONTINUE]) to resolve this texture.
[b]Note:[/b] [param to_texture] texture requires the [constant TEXTURE_USAGE_CAN_COPY_TO_BIT] to be retrieved.
[b]Note:[/b] [param to_texture] texture must [b]not[/b] be multisampled and must also be 2D (or a slice of a 3D/cubemap texture).
*/
func (self Instance) TextureResolveMultisample(from_texture RID.Texture, to_texture RID.Texture) error { //gd:RenderingDevice.texture_resolve_multisample
	return error(gd.ToError(Advanced(self).TextureResolveMultisample(RID.Any(from_texture), RID.Any(to_texture))))
}

/*
Returns the data format used to create this texture.
*/
func (self Instance) TextureGetFormat(texture RID.Texture) [1]gdclass.RDTextureFormat { //gd:RenderingDevice.texture_get_format
	return [1]gdclass.RDTextureFormat(Advanced(self).TextureGetFormat(RID.Any(texture)))
}

/*
Returns the internal graphics handle for this texture object. For use when communicating with third-party APIs mostly with GDExtension.
[b]Note:[/b] This function returns a [code]uint64_t[/code] which internally maps to a [code]GLuint[/code] (OpenGL) or [code]VkImage[/code] (Vulkan).
*/
func (self Instance) TextureGetNativeHandle(texture RID.Texture) int { //gd:RenderingDevice.texture_get_native_handle
	return int(int(Advanced(self).TextureGetNativeHandle(RID.Any(texture))))
}

/*
Creates a new framebuffer format with the specified [param attachments] and [param view_count]. Returns the new framebuffer's unique framebuffer format ID.
If [param view_count] is greater than or equal to [code]2[/code], enables multiview which is used for VR rendering. This requires support for the Vulkan multiview extension.
*/
func (self Instance) FramebufferFormatCreate(attachments [][1]gdclass.RDAttachmentFormat) int { //gd:RenderingDevice.framebuffer_format_create
	return int(int(Advanced(self).FramebufferFormatCreate(gd.ArrayFromSlice[Array.Contains[[1]gdclass.RDAttachmentFormat]](attachments), int64(1))))
}

/*
Creates a new framebuffer format with the specified [param attachments] and [param view_count]. Returns the new framebuffer's unique framebuffer format ID.
If [param view_count] is greater than or equal to [code]2[/code], enables multiview which is used for VR rendering. This requires support for the Vulkan multiview extension.
*/
func (self Expanded) FramebufferFormatCreate(attachments [][1]gdclass.RDAttachmentFormat, view_count int) int { //gd:RenderingDevice.framebuffer_format_create
	return int(int(Advanced(self).FramebufferFormatCreate(gd.ArrayFromSlice[Array.Contains[[1]gdclass.RDAttachmentFormat]](attachments), int64(view_count))))
}

/*
Creates a multipass framebuffer format with the specified [param attachments], [param passes] and [param view_count] and returns its ID. If [param view_count] is greater than or equal to [code]2[/code], enables multiview which is used for VR rendering. This requires support for the Vulkan multiview extension.
*/
func (self Instance) FramebufferFormatCreateMultipass(attachments [][1]gdclass.RDAttachmentFormat, passes [][1]gdclass.RDFramebufferPass) int { //gd:RenderingDevice.framebuffer_format_create_multipass
	return int(int(Advanced(self).FramebufferFormatCreateMultipass(gd.ArrayFromSlice[Array.Contains[[1]gdclass.RDAttachmentFormat]](attachments), gd.ArrayFromSlice[Array.Contains[[1]gdclass.RDFramebufferPass]](passes), int64(1))))
}

/*
Creates a multipass framebuffer format with the specified [param attachments], [param passes] and [param view_count] and returns its ID. If [param view_count] is greater than or equal to [code]2[/code], enables multiview which is used for VR rendering. This requires support for the Vulkan multiview extension.
*/
func (self Expanded) FramebufferFormatCreateMultipass(attachments [][1]gdclass.RDAttachmentFormat, passes [][1]gdclass.RDFramebufferPass, view_count int) int { //gd:RenderingDevice.framebuffer_format_create_multipass
	return int(int(Advanced(self).FramebufferFormatCreateMultipass(gd.ArrayFromSlice[Array.Contains[[1]gdclass.RDAttachmentFormat]](attachments), gd.ArrayFromSlice[Array.Contains[[1]gdclass.RDFramebufferPass]](passes), int64(view_count))))
}

/*
Creates a new empty framebuffer format with the specified number of [param samples] and returns its ID.
*/
func (self Instance) FramebufferFormatCreateEmpty() int { //gd:RenderingDevice.framebuffer_format_create_empty
	return int(int(Advanced(self).FramebufferFormatCreateEmpty(0)))
}

/*
Creates a new empty framebuffer format with the specified number of [param samples] and returns its ID.
*/
func (self Expanded) FramebufferFormatCreateEmpty(samples gdclass.RenderingDeviceTextureSamples) int { //gd:RenderingDevice.framebuffer_format_create_empty
	return int(int(Advanced(self).FramebufferFormatCreateEmpty(samples)))
}

/*
Returns the number of texture samples used for the given framebuffer [param format] ID (returned by [method framebuffer_get_format]).
*/
func (self Instance) FramebufferFormatGetTextureSamples(format int) gdclass.RenderingDeviceTextureSamples { //gd:RenderingDevice.framebuffer_format_get_texture_samples
	return gdclass.RenderingDeviceTextureSamples(Advanced(self).FramebufferFormatGetTextureSamples(int64(format), int64(0)))
}

/*
Returns the number of texture samples used for the given framebuffer [param format] ID (returned by [method framebuffer_get_format]).
*/
func (self Expanded) FramebufferFormatGetTextureSamples(format int, render_pass int) gdclass.RenderingDeviceTextureSamples { //gd:RenderingDevice.framebuffer_format_get_texture_samples
	return gdclass.RenderingDeviceTextureSamples(Advanced(self).FramebufferFormatGetTextureSamples(int64(format), int64(render_pass)))
}

/*
Creates a new framebuffer. It can be accessed with the RID that is returned.
Once finished with your RID, you will want to free the RID using the RenderingDevice's [method free_rid] method.
*/
func (self Instance) FramebufferCreate(textures [][]RID.Texture) RID.Framebuffer { //gd:RenderingDevice.framebuffer_create
	return RID.Framebuffer(Advanced(self).FramebufferCreate(gd.ArrayFromSlice[Array.Contains[RID.Any]](textures), int64(-1), int64(1)))
}

/*
Creates a new framebuffer. It can be accessed with the RID that is returned.
Once finished with your RID, you will want to free the RID using the RenderingDevice's [method free_rid] method.
*/
func (self Expanded) FramebufferCreate(textures [][]RID.Texture, validate_with_format int, view_count int) RID.Framebuffer { //gd:RenderingDevice.framebuffer_create
	return RID.Framebuffer(Advanced(self).FramebufferCreate(gd.ArrayFromSlice[Array.Contains[RID.Any]](textures), int64(validate_with_format), int64(view_count)))
}

/*
Creates a new multipass framebuffer. It can be accessed with the RID that is returned.
Once finished with your RID, you will want to free the RID using the RenderingDevice's [method free_rid] method.
*/
func (self Instance) FramebufferCreateMultipass(textures [][]RID.Texture, passes [][1]gdclass.RDFramebufferPass) RID.Framebuffer { //gd:RenderingDevice.framebuffer_create_multipass
	return RID.Framebuffer(Advanced(self).FramebufferCreateMultipass(gd.ArrayFromSlice[Array.Contains[RID.Any]](textures), gd.ArrayFromSlice[Array.Contains[[1]gdclass.RDFramebufferPass]](passes), int64(-1), int64(1)))
}

/*
Creates a new multipass framebuffer. It can be accessed with the RID that is returned.
Once finished with your RID, you will want to free the RID using the RenderingDevice's [method free_rid] method.
*/
func (self Expanded) FramebufferCreateMultipass(textures [][]RID.Texture, passes [][1]gdclass.RDFramebufferPass, validate_with_format int, view_count int) RID.Framebuffer { //gd:RenderingDevice.framebuffer_create_multipass
	return RID.Framebuffer(Advanced(self).FramebufferCreateMultipass(gd.ArrayFromSlice[Array.Contains[RID.Any]](textures), gd.ArrayFromSlice[Array.Contains[[1]gdclass.RDFramebufferPass]](passes), int64(validate_with_format), int64(view_count)))
}

/*
Creates a new empty framebuffer. It can be accessed with the RID that is returned.
Once finished with your RID, you will want to free the RID using the RenderingDevice's [method free_rid] method.
*/
func (self Instance) FramebufferCreateEmpty(size Vector2i.XY) RID.Framebuffer { //gd:RenderingDevice.framebuffer_create_empty
	return RID.Framebuffer(Advanced(self).FramebufferCreateEmpty(Vector2i.XY(size), 0, int64(-1)))
}

/*
Creates a new empty framebuffer. It can be accessed with the RID that is returned.
Once finished with your RID, you will want to free the RID using the RenderingDevice's [method free_rid] method.
*/
func (self Expanded) FramebufferCreateEmpty(size Vector2i.XY, samples gdclass.RenderingDeviceTextureSamples, validate_with_format int) RID.Framebuffer { //gd:RenderingDevice.framebuffer_create_empty
	return RID.Framebuffer(Advanced(self).FramebufferCreateEmpty(Vector2i.XY(size), samples, int64(validate_with_format)))
}

/*
Returns the format ID of the framebuffer specified by the [param framebuffer] RID. This ID is guaranteed to be unique for the same formats and does not need to be freed.
*/
func (self Instance) FramebufferGetFormat(framebuffer RID.Framebuffer) int { //gd:RenderingDevice.framebuffer_get_format
	return int(int(Advanced(self).FramebufferGetFormat(RID.Any(framebuffer))))
}

/*
Returns [code]true[/code] if the framebuffer specified by the [param framebuffer] RID is valid, [code]false[/code] otherwise.
*/
func (self Instance) FramebufferIsValid(framebuffer RID.Framebuffer) bool { //gd:RenderingDevice.framebuffer_is_valid
	return bool(Advanced(self).FramebufferIsValid(RID.Any(framebuffer)))
}

/*
Creates a new sampler. It can be accessed with the RID that is returned.
Once finished with your RID, you will want to free the RID using the RenderingDevice's [method free_rid] method.
*/
func (self Instance) SamplerCreate(state [1]gdclass.RDSamplerState) RID.Sampler { //gd:RenderingDevice.sampler_create
	return RID.Sampler(Advanced(self).SamplerCreate(state))
}

/*
Returns [code]true[/code] if implementation supports using a texture of [param format] with the given [param sampler_filter].
*/
func (self Instance) SamplerIsFormatSupportedForFilter(format gdclass.RenderingDeviceDataFormat, sampler_filter gdclass.RenderingDeviceSamplerFilter) bool { //gd:RenderingDevice.sampler_is_format_supported_for_filter
	return bool(Advanced(self).SamplerIsFormatSupportedForFilter(format, sampler_filter))
}

/*
It can be accessed with the RID that is returned.
Once finished with your RID, you will want to free the RID using the RenderingDevice's [method free_rid] method.
*/
func (self Instance) VertexBufferCreate(size_bytes int) RID.VertexBuffer { //gd:RenderingDevice.vertex_buffer_create
	return RID.VertexBuffer(Advanced(self).VertexBufferCreate(int64(size_bytes), Packed.Bytes(Packed.New([1][]byte{}[0]...)), 0))
}

/*
It can be accessed with the RID that is returned.
Once finished with your RID, you will want to free the RID using the RenderingDevice's [method free_rid] method.
*/
func (self Expanded) VertexBufferCreate(size_bytes int, data []byte, creation_bits gdclass.RenderingDeviceBufferCreationBits) RID.VertexBuffer { //gd:RenderingDevice.vertex_buffer_create
	return RID.VertexBuffer(Advanced(self).VertexBufferCreate(int64(size_bytes), Packed.Bytes(Packed.New(data...)), creation_bits))
}

/*
Creates a new vertex format with the specified [param vertex_descriptions]. Returns a unique vertex format ID corresponding to the newly created vertex format.
*/
func (self Instance) VertexFormatCreate(vertex_descriptions [][1]gdclass.RDVertexAttribute) int { //gd:RenderingDevice.vertex_format_create
	return int(int(Advanced(self).VertexFormatCreate(gd.ArrayFromSlice[Array.Contains[[1]gdclass.RDVertexAttribute]](vertex_descriptions))))
}

/*
Creates a vertex array based on the specified buffers. Optionally, [param offsets] (in bytes) may be defined for each buffer.
*/
func (self Instance) VertexArrayCreate(vertex_count int, vertex_format int, src_buffers [][]RID.VertexBuffer) RID.VertexArray { //gd:RenderingDevice.vertex_array_create
	return RID.VertexArray(Advanced(self).VertexArrayCreate(int64(vertex_count), int64(vertex_format), gd.ArrayFromSlice[Array.Contains[RID.Any]](src_buffers), Packed.New([1][]int64{}[0]...)))
}

/*
Creates a vertex array based on the specified buffers. Optionally, [param offsets] (in bytes) may be defined for each buffer.
*/
func (self Expanded) VertexArrayCreate(vertex_count int, vertex_format int, src_buffers [][]RID.VertexBuffer, offsets []int64) RID.VertexArray { //gd:RenderingDevice.vertex_array_create
	return RID.VertexArray(Advanced(self).VertexArrayCreate(int64(vertex_count), int64(vertex_format), gd.ArrayFromSlice[Array.Contains[RID.Any]](src_buffers), Packed.New(offsets...)))
}

/*
Creates a new index buffer. It can be accessed with the RID that is returned.
Once finished with your RID, you will want to free the RID using the RenderingDevice's [method free_rid] method.
*/
func (self Instance) IndexBufferCreate(size_indices int, format gdclass.RenderingDeviceIndexBufferFormat) RID.IndexBuffer { //gd:RenderingDevice.index_buffer_create
	return RID.IndexBuffer(Advanced(self).IndexBufferCreate(int64(size_indices), format, Packed.Bytes(Packed.New([1][]byte{}[0]...)), false, 0))
}

/*
Creates a new index buffer. It can be accessed with the RID that is returned.
Once finished with your RID, you will want to free the RID using the RenderingDevice's [method free_rid] method.
*/
func (self Expanded) IndexBufferCreate(size_indices int, format gdclass.RenderingDeviceIndexBufferFormat, data []byte, use_restart_indices bool, creation_bits gdclass.RenderingDeviceBufferCreationBits) RID.IndexBuffer { //gd:RenderingDevice.index_buffer_create
	return RID.IndexBuffer(Advanced(self).IndexBufferCreate(int64(size_indices), format, Packed.Bytes(Packed.New(data...)), use_restart_indices, creation_bits))
}

/*
Creates a new index array. It can be accessed with the RID that is returned.
Once finished with your RID, you will want to free the RID using the RenderingDevice's [method free_rid] method.
*/
func (self Instance) IndexArrayCreate(index_buffer RID.IndexBuffer, index_offset int, index_count int) RID.IndexArray { //gd:RenderingDevice.index_array_create
	return RID.IndexArray(Advanced(self).IndexArrayCreate(RID.Any(index_buffer), int64(index_offset), int64(index_count)))
}

/*
Compiles a SPIR-V from the shader source code in [param shader_source] and returns the SPIR-V as a [RDShaderSPIRV]. This intermediate language shader is portable across different GPU models and driver versions, but cannot be run directly by GPUs until compiled into a binary shader using [method shader_compile_binary_from_spirv].
If [param allow_cache] is [code]true[/code], make use of the shader cache generated by Godot. This avoids a potentially lengthy shader compilation step if the shader is already in cache. If [param allow_cache] is [code]false[/code], Godot's shader cache is ignored and the shader will always be recompiled.
*/
func (self Instance) ShaderCompileSpirvFromSource(shader_source [1]gdclass.RDShaderSource) [1]gdclass.RDShaderSPIRV { //gd:RenderingDevice.shader_compile_spirv_from_source
	return [1]gdclass.RDShaderSPIRV(Advanced(self).ShaderCompileSpirvFromSource(shader_source, true))
}

/*
Compiles a SPIR-V from the shader source code in [param shader_source] and returns the SPIR-V as a [RDShaderSPIRV]. This intermediate language shader is portable across different GPU models and driver versions, but cannot be run directly by GPUs until compiled into a binary shader using [method shader_compile_binary_from_spirv].
If [param allow_cache] is [code]true[/code], make use of the shader cache generated by Godot. This avoids a potentially lengthy shader compilation step if the shader is already in cache. If [param allow_cache] is [code]false[/code], Godot's shader cache is ignored and the shader will always be recompiled.
*/
func (self Expanded) ShaderCompileSpirvFromSource(shader_source [1]gdclass.RDShaderSource, allow_cache bool) [1]gdclass.RDShaderSPIRV { //gd:RenderingDevice.shader_compile_spirv_from_source
	return [1]gdclass.RDShaderSPIRV(Advanced(self).ShaderCompileSpirvFromSource(shader_source, allow_cache))
}

/*
Compiles a binary shader from [param spirv_data] and returns the compiled binary data as a [PackedByteArray]. This compiled shader is specific to the GPU model and driver version used; it will not work on different GPU models or even different driver versions. See also [method shader_compile_spirv_from_source].
[param name] is an optional human-readable name that can be given to the compiled shader for organizational purposes.
*/
func (self Instance) ShaderCompileBinaryFromSpirv(spirv_data [1]gdclass.RDShaderSPIRV) []byte { //gd:RenderingDevice.shader_compile_binary_from_spirv
	return []byte(Advanced(self).ShaderCompileBinaryFromSpirv(spirv_data, String.New("")).Bytes())
}

/*
Compiles a binary shader from [param spirv_data] and returns the compiled binary data as a [PackedByteArray]. This compiled shader is specific to the GPU model and driver version used; it will not work on different GPU models or even different driver versions. See also [method shader_compile_spirv_from_source].
[param name] is an optional human-readable name that can be given to the compiled shader for organizational purposes.
*/
func (self Expanded) ShaderCompileBinaryFromSpirv(spirv_data [1]gdclass.RDShaderSPIRV, name string) []byte { //gd:RenderingDevice.shader_compile_binary_from_spirv
	return []byte(Advanced(self).ShaderCompileBinaryFromSpirv(spirv_data, String.New(name)).Bytes())
}

/*
Creates a new shader instance from SPIR-V intermediate code. It can be accessed with the RID that is returned.
Once finished with your RID, you will want to free the RID using the RenderingDevice's [method free_rid] method. See also [method shader_compile_spirv_from_source] and [method shader_create_from_bytecode].
*/
func (self Instance) ShaderCreateFromSpirv(spirv_data [1]gdclass.RDShaderSPIRV) RID.Shader { //gd:RenderingDevice.shader_create_from_spirv
	return RID.Shader(Advanced(self).ShaderCreateFromSpirv(spirv_data, String.New("")))
}

/*
Creates a new shader instance from SPIR-V intermediate code. It can be accessed with the RID that is returned.
Once finished with your RID, you will want to free the RID using the RenderingDevice's [method free_rid] method. See also [method shader_compile_spirv_from_source] and [method shader_create_from_bytecode].
*/
func (self Expanded) ShaderCreateFromSpirv(spirv_data [1]gdclass.RDShaderSPIRV, name string) RID.Shader { //gd:RenderingDevice.shader_create_from_spirv
	return RID.Shader(Advanced(self).ShaderCreateFromSpirv(spirv_data, String.New(name)))
}

/*
Creates a new shader instance from a binary compiled shader. It can be accessed with the RID that is returned.
Once finished with your RID, you will want to free the RID using the RenderingDevice's [method free_rid] method. See also [method shader_compile_binary_from_spirv] and [method shader_create_from_spirv].
*/
func (self Instance) ShaderCreateFromBytecode(binary_data []byte) RID.Shader { //gd:RenderingDevice.shader_create_from_bytecode
	return RID.Shader(Advanced(self).ShaderCreateFromBytecode(Packed.Bytes(Packed.New(binary_data...)), RID.Any([1]RID.Any{}[0])))
}

/*
Creates a new shader instance from a binary compiled shader. It can be accessed with the RID that is returned.
Once finished with your RID, you will want to free the RID using the RenderingDevice's [method free_rid] method. See also [method shader_compile_binary_from_spirv] and [method shader_create_from_spirv].
*/
func (self Expanded) ShaderCreateFromBytecode(binary_data []byte, placeholder_rid RID.ShaderPlaceholder) RID.Shader { //gd:RenderingDevice.shader_create_from_bytecode
	return RID.Shader(Advanced(self).ShaderCreateFromBytecode(Packed.Bytes(Packed.New(binary_data...)), RID.Any(placeholder_rid)))
}

/*
Create a placeholder RID by allocating an RID without initializing it for use in [method shader_create_from_bytecode]. This allows you to create an RID for a shader and pass it around, but defer compiling the shader to a later time.
*/
func (self Instance) ShaderCreatePlaceholder() RID.ShaderPlaceholder { //gd:RenderingDevice.shader_create_placeholder
	return RID.ShaderPlaceholder(Advanced(self).ShaderCreatePlaceholder())
}

/*
Returns the internal vertex input mask. Internally, the vertex input mask is an unsigned integer consisting of the locations (specified in GLSL via. [code]layout(location = ...)[/code]) of the input variables (specified in GLSL by the [code]in[/code] keyword).
*/
func (self Instance) ShaderGetVertexInputAttributeMask(shader RID.Shader) int { //gd:RenderingDevice.shader_get_vertex_input_attribute_mask
	return int(int(Advanced(self).ShaderGetVertexInputAttributeMask(RID.Any(shader))))
}

/*
Creates a new uniform buffer. It can be accessed with the RID that is returned.
Once finished with your RID, you will want to free the RID using the RenderingDevice's [method free_rid] method.
*/
func (self Instance) UniformBufferCreate(size_bytes int) RID.UniformBuffer { //gd:RenderingDevice.uniform_buffer_create
	return RID.UniformBuffer(Advanced(self).UniformBufferCreate(int64(size_bytes), Packed.Bytes(Packed.New([1][]byte{}[0]...)), 0))
}

/*
Creates a new uniform buffer. It can be accessed with the RID that is returned.
Once finished with your RID, you will want to free the RID using the RenderingDevice's [method free_rid] method.
*/
func (self Expanded) UniformBufferCreate(size_bytes int, data []byte, creation_bits gdclass.RenderingDeviceBufferCreationBits) RID.UniformBuffer { //gd:RenderingDevice.uniform_buffer_create
	return RID.UniformBuffer(Advanced(self).UniformBufferCreate(int64(size_bytes), Packed.Bytes(Packed.New(data...)), creation_bits))
}

/*
Creates a [url=https://vkguide.dev/docs/chapter-4/storage_buffers/]storage buffer[/url] with the specified [param data] and [param usage]. It can be accessed with the RID that is returned.
Once finished with your RID, you will want to free the RID using the RenderingDevice's [method free_rid] method.
*/
func (self Instance) StorageBufferCreate(size_bytes int) RID.StorageBuffer { //gd:RenderingDevice.storage_buffer_create
	return RID.StorageBuffer(Advanced(self).StorageBufferCreate(int64(size_bytes), Packed.Bytes(Packed.New([1][]byte{}[0]...)), 0, 0))
}

/*
Creates a [url=https://vkguide.dev/docs/chapter-4/storage_buffers/]storage buffer[/url] with the specified [param data] and [param usage]. It can be accessed with the RID that is returned.
Once finished with your RID, you will want to free the RID using the RenderingDevice's [method free_rid] method.
*/
func (self Expanded) StorageBufferCreate(size_bytes int, data []byte, usage gdclass.RenderingDeviceStorageBufferUsage, creation_bits gdclass.RenderingDeviceBufferCreationBits) RID.StorageBuffer { //gd:RenderingDevice.storage_buffer_create
	return RID.StorageBuffer(Advanced(self).StorageBufferCreate(int64(size_bytes), Packed.Bytes(Packed.New(data...)), usage, creation_bits))
}

/*
Creates a new texture buffer. It can be accessed with the RID that is returned.
Once finished with your RID, you will want to free the RID using the RenderingDevice's [method free_rid] method.
*/
func (self Instance) TextureBufferCreate(size_bytes int, format gdclass.RenderingDeviceDataFormat) RID.TextureBuffer { //gd:RenderingDevice.texture_buffer_create
	return RID.TextureBuffer(Advanced(self).TextureBufferCreate(int64(size_bytes), format, Packed.Bytes(Packed.New([1][]byte{}[0]...))))
}

/*
Creates a new texture buffer. It can be accessed with the RID that is returned.
Once finished with your RID, you will want to free the RID using the RenderingDevice's [method free_rid] method.
*/
func (self Expanded) TextureBufferCreate(size_bytes int, format gdclass.RenderingDeviceDataFormat, data []byte) RID.TextureBuffer { //gd:RenderingDevice.texture_buffer_create
	return RID.TextureBuffer(Advanced(self).TextureBufferCreate(int64(size_bytes), format, Packed.Bytes(Packed.New(data...))))
}

/*
Creates a new uniform set. It can be accessed with the RID that is returned.
Once finished with your RID, you will want to free the RID using the RenderingDevice's [method free_rid] method.
*/
func (self Instance) UniformSetCreate(uniforms [][1]gdclass.RDUniform, shader RID.Shader, shader_set int) RID.UniformSet { //gd:RenderingDevice.uniform_set_create
	return RID.UniformSet(Advanced(self).UniformSetCreate(gd.ArrayFromSlice[Array.Contains[[1]gdclass.RDUniform]](uniforms), RID.Any(shader), int64(shader_set)))
}

/*
Checks if the [param uniform_set] is valid, i.e. is owned.
*/
func (self Instance) UniformSetIsValid(uniform_set RID.UniformSet) bool { //gd:RenderingDevice.uniform_set_is_valid
	return bool(Advanced(self).UniformSetIsValid(RID.Any(uniform_set)))
}

/*
Copies [param size] bytes from the [param src_buffer] at [param src_offset] into [param dst_buffer] at [param dst_offset].
Prints an error if:
- [param size] exceeds the size of either [param src_buffer] or [param dst_buffer] at their corresponding offsets
- a draw list is currently active (created by [method draw_list_begin])
- a compute list is currently active (created by [method compute_list_begin])
*/
func (self Instance) BufferCopy(src_buffer RID.Buffer, dst_buffer RID.Buffer, src_offset int, dst_offset int, size int) error { //gd:RenderingDevice.buffer_copy
	return error(gd.ToError(Advanced(self).BufferCopy(RID.Any(src_buffer), RID.Any(dst_buffer), int64(src_offset), int64(dst_offset), int64(size))))
}

/*
Updates a region of [param size_bytes] bytes, starting at [param offset], in the buffer, with the specified [param data].
Prints an error if:
- the region specified by [param offset] + [param size_bytes] exceeds the buffer
- a draw list is currently active (created by [method draw_list_begin])
- a compute list is currently active (created by [method compute_list_begin])
*/
func (self Instance) BufferUpdate(buffer RID.Buffer, offset int, size_bytes int, data []byte) error { //gd:RenderingDevice.buffer_update
	return error(gd.ToError(Advanced(self).BufferUpdate(RID.Any(buffer), int64(offset), int64(size_bytes), Packed.Bytes(Packed.New(data...)))))
}

/*
Clears the contents of the [param buffer], clearing [param size_bytes] bytes, starting at [param offset].
Prints an error if:
- the size isn't a multiple of four
- the region specified by [param offset] + [param size_bytes] exceeds the buffer
- a draw list is currently active (created by [method draw_list_begin])
- a compute list is currently active (created by [method compute_list_begin])
*/
func (self Instance) BufferClear(buffer RID.Buffer, offset int, size_bytes int) error { //gd:RenderingDevice.buffer_clear
	return error(gd.ToError(Advanced(self).BufferClear(RID.Any(buffer), int64(offset), int64(size_bytes))))
}

/*
Returns a copy of the data of the specified [param buffer], optionally [param offset_bytes] and [param size_bytes] can be set to copy only a portion of the buffer.
[b]Note:[/b] This method will block the GPU from working until the data is retrieved. Refer to [method buffer_get_data_async] for an alternative that returns the data in more performant way.
*/
func (self Instance) BufferGetData(buffer RID.Buffer) []byte { //gd:RenderingDevice.buffer_get_data
	return []byte(Advanced(self).BufferGetData(RID.Any(buffer), int64(0), int64(0)).Bytes())
}

/*
Returns a copy of the data of the specified [param buffer], optionally [param offset_bytes] and [param size_bytes] can be set to copy only a portion of the buffer.
[b]Note:[/b] This method will block the GPU from working until the data is retrieved. Refer to [method buffer_get_data_async] for an alternative that returns the data in more performant way.
*/
func (self Expanded) BufferGetData(buffer RID.Buffer, offset_bytes int, size_bytes int) []byte { //gd:RenderingDevice.buffer_get_data
	return []byte(Advanced(self).BufferGetData(RID.Any(buffer), int64(offset_bytes), int64(size_bytes)).Bytes())
}

/*
Asynchronous version of [method buffer_get_data]. RenderingDevice will call [param callback] in a certain amount of frames with the data the buffer had at the time of the request.
[b]Note:[/b] At the moment, the delay corresponds to the amount of frames specified by [member ProjectSettings.rendering/rendering_device/vsync/frame_queue_size].
[b]Note:[/b] Downloading large buffers can have a prohibitive cost for real-time even when using the asynchronous method due to hardware bandwidth limitations. When dealing with large resources, you can adjust settings such as [member ProjectSettings.rendering/rendering_device/staging_buffer/block_size_kb] to improve the transfer speed at the cost of extra memory.
[codeblock]
func _buffer_get_data_callback(array):

	value = array.decode_u32(0)

...

rd.buffer_get_data_async(buffer, _buffer_get_data_callback)
[/codeblock]
*/
func (self Instance) BufferGetDataAsync(buffer RID.Buffer, callback func(data []byte)) error { //gd:RenderingDevice.buffer_get_data_async
	return error(gd.ToError(Advanced(self).BufferGetDataAsync(RID.Any(buffer), Callable.New(callback), int64(0), int64(0))))
}

/*
Asynchronous version of [method buffer_get_data]. RenderingDevice will call [param callback] in a certain amount of frames with the data the buffer had at the time of the request.
[b]Note:[/b] At the moment, the delay corresponds to the amount of frames specified by [member ProjectSettings.rendering/rendering_device/vsync/frame_queue_size].
[b]Note:[/b] Downloading large buffers can have a prohibitive cost for real-time even when using the asynchronous method due to hardware bandwidth limitations. When dealing with large resources, you can adjust settings such as [member ProjectSettings.rendering/rendering_device/staging_buffer/block_size_kb] to improve the transfer speed at the cost of extra memory.
[codeblock]
func _buffer_get_data_callback(array):

	value = array.decode_u32(0)

...

rd.buffer_get_data_async(buffer, _buffer_get_data_callback)
[/codeblock]
*/
func (self Expanded) BufferGetDataAsync(buffer RID.Buffer, callback func(data []byte), offset_bytes int, size_bytes int) error { //gd:RenderingDevice.buffer_get_data_async
	return error(gd.ToError(Advanced(self).BufferGetDataAsync(RID.Any(buffer), Callable.New(callback), int64(offset_bytes), int64(size_bytes))))
}

/*
Returns the address of the given [param buffer] which can be passed to shaders in any way to access underlying data. Buffer must have been created with this feature enabled.
[b]Note:[/b] You must check that the GPU supports this functionality by calling [method has_feature] with [constant SUPPORTS_BUFFER_DEVICE_ADDRESS] as a parameter.
*/
func (self Instance) BufferGetDeviceAddress(buffer RID.Buffer) int { //gd:RenderingDevice.buffer_get_device_address
	return int(int(Advanced(self).BufferGetDeviceAddress(RID.Any(buffer))))
}

/*
Creates a new render pipeline. It can be accessed with the RID that is returned.
Once finished with your RID, you will want to free the RID using the RenderingDevice's [method free_rid] method.
*/
func (self Instance) RenderPipelineCreate(shader RID.Shader, framebuffer_format int, vertex_format int, primitive gdclass.RenderingDeviceRenderPrimitive, rasterization_state [1]gdclass.RDPipelineRasterizationState, multisample_state [1]gdclass.RDPipelineMultisampleState, stencil_state [1]gdclass.RDPipelineDepthStencilState, color_blend_state [1]gdclass.RDPipelineColorBlendState) RID.RenderPipeline { //gd:RenderingDevice.render_pipeline_create
	return RID.RenderPipeline(Advanced(self).RenderPipelineCreate(RID.Any(shader), int64(framebuffer_format), int64(vertex_format), primitive, rasterization_state, multisample_state, stencil_state, color_blend_state, 0, int64(0), gd.ArrayFromSlice[Array.Contains[[1]gdclass.RDPipelineSpecializationConstant]]([1][][1]gdclass.RDPipelineSpecializationConstant{}[0])))
}

/*
Creates a new render pipeline. It can be accessed with the RID that is returned.
Once finished with your RID, you will want to free the RID using the RenderingDevice's [method free_rid] method.
*/
func (self Expanded) RenderPipelineCreate(shader RID.Shader, framebuffer_format int, vertex_format int, primitive gdclass.RenderingDeviceRenderPrimitive, rasterization_state [1]gdclass.RDPipelineRasterizationState, multisample_state [1]gdclass.RDPipelineMultisampleState, stencil_state [1]gdclass.RDPipelineDepthStencilState, color_blend_state [1]gdclass.RDPipelineColorBlendState, dynamic_state_flags gdclass.RenderingDevicePipelineDynamicStateFlags, for_render_pass int, specialization_constants [][1]gdclass.RDPipelineSpecializationConstant) RID.RenderPipeline { //gd:RenderingDevice.render_pipeline_create
	return RID.RenderPipeline(Advanced(self).RenderPipelineCreate(RID.Any(shader), int64(framebuffer_format), int64(vertex_format), primitive, rasterization_state, multisample_state, stencil_state, color_blend_state, dynamic_state_flags, int64(for_render_pass), gd.ArrayFromSlice[Array.Contains[[1]gdclass.RDPipelineSpecializationConstant]](specialization_constants)))
}

/*
Returns [code]true[/code] if the render pipeline specified by the [param render_pipeline] RID is valid, [code]false[/code] otherwise.
*/
func (self Instance) RenderPipelineIsValid(render_pipeline RID.RenderPipeline) bool { //gd:RenderingDevice.render_pipeline_is_valid
	return bool(Advanced(self).RenderPipelineIsValid(RID.Any(render_pipeline)))
}

/*
Creates a new compute pipeline. It can be accessed with the RID that is returned.
Once finished with your RID, you will want to free the RID using the RenderingDevice's [method free_rid] method.
*/
func (self Instance) ComputePipelineCreate(shader RID.Shader) RID.ComputePipeline { //gd:RenderingDevice.compute_pipeline_create
	return RID.ComputePipeline(Advanced(self).ComputePipelineCreate(RID.Any(shader), gd.ArrayFromSlice[Array.Contains[[1]gdclass.RDPipelineSpecializationConstant]]([1][][1]gdclass.RDPipelineSpecializationConstant{}[0])))
}

/*
Creates a new compute pipeline. It can be accessed with the RID that is returned.
Once finished with your RID, you will want to free the RID using the RenderingDevice's [method free_rid] method.
*/
func (self Expanded) ComputePipelineCreate(shader RID.Shader, specialization_constants [][1]gdclass.RDPipelineSpecializationConstant) RID.ComputePipeline { //gd:RenderingDevice.compute_pipeline_create
	return RID.ComputePipeline(Advanced(self).ComputePipelineCreate(RID.Any(shader), gd.ArrayFromSlice[Array.Contains[[1]gdclass.RDPipelineSpecializationConstant]](specialization_constants)))
}

/*
Returns [code]true[/code] if the compute pipeline specified by the [param compute_pipeline] RID is valid, [code]false[/code] otherwise.
*/
func (self Instance) ComputePipelineIsValid(compute_pipeline RID.ComputePipeline) bool { //gd:RenderingDevice.compute_pipeline_is_valid
	return bool(Advanced(self).ComputePipelineIsValid(RID.Any(compute_pipeline)))
}

/*
Returns the window width matching the graphics API context for the given window ID (in pixels). Despite the parameter being named [param screen], this returns the [i]window[/i] size. See also [method screen_get_height].
[b]Note:[/b] Only the main [RenderingDevice] returned by [method RenderingServer.get_rendering_device] has a width. If called on a local [RenderingDevice], this method prints an error and returns [constant INVALID_ID].
*/
func (self Instance) ScreenGetWidth() int { //gd:RenderingDevice.screen_get_width
	return int(int(Advanced(self).ScreenGetWidth(int64(0))))
}

/*
Returns the window width matching the graphics API context for the given window ID (in pixels). Despite the parameter being named [param screen], this returns the [i]window[/i] size. See also [method screen_get_height].
[b]Note:[/b] Only the main [RenderingDevice] returned by [method RenderingServer.get_rendering_device] has a width. If called on a local [RenderingDevice], this method prints an error and returns [constant INVALID_ID].
*/
func (self Expanded) ScreenGetWidth(screen int) int { //gd:RenderingDevice.screen_get_width
	return int(int(Advanced(self).ScreenGetWidth(int64(screen))))
}

/*
Returns the window height matching the graphics API context for the given window ID (in pixels). Despite the parameter being named [param screen], this returns the [i]window[/i] size. See also [method screen_get_width].
[b]Note:[/b] Only the main [RenderingDevice] returned by [method RenderingServer.get_rendering_device] has a height. If called on a local [RenderingDevice], this method prints an error and returns [constant INVALID_ID].
*/
func (self Instance) ScreenGetHeight() int { //gd:RenderingDevice.screen_get_height
	return int(int(Advanced(self).ScreenGetHeight(int64(0))))
}

/*
Returns the window height matching the graphics API context for the given window ID (in pixels). Despite the parameter being named [param screen], this returns the [i]window[/i] size. See also [method screen_get_width].
[b]Note:[/b] Only the main [RenderingDevice] returned by [method RenderingServer.get_rendering_device] has a height. If called on a local [RenderingDevice], this method prints an error and returns [constant INVALID_ID].
*/
func (self Expanded) ScreenGetHeight(screen int) int { //gd:RenderingDevice.screen_get_height
	return int(int(Advanced(self).ScreenGetHeight(int64(screen))))
}

/*
Returns the framebuffer format of the given screen.
[b]Note:[/b] Only the main [RenderingDevice] returned by [method RenderingServer.get_rendering_device] has a format. If called on a local [RenderingDevice], this method prints an error and returns [constant INVALID_ID].
*/
func (self Instance) ScreenGetFramebufferFormat() int { //gd:RenderingDevice.screen_get_framebuffer_format
	return int(int(Advanced(self).ScreenGetFramebufferFormat(int64(0))))
}

/*
Returns the framebuffer format of the given screen.
[b]Note:[/b] Only the main [RenderingDevice] returned by [method RenderingServer.get_rendering_device] has a format. If called on a local [RenderingDevice], this method prints an error and returns [constant INVALID_ID].
*/
func (self Expanded) ScreenGetFramebufferFormat(screen int) int { //gd:RenderingDevice.screen_get_framebuffer_format
	return int(int(Advanced(self).ScreenGetFramebufferFormat(int64(screen))))
}

/*
High-level variant of [method draw_list_begin], with the parameters automatically being adjusted for drawing onto the window specified by the [param screen] ID.
[b]Note:[/b] Cannot be used with local RenderingDevices, as these don't have a screen. If called on a local RenderingDevice, [method draw_list_begin_for_screen] returns [constant INVALID_ID].
*/
func (self Instance) DrawListBeginForScreen() int { //gd:RenderingDevice.draw_list_begin_for_screen
	return int(int(Advanced(self).DrawListBeginForScreen(int64(0), Color.RGBA(gd.Color{0, 0, 0, 1}))))
}

/*
High-level variant of [method draw_list_begin], with the parameters automatically being adjusted for drawing onto the window specified by the [param screen] ID.
[b]Note:[/b] Cannot be used with local RenderingDevices, as these don't have a screen. If called on a local RenderingDevice, [method draw_list_begin_for_screen] returns [constant INVALID_ID].
*/
func (self Expanded) DrawListBeginForScreen(screen int, clear_color Color.RGBA) int { //gd:RenderingDevice.draw_list_begin_for_screen
	return int(int(Advanced(self).DrawListBeginForScreen(int64(screen), Color.RGBA(clear_color))))
}

/*
Starts a list of raster drawing commands created with the [code]draw_*[/code] methods. The returned value should be passed to other [code]draw_list_*[/code] functions.
Multiple draw lists cannot be created at the same time; you must finish the previous draw list first using [method draw_list_end].
A simple drawing operation might look like this (code is not a complete example):
[codeblock]
var rd = RenderingDevice.new()
var clear_colors = PackedColorArray([Color(0, 0, 0, 0), Color(0, 0, 0, 0), Color(0, 0, 0, 0)])
var draw_list = rd.draw_list_begin(framebuffers[i], RenderingDevice.CLEAR_COLOR_ALL, clear_colors, true, 1.0f, true, 0, Rect2(), RenderingDevice.OPAQUE_PASS)

# Draw opaque.
rd.draw_list_bind_render_pipeline(draw_list, raster_pipeline)
rd.draw_list_bind_uniform_set(draw_list, raster_base_uniform, 0)
rd.draw_list_set_push_constant(draw_list, raster_push_constant, raster_push_constant.size())
rd.draw_list_draw(draw_list, false, 1, slice_triangle_count[i] * 3)
# Draw wire.
rd.draw_list_bind_render_pipeline(draw_list, raster_pipeline_wire)
rd.draw_list_bind_uniform_set(draw_list, raster_base_uniform, 0)
rd.draw_list_set_push_constant(draw_list, raster_push_constant, raster_push_constant.size())
rd.draw_list_draw(draw_list, false, 1, slice_triangle_count[i] * 3)

rd.draw_list_end()
[/codeblock]
The [param draw_flags] indicates if the texture attachments of the framebuffer should be cleared or ignored. Only one of the two flags can be used for each individual attachment. Ignoring an attachment means that any contents that existed before the draw list will be completely discarded, reducing the memory bandwidth used by the render pass but producing garbage results if the pixels aren't replaced. The default behavior allows the engine to figure out the right operation to use if the texture is discardable, which can result in increased performance. See [RDTextureFormat] or [method texture_set_discardable].
The [param breadcrumb] parameter can be an arbitrary 32-bit integer that is useful to diagnose GPU crashes. If Godot is built in dev or debug mode; when the GPU crashes Godot will dump all shaders that were being executed at the time of the crash and the breadcrumb is useful to diagnose what passes did those shaders belong to.
It does not affect rendering behavior and can be set to 0. It is recommended to use [enum BreadcrumbMarker] enumerations for consistency but it's not required. It is also possible to use bitwise operations to add extra data. e.g.
[codeblock]
rd.draw_list_begin(fb[i], RenderingDevice.CLEAR_COLOR_ALL, clear_colors, true, 1.0f, true, 0, Rect2(), RenderingDevice.OPAQUE_PASS | 5)
[/codeblock]
*/
func (self Instance) DrawListBegin(framebuffer RID.Framebuffer) int { //gd:RenderingDevice.draw_list_begin
	return int(int(Advanced(self).DrawListBegin(RID.Any(framebuffer), 0, Packed.New([1][]Color.RGBA{}[0]...), float64(1.0), int64(0), Rect2.PositionSize(gd.NewRect2(0, 0, 0, 0)), int64(0))))
}

/*
Starts a list of raster drawing commands created with the [code]draw_*[/code] methods. The returned value should be passed to other [code]draw_list_*[/code] functions.
Multiple draw lists cannot be created at the same time; you must finish the previous draw list first using [method draw_list_end].
A simple drawing operation might look like this (code is not a complete example):
[codeblock]
var rd = RenderingDevice.new()
var clear_colors = PackedColorArray([Color(0, 0, 0, 0), Color(0, 0, 0, 0), Color(0, 0, 0, 0)])
var draw_list = rd.draw_list_begin(framebuffers[i], RenderingDevice.CLEAR_COLOR_ALL, clear_colors, true, 1.0f, true, 0, Rect2(), RenderingDevice.OPAQUE_PASS)

# Draw opaque.
rd.draw_list_bind_render_pipeline(draw_list, raster_pipeline)
rd.draw_list_bind_uniform_set(draw_list, raster_base_uniform, 0)
rd.draw_list_set_push_constant(draw_list, raster_push_constant, raster_push_constant.size())
rd.draw_list_draw(draw_list, false, 1, slice_triangle_count[i] * 3)
# Draw wire.
rd.draw_list_bind_render_pipeline(draw_list, raster_pipeline_wire)
rd.draw_list_bind_uniform_set(draw_list, raster_base_uniform, 0)
rd.draw_list_set_push_constant(draw_list, raster_push_constant, raster_push_constant.size())
rd.draw_list_draw(draw_list, false, 1, slice_triangle_count[i] * 3)

rd.draw_list_end()
[/codeblock]
The [param draw_flags] indicates if the texture attachments of the framebuffer should be cleared or ignored. Only one of the two flags can be used for each individual attachment. Ignoring an attachment means that any contents that existed before the draw list will be completely discarded, reducing the memory bandwidth used by the render pass but producing garbage results if the pixels aren't replaced. The default behavior allows the engine to figure out the right operation to use if the texture is discardable, which can result in increased performance. See [RDTextureFormat] or [method texture_set_discardable].
The [param breadcrumb] parameter can be an arbitrary 32-bit integer that is useful to diagnose GPU crashes. If Godot is built in dev or debug mode; when the GPU crashes Godot will dump all shaders that were being executed at the time of the crash and the breadcrumb is useful to diagnose what passes did those shaders belong to.
It does not affect rendering behavior and can be set to 0. It is recommended to use [enum BreadcrumbMarker] enumerations for consistency but it's not required. It is also possible to use bitwise operations to add extra data. e.g.
[codeblock]
rd.draw_list_begin(fb[i], RenderingDevice.CLEAR_COLOR_ALL, clear_colors, true, 1.0f, true, 0, Rect2(), RenderingDevice.OPAQUE_PASS | 5)
[/codeblock]
*/
func (self Expanded) DrawListBegin(framebuffer RID.Framebuffer, draw_flags gdclass.RenderingDeviceDrawFlags, clear_color_values []Color.RGBA, clear_depth_value Float.X, clear_stencil_value int, region Rect2.PositionSize, breadcrumb int) int { //gd:RenderingDevice.draw_list_begin
	return int(int(Advanced(self).DrawListBegin(RID.Any(framebuffer), draw_flags, Packed.New(clear_color_values...), float64(clear_depth_value), int64(clear_stencil_value), Rect2.PositionSize(region), int64(breadcrumb))))
}

/*
This method does nothing and always returns an empty [PackedInt64Array].
*/
func (self Instance) DrawListBeginSplit(framebuffer RID.Framebuffer, splits int, initial_color_action gdclass.RenderingDeviceInitialAction, final_color_action gdclass.RenderingDeviceFinalAction, initial_depth_action gdclass.RenderingDeviceInitialAction, final_depth_action gdclass.RenderingDeviceFinalAction) []int64 { //gd:RenderingDevice.draw_list_begin_split
	return []int64(slices.Collect(Advanced(self).DrawListBeginSplit(RID.Any(framebuffer), int64(splits), initial_color_action, final_color_action, initial_depth_action, final_depth_action, Packed.New([1][]Color.RGBA{}[0]...), float64(1.0), int64(0), Rect2.PositionSize(gd.NewRect2(0, 0, 0, 0)), gd.ArrayFromSlice[Array.Contains[RID.Any]]([1][]RID.Any{}[0])).Values()))
}

/*
This method does nothing and always returns an empty [PackedInt64Array].
*/
func (self Expanded) DrawListBeginSplit(framebuffer RID.Framebuffer, splits int, initial_color_action gdclass.RenderingDeviceInitialAction, final_color_action gdclass.RenderingDeviceFinalAction, initial_depth_action gdclass.RenderingDeviceInitialAction, final_depth_action gdclass.RenderingDeviceFinalAction, clear_color_values []Color.RGBA, clear_depth Float.X, clear_stencil int, region Rect2.PositionSize, storage_textures [][]RID.Texture) []int64 { //gd:RenderingDevice.draw_list_begin_split
	return []int64(slices.Collect(Advanced(self).DrawListBeginSplit(RID.Any(framebuffer), int64(splits), initial_color_action, final_color_action, initial_depth_action, final_depth_action, Packed.New(clear_color_values...), float64(clear_depth), int64(clear_stencil), Rect2.PositionSize(region), gd.ArrayFromSlice[Array.Contains[RID.Any]](storage_textures)).Values()))
}

/*
Sets blend constants for the specified [param draw_list] to [param color]. Blend constants are used only if the graphics pipeline is created with [constant DYNAMIC_STATE_BLEND_CONSTANTS] flag set.
*/
func (self Instance) DrawListSetBlendConstants(draw_list int, color Color.RGBA) { //gd:RenderingDevice.draw_list_set_blend_constants
	Advanced(self).DrawListSetBlendConstants(int64(draw_list), Color.RGBA(color))
}

/*
Binds [param render_pipeline] to the specified [param draw_list].
*/
func (self Instance) DrawListBindRenderPipeline(draw_list int, render_pipeline RID.RenderPipeline) { //gd:RenderingDevice.draw_list_bind_render_pipeline
	Advanced(self).DrawListBindRenderPipeline(int64(draw_list), RID.Any(render_pipeline))
}

/*
Binds [param uniform_set] to the specified [param draw_list]. A [param set_index] must also be specified, which is an identifier starting from [code]0[/code] that must match the one expected by the draw list.
*/
func (self Instance) DrawListBindUniformSet(draw_list int, uniform_set RID.UniformSet, set_index int) { //gd:RenderingDevice.draw_list_bind_uniform_set
	Advanced(self).DrawListBindUniformSet(int64(draw_list), RID.Any(uniform_set), int64(set_index))
}

/*
Binds [param vertex_array] to the specified [param draw_list].
*/
func (self Instance) DrawListBindVertexArray(draw_list int, vertex_array RID.VertexArray) { //gd:RenderingDevice.draw_list_bind_vertex_array
	Advanced(self).DrawListBindVertexArray(int64(draw_list), RID.Any(vertex_array))
}

/*
Binds [param index_array] to the specified [param draw_list].
*/
func (self Instance) DrawListBindIndexArray(draw_list int, index_array RID.IndexArray) { //gd:RenderingDevice.draw_list_bind_index_array
	Advanced(self).DrawListBindIndexArray(int64(draw_list), RID.Any(index_array))
}

/*
Sets the push constant data to [param buffer] for the specified [param draw_list]. The shader determines how this binary data is used. The buffer's size in bytes must also be specified in [param size_bytes] (this can be obtained by calling the [method PackedByteArray.size] method on the passed [param buffer]).
*/
func (self Instance) DrawListSetPushConstant(draw_list int, buffer []byte, size_bytes int) { //gd:RenderingDevice.draw_list_set_push_constant
	Advanced(self).DrawListSetPushConstant(int64(draw_list), Packed.Bytes(Packed.New(buffer...)), int64(size_bytes))
}

/*
Submits [param draw_list] for rendering on the GPU. This is the raster equivalent to [method compute_list_dispatch].
*/
func (self Instance) DrawListDraw(draw_list int, use_indices bool, instances int) { //gd:RenderingDevice.draw_list_draw
	Advanced(self).DrawListDraw(int64(draw_list), use_indices, int64(instances), int64(0))
}

/*
Submits [param draw_list] for rendering on the GPU. This is the raster equivalent to [method compute_list_dispatch].
*/
func (self Expanded) DrawListDraw(draw_list int, use_indices bool, instances int, procedural_vertex_count int) { //gd:RenderingDevice.draw_list_draw
	Advanced(self).DrawListDraw(int64(draw_list), use_indices, int64(instances), int64(procedural_vertex_count))
}

/*
Submits [param draw_list] for rendering on the GPU with the given parameters stored in the [param buffer] at [param offset]. Parameters being integers: vertex count, instance count, first vertex, first instance. And when using indices: index count, instance count, first index, vertex offset, first instance. Buffer must have been created with [constant STORAGE_BUFFER_USAGE_DISPATCH_INDIRECT] flag.
*/
func (self Instance) DrawListDrawIndirect(draw_list int, use_indices bool, buffer RID.Buffer) { //gd:RenderingDevice.draw_list_draw_indirect
	Advanced(self).DrawListDrawIndirect(int64(draw_list), use_indices, RID.Any(buffer), int64(0), int64(1), int64(0))
}

/*
Submits [param draw_list] for rendering on the GPU with the given parameters stored in the [param buffer] at [param offset]. Parameters being integers: vertex count, instance count, first vertex, first instance. And when using indices: index count, instance count, first index, vertex offset, first instance. Buffer must have been created with [constant STORAGE_BUFFER_USAGE_DISPATCH_INDIRECT] flag.
*/
func (self Expanded) DrawListDrawIndirect(draw_list int, use_indices bool, buffer RID.Buffer, offset int, draw_count int, stride int) { //gd:RenderingDevice.draw_list_draw_indirect
	Advanced(self).DrawListDrawIndirect(int64(draw_list), use_indices, RID.Any(buffer), int64(offset), int64(draw_count), int64(stride))
}

/*
Creates a scissor rectangle and enables it for the specified [param draw_list]. Scissor rectangles are used for clipping by discarding fragments that fall outside a specified rectangular portion of the screen. See also [method draw_list_disable_scissor].
[b]Note:[/b] The specified [param rect] is automatically intersected with the screen's dimensions, which means it cannot exceed the screen's dimensions.
*/
func (self Instance) DrawListEnableScissor(draw_list int) { //gd:RenderingDevice.draw_list_enable_scissor
	Advanced(self).DrawListEnableScissor(int64(draw_list), Rect2.PositionSize(gd.NewRect2(0, 0, 0, 0)))
}

/*
Creates a scissor rectangle and enables it for the specified [param draw_list]. Scissor rectangles are used for clipping by discarding fragments that fall outside a specified rectangular portion of the screen. See also [method draw_list_disable_scissor].
[b]Note:[/b] The specified [param rect] is automatically intersected with the screen's dimensions, which means it cannot exceed the screen's dimensions.
*/
func (self Expanded) DrawListEnableScissor(draw_list int, rect Rect2.PositionSize) { //gd:RenderingDevice.draw_list_enable_scissor
	Advanced(self).DrawListEnableScissor(int64(draw_list), Rect2.PositionSize(rect))
}

/*
Removes and disables the scissor rectangle for the specified [param draw_list]. See also [method draw_list_enable_scissor].
*/
func (self Instance) DrawListDisableScissor(draw_list int) { //gd:RenderingDevice.draw_list_disable_scissor
	Advanced(self).DrawListDisableScissor(int64(draw_list))
}

/*
Switches to the next draw pass.
*/
func (self Instance) DrawListSwitchToNextPass() int { //gd:RenderingDevice.draw_list_switch_to_next_pass
	return int(int(Advanced(self).DrawListSwitchToNextPass()))
}

/*
This method does nothing and always returns an empty [PackedInt64Array].
*/
func (self Instance) DrawListSwitchToNextPassSplit(splits int) []int64 { //gd:RenderingDevice.draw_list_switch_to_next_pass_split
	return []int64(slices.Collect(Advanced(self).DrawListSwitchToNextPassSplit(int64(splits)).Values()))
}

/*
Finishes a list of raster drawing commands created with the [code]draw_*[/code] methods.
*/
func (self Instance) DrawListEnd() { //gd:RenderingDevice.draw_list_end
	Advanced(self).DrawListEnd()
}

/*
Starts a list of compute commands created with the [code]compute_*[/code] methods. The returned value should be passed to other [code]compute_list_*[/code] functions.
Multiple compute lists cannot be created at the same time; you must finish the previous compute list first using [method compute_list_end].
A simple compute operation might look like this (code is not a complete example):
[codeblock]
var rd = RenderingDevice.new()
var compute_list = rd.compute_list_begin()

rd.compute_list_bind_compute_pipeline(compute_list, compute_shader_dilate_pipeline)
rd.compute_list_bind_uniform_set(compute_list, compute_base_uniform_set, 0)
rd.compute_list_bind_uniform_set(compute_list, dilate_uniform_set, 1)

for i in atlas_slices:

	rd.compute_list_set_push_constant(compute_list, push_constant, push_constant.size())
	rd.compute_list_dispatch(compute_list, group_size.x, group_size.y, group_size.z)
	# No barrier, let them run all together.

rd.compute_list_end()
[/codeblock]
*/
func (self Instance) ComputeListBegin() int { //gd:RenderingDevice.compute_list_begin
	return int(int(Advanced(self).ComputeListBegin()))
}

/*
Tells the GPU what compute pipeline to use when processing the compute list. If the shader has changed since the last time this function was called, Godot will unbind all descriptor sets and will re-bind them inside [method compute_list_dispatch].
*/
func (self Instance) ComputeListBindComputePipeline(compute_list int, compute_pipeline RID.ComputePipeline) { //gd:RenderingDevice.compute_list_bind_compute_pipeline
	Advanced(self).ComputeListBindComputePipeline(int64(compute_list), RID.Any(compute_pipeline))
}

/*
Sets the push constant data to [param buffer] for the specified [param compute_list]. The shader determines how this binary data is used. The buffer's size in bytes must also be specified in [param size_bytes] (this can be obtained by calling the [method PackedByteArray.size] method on the passed [param buffer]).
*/
func (self Instance) ComputeListSetPushConstant(compute_list int, buffer []byte, size_bytes int) { //gd:RenderingDevice.compute_list_set_push_constant
	Advanced(self).ComputeListSetPushConstant(int64(compute_list), Packed.Bytes(Packed.New(buffer...)), int64(size_bytes))
}

/*
Binds the [param uniform_set] to this [param compute_list]. Godot ensures that all textures in the uniform set have the correct Vulkan access masks. If Godot had to change access masks of textures, it will raise a Vulkan image memory barrier.
*/
func (self Instance) ComputeListBindUniformSet(compute_list int, uniform_set RID.UniformSet, set_index int) { //gd:RenderingDevice.compute_list_bind_uniform_set
	Advanced(self).ComputeListBindUniformSet(int64(compute_list), RID.Any(uniform_set), int64(set_index))
}

/*
Submits the compute list for processing on the GPU. This is the compute equivalent to [method draw_list_draw].
*/
func (self Instance) ComputeListDispatch(compute_list int, x_groups int, y_groups int, z_groups int) { //gd:RenderingDevice.compute_list_dispatch
	Advanced(self).ComputeListDispatch(int64(compute_list), int64(x_groups), int64(y_groups), int64(z_groups))
}

/*
Submits the compute list for processing on the GPU with the given group counts stored in the [param buffer] at [param offset]. Buffer must have been created with [constant STORAGE_BUFFER_USAGE_DISPATCH_INDIRECT] flag.
*/
func (self Instance) ComputeListDispatchIndirect(compute_list int, buffer RID.Buffer, offset int) { //gd:RenderingDevice.compute_list_dispatch_indirect
	Advanced(self).ComputeListDispatchIndirect(int64(compute_list), RID.Any(buffer), int64(offset))
}

/*
Raises a Vulkan compute barrier in the specified [param compute_list].
*/
func (self Instance) ComputeListAddBarrier(compute_list int) { //gd:RenderingDevice.compute_list_add_barrier
	Advanced(self).ComputeListAddBarrier(int64(compute_list))
}

/*
Finishes a list of compute commands created with the [code]compute_*[/code] methods.
*/
func (self Instance) ComputeListEnd() { //gd:RenderingDevice.compute_list_end
	Advanced(self).ComputeListEnd()
}

/*
Tries to free an object in the RenderingDevice. To avoid memory leaks, this should be called after using an object as memory management does not occur automatically when using RenderingDevice directly.
*/
func (self Instance) FreeRid(rid RID.Any) { //gd:RenderingDevice.free_rid
	Advanced(self).FreeRid(RID.Any(rid))
}

/*
Creates a timestamp marker with the specified [param name]. This is used for performance reporting with the [method get_captured_timestamp_cpu_time], [method get_captured_timestamp_gpu_time] and [method get_captured_timestamp_name] methods.
*/
func (self Instance) CaptureTimestamp(name string) { //gd:RenderingDevice.capture_timestamp
	Advanced(self).CaptureTimestamp(String.New(name))
}

/*
Returns the total number of timestamps (rendering steps) available for profiling.
*/
func (self Instance) GetCapturedTimestampsCount() int { //gd:RenderingDevice.get_captured_timestamps_count
	return int(int(Advanced(self).GetCapturedTimestampsCount()))
}

/*
Returns the index of the last frame rendered that has rendering timestamps available for querying.
*/
func (self Instance) GetCapturedTimestampsFrame() int { //gd:RenderingDevice.get_captured_timestamps_frame
	return int(int(Advanced(self).GetCapturedTimestampsFrame()))
}

/*
Returns the timestamp in GPU time for the rendering step specified by [param index] (in microseconds since the engine started). See also [method get_captured_timestamp_cpu_time] and [method capture_timestamp].
*/
func (self Instance) GetCapturedTimestampGpuTime(index int) int { //gd:RenderingDevice.get_captured_timestamp_gpu_time
	return int(int(Advanced(self).GetCapturedTimestampGpuTime(int64(index))))
}

/*
Returns the timestamp in CPU time for the rendering step specified by [param index] (in microseconds since the engine started). See also [method get_captured_timestamp_gpu_time] and [method capture_timestamp].
*/
func (self Instance) GetCapturedTimestampCpuTime(index int) int { //gd:RenderingDevice.get_captured_timestamp_cpu_time
	return int(int(Advanced(self).GetCapturedTimestampCpuTime(int64(index))))
}

/*
Returns the timestamp's name for the rendering step specified by [param index]. See also [method capture_timestamp].
*/
func (self Instance) GetCapturedTimestampName(index int) string { //gd:RenderingDevice.get_captured_timestamp_name
	return string(Advanced(self).GetCapturedTimestampName(int64(index)).String())
}

/*
Returns [code]true[/code] if the [param feature] is supported by the GPU.
*/
func (self Instance) HasFeature(feature gdclass.RenderingDeviceFeatures) bool { //gd:RenderingDevice.has_feature
	return bool(Advanced(self).HasFeature(feature))
}

/*
Returns the value of the specified [param limit]. This limit varies depending on the current graphics hardware (and sometimes the driver version). If the given limit is exceeded, rendering errors will occur.
Limits for various graphics hardware can be found in the [url=https://vulkan.gpuinfo.org/]Vulkan Hardware Database[/url].
*/
func (self Instance) LimitGet(limit gdclass.RenderingDeviceLimit) int { //gd:RenderingDevice.limit_get
	return int(int(Advanced(self).LimitGet(limit)))
}

/*
Returns the frame count kept by the graphics API. Higher values result in higher input lag, but with more consistent throughput. For the main [RenderingDevice], frames are cycled (usually 3 with triple-buffered V-Sync enabled). However, local [RenderingDevice]s only have 1 frame.
*/
func (self Instance) GetFrameDelay() int { //gd:RenderingDevice.get_frame_delay
	return int(int(Advanced(self).GetFrameDelay()))
}

/*
Pushes the frame setup and draw command buffers then marks the local device as currently processing (which allows calling [method sync]).
[b]Note:[/b] Only available in local RenderingDevices.
*/
func (self Instance) Submit() { //gd:RenderingDevice.submit
	Advanced(self).Submit()
}

/*
Forces a synchronization between the CPU and GPU, which may be required in certain cases. Only call this when needed, as CPU-GPU synchronization has a performance cost.
[b]Note:[/b] Only available in local RenderingDevices.
[b]Note:[/b] [method sync] can only be called after a [method submit].
*/
func (self Instance) Sync() { //gd:RenderingDevice.sync
	Advanced(self).Sync()
}

/*
This method does nothing.
*/
func (self Instance) Barrier() { //gd:RenderingDevice.barrier
	Advanced(self).Barrier(32767, 32767)
}

/*
This method does nothing.
*/
func (self Expanded) Barrier(from gdclass.RenderingDeviceBarrierMask, to gdclass.RenderingDeviceBarrierMask) { //gd:RenderingDevice.barrier
	Advanced(self).Barrier(from, to)
}

/*
This method does nothing.
*/
func (self Instance) FullBarrier() { //gd:RenderingDevice.full_barrier
	Advanced(self).FullBarrier()
}

/*
Create a new local [RenderingDevice]. This is most useful for performing compute operations on the GPU independently from the rest of the engine.
*/
func (self Instance) CreateLocalDevice() [1]gdclass.RenderingDevice { //gd:RenderingDevice.create_local_device
	return [1]gdclass.RenderingDevice(Advanced(self).CreateLocalDevice())
}

/*
Sets the resource name for [param id] to [param name]. This is used for debugging with third-party tools such as [url=https://renderdoc.org/]RenderDoc[/url].
The following types of resources can be named: texture, sampler, vertex buffer, index buffer, uniform buffer, texture buffer, storage buffer, uniform set buffer, shader, render pipeline and compute pipeline. Framebuffers cannot be named. Attempting to name an incompatible resource type will print an error.
[b]Note:[/b] Resource names are only set when the engine runs in verbose mode ([method OS.is_stdout_verbose] = [code]true[/code]), or when using an engine build compiled with the [code]dev_mode=yes[/code] SCons option. The graphics driver must also support the [code]VK_EXT_DEBUG_UTILS_EXTENSION_NAME[/code] Vulkan extension for named resources to work.
*/
func (self Instance) SetResourceName(id RID.Any, name string) { //gd:RenderingDevice.set_resource_name
	Advanced(self).SetResourceName(RID.Any(id), String.New(name))
}

/*
Create a command buffer debug label region that can be displayed in third-party tools such as [url=https://renderdoc.org/]RenderDoc[/url]. All regions must be ended with a [method draw_command_end_label] call. When viewed from the linear series of submissions to a single queue, calls to [method draw_command_begin_label] and [method draw_command_end_label] must be matched and balanced.
The [code]VK_EXT_DEBUG_UTILS_EXTENSION_NAME[/code] Vulkan extension must be available and enabled for command buffer debug label region to work. See also [method draw_command_end_label].
*/
func (self Instance) DrawCommandBeginLabel(name string, color Color.RGBA) { //gd:RenderingDevice.draw_command_begin_label
	Advanced(self).DrawCommandBeginLabel(String.New(name), Color.RGBA(color))
}

/*
This method does nothing.
*/
func (self Instance) DrawCommandInsertLabel(name string, color Color.RGBA) { //gd:RenderingDevice.draw_command_insert_label
	Advanced(self).DrawCommandInsertLabel(String.New(name), Color.RGBA(color))
}

/*
Ends the command buffer debug label region started by a [method draw_command_begin_label] call.
*/
func (self Instance) DrawCommandEndLabel() { //gd:RenderingDevice.draw_command_end_label
	Advanced(self).DrawCommandEndLabel()
}

/*
Returns the vendor of the video adapter (e.g. "NVIDIA Corporation"). Equivalent to [method RenderingServer.get_video_adapter_vendor]. See also [method get_device_name].
*/
func (self Instance) GetDeviceVendorName() string { //gd:RenderingDevice.get_device_vendor_name
	return string(Advanced(self).GetDeviceVendorName().String())
}

/*
Returns the name of the video adapter (e.g. "GeForce GTX 1080/PCIe/SSE2"). Equivalent to [method RenderingServer.get_video_adapter_name]. See also [method get_device_vendor_name].
*/
func (self Instance) GetDeviceName() string { //gd:RenderingDevice.get_device_name
	return string(Advanced(self).GetDeviceName().String())
}

/*
Returns the universally unique identifier for the pipeline cache. This is used to cache shader files on disk, which avoids shader recompilations on subsequent engine runs. This UUID varies depending on the graphics card model, but also the driver version. Therefore, updating graphics drivers will invalidate the shader cache.
*/
func (self Instance) GetDevicePipelineCacheUuid() string { //gd:RenderingDevice.get_device_pipeline_cache_uuid
	return string(Advanced(self).GetDevicePipelineCacheUuid().String())
}

/*
Returns the memory usage in bytes corresponding to the given [param type]. When using Vulkan, these statistics are calculated by [url=https://github.com/GPUOpen-LibrariesAndSDKs/VulkanMemoryAllocator]Vulkan Memory Allocator[/url].
*/
func (self Instance) GetMemoryUsage(atype gdclass.RenderingDeviceMemoryType) int { //gd:RenderingDevice.get_memory_usage
	return int(int(Advanced(self).GetMemoryUsage(atype)))
}

/*
Returns the unique identifier of the driver [param resource] for the specified [param rid]. Some driver resource types ignore the specified [param rid] (see [enum DriverResource] descriptions). [param index] is always ignored but must be specified anyway.
*/
func (self Instance) GetDriverResource(resource gdclass.RenderingDeviceDriverResource, rid RID.Any, index int) int { //gd:RenderingDevice.get_driver_resource
	return int(int(Advanced(self).GetDriverResource(resource, RID.Any(rid), int64(index))))
}

/*
Returns a string with a performance report from the past frame. Updates every frame.
*/
func (self Instance) GetPerfReport() string { //gd:RenderingDevice.get_perf_report
	return string(Advanced(self).GetPerfReport().String())
}

/*
Returns string report in CSV format using the following methods:
- [method get_tracked_object_name]
- [method get_tracked_object_type_count]
- [method get_driver_total_memory]
- [method get_driver_allocation_count]
- [method get_driver_memory_by_object_type]
- [method get_driver_allocs_by_object_type]
- [method get_device_total_memory]
- [method get_device_allocation_count]
- [method get_device_memory_by_object_type]
- [method get_device_allocs_by_object_type]
This is only used by Vulkan in debug builds. Godot must also be started with the [code]--extra-gpu-memory-tracking[/code] [url=$DOCS_URL/tutorials/editor/command_line_tutorial.html]command line argument[/url].
*/
func (self Instance) GetDriverAndDeviceMemoryReport() string { //gd:RenderingDevice.get_driver_and_device_memory_report
	return string(Advanced(self).GetDriverAndDeviceMemoryReport().String())
}

/*
Returns the name of the type of object for the given [param type_index]. This value must be in range [code][0; get_tracked_object_type_count - 1][/code]. If [method get_tracked_object_type_count] is 0, then type argument is ignored and always returns the same string.
The return value is important because it gives meaning to the types passed to [method get_driver_memory_by_object_type], [method get_driver_allocs_by_object_type], [method get_device_memory_by_object_type], and [method get_device_allocs_by_object_type]. Examples of strings it can return (not exhaustive):
- DEVICE_MEMORY
- PIPELINE_CACHE
- SWAPCHAIN_KHR
- COMMAND_POOL
Thus if e.g. [code]get_tracked_object_name(5)[/code] returns "COMMAND_POOL", then [code]get_device_memory_by_object_type(5)[/code] returns the bytes used by the GPU for command pools.
This is only used by Vulkan in debug builds. Godot must also be started with the [code]--extra-gpu-memory-tracking[/code] [url=$DOCS_URL/tutorials/editor/command_line_tutorial.html]command line argument[/url].
*/
func (self Instance) GetTrackedObjectName(type_index int) string { //gd:RenderingDevice.get_tracked_object_name
	return string(Advanced(self).GetTrackedObjectName(int64(type_index)).String())
}

/*
Returns how many types of trackable objects are.
This is only used by Vulkan in debug builds. Godot must also be started with the [code]--extra-gpu-memory-tracking[/code] [url=$DOCS_URL/tutorials/editor/command_line_tutorial.html]command line argument[/url].
*/
func (self Instance) GetTrackedObjectTypeCount() int { //gd:RenderingDevice.get_tracked_object_type_count
	return int(int(Advanced(self).GetTrackedObjectTypeCount()))
}

/*
Returns how much bytes the GPU driver is using for internal driver structures.
This is only used by Vulkan in debug builds and can return 0 when this information is not tracked or unknown.
*/
func (self Instance) GetDriverTotalMemory() int { //gd:RenderingDevice.get_driver_total_memory
	return int(int(Advanced(self).GetDriverTotalMemory()))
}

/*
Returns how many allocations the GPU driver has performed for internal driver structures.
This is only used by Vulkan in debug builds and can return 0 when this information is not tracked or unknown.
*/
func (self Instance) GetDriverAllocationCount() int { //gd:RenderingDevice.get_driver_allocation_count
	return int(int(Advanced(self).GetDriverAllocationCount()))
}

/*
Same as [method get_driver_total_memory] but filtered for a given object type.
The type argument must be in range [code][0; get_tracked_object_type_count - 1][/code]. If [method get_tracked_object_type_count] is 0, then type argument is ignored and always returns 0.
This is only used by Vulkan in debug builds and can return 0 when this information is not tracked or unknown.
*/
func (self Instance) GetDriverMemoryByObjectType(atype int) int { //gd:RenderingDevice.get_driver_memory_by_object_type
	return int(int(Advanced(self).GetDriverMemoryByObjectType(int64(atype))))
}

/*
Same as [method get_driver_allocation_count] but filtered for a given object type.
The type argument must be in range [code][0; get_tracked_object_type_count - 1][/code]. If [method get_tracked_object_type_count] is 0, then type argument is ignored and always returns 0.
This is only used by Vulkan in debug builds and can return 0 when this information is not tracked or unknown.
*/
func (self Instance) GetDriverAllocsByObjectType(atype int) int { //gd:RenderingDevice.get_driver_allocs_by_object_type
	return int(int(Advanced(self).GetDriverAllocsByObjectType(int64(atype))))
}

/*
Returns how much bytes the GPU is using.
This is only used by Vulkan in debug builds and can return 0 when this information is not tracked or unknown.
*/
func (self Instance) GetDeviceTotalMemory() int { //gd:RenderingDevice.get_device_total_memory
	return int(int(Advanced(self).GetDeviceTotalMemory()))
}

/*
Returns how many allocations the GPU has performed for internal driver structures.
This is only used by Vulkan in debug builds and can return 0 when this information is not tracked or unknown.
*/
func (self Instance) GetDeviceAllocationCount() int { //gd:RenderingDevice.get_device_allocation_count
	return int(int(Advanced(self).GetDeviceAllocationCount()))
}

/*
Same as [method get_device_total_memory] but filtered for a given object type.
The type argument must be in range [code][0; get_tracked_object_type_count - 1][/code]. If [method get_tracked_object_type_count] is 0, then type argument is ignored and always returns 0.
This is only used by Vulkan in debug builds and can return 0 when this information is not tracked or unknown.
*/
func (self Instance) GetDeviceMemoryByObjectType(atype int) int { //gd:RenderingDevice.get_device_memory_by_object_type
	return int(int(Advanced(self).GetDeviceMemoryByObjectType(int64(atype))))
}

/*
Same as [method get_device_allocation_count] but filtered for a given object type.
The type argument must be in range [code][0; get_tracked_object_type_count - 1][/code]. If [method get_tracked_object_type_count] is 0, then type argument is ignored and always returns 0.
This is only used by Vulkan in debug builds and can return 0 when this information is not tracked or unknown.
*/
func (self Instance) GetDeviceAllocsByObjectType(atype int) int { //gd:RenderingDevice.get_device_allocs_by_object_type
	return int(int(Advanced(self).GetDeviceAllocsByObjectType(int64(atype))))
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.RenderingDevice

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("RenderingDevice"))
	casted := Instance{*(*gdclass.RenderingDevice)(unsafe.Pointer(&object))}
	return casted
}

/*
Creates a new texture. It can be accessed with the RID that is returned.
Once finished with your RID, you will want to free the RID using the RenderingDevice's [method free_rid] method.
[b]Note:[/b] Not to be confused with [method RenderingServer.texture_2d_create], which creates the Godot-specific [Texture2D] resource as opposed to the graphics API's own texture type.
*/
//go:nosplit
func (self class) TextureCreate(format [1]gdclass.RDTextureFormat, view [1]gdclass.RDTextureView, data Array.Contains[Packed.Bytes]) RID.Any { //gd:RenderingDevice.texture_create
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(format[0])[0])
	callframe.Arg(frame, pointers.Get(view[0])[0])
	callframe.Arg(frame, pointers.Get(gd.InternalArray(data)))
	var r_ret = callframe.Ret[RID.Any](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RenderingDevice.Bind_texture_create, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Creates a shared texture using the specified [param view] and the texture information from [param with_texture].
*/
//go:nosplit
func (self class) TextureCreateShared(view [1]gdclass.RDTextureView, with_texture RID.Any) RID.Any { //gd:RenderingDevice.texture_create_shared
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(view[0])[0])
	callframe.Arg(frame, with_texture)
	var r_ret = callframe.Ret[RID.Any](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RenderingDevice.Bind_texture_create_shared, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Creates a shared texture using the specified [param view] and the texture information from [param with_texture]'s [param layer] and [param mipmap]. The number of included mipmaps from the original texture can be controlled using the [param mipmaps] parameter. Only relevant for textures with multiple layers, such as 3D textures, texture arrays and cubemaps. For single-layer textures, use [method texture_create_shared].
For 2D textures (which only have one layer), [param layer] must be [code]0[/code].
[b]Note:[/b] Layer slicing is only supported for 2D texture arrays, not 3D textures or cubemaps.
*/
//go:nosplit
func (self class) TextureCreateSharedFromSlice(view [1]gdclass.RDTextureView, with_texture RID.Any, layer int64, mipmap int64, mipmaps int64, slice_type gdclass.RenderingDeviceTextureSliceType) RID.Any { //gd:RenderingDevice.texture_create_shared_from_slice
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(view[0])[0])
	callframe.Arg(frame, with_texture)
	callframe.Arg(frame, layer)
	callframe.Arg(frame, mipmap)
	callframe.Arg(frame, mipmaps)
	callframe.Arg(frame, slice_type)
	var r_ret = callframe.Ret[RID.Any](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RenderingDevice.Bind_texture_create_shared_from_slice, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns an RID for an existing [param image] ([code]VkImage[/code]) with the given [param type], [param format], [param samples], [param usage_flags], [param width], [param height], [param depth], and [param layers]. This can be used to allow Godot to render onto foreign images.
*/
//go:nosplit
func (self class) TextureCreateFromExtension(atype gdclass.RenderingDeviceTextureType, format gdclass.RenderingDeviceDataFormat, samples gdclass.RenderingDeviceTextureSamples, usage_flags gdclass.RenderingDeviceTextureUsageBits, image int64, width int64, height int64, depth int64, layers int64) RID.Any { //gd:RenderingDevice.texture_create_from_extension
	var frame = callframe.New()
	callframe.Arg(frame, atype)
	callframe.Arg(frame, format)
	callframe.Arg(frame, samples)
	callframe.Arg(frame, usage_flags)
	callframe.Arg(frame, image)
	callframe.Arg(frame, width)
	callframe.Arg(frame, height)
	callframe.Arg(frame, depth)
	callframe.Arg(frame, layers)
	var r_ret = callframe.Ret[RID.Any](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RenderingDevice.Bind_texture_create_from_extension, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Updates texture data with new data, replacing the previous data in place. The updated texture data must have the same dimensions and format. For 2D textures (which only have one layer), [param layer] must be [code]0[/code]. Returns [constant @GlobalScope.OK] if the update was successful, [constant @GlobalScope.ERR_INVALID_PARAMETER] otherwise.
[b]Note:[/b] Updating textures is forbidden during creation of a draw or compute list.
[b]Note:[/b] The existing [param texture] can't be updated while a draw list that uses it as part of a framebuffer is being created. Ensure the draw list is finalized (and that the color/depth texture using it is not set to [constant FINAL_ACTION_CONTINUE]) to update this texture.
[b]Note:[/b] The existing [param texture] requires the [constant TEXTURE_USAGE_CAN_UPDATE_BIT] to be updatable.
*/
//go:nosplit
func (self class) TextureUpdate(texture RID.Any, layer int64, data Packed.Bytes) Error.Code { //gd:RenderingDevice.texture_update
	var frame = callframe.New()
	callframe.Arg(frame, texture)
	callframe.Arg(frame, layer)
	callframe.Arg(frame, pointers.Get(gd.InternalPacked[gd.PackedByteArray, byte](Packed.Array[byte](data))))
	var r_ret = callframe.Ret[int64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RenderingDevice.Bind_texture_update, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = Error.Code(r_ret.Get())
	frame.Free()
	return ret
}

/*
Returns the [param texture] data for the specified [param layer] as raw binary data. For 2D textures (which only have one layer), [param layer] must be [code]0[/code].
[b]Note:[/b] [param texture] can't be retrieved while a draw list that uses it as part of a framebuffer is being created. Ensure the draw list is finalized (and that the color/depth texture using it is not set to [constant FINAL_ACTION_CONTINUE]) to retrieve this texture. Otherwise, an error is printed and a empty [PackedByteArray] is returned.
[b]Note:[/b] [param texture] requires the [constant TEXTURE_USAGE_CAN_COPY_FROM_BIT] to be retrieved. Otherwise, an error is printed and a empty [PackedByteArray] is returned.
[b]Note:[/b] This method will block the GPU from working until the data is retrieved. Refer to [method texture_get_data_async] for an alternative that returns the data in more performant way.
*/
//go:nosplit
func (self class) TextureGetData(texture RID.Any, layer int64) Packed.Bytes { //gd:RenderingDevice.texture_get_data
	var frame = callframe.New()
	callframe.Arg(frame, texture)
	callframe.Arg(frame, layer)
	var r_ret = callframe.Ret[gd.PackedPointers](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RenderingDevice.Bind_texture_get_data, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = Packed.Bytes(Array.Through(gd.PackedProxy[gd.PackedByteArray, byte]{}, pointers.Pack(pointers.New[gd.PackedByteArray](r_ret.Get()))))
	frame.Free()
	return ret
}

/*
Asynchronous version of [method texture_get_data]. RenderingDevice will call [param callback] in a certain amount of frames with the data the texture had at the time of the request.
[b]Note:[/b] At the moment, the delay corresponds to the amount of frames specified by [member ProjectSettings.rendering/rendering_device/vsync/frame_queue_size].
[b]Note:[/b] Downloading large textures can have a prohibitive cost for real-time even when using the asynchronous method due to hardware bandwidth limitations. When dealing with large resources, you can adjust settings such as [member ProjectSettings.rendering/rendering_device/staging_buffer/texture_download_region_size_px] and [member ProjectSettings.rendering/rendering_device/staging_buffer/block_size_kb] to improve the transfer speed at the cost of extra memory.
[codeblock]
func _texture_get_data_callback(array):
    value = array.decode_u32(0)

...

rd.texture_get_data_async(texture, 0, _texture_get_data_callback)
[/codeblock]
*/
//go:nosplit
func (self class) TextureGetDataAsync(texture RID.Any, layer int64, callback Callable.Function) Error.Code { //gd:RenderingDevice.texture_get_data_async
	var frame = callframe.New()
	callframe.Arg(frame, texture)
	callframe.Arg(frame, layer)
	callframe.Arg(frame, pointers.Get(gd.InternalCallable(callback)))
	var r_ret = callframe.Ret[int64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RenderingDevice.Bind_texture_get_data_async, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = Error.Code(r_ret.Get())
	frame.Free()
	return ret
}

/*
Returns [code]true[/code] if the specified [param format] is supported for the given [param usage_flags], [code]false[/code] otherwise.
*/
//go:nosplit
func (self class) TextureIsFormatSupportedForUsage(format gdclass.RenderingDeviceDataFormat, usage_flags gdclass.RenderingDeviceTextureUsageBits) bool { //gd:RenderingDevice.texture_is_format_supported_for_usage
	var frame = callframe.New()
	callframe.Arg(frame, format)
	callframe.Arg(frame, usage_flags)
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RenderingDevice.Bind_texture_is_format_supported_for_usage, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns [code]true[/code] if the [param texture] is shared, [code]false[/code] otherwise. See [RDTextureView].
*/
//go:nosplit
func (self class) TextureIsShared(texture RID.Any) bool { //gd:RenderingDevice.texture_is_shared
	var frame = callframe.New()
	callframe.Arg(frame, texture)
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RenderingDevice.Bind_texture_is_shared, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns [code]true[/code] if the [param texture] is valid, [code]false[/code] otherwise.
*/
//go:nosplit
func (self class) TextureIsValid(texture RID.Any) bool { //gd:RenderingDevice.texture_is_valid
	var frame = callframe.New()
	callframe.Arg(frame, texture)
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RenderingDevice.Bind_texture_is_valid, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Updates the discardable property of [param texture].
If a texture is discardable, its contents do not need to be preserved between frames. This flag is only relevant when the texture is used as target in a draw list.
This information is used by [RenderingDevice] to figure out if a texture's contents can be discarded, eliminating unnecessary writes to memory and boosting performance.
*/
//go:nosplit
func (self class) TextureSetDiscardable(texture RID.Any, discardable bool) { //gd:RenderingDevice.texture_set_discardable
	var frame = callframe.New()
	callframe.Arg(frame, texture)
	callframe.Arg(frame, discardable)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RenderingDevice.Bind_texture_set_discardable, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns [code]true[/code] if the [param texture] is discardable, [code]false[/code] otherwise. See [RDTextureFormat] or [method texture_set_discardable].
*/
//go:nosplit
func (self class) TextureIsDiscardable(texture RID.Any) bool { //gd:RenderingDevice.texture_is_discardable
	var frame = callframe.New()
	callframe.Arg(frame, texture)
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RenderingDevice.Bind_texture_is_discardable, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Copies the [param from_texture] to [param to_texture] with the specified [param from_pos], [param to_pos] and [param size] coordinates. The Z axis of the [param from_pos], [param to_pos] and [param size] must be [code]0[/code] for 2-dimensional textures. Source and destination mipmaps/layers must also be specified, with these parameters being [code]0[/code] for textures without mipmaps or single-layer textures. Returns [constant @GlobalScope.OK] if the texture copy was successful or [constant @GlobalScope.ERR_INVALID_PARAMETER] otherwise.
[b]Note:[/b] [param from_texture] texture can't be copied while a draw list that uses it as part of a framebuffer is being created. Ensure the draw list is finalized (and that the color/depth texture using it is not set to [constant FINAL_ACTION_CONTINUE]) to copy this texture.
[b]Note:[/b] [param from_texture] texture requires the [constant TEXTURE_USAGE_CAN_COPY_FROM_BIT] to be retrieved.
[b]Note:[/b] [param to_texture] can't be copied while a draw list that uses it as part of a framebuffer is being created. Ensure the draw list is finalized (and that the color/depth texture using it is not set to [constant FINAL_ACTION_CONTINUE]) to copy this texture.
[b]Note:[/b] [param to_texture] requires the [constant TEXTURE_USAGE_CAN_COPY_TO_BIT] to be retrieved.
[b]Note:[/b] [param from_texture] and [param to_texture] must be of the same type (color or depth).
*/
//go:nosplit
func (self class) TextureCopy(from_texture RID.Any, to_texture RID.Any, from_pos Vector3.XYZ, to_pos Vector3.XYZ, size Vector3.XYZ, src_mipmap int64, dst_mipmap int64, src_layer int64, dst_layer int64) Error.Code { //gd:RenderingDevice.texture_copy
	var frame = callframe.New()
	callframe.Arg(frame, from_texture)
	callframe.Arg(frame, to_texture)
	callframe.Arg(frame, from_pos)
	callframe.Arg(frame, to_pos)
	callframe.Arg(frame, size)
	callframe.Arg(frame, src_mipmap)
	callframe.Arg(frame, dst_mipmap)
	callframe.Arg(frame, src_layer)
	callframe.Arg(frame, dst_layer)
	var r_ret = callframe.Ret[int64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RenderingDevice.Bind_texture_copy, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = Error.Code(r_ret.Get())
	frame.Free()
	return ret
}

/*
Clears the specified [param texture] by replacing all of its pixels with the specified [param color]. [param base_mipmap] and [param mipmap_count] determine which mipmaps of the texture are affected by this clear operation, while [param base_layer] and [param layer_count] determine which layers of a 3D texture (or texture array) are affected by this clear operation. For 2D textures (which only have one layer by design), [param base_layer] must be [code]0[/code] and [param layer_count] must be [code]1[/code].
[b]Note:[/b] [param texture] can't be cleared while a draw list that uses it as part of a framebuffer is being created. Ensure the draw list is finalized (and that the color/depth texture using it is not set to [constant FINAL_ACTION_CONTINUE]) to clear this texture.
*/
//go:nosplit
func (self class) TextureClear(texture RID.Any, color Color.RGBA, base_mipmap int64, mipmap_count int64, base_layer int64, layer_count int64) Error.Code { //gd:RenderingDevice.texture_clear
	var frame = callframe.New()
	callframe.Arg(frame, texture)
	callframe.Arg(frame, color)
	callframe.Arg(frame, base_mipmap)
	callframe.Arg(frame, mipmap_count)
	callframe.Arg(frame, base_layer)
	callframe.Arg(frame, layer_count)
	var r_ret = callframe.Ret[int64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RenderingDevice.Bind_texture_clear, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = Error.Code(r_ret.Get())
	frame.Free()
	return ret
}

/*
Resolves the [param from_texture] texture onto [param to_texture] with multisample antialiasing enabled. This must be used when rendering a framebuffer for MSAA to work. Returns [constant @GlobalScope.OK] if successful, [constant @GlobalScope.ERR_INVALID_PARAMETER] otherwise.
[b]Note:[/b] [param from_texture] and [param to_texture] textures must have the same dimension, format and type (color or depth).
[b]Note:[/b] [param from_texture] can't be copied while a draw list that uses it as part of a framebuffer is being created. Ensure the draw list is finalized (and that the color/depth texture using it is not set to [constant FINAL_ACTION_CONTINUE]) to resolve this texture.
[b]Note:[/b] [param from_texture] requires the [constant TEXTURE_USAGE_CAN_COPY_FROM_BIT] to be retrieved.
[b]Note:[/b] [param from_texture] must be multisampled and must also be 2D (or a slice of a 3D/cubemap texture).
[b]Note:[/b] [param to_texture] can't be copied while a draw list that uses it as part of a framebuffer is being created. Ensure the draw list is finalized (and that the color/depth texture using it is not set to [constant FINAL_ACTION_CONTINUE]) to resolve this texture.
[b]Note:[/b] [param to_texture] texture requires the [constant TEXTURE_USAGE_CAN_COPY_TO_BIT] to be retrieved.
[b]Note:[/b] [param to_texture] texture must [b]not[/b] be multisampled and must also be 2D (or a slice of a 3D/cubemap texture).
*/
//go:nosplit
func (self class) TextureResolveMultisample(from_texture RID.Any, to_texture RID.Any) Error.Code { //gd:RenderingDevice.texture_resolve_multisample
	var frame = callframe.New()
	callframe.Arg(frame, from_texture)
	callframe.Arg(frame, to_texture)
	var r_ret = callframe.Ret[int64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RenderingDevice.Bind_texture_resolve_multisample, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = Error.Code(r_ret.Get())
	frame.Free()
	return ret
}

/*
Returns the data format used to create this texture.
*/
//go:nosplit
func (self class) TextureGetFormat(texture RID.Any) [1]gdclass.RDTextureFormat { //gd:RenderingDevice.texture_get_format
	var frame = callframe.New()
	callframe.Arg(frame, texture)
	var r_ret = callframe.Ret[gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RenderingDevice.Bind_texture_get_format, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = [1]gdclass.RDTextureFormat{gd.PointerWithOwnershipTransferredToGo[gdclass.RDTextureFormat](r_ret.Get())}
	frame.Free()
	return ret
}

/*
Returns the internal graphics handle for this texture object. For use when communicating with third-party APIs mostly with GDExtension.
[b]Note:[/b] This function returns a [code]uint64_t[/code] which internally maps to a [code]GLuint[/code] (OpenGL) or [code]VkImage[/code] (Vulkan).
*/
//go:nosplit
func (self class) TextureGetNativeHandle(texture RID.Any) int64 { //gd:RenderingDevice.texture_get_native_handle
	var frame = callframe.New()
	callframe.Arg(frame, texture)
	var r_ret = callframe.Ret[int64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RenderingDevice.Bind_texture_get_native_handle, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Creates a new framebuffer format with the specified [param attachments] and [param view_count]. Returns the new framebuffer's unique framebuffer format ID.
If [param view_count] is greater than or equal to [code]2[/code], enables multiview which is used for VR rendering. This requires support for the Vulkan multiview extension.
*/
//go:nosplit
func (self class) FramebufferFormatCreate(attachments Array.Contains[[1]gdclass.RDAttachmentFormat], view_count int64) int64 { //gd:RenderingDevice.framebuffer_format_create
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalArray(attachments)))
	callframe.Arg(frame, view_count)
	var r_ret = callframe.Ret[int64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RenderingDevice.Bind_framebuffer_format_create, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Creates a multipass framebuffer format with the specified [param attachments], [param passes] and [param view_count] and returns its ID. If [param view_count] is greater than or equal to [code]2[/code], enables multiview which is used for VR rendering. This requires support for the Vulkan multiview extension.
*/
//go:nosplit
func (self class) FramebufferFormatCreateMultipass(attachments Array.Contains[[1]gdclass.RDAttachmentFormat], passes Array.Contains[[1]gdclass.RDFramebufferPass], view_count int64) int64 { //gd:RenderingDevice.framebuffer_format_create_multipass
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalArray(attachments)))
	callframe.Arg(frame, pointers.Get(gd.InternalArray(passes)))
	callframe.Arg(frame, view_count)
	var r_ret = callframe.Ret[int64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RenderingDevice.Bind_framebuffer_format_create_multipass, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Creates a new empty framebuffer format with the specified number of [param samples] and returns its ID.
*/
//go:nosplit
func (self class) FramebufferFormatCreateEmpty(samples gdclass.RenderingDeviceTextureSamples) int64 { //gd:RenderingDevice.framebuffer_format_create_empty
	var frame = callframe.New()
	callframe.Arg(frame, samples)
	var r_ret = callframe.Ret[int64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RenderingDevice.Bind_framebuffer_format_create_empty, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the number of texture samples used for the given framebuffer [param format] ID (returned by [method framebuffer_get_format]).
*/
//go:nosplit
func (self class) FramebufferFormatGetTextureSamples(format int64, render_pass int64) gdclass.RenderingDeviceTextureSamples { //gd:RenderingDevice.framebuffer_format_get_texture_samples
	var frame = callframe.New()
	callframe.Arg(frame, format)
	callframe.Arg(frame, render_pass)
	var r_ret = callframe.Ret[gdclass.RenderingDeviceTextureSamples](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RenderingDevice.Bind_framebuffer_format_get_texture_samples, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Creates a new framebuffer. It can be accessed with the RID that is returned.
Once finished with your RID, you will want to free the RID using the RenderingDevice's [method free_rid] method.
*/
//go:nosplit
func (self class) FramebufferCreate(textures Array.Contains[RID.Any], validate_with_format int64, view_count int64) RID.Any { //gd:RenderingDevice.framebuffer_create
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalArray(textures)))
	callframe.Arg(frame, validate_with_format)
	callframe.Arg(frame, view_count)
	var r_ret = callframe.Ret[RID.Any](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RenderingDevice.Bind_framebuffer_create, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Creates a new multipass framebuffer. It can be accessed with the RID that is returned.
Once finished with your RID, you will want to free the RID using the RenderingDevice's [method free_rid] method.
*/
//go:nosplit
func (self class) FramebufferCreateMultipass(textures Array.Contains[RID.Any], passes Array.Contains[[1]gdclass.RDFramebufferPass], validate_with_format int64, view_count int64) RID.Any { //gd:RenderingDevice.framebuffer_create_multipass
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalArray(textures)))
	callframe.Arg(frame, pointers.Get(gd.InternalArray(passes)))
	callframe.Arg(frame, validate_with_format)
	callframe.Arg(frame, view_count)
	var r_ret = callframe.Ret[RID.Any](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RenderingDevice.Bind_framebuffer_create_multipass, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Creates a new empty framebuffer. It can be accessed with the RID that is returned.
Once finished with your RID, you will want to free the RID using the RenderingDevice's [method free_rid] method.
*/
//go:nosplit
func (self class) FramebufferCreateEmpty(size Vector2i.XY, samples gdclass.RenderingDeviceTextureSamples, validate_with_format int64) RID.Any { //gd:RenderingDevice.framebuffer_create_empty
	var frame = callframe.New()
	callframe.Arg(frame, size)
	callframe.Arg(frame, samples)
	callframe.Arg(frame, validate_with_format)
	var r_ret = callframe.Ret[RID.Any](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RenderingDevice.Bind_framebuffer_create_empty, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the format ID of the framebuffer specified by the [param framebuffer] RID. This ID is guaranteed to be unique for the same formats and does not need to be freed.
*/
//go:nosplit
func (self class) FramebufferGetFormat(framebuffer RID.Any) int64 { //gd:RenderingDevice.framebuffer_get_format
	var frame = callframe.New()
	callframe.Arg(frame, framebuffer)
	var r_ret = callframe.Ret[int64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RenderingDevice.Bind_framebuffer_get_format, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns [code]true[/code] if the framebuffer specified by the [param framebuffer] RID is valid, [code]false[/code] otherwise.
*/
//go:nosplit
func (self class) FramebufferIsValid(framebuffer RID.Any) bool { //gd:RenderingDevice.framebuffer_is_valid
	var frame = callframe.New()
	callframe.Arg(frame, framebuffer)
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RenderingDevice.Bind_framebuffer_is_valid, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Creates a new sampler. It can be accessed with the RID that is returned.
Once finished with your RID, you will want to free the RID using the RenderingDevice's [method free_rid] method.
*/
//go:nosplit
func (self class) SamplerCreate(state [1]gdclass.RDSamplerState) RID.Any { //gd:RenderingDevice.sampler_create
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(state[0])[0])
	var r_ret = callframe.Ret[RID.Any](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RenderingDevice.Bind_sampler_create, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns [code]true[/code] if implementation supports using a texture of [param format] with the given [param sampler_filter].
*/
//go:nosplit
func (self class) SamplerIsFormatSupportedForFilter(format gdclass.RenderingDeviceDataFormat, sampler_filter gdclass.RenderingDeviceSamplerFilter) bool { //gd:RenderingDevice.sampler_is_format_supported_for_filter
	var frame = callframe.New()
	callframe.Arg(frame, format)
	callframe.Arg(frame, sampler_filter)
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RenderingDevice.Bind_sampler_is_format_supported_for_filter, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
It can be accessed with the RID that is returned.
Once finished with your RID, you will want to free the RID using the RenderingDevice's [method free_rid] method.
*/
//go:nosplit
func (self class) VertexBufferCreate(size_bytes int64, data Packed.Bytes, creation_bits gdclass.RenderingDeviceBufferCreationBits) RID.Any { //gd:RenderingDevice.vertex_buffer_create
	var frame = callframe.New()
	callframe.Arg(frame, size_bytes)
	callframe.Arg(frame, pointers.Get(gd.InternalPacked[gd.PackedByteArray, byte](Packed.Array[byte](data))))
	callframe.Arg(frame, creation_bits)
	var r_ret = callframe.Ret[RID.Any](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RenderingDevice.Bind_vertex_buffer_create, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Creates a new vertex format with the specified [param vertex_descriptions]. Returns a unique vertex format ID corresponding to the newly created vertex format.
*/
//go:nosplit
func (self class) VertexFormatCreate(vertex_descriptions Array.Contains[[1]gdclass.RDVertexAttribute]) int64 { //gd:RenderingDevice.vertex_format_create
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalArray(vertex_descriptions)))
	var r_ret = callframe.Ret[int64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RenderingDevice.Bind_vertex_format_create, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Creates a vertex array based on the specified buffers. Optionally, [param offsets] (in bytes) may be defined for each buffer.
*/
//go:nosplit
func (self class) VertexArrayCreate(vertex_count int64, vertex_format int64, src_buffers Array.Contains[RID.Any], offsets Packed.Array[int64]) RID.Any { //gd:RenderingDevice.vertex_array_create
	var frame = callframe.New()
	callframe.Arg(frame, vertex_count)
	callframe.Arg(frame, vertex_format)
	callframe.Arg(frame, pointers.Get(gd.InternalArray(src_buffers)))
	callframe.Arg(frame, pointers.Get(gd.InternalPacked[gd.PackedInt64Array, int64](offsets)))
	var r_ret = callframe.Ret[RID.Any](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RenderingDevice.Bind_vertex_array_create, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Creates a new index buffer. It can be accessed with the RID that is returned.
Once finished with your RID, you will want to free the RID using the RenderingDevice's [method free_rid] method.
*/
//go:nosplit
func (self class) IndexBufferCreate(size_indices int64, format gdclass.RenderingDeviceIndexBufferFormat, data Packed.Bytes, use_restart_indices bool, creation_bits gdclass.RenderingDeviceBufferCreationBits) RID.Any { //gd:RenderingDevice.index_buffer_create
	var frame = callframe.New()
	callframe.Arg(frame, size_indices)
	callframe.Arg(frame, format)
	callframe.Arg(frame, pointers.Get(gd.InternalPacked[gd.PackedByteArray, byte](Packed.Array[byte](data))))
	callframe.Arg(frame, use_restart_indices)
	callframe.Arg(frame, creation_bits)
	var r_ret = callframe.Ret[RID.Any](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RenderingDevice.Bind_index_buffer_create, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Creates a new index array. It can be accessed with the RID that is returned.
Once finished with your RID, you will want to free the RID using the RenderingDevice's [method free_rid] method.
*/
//go:nosplit
func (self class) IndexArrayCreate(index_buffer RID.Any, index_offset int64, index_count int64) RID.Any { //gd:RenderingDevice.index_array_create
	var frame = callframe.New()
	callframe.Arg(frame, index_buffer)
	callframe.Arg(frame, index_offset)
	callframe.Arg(frame, index_count)
	var r_ret = callframe.Ret[RID.Any](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RenderingDevice.Bind_index_array_create, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Compiles a SPIR-V from the shader source code in [param shader_source] and returns the SPIR-V as a [RDShaderSPIRV]. This intermediate language shader is portable across different GPU models and driver versions, but cannot be run directly by GPUs until compiled into a binary shader using [method shader_compile_binary_from_spirv].
If [param allow_cache] is [code]true[/code], make use of the shader cache generated by Godot. This avoids a potentially lengthy shader compilation step if the shader is already in cache. If [param allow_cache] is [code]false[/code], Godot's shader cache is ignored and the shader will always be recompiled.
*/
//go:nosplit
func (self class) ShaderCompileSpirvFromSource(shader_source [1]gdclass.RDShaderSource, allow_cache bool) [1]gdclass.RDShaderSPIRV { //gd:RenderingDevice.shader_compile_spirv_from_source
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(shader_source[0])[0])
	callframe.Arg(frame, allow_cache)
	var r_ret = callframe.Ret[gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RenderingDevice.Bind_shader_compile_spirv_from_source, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = [1]gdclass.RDShaderSPIRV{gd.PointerWithOwnershipTransferredToGo[gdclass.RDShaderSPIRV](r_ret.Get())}
	frame.Free()
	return ret
}

/*
Compiles a binary shader from [param spirv_data] and returns the compiled binary data as a [PackedByteArray]. This compiled shader is specific to the GPU model and driver version used; it will not work on different GPU models or even different driver versions. See also [method shader_compile_spirv_from_source].
[param name] is an optional human-readable name that can be given to the compiled shader for organizational purposes.
*/
//go:nosplit
func (self class) ShaderCompileBinaryFromSpirv(spirv_data [1]gdclass.RDShaderSPIRV, name String.Readable) Packed.Bytes { //gd:RenderingDevice.shader_compile_binary_from_spirv
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(spirv_data[0])[0])
	callframe.Arg(frame, pointers.Get(gd.InternalString(name)))
	var r_ret = callframe.Ret[gd.PackedPointers](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RenderingDevice.Bind_shader_compile_binary_from_spirv, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = Packed.Bytes(Array.Through(gd.PackedProxy[gd.PackedByteArray, byte]{}, pointers.Pack(pointers.New[gd.PackedByteArray](r_ret.Get()))))
	frame.Free()
	return ret
}

/*
Creates a new shader instance from SPIR-V intermediate code. It can be accessed with the RID that is returned.
Once finished with your RID, you will want to free the RID using the RenderingDevice's [method free_rid] method. See also [method shader_compile_spirv_from_source] and [method shader_create_from_bytecode].
*/
//go:nosplit
func (self class) ShaderCreateFromSpirv(spirv_data [1]gdclass.RDShaderSPIRV, name String.Readable) RID.Any { //gd:RenderingDevice.shader_create_from_spirv
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(spirv_data[0])[0])
	callframe.Arg(frame, pointers.Get(gd.InternalString(name)))
	var r_ret = callframe.Ret[RID.Any](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RenderingDevice.Bind_shader_create_from_spirv, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Creates a new shader instance from a binary compiled shader. It can be accessed with the RID that is returned.
Once finished with your RID, you will want to free the RID using the RenderingDevice's [method free_rid] method. See also [method shader_compile_binary_from_spirv] and [method shader_create_from_spirv].
*/
//go:nosplit
func (self class) ShaderCreateFromBytecode(binary_data Packed.Bytes, placeholder_rid RID.Any) RID.Any { //gd:RenderingDevice.shader_create_from_bytecode
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalPacked[gd.PackedByteArray, byte](Packed.Array[byte](binary_data))))
	callframe.Arg(frame, placeholder_rid)
	var r_ret = callframe.Ret[RID.Any](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RenderingDevice.Bind_shader_create_from_bytecode, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Create a placeholder RID by allocating an RID without initializing it for use in [method shader_create_from_bytecode]. This allows you to create an RID for a shader and pass it around, but defer compiling the shader to a later time.
*/
//go:nosplit
func (self class) ShaderCreatePlaceholder() RID.Any { //gd:RenderingDevice.shader_create_placeholder
	var frame = callframe.New()
	var r_ret = callframe.Ret[RID.Any](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RenderingDevice.Bind_shader_create_placeholder, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the internal vertex input mask. Internally, the vertex input mask is an unsigned integer consisting of the locations (specified in GLSL via. [code]layout(location = ...)[/code]) of the input variables (specified in GLSL by the [code]in[/code] keyword).
*/
//go:nosplit
func (self class) ShaderGetVertexInputAttributeMask(shader RID.Any) int64 { //gd:RenderingDevice.shader_get_vertex_input_attribute_mask
	var frame = callframe.New()
	callframe.Arg(frame, shader)
	var r_ret = callframe.Ret[int64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RenderingDevice.Bind_shader_get_vertex_input_attribute_mask, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Creates a new uniform buffer. It can be accessed with the RID that is returned.
Once finished with your RID, you will want to free the RID using the RenderingDevice's [method free_rid] method.
*/
//go:nosplit
func (self class) UniformBufferCreate(size_bytes int64, data Packed.Bytes, creation_bits gdclass.RenderingDeviceBufferCreationBits) RID.Any { //gd:RenderingDevice.uniform_buffer_create
	var frame = callframe.New()
	callframe.Arg(frame, size_bytes)
	callframe.Arg(frame, pointers.Get(gd.InternalPacked[gd.PackedByteArray, byte](Packed.Array[byte](data))))
	callframe.Arg(frame, creation_bits)
	var r_ret = callframe.Ret[RID.Any](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RenderingDevice.Bind_uniform_buffer_create, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Creates a [url=https://vkguide.dev/docs/chapter-4/storage_buffers/]storage buffer[/url] with the specified [param data] and [param usage]. It can be accessed with the RID that is returned.
Once finished with your RID, you will want to free the RID using the RenderingDevice's [method free_rid] method.
*/
//go:nosplit
func (self class) StorageBufferCreate(size_bytes int64, data Packed.Bytes, usage gdclass.RenderingDeviceStorageBufferUsage, creation_bits gdclass.RenderingDeviceBufferCreationBits) RID.Any { //gd:RenderingDevice.storage_buffer_create
	var frame = callframe.New()
	callframe.Arg(frame, size_bytes)
	callframe.Arg(frame, pointers.Get(gd.InternalPacked[gd.PackedByteArray, byte](Packed.Array[byte](data))))
	callframe.Arg(frame, usage)
	callframe.Arg(frame, creation_bits)
	var r_ret = callframe.Ret[RID.Any](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RenderingDevice.Bind_storage_buffer_create, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Creates a new texture buffer. It can be accessed with the RID that is returned.
Once finished with your RID, you will want to free the RID using the RenderingDevice's [method free_rid] method.
*/
//go:nosplit
func (self class) TextureBufferCreate(size_bytes int64, format gdclass.RenderingDeviceDataFormat, data Packed.Bytes) RID.Any { //gd:RenderingDevice.texture_buffer_create
	var frame = callframe.New()
	callframe.Arg(frame, size_bytes)
	callframe.Arg(frame, format)
	callframe.Arg(frame, pointers.Get(gd.InternalPacked[gd.PackedByteArray, byte](Packed.Array[byte](data))))
	var r_ret = callframe.Ret[RID.Any](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RenderingDevice.Bind_texture_buffer_create, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Creates a new uniform set. It can be accessed with the RID that is returned.
Once finished with your RID, you will want to free the RID using the RenderingDevice's [method free_rid] method.
*/
//go:nosplit
func (self class) UniformSetCreate(uniforms Array.Contains[[1]gdclass.RDUniform], shader RID.Any, shader_set int64) RID.Any { //gd:RenderingDevice.uniform_set_create
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalArray(uniforms)))
	callframe.Arg(frame, shader)
	callframe.Arg(frame, shader_set)
	var r_ret = callframe.Ret[RID.Any](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RenderingDevice.Bind_uniform_set_create, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Checks if the [param uniform_set] is valid, i.e. is owned.
*/
//go:nosplit
func (self class) UniformSetIsValid(uniform_set RID.Any) bool { //gd:RenderingDevice.uniform_set_is_valid
	var frame = callframe.New()
	callframe.Arg(frame, uniform_set)
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RenderingDevice.Bind_uniform_set_is_valid, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Copies [param size] bytes from the [param src_buffer] at [param src_offset] into [param dst_buffer] at [param dst_offset].
Prints an error if:
- [param size] exceeds the size of either [param src_buffer] or [param dst_buffer] at their corresponding offsets
- a draw list is currently active (created by [method draw_list_begin])
- a compute list is currently active (created by [method compute_list_begin])
*/
//go:nosplit
func (self class) BufferCopy(src_buffer RID.Any, dst_buffer RID.Any, src_offset int64, dst_offset int64, size int64) Error.Code { //gd:RenderingDevice.buffer_copy
	var frame = callframe.New()
	callframe.Arg(frame, src_buffer)
	callframe.Arg(frame, dst_buffer)
	callframe.Arg(frame, src_offset)
	callframe.Arg(frame, dst_offset)
	callframe.Arg(frame, size)
	var r_ret = callframe.Ret[int64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RenderingDevice.Bind_buffer_copy, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = Error.Code(r_ret.Get())
	frame.Free()
	return ret
}

/*
Updates a region of [param size_bytes] bytes, starting at [param offset], in the buffer, with the specified [param data].
Prints an error if:
- the region specified by [param offset] + [param size_bytes] exceeds the buffer
- a draw list is currently active (created by [method draw_list_begin])
- a compute list is currently active (created by [method compute_list_begin])
*/
//go:nosplit
func (self class) BufferUpdate(buffer RID.Any, offset int64, size_bytes int64, data Packed.Bytes) Error.Code { //gd:RenderingDevice.buffer_update
	var frame = callframe.New()
	callframe.Arg(frame, buffer)
	callframe.Arg(frame, offset)
	callframe.Arg(frame, size_bytes)
	callframe.Arg(frame, pointers.Get(gd.InternalPacked[gd.PackedByteArray, byte](Packed.Array[byte](data))))
	var r_ret = callframe.Ret[int64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RenderingDevice.Bind_buffer_update, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = Error.Code(r_ret.Get())
	frame.Free()
	return ret
}

/*
Clears the contents of the [param buffer], clearing [param size_bytes] bytes, starting at [param offset].
Prints an error if:
- the size isn't a multiple of four
- the region specified by [param offset] + [param size_bytes] exceeds the buffer
- a draw list is currently active (created by [method draw_list_begin])
- a compute list is currently active (created by [method compute_list_begin])
*/
//go:nosplit
func (self class) BufferClear(buffer RID.Any, offset int64, size_bytes int64) Error.Code { //gd:RenderingDevice.buffer_clear
	var frame = callframe.New()
	callframe.Arg(frame, buffer)
	callframe.Arg(frame, offset)
	callframe.Arg(frame, size_bytes)
	var r_ret = callframe.Ret[int64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RenderingDevice.Bind_buffer_clear, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = Error.Code(r_ret.Get())
	frame.Free()
	return ret
}

/*
Returns a copy of the data of the specified [param buffer], optionally [param offset_bytes] and [param size_bytes] can be set to copy only a portion of the buffer.
[b]Note:[/b] This method will block the GPU from working until the data is retrieved. Refer to [method buffer_get_data_async] for an alternative that returns the data in more performant way.
*/
//go:nosplit
func (self class) BufferGetData(buffer RID.Any, offset_bytes int64, size_bytes int64) Packed.Bytes { //gd:RenderingDevice.buffer_get_data
	var frame = callframe.New()
	callframe.Arg(frame, buffer)
	callframe.Arg(frame, offset_bytes)
	callframe.Arg(frame, size_bytes)
	var r_ret = callframe.Ret[gd.PackedPointers](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RenderingDevice.Bind_buffer_get_data, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = Packed.Bytes(Array.Through(gd.PackedProxy[gd.PackedByteArray, byte]{}, pointers.Pack(pointers.New[gd.PackedByteArray](r_ret.Get()))))
	frame.Free()
	return ret
}

/*
Asynchronous version of [method buffer_get_data]. RenderingDevice will call [param callback] in a certain amount of frames with the data the buffer had at the time of the request.
[b]Note:[/b] At the moment, the delay corresponds to the amount of frames specified by [member ProjectSettings.rendering/rendering_device/vsync/frame_queue_size].
[b]Note:[/b] Downloading large buffers can have a prohibitive cost for real-time even when using the asynchronous method due to hardware bandwidth limitations. When dealing with large resources, you can adjust settings such as [member ProjectSettings.rendering/rendering_device/staging_buffer/block_size_kb] to improve the transfer speed at the cost of extra memory.
[codeblock]
func _buffer_get_data_callback(array):
    value = array.decode_u32(0)

...

rd.buffer_get_data_async(buffer, _buffer_get_data_callback)
[/codeblock]
*/
//go:nosplit
func (self class) BufferGetDataAsync(buffer RID.Any, callback Callable.Function, offset_bytes int64, size_bytes int64) Error.Code { //gd:RenderingDevice.buffer_get_data_async
	var frame = callframe.New()
	callframe.Arg(frame, buffer)
	callframe.Arg(frame, pointers.Get(gd.InternalCallable(callback)))
	callframe.Arg(frame, offset_bytes)
	callframe.Arg(frame, size_bytes)
	var r_ret = callframe.Ret[int64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RenderingDevice.Bind_buffer_get_data_async, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = Error.Code(r_ret.Get())
	frame.Free()
	return ret
}

/*
Returns the address of the given [param buffer] which can be passed to shaders in any way to access underlying data. Buffer must have been created with this feature enabled.
[b]Note:[/b] You must check that the GPU supports this functionality by calling [method has_feature] with [constant SUPPORTS_BUFFER_DEVICE_ADDRESS] as a parameter.
*/
//go:nosplit
func (self class) BufferGetDeviceAddress(buffer RID.Any) int64 { //gd:RenderingDevice.buffer_get_device_address
	var frame = callframe.New()
	callframe.Arg(frame, buffer)
	var r_ret = callframe.Ret[int64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RenderingDevice.Bind_buffer_get_device_address, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Creates a new render pipeline. It can be accessed with the RID that is returned.
Once finished with your RID, you will want to free the RID using the RenderingDevice's [method free_rid] method.
*/
//go:nosplit
func (self class) RenderPipelineCreate(shader RID.Any, framebuffer_format int64, vertex_format int64, primitive gdclass.RenderingDeviceRenderPrimitive, rasterization_state [1]gdclass.RDPipelineRasterizationState, multisample_state [1]gdclass.RDPipelineMultisampleState, stencil_state [1]gdclass.RDPipelineDepthStencilState, color_blend_state [1]gdclass.RDPipelineColorBlendState, dynamic_state_flags gdclass.RenderingDevicePipelineDynamicStateFlags, for_render_pass int64, specialization_constants Array.Contains[[1]gdclass.RDPipelineSpecializationConstant]) RID.Any { //gd:RenderingDevice.render_pipeline_create
	var frame = callframe.New()
	callframe.Arg(frame, shader)
	callframe.Arg(frame, framebuffer_format)
	callframe.Arg(frame, vertex_format)
	callframe.Arg(frame, primitive)
	callframe.Arg(frame, pointers.Get(rasterization_state[0])[0])
	callframe.Arg(frame, pointers.Get(multisample_state[0])[0])
	callframe.Arg(frame, pointers.Get(stencil_state[0])[0])
	callframe.Arg(frame, pointers.Get(color_blend_state[0])[0])
	callframe.Arg(frame, dynamic_state_flags)
	callframe.Arg(frame, for_render_pass)
	callframe.Arg(frame, pointers.Get(gd.InternalArray(specialization_constants)))
	var r_ret = callframe.Ret[RID.Any](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RenderingDevice.Bind_render_pipeline_create, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns [code]true[/code] if the render pipeline specified by the [param render_pipeline] RID is valid, [code]false[/code] otherwise.
*/
//go:nosplit
func (self class) RenderPipelineIsValid(render_pipeline RID.Any) bool { //gd:RenderingDevice.render_pipeline_is_valid
	var frame = callframe.New()
	callframe.Arg(frame, render_pipeline)
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RenderingDevice.Bind_render_pipeline_is_valid, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Creates a new compute pipeline. It can be accessed with the RID that is returned.
Once finished with your RID, you will want to free the RID using the RenderingDevice's [method free_rid] method.
*/
//go:nosplit
func (self class) ComputePipelineCreate(shader RID.Any, specialization_constants Array.Contains[[1]gdclass.RDPipelineSpecializationConstant]) RID.Any { //gd:RenderingDevice.compute_pipeline_create
	var frame = callframe.New()
	callframe.Arg(frame, shader)
	callframe.Arg(frame, pointers.Get(gd.InternalArray(specialization_constants)))
	var r_ret = callframe.Ret[RID.Any](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RenderingDevice.Bind_compute_pipeline_create, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns [code]true[/code] if the compute pipeline specified by the [param compute_pipeline] RID is valid, [code]false[/code] otherwise.
*/
//go:nosplit
func (self class) ComputePipelineIsValid(compute_pipeline RID.Any) bool { //gd:RenderingDevice.compute_pipeline_is_valid
	var frame = callframe.New()
	callframe.Arg(frame, compute_pipeline)
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RenderingDevice.Bind_compute_pipeline_is_valid, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the window width matching the graphics API context for the given window ID (in pixels). Despite the parameter being named [param screen], this returns the [i]window[/i] size. See also [method screen_get_height].
[b]Note:[/b] Only the main [RenderingDevice] returned by [method RenderingServer.get_rendering_device] has a width. If called on a local [RenderingDevice], this method prints an error and returns [constant INVALID_ID].
*/
//go:nosplit
func (self class) ScreenGetWidth(screen int64) int64 { //gd:RenderingDevice.screen_get_width
	var frame = callframe.New()
	callframe.Arg(frame, screen)
	var r_ret = callframe.Ret[int64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RenderingDevice.Bind_screen_get_width, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the window height matching the graphics API context for the given window ID (in pixels). Despite the parameter being named [param screen], this returns the [i]window[/i] size. See also [method screen_get_width].
[b]Note:[/b] Only the main [RenderingDevice] returned by [method RenderingServer.get_rendering_device] has a height. If called on a local [RenderingDevice], this method prints an error and returns [constant INVALID_ID].
*/
//go:nosplit
func (self class) ScreenGetHeight(screen int64) int64 { //gd:RenderingDevice.screen_get_height
	var frame = callframe.New()
	callframe.Arg(frame, screen)
	var r_ret = callframe.Ret[int64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RenderingDevice.Bind_screen_get_height, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the framebuffer format of the given screen.
[b]Note:[/b] Only the main [RenderingDevice] returned by [method RenderingServer.get_rendering_device] has a format. If called on a local [RenderingDevice], this method prints an error and returns [constant INVALID_ID].
*/
//go:nosplit
func (self class) ScreenGetFramebufferFormat(screen int64) int64 { //gd:RenderingDevice.screen_get_framebuffer_format
	var frame = callframe.New()
	callframe.Arg(frame, screen)
	var r_ret = callframe.Ret[int64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RenderingDevice.Bind_screen_get_framebuffer_format, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
High-level variant of [method draw_list_begin], with the parameters automatically being adjusted for drawing onto the window specified by the [param screen] ID.
[b]Note:[/b] Cannot be used with local RenderingDevices, as these don't have a screen. If called on a local RenderingDevice, [method draw_list_begin_for_screen] returns [constant INVALID_ID].
*/
//go:nosplit
func (self class) DrawListBeginForScreen(screen int64, clear_color Color.RGBA) int64 { //gd:RenderingDevice.draw_list_begin_for_screen
	var frame = callframe.New()
	callframe.Arg(frame, screen)
	callframe.Arg(frame, clear_color)
	var r_ret = callframe.Ret[int64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RenderingDevice.Bind_draw_list_begin_for_screen, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Starts a list of raster drawing commands created with the [code]draw_*[/code] methods. The returned value should be passed to other [code]draw_list_*[/code] functions.
Multiple draw lists cannot be created at the same time; you must finish the previous draw list first using [method draw_list_end].
A simple drawing operation might look like this (code is not a complete example):
[codeblock]
var rd = RenderingDevice.new()
var clear_colors = PackedColorArray([Color(0, 0, 0, 0), Color(0, 0, 0, 0), Color(0, 0, 0, 0)])
var draw_list = rd.draw_list_begin(framebuffers[i], RenderingDevice.CLEAR_COLOR_ALL, clear_colors, true, 1.0f, true, 0, Rect2(), RenderingDevice.OPAQUE_PASS)

# Draw opaque.
rd.draw_list_bind_render_pipeline(draw_list, raster_pipeline)
rd.draw_list_bind_uniform_set(draw_list, raster_base_uniform, 0)
rd.draw_list_set_push_constant(draw_list, raster_push_constant, raster_push_constant.size())
rd.draw_list_draw(draw_list, false, 1, slice_triangle_count[i] * 3)
# Draw wire.
rd.draw_list_bind_render_pipeline(draw_list, raster_pipeline_wire)
rd.draw_list_bind_uniform_set(draw_list, raster_base_uniform, 0)
rd.draw_list_set_push_constant(draw_list, raster_push_constant, raster_push_constant.size())
rd.draw_list_draw(draw_list, false, 1, slice_triangle_count[i] * 3)

rd.draw_list_end()
[/codeblock]
The [param draw_flags] indicates if the texture attachments of the framebuffer should be cleared or ignored. Only one of the two flags can be used for each individual attachment. Ignoring an attachment means that any contents that existed before the draw list will be completely discarded, reducing the memory bandwidth used by the render pass but producing garbage results if the pixels aren't replaced. The default behavior allows the engine to figure out the right operation to use if the texture is discardable, which can result in increased performance. See [RDTextureFormat] or [method texture_set_discardable].
The [param breadcrumb] parameter can be an arbitrary 32-bit integer that is useful to diagnose GPU crashes. If Godot is built in dev or debug mode; when the GPU crashes Godot will dump all shaders that were being executed at the time of the crash and the breadcrumb is useful to diagnose what passes did those shaders belong to.
It does not affect rendering behavior and can be set to 0. It is recommended to use [enum BreadcrumbMarker] enumerations for consistency but it's not required. It is also possible to use bitwise operations to add extra data. e.g.
[codeblock]
rd.draw_list_begin(fb[i], RenderingDevice.CLEAR_COLOR_ALL, clear_colors, true, 1.0f, true, 0, Rect2(), RenderingDevice.OPAQUE_PASS | 5)
[/codeblock]
*/
//go:nosplit
func (self class) DrawListBegin(framebuffer RID.Any, draw_flags gdclass.RenderingDeviceDrawFlags, clear_color_values Packed.Array[Color.RGBA], clear_depth_value float64, clear_stencil_value int64, region Rect2.PositionSize, breadcrumb int64) int64 { //gd:RenderingDevice.draw_list_begin
	var frame = callframe.New()
	callframe.Arg(frame, framebuffer)
	callframe.Arg(frame, draw_flags)
	callframe.Arg(frame, pointers.Get(gd.InternalPacked[gd.PackedColorArray, Color.RGBA](clear_color_values)))
	callframe.Arg(frame, clear_depth_value)
	callframe.Arg(frame, clear_stencil_value)
	callframe.Arg(frame, region)
	callframe.Arg(frame, breadcrumb)
	var r_ret = callframe.Ret[int64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RenderingDevice.Bind_draw_list_begin, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
This method does nothing and always returns an empty [PackedInt64Array].
*/
//go:nosplit
func (self class) DrawListBeginSplit(framebuffer RID.Any, splits int64, initial_color_action gdclass.RenderingDeviceInitialAction, final_color_action gdclass.RenderingDeviceFinalAction, initial_depth_action gdclass.RenderingDeviceInitialAction, final_depth_action gdclass.RenderingDeviceFinalAction, clear_color_values Packed.Array[Color.RGBA], clear_depth float64, clear_stencil int64, region Rect2.PositionSize, storage_textures Array.Contains[RID.Any]) Packed.Array[int64] { //gd:RenderingDevice.draw_list_begin_split
	var frame = callframe.New()
	callframe.Arg(frame, framebuffer)
	callframe.Arg(frame, splits)
	callframe.Arg(frame, initial_color_action)
	callframe.Arg(frame, final_color_action)
	callframe.Arg(frame, initial_depth_action)
	callframe.Arg(frame, final_depth_action)
	callframe.Arg(frame, pointers.Get(gd.InternalPacked[gd.PackedColorArray, Color.RGBA](clear_color_values)))
	callframe.Arg(frame, clear_depth)
	callframe.Arg(frame, clear_stencil)
	callframe.Arg(frame, region)
	callframe.Arg(frame, pointers.Get(gd.InternalArray(storage_textures)))
	var r_ret = callframe.Ret[gd.PackedPointers](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RenderingDevice.Bind_draw_list_begin_split, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = Packed.Array[int64](Array.Through(gd.PackedProxy[gd.PackedInt64Array, int64]{}, pointers.Pack(pointers.New[gd.PackedStringArray](r_ret.Get()))))
	frame.Free()
	return ret
}

/*
Sets blend constants for the specified [param draw_list] to [param color]. Blend constants are used only if the graphics pipeline is created with [constant DYNAMIC_STATE_BLEND_CONSTANTS] flag set.
*/
//go:nosplit
func (self class) DrawListSetBlendConstants(draw_list int64, color Color.RGBA) { //gd:RenderingDevice.draw_list_set_blend_constants
	var frame = callframe.New()
	callframe.Arg(frame, draw_list)
	callframe.Arg(frame, color)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RenderingDevice.Bind_draw_list_set_blend_constants, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Binds [param render_pipeline] to the specified [param draw_list].
*/
//go:nosplit
func (self class) DrawListBindRenderPipeline(draw_list int64, render_pipeline RID.Any) { //gd:RenderingDevice.draw_list_bind_render_pipeline
	var frame = callframe.New()
	callframe.Arg(frame, draw_list)
	callframe.Arg(frame, render_pipeline)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RenderingDevice.Bind_draw_list_bind_render_pipeline, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Binds [param uniform_set] to the specified [param draw_list]. A [param set_index] must also be specified, which is an identifier starting from [code]0[/code] that must match the one expected by the draw list.
*/
//go:nosplit
func (self class) DrawListBindUniformSet(draw_list int64, uniform_set RID.Any, set_index int64) { //gd:RenderingDevice.draw_list_bind_uniform_set
	var frame = callframe.New()
	callframe.Arg(frame, draw_list)
	callframe.Arg(frame, uniform_set)
	callframe.Arg(frame, set_index)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RenderingDevice.Bind_draw_list_bind_uniform_set, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Binds [param vertex_array] to the specified [param draw_list].
*/
//go:nosplit
func (self class) DrawListBindVertexArray(draw_list int64, vertex_array RID.Any) { //gd:RenderingDevice.draw_list_bind_vertex_array
	var frame = callframe.New()
	callframe.Arg(frame, draw_list)
	callframe.Arg(frame, vertex_array)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RenderingDevice.Bind_draw_list_bind_vertex_array, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Binds [param index_array] to the specified [param draw_list].
*/
//go:nosplit
func (self class) DrawListBindIndexArray(draw_list int64, index_array RID.Any) { //gd:RenderingDevice.draw_list_bind_index_array
	var frame = callframe.New()
	callframe.Arg(frame, draw_list)
	callframe.Arg(frame, index_array)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RenderingDevice.Bind_draw_list_bind_index_array, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Sets the push constant data to [param buffer] for the specified [param draw_list]. The shader determines how this binary data is used. The buffer's size in bytes must also be specified in [param size_bytes] (this can be obtained by calling the [method PackedByteArray.size] method on the passed [param buffer]).
*/
//go:nosplit
func (self class) DrawListSetPushConstant(draw_list int64, buffer Packed.Bytes, size_bytes int64) { //gd:RenderingDevice.draw_list_set_push_constant
	var frame = callframe.New()
	callframe.Arg(frame, draw_list)
	callframe.Arg(frame, pointers.Get(gd.InternalPacked[gd.PackedByteArray, byte](Packed.Array[byte](buffer))))
	callframe.Arg(frame, size_bytes)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RenderingDevice.Bind_draw_list_set_push_constant, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Submits [param draw_list] for rendering on the GPU. This is the raster equivalent to [method compute_list_dispatch].
*/
//go:nosplit
func (self class) DrawListDraw(draw_list int64, use_indices bool, instances int64, procedural_vertex_count int64) { //gd:RenderingDevice.draw_list_draw
	var frame = callframe.New()
	callframe.Arg(frame, draw_list)
	callframe.Arg(frame, use_indices)
	callframe.Arg(frame, instances)
	callframe.Arg(frame, procedural_vertex_count)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RenderingDevice.Bind_draw_list_draw, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Submits [param draw_list] for rendering on the GPU with the given parameters stored in the [param buffer] at [param offset]. Parameters being integers: vertex count, instance count, first vertex, first instance. And when using indices: index count, instance count, first index, vertex offset, first instance. Buffer must have been created with [constant STORAGE_BUFFER_USAGE_DISPATCH_INDIRECT] flag.
*/
//go:nosplit
func (self class) DrawListDrawIndirect(draw_list int64, use_indices bool, buffer RID.Any, offset int64, draw_count int64, stride int64) { //gd:RenderingDevice.draw_list_draw_indirect
	var frame = callframe.New()
	callframe.Arg(frame, draw_list)
	callframe.Arg(frame, use_indices)
	callframe.Arg(frame, buffer)
	callframe.Arg(frame, offset)
	callframe.Arg(frame, draw_count)
	callframe.Arg(frame, stride)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RenderingDevice.Bind_draw_list_draw_indirect, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Creates a scissor rectangle and enables it for the specified [param draw_list]. Scissor rectangles are used for clipping by discarding fragments that fall outside a specified rectangular portion of the screen. See also [method draw_list_disable_scissor].
[b]Note:[/b] The specified [param rect] is automatically intersected with the screen's dimensions, which means it cannot exceed the screen's dimensions.
*/
//go:nosplit
func (self class) DrawListEnableScissor(draw_list int64, rect Rect2.PositionSize) { //gd:RenderingDevice.draw_list_enable_scissor
	var frame = callframe.New()
	callframe.Arg(frame, draw_list)
	callframe.Arg(frame, rect)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RenderingDevice.Bind_draw_list_enable_scissor, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Removes and disables the scissor rectangle for the specified [param draw_list]. See also [method draw_list_enable_scissor].
*/
//go:nosplit
func (self class) DrawListDisableScissor(draw_list int64) { //gd:RenderingDevice.draw_list_disable_scissor
	var frame = callframe.New()
	callframe.Arg(frame, draw_list)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RenderingDevice.Bind_draw_list_disable_scissor, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Switches to the next draw pass.
*/
//go:nosplit
func (self class) DrawListSwitchToNextPass() int64 { //gd:RenderingDevice.draw_list_switch_to_next_pass
	var frame = callframe.New()
	var r_ret = callframe.Ret[int64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RenderingDevice.Bind_draw_list_switch_to_next_pass, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
This method does nothing and always returns an empty [PackedInt64Array].
*/
//go:nosplit
func (self class) DrawListSwitchToNextPassSplit(splits int64) Packed.Array[int64] { //gd:RenderingDevice.draw_list_switch_to_next_pass_split
	var frame = callframe.New()
	callframe.Arg(frame, splits)
	var r_ret = callframe.Ret[gd.PackedPointers](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RenderingDevice.Bind_draw_list_switch_to_next_pass_split, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = Packed.Array[int64](Array.Through(gd.PackedProxy[gd.PackedInt64Array, int64]{}, pointers.Pack(pointers.New[gd.PackedStringArray](r_ret.Get()))))
	frame.Free()
	return ret
}

/*
Finishes a list of raster drawing commands created with the [code]draw_*[/code] methods.
*/
//go:nosplit
func (self class) DrawListEnd() { //gd:RenderingDevice.draw_list_end
	var frame = callframe.New()
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RenderingDevice.Bind_draw_list_end, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Starts a list of compute commands created with the [code]compute_*[/code] methods. The returned value should be passed to other [code]compute_list_*[/code] functions.
Multiple compute lists cannot be created at the same time; you must finish the previous compute list first using [method compute_list_end].
A simple compute operation might look like this (code is not a complete example):
[codeblock]
var rd = RenderingDevice.new()
var compute_list = rd.compute_list_begin()

rd.compute_list_bind_compute_pipeline(compute_list, compute_shader_dilate_pipeline)
rd.compute_list_bind_uniform_set(compute_list, compute_base_uniform_set, 0)
rd.compute_list_bind_uniform_set(compute_list, dilate_uniform_set, 1)

for i in atlas_slices:
    rd.compute_list_set_push_constant(compute_list, push_constant, push_constant.size())
    rd.compute_list_dispatch(compute_list, group_size.x, group_size.y, group_size.z)
    # No barrier, let them run all together.

rd.compute_list_end()
[/codeblock]
*/
//go:nosplit
func (self class) ComputeListBegin() int64 { //gd:RenderingDevice.compute_list_begin
	var frame = callframe.New()
	var r_ret = callframe.Ret[int64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RenderingDevice.Bind_compute_list_begin, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Tells the GPU what compute pipeline to use when processing the compute list. If the shader has changed since the last time this function was called, Godot will unbind all descriptor sets and will re-bind them inside [method compute_list_dispatch].
*/
//go:nosplit
func (self class) ComputeListBindComputePipeline(compute_list int64, compute_pipeline RID.Any) { //gd:RenderingDevice.compute_list_bind_compute_pipeline
	var frame = callframe.New()
	callframe.Arg(frame, compute_list)
	callframe.Arg(frame, compute_pipeline)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RenderingDevice.Bind_compute_list_bind_compute_pipeline, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Sets the push constant data to [param buffer] for the specified [param compute_list]. The shader determines how this binary data is used. The buffer's size in bytes must also be specified in [param size_bytes] (this can be obtained by calling the [method PackedByteArray.size] method on the passed [param buffer]).
*/
//go:nosplit
func (self class) ComputeListSetPushConstant(compute_list int64, buffer Packed.Bytes, size_bytes int64) { //gd:RenderingDevice.compute_list_set_push_constant
	var frame = callframe.New()
	callframe.Arg(frame, compute_list)
	callframe.Arg(frame, pointers.Get(gd.InternalPacked[gd.PackedByteArray, byte](Packed.Array[byte](buffer))))
	callframe.Arg(frame, size_bytes)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RenderingDevice.Bind_compute_list_set_push_constant, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Binds the [param uniform_set] to this [param compute_list]. Godot ensures that all textures in the uniform set have the correct Vulkan access masks. If Godot had to change access masks of textures, it will raise a Vulkan image memory barrier.
*/
//go:nosplit
func (self class) ComputeListBindUniformSet(compute_list int64, uniform_set RID.Any, set_index int64) { //gd:RenderingDevice.compute_list_bind_uniform_set
	var frame = callframe.New()
	callframe.Arg(frame, compute_list)
	callframe.Arg(frame, uniform_set)
	callframe.Arg(frame, set_index)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RenderingDevice.Bind_compute_list_bind_uniform_set, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Submits the compute list for processing on the GPU. This is the compute equivalent to [method draw_list_draw].
*/
//go:nosplit
func (self class) ComputeListDispatch(compute_list int64, x_groups int64, y_groups int64, z_groups int64) { //gd:RenderingDevice.compute_list_dispatch
	var frame = callframe.New()
	callframe.Arg(frame, compute_list)
	callframe.Arg(frame, x_groups)
	callframe.Arg(frame, y_groups)
	callframe.Arg(frame, z_groups)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RenderingDevice.Bind_compute_list_dispatch, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Submits the compute list for processing on the GPU with the given group counts stored in the [param buffer] at [param offset]. Buffer must have been created with [constant STORAGE_BUFFER_USAGE_DISPATCH_INDIRECT] flag.
*/
//go:nosplit
func (self class) ComputeListDispatchIndirect(compute_list int64, buffer RID.Any, offset int64) { //gd:RenderingDevice.compute_list_dispatch_indirect
	var frame = callframe.New()
	callframe.Arg(frame, compute_list)
	callframe.Arg(frame, buffer)
	callframe.Arg(frame, offset)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RenderingDevice.Bind_compute_list_dispatch_indirect, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Raises a Vulkan compute barrier in the specified [param compute_list].
*/
//go:nosplit
func (self class) ComputeListAddBarrier(compute_list int64) { //gd:RenderingDevice.compute_list_add_barrier
	var frame = callframe.New()
	callframe.Arg(frame, compute_list)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RenderingDevice.Bind_compute_list_add_barrier, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Finishes a list of compute commands created with the [code]compute_*[/code] methods.
*/
//go:nosplit
func (self class) ComputeListEnd() { //gd:RenderingDevice.compute_list_end
	var frame = callframe.New()
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RenderingDevice.Bind_compute_list_end, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Tries to free an object in the RenderingDevice. To avoid memory leaks, this should be called after using an object as memory management does not occur automatically when using RenderingDevice directly.
*/
//go:nosplit
func (self class) FreeRid(rid RID.Any) { //gd:RenderingDevice.free_rid
	var frame = callframe.New()
	callframe.Arg(frame, rid)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RenderingDevice.Bind_free_rid, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Creates a timestamp marker with the specified [param name]. This is used for performance reporting with the [method get_captured_timestamp_cpu_time], [method get_captured_timestamp_gpu_time] and [method get_captured_timestamp_name] methods.
*/
//go:nosplit
func (self class) CaptureTimestamp(name String.Readable) { //gd:RenderingDevice.capture_timestamp
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalString(name)))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RenderingDevice.Bind_capture_timestamp, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns the total number of timestamps (rendering steps) available for profiling.
*/
//go:nosplit
func (self class) GetCapturedTimestampsCount() int64 { //gd:RenderingDevice.get_captured_timestamps_count
	var frame = callframe.New()
	var r_ret = callframe.Ret[int64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RenderingDevice.Bind_get_captured_timestamps_count, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the index of the last frame rendered that has rendering timestamps available for querying.
*/
//go:nosplit
func (self class) GetCapturedTimestampsFrame() int64 { //gd:RenderingDevice.get_captured_timestamps_frame
	var frame = callframe.New()
	var r_ret = callframe.Ret[int64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RenderingDevice.Bind_get_captured_timestamps_frame, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the timestamp in GPU time for the rendering step specified by [param index] (in microseconds since the engine started). See also [method get_captured_timestamp_cpu_time] and [method capture_timestamp].
*/
//go:nosplit
func (self class) GetCapturedTimestampGpuTime(index int64) int64 { //gd:RenderingDevice.get_captured_timestamp_gpu_time
	var frame = callframe.New()
	callframe.Arg(frame, index)
	var r_ret = callframe.Ret[int64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RenderingDevice.Bind_get_captured_timestamp_gpu_time, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the timestamp in CPU time for the rendering step specified by [param index] (in microseconds since the engine started). See also [method get_captured_timestamp_gpu_time] and [method capture_timestamp].
*/
//go:nosplit
func (self class) GetCapturedTimestampCpuTime(index int64) int64 { //gd:RenderingDevice.get_captured_timestamp_cpu_time
	var frame = callframe.New()
	callframe.Arg(frame, index)
	var r_ret = callframe.Ret[int64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RenderingDevice.Bind_get_captured_timestamp_cpu_time, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the timestamp's name for the rendering step specified by [param index]. See also [method capture_timestamp].
*/
//go:nosplit
func (self class) GetCapturedTimestampName(index int64) String.Readable { //gd:RenderingDevice.get_captured_timestamp_name
	var frame = callframe.New()
	callframe.Arg(frame, index)
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RenderingDevice.Bind_get_captured_timestamp_name, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = String.Via(gd.StringProxy{}, pointers.Pack(pointers.New[gd.String](r_ret.Get())))
	frame.Free()
	return ret
}

/*
Returns [code]true[/code] if the [param feature] is supported by the GPU.
*/
//go:nosplit
func (self class) HasFeature(feature gdclass.RenderingDeviceFeatures) bool { //gd:RenderingDevice.has_feature
	var frame = callframe.New()
	callframe.Arg(frame, feature)
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RenderingDevice.Bind_has_feature, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the value of the specified [param limit]. This limit varies depending on the current graphics hardware (and sometimes the driver version). If the given limit is exceeded, rendering errors will occur.
Limits for various graphics hardware can be found in the [url=https://vulkan.gpuinfo.org/]Vulkan Hardware Database[/url].
*/
//go:nosplit
func (self class) LimitGet(limit gdclass.RenderingDeviceLimit) int64 { //gd:RenderingDevice.limit_get
	var frame = callframe.New()
	callframe.Arg(frame, limit)
	var r_ret = callframe.Ret[int64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RenderingDevice.Bind_limit_get, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the frame count kept by the graphics API. Higher values result in higher input lag, but with more consistent throughput. For the main [RenderingDevice], frames are cycled (usually 3 with triple-buffered V-Sync enabled). However, local [RenderingDevice]s only have 1 frame.
*/
//go:nosplit
func (self class) GetFrameDelay() int64 { //gd:RenderingDevice.get_frame_delay
	var frame = callframe.New()
	var r_ret = callframe.Ret[int64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RenderingDevice.Bind_get_frame_delay, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Pushes the frame setup and draw command buffers then marks the local device as currently processing (which allows calling [method sync]).
[b]Note:[/b] Only available in local RenderingDevices.
*/
//go:nosplit
func (self class) Submit() { //gd:RenderingDevice.submit
	var frame = callframe.New()
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RenderingDevice.Bind_submit, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Forces a synchronization between the CPU and GPU, which may be required in certain cases. Only call this when needed, as CPU-GPU synchronization has a performance cost.
[b]Note:[/b] Only available in local RenderingDevices.
[b]Note:[/b] [method sync] can only be called after a [method submit].
*/
//go:nosplit
func (self class) Sync() { //gd:RenderingDevice.sync
	var frame = callframe.New()
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RenderingDevice.Bind_sync, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
This method does nothing.
*/
//go:nosplit
func (self class) Barrier(from gdclass.RenderingDeviceBarrierMask, to gdclass.RenderingDeviceBarrierMask) { //gd:RenderingDevice.barrier
	var frame = callframe.New()
	callframe.Arg(frame, from)
	callframe.Arg(frame, to)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RenderingDevice.Bind_barrier, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
This method does nothing.
*/
//go:nosplit
func (self class) FullBarrier() { //gd:RenderingDevice.full_barrier
	var frame = callframe.New()
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RenderingDevice.Bind_full_barrier, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Create a new local [RenderingDevice]. This is most useful for performing compute operations on the GPU independently from the rest of the engine.
*/
//go:nosplit
func (self class) CreateLocalDevice() [1]gdclass.RenderingDevice { //gd:RenderingDevice.create_local_device
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RenderingDevice.Bind_create_local_device, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = [1]gdclass.RenderingDevice{gd.PointerWithOwnershipTransferredToGo[gdclass.RenderingDevice](r_ret.Get())}
	frame.Free()
	return ret
}

/*
Sets the resource name for [param id] to [param name]. This is used for debugging with third-party tools such as [url=https://renderdoc.org/]RenderDoc[/url].
The following types of resources can be named: texture, sampler, vertex buffer, index buffer, uniform buffer, texture buffer, storage buffer, uniform set buffer, shader, render pipeline and compute pipeline. Framebuffers cannot be named. Attempting to name an incompatible resource type will print an error.
[b]Note:[/b] Resource names are only set when the engine runs in verbose mode ([method OS.is_stdout_verbose] = [code]true[/code]), or when using an engine build compiled with the [code]dev_mode=yes[/code] SCons option. The graphics driver must also support the [code]VK_EXT_DEBUG_UTILS_EXTENSION_NAME[/code] Vulkan extension for named resources to work.
*/
//go:nosplit
func (self class) SetResourceName(id RID.Any, name String.Readable) { //gd:RenderingDevice.set_resource_name
	var frame = callframe.New()
	callframe.Arg(frame, id)
	callframe.Arg(frame, pointers.Get(gd.InternalString(name)))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RenderingDevice.Bind_set_resource_name, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Create a command buffer debug label region that can be displayed in third-party tools such as [url=https://renderdoc.org/]RenderDoc[/url]. All regions must be ended with a [method draw_command_end_label] call. When viewed from the linear series of submissions to a single queue, calls to [method draw_command_begin_label] and [method draw_command_end_label] must be matched and balanced.
The [code]VK_EXT_DEBUG_UTILS_EXTENSION_NAME[/code] Vulkan extension must be available and enabled for command buffer debug label region to work. See also [method draw_command_end_label].
*/
//go:nosplit
func (self class) DrawCommandBeginLabel(name String.Readable, color Color.RGBA) { //gd:RenderingDevice.draw_command_begin_label
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalString(name)))
	callframe.Arg(frame, color)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RenderingDevice.Bind_draw_command_begin_label, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
This method does nothing.
*/
//go:nosplit
func (self class) DrawCommandInsertLabel(name String.Readable, color Color.RGBA) { //gd:RenderingDevice.draw_command_insert_label
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalString(name)))
	callframe.Arg(frame, color)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RenderingDevice.Bind_draw_command_insert_label, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Ends the command buffer debug label region started by a [method draw_command_begin_label] call.
*/
//go:nosplit
func (self class) DrawCommandEndLabel() { //gd:RenderingDevice.draw_command_end_label
	var frame = callframe.New()
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RenderingDevice.Bind_draw_command_end_label, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns the vendor of the video adapter (e.g. "NVIDIA Corporation"). Equivalent to [method RenderingServer.get_video_adapter_vendor]. See also [method get_device_name].
*/
//go:nosplit
func (self class) GetDeviceVendorName() String.Readable { //gd:RenderingDevice.get_device_vendor_name
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RenderingDevice.Bind_get_device_vendor_name, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = String.Via(gd.StringProxy{}, pointers.Pack(pointers.New[gd.String](r_ret.Get())))
	frame.Free()
	return ret
}

/*
Returns the name of the video adapter (e.g. "GeForce GTX 1080/PCIe/SSE2"). Equivalent to [method RenderingServer.get_video_adapter_name]. See also [method get_device_vendor_name].
*/
//go:nosplit
func (self class) GetDeviceName() String.Readable { //gd:RenderingDevice.get_device_name
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RenderingDevice.Bind_get_device_name, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = String.Via(gd.StringProxy{}, pointers.Pack(pointers.New[gd.String](r_ret.Get())))
	frame.Free()
	return ret
}

/*
Returns the universally unique identifier for the pipeline cache. This is used to cache shader files on disk, which avoids shader recompilations on subsequent engine runs. This UUID varies depending on the graphics card model, but also the driver version. Therefore, updating graphics drivers will invalidate the shader cache.
*/
//go:nosplit
func (self class) GetDevicePipelineCacheUuid() String.Readable { //gd:RenderingDevice.get_device_pipeline_cache_uuid
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RenderingDevice.Bind_get_device_pipeline_cache_uuid, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = String.Via(gd.StringProxy{}, pointers.Pack(pointers.New[gd.String](r_ret.Get())))
	frame.Free()
	return ret
}

/*
Returns the memory usage in bytes corresponding to the given [param type]. When using Vulkan, these statistics are calculated by [url=https://github.com/GPUOpen-LibrariesAndSDKs/VulkanMemoryAllocator]Vulkan Memory Allocator[/url].
*/
//go:nosplit
func (self class) GetMemoryUsage(atype gdclass.RenderingDeviceMemoryType) int64 { //gd:RenderingDevice.get_memory_usage
	var frame = callframe.New()
	callframe.Arg(frame, atype)
	var r_ret = callframe.Ret[int64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RenderingDevice.Bind_get_memory_usage, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the unique identifier of the driver [param resource] for the specified [param rid]. Some driver resource types ignore the specified [param rid] (see [enum DriverResource] descriptions). [param index] is always ignored but must be specified anyway.
*/
//go:nosplit
func (self class) GetDriverResource(resource gdclass.RenderingDeviceDriverResource, rid RID.Any, index int64) int64 { //gd:RenderingDevice.get_driver_resource
	var frame = callframe.New()
	callframe.Arg(frame, resource)
	callframe.Arg(frame, rid)
	callframe.Arg(frame, index)
	var r_ret = callframe.Ret[int64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RenderingDevice.Bind_get_driver_resource, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns a string with a performance report from the past frame. Updates every frame.
*/
//go:nosplit
func (self class) GetPerfReport() String.Readable { //gd:RenderingDevice.get_perf_report
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RenderingDevice.Bind_get_perf_report, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = String.Via(gd.StringProxy{}, pointers.Pack(pointers.New[gd.String](r_ret.Get())))
	frame.Free()
	return ret
}

/*
Returns string report in CSV format using the following methods:
- [method get_tracked_object_name]
- [method get_tracked_object_type_count]
- [method get_driver_total_memory]
- [method get_driver_allocation_count]
- [method get_driver_memory_by_object_type]
- [method get_driver_allocs_by_object_type]
- [method get_device_total_memory]
- [method get_device_allocation_count]
- [method get_device_memory_by_object_type]
- [method get_device_allocs_by_object_type]
This is only used by Vulkan in debug builds. Godot must also be started with the [code]--extra-gpu-memory-tracking[/code] [url=$DOCS_URL/tutorials/editor/command_line_tutorial.html]command line argument[/url].
*/
//go:nosplit
func (self class) GetDriverAndDeviceMemoryReport() String.Readable { //gd:RenderingDevice.get_driver_and_device_memory_report
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RenderingDevice.Bind_get_driver_and_device_memory_report, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = String.Via(gd.StringProxy{}, pointers.Pack(pointers.New[gd.String](r_ret.Get())))
	frame.Free()
	return ret
}

/*
Returns the name of the type of object for the given [param type_index]. This value must be in range [code][0; get_tracked_object_type_count - 1][/code]. If [method get_tracked_object_type_count] is 0, then type argument is ignored and always returns the same string.
The return value is important because it gives meaning to the types passed to [method get_driver_memory_by_object_type], [method get_driver_allocs_by_object_type], [method get_device_memory_by_object_type], and [method get_device_allocs_by_object_type]. Examples of strings it can return (not exhaustive):
- DEVICE_MEMORY
- PIPELINE_CACHE
- SWAPCHAIN_KHR
- COMMAND_POOL
Thus if e.g. [code]get_tracked_object_name(5)[/code] returns "COMMAND_POOL", then [code]get_device_memory_by_object_type(5)[/code] returns the bytes used by the GPU for command pools.
This is only used by Vulkan in debug builds. Godot must also be started with the [code]--extra-gpu-memory-tracking[/code] [url=$DOCS_URL/tutorials/editor/command_line_tutorial.html]command line argument[/url].
*/
//go:nosplit
func (self class) GetTrackedObjectName(type_index int64) String.Readable { //gd:RenderingDevice.get_tracked_object_name
	var frame = callframe.New()
	callframe.Arg(frame, type_index)
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RenderingDevice.Bind_get_tracked_object_name, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = String.Via(gd.StringProxy{}, pointers.Pack(pointers.New[gd.String](r_ret.Get())))
	frame.Free()
	return ret
}

/*
Returns how many types of trackable objects are.
This is only used by Vulkan in debug builds. Godot must also be started with the [code]--extra-gpu-memory-tracking[/code] [url=$DOCS_URL/tutorials/editor/command_line_tutorial.html]command line argument[/url].
*/
//go:nosplit
func (self class) GetTrackedObjectTypeCount() int64 { //gd:RenderingDevice.get_tracked_object_type_count
	var frame = callframe.New()
	var r_ret = callframe.Ret[int64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RenderingDevice.Bind_get_tracked_object_type_count, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns how much bytes the GPU driver is using for internal driver structures.
This is only used by Vulkan in debug builds and can return 0 when this information is not tracked or unknown.
*/
//go:nosplit
func (self class) GetDriverTotalMemory() int64 { //gd:RenderingDevice.get_driver_total_memory
	var frame = callframe.New()
	var r_ret = callframe.Ret[int64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RenderingDevice.Bind_get_driver_total_memory, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns how many allocations the GPU driver has performed for internal driver structures.
This is only used by Vulkan in debug builds and can return 0 when this information is not tracked or unknown.
*/
//go:nosplit
func (self class) GetDriverAllocationCount() int64 { //gd:RenderingDevice.get_driver_allocation_count
	var frame = callframe.New()
	var r_ret = callframe.Ret[int64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RenderingDevice.Bind_get_driver_allocation_count, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Same as [method get_driver_total_memory] but filtered for a given object type.
The type argument must be in range [code][0; get_tracked_object_type_count - 1][/code]. If [method get_tracked_object_type_count] is 0, then type argument is ignored and always returns 0.
This is only used by Vulkan in debug builds and can return 0 when this information is not tracked or unknown.
*/
//go:nosplit
func (self class) GetDriverMemoryByObjectType(atype int64) int64 { //gd:RenderingDevice.get_driver_memory_by_object_type
	var frame = callframe.New()
	callframe.Arg(frame, atype)
	var r_ret = callframe.Ret[int64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RenderingDevice.Bind_get_driver_memory_by_object_type, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Same as [method get_driver_allocation_count] but filtered for a given object type.
The type argument must be in range [code][0; get_tracked_object_type_count - 1][/code]. If [method get_tracked_object_type_count] is 0, then type argument is ignored and always returns 0.
This is only used by Vulkan in debug builds and can return 0 when this information is not tracked or unknown.
*/
//go:nosplit
func (self class) GetDriverAllocsByObjectType(atype int64) int64 { //gd:RenderingDevice.get_driver_allocs_by_object_type
	var frame = callframe.New()
	callframe.Arg(frame, atype)
	var r_ret = callframe.Ret[int64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RenderingDevice.Bind_get_driver_allocs_by_object_type, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns how much bytes the GPU is using.
This is only used by Vulkan in debug builds and can return 0 when this information is not tracked or unknown.
*/
//go:nosplit
func (self class) GetDeviceTotalMemory() int64 { //gd:RenderingDevice.get_device_total_memory
	var frame = callframe.New()
	var r_ret = callframe.Ret[int64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RenderingDevice.Bind_get_device_total_memory, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns how many allocations the GPU has performed for internal driver structures.
This is only used by Vulkan in debug builds and can return 0 when this information is not tracked or unknown.
*/
//go:nosplit
func (self class) GetDeviceAllocationCount() int64 { //gd:RenderingDevice.get_device_allocation_count
	var frame = callframe.New()
	var r_ret = callframe.Ret[int64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RenderingDevice.Bind_get_device_allocation_count, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Same as [method get_device_total_memory] but filtered for a given object type.
The type argument must be in range [code][0; get_tracked_object_type_count - 1][/code]. If [method get_tracked_object_type_count] is 0, then type argument is ignored and always returns 0.
This is only used by Vulkan in debug builds and can return 0 when this information is not tracked or unknown.
*/
//go:nosplit
func (self class) GetDeviceMemoryByObjectType(atype int64) int64 { //gd:RenderingDevice.get_device_memory_by_object_type
	var frame = callframe.New()
	callframe.Arg(frame, atype)
	var r_ret = callframe.Ret[int64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RenderingDevice.Bind_get_device_memory_by_object_type, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Same as [method get_device_allocation_count] but filtered for a given object type.
The type argument must be in range [code][0; get_tracked_object_type_count - 1][/code]. If [method get_tracked_object_type_count] is 0, then type argument is ignored and always returns 0.
This is only used by Vulkan in debug builds and can return 0 when this information is not tracked or unknown.
*/
//go:nosplit
func (self class) GetDeviceAllocsByObjectType(atype int64) int64 { //gd:RenderingDevice.get_device_allocs_by_object_type
	var frame = callframe.New()
	callframe.Arg(frame, atype)
	var r_ret = callframe.Ret[int64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RenderingDevice.Bind_get_device_allocs_by_object_type, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self class) AsRenderingDevice() Advanced    { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsRenderingDevice() Instance { return *((*Instance)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(Object.Advanced(self.AsObject()), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(Object.Instance(self.AsObject()), name)
	}
}
func init() {
	gdclass.Register("RenderingDevice", func(ptr gd.Object) any {
		return [1]gdclass.RenderingDevice{*(*gdclass.RenderingDevice)(unsafe.Pointer(&ptr))}
	})
}

type DeviceType = gdclass.RenderingDeviceDeviceType //gd:RenderingDevice.DeviceType

const (
	/*Rendering device type does not match any of the other enum values or is unknown.*/
	DeviceTypeOther DeviceType = 0
	/*Rendering device is an integrated GPU, which is typically [i](but not always)[/i] slower than dedicated GPUs ([constant DEVICE_TYPE_DISCRETE_GPU]). On Android and iOS, the rendering device type is always considered to be [constant DEVICE_TYPE_INTEGRATED_GPU].*/
	DeviceTypeIntegratedGpu DeviceType = 1
	/*Rendering device is a dedicated GPU, which is typically [i](but not always)[/i] faster than integrated GPUs ([constant DEVICE_TYPE_INTEGRATED_GPU]).*/
	DeviceTypeDiscreteGpu DeviceType = 2
	/*Rendering device is an emulated GPU in a virtual environment. This is typically much slower than the host GPU, which means the expected performance level on a dedicated GPU will be roughly equivalent to [constant DEVICE_TYPE_INTEGRATED_GPU]. Virtual machine GPU passthrough (such as VFIO) will not report the device type as [constant DEVICE_TYPE_VIRTUAL_GPU]. Instead, the host GPU's device type will be reported as if the GPU was not emulated.*/
	DeviceTypeVirtualGpu DeviceType = 3
	/*Rendering device is provided by software emulation (such as Lavapipe or [url=https://github.com/google/swiftshader]SwiftShader[/url]). This is the slowest kind of rendering device available; it's typically much slower than [constant DEVICE_TYPE_INTEGRATED_GPU].*/
	DeviceTypeCpu DeviceType = 4
	/*Represents the size of the [enum DeviceType] enum.*/
	DeviceTypeMax DeviceType = 5
)

type DriverResource = gdclass.RenderingDeviceDriverResource //gd:RenderingDevice.DriverResource

const (
	/*Specific device object based on a physical device.
	  - Vulkan: Vulkan device driver resource ([code]VkDevice[/code]). ([code]rid[/code] argument doesn't apply.)*/
	DriverResourceLogicalDevice DriverResource = 0
	/*Physical device the specific logical device is based on.
	  - Vulkan: [code]VkDevice[/code]. ([code]rid[/code] argument doesn't apply.)*/
	DriverResourcePhysicalDevice DriverResource = 1
	/*Top-most graphics API entry object.
	  - Vulkan: [code]VkInstance[/code]. ([code]rid[/code] argument doesn't apply.)*/
	DriverResourceTopmostObject DriverResource = 2
	/*The main graphics-compute command queue.
	  - Vulkan: [code]VkQueue[/code]. ([code]rid[/code] argument doesn't apply.)*/
	DriverResourceCommandQueue DriverResource = 3
	/*The specific family the main queue belongs to.
	  - Vulkan: the queue family index, an [code]uint32_t[/code]. ([code]rid[/code] argument doesn't apply.)*/
	DriverResourceQueueFamily DriverResource = 4
	/*- Vulkan: [code]VkImage[/code].*/
	DriverResourceTexture DriverResource = 5
	/*The view of an owned or shared texture.
	  - Vulkan: [code]VkImageView[/code].*/
	DriverResourceTextureView DriverResource = 6
	/*The native id of the data format of the texture.
	  - Vulkan: [code]VkFormat[/code].*/
	DriverResourceTextureDataFormat DriverResource = 7
	/*- Vulkan: [code]VkSampler[/code].*/
	DriverResourceSampler DriverResource = 8
	/*- Vulkan: [code]VkDescriptorSet[/code].*/
	DriverResourceUniformSet DriverResource = 9
	/*Buffer of any kind of (storage, vertex, etc.).
	  - Vulkan: [code]VkBuffer[/code].*/
	DriverResourceBuffer DriverResource = 10
	/*- Vulkan: [code]VkPipeline[/code].*/
	DriverResourceComputePipeline DriverResource = 11
	/*- Vulkan: [code]VkPipeline[/code].*/
	DriverResourceRenderPipeline                 DriverResource = 12
	DriverResourceVulkanDevice                   DriverResource = 0
	DriverResourceVulkanPhysicalDevice           DriverResource = 1
	DriverResourceVulkanInstance                 DriverResource = 2
	DriverResourceVulkanQueue                    DriverResource = 3
	DriverResourceVulkanQueueFamilyIndex         DriverResource = 4
	DriverResourceVulkanImage                    DriverResource = 5
	DriverResourceVulkanImageView                DriverResource = 6
	DriverResourceVulkanImageNativeTextureFormat DriverResource = 7
	DriverResourceVulkanSampler                  DriverResource = 8
	DriverResourceVulkanDescriptorSet            DriverResource = 9
	DriverResourceVulkanBuffer                   DriverResource = 10
	DriverResourceVulkanComputePipeline          DriverResource = 11
	DriverResourceVulkanRenderPipeline           DriverResource = 12
)

type DataFormat = gdclass.RenderingDeviceDataFormat //gd:RenderingDevice.DataFormat

const (
	/*4-bit-per-channel red/green channel data format, packed into 8 bits. Values are in the [code][0.0, 1.0][/code] range.
	  [b]Note:[/b] More information on all data formats can be found on the [url=https://registry.khronos.org/vulkan/specs/1.1/html/vkspec.html#_identification_of_formats]Identification of formats[/url] section of the Vulkan specification, as well as the [url=https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFormat.html]VkFormat[/url] enum.*/
	DataFormatR4g4UnormPack8 DataFormat = 0
	/*4-bit-per-channel red/green/blue/alpha channel data format, packed into 16 bits. Values are in the [code][0.0, 1.0][/code] range.*/
	DataFormatR4g4b4a4UnormPack16 DataFormat = 1
	/*4-bit-per-channel blue/green/red/alpha channel data format, packed into 16 bits. Values are in the [code][0.0, 1.0][/code] range.*/
	DataFormatB4g4r4a4UnormPack16 DataFormat = 2
	/*Red/green/blue channel data format with 5 bits of red, 6 bits of green and 5 bits of blue, packed into 16 bits. Values are in the [code][0.0, 1.0][/code] range.*/
	DataFormatR5g6b5UnormPack16 DataFormat = 3
	/*Blue/green/red channel data format with 5 bits of blue, 6 bits of green and 5 bits of red, packed into 16 bits. Values are in the [code][0.0, 1.0][/code] range.*/
	DataFormatB5g6r5UnormPack16 DataFormat = 4
	/*Red/green/blue/alpha channel data format with 5 bits of red, 6 bits of green, 5 bits of blue and 1 bit of alpha, packed into 16 bits. Values are in the [code][0.0, 1.0][/code] range.*/
	DataFormatR5g5b5a1UnormPack16 DataFormat = 5
	/*Blue/green/red/alpha channel data format with 5 bits of blue, 6 bits of green, 5 bits of red and 1 bit of alpha, packed into 16 bits. Values are in the [code][0.0, 1.0][/code] range.*/
	DataFormatB5g5r5a1UnormPack16 DataFormat = 6
	/*Alpha/red/green/blue channel data format with 1 bit of alpha, 5 bits of red, 6 bits of green and 5 bits of blue, packed into 16 bits. Values are in the [code][0.0, 1.0][/code] range.*/
	DataFormatA1r5g5b5UnormPack16 DataFormat = 7
	/*8-bit-per-channel unsigned floating-point red channel data format with normalized value. Values are in the [code][0.0, 1.0][/code] range.*/
	DataFormatR8Unorm DataFormat = 8
	/*8-bit-per-channel signed floating-point red channel data format with normalized value. Values are in the [code][-1.0, 1.0][/code] range.*/
	DataFormatR8Snorm DataFormat = 9
	/*8-bit-per-channel unsigned floating-point red channel data format with scaled value (value is converted from integer to float). Values are in the [code][0.0, 255.0][/code] range.*/
	DataFormatR8Uscaled DataFormat = 10
	/*8-bit-per-channel signed floating-point red channel data format with scaled value (value is converted from integer to float). Values are in the [code][-127.0, 127.0][/code] range.*/
	DataFormatR8Sscaled DataFormat = 11
	/*8-bit-per-channel unsigned integer red channel data format. Values are in the [code][0, 255][/code] range.*/
	DataFormatR8Uint DataFormat = 12
	/*8-bit-per-channel signed integer red channel data format. Values are in the [code][-127, 127][/code] range.*/
	DataFormatR8Sint DataFormat = 13
	/*8-bit-per-channel unsigned floating-point red channel data format with normalized value and non-linear sRGB encoding. Values are in the [code][0.0, 1.0][/code] range.*/
	DataFormatR8Srgb DataFormat = 14
	/*8-bit-per-channel unsigned floating-point red/green channel data format with normalized value. Values are in the [code][0.0, 1.0][/code] range.*/
	DataFormatR8g8Unorm DataFormat = 15
	/*8-bit-per-channel signed floating-point red/green channel data format with normalized value. Values are in the [code][-1.0, 1.0][/code] range.*/
	DataFormatR8g8Snorm DataFormat = 16
	/*8-bit-per-channel unsigned floating-point red/green channel data format with scaled value (value is converted from integer to float). Values are in the [code][0.0, 255.0][/code] range.*/
	DataFormatR8g8Uscaled DataFormat = 17
	/*8-bit-per-channel signed floating-point red/green channel data format with scaled value (value is converted from integer to float). Values are in the [code][-127.0, 127.0][/code] range.*/
	DataFormatR8g8Sscaled DataFormat = 18
	/*8-bit-per-channel unsigned integer red/green channel data format. Values are in the [code][0, 255][/code] range.*/
	DataFormatR8g8Uint DataFormat = 19
	/*8-bit-per-channel signed integer red/green channel data format. Values are in the [code][-127, 127][/code] range.*/
	DataFormatR8g8Sint DataFormat = 20
	/*8-bit-per-channel unsigned floating-point red/green channel data format with normalized value and non-linear sRGB encoding. Values are in the [code][0.0, 1.0][/code] range.*/
	DataFormatR8g8Srgb DataFormat = 21
	/*8-bit-per-channel unsigned floating-point red/green/blue channel data format with normalized value. Values are in the [code][0.0, 1.0][/code] range.*/
	DataFormatR8g8b8Unorm DataFormat = 22
	/*8-bit-per-channel signed floating-point red/green/blue channel data format with normalized value. Values are in the [code][-1.0, 1.0][/code] range.*/
	DataFormatR8g8b8Snorm DataFormat = 23
	/*8-bit-per-channel unsigned floating-point red/green/blue channel data format with scaled value (value is converted from integer to float). Values are in the [code][0.0, 255.0][/code] range.*/
	DataFormatR8g8b8Uscaled DataFormat = 24
	/*8-bit-per-channel signed floating-point red/green/blue channel data format with scaled value (value is converted from integer to float). Values are in the [code][-127.0, 127.0][/code] range.*/
	DataFormatR8g8b8Sscaled DataFormat = 25
	/*8-bit-per-channel unsigned integer red/green/blue channel data format. Values are in the [code][0, 255][/code] range.*/
	DataFormatR8g8b8Uint DataFormat = 26
	/*8-bit-per-channel signed integer red/green/blue channel data format. Values are in the [code][-127, 127][/code] range.*/
	DataFormatR8g8b8Sint DataFormat = 27
	/*8-bit-per-channel unsigned floating-point red/green/blue/blue channel data format with normalized value and non-linear sRGB encoding. Values are in the [code][0.0, 1.0][/code] range.*/
	DataFormatR8g8b8Srgb DataFormat = 28
	/*8-bit-per-channel unsigned floating-point blue/green/red channel data format with normalized value. Values are in the [code][0.0, 1.0][/code] range.*/
	DataFormatB8g8r8Unorm DataFormat = 29
	/*8-bit-per-channel signed floating-point blue/green/red channel data format with normalized value. Values are in the [code][-1.0, 1.0][/code] range.*/
	DataFormatB8g8r8Snorm DataFormat = 30
	/*8-bit-per-channel unsigned floating-point blue/green/red channel data format with scaled value (value is converted from integer to float). Values are in the [code][0.0, 255.0][/code] range.*/
	DataFormatB8g8r8Uscaled DataFormat = 31
	/*8-bit-per-channel signed floating-point blue/green/red channel data format with scaled value (value is converted from integer to float). Values are in the [code][-127.0, 127.0][/code] range.*/
	DataFormatB8g8r8Sscaled DataFormat = 32
	/*8-bit-per-channel unsigned integer blue/green/red channel data format. Values are in the [code][0, 255][/code] range.*/
	DataFormatB8g8r8Uint DataFormat = 33
	/*8-bit-per-channel signed integer blue/green/red channel data format. Values are in the [code][-127, 127][/code] range.*/
	DataFormatB8g8r8Sint DataFormat = 34
	/*8-bit-per-channel unsigned floating-point blue/green/red data format with normalized value and non-linear sRGB encoding. Values are in the [code][0.0, 1.0][/code] range.*/
	DataFormatB8g8r8Srgb DataFormat = 35
	/*8-bit-per-channel unsigned floating-point red/green/blue/alpha channel data format with normalized value. Values are in the [code][0.0, 1.0][/code] range.*/
	DataFormatR8g8b8a8Unorm DataFormat = 36
	/*8-bit-per-channel signed floating-point red/green/blue/alpha channel data format with normalized value. Values are in the [code][-1.0, 1.0][/code] range.*/
	DataFormatR8g8b8a8Snorm DataFormat = 37
	/*8-bit-per-channel unsigned floating-point red/green/blue/alpha channel data format with scaled value (value is converted from integer to float). Values are in the [code][0.0, 255.0][/code] range.*/
	DataFormatR8g8b8a8Uscaled DataFormat = 38
	/*8-bit-per-channel signed floating-point red/green/blue/alpha channel data format with scaled value (value is converted from integer to float). Values are in the [code][-127.0, 127.0][/code] range.*/
	DataFormatR8g8b8a8Sscaled DataFormat = 39
	/*8-bit-per-channel unsigned integer red/green/blue/alpha channel data format. Values are in the [code][0, 255][/code] range.*/
	DataFormatR8g8b8a8Uint DataFormat = 40
	/*8-bit-per-channel signed integer red/green/blue/alpha channel data format. Values are in the [code][-127, 127][/code] range.*/
	DataFormatR8g8b8a8Sint DataFormat = 41
	/*8-bit-per-channel unsigned floating-point red/green/blue/alpha channel data format with normalized value and non-linear sRGB encoding. Values are in the [code][0.0, 1.0][/code] range.*/
	DataFormatR8g8b8a8Srgb DataFormat = 42
	/*8-bit-per-channel unsigned floating-point blue/green/red/alpha channel data format with normalized value. Values are in the [code][0.0, 1.0][/code] range.*/
	DataFormatB8g8r8a8Unorm DataFormat = 43
	/*8-bit-per-channel signed floating-point blue/green/red/alpha channel data format with normalized value. Values are in the [code][-1.0, 1.0][/code] range.*/
	DataFormatB8g8r8a8Snorm DataFormat = 44
	/*8-bit-per-channel unsigned floating-point blue/green/red/alpha channel data format with scaled value (value is converted from integer to float). Values are in the [code][0.0, 255.0][/code] range.*/
	DataFormatB8g8r8a8Uscaled DataFormat = 45
	/*8-bit-per-channel signed floating-point blue/green/red/alpha channel data format with scaled value (value is converted from integer to float). Values are in the [code][-127.0, 127.0][/code] range.*/
	DataFormatB8g8r8a8Sscaled DataFormat = 46
	/*8-bit-per-channel unsigned integer blue/green/red/alpha channel data format. Values are in the [code][0, 255][/code] range.*/
	DataFormatB8g8r8a8Uint DataFormat = 47
	/*8-bit-per-channel signed integer blue/green/red/alpha channel data format. Values are in the [code][-127, 127][/code] range.*/
	DataFormatB8g8r8a8Sint DataFormat = 48
	/*8-bit-per-channel unsigned floating-point blue/green/red/alpha channel data format with normalized value and non-linear sRGB encoding. Values are in the [code][0.0, 1.0][/code] range.*/
	DataFormatB8g8r8a8Srgb DataFormat = 49
	/*8-bit-per-channel unsigned floating-point alpha/red/green/blue channel data format with normalized value, packed in 32 bits. Values are in the [code][0.0, 1.0][/code] range.*/
	DataFormatA8b8g8r8UnormPack32 DataFormat = 50
	/*8-bit-per-channel signed floating-point alpha/red/green/blue channel data format with normalized value, packed in 32 bits. Values are in the [code][-1.0, 1.0][/code] range.*/
	DataFormatA8b8g8r8SnormPack32 DataFormat = 51
	/*8-bit-per-channel unsigned floating-point alpha/red/green/blue channel data format with scaled value (value is converted from integer to float), packed in 32 bits. Values are in the [code][0.0, 255.0][/code] range.*/
	DataFormatA8b8g8r8UscaledPack32 DataFormat = 52
	/*8-bit-per-channel signed floating-point alpha/red/green/blue channel data format with scaled value (value is converted from integer to float), packed in 32 bits. Values are in the [code][-127.0, 127.0][/code] range.*/
	DataFormatA8b8g8r8SscaledPack32 DataFormat = 53
	/*8-bit-per-channel unsigned integer alpha/red/green/blue channel data format, packed in 32 bits. Values are in the [code][0, 255][/code] range.*/
	DataFormatA8b8g8r8UintPack32 DataFormat = 54
	/*8-bit-per-channel signed integer alpha/red/green/blue channel data format, packed in 32 bits. Values are in the [code][-127, 127][/code] range.*/
	DataFormatA8b8g8r8SintPack32 DataFormat = 55
	/*8-bit-per-channel unsigned floating-point alpha/red/green/blue channel data format with normalized value and non-linear sRGB encoding, packed in 32 bits. Values are in the [code][0.0, 1.0][/code] range.*/
	DataFormatA8b8g8r8SrgbPack32 DataFormat = 56
	/*Unsigned floating-point alpha/red/green/blue channel data format with normalized value, packed in 32 bits. Format contains 2 bits of alpha, 10 bits of red, 10 bits of green and 10 bits of blue. Values are in the [code][0.0, 1.0][/code] range.*/
	DataFormatA2r10g10b10UnormPack32 DataFormat = 57
	/*Signed floating-point alpha/red/green/blue channel data format with normalized value, packed in 32 bits. Format contains 2 bits of alpha, 10 bits of red, 10 bits of green and 10 bits of blue. Values are in the [code][-1.0, 1.0][/code] range.*/
	DataFormatA2r10g10b10SnormPack32 DataFormat = 58
	/*Unsigned floating-point alpha/red/green/blue channel data format with normalized value, packed in 32 bits. Format contains 2 bits of alpha, 10 bits of red, 10 bits of green and 10 bits of blue. Values are in the [code][0.0, 1023.0][/code] range for red/green/blue and [code][0.0, 3.0][/code] for alpha.*/
	DataFormatA2r10g10b10UscaledPack32 DataFormat = 59
	/*Signed floating-point alpha/red/green/blue channel data format with normalized value, packed in 32 bits. Format contains 2 bits of alpha, 10 bits of red, 10 bits of green and 10 bits of blue. Values are in the [code][-511.0, 511.0][/code] range for red/green/blue and [code][-1.0, 1.0][/code] for alpha.*/
	DataFormatA2r10g10b10SscaledPack32 DataFormat = 60
	/*Unsigned integer alpha/red/green/blue channel data format with normalized value, packed in 32 bits. Format contains 2 bits of alpha, 10 bits of red, 10 bits of green and 10 bits of blue. Values are in the [code][0, 1023][/code] range for red/green/blue and [code][0, 3][/code] for alpha.*/
	DataFormatA2r10g10b10UintPack32 DataFormat = 61
	/*Signed integer alpha/red/green/blue channel data format with normalized value, packed in 32 bits. Format contains 2 bits of alpha, 10 bits of red, 10 bits of green and 10 bits of blue. Values are in the [code][-511, 511][/code] range for red/green/blue and [code][-1, 1][/code] for alpha.*/
	DataFormatA2r10g10b10SintPack32 DataFormat = 62
	/*Unsigned floating-point alpha/blue/green/red channel data format with normalized value, packed in 32 bits. Format contains 2 bits of alpha, 10 bits of blue, 10 bits of green and 10 bits of red. Values are in the [code][0.0, 1.0][/code] range.*/
	DataFormatA2b10g10r10UnormPack32 DataFormat = 63
	/*Signed floating-point alpha/blue/green/red channel data format with normalized value, packed in 32 bits. Format contains 2 bits of alpha, 10 bits of blue, 10 bits of green and 10 bits of red. Values are in the [code][-1.0, 1.0][/code] range.*/
	DataFormatA2b10g10r10SnormPack32 DataFormat = 64
	/*Unsigned floating-point alpha/blue/green/red channel data format with normalized value, packed in 32 bits. Format contains 2 bits of alpha, 10 bits of blue, 10 bits of green and 10 bits of red. Values are in the [code][0.0, 1023.0][/code] range for blue/green/red and [code][0.0, 3.0][/code] for alpha.*/
	DataFormatA2b10g10r10UscaledPack32 DataFormat = 65
	/*Signed floating-point alpha/blue/green/red channel data format with normalized value, packed in 32 bits. Format contains 2 bits of alpha, 10 bits of blue, 10 bits of green and 10 bits of red. Values are in the [code][-511.0, 511.0][/code] range for blue/green/red and [code][-1.0, 1.0][/code] for alpha.*/
	DataFormatA2b10g10r10SscaledPack32 DataFormat = 66
	/*Unsigned integer alpha/blue/green/red channel data format with normalized value, packed in 32 bits. Format contains 2 bits of alpha, 10 bits of blue, 10 bits of green and 10 bits of red. Values are in the [code][0, 1023][/code] range for blue/green/red and [code][0, 3][/code] for alpha.*/
	DataFormatA2b10g10r10UintPack32 DataFormat = 67
	/*Signed integer alpha/blue/green/red channel data format with normalized value, packed in 32 bits. Format contains 2 bits of alpha, 10 bits of blue, 10 bits of green and 10 bits of red. Values are in the [code][-511, 511][/code] range for blue/green/red and [code][-1, 1][/code] for alpha.*/
	DataFormatA2b10g10r10SintPack32 DataFormat = 68
	/*16-bit-per-channel unsigned floating-point red channel data format with normalized value. Values are in the [code][0.0, 1.0][/code] range.*/
	DataFormatR16Unorm DataFormat = 69
	/*16-bit-per-channel signed floating-point red channel data format with normalized value. Values are in the [code][-1.0, 1.0][/code] range.*/
	DataFormatR16Snorm DataFormat = 70
	/*16-bit-per-channel unsigned floating-point red channel data format with scaled value (value is converted from integer to float). Values are in the [code][0.0, 65535.0][/code] range.*/
	DataFormatR16Uscaled DataFormat = 71
	/*16-bit-per-channel signed floating-point red channel data format with scaled value (value is converted from integer to float). Values are in the [code][-32767.0, 32767.0][/code] range.*/
	DataFormatR16Sscaled DataFormat = 72
	/*16-bit-per-channel unsigned integer red channel data format. Values are in the [code][0.0, 65535][/code] range.*/
	DataFormatR16Uint DataFormat = 73
	/*16-bit-per-channel signed integer red channel data format. Values are in the [code][-32767, 32767][/code] range.*/
	DataFormatR16Sint DataFormat = 74
	/*16-bit-per-channel signed floating-point red channel data format with the value stored as-is.*/
	DataFormatR16Sfloat DataFormat = 75
	/*16-bit-per-channel unsigned floating-point red/green channel data format with normalized value. Values are in the [code][0.0, 1.0][/code] range.*/
	DataFormatR16g16Unorm DataFormat = 76
	/*16-bit-per-channel signed floating-point red/green channel data format with normalized value. Values are in the [code][-1.0, 1.0][/code] range.*/
	DataFormatR16g16Snorm DataFormat = 77
	/*16-bit-per-channel unsigned floating-point red/green channel data format with scaled value (value is converted from integer to float). Values are in the [code][0.0, 65535.0][/code] range.*/
	DataFormatR16g16Uscaled DataFormat = 78
	/*16-bit-per-channel signed floating-point red/green channel data format with scaled value (value is converted from integer to float). Values are in the [code][-32767.0, 32767.0][/code] range.*/
	DataFormatR16g16Sscaled DataFormat = 79
	/*16-bit-per-channel unsigned integer red/green channel data format. Values are in the [code][0.0, 65535][/code] range.*/
	DataFormatR16g16Uint DataFormat = 80
	/*16-bit-per-channel signed integer red/green channel data format. Values are in the [code][-32767, 32767][/code] range.*/
	DataFormatR16g16Sint DataFormat = 81
	/*16-bit-per-channel signed floating-point red/green channel data format with the value stored as-is.*/
	DataFormatR16g16Sfloat DataFormat = 82
	/*16-bit-per-channel unsigned floating-point red/green/blue channel data format with normalized value. Values are in the [code][0.0, 1.0][/code] range.*/
	DataFormatR16g16b16Unorm DataFormat = 83
	/*16-bit-per-channel signed floating-point red/green/blue channel data format with normalized value. Values are in the [code][-1.0, 1.0][/code] range.*/
	DataFormatR16g16b16Snorm DataFormat = 84
	/*16-bit-per-channel unsigned floating-point red/green/blue channel data format with scaled value (value is converted from integer to float). Values are in the [code][0.0, 65535.0][/code] range.*/
	DataFormatR16g16b16Uscaled DataFormat = 85
	/*16-bit-per-channel signed floating-point red/green/blue channel data format with scaled value (value is converted from integer to float). Values are in the [code][-32767.0, 32767.0][/code] range.*/
	DataFormatR16g16b16Sscaled DataFormat = 86
	/*16-bit-per-channel unsigned integer red/green/blue channel data format. Values are in the [code][0.0, 65535][/code] range.*/
	DataFormatR16g16b16Uint DataFormat = 87
	/*16-bit-per-channel signed integer red/green/blue channel data format. Values are in the [code][-32767, 32767][/code] range.*/
	DataFormatR16g16b16Sint DataFormat = 88
	/*16-bit-per-channel signed floating-point red/green/blue channel data format with the value stored as-is.*/
	DataFormatR16g16b16Sfloat DataFormat = 89
	/*16-bit-per-channel unsigned floating-point red/green/blue/alpha channel data format with normalized value. Values are in the [code][0.0, 1.0][/code] range.*/
	DataFormatR16g16b16a16Unorm DataFormat = 90
	/*16-bit-per-channel signed floating-point red/green/blue/alpha channel data format with normalized value. Values are in the [code][-1.0, 1.0][/code] range.*/
	DataFormatR16g16b16a16Snorm DataFormat = 91
	/*16-bit-per-channel unsigned floating-point red/green/blue/alpha channel data format with scaled value (value is converted from integer to float). Values are in the [code][0.0, 65535.0][/code] range.*/
	DataFormatR16g16b16a16Uscaled DataFormat = 92
	/*16-bit-per-channel signed floating-point red/green/blue/alpha channel data format with scaled value (value is converted from integer to float). Values are in the [code][-32767.0, 32767.0][/code] range.*/
	DataFormatR16g16b16a16Sscaled DataFormat = 93
	/*16-bit-per-channel unsigned integer red/green/blue/alpha channel data format. Values are in the [code][0.0, 65535][/code] range.*/
	DataFormatR16g16b16a16Uint DataFormat = 94
	/*16-bit-per-channel signed integer red/green/blue/alpha channel data format. Values are in the [code][-32767, 32767][/code] range.*/
	DataFormatR16g16b16a16Sint DataFormat = 95
	/*16-bit-per-channel signed floating-point red/green/blue/alpha channel data format with the value stored as-is.*/
	DataFormatR16g16b16a16Sfloat DataFormat = 96
	/*32-bit-per-channel unsigned integer red channel data format. Values are in the [code][0, 2^32 - 1][/code] range.*/
	DataFormatR32Uint DataFormat = 97
	/*32-bit-per-channel signed integer red channel data format. Values are in the [code][2^31 + 1, 2^31 - 1][/code] range.*/
	DataFormatR32Sint DataFormat = 98
	/*32-bit-per-channel signed floating-point red channel data format with the value stored as-is.*/
	DataFormatR32Sfloat DataFormat = 99
	/*32-bit-per-channel unsigned integer red/green channel data format. Values are in the [code][0, 2^32 - 1][/code] range.*/
	DataFormatR32g32Uint DataFormat = 100
	/*32-bit-per-channel signed integer red/green channel data format. Values are in the [code][2^31 + 1, 2^31 - 1][/code] range.*/
	DataFormatR32g32Sint DataFormat = 101
	/*32-bit-per-channel signed floating-point red/green channel data format with the value stored as-is.*/
	DataFormatR32g32Sfloat DataFormat = 102
	/*32-bit-per-channel unsigned integer red/green/blue channel data format. Values are in the [code][0, 2^32 - 1][/code] range.*/
	DataFormatR32g32b32Uint DataFormat = 103
	/*32-bit-per-channel signed integer red/green/blue channel data format. Values are in the [code][2^31 + 1, 2^31 - 1][/code] range.*/
	DataFormatR32g32b32Sint DataFormat = 104
	/*32-bit-per-channel signed floating-point red/green/blue channel data format with the value stored as-is.*/
	DataFormatR32g32b32Sfloat DataFormat = 105
	/*32-bit-per-channel unsigned integer red/green/blue/alpha channel data format. Values are in the [code][0, 2^32 - 1][/code] range.*/
	DataFormatR32g32b32a32Uint DataFormat = 106
	/*32-bit-per-channel signed integer red/green/blue/alpha channel data format. Values are in the [code][2^31 + 1, 2^31 - 1][/code] range.*/
	DataFormatR32g32b32a32Sint DataFormat = 107
	/*32-bit-per-channel signed floating-point red/green/blue/alpha channel data format with the value stored as-is.*/
	DataFormatR32g32b32a32Sfloat DataFormat = 108
	/*64-bit-per-channel unsigned integer red channel data format. Values are in the [code][0, 2^64 - 1][/code] range.*/
	DataFormatR64Uint DataFormat = 109
	/*64-bit-per-channel signed integer red channel data format. Values are in the [code][2^63 + 1, 2^63 - 1][/code] range.*/
	DataFormatR64Sint DataFormat = 110
	/*64-bit-per-channel signed floating-point red channel data format with the value stored as-is.*/
	DataFormatR64Sfloat DataFormat = 111
	/*64-bit-per-channel unsigned integer red/green channel data format. Values are in the [code][0, 2^64 - 1][/code] range.*/
	DataFormatR64g64Uint DataFormat = 112
	/*64-bit-per-channel signed integer red/green channel data format. Values are in the [code][2^63 + 1, 2^63 - 1][/code] range.*/
	DataFormatR64g64Sint DataFormat = 113
	/*64-bit-per-channel signed floating-point red/green channel data format with the value stored as-is.*/
	DataFormatR64g64Sfloat DataFormat = 114
	/*64-bit-per-channel unsigned integer red/green/blue channel data format. Values are in the [code][0, 2^64 - 1][/code] range.*/
	DataFormatR64g64b64Uint DataFormat = 115
	/*64-bit-per-channel signed integer red/green/blue channel data format. Values are in the [code][2^63 + 1, 2^63 - 1][/code] range.*/
	DataFormatR64g64b64Sint DataFormat = 116
	/*64-bit-per-channel signed floating-point red/green/blue channel data format with the value stored as-is.*/
	DataFormatR64g64b64Sfloat DataFormat = 117
	/*64-bit-per-channel unsigned integer red/green/blue/alpha channel data format. Values are in the [code][0, 2^64 - 1][/code] range.*/
	DataFormatR64g64b64a64Uint DataFormat = 118
	/*64-bit-per-channel signed integer red/green/blue/alpha channel data format. Values are in the [code][2^63 + 1, 2^63 - 1][/code] range.*/
	DataFormatR64g64b64a64Sint DataFormat = 119
	/*64-bit-per-channel signed floating-point red/green/blue/alpha channel data format with the value stored as-is.*/
	DataFormatR64g64b64a64Sfloat DataFormat = 120
	/*Unsigned floating-point blue/green/red data format with the value stored as-is, packed in 32 bits. The format's precision is 10 bits of blue channel, 11 bits of green channel and 11 bits of red channel.*/
	DataFormatB10g11r11UfloatPack32 DataFormat = 121
	/*Unsigned floating-point exposure/blue/green/red data format with the value stored as-is, packed in 32 bits. The format's precision is 5 bits of exposure, 9 bits of blue channel, 9 bits of green channel and 9 bits of red channel.*/
	DataFormatE5b9g9r9UfloatPack32 DataFormat = 122
	/*16-bit unsigned floating-point depth data format with normalized value. Values are in the [code][0.0, 1.0][/code] range.*/
	DataFormatD16Unorm DataFormat = 123
	/*24-bit unsigned floating-point depth data format with normalized value, plus 8 unused bits, packed in 32 bits. Values for depth are in the [code][0.0, 1.0][/code] range.*/
	DataFormatX8D24UnormPack32 DataFormat = 124
	/*32-bit signed floating-point depth data format with the value stored as-is.*/
	DataFormatD32Sfloat DataFormat = 125
	/*8-bit unsigned integer stencil data format.*/
	DataFormatS8Uint DataFormat = 126
	/*16-bit unsigned floating-point depth data format with normalized value, plus 8 bits of stencil in unsigned integer format. Values for depth are in the [code][0.0, 1.0][/code] range. Values for stencil are in the [code][0, 255][/code] range.*/
	DataFormatD16UnormS8Uint DataFormat = 127
	/*24-bit unsigned floating-point depth data format with normalized value, plus 8 bits of stencil in unsigned integer format. Values for depth are in the [code][0.0, 1.0][/code] range. Values for stencil are in the [code][0, 255][/code] range.*/
	DataFormatD24UnormS8Uint DataFormat = 128
	/*32-bit signed floating-point depth data format with the value stored as-is, plus 8 bits of stencil in unsigned integer format. Values for stencil are in the [code][0, 255][/code] range.*/
	DataFormatD32SfloatS8Uint DataFormat = 129
	/*VRAM-compressed unsigned red/green/blue channel data format with normalized value. Values are in the [code][0.0, 1.0][/code] range. The format's precision is 5 bits of red channel, 6 bits of green channel and 5 bits of blue channel. Using BC1 texture compression (also known as S3TC DXT1).*/
	DataFormatBc1RgbUnormBlock DataFormat = 130
	/*VRAM-compressed unsigned red/green/blue channel data format with normalized value and non-linear sRGB encoding. Values are in the [code][0.0, 1.0][/code] range. The format's precision is 5 bits of red channel, 6 bits of green channel and 5 bits of blue channel. Using BC1 texture compression (also known as S3TC DXT1).*/
	DataFormatBc1RgbSrgbBlock DataFormat = 131
	/*VRAM-compressed unsigned red/green/blue/alpha channel data format with normalized value. Values are in the [code][0.0, 1.0][/code] range. The format's precision is 5 bits of red channel, 6 bits of green channel, 5 bits of blue channel and 1 bit of alpha channel. Using BC1 texture compression (also known as S3TC DXT1).*/
	DataFormatBc1RgbaUnormBlock DataFormat = 132
	/*VRAM-compressed unsigned red/green/blue/alpha channel data format with normalized value and non-linear sRGB encoding. Values are in the [code][0.0, 1.0][/code] range. The format's precision is 5 bits of red channel, 6 bits of green channel, 5 bits of blue channel and 1 bit of alpha channel. Using BC1 texture compression (also known as S3TC DXT1).*/
	DataFormatBc1RgbaSrgbBlock DataFormat = 133
	/*VRAM-compressed unsigned red/green/blue/alpha channel data format with normalized value. Values are in the [code][0.0, 1.0][/code] range. The format's precision is 5 bits of red channel, 6 bits of green channel, 5 bits of blue channel and 4 bits of alpha channel. Using BC2 texture compression (also known as S3TC DXT3).*/
	DataFormatBc2UnormBlock DataFormat = 134
	/*VRAM-compressed unsigned red/green/blue/alpha channel data format with normalized value and non-linear sRGB encoding. Values are in the [code][0.0, 1.0][/code] range. The format's precision is 5 bits of red channel, 6 bits of green channel, 5 bits of blue channel and 4 bits of alpha channel. Using BC2 texture compression (also known as S3TC DXT3).*/
	DataFormatBc2SrgbBlock DataFormat = 135
	/*VRAM-compressed unsigned red/green/blue/alpha channel data format with normalized value. Values are in the [code][0.0, 1.0][/code] range. The format's precision is 5 bits of red channel, 6 bits of green channel, 5 bits of blue channel and 8 bits of alpha channel. Using BC3 texture compression (also known as S3TC DXT5).*/
	DataFormatBc3UnormBlock DataFormat = 136
	/*VRAM-compressed unsigned red/green/blue/alpha channel data format with normalized value and non-linear sRGB encoding. Values are in the [code][0.0, 1.0][/code] range. The format's precision is 5 bits of red channel, 6 bits of green channel, 5 bits of blue channel and 8 bits of alpha channel. Using BC3 texture compression (also known as S3TC DXT5).*/
	DataFormatBc3SrgbBlock DataFormat = 137
	/*VRAM-compressed unsigned red channel data format with normalized value. Values are in the [code][0.0, 1.0][/code] range. The format's precision is 8 bits of red channel. Using BC4 texture compression.*/
	DataFormatBc4UnormBlock DataFormat = 138
	/*VRAM-compressed signed red channel data format with normalized value. Values are in the [code][-1.0, 1.0][/code] range. The format's precision is 8 bits of red channel. Using BC4 texture compression.*/
	DataFormatBc4SnormBlock DataFormat = 139
	/*VRAM-compressed unsigned red/green channel data format with normalized value. Values are in the [code][0.0, 1.0][/code] range. The format's precision is 8 bits of red channel and 8 bits of green channel. Using BC5 texture compression (also known as S3TC RGTC).*/
	DataFormatBc5UnormBlock DataFormat = 140
	/*VRAM-compressed signed red/green channel data format with normalized value. Values are in the [code][-1.0, 1.0][/code] range. The format's precision is 8 bits of red channel and 8 bits of green channel. Using BC5 texture compression (also known as S3TC RGTC).*/
	DataFormatBc5SnormBlock DataFormat = 141
	/*VRAM-compressed unsigned red/green/blue channel data format with the floating-point value stored as-is. The format's precision is between 10 and 13 bits for the red/green/blue channels. Using BC6H texture compression (also known as BPTC HDR).*/
	DataFormatBc6hUfloatBlock DataFormat = 142
	/*VRAM-compressed signed red/green/blue channel data format with the floating-point value stored as-is. The format's precision is between 10 and 13 bits for the red/green/blue channels. Using BC6H texture compression (also known as BPTC HDR).*/
	DataFormatBc6hSfloatBlock DataFormat = 143
	/*VRAM-compressed unsigned red/green/blue/alpha channel data format with normalized value. Values are in the [code][0.0, 1.0][/code] range. The format's precision is between 4 and 7 bits for the red/green/blue channels and between 0 and 8 bits for the alpha channel. Also known as BPTC LDR.*/
	DataFormatBc7UnormBlock DataFormat = 144
	/*VRAM-compressed unsigned red/green/blue/alpha channel data format with normalized value and non-linear sRGB encoding. Values are in the [code][0.0, 1.0][/code] range. The format's precision is between 4 and 7 bits for the red/green/blue channels and between 0 and 8 bits for the alpha channel. Also known as BPTC LDR.*/
	DataFormatBc7SrgbBlock DataFormat = 145
	/*VRAM-compressed unsigned red/green/blue channel data format with normalized value. Values are in the [code][0.0, 1.0][/code] range. Using ETC2 texture compression.*/
	DataFormatEtc2R8g8b8UnormBlock DataFormat = 146
	/*VRAM-compressed unsigned red/green/blue channel data format with normalized value and non-linear sRGB encoding. Values are in the [code][0.0, 1.0][/code] range. Using ETC2 texture compression.*/
	DataFormatEtc2R8g8b8SrgbBlock DataFormat = 147
	/*VRAM-compressed unsigned red/green/blue/alpha channel data format with normalized value. Values are in the [code][0.0, 1.0][/code] range. Red/green/blue use 8 bit of precision each, with alpha using 1 bit of precision. Using ETC2 texture compression.*/
	DataFormatEtc2R8g8b8a1UnormBlock DataFormat = 148
	/*VRAM-compressed unsigned red/green/blue/alpha channel data format with normalized value and non-linear sRGB encoding. Values are in the [code][0.0, 1.0][/code] range. Red/green/blue use 8 bit of precision each, with alpha using 1 bit of precision. Using ETC2 texture compression.*/
	DataFormatEtc2R8g8b8a1SrgbBlock DataFormat = 149
	/*VRAM-compressed unsigned red/green/blue/alpha channel data format with normalized value. Values are in the [code][0.0, 1.0][/code] range. Red/green/blue use 8 bits of precision each, with alpha using 8 bits of precision. Using ETC2 texture compression.*/
	DataFormatEtc2R8g8b8a8UnormBlock DataFormat = 150
	/*VRAM-compressed unsigned red/green/blue/alpha channel data format with normalized value and non-linear sRGB encoding. Values are in the [code][0.0, 1.0][/code] range. Red/green/blue use 8 bits of precision each, with alpha using 8 bits of precision. Using ETC2 texture compression.*/
	DataFormatEtc2R8g8b8a8SrgbBlock DataFormat = 151
	/*11-bit VRAM-compressed unsigned red channel data format with normalized value. Values are in the [code][0.0, 1.0][/code] range. Using ETC2 texture compression.*/
	DataFormatEacR11UnormBlock DataFormat = 152
	/*11-bit VRAM-compressed signed red channel data format with normalized value. Values are in the [code][-1.0, 1.0][/code] range. Using ETC2 texture compression.*/
	DataFormatEacR11SnormBlock DataFormat = 153
	/*11-bit VRAM-compressed unsigned red/green channel data format with normalized value. Values are in the [code][0.0, 1.0][/code] range. Using ETC2 texture compression.*/
	DataFormatEacR11g11UnormBlock DataFormat = 154
	/*11-bit VRAM-compressed signed red/green channel data format with normalized value. Values are in the [code][-1.0, 1.0][/code] range. Using ETC2 texture compression.*/
	DataFormatEacR11g11SnormBlock DataFormat = 155
	/*VRAM-compressed unsigned floating-point data format with normalized value, packed in 4×4 blocks (highest quality). Values are in the [code][0.0, 1.0][/code] range. Using ASTC compression.*/
	DataFormatAstc4x4UnormBlock DataFormat = 156
	/*VRAM-compressed unsigned floating-point data format with normalized value and non-linear sRGB encoding, packed in 4×4 blocks (highest quality). Values are in the [code][0.0, 1.0][/code] range. Using ASTC compression.*/
	DataFormatAstc4x4SrgbBlock DataFormat = 157
	/*VRAM-compressed unsigned floating-point data format with normalized value, packed in 5×4 blocks. Values are in the [code][0.0, 1.0][/code] range. Using ASTC compression.*/
	DataFormatAstc5x4UnormBlock DataFormat = 158
	/*VRAM-compressed unsigned floating-point data format with normalized value and non-linear sRGB encoding, packed in 5×4 blocks. Values are in the [code][0.0, 1.0][/code] range. Using ASTC compression.*/
	DataFormatAstc5x4SrgbBlock DataFormat = 159
	/*VRAM-compressed unsigned floating-point data format with normalized value, packed in 5×5 blocks. Values are in the [code][0.0, 1.0][/code] range. Using ASTC compression.*/
	DataFormatAstc5x5UnormBlock DataFormat = 160
	/*VRAM-compressed unsigned floating-point data format with normalized value and non-linear sRGB encoding, packed in 5×5 blocks. Values are in the [code][0.0, 1.0][/code] range. Using ASTC compression.*/
	DataFormatAstc5x5SrgbBlock DataFormat = 161
	/*VRAM-compressed unsigned floating-point data format with normalized value, packed in 6×5 blocks. Values are in the [code][0.0, 1.0][/code] range. Using ASTC compression.*/
	DataFormatAstc6x5UnormBlock DataFormat = 162
	/*VRAM-compressed unsigned floating-point data format with normalized value and non-linear sRGB encoding, packed in 6×5 blocks. Values are in the [code][0.0, 1.0][/code] range. Using ASTC compression.*/
	DataFormatAstc6x5SrgbBlock DataFormat = 163
	/*VRAM-compressed unsigned floating-point data format with normalized value, packed in 6×6 blocks. Values are in the [code][0.0, 1.0][/code] range. Using ASTC compression.*/
	DataFormatAstc6x6UnormBlock DataFormat = 164
	/*VRAM-compressed unsigned floating-point data format with normalized value and non-linear sRGB encoding, packed in 6×6 blocks. Values are in the [code][0.0, 1.0][/code] range. Using ASTC compression.*/
	DataFormatAstc6x6SrgbBlock DataFormat = 165
	/*VRAM-compressed unsigned floating-point data format with normalized value, packed in 8×5 blocks. Values are in the [code][0.0, 1.0][/code] range. Using ASTC compression.*/
	DataFormatAstc8x5UnormBlock DataFormat = 166
	/*VRAM-compressed unsigned floating-point data format with normalized value and non-linear sRGB encoding, packed in 8×5 blocks. Values are in the [code][0.0, 1.0][/code] range. Using ASTC compression.*/
	DataFormatAstc8x5SrgbBlock DataFormat = 167
	/*VRAM-compressed unsigned floating-point data format with normalized value, packed in 8×6 blocks. Values are in the [code][0.0, 1.0][/code] range. Using ASTC compression.*/
	DataFormatAstc8x6UnormBlock DataFormat = 168
	/*VRAM-compressed unsigned floating-point data format with normalized value and non-linear sRGB encoding, packed in 8×6 blocks. Values are in the [code][0.0, 1.0][/code] range. Using ASTC compression.*/
	DataFormatAstc8x6SrgbBlock DataFormat = 169
	/*VRAM-compressed unsigned floating-point data format with normalized value, packed in 8×8 blocks. Values are in the [code][0.0, 1.0][/code] range. Using ASTC compression.*/
	DataFormatAstc8x8UnormBlock DataFormat = 170
	/*VRAM-compressed unsigned floating-point data format with normalized value and non-linear sRGB encoding, packed in 8×8 blocks. Values are in the [code][0.0, 1.0][/code] range. Using ASTC compression.*/
	DataFormatAstc8x8SrgbBlock DataFormat = 171
	/*VRAM-compressed unsigned floating-point data format with normalized value, packed in 10×5 blocks. Values are in the [code][0.0, 1.0][/code] range. Using ASTC compression.*/
	DataFormatAstc10x5UnormBlock DataFormat = 172
	/*VRAM-compressed unsigned floating-point data format with normalized value and non-linear sRGB encoding, packed in 10×5 blocks. Values are in the [code][0.0, 1.0][/code] range. Using ASTC compression.*/
	DataFormatAstc10x5SrgbBlock DataFormat = 173
	/*VRAM-compressed unsigned floating-point data format with normalized value, packed in 10×6 blocks. Values are in the [code][0.0, 1.0][/code] range. Using ASTC compression.*/
	DataFormatAstc10x6UnormBlock DataFormat = 174
	/*VRAM-compressed unsigned floating-point data format with normalized value and non-linear sRGB encoding, packed in 10×6 blocks. Values are in the [code][0.0, 1.0][/code] range. Using ASTC compression.*/
	DataFormatAstc10x6SrgbBlock DataFormat = 175
	/*VRAM-compressed unsigned floating-point data format with normalized value, packed in 10×8 blocks. Values are in the [code][0.0, 1.0][/code] range. Using ASTC compression.*/
	DataFormatAstc10x8UnormBlock DataFormat = 176
	/*VRAM-compressed unsigned floating-point data format with normalized value and non-linear sRGB encoding, packed in 10×8 blocks. Values are in the [code][0.0, 1.0][/code] range. Using ASTC compression.*/
	DataFormatAstc10x8SrgbBlock DataFormat = 177
	/*VRAM-compressed unsigned floating-point data format with normalized value, packed in 10×10 blocks. Values are in the [code][0.0, 1.0][/code] range. Using ASTC compression.*/
	DataFormatAstc10x10UnormBlock DataFormat = 178
	/*VRAM-compressed unsigned floating-point data format with normalized value and non-linear sRGB encoding, packed in 10×10 blocks. Values are in the [code][0.0, 1.0][/code] range. Using ASTC compression.*/
	DataFormatAstc10x10SrgbBlock DataFormat = 179
	/*VRAM-compressed unsigned floating-point data format with normalized value, packed in 12×10 blocks. Values are in the [code][0.0, 1.0][/code] range. Using ASTC compression.*/
	DataFormatAstc12x10UnormBlock DataFormat = 180
	/*VRAM-compressed unsigned floating-point data format with normalized value and non-linear sRGB encoding, packed in 12×10 blocks. Values are in the [code][0.0, 1.0][/code] range. Using ASTC compression.*/
	DataFormatAstc12x10SrgbBlock DataFormat = 181
	/*VRAM-compressed unsigned floating-point data format with normalized value, packed in 12 blocks (lowest quality). Values are in the [code][0.0, 1.0][/code] range. Using ASTC compression.*/
	DataFormatAstc12x12UnormBlock DataFormat = 182
	/*VRAM-compressed unsigned floating-point data format with normalized value and non-linear sRGB encoding, packed in 12 blocks (lowest quality). Values are in the [code][0.0, 1.0][/code] range. Using ASTC compression.*/
	DataFormatAstc12x12SrgbBlock DataFormat = 183
	/*8-bit-per-channel unsigned floating-point green/blue/red channel data format with normalized value. Values are in the [code][0.0, 1.0][/code] range. Blue and red channel data is stored at halved horizontal resolution (i.e. 2 horizontally adjacent pixels will share the same value for the blue/red channel).*/
	DataFormatG8b8g8r8422Unorm DataFormat = 184
	/*8-bit-per-channel unsigned floating-point blue/green/red channel data format with normalized value. Values are in the [code][0.0, 1.0][/code] range. Blue and red channel data is stored at halved horizontal resolution (i.e. 2 horizontally adjacent pixels will share the same value for the blue/red channel).*/
	DataFormatB8g8r8g8422Unorm DataFormat = 185
	/*8-bit-per-channel unsigned floating-point green/blue/red channel data with normalized value, stored across 3 separate planes (green + blue + red). Values are in the [code][0.0, 1.0][/code] range. Blue and red channel data is stored at halved horizontal and vertical resolution (i.e. 2×2 adjacent pixels will share the same value for the blue/red channel).*/
	DataFormatG8B8R83plane420Unorm DataFormat = 186
	/*8-bit-per-channel unsigned floating-point green/blue/red channel data with normalized value, stored across 2 separate planes (green + blue/red). Values are in the [code][0.0, 1.0][/code] range. Blue and red channel data is stored at halved horizontal and vertical resolution (i.e. 2×2 adjacent pixels will share the same value for the blue/red channel).*/
	DataFormatG8B8r82plane420Unorm DataFormat = 187
	/*8-bit-per-channel unsigned floating-point green/blue/red channel data with normalized value, stored across 2 separate planes (green + blue + red). Values are in the [code][0.0, 1.0][/code] range. Blue and red channel data is stored at halved horizontal resolution (i.e. 2 horizontally adjacent pixels will share the same value for the blue/red channel).*/
	DataFormatG8B8R83plane422Unorm DataFormat = 188
	/*8-bit-per-channel unsigned floating-point green/blue/red channel data with normalized value, stored across 2 separate planes (green + blue/red). Values are in the [code][0.0, 1.0][/code] range. Blue and red channel data is stored at halved horizontal resolution (i.e. 2 horizontally adjacent pixels will share the same value for the blue/red channel).*/
	DataFormatG8B8r82plane422Unorm DataFormat = 189
	/*8-bit-per-channel unsigned floating-point green/blue/red channel data with normalized value, stored across 3 separate planes. Values are in the [code][0.0, 1.0][/code] range.*/
	DataFormatG8B8R83plane444Unorm DataFormat = 190
	/*10-bit-per-channel unsigned floating-point red channel data with normalized value, plus 6 unused bits, packed in 16 bits. Values are in the [code][0.0, 1.0][/code] range.*/
	DataFormatR10x6UnormPack16 DataFormat = 191
	/*10-bit-per-channel unsigned floating-point red/green channel data with normalized value, plus 6 unused bits after each channel, packed in 2×16 bits. Values are in the [code][0.0, 1.0][/code] range.*/
	DataFormatR10x6g10x6Unorm2pack16 DataFormat = 192
	/*10-bit-per-channel unsigned floating-point red/green/blue/alpha channel data with normalized value, plus 6 unused bits after each channel, packed in 4×16 bits. Values are in the [code][0.0, 1.0][/code] range.*/
	DataFormatR10x6g10x6b10x6a10x6Unorm4pack16 DataFormat = 193
	/*10-bit-per-channel unsigned floating-point green/blue/green/red channel data with normalized value, plus 6 unused bits after each channel, packed in 4×16 bits. Values are in the [code][0.0, 1.0][/code] range. Blue and red channel data is stored at halved horizontal resolution (i.e. 2 horizontally adjacent pixels will share the same value for the blue/red channel). The green channel is listed twice, but contains different values to allow it to be represented at full resolution.*/
	DataFormatG10x6b10x6g10x6r10x6422Unorm4pack16 DataFormat = 194
	/*10-bit-per-channel unsigned floating-point blue/green/red/green channel data with normalized value, plus 6 unused bits after each channel, packed in 4×16 bits. Values are in the [code][0.0, 1.0][/code] range. Blue and red channel data is stored at halved horizontal resolution (i.e. 2 horizontally adjacent pixels will share the same value for the blue/red channel). The green channel is listed twice, but contains different values to allow it to be represented at full resolution.*/
	DataFormatB10x6g10x6r10x6g10x6422Unorm4pack16 DataFormat = 195
	/*10-bit-per-channel unsigned floating-point green/blue/red channel data with normalized value, plus 6 unused bits after each channel. Packed in 3×16 bits and stored across 2 separate planes (green + blue + red). Values are in the [code][0.0, 1.0][/code] range. Blue and red channel data is stored at halved horizontal and vertical resolution (i.e. 2×2 adjacent pixels will share the same value for the blue/red channel).*/
	DataFormatG10x6B10x6R10x63plane420Unorm3pack16 DataFormat = 196
	/*10-bit-per-channel unsigned floating-point green/blue/red channel data with normalized value, plus 6 unused bits after each channel. Packed in 3×16 bits and stored across 2 separate planes (green + blue/red). Values are in the [code][0.0, 1.0][/code] range. Blue and red channel data is stored at halved horizontal and vertical resolution (i.e. 2×2 adjacent pixels will share the same value for the blue/red channel).*/
	DataFormatG10x6B10x6r10x62plane420Unorm3pack16 DataFormat = 197
	/*10-bit-per-channel unsigned floating-point green/blue/red channel data with normalized value, plus 6 unused bits after each channel. Packed in 3×16 bits and stored across 3 separate planes (green + blue + red). Values are in the [code][0.0, 1.0][/code] range. Blue and red channel data is stored at halved horizontal resolution (i.e. 2 horizontally adjacent pixels will share the same value for the blue/red channel).*/
	DataFormatG10x6B10x6R10x63plane422Unorm3pack16 DataFormat = 198
	/*10-bit-per-channel unsigned floating-point green/blue/red channel data with normalized value, plus 6 unused bits after each channel. Packed in 3×16 bits and stored across 3 separate planes (green + blue/red). Values are in the [code][0.0, 1.0][/code] range. Blue and red channel data is stored at halved horizontal resolution (i.e. 2 horizontally adjacent pixels will share the same value for the blue/red channel).*/
	DataFormatG10x6B10x6r10x62plane422Unorm3pack16 DataFormat = 199
	/*10-bit-per-channel unsigned floating-point green/blue/red channel data with normalized value, plus 6 unused bits after each channel. Packed in 3×16 bits and stored across 3 separate planes (green + blue + red). Values are in the [code][0.0, 1.0][/code] range.*/
	DataFormatG10x6B10x6R10x63plane444Unorm3pack16 DataFormat = 200
	/*12-bit-per-channel unsigned floating-point red channel data with normalized value, plus 6 unused bits, packed in 16 bits. Values are in the [code][0.0, 1.0][/code] range.*/
	DataFormatR12x4UnormPack16 DataFormat = 201
	/*12-bit-per-channel unsigned floating-point red/green channel data with normalized value, plus 6 unused bits after each channel, packed in 2×16 bits. Values are in the [code][0.0, 1.0][/code] range.*/
	DataFormatR12x4g12x4Unorm2pack16 DataFormat = 202
	/*12-bit-per-channel unsigned floating-point red/green/blue/alpha channel data with normalized value, plus 6 unused bits after each channel, packed in 4×16 bits. Values are in the [code][0.0, 1.0][/code] range.*/
	DataFormatR12x4g12x4b12x4a12x4Unorm4pack16 DataFormat = 203
	/*12-bit-per-channel unsigned floating-point green/blue/green/red channel data with normalized value, plus 6 unused bits after each channel, packed in 4×16 bits. Values are in the [code][0.0, 1.0][/code] range. Blue and red channel data is stored at halved horizontal resolution (i.e. 2 horizontally adjacent pixels will share the same value for the blue/red channel). The green channel is listed twice, but contains different values to allow it to be represented at full resolution.*/
	DataFormatG12x4b12x4g12x4r12x4422Unorm4pack16 DataFormat = 204
	/*12-bit-per-channel unsigned floating-point blue/green/red/green channel data with normalized value, plus 6 unused bits after each channel, packed in 4×16 bits. Values are in the [code][0.0, 1.0][/code] range. Blue and red channel data is stored at halved horizontal resolution (i.e. 2 horizontally adjacent pixels will share the same value for the blue/red channel). The green channel is listed twice, but contains different values to allow it to be represented at full resolution.*/
	DataFormatB12x4g12x4r12x4g12x4422Unorm4pack16 DataFormat = 205
	/*12-bit-per-channel unsigned floating-point green/blue/red channel data with normalized value, plus 6 unused bits after each channel. Packed in 3×16 bits and stored across 2 separate planes (green + blue + red). Values are in the [code][0.0, 1.0][/code] range. Blue and red channel data is stored at halved horizontal and vertical resolution (i.e. 2×2 adjacent pixels will share the same value for the blue/red channel).*/
	DataFormatG12x4B12x4R12x43plane420Unorm3pack16 DataFormat = 206
	/*12-bit-per-channel unsigned floating-point green/blue/red channel data with normalized value, plus 6 unused bits after each channel. Packed in 3×16 bits and stored across 2 separate planes (green + blue/red). Values are in the [code][0.0, 1.0][/code] range. Blue and red channel data is stored at halved horizontal and vertical resolution (i.e. 2×2 adjacent pixels will share the same value for the blue/red channel).*/
	DataFormatG12x4B12x4r12x42plane420Unorm3pack16 DataFormat = 207
	/*12-bit-per-channel unsigned floating-point green/blue/red channel data with normalized value, plus 6 unused bits after each channel. Packed in 3×16 bits and stored across 3 separate planes (green + blue + red). Values are in the [code][0.0, 1.0][/code] range. Blue and red channel data is stored at halved horizontal resolution (i.e. 2 horizontally adjacent pixels will share the same value for the blue/red channel).*/
	DataFormatG12x4B12x4R12x43plane422Unorm3pack16 DataFormat = 208
	/*12-bit-per-channel unsigned floating-point green/blue/red channel data with normalized value, plus 6 unused bits after each channel. Packed in 3×16 bits and stored across 3 separate planes (green + blue/red). Values are in the [code][0.0, 1.0][/code] range. Blue and red channel data is stored at halved horizontal resolution (i.e. 2 horizontally adjacent pixels will share the same value for the blue/red channel).*/
	DataFormatG12x4B12x4r12x42plane422Unorm3pack16 DataFormat = 209
	/*12-bit-per-channel unsigned floating-point green/blue/red channel data with normalized value, plus 6 unused bits after each channel. Packed in 3×16 bits and stored across 3 separate planes (green + blue + red). Values are in the [code][0.0, 1.0][/code] range.*/
	DataFormatG12x4B12x4R12x43plane444Unorm3pack16 DataFormat = 210
	/*16-bit-per-channel unsigned floating-point green/blue/red channel data format with normalized value. Values are in the [code][0.0, 1.0][/code] range. Blue and red channel data is stored at halved horizontal resolution (i.e. 2 horizontally adjacent pixels will share the same value for the blue/red channel).*/
	DataFormatG16b16g16r16422Unorm DataFormat = 211
	/*16-bit-per-channel unsigned floating-point blue/green/red channel data format with normalized value. Values are in the [code][0.0, 1.0][/code] range. Blue and red channel data is stored at halved horizontal resolution (i.e. 2 horizontally adjacent pixels will share the same value for the blue/red channel).*/
	DataFormatB16g16r16g16422Unorm DataFormat = 212
	/*16-bit-per-channel unsigned floating-point green/blue/red channel data with normalized value, plus 6 unused bits after each channel. Stored across 2 separate planes (green + blue + red). Values are in the [code][0.0, 1.0][/code] range. Blue and red channel data is stored at halved horizontal and vertical resolution (i.e. 2×2 adjacent pixels will share the same value for the blue/red channel).*/
	DataFormatG16B16R163plane420Unorm DataFormat = 213
	/*16-bit-per-channel unsigned floating-point green/blue/red channel data with normalized value, plus 6 unused bits after each channel. Stored across 2 separate planes (green + blue/red). Values are in the [code][0.0, 1.0][/code] range. Blue and red channel data is stored at halved horizontal and vertical resolution (i.e. 2×2 adjacent pixels will share the same value for the blue/red channel).*/
	DataFormatG16B16r162plane420Unorm DataFormat = 214
	/*16-bit-per-channel unsigned floating-point green/blue/red channel data with normalized value, plus 6 unused bits after each channel. Stored across 3 separate planes (green + blue + red). Values are in the [code][0.0, 1.0][/code] range. Blue and red channel data is stored at halved horizontal resolution (i.e. 2 horizontally adjacent pixels will share the same value for the blue/red channel).*/
	DataFormatG16B16R163plane422Unorm DataFormat = 215
	/*16-bit-per-channel unsigned floating-point green/blue/red channel data with normalized value, plus 6 unused bits after each channel. Stored across 3 separate planes (green + blue/red). Values are in the [code][0.0, 1.0][/code] range. Blue and red channel data is stored at halved horizontal resolution (i.e. 2 horizontally adjacent pixels will share the same value for the blue/red channel).*/
	DataFormatG16B16r162plane422Unorm DataFormat = 216
	/*16-bit-per-channel unsigned floating-point green/blue/red channel data with normalized value, plus 6 unused bits after each channel. Stored across 3 separate planes (green + blue + red). Values are in the [code][0.0, 1.0][/code] range.*/
	DataFormatG16B16R163plane444Unorm DataFormat = 217
	/*Represents the size of the [enum DataFormat] enum.*/
	DataFormatMax DataFormat = 218
)

type BarrierMask = gdclass.RenderingDeviceBarrierMask //gd:RenderingDevice.BarrierMask

const (
	/*Vertex shader barrier mask.*/
	BarrierMaskVertex BarrierMask = 1
	/*Fragment shader barrier mask.*/
	BarrierMaskFragment BarrierMask = 8
	/*Compute barrier mask.*/
	BarrierMaskCompute BarrierMask = 2
	/*Transfer barrier mask.*/
	BarrierMaskTransfer BarrierMask = 4
	/*Raster barrier mask (vertex and fragment). Equivalent to [code]BARRIER_MASK_VERTEX | BARRIER_MASK_FRAGMENT[/code].*/
	BarrierMaskRaster BarrierMask = 9
	/*Barrier mask for all types (vertex, fragment, compute, transfer).*/
	BarrierMaskAllBarriers BarrierMask = 32767
	/*No barrier for any type.*/
	BarrierMaskNoBarrier BarrierMask = 32768
)

type TextureType = gdclass.RenderingDeviceTextureType //gd:RenderingDevice.TextureType

const (
	/*1-dimensional texture.*/
	TextureType1d TextureType = 0
	/*2-dimensional texture.*/
	TextureType2d TextureType = 1
	/*3-dimensional texture.*/
	TextureType3d TextureType = 2
	/*[Cubemap] texture.*/
	TextureTypeCube TextureType = 3
	/*Array of 1-dimensional textures.*/
	TextureType1dArray TextureType = 4
	/*Array of 2-dimensional textures.*/
	TextureType2dArray TextureType = 5
	/*Array of [Cubemap] textures.*/
	TextureTypeCubeArray TextureType = 6
	/*Represents the size of the [enum TextureType] enum.*/
	TextureTypeMax TextureType = 7
)

type TextureSamples = gdclass.RenderingDeviceTextureSamples //gd:RenderingDevice.TextureSamples

const (
	/*Perform 1 texture sample (this is the fastest but lowest-quality for antialiasing).*/
	TextureSamples1 TextureSamples = 0
	/*Perform 2 texture samples.*/
	TextureSamples2 TextureSamples = 1
	/*Perform 4 texture samples.*/
	TextureSamples4 TextureSamples = 2
	/*Perform 8 texture samples. Not supported on mobile GPUs (including Apple Silicon).*/
	TextureSamples8 TextureSamples = 3
	/*Perform 16 texture samples. Not supported on mobile GPUs and many desktop GPUs.*/
	TextureSamples16 TextureSamples = 4
	/*Perform 32 texture samples. Not supported on most GPUs.*/
	TextureSamples32 TextureSamples = 5
	/*Perform 64 texture samples (this is the slowest but highest-quality for antialiasing). Not supported on most GPUs.*/
	TextureSamples64 TextureSamples = 6
	/*Represents the size of the [enum TextureSamples] enum.*/
	TextureSamplesMax TextureSamples = 7
)

type TextureUsageBits = gdclass.RenderingDeviceTextureUsageBits //gd:RenderingDevice.TextureUsageBits

const (
	/*Texture can be sampled.*/
	TextureUsageSamplingBit TextureUsageBits = 1
	/*Texture can be used as a color attachment in a framebuffer.*/
	TextureUsageColorAttachmentBit TextureUsageBits = 2
	/*Texture can be used as a depth/stencil attachment in a framebuffer.*/
	TextureUsageDepthStencilAttachmentBit TextureUsageBits = 4
	/*Texture can be used as a [url=https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#descriptorsets-storageimage]storage image[/url].*/
	TextureUsageStorageBit TextureUsageBits = 8
	/*Texture can be used as a [url=https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#descriptorsets-storageimage]storage image[/url] with support for atomic operations.*/
	TextureUsageStorageAtomicBit TextureUsageBits = 16
	/*Texture can be read back on the CPU using [method texture_get_data] faster than without this bit, since it is always kept in the system memory.*/
	TextureUsageCpuReadBit TextureUsageBits = 32
	/*Texture can be updated using [method texture_update].*/
	TextureUsageCanUpdateBit TextureUsageBits = 64
	/*Texture can be a source for [method texture_copy].*/
	TextureUsageCanCopyFromBit TextureUsageBits = 128
	/*Texture can be a destination for [method texture_copy].*/
	TextureUsageCanCopyToBit TextureUsageBits = 256
	/*Texture can be used as a [url=https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#descriptorsets-inputattachment]input attachment[/url] in a framebuffer.*/
	TextureUsageInputAttachmentBit TextureUsageBits = 512
)

type TextureSwizzle = gdclass.RenderingDeviceTextureSwizzle //gd:RenderingDevice.TextureSwizzle

const (
	/*Return the sampled value as-is.*/
	TextureSwizzleIdentity TextureSwizzle = 0
	/*Always return [code]0.0[/code] when sampling.*/
	TextureSwizzleZero TextureSwizzle = 1
	/*Always return [code]1.0[/code] when sampling.*/
	TextureSwizzleOne TextureSwizzle = 2
	/*Sample the red color channel.*/
	TextureSwizzleR TextureSwizzle = 3
	/*Sample the green color channel.*/
	TextureSwizzleG TextureSwizzle = 4
	/*Sample the blue color channel.*/
	TextureSwizzleB TextureSwizzle = 5
	/*Sample the alpha channel.*/
	TextureSwizzleA TextureSwizzle = 6
	/*Represents the size of the [enum TextureSwizzle] enum.*/
	TextureSwizzleMax TextureSwizzle = 7
)

type TextureSliceType = gdclass.RenderingDeviceTextureSliceType //gd:RenderingDevice.TextureSliceType

const (
	/*2-dimensional texture slice.*/
	TextureSlice2d TextureSliceType = 0
	/*Cubemap texture slice.*/
	TextureSliceCubemap TextureSliceType = 1
	/*3-dimensional texture slice.*/
	TextureSlice3d TextureSliceType = 2
)

type SamplerFilter = gdclass.RenderingDeviceSamplerFilter //gd:RenderingDevice.SamplerFilter

const (
	/*Nearest-neighbor sampler filtering. Sampling at higher resolutions than the source will result in a pixelated look.*/
	SamplerFilterNearest SamplerFilter = 0
	/*Bilinear sampler filtering. Sampling at higher resolutions than the source will result in a blurry look.*/
	SamplerFilterLinear SamplerFilter = 1
)

type SamplerRepeatMode = gdclass.RenderingDeviceSamplerRepeatMode //gd:RenderingDevice.SamplerRepeatMode

const (
	/*Sample with repeating enabled.*/
	SamplerRepeatModeRepeat SamplerRepeatMode = 0
	/*Sample with mirrored repeating enabled. When sampling outside the [code][0.0, 1.0][/code] range, return a mirrored version of the sampler. This mirrored version is mirrored again if sampling further away, with the pattern repeating indefinitely.*/
	SamplerRepeatModeMirroredRepeat SamplerRepeatMode = 1
	/*Sample with repeating disabled. When sampling outside the [code][0.0, 1.0][/code] range, return the color of the last pixel on the edge.*/
	SamplerRepeatModeClampToEdge SamplerRepeatMode = 2
	/*Sample with repeating disabled. When sampling outside the [code][0.0, 1.0][/code] range, return the specified [member RDSamplerState.border_color].*/
	SamplerRepeatModeClampToBorder SamplerRepeatMode = 3
	/*Sample with mirrored repeating enabled, but only once. When sampling in the [code][-1.0, 0.0][/code] range, return a mirrored version of the sampler. When sampling outside the [code][-1.0, 1.0][/code] range, return the color of the last pixel on the edge.*/
	SamplerRepeatModeMirrorClampToEdge SamplerRepeatMode = 4
	/*Represents the size of the [enum SamplerRepeatMode] enum.*/
	SamplerRepeatModeMax SamplerRepeatMode = 5
)

type SamplerBorderColor = gdclass.RenderingDeviceSamplerBorderColor //gd:RenderingDevice.SamplerBorderColor

const (
	/*Return a floating-point transparent black color when sampling outside the [code][0.0, 1.0][/code] range. Only effective if the sampler repeat mode is [constant SAMPLER_REPEAT_MODE_CLAMP_TO_BORDER].*/
	SamplerBorderColorFloatTransparentBlack SamplerBorderColor = 0
	/*Return a integer transparent black color when sampling outside the [code][0.0, 1.0][/code] range. Only effective if the sampler repeat mode is [constant SAMPLER_REPEAT_MODE_CLAMP_TO_BORDER].*/
	SamplerBorderColorIntTransparentBlack SamplerBorderColor = 1
	/*Return a floating-point opaque black color when sampling outside the [code][0.0, 1.0][/code] range. Only effective if the sampler repeat mode is [constant SAMPLER_REPEAT_MODE_CLAMP_TO_BORDER].*/
	SamplerBorderColorFloatOpaqueBlack SamplerBorderColor = 2
	/*Return a integer opaque black color when sampling outside the [code][0.0, 1.0][/code] range. Only effective if the sampler repeat mode is [constant SAMPLER_REPEAT_MODE_CLAMP_TO_BORDER].*/
	SamplerBorderColorIntOpaqueBlack SamplerBorderColor = 3
	/*Return a floating-point opaque white color when sampling outside the [code][0.0, 1.0][/code] range. Only effective if the sampler repeat mode is [constant SAMPLER_REPEAT_MODE_CLAMP_TO_BORDER].*/
	SamplerBorderColorFloatOpaqueWhite SamplerBorderColor = 4
	/*Return a integer opaque white color when sampling outside the [code][0.0, 1.0][/code] range. Only effective if the sampler repeat mode is [constant SAMPLER_REPEAT_MODE_CLAMP_TO_BORDER].*/
	SamplerBorderColorIntOpaqueWhite SamplerBorderColor = 5
	/*Represents the size of the [enum SamplerBorderColor] enum.*/
	SamplerBorderColorMax SamplerBorderColor = 6
)

type VertexFrequency = gdclass.RenderingDeviceVertexFrequency //gd:RenderingDevice.VertexFrequency

const (
	/*Vertex attribute addressing is a function of the vertex. This is used to specify the rate at which vertex attributes are pulled from buffers.*/
	VertexFrequencyVertex VertexFrequency = 0
	/*Vertex attribute addressing is a function of the instance index. This is used to specify the rate at which vertex attributes are pulled from buffers.*/
	VertexFrequencyInstance VertexFrequency = 1
)

type IndexBufferFormat = gdclass.RenderingDeviceIndexBufferFormat //gd:RenderingDevice.IndexBufferFormat

const (
	/*Index buffer in 16-bit unsigned integer format. This limits the maximum index that can be specified to [code]65535[/code].*/
	IndexBufferFormatUint16 IndexBufferFormat = 0
	/*Index buffer in 32-bit unsigned integer format. This limits the maximum index that can be specified to [code]4294967295[/code].*/
	IndexBufferFormatUint32 IndexBufferFormat = 1
)

type StorageBufferUsage = gdclass.RenderingDeviceStorageBufferUsage //gd:RenderingDevice.StorageBufferUsage

const (
	StorageBufferUsageDispatchIndirect StorageBufferUsage = 1
)

type BufferCreationBits = gdclass.RenderingDeviceBufferCreationBits //gd:RenderingDevice.BufferCreationBits

const (
	/*Optionally, set this flag if you wish to use [method buffer_get_device_address] functionality. You must first check the GPU supports it:
	  [codeblocks]
	  [gdscript]
	  rd = RenderingServer.get_rendering_device()

	  if rd.has_feature(RenderingDevice.SUPPORTS_BUFFER_DEVICE_ADDRESS):
	        storage_buffer = rd.storage_buffer_create(bytes.size(), bytes, RenderingDevice.STORAGE_BUFFER_USAGE_SHADER_DEVICE_ADDRESS):
	        storage_buffer_address = rd.buffer_get_device_address(storage_buffer)
	  [/gdscript]
	  [/codeblocks]*/
	BufferCreationDeviceAddressBit BufferCreationBits = 1
	/*Set this flag so that it is created as storage. This is useful if Compute Shaders need access (for reading or writing) to the buffer, e.g. skeletal animations are processed in Compute Shaders which need access to vertex buffers, to be later consumed by vertex shaders as part of the regular rasterization pipeline.*/
	BufferCreationAsStorageBit BufferCreationBits = 2
)

type UniformType = gdclass.RenderingDeviceUniformType //gd:RenderingDevice.UniformType

const (
	/*Sampler uniform.*/
	UniformTypeSampler UniformType = 0
	/*Sampler uniform with a texture.*/
	UniformTypeSamplerWithTexture UniformType = 1
	/*Texture uniform.*/
	UniformTypeTexture UniformType = 2
	/*Image uniform.*/
	UniformTypeImage UniformType = 3
	/*Texture buffer uniform.*/
	UniformTypeTextureBuffer UniformType = 4
	/*Sampler uniform with a texture buffer.*/
	UniformTypeSamplerWithTextureBuffer UniformType = 5
	/*Image buffer uniform.*/
	UniformTypeImageBuffer UniformType = 6
	/*Uniform buffer uniform.*/
	UniformTypeUniformBuffer UniformType = 7
	/*[url=https://vkguide.dev/docs/chapter-4/storage_buffers/]Storage buffer[/url] uniform.*/
	UniformTypeStorageBuffer UniformType = 8
	/*Input attachment uniform.*/
	UniformTypeInputAttachment UniformType = 9
	/*Represents the size of the [enum UniformType] enum.*/
	UniformTypeMax UniformType = 10
)

type RenderPrimitive = gdclass.RenderingDeviceRenderPrimitive //gd:RenderingDevice.RenderPrimitive

const (
	/*Point rendering primitive (with constant size, regardless of distance from camera).*/
	RenderPrimitivePoints RenderPrimitive = 0
	/*Line list rendering primitive. Lines are drawn separated from each other.*/
	RenderPrimitiveLines RenderPrimitive = 1
	/*[url=https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#drawing-line-lists-with-adjacency]Line list rendering primitive with adjacency.[/url]
	  [b]Note:[/b] Adjacency is only useful with geometry shaders, which Godot does not expose.*/
	RenderPrimitiveLinesWithAdjacency RenderPrimitive = 2
	/*Line strip rendering primitive. Lines drawn are connected to the previous vertex.*/
	RenderPrimitiveLinestrips RenderPrimitive = 3
	/*[url=https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#drawing-line-strips-with-adjacency]Line strip rendering primitive with adjacency.[/url]
	  [b]Note:[/b] Adjacency is only useful with geometry shaders, which Godot does not expose.*/
	RenderPrimitiveLinestripsWithAdjacency RenderPrimitive = 4
	/*Triangle list rendering primitive. Triangles are drawn separated from each other.*/
	RenderPrimitiveTriangles RenderPrimitive = 5
	/*[url=https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#drawing-triangle-lists-with-adjacency]Triangle list rendering primitive with adjacency.[/url]
	  [b]Note:[/b] Adjacency is only useful with geometry shaders, which Godot does not expose.*/
	RenderPrimitiveTrianglesWithAdjacency RenderPrimitive = 6
	/*Triangle strip rendering primitive. Triangles drawn are connected to the previous triangle.*/
	RenderPrimitiveTriangleStrips RenderPrimitive = 7
	/*[url=https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#drawing-triangle-strips-with-adjacency]Triangle strip rendering primitive with adjacency.[/url]
	  [b]Note:[/b] Adjacency is only useful with geometry shaders, which Godot does not expose.*/
	RenderPrimitiveTriangleStripsWithAjacency RenderPrimitive = 8
	/*Triangle strip rendering primitive with [i]primitive restart[/i] enabled. Triangles drawn are connected to the previous triangle, but a primitive restart index can be specified before drawing to create a second triangle strip after the specified index.
	  [b]Note:[/b] Only compatible with indexed draws.*/
	RenderPrimitiveTriangleStripsWithRestartIndex RenderPrimitive = 9
	/*Tessellation patch rendering primitive. Only useful with tessellation shaders, which can be used to deform these patches.*/
	RenderPrimitiveTesselationPatch RenderPrimitive = 10
	/*Represents the size of the [enum RenderPrimitive] enum.*/
	RenderPrimitiveMax RenderPrimitive = 11
)

type PolygonCullMode = gdclass.RenderingDevicePolygonCullMode //gd:RenderingDevice.PolygonCullMode

const (
	/*Do not use polygon front face or backface culling.*/
	PolygonCullDisabled PolygonCullMode = 0
	/*Use polygon frontface culling (faces pointing towards the camera are hidden).*/
	PolygonCullFront PolygonCullMode = 1
	/*Use polygon backface culling (faces pointing away from the camera are hidden).*/
	PolygonCullBack PolygonCullMode = 2
)

type PolygonFrontFace = gdclass.RenderingDevicePolygonFrontFace //gd:RenderingDevice.PolygonFrontFace

const (
	/*Clockwise winding order to determine which face of a polygon is its front face.*/
	PolygonFrontFaceClockwise PolygonFrontFace = 0
	/*Counter-clockwise winding order to determine which face of a polygon is its front face.*/
	PolygonFrontFaceCounterClockwise PolygonFrontFace = 1
)

type StencilOperation = gdclass.RenderingDeviceStencilOperation //gd:RenderingDevice.StencilOperation

const (
	/*Keep the current stencil value.*/
	StencilOpKeep StencilOperation = 0
	/*Set the stencil value to [code]0[/code].*/
	StencilOpZero StencilOperation = 1
	/*Replace the existing stencil value with the new one.*/
	StencilOpReplace StencilOperation = 2
	/*Increment the existing stencil value and clamp to the maximum representable unsigned value if reached. Stencil bits are considered as an unsigned integer.*/
	StencilOpIncrementAndClamp StencilOperation = 3
	/*Decrement the existing stencil value and clamp to the minimum value if reached. Stencil bits are considered as an unsigned integer.*/
	StencilOpDecrementAndClamp StencilOperation = 4
	/*Bitwise-invert the existing stencil value.*/
	StencilOpInvert StencilOperation = 5
	/*Increment the stencil value and wrap around to [code]0[/code] if reaching the maximum representable unsigned. Stencil bits are considered as an unsigned integer.*/
	StencilOpIncrementAndWrap StencilOperation = 6
	/*Decrement the stencil value and wrap around to the maximum representable unsigned if reaching the minimum. Stencil bits are considered as an unsigned integer.*/
	StencilOpDecrementAndWrap StencilOperation = 7
	/*Represents the size of the [enum StencilOperation] enum.*/
	StencilOpMax StencilOperation = 8
)

type CompareOperator = gdclass.RenderingDeviceCompareOperator //gd:RenderingDevice.CompareOperator

const (
	/*"Never" comparison (opposite of [constant COMPARE_OP_ALWAYS]).*/
	CompareOpNever CompareOperator = 0
	/*"Less than" comparison.*/
	CompareOpLess CompareOperator = 1
	/*"Equal" comparison.*/
	CompareOpEqual CompareOperator = 2
	/*"Less than or equal" comparison.*/
	CompareOpLessOrEqual CompareOperator = 3
	/*"Greater than" comparison.*/
	CompareOpGreater CompareOperator = 4
	/*"Not equal" comparison.*/
	CompareOpNotEqual CompareOperator = 5
	/*"Greater than or equal" comparison.*/
	CompareOpGreaterOrEqual CompareOperator = 6
	/*"Always" comparison (opposite of [constant COMPARE_OP_NEVER]).*/
	CompareOpAlways CompareOperator = 7
	/*Represents the size of the [enum CompareOperator] enum.*/
	CompareOpMax CompareOperator = 8
)

type LogicOperation = gdclass.RenderingDeviceLogicOperation //gd:RenderingDevice.LogicOperation

const (
	/*Clear logic operation (result is always [code]0[/code]). See also [constant LOGIC_OP_SET].*/
	LogicOpClear LogicOperation = 0
	/*AND logic operation.*/
	LogicOpAnd LogicOperation = 1
	/*AND logic operation with the [i]destination[/i] operand being inverted. See also [constant LOGIC_OP_AND_INVERTED].*/
	LogicOpAndReverse LogicOperation = 2
	/*Copy logic operation (keeps the [i]source[/i] value as-is). See also [constant LOGIC_OP_COPY_INVERTED] and [constant LOGIC_OP_NO_OP].*/
	LogicOpCopy LogicOperation = 3
	/*AND logic operation with the [i]source[/i] operand being inverted. See also [constant LOGIC_OP_AND_REVERSE].*/
	LogicOpAndInverted LogicOperation = 4
	/*No-op logic operation (keeps the [i]destination[/i] value as-is). See also [constant LOGIC_OP_COPY].*/
	LogicOpNoOp LogicOperation = 5
	/*Exclusive or (XOR) logic operation.*/
	LogicOpXor LogicOperation = 6
	/*OR logic operation.*/
	LogicOpOr LogicOperation = 7
	/*Not-OR (NOR) logic operation.*/
	LogicOpNor LogicOperation = 8
	/*Not-XOR (XNOR) logic operation.*/
	LogicOpEquivalent LogicOperation = 9
	/*Invert logic operation.*/
	LogicOpInvert LogicOperation = 10
	/*OR logic operation with the [i]destination[/i] operand being inverted. See also [constant LOGIC_OP_OR_REVERSE].*/
	LogicOpOrReverse LogicOperation = 11
	/*NOT logic operation (inverts the value). See also [constant LOGIC_OP_COPY].*/
	LogicOpCopyInverted LogicOperation = 12
	/*OR logic operation with the [i]source[/i] operand being inverted. See also [constant LOGIC_OP_OR_REVERSE].*/
	LogicOpOrInverted LogicOperation = 13
	/*Not-AND (NAND) logic operation.*/
	LogicOpNand LogicOperation = 14
	/*SET logic operation (result is always [code]1[/code]). See also [constant LOGIC_OP_CLEAR].*/
	LogicOpSet LogicOperation = 15
	/*Represents the size of the [enum LogicOperation] enum.*/
	LogicOpMax LogicOperation = 16
)

type BlendFactor = gdclass.RenderingDeviceBlendFactor //gd:RenderingDevice.BlendFactor

const (
	/*Constant [code]0.0[/code] blend factor.*/
	BlendFactorZero BlendFactor = 0
	/*Constant [code]1.0[/code] blend factor.*/
	BlendFactorOne BlendFactor = 1
	/*Color blend factor is [code]source color[/code]. Alpha blend factor is [code]source alpha[/code].*/
	BlendFactorSrcColor BlendFactor = 2
	/*Color blend factor is [code]1.0 - source color[/code]. Alpha blend factor is [code]1.0 - source alpha[/code].*/
	BlendFactorOneMinusSrcColor BlendFactor = 3
	/*Color blend factor is [code]destination color[/code]. Alpha blend factor is [code]destination alpha[/code].*/
	BlendFactorDstColor BlendFactor = 4
	/*Color blend factor is [code]1.0 - destination color[/code]. Alpha blend factor is [code]1.0 - destination alpha[/code].*/
	BlendFactorOneMinusDstColor BlendFactor = 5
	/*Color and alpha blend factor is [code]source alpha[/code].*/
	BlendFactorSrcAlpha BlendFactor = 6
	/*Color and alpha blend factor is [code]1.0 - source alpha[/code].*/
	BlendFactorOneMinusSrcAlpha BlendFactor = 7
	/*Color and alpha blend factor is [code]destination alpha[/code].*/
	BlendFactorDstAlpha BlendFactor = 8
	/*Color and alpha blend factor is [code]1.0 - destination alpha[/code].*/
	BlendFactorOneMinusDstAlpha BlendFactor = 9
	/*Color blend factor is [code]blend constant color[/code]. Alpha blend factor is [code]blend constant alpha[/code] (see [method draw_list_set_blend_constants]).*/
	BlendFactorConstantColor BlendFactor = 10
	/*Color blend factor is [code]1.0 - blend constant color[/code]. Alpha blend factor is [code]1.0 - blend constant alpha[/code] (see [method draw_list_set_blend_constants]).*/
	BlendFactorOneMinusConstantColor BlendFactor = 11
	/*Color and alpha blend factor is [code]blend constant alpha[/code] (see [method draw_list_set_blend_constants]).*/
	BlendFactorConstantAlpha BlendFactor = 12
	/*Color and alpha blend factor is [code]1.0 - blend constant alpha[/code] (see [method draw_list_set_blend_constants]).*/
	BlendFactorOneMinusConstantAlpha BlendFactor = 13
	/*Color blend factor is [code]min(source alpha, 1.0 - destination alpha)[/code]. Alpha blend factor is [code]1.0[/code].*/
	BlendFactorSrcAlphaSaturate BlendFactor = 14
	/*Color blend factor is [code]second source color[/code]. Alpha blend factor is [code]second source alpha[/code]. Only relevant for dual-source blending.*/
	BlendFactorSrc1Color BlendFactor = 15
	/*Color blend factor is [code]1.0 - second source color[/code]. Alpha blend factor is [code]1.0 - second source alpha[/code]. Only relevant for dual-source blending.*/
	BlendFactorOneMinusSrc1Color BlendFactor = 16
	/*Color and alpha blend factor is [code]second source alpha[/code]. Only relevant for dual-source blending.*/
	BlendFactorSrc1Alpha BlendFactor = 17
	/*Color and alpha blend factor is [code]1.0 - second source alpha[/code]. Only relevant for dual-source blending.*/
	BlendFactorOneMinusSrc1Alpha BlendFactor = 18
	/*Represents the size of the [enum BlendFactor] enum.*/
	BlendFactorMax BlendFactor = 19
)

type BlendOperation = gdclass.RenderingDeviceBlendOperation //gd:RenderingDevice.BlendOperation

const (
	/*Additive blending operation ([code]source + destination[/code]).*/
	BlendOpAdd BlendOperation = 0
	/*Subtractive blending operation ([code]source - destination[/code]).*/
	BlendOpSubtract BlendOperation = 1
	/*Reverse subtractive blending operation ([code]destination - source[/code]).*/
	BlendOpReverseSubtract BlendOperation = 2
	/*Minimum blending operation (keep the lowest value of the two).*/
	BlendOpMinimum BlendOperation = 3
	/*Maximum blending operation (keep the highest value of the two).*/
	BlendOpMaximum BlendOperation = 4
	/*Represents the size of the [enum BlendOperation] enum.*/
	BlendOpMax BlendOperation = 5
)

type PipelineDynamicStateFlags = gdclass.RenderingDevicePipelineDynamicStateFlags //gd:RenderingDevice.PipelineDynamicStateFlags

const (
	/*Allows dynamically changing the width of rendering lines.*/
	DynamicStateLineWidth PipelineDynamicStateFlags = 1
	/*Allows dynamically changing the depth bias.*/
	DynamicStateDepthBias          PipelineDynamicStateFlags = 2
	DynamicStateBlendConstants     PipelineDynamicStateFlags = 4
	DynamicStateDepthBounds        PipelineDynamicStateFlags = 8
	DynamicStateStencilCompareMask PipelineDynamicStateFlags = 16
	DynamicStateStencilWriteMask   PipelineDynamicStateFlags = 32
	DynamicStateStencilReference   PipelineDynamicStateFlags = 64
)

type InitialAction = gdclass.RenderingDeviceInitialAction //gd:RenderingDevice.InitialAction

const (
	/*Load the previous contents of the framebuffer.*/
	InitialActionLoad InitialAction = 0
	/*Clear the whole framebuffer or its specified region.*/
	InitialActionClear InitialAction = 1
	/*Ignore the previous contents of the framebuffer. This is the fastest option if you'll overwrite all of the pixels and don't need to read any of them.*/
	InitialActionDiscard InitialAction = 2
	/*Represents the size of the [enum InitialAction] enum.*/
	InitialActionMax                 InitialAction = 3
	InitialActionClearRegion         InitialAction = 1
	InitialActionClearRegionContinue InitialAction = 1
	InitialActionKeep                InitialAction = 0
	InitialActionDrop                InitialAction = 2
	InitialActionContinue            InitialAction = 0
)

type FinalAction = gdclass.RenderingDeviceFinalAction //gd:RenderingDevice.FinalAction

const (
	/*Store the result of the draw list in the framebuffer. This is generally what you want to do.*/
	FinalActionStore FinalAction = 0
	/*Discard the contents of the framebuffer. This is the fastest option if you don't need to use the results of the draw list.*/
	FinalActionDiscard FinalAction = 1
	/*Represents the size of the [enum FinalAction] enum.*/
	FinalActionMax      FinalAction = 2
	FinalActionRead     FinalAction = 0
	FinalActionContinue FinalAction = 0
)

type ShaderStage = gdclass.RenderingDeviceShaderStage //gd:RenderingDevice.ShaderStage

const (
	/*Vertex shader stage. This can be used to manipulate vertices from a shader (but not create new vertices).*/
	ShaderStageVertex ShaderStage = 0
	/*Fragment shader stage (called "pixel shader" in Direct3D). This can be used to manipulate pixels from a shader.*/
	ShaderStageFragment ShaderStage = 1
	/*Tessellation control shader stage. This can be used to create additional geometry from a shader.*/
	ShaderStageTesselationControl ShaderStage = 2
	/*Tessellation evaluation shader stage. This can be used to create additional geometry from a shader.*/
	ShaderStageTesselationEvaluation ShaderStage = 3
	/*Compute shader stage. This can be used to run arbitrary computing tasks in a shader, performing them on the GPU instead of the CPU.*/
	ShaderStageCompute ShaderStage = 4
	/*Represents the size of the [enum ShaderStage] enum.*/
	ShaderStageMax ShaderStage = 5
	/*Vertex shader stage bit (see also [constant SHADER_STAGE_VERTEX]).*/
	ShaderStageVertexBit ShaderStage = 1
	/*Fragment shader stage bit (see also [constant SHADER_STAGE_FRAGMENT]).*/
	ShaderStageFragmentBit ShaderStage = 2
	/*Tessellation control shader stage bit (see also [constant SHADER_STAGE_TESSELATION_CONTROL]).*/
	ShaderStageTesselationControlBit ShaderStage = 4
	/*Tessellation evaluation shader stage bit (see also [constant SHADER_STAGE_TESSELATION_EVALUATION]).*/
	ShaderStageTesselationEvaluationBit ShaderStage = 8
	/*Compute shader stage bit (see also [constant SHADER_STAGE_COMPUTE]).*/
	ShaderStageComputeBit ShaderStage = 16
)

type ShaderLanguage = gdclass.RenderingDeviceShaderLanguage //gd:RenderingDevice.ShaderLanguage

const (
	/*Khronos' GLSL shading language (used natively by OpenGL and Vulkan). This is the language used for core Godot shaders.*/
	ShaderLanguageGlsl ShaderLanguage = 0
	/*Microsoft's High-Level Shading Language (used natively by Direct3D, but can also be used in Vulkan).*/
	ShaderLanguageHlsl ShaderLanguage = 1
)

type PipelineSpecializationConstantType = gdclass.RenderingDevicePipelineSpecializationConstantType //gd:RenderingDevice.PipelineSpecializationConstantType

const (
	/*Boolean specialization constant.*/
	PipelineSpecializationConstantTypeBool PipelineSpecializationConstantType = 0
	/*Integer specialization constant.*/
	PipelineSpecializationConstantTypeInt PipelineSpecializationConstantType = 1
	/*Floating-point specialization constant.*/
	PipelineSpecializationConstantTypeFloat PipelineSpecializationConstantType = 2
)

type Features = gdclass.RenderingDeviceFeatures //gd:RenderingDevice.Features

const (
	/*Features support for buffer device address extension.*/
	SupportsBufferDeviceAddress Features = 6
)

type Limit = gdclass.RenderingDeviceLimit //gd:RenderingDevice.Limit

const (
	/*Maximum number of uniform sets that can be bound at a given time.*/
	LimitMaxBoundUniformSets Limit = 0
	/*Maximum number of color framebuffer attachments that can be used at a given time.*/
	LimitMaxFramebufferColorAttachments Limit = 1
	/*Maximum number of textures that can be used per uniform set.*/
	LimitMaxTexturesPerUniformSet Limit = 2
	/*Maximum number of samplers that can be used per uniform set.*/
	LimitMaxSamplersPerUniformSet Limit = 3
	/*Maximum number of [url=https://vkguide.dev/docs/chapter-4/storage_buffers/]storage buffers[/url] per uniform set.*/
	LimitMaxStorageBuffersPerUniformSet Limit = 4
	/*Maximum number of storage images per uniform set.*/
	LimitMaxStorageImagesPerUniformSet Limit = 5
	/*Maximum number of uniform buffers per uniform set.*/
	LimitMaxUniformBuffersPerUniformSet Limit = 6
	/*Maximum index for an indexed draw command.*/
	LimitMaxDrawIndexedIndex Limit = 7
	/*Maximum height of a framebuffer (in pixels).*/
	LimitMaxFramebufferHeight Limit = 8
	/*Maximum width of a framebuffer (in pixels).*/
	LimitMaxFramebufferWidth Limit = 9
	/*Maximum number of texture array layers.*/
	LimitMaxTextureArrayLayers Limit = 10
	/*Maximum supported 1-dimensional texture size (in pixels on a single axis).*/
	LimitMaxTextureSize1d Limit = 11
	/*Maximum supported 2-dimensional texture size (in pixels on a single axis).*/
	LimitMaxTextureSize2d Limit = 12
	/*Maximum supported 3-dimensional texture size (in pixels on a single axis).*/
	LimitMaxTextureSize3d Limit = 13
	/*Maximum supported cubemap texture size (in pixels on a single axis of a single face).*/
	LimitMaxTextureSizeCube Limit = 14
	/*Maximum number of textures per shader stage.*/
	LimitMaxTexturesPerShaderStage Limit = 15
	/*Maximum number of samplers per shader stage.*/
	LimitMaxSamplersPerShaderStage Limit = 16
	/*Maximum number of [url=https://vkguide.dev/docs/chapter-4/storage_buffers/]storage buffers[/url] per shader stage.*/
	LimitMaxStorageBuffersPerShaderStage Limit = 17
	/*Maximum number of storage images per shader stage.*/
	LimitMaxStorageImagesPerShaderStage Limit = 18
	/*Maximum number of uniform buffers per uniform set.*/
	LimitMaxUniformBuffersPerShaderStage Limit = 19
	/*Maximum size of a push constant. A lot of devices are limited to 128 bytes, so try to avoid exceeding 128 bytes in push constants to ensure compatibility even if your GPU is reporting a higher value.*/
	LimitMaxPushConstantSize Limit = 20
	/*Maximum size of a uniform buffer.*/
	LimitMaxUniformBufferSize Limit = 21
	/*Maximum vertex input attribute offset.*/
	LimitMaxVertexInputAttributeOffset Limit = 22
	/*Maximum number of vertex input attributes.*/
	LimitMaxVertexInputAttributes Limit = 23
	/*Maximum number of vertex input bindings.*/
	LimitMaxVertexInputBindings Limit = 24
	/*Maximum vertex input binding stride.*/
	LimitMaxVertexInputBindingStride Limit = 25
	/*Minimum uniform buffer offset alignment.*/
	LimitMinUniformBufferOffsetAlignment Limit = 26
	/*Maximum shared memory size for compute shaders.*/
	LimitMaxComputeSharedMemorySize Limit = 27
	/*Maximum number of workgroups for compute shaders on the X axis.*/
	LimitMaxComputeWorkgroupCountX Limit = 28
	/*Maximum number of workgroups for compute shaders on the Y axis.*/
	LimitMaxComputeWorkgroupCountY Limit = 29
	/*Maximum number of workgroups for compute shaders on the Z axis.*/
	LimitMaxComputeWorkgroupCountZ Limit = 30
	/*Maximum number of workgroup invocations for compute shaders.*/
	LimitMaxComputeWorkgroupInvocations Limit = 31
	/*Maximum workgroup size for compute shaders on the X axis.*/
	LimitMaxComputeWorkgroupSizeX Limit = 32
	/*Maximum workgroup size for compute shaders on the Y axis.*/
	LimitMaxComputeWorkgroupSizeY Limit = 33
	/*Maximum workgroup size for compute shaders on the Z axis.*/
	LimitMaxComputeWorkgroupSizeZ Limit = 34
	/*Maximum viewport width (in pixels).*/
	LimitMaxViewportDimensionsX Limit = 35
	/*Maximum viewport height (in pixels).*/
	LimitMaxViewportDimensionsY Limit = 36
	/*Returns the smallest value for [member ProjectSettings.rendering/scaling_3d/scale] when using the MetalFX temporal upscaler.
	  [b]Note:[/b] The returned value is multiplied by a factor of [code]1000000[/code] to preserve 6 digits of precision. It must be divided by [code]1000000.0[/code] to convert the value to a floating point number.*/
	LimitMetalfxTemporalScalerMinScale Limit = 46
	/*Returns the largest value for [member ProjectSettings.rendering/scaling_3d/scale] when using the MetalFX temporal upscaler.
	  [b]Note:[/b] The returned value is multiplied by a factor of [code]1000000[/code] to preserve 6 digits of precision. It must be divided by [code]1000000.0[/code] to convert the value to a floating point number.*/
	LimitMetalfxTemporalScalerMaxScale Limit = 47
)

type MemoryType = gdclass.RenderingDeviceMemoryType //gd:RenderingDevice.MemoryType

const (
	/*Memory taken by textures.*/
	MemoryTextures MemoryType = 0
	/*Memory taken by buffers.*/
	MemoryBuffers MemoryType = 1
	/*Total memory taken. This is greater than the sum of [constant MEMORY_TEXTURES] and [constant MEMORY_BUFFERS], as it also includes miscellaneous memory usage.*/
	MemoryTotal MemoryType = 2
)

type BreadcrumbMarker = gdclass.RenderingDeviceBreadcrumbMarker //gd:RenderingDevice.BreadcrumbMarker

const (
	None                  BreadcrumbMarker = 0
	ReflectionProbes      BreadcrumbMarker = 65536
	SkyPass               BreadcrumbMarker = 131072
	LightmapperPass       BreadcrumbMarker = 196608
	ShadowPassDirectional BreadcrumbMarker = 262144
	ShadowPassCube        BreadcrumbMarker = 327680
	OpaquePass            BreadcrumbMarker = 393216
	AlphaPass             BreadcrumbMarker = 458752
	TransparentPass       BreadcrumbMarker = 524288
	PostProcessingPass    BreadcrumbMarker = 589824
	BlitPass              BreadcrumbMarker = 655360
	UiPass                BreadcrumbMarker = 720896
	DebugPass             BreadcrumbMarker = 786432
)

type DrawFlags = gdclass.RenderingDeviceDrawFlags //gd:RenderingDevice.DrawFlags

const (
	/*Do not clear or ignore any attachments.*/
	DrawDefaultAll DrawFlags = 0
	/*Clear the first color attachment.*/
	DrawClearColor0 DrawFlags = 1
	/*Clear the second color attachment.*/
	DrawClearColor1 DrawFlags = 2
	/*Clear the third color attachment.*/
	DrawClearColor2 DrawFlags = 4
	/*Clear the fourth color attachment.*/
	DrawClearColor3 DrawFlags = 8
	/*Clear the fifth color attachment.*/
	DrawClearColor4 DrawFlags = 16
	/*Clear the sixth color attachment.*/
	DrawClearColor5 DrawFlags = 32
	/*Clear the seventh color attachment.*/
	DrawClearColor6 DrawFlags = 64
	/*Clear the eighth color attachment.*/
	DrawClearColor7 DrawFlags = 128
	/*Mask for clearing all color attachments.*/
	DrawClearColorMask DrawFlags = 255
	/*Clear all color attachments.*/
	DrawClearColorAll DrawFlags = 255
	/*Ignore the previous contents of the first color attachment.*/
	DrawIgnoreColor0 DrawFlags = 256
	/*Ignore the previous contents of the second color attachment.*/
	DrawIgnoreColor1 DrawFlags = 512
	/*Ignore the previous contents of the third color attachment.*/
	DrawIgnoreColor2 DrawFlags = 1024
	/*Ignore the previous contents of the fourth color attachment.*/
	DrawIgnoreColor3 DrawFlags = 2048
	/*Ignore the previous contents of the fifth color attachment.*/
	DrawIgnoreColor4 DrawFlags = 4096
	/*Ignore the previous contents of the sixth color attachment.*/
	DrawIgnoreColor5 DrawFlags = 8192
	/*Ignore the previous contents of the seventh color attachment.*/
	DrawIgnoreColor6 DrawFlags = 16384
	/*Ignore the previous contents of the eighth color attachment.*/
	DrawIgnoreColor7 DrawFlags = 32768
	/*Mask for ignoring all the previous contents of the color attachments.*/
	DrawIgnoreColorMask DrawFlags = 65280
	/*Ignore the previous contents of all color attachments.*/
	DrawIgnoreColorAll DrawFlags = 65280
	/*Clear the depth attachment.*/
	DrawClearDepth DrawFlags = 65536
	/*Ignore the previous contents of the depth attachment.*/
	DrawIgnoreDepth DrawFlags = 131072
	/*Clear the stencil attachment.*/
	DrawClearStencil DrawFlags = 262144
	/*Ignore the previous contents of the stencil attachment.*/
	DrawIgnoreStencil DrawFlags = 524288
	/*Clear all attachments.*/
	DrawClearAll DrawFlags = 327935
	/*Ignore the previous contents of all attachments.*/
	DrawIgnoreAll DrawFlags = 720640
)
