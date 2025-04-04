// Package PopupMenu provides methods for working with PopupMenu object instances.
package PopupMenu

import "unsafe"
import "reflect"
import "slices"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/internal/gdclass"
import "graphics.gd/variant"
import "graphics.gd/classdb/Node"
import "graphics.gd/classdb/Popup"
import "graphics.gd/classdb/Viewport"
import "graphics.gd/classdb/Window"
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
[PopupMenu] is a modal window used to display a list of options. Useful for toolbars and context menus.
The size of a [PopupMenu] can be limited by using [member Window.max_size]. If the height of the list of items is larger than the maximum height of the [PopupMenu], a [ScrollContainer] within the popup will allow the user to scroll the contents. If no maximum size is set, or if it is set to [code]0[/code], the [PopupMenu] height will be limited by its parent rect.
All [code]set_*[/code] methods allow negative item indices, i.e. [code]-1[/code] to access the last item, [code]-2[/code] to select the second-to-last item, and so on.
[b]Incremental search:[/b] Like [ItemList] and [Tree], [PopupMenu] supports searching within the list while the control is focused. Press a key that matches the first letter of an item's name to select the first item starting with the given letter. After that point, there are two ways to perform incremental search: 1) Press the same key again before the timeout duration to select the next item starting with the same letter. 2) Press letter keys that match the rest of the word before the timeout duration to match to select the item in question directly. Both of these actions will be reset to the beginning of the list if the timeout duration has passed since the last keystroke was registered. You can adjust the timeout duration by changing [member ProjectSettings.gui/timers/incremental_search_max_interval_msec].
[b]Note:[/b] The ID values used for items are limited to 32 bits, not full 64 bits of [int]. This has a range of [code]-2^32[/code] to [code]2^32 - 1[/code], i.e. [code]-2147483648[/code] to [code]2147483647[/code].
*/
type Instance [1]gdclass.PopupMenu
type Expanded [1]gdclass.PopupMenu

// Nil is a nil/null instance of the class. Equivalent to the zero value.
var Nil Instance

type Any interface {
	gd.IsClass
	AsPopupMenu() Instance
}

/*
Checks the provided [param event] against the [PopupMenu]'s shortcuts and accelerators, and activates the first item with matching events. If [param for_global_only] is [code]true[/code], only shortcuts and accelerators with [code]global[/code] set to [code]true[/code] will be called.
Returns [code]true[/code] if an item was successfully activated.
[b]Note:[/b] Certain [Control]s, such as [MenuButton], will call this method automatically.
*/
func (self Instance) ActivateItemByEvent(event [1]gdclass.InputEvent) bool { //gd:PopupMenu.activate_item_by_event
	return bool(Advanced(self).ActivateItemByEvent(event, false))
}

/*
Checks the provided [param event] against the [PopupMenu]'s shortcuts and accelerators, and activates the first item with matching events. If [param for_global_only] is [code]true[/code], only shortcuts and accelerators with [code]global[/code] set to [code]true[/code] will be called.
Returns [code]true[/code] if an item was successfully activated.
[b]Note:[/b] Certain [Control]s, such as [MenuButton], will call this method automatically.
*/
func (self Expanded) ActivateItemByEvent(event [1]gdclass.InputEvent, for_global_only bool) bool { //gd:PopupMenu.activate_item_by_event
	return bool(Advanced(self).ActivateItemByEvent(event, for_global_only))
}

/*
Returns [code]true[/code] if the system native menu is supported and currently used by this [PopupMenu].
*/
func (self Instance) IsNativeMenu() bool { //gd:PopupMenu.is_native_menu
	return bool(Advanced(self).IsNativeMenu())
}

/*
Adds a new item with text [param label].
An [param id] can optionally be provided, as well as an accelerator ([param accel]). If no [param id] is provided, one will be created from the index. If no [param accel] is provided, then the default value of 0 (corresponding to [constant @GlobalScope.KEY_NONE]) will be assigned to the item (which means it won't have any accelerator). See [method get_item_accelerator] for more info on accelerators.
[b]Note:[/b] The provided [param id] is used only in [signal id_pressed] and [signal id_focused] signals. It's not related to the [code]index[/code] arguments in e.g. [method set_item_checked].
*/
func (self Instance) AddItem(label string) { //gd:PopupMenu.add_item
	Advanced(self).AddItem(String.New(label), int64(-1), 0)
}

/*
Adds a new item with text [param label].
An [param id] can optionally be provided, as well as an accelerator ([param accel]). If no [param id] is provided, one will be created from the index. If no [param accel] is provided, then the default value of 0 (corresponding to [constant @GlobalScope.KEY_NONE]) will be assigned to the item (which means it won't have any accelerator). See [method get_item_accelerator] for more info on accelerators.
[b]Note:[/b] The provided [param id] is used only in [signal id_pressed] and [signal id_focused] signals. It's not related to the [code]index[/code] arguments in e.g. [method set_item_checked].
*/
func (self Expanded) AddItem(label string, id int, accel Key) { //gd:PopupMenu.add_item
	Advanced(self).AddItem(String.New(label), int64(id), accel)
}

/*
Adds a new item with text [param label] and icon [param texture].
An [param id] can optionally be provided, as well as an accelerator ([param accel]). If no [param id] is provided, one will be created from the index. If no [param accel] is provided, then the default value of 0 (corresponding to [constant @GlobalScope.KEY_NONE]) will be assigned to the item (which means it won't have any accelerator). See [method get_item_accelerator] for more info on accelerators.
*/
func (self Instance) AddIconItem(texture [1]gdclass.Texture2D, label string) { //gd:PopupMenu.add_icon_item
	Advanced(self).AddIconItem(texture, String.New(label), int64(-1), 0)
}

/*
Adds a new item with text [param label] and icon [param texture].
An [param id] can optionally be provided, as well as an accelerator ([param accel]). If no [param id] is provided, one will be created from the index. If no [param accel] is provided, then the default value of 0 (corresponding to [constant @GlobalScope.KEY_NONE]) will be assigned to the item (which means it won't have any accelerator). See [method get_item_accelerator] for more info on accelerators.
*/
func (self Expanded) AddIconItem(texture [1]gdclass.Texture2D, label string, id int, accel Key) { //gd:PopupMenu.add_icon_item
	Advanced(self).AddIconItem(texture, String.New(label), int64(id), accel)
}

/*
Adds a new checkable item with text [param label].
An [param id] can optionally be provided, as well as an accelerator ([param accel]). If no [param id] is provided, one will be created from the index. If no [param accel] is provided, then the default value of 0 (corresponding to [constant @GlobalScope.KEY_NONE]) will be assigned to the item (which means it won't have any accelerator). See [method get_item_accelerator] for more info on accelerators.
[b]Note:[/b] Checkable items just display a checkmark, but don't have any built-in checking behavior and must be checked/unchecked manually. See [method set_item_checked] for more info on how to control it.
*/
func (self Instance) AddCheckItem(label string) { //gd:PopupMenu.add_check_item
	Advanced(self).AddCheckItem(String.New(label), int64(-1), 0)
}

/*
Adds a new checkable item with text [param label].
An [param id] can optionally be provided, as well as an accelerator ([param accel]). If no [param id] is provided, one will be created from the index. If no [param accel] is provided, then the default value of 0 (corresponding to [constant @GlobalScope.KEY_NONE]) will be assigned to the item (which means it won't have any accelerator). See [method get_item_accelerator] for more info on accelerators.
[b]Note:[/b] Checkable items just display a checkmark, but don't have any built-in checking behavior and must be checked/unchecked manually. See [method set_item_checked] for more info on how to control it.
*/
func (self Expanded) AddCheckItem(label string, id int, accel Key) { //gd:PopupMenu.add_check_item
	Advanced(self).AddCheckItem(String.New(label), int64(id), accel)
}

/*
Adds a new checkable item with text [param label] and icon [param texture].
An [param id] can optionally be provided, as well as an accelerator ([param accel]). If no [param id] is provided, one will be created from the index. If no [param accel] is provided, then the default value of 0 (corresponding to [constant @GlobalScope.KEY_NONE]) will be assigned to the item (which means it won't have any accelerator). See [method get_item_accelerator] for more info on accelerators.
[b]Note:[/b] Checkable items just display a checkmark, but don't have any built-in checking behavior and must be checked/unchecked manually. See [method set_item_checked] for more info on how to control it.
*/
func (self Instance) AddIconCheckItem(texture [1]gdclass.Texture2D, label string) { //gd:PopupMenu.add_icon_check_item
	Advanced(self).AddIconCheckItem(texture, String.New(label), int64(-1), 0)
}

/*
Adds a new checkable item with text [param label] and icon [param texture].
An [param id] can optionally be provided, as well as an accelerator ([param accel]). If no [param id] is provided, one will be created from the index. If no [param accel] is provided, then the default value of 0 (corresponding to [constant @GlobalScope.KEY_NONE]) will be assigned to the item (which means it won't have any accelerator). See [method get_item_accelerator] for more info on accelerators.
[b]Note:[/b] Checkable items just display a checkmark, but don't have any built-in checking behavior and must be checked/unchecked manually. See [method set_item_checked] for more info on how to control it.
*/
func (self Expanded) AddIconCheckItem(texture [1]gdclass.Texture2D, label string, id int, accel Key) { //gd:PopupMenu.add_icon_check_item
	Advanced(self).AddIconCheckItem(texture, String.New(label), int64(id), accel)
}

/*
Adds a new radio check button with text [param label].
An [param id] can optionally be provided, as well as an accelerator ([param accel]). If no [param id] is provided, one will be created from the index. If no [param accel] is provided, then the default value of 0 (corresponding to [constant @GlobalScope.KEY_NONE]) will be assigned to the item (which means it won't have any accelerator). See [method get_item_accelerator] for more info on accelerators.
[b]Note:[/b] Checkable items just display a checkmark, but don't have any built-in checking behavior and must be checked/unchecked manually. See [method set_item_checked] for more info on how to control it.
*/
func (self Instance) AddRadioCheckItem(label string) { //gd:PopupMenu.add_radio_check_item
	Advanced(self).AddRadioCheckItem(String.New(label), int64(-1), 0)
}

/*
Adds a new radio check button with text [param label].
An [param id] can optionally be provided, as well as an accelerator ([param accel]). If no [param id] is provided, one will be created from the index. If no [param accel] is provided, then the default value of 0 (corresponding to [constant @GlobalScope.KEY_NONE]) will be assigned to the item (which means it won't have any accelerator). See [method get_item_accelerator] for more info on accelerators.
[b]Note:[/b] Checkable items just display a checkmark, but don't have any built-in checking behavior and must be checked/unchecked manually. See [method set_item_checked] for more info on how to control it.
*/
func (self Expanded) AddRadioCheckItem(label string, id int, accel Key) { //gd:PopupMenu.add_radio_check_item
	Advanced(self).AddRadioCheckItem(String.New(label), int64(id), accel)
}

/*
Same as [method add_icon_check_item], but uses a radio check button.
*/
func (self Instance) AddIconRadioCheckItem(texture [1]gdclass.Texture2D, label string) { //gd:PopupMenu.add_icon_radio_check_item
	Advanced(self).AddIconRadioCheckItem(texture, String.New(label), int64(-1), 0)
}

/*
Same as [method add_icon_check_item], but uses a radio check button.
*/
func (self Expanded) AddIconRadioCheckItem(texture [1]gdclass.Texture2D, label string, id int, accel Key) { //gd:PopupMenu.add_icon_radio_check_item
	Advanced(self).AddIconRadioCheckItem(texture, String.New(label), int64(id), accel)
}

/*
Adds a new multistate item with text [param label].
Contrarily to normal binary items, multistate items can have more than two states, as defined by [param max_states]. The default value is defined by [param default_state].
An [param id] can optionally be provided, as well as an accelerator ([param accel]). If no [param id] is provided, one will be created from the index. If no [param accel] is provided, then the default value of 0 (corresponding to [constant @GlobalScope.KEY_NONE]) will be assigned to the item (which means it won't have any accelerator). See [method get_item_accelerator] for more info on accelerators.
[codeblock]
func _ready():

	add_multistate_item("Item", 3, 0)

	index_pressed.connect(func(index: int):
	        toggle_item_multistate(index)
	        match get_item_multistate(index):
	            0:
	                print("First state")
	            1:
	                print("Second state")
	            2:
	                print("Third state")
	    )

[/codeblock]
[b]Note:[/b] Multistate items don't update their state automatically and must be done manually. See [method toggle_item_multistate], [method set_item_multistate] and [method get_item_multistate] for more info on how to control it.
*/
func (self Instance) AddMultistateItem(label string, max_states int) { //gd:PopupMenu.add_multistate_item
	Advanced(self).AddMultistateItem(String.New(label), int64(max_states), int64(0), int64(-1), 0)
}

/*
Adds a new multistate item with text [param label].
Contrarily to normal binary items, multistate items can have more than two states, as defined by [param max_states]. The default value is defined by [param default_state].
An [param id] can optionally be provided, as well as an accelerator ([param accel]). If no [param id] is provided, one will be created from the index. If no [param accel] is provided, then the default value of 0 (corresponding to [constant @GlobalScope.KEY_NONE]) will be assigned to the item (which means it won't have any accelerator). See [method get_item_accelerator] for more info on accelerators.
[codeblock]
func _ready():

	add_multistate_item("Item", 3, 0)

	index_pressed.connect(func(index: int):
	        toggle_item_multistate(index)
	        match get_item_multistate(index):
	            0:
	                print("First state")
	            1:
	                print("Second state")
	            2:
	                print("Third state")
	    )

[/codeblock]
[b]Note:[/b] Multistate items don't update their state automatically and must be done manually. See [method toggle_item_multistate], [method set_item_multistate] and [method get_item_multistate] for more info on how to control it.
*/
func (self Expanded) AddMultistateItem(label string, max_states int, default_state int, id int, accel Key) { //gd:PopupMenu.add_multistate_item
	Advanced(self).AddMultistateItem(String.New(label), int64(max_states), int64(default_state), int64(id), accel)
}

/*
Adds a [Shortcut].
An [param id] can optionally be provided. If no [param id] is provided, one will be created from the index.
If [param allow_echo] is [code]true[/code], the shortcut can be activated with echo events.
*/
func (self Instance) AddShortcut(shortcut [1]gdclass.Shortcut) { //gd:PopupMenu.add_shortcut
	Advanced(self).AddShortcut(shortcut, int64(-1), false, false)
}

/*
Adds a [Shortcut].
An [param id] can optionally be provided. If no [param id] is provided, one will be created from the index.
If [param allow_echo] is [code]true[/code], the shortcut can be activated with echo events.
*/
func (self Expanded) AddShortcut(shortcut [1]gdclass.Shortcut, id int, global bool, allow_echo bool) { //gd:PopupMenu.add_shortcut
	Advanced(self).AddShortcut(shortcut, int64(id), global, allow_echo)
}

/*
Adds a new item and assigns the specified [Shortcut] and icon [param texture] to it. Sets the label of the checkbox to the [Shortcut]'s name.
An [param id] can optionally be provided. If no [param id] is provided, one will be created from the index.
If [param allow_echo] is [code]true[/code], the shortcut can be activated with echo events.
*/
func (self Instance) AddIconShortcut(texture [1]gdclass.Texture2D, shortcut [1]gdclass.Shortcut) { //gd:PopupMenu.add_icon_shortcut
	Advanced(self).AddIconShortcut(texture, shortcut, int64(-1), false, false)
}

/*
Adds a new item and assigns the specified [Shortcut] and icon [param texture] to it. Sets the label of the checkbox to the [Shortcut]'s name.
An [param id] can optionally be provided. If no [param id] is provided, one will be created from the index.
If [param allow_echo] is [code]true[/code], the shortcut can be activated with echo events.
*/
func (self Expanded) AddIconShortcut(texture [1]gdclass.Texture2D, shortcut [1]gdclass.Shortcut, id int, global bool, allow_echo bool) { //gd:PopupMenu.add_icon_shortcut
	Advanced(self).AddIconShortcut(texture, shortcut, int64(id), global, allow_echo)
}

/*
Adds a new checkable item and assigns the specified [Shortcut] to it. Sets the label of the checkbox to the [Shortcut]'s name.
An [param id] can optionally be provided. If no [param id] is provided, one will be created from the index.
[b]Note:[/b] Checkable items just display a checkmark, but don't have any built-in checking behavior and must be checked/unchecked manually. See [method set_item_checked] for more info on how to control it.
*/
func (self Instance) AddCheckShortcut(shortcut [1]gdclass.Shortcut) { //gd:PopupMenu.add_check_shortcut
	Advanced(self).AddCheckShortcut(shortcut, int64(-1), false)
}

/*
Adds a new checkable item and assigns the specified [Shortcut] to it. Sets the label of the checkbox to the [Shortcut]'s name.
An [param id] can optionally be provided. If no [param id] is provided, one will be created from the index.
[b]Note:[/b] Checkable items just display a checkmark, but don't have any built-in checking behavior and must be checked/unchecked manually. See [method set_item_checked] for more info on how to control it.
*/
func (self Expanded) AddCheckShortcut(shortcut [1]gdclass.Shortcut, id int, global bool) { //gd:PopupMenu.add_check_shortcut
	Advanced(self).AddCheckShortcut(shortcut, int64(id), global)
}

/*
Adds a new checkable item and assigns the specified [Shortcut] and icon [param texture] to it. Sets the label of the checkbox to the [Shortcut]'s name.
An [param id] can optionally be provided. If no [param id] is provided, one will be created from the index.
[b]Note:[/b] Checkable items just display a checkmark, but don't have any built-in checking behavior and must be checked/unchecked manually. See [method set_item_checked] for more info on how to control it.
*/
func (self Instance) AddIconCheckShortcut(texture [1]gdclass.Texture2D, shortcut [1]gdclass.Shortcut) { //gd:PopupMenu.add_icon_check_shortcut
	Advanced(self).AddIconCheckShortcut(texture, shortcut, int64(-1), false)
}

/*
Adds a new checkable item and assigns the specified [Shortcut] and icon [param texture] to it. Sets the label of the checkbox to the [Shortcut]'s name.
An [param id] can optionally be provided. If no [param id] is provided, one will be created from the index.
[b]Note:[/b] Checkable items just display a checkmark, but don't have any built-in checking behavior and must be checked/unchecked manually. See [method set_item_checked] for more info on how to control it.
*/
func (self Expanded) AddIconCheckShortcut(texture [1]gdclass.Texture2D, shortcut [1]gdclass.Shortcut, id int, global bool) { //gd:PopupMenu.add_icon_check_shortcut
	Advanced(self).AddIconCheckShortcut(texture, shortcut, int64(id), global)
}

/*
Adds a new radio check button and assigns a [Shortcut] to it. Sets the label of the checkbox to the [Shortcut]'s name.
An [param id] can optionally be provided. If no [param id] is provided, one will be created from the index.
[b]Note:[/b] Checkable items just display a checkmark, but don't have any built-in checking behavior and must be checked/unchecked manually. See [method set_item_checked] for more info on how to control it.
*/
func (self Instance) AddRadioCheckShortcut(shortcut [1]gdclass.Shortcut) { //gd:PopupMenu.add_radio_check_shortcut
	Advanced(self).AddRadioCheckShortcut(shortcut, int64(-1), false)
}

/*
Adds a new radio check button and assigns a [Shortcut] to it. Sets the label of the checkbox to the [Shortcut]'s name.
An [param id] can optionally be provided. If no [param id] is provided, one will be created from the index.
[b]Note:[/b] Checkable items just display a checkmark, but don't have any built-in checking behavior and must be checked/unchecked manually. See [method set_item_checked] for more info on how to control it.
*/
func (self Expanded) AddRadioCheckShortcut(shortcut [1]gdclass.Shortcut, id int, global bool) { //gd:PopupMenu.add_radio_check_shortcut
	Advanced(self).AddRadioCheckShortcut(shortcut, int64(id), global)
}

/*
Same as [method add_icon_check_shortcut], but uses a radio check button.
*/
func (self Instance) AddIconRadioCheckShortcut(texture [1]gdclass.Texture2D, shortcut [1]gdclass.Shortcut) { //gd:PopupMenu.add_icon_radio_check_shortcut
	Advanced(self).AddIconRadioCheckShortcut(texture, shortcut, int64(-1), false)
}

/*
Same as [method add_icon_check_shortcut], but uses a radio check button.
*/
func (self Expanded) AddIconRadioCheckShortcut(texture [1]gdclass.Texture2D, shortcut [1]gdclass.Shortcut, id int, global bool) { //gd:PopupMenu.add_icon_radio_check_shortcut
	Advanced(self).AddIconRadioCheckShortcut(texture, shortcut, int64(id), global)
}

/*
Adds an item that will act as a submenu of the parent [PopupMenu] node when clicked. The [param submenu] argument must be the name of an existing [PopupMenu] that has been added as a child to this node. This submenu will be shown when the item is clicked, hovered for long enough, or activated using the [code]ui_select[/code] or [code]ui_right[/code] input actions.
An [param id] can optionally be provided. If no [param id] is provided, one will be created from the index.
*/
func (self Instance) AddSubmenuItem(label string, submenu string) { //gd:PopupMenu.add_submenu_item
	Advanced(self).AddSubmenuItem(String.New(label), String.New(submenu), int64(-1))
}

/*
Adds an item that will act as a submenu of the parent [PopupMenu] node when clicked. The [param submenu] argument must be the name of an existing [PopupMenu] that has been added as a child to this node. This submenu will be shown when the item is clicked, hovered for long enough, or activated using the [code]ui_select[/code] or [code]ui_right[/code] input actions.
An [param id] can optionally be provided. If no [param id] is provided, one will be created from the index.
*/
func (self Expanded) AddSubmenuItem(label string, submenu string, id int) { //gd:PopupMenu.add_submenu_item
	Advanced(self).AddSubmenuItem(String.New(label), String.New(submenu), int64(id))
}

/*
Adds an item that will act as a submenu of the parent [PopupMenu] node when clicked. This submenu will be shown when the item is clicked, hovered for long enough, or activated using the [code]ui_select[/code] or [code]ui_right[/code] input actions.
[param submenu] must be either child of this [PopupMenu] or has no parent node (in which case it will be automatically added as a child). If the [param submenu] popup has another parent, this method will fail.
An [param id] can optionally be provided. If no [param id] is provided, one will be created from the index.
*/
func (self Instance) AddSubmenuNodeItem(label string, submenu [1]gdclass.PopupMenu) { //gd:PopupMenu.add_submenu_node_item
	Advanced(self).AddSubmenuNodeItem(String.New(label), submenu, int64(-1))
}

/*
Adds an item that will act as a submenu of the parent [PopupMenu] node when clicked. This submenu will be shown when the item is clicked, hovered for long enough, or activated using the [code]ui_select[/code] or [code]ui_right[/code] input actions.
[param submenu] must be either child of this [PopupMenu] or has no parent node (in which case it will be automatically added as a child). If the [param submenu] popup has another parent, this method will fail.
An [param id] can optionally be provided. If no [param id] is provided, one will be created from the index.
*/
func (self Expanded) AddSubmenuNodeItem(label string, submenu [1]gdclass.PopupMenu, id int) { //gd:PopupMenu.add_submenu_node_item
	Advanced(self).AddSubmenuNodeItem(String.New(label), submenu, int64(id))
}

/*
Sets the text of the item at the given [param index].
*/
func (self Instance) SetItemText(index int, text string) { //gd:PopupMenu.set_item_text
	Advanced(self).SetItemText(int64(index), String.New(text))
}

/*
Sets item's text base writing direction.
*/
func (self Instance) SetItemTextDirection(index int, direction gdclass.ControlTextDirection) { //gd:PopupMenu.set_item_text_direction
	Advanced(self).SetItemTextDirection(int64(index), direction)
}

/*
Sets language code of item's text used for line-breaking and text shaping algorithms, if left empty current locale is used instead.
*/
func (self Instance) SetItemLanguage(index int, language string) { //gd:PopupMenu.set_item_language
	Advanced(self).SetItemLanguage(int64(index), String.New(language))
}

/*
Replaces the [Texture2D] icon of the item at the given [param index].
*/
func (self Instance) SetItemIcon(index int, icon [1]gdclass.Texture2D) { //gd:PopupMenu.set_item_icon
	Advanced(self).SetItemIcon(int64(index), icon)
}

/*
Sets the maximum allowed width of the icon for the item at the given [param index]. This limit is applied on top of the default size of the icon and on top of [theme_item icon_max_width]. The height is adjusted according to the icon's ratio.
*/
func (self Instance) SetItemIconMaxWidth(index int, width int) { //gd:PopupMenu.set_item_icon_max_width
	Advanced(self).SetItemIconMaxWidth(int64(index), int64(width))
}

/*
Sets a modulating [Color] of the item's icon at the given [param index].
*/
func (self Instance) SetItemIconModulate(index int, modulate Color.RGBA) { //gd:PopupMenu.set_item_icon_modulate
	Advanced(self).SetItemIconModulate(int64(index), Color.RGBA(modulate))
}

/*
Sets the checkstate status of the item at the given [param index].
*/
func (self Instance) SetItemChecked(index int, checked bool) { //gd:PopupMenu.set_item_checked
	Advanced(self).SetItemChecked(int64(index), checked)
}

/*
Sets the [param id] of the item at the given [param index].
The [param id] is used in [signal id_pressed] and [signal id_focused] signals.
*/
func (self Instance) SetItemId(index int, id int) { //gd:PopupMenu.set_item_id
	Advanced(self).SetItemId(int64(index), int64(id))
}

/*
Sets the accelerator of the item at the given [param index]. An accelerator is a keyboard shortcut that can be pressed to trigger the menu button even if it's not currently open. [param accel] is generally a combination of [enum KeyModifierMask]s and [enum Key]s using bitwise OR such as [code]KEY_MASK_CTRL | KEY_A[/code] ([kbd]Ctrl + A[/kbd]).
*/
func (self Instance) SetItemAccelerator(index int, accel Key) { //gd:PopupMenu.set_item_accelerator
	Advanced(self).SetItemAccelerator(int64(index), accel)
}

/*
Sets the metadata of an item, which may be of any type. You can later get it with [method get_item_metadata], which provides a simple way of assigning context data to items.
*/
func (self Instance) SetItemMetadata(index int, metadata any) { //gd:PopupMenu.set_item_metadata
	Advanced(self).SetItemMetadata(int64(index), variant.New(metadata))
}

/*
Enables/disables the item at the given [param index]. When it is disabled, it can't be selected and its action can't be invoked.
*/
func (self Instance) SetItemDisabled(index int, disabled bool) { //gd:PopupMenu.set_item_disabled
	Advanced(self).SetItemDisabled(int64(index), disabled)
}

/*
Sets the submenu of the item at the given [param index]. The submenu is the name of a child [PopupMenu] node that would be shown when the item is clicked.
*/
func (self Instance) SetItemSubmenu(index int, submenu string) { //gd:PopupMenu.set_item_submenu
	Advanced(self).SetItemSubmenu(int64(index), String.New(submenu))
}

/*
Sets the submenu of the item at the given [param index]. The submenu is a [PopupMenu] node that would be shown when the item is clicked. It must either be a child of this [PopupMenu] or has no parent (in which case it will be automatically added as a child). If the [param submenu] popup has another parent, this method will fail.
*/
func (self Instance) SetItemSubmenuNode(index int, submenu [1]gdclass.PopupMenu) { //gd:PopupMenu.set_item_submenu_node
	Advanced(self).SetItemSubmenuNode(int64(index), submenu)
}

/*
Mark the item at the given [param index] as a separator, which means that it would be displayed as a line. If [code]false[/code], sets the type of the item to plain text.
*/
func (self Instance) SetItemAsSeparator(index int, enable bool) { //gd:PopupMenu.set_item_as_separator
	Advanced(self).SetItemAsSeparator(int64(index), enable)
}

/*
Sets whether the item at the given [param index] has a checkbox. If [code]false[/code], sets the type of the item to plain text.
[b]Note:[/b] Checkable items just display a checkmark, but don't have any built-in checking behavior and must be checked/unchecked manually.
*/
func (self Instance) SetItemAsCheckable(index int, enable bool) { //gd:PopupMenu.set_item_as_checkable
	Advanced(self).SetItemAsCheckable(int64(index), enable)
}

/*
Sets the type of the item at the given [param index] to radio button. If [code]false[/code], sets the type of the item to plain text.
*/
func (self Instance) SetItemAsRadioCheckable(index int, enable bool) { //gd:PopupMenu.set_item_as_radio_checkable
	Advanced(self).SetItemAsRadioCheckable(int64(index), enable)
}

/*
Sets the [String] tooltip of the item at the given [param index].
*/
func (self Instance) SetItemTooltip(index int, tooltip string) { //gd:PopupMenu.set_item_tooltip
	Advanced(self).SetItemTooltip(int64(index), String.New(tooltip))
}

/*
Sets a [Shortcut] for the item at the given [param index].
*/
func (self Instance) SetItemShortcut(index int, shortcut [1]gdclass.Shortcut) { //gd:PopupMenu.set_item_shortcut
	Advanced(self).SetItemShortcut(int64(index), shortcut, false)
}

/*
Sets a [Shortcut] for the item at the given [param index].
*/
func (self Expanded) SetItemShortcut(index int, shortcut [1]gdclass.Shortcut, global bool) { //gd:PopupMenu.set_item_shortcut
	Advanced(self).SetItemShortcut(int64(index), shortcut, global)
}

/*
Sets the horizontal offset of the item at the given [param index].
*/
func (self Instance) SetItemIndent(index int, indent int) { //gd:PopupMenu.set_item_indent
	Advanced(self).SetItemIndent(int64(index), int64(indent))
}

/*
Sets the state of a multistate item. See [method add_multistate_item] for details.
*/
func (self Instance) SetItemMultistate(index int, state int) { //gd:PopupMenu.set_item_multistate
	Advanced(self).SetItemMultistate(int64(index), int64(state))
}

/*
Sets the max states of a multistate item. See [method add_multistate_item] for details.
*/
func (self Instance) SetItemMultistateMax(index int, max_states int) { //gd:PopupMenu.set_item_multistate_max
	Advanced(self).SetItemMultistateMax(int64(index), int64(max_states))
}

/*
Disables the [Shortcut] of the item at the given [param index].
*/
func (self Instance) SetItemShortcutDisabled(index int, disabled bool) { //gd:PopupMenu.set_item_shortcut_disabled
	Advanced(self).SetItemShortcutDisabled(int64(index), disabled)
}

/*
Toggles the check state of the item at the given [param index].
*/
func (self Instance) ToggleItemChecked(index int) { //gd:PopupMenu.toggle_item_checked
	Advanced(self).ToggleItemChecked(int64(index))
}

/*
Cycle to the next state of a multistate item. See [method add_multistate_item] for details.
*/
func (self Instance) ToggleItemMultistate(index int) { //gd:PopupMenu.toggle_item_multistate
	Advanced(self).ToggleItemMultistate(int64(index))
}

/*
Returns the text of the item at the given [param index].
*/
func (self Instance) GetItemText(index int) string { //gd:PopupMenu.get_item_text
	return string(Advanced(self).GetItemText(int64(index)).String())
}

/*
Returns item's text base writing direction.
*/
func (self Instance) GetItemTextDirection(index int) gdclass.ControlTextDirection { //gd:PopupMenu.get_item_text_direction
	return gdclass.ControlTextDirection(Advanced(self).GetItemTextDirection(int64(index)))
}

/*
Returns item's text language code.
*/
func (self Instance) GetItemLanguage(index int) string { //gd:PopupMenu.get_item_language
	return string(Advanced(self).GetItemLanguage(int64(index)).String())
}

/*
Returns the icon of the item at the given [param index].
*/
func (self Instance) GetItemIcon(index int) [1]gdclass.Texture2D { //gd:PopupMenu.get_item_icon
	return [1]gdclass.Texture2D(Advanced(self).GetItemIcon(int64(index)))
}

/*
Returns the maximum allowed width of the icon for the item at the given [param index].
*/
func (self Instance) GetItemIconMaxWidth(index int) int { //gd:PopupMenu.get_item_icon_max_width
	return int(int(Advanced(self).GetItemIconMaxWidth(int64(index))))
}

/*
Returns a [Color] modulating the item's icon at the given [param index].
*/
func (self Instance) GetItemIconModulate(index int) Color.RGBA { //gd:PopupMenu.get_item_icon_modulate
	return Color.RGBA(Advanced(self).GetItemIconModulate(int64(index)))
}

/*
Returns [code]true[/code] if the item at the given [param index] is checked.
*/
func (self Instance) IsItemChecked(index int) bool { //gd:PopupMenu.is_item_checked
	return bool(Advanced(self).IsItemChecked(int64(index)))
}

/*
Returns the ID of the item at the given [param index]. [code]id[/code] can be manually assigned, while index can not.
*/
func (self Instance) GetItemId(index int) int { //gd:PopupMenu.get_item_id
	return int(int(Advanced(self).GetItemId(int64(index))))
}

/*
Returns the index of the item containing the specified [param id]. Index is automatically assigned to each item by the engine and can not be set manually.
*/
func (self Instance) GetItemIndex(id int) int { //gd:PopupMenu.get_item_index
	return int(int(Advanced(self).GetItemIndex(int64(id))))
}

/*
Returns the accelerator of the item at the given [param index]. An accelerator is a keyboard shortcut that can be pressed to trigger the menu button even if it's not currently open. The return value is an integer which is generally a combination of [enum KeyModifierMask]s and [enum Key]s using bitwise OR such as [code]KEY_MASK_CTRL | KEY_A[/code] ([kbd]Ctrl + A[/kbd]). If no accelerator is defined for the specified [param index], [method get_item_accelerator] returns [code]0[/code] (corresponding to [constant @GlobalScope.KEY_NONE]).
*/
func (self Instance) GetItemAccelerator(index int) Key { //gd:PopupMenu.get_item_accelerator
	return Key(Advanced(self).GetItemAccelerator(int64(index)))
}

/*
Returns the metadata of the specified item, which might be of any type. You can set it with [method set_item_metadata], which provides a simple way of assigning context data to items.
*/
func (self Instance) GetItemMetadata(index int) any { //gd:PopupMenu.get_item_metadata
	return any(Advanced(self).GetItemMetadata(int64(index)).Interface())
}

/*
Returns [code]true[/code] if the item at the given [param index] is disabled. When it is disabled it can't be selected, or its action invoked.
See [method set_item_disabled] for more info on how to disable an item.
*/
func (self Instance) IsItemDisabled(index int) bool { //gd:PopupMenu.is_item_disabled
	return bool(Advanced(self).IsItemDisabled(int64(index)))
}

/*
Returns the submenu name of the item at the given [param index]. See [method add_submenu_item] for more info on how to add a submenu.
*/
func (self Instance) GetItemSubmenu(index int) string { //gd:PopupMenu.get_item_submenu
	return string(Advanced(self).GetItemSubmenu(int64(index)).String())
}

/*
Returns the submenu of the item at the given [param index], or [code]null[/code] if no submenu was added. See [method add_submenu_node_item] for more info on how to add a submenu.
*/
func (self Instance) GetItemSubmenuNode(index int) [1]gdclass.PopupMenu { //gd:PopupMenu.get_item_submenu_node
	return [1]gdclass.PopupMenu(Advanced(self).GetItemSubmenuNode(int64(index)))
}

/*
Returns [code]true[/code] if the item is a separator. If it is, it will be displayed as a line. See [method add_separator] for more info on how to add a separator.
*/
func (self Instance) IsItemSeparator(index int) bool { //gd:PopupMenu.is_item_separator
	return bool(Advanced(self).IsItemSeparator(int64(index)))
}

/*
Returns [code]true[/code] if the item at the given [param index] is checkable in some way, i.e. if it has a checkbox or radio button.
[b]Note:[/b] Checkable items just display a checkmark or radio button, but don't have any built-in checking behavior and must be checked/unchecked manually.
*/
func (self Instance) IsItemCheckable(index int) bool { //gd:PopupMenu.is_item_checkable
	return bool(Advanced(self).IsItemCheckable(int64(index)))
}

/*
Returns [code]true[/code] if the item at the given [param index] has radio button-style checkability.
[b]Note:[/b] This is purely cosmetic; you must add the logic for checking/unchecking items in radio groups.
*/
func (self Instance) IsItemRadioCheckable(index int) bool { //gd:PopupMenu.is_item_radio_checkable
	return bool(Advanced(self).IsItemRadioCheckable(int64(index)))
}

/*
Returns [code]true[/code] if the specified item's shortcut is disabled.
*/
func (self Instance) IsItemShortcutDisabled(index int) bool { //gd:PopupMenu.is_item_shortcut_disabled
	return bool(Advanced(self).IsItemShortcutDisabled(int64(index)))
}

/*
Returns the tooltip associated with the item at the given [param index].
*/
func (self Instance) GetItemTooltip(index int) string { //gd:PopupMenu.get_item_tooltip
	return string(Advanced(self).GetItemTooltip(int64(index)).String())
}

/*
Returns the [Shortcut] associated with the item at the given [param index].
*/
func (self Instance) GetItemShortcut(index int) [1]gdclass.Shortcut { //gd:PopupMenu.get_item_shortcut
	return [1]gdclass.Shortcut(Advanced(self).GetItemShortcut(int64(index)))
}

/*
Returns the horizontal offset of the item at the given [param index].
*/
func (self Instance) GetItemIndent(index int) int { //gd:PopupMenu.get_item_indent
	return int(int(Advanced(self).GetItemIndent(int64(index))))
}

/*
Returns the max states of the item at the given [param index].
*/
func (self Instance) GetItemMultistateMax(index int) int { //gd:PopupMenu.get_item_multistate_max
	return int(int(Advanced(self).GetItemMultistateMax(int64(index))))
}

/*
Returns the state of the item at the given [param index].
*/
func (self Instance) GetItemMultistate(index int) int { //gd:PopupMenu.get_item_multistate
	return int(int(Advanced(self).GetItemMultistate(int64(index))))
}

/*
Sets the currently focused item as the given [param index].
Passing [code]-1[/code] as the index makes so that no item is focused.
*/
func (self Instance) SetFocusedItem(index int) { //gd:PopupMenu.set_focused_item
	Advanced(self).SetFocusedItem(int64(index))
}

/*
Returns the index of the currently focused item. Returns [code]-1[/code] if no item is focused.
*/
func (self Instance) GetFocusedItem() int { //gd:PopupMenu.get_focused_item
	return int(int(Advanced(self).GetFocusedItem()))
}

/*
Moves the scroll view to make the item at the given [param index] visible.
*/
func (self Instance) ScrollToItem(index int) { //gd:PopupMenu.scroll_to_item
	Advanced(self).ScrollToItem(int64(index))
}

/*
Removes the item at the given [param index] from the menu.
[b]Note:[/b] The indices of items after the removed item will be shifted by one.
*/
func (self Instance) RemoveItem(index int) { //gd:PopupMenu.remove_item
	Advanced(self).RemoveItem(int64(index))
}

/*
Adds a separator between items. Separators also occupy an index, which you can set by using the [param id] parameter.
A [param label] can optionally be provided, which will appear at the center of the separator.
*/
func (self Instance) AddSeparator() { //gd:PopupMenu.add_separator
	Advanced(self).AddSeparator(String.New(""), int64(-1))
}

/*
Adds a separator between items. Separators also occupy an index, which you can set by using the [param id] parameter.
A [param label] can optionally be provided, which will appear at the center of the separator.
*/
func (self Expanded) AddSeparator(label string, id int) { //gd:PopupMenu.add_separator
	Advanced(self).AddSeparator(String.New(label), int64(id))
}

/*
Removes all items from the [PopupMenu]. If [param free_submenus] is [code]true[/code], the submenu nodes are automatically freed.
*/
func (self Instance) Clear() { //gd:PopupMenu.clear
	Advanced(self).Clear(false)
}

/*
Removes all items from the [PopupMenu]. If [param free_submenus] is [code]true[/code], the submenu nodes are automatically freed.
*/
func (self Expanded) Clear(free_submenus bool) { //gd:PopupMenu.clear
	Advanced(self).Clear(free_submenus)
}

/*
Returns [code]true[/code] if the menu is bound to the special system menu.
*/
func (self Instance) IsSystemMenu() bool { //gd:PopupMenu.is_system_menu
	return bool(Advanced(self).IsSystemMenu())
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.PopupMenu

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("PopupMenu"))
	casted := Instance{*(*gdclass.PopupMenu)(unsafe.Pointer(&object))}
	return casted
}

func (self Instance) HideOnItemSelection() bool {
	return bool(class(self).IsHideOnItemSelection())
}

func (self Instance) SetHideOnItemSelection(value bool) {
	class(self).SetHideOnItemSelection(value)
}

func (self Instance) HideOnCheckableItemSelection() bool {
	return bool(class(self).IsHideOnCheckableItemSelection())
}

func (self Instance) SetHideOnCheckableItemSelection(value bool) {
	class(self).SetHideOnCheckableItemSelection(value)
}

func (self Instance) HideOnStateItemSelection() bool {
	return bool(class(self).IsHideOnStateItemSelection())
}

func (self Instance) SetHideOnStateItemSelection(value bool) {
	class(self).SetHideOnStateItemSelection(value)
}

func (self Instance) SubmenuPopupDelay() Float.X {
	return Float.X(Float.X(class(self).GetSubmenuPopupDelay()))
}

func (self Instance) SetSubmenuPopupDelay(value Float.X) {
	class(self).SetSubmenuPopupDelay(float64(value))
}

func (self Instance) AllowSearch() bool {
	return bool(class(self).GetAllowSearch())
}

func (self Instance) SetAllowSearch(value bool) {
	class(self).SetAllowSearch(value)
}

func (self Instance) SystemMenuId() gdclass.NativeMenuSystemMenus {
	return gdclass.NativeMenuSystemMenus(class(self).GetSystemMenu())
}

func (self Instance) SetSystemMenuId(value gdclass.NativeMenuSystemMenus) {
	class(self).SetSystemMenu(value)
}

func (self Instance) PreferNativeMenu() bool {
	return bool(class(self).IsPreferNativeMenu())
}

func (self Instance) SetPreferNativeMenu(value bool) {
	class(self).SetPreferNativeMenu(value)
}

func (self Instance) ItemCount() int {
	return int(int(class(self).GetItemCount()))
}

func (self Instance) SetItemCount(value int) {
	class(self).SetItemCount(int64(value))
}

/*
Checks the provided [param event] against the [PopupMenu]'s shortcuts and accelerators, and activates the first item with matching events. If [param for_global_only] is [code]true[/code], only shortcuts and accelerators with [code]global[/code] set to [code]true[/code] will be called.
Returns [code]true[/code] if an item was successfully activated.
[b]Note:[/b] Certain [Control]s, such as [MenuButton], will call this method automatically.
*/
//go:nosplit
func (self class) ActivateItemByEvent(event [1]gdclass.InputEvent, for_global_only bool) bool { //gd:PopupMenu.activate_item_by_event
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(event[0])[0])
	callframe.Arg(frame, for_global_only)
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PopupMenu.Bind_activate_item_by_event, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetPreferNativeMenu(enabled bool) { //gd:PopupMenu.set_prefer_native_menu
	var frame = callframe.New()
	callframe.Arg(frame, enabled)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PopupMenu.Bind_set_prefer_native_menu, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) IsPreferNativeMenu() bool { //gd:PopupMenu.is_prefer_native_menu
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PopupMenu.Bind_is_prefer_native_menu, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns [code]true[/code] if the system native menu is supported and currently used by this [PopupMenu].
*/
//go:nosplit
func (self class) IsNativeMenu() bool { //gd:PopupMenu.is_native_menu
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PopupMenu.Bind_is_native_menu, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Adds a new item with text [param label].
An [param id] can optionally be provided, as well as an accelerator ([param accel]). If no [param id] is provided, one will be created from the index. If no [param accel] is provided, then the default value of 0 (corresponding to [constant @GlobalScope.KEY_NONE]) will be assigned to the item (which means it won't have any accelerator). See [method get_item_accelerator] for more info on accelerators.
[b]Note:[/b] The provided [param id] is used only in [signal id_pressed] and [signal id_focused] signals. It's not related to the [code]index[/code] arguments in e.g. [method set_item_checked].
*/
//go:nosplit
func (self class) AddItem(label String.Readable, id int64, accel Key) { //gd:PopupMenu.add_item
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalString(label)))
	callframe.Arg(frame, id)
	callframe.Arg(frame, accel)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PopupMenu.Bind_add_item, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Adds a new item with text [param label] and icon [param texture].
An [param id] can optionally be provided, as well as an accelerator ([param accel]). If no [param id] is provided, one will be created from the index. If no [param accel] is provided, then the default value of 0 (corresponding to [constant @GlobalScope.KEY_NONE]) will be assigned to the item (which means it won't have any accelerator). See [method get_item_accelerator] for more info on accelerators.
*/
//go:nosplit
func (self class) AddIconItem(texture [1]gdclass.Texture2D, label String.Readable, id int64, accel Key) { //gd:PopupMenu.add_icon_item
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(texture[0])[0])
	callframe.Arg(frame, pointers.Get(gd.InternalString(label)))
	callframe.Arg(frame, id)
	callframe.Arg(frame, accel)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PopupMenu.Bind_add_icon_item, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Adds a new checkable item with text [param label].
An [param id] can optionally be provided, as well as an accelerator ([param accel]). If no [param id] is provided, one will be created from the index. If no [param accel] is provided, then the default value of 0 (corresponding to [constant @GlobalScope.KEY_NONE]) will be assigned to the item (which means it won't have any accelerator). See [method get_item_accelerator] for more info on accelerators.
[b]Note:[/b] Checkable items just display a checkmark, but don't have any built-in checking behavior and must be checked/unchecked manually. See [method set_item_checked] for more info on how to control it.
*/
//go:nosplit
func (self class) AddCheckItem(label String.Readable, id int64, accel Key) { //gd:PopupMenu.add_check_item
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalString(label)))
	callframe.Arg(frame, id)
	callframe.Arg(frame, accel)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PopupMenu.Bind_add_check_item, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Adds a new checkable item with text [param label] and icon [param texture].
An [param id] can optionally be provided, as well as an accelerator ([param accel]). If no [param id] is provided, one will be created from the index. If no [param accel] is provided, then the default value of 0 (corresponding to [constant @GlobalScope.KEY_NONE]) will be assigned to the item (which means it won't have any accelerator). See [method get_item_accelerator] for more info on accelerators.
[b]Note:[/b] Checkable items just display a checkmark, but don't have any built-in checking behavior and must be checked/unchecked manually. See [method set_item_checked] for more info on how to control it.
*/
//go:nosplit
func (self class) AddIconCheckItem(texture [1]gdclass.Texture2D, label String.Readable, id int64, accel Key) { //gd:PopupMenu.add_icon_check_item
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(texture[0])[0])
	callframe.Arg(frame, pointers.Get(gd.InternalString(label)))
	callframe.Arg(frame, id)
	callframe.Arg(frame, accel)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PopupMenu.Bind_add_icon_check_item, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Adds a new radio check button with text [param label].
An [param id] can optionally be provided, as well as an accelerator ([param accel]). If no [param id] is provided, one will be created from the index. If no [param accel] is provided, then the default value of 0 (corresponding to [constant @GlobalScope.KEY_NONE]) will be assigned to the item (which means it won't have any accelerator). See [method get_item_accelerator] for more info on accelerators.
[b]Note:[/b] Checkable items just display a checkmark, but don't have any built-in checking behavior and must be checked/unchecked manually. See [method set_item_checked] for more info on how to control it.
*/
//go:nosplit
func (self class) AddRadioCheckItem(label String.Readable, id int64, accel Key) { //gd:PopupMenu.add_radio_check_item
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalString(label)))
	callframe.Arg(frame, id)
	callframe.Arg(frame, accel)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PopupMenu.Bind_add_radio_check_item, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Same as [method add_icon_check_item], but uses a radio check button.
*/
//go:nosplit
func (self class) AddIconRadioCheckItem(texture [1]gdclass.Texture2D, label String.Readable, id int64, accel Key) { //gd:PopupMenu.add_icon_radio_check_item
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(texture[0])[0])
	callframe.Arg(frame, pointers.Get(gd.InternalString(label)))
	callframe.Arg(frame, id)
	callframe.Arg(frame, accel)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PopupMenu.Bind_add_icon_radio_check_item, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Adds a new multistate item with text [param label].
Contrarily to normal binary items, multistate items can have more than two states, as defined by [param max_states]. The default value is defined by [param default_state].
An [param id] can optionally be provided, as well as an accelerator ([param accel]). If no [param id] is provided, one will be created from the index. If no [param accel] is provided, then the default value of 0 (corresponding to [constant @GlobalScope.KEY_NONE]) will be assigned to the item (which means it won't have any accelerator). See [method get_item_accelerator] for more info on accelerators.
[codeblock]
func _ready():
    add_multistate_item("Item", 3, 0)

    index_pressed.connect(func(index: int):
            toggle_item_multistate(index)
            match get_item_multistate(index):
                0:
                    print("First state")
                1:
                    print("Second state")
                2:
                    print("Third state")
        )
[/codeblock]
[b]Note:[/b] Multistate items don't update their state automatically and must be done manually. See [method toggle_item_multistate], [method set_item_multistate] and [method get_item_multistate] for more info on how to control it.
*/
//go:nosplit
func (self class) AddMultistateItem(label String.Readable, max_states int64, default_state int64, id int64, accel Key) { //gd:PopupMenu.add_multistate_item
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalString(label)))
	callframe.Arg(frame, max_states)
	callframe.Arg(frame, default_state)
	callframe.Arg(frame, id)
	callframe.Arg(frame, accel)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PopupMenu.Bind_add_multistate_item, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Adds a [Shortcut].
An [param id] can optionally be provided. If no [param id] is provided, one will be created from the index.
If [param allow_echo] is [code]true[/code], the shortcut can be activated with echo events.
*/
//go:nosplit
func (self class) AddShortcut(shortcut [1]gdclass.Shortcut, id int64, global bool, allow_echo bool) { //gd:PopupMenu.add_shortcut
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(shortcut[0])[0])
	callframe.Arg(frame, id)
	callframe.Arg(frame, global)
	callframe.Arg(frame, allow_echo)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PopupMenu.Bind_add_shortcut, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Adds a new item and assigns the specified [Shortcut] and icon [param texture] to it. Sets the label of the checkbox to the [Shortcut]'s name.
An [param id] can optionally be provided. If no [param id] is provided, one will be created from the index.
If [param allow_echo] is [code]true[/code], the shortcut can be activated with echo events.
*/
//go:nosplit
func (self class) AddIconShortcut(texture [1]gdclass.Texture2D, shortcut [1]gdclass.Shortcut, id int64, global bool, allow_echo bool) { //gd:PopupMenu.add_icon_shortcut
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(texture[0])[0])
	callframe.Arg(frame, pointers.Get(shortcut[0])[0])
	callframe.Arg(frame, id)
	callframe.Arg(frame, global)
	callframe.Arg(frame, allow_echo)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PopupMenu.Bind_add_icon_shortcut, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Adds a new checkable item and assigns the specified [Shortcut] to it. Sets the label of the checkbox to the [Shortcut]'s name.
An [param id] can optionally be provided. If no [param id] is provided, one will be created from the index.
[b]Note:[/b] Checkable items just display a checkmark, but don't have any built-in checking behavior and must be checked/unchecked manually. See [method set_item_checked] for more info on how to control it.
*/
//go:nosplit
func (self class) AddCheckShortcut(shortcut [1]gdclass.Shortcut, id int64, global bool) { //gd:PopupMenu.add_check_shortcut
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(shortcut[0])[0])
	callframe.Arg(frame, id)
	callframe.Arg(frame, global)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PopupMenu.Bind_add_check_shortcut, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Adds a new checkable item and assigns the specified [Shortcut] and icon [param texture] to it. Sets the label of the checkbox to the [Shortcut]'s name.
An [param id] can optionally be provided. If no [param id] is provided, one will be created from the index.
[b]Note:[/b] Checkable items just display a checkmark, but don't have any built-in checking behavior and must be checked/unchecked manually. See [method set_item_checked] for more info on how to control it.
*/
//go:nosplit
func (self class) AddIconCheckShortcut(texture [1]gdclass.Texture2D, shortcut [1]gdclass.Shortcut, id int64, global bool) { //gd:PopupMenu.add_icon_check_shortcut
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(texture[0])[0])
	callframe.Arg(frame, pointers.Get(shortcut[0])[0])
	callframe.Arg(frame, id)
	callframe.Arg(frame, global)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PopupMenu.Bind_add_icon_check_shortcut, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Adds a new radio check button and assigns a [Shortcut] to it. Sets the label of the checkbox to the [Shortcut]'s name.
An [param id] can optionally be provided. If no [param id] is provided, one will be created from the index.
[b]Note:[/b] Checkable items just display a checkmark, but don't have any built-in checking behavior and must be checked/unchecked manually. See [method set_item_checked] for more info on how to control it.
*/
//go:nosplit
func (self class) AddRadioCheckShortcut(shortcut [1]gdclass.Shortcut, id int64, global bool) { //gd:PopupMenu.add_radio_check_shortcut
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(shortcut[0])[0])
	callframe.Arg(frame, id)
	callframe.Arg(frame, global)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PopupMenu.Bind_add_radio_check_shortcut, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Same as [method add_icon_check_shortcut], but uses a radio check button.
*/
//go:nosplit
func (self class) AddIconRadioCheckShortcut(texture [1]gdclass.Texture2D, shortcut [1]gdclass.Shortcut, id int64, global bool) { //gd:PopupMenu.add_icon_radio_check_shortcut
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(texture[0])[0])
	callframe.Arg(frame, pointers.Get(shortcut[0])[0])
	callframe.Arg(frame, id)
	callframe.Arg(frame, global)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PopupMenu.Bind_add_icon_radio_check_shortcut, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Adds an item that will act as a submenu of the parent [PopupMenu] node when clicked. The [param submenu] argument must be the name of an existing [PopupMenu] that has been added as a child to this node. This submenu will be shown when the item is clicked, hovered for long enough, or activated using the [code]ui_select[/code] or [code]ui_right[/code] input actions.
An [param id] can optionally be provided. If no [param id] is provided, one will be created from the index.
*/
//go:nosplit
func (self class) AddSubmenuItem(label String.Readable, submenu String.Readable, id int64) { //gd:PopupMenu.add_submenu_item
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalString(label)))
	callframe.Arg(frame, pointers.Get(gd.InternalString(submenu)))
	callframe.Arg(frame, id)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PopupMenu.Bind_add_submenu_item, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Adds an item that will act as a submenu of the parent [PopupMenu] node when clicked. This submenu will be shown when the item is clicked, hovered for long enough, or activated using the [code]ui_select[/code] or [code]ui_right[/code] input actions.
[param submenu] must be either child of this [PopupMenu] or has no parent node (in which case it will be automatically added as a child). If the [param submenu] popup has another parent, this method will fail.
An [param id] can optionally be provided. If no [param id] is provided, one will be created from the index.
*/
//go:nosplit
func (self class) AddSubmenuNodeItem(label String.Readable, submenu [1]gdclass.PopupMenu, id int64) { //gd:PopupMenu.add_submenu_node_item
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalString(label)))
	callframe.Arg(frame, gd.PointerWithOwnershipTransferredToGodot(submenu[0].AsObject()[0]))
	callframe.Arg(frame, id)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PopupMenu.Bind_add_submenu_node_item, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Sets the text of the item at the given [param index].
*/
//go:nosplit
func (self class) SetItemText(index int64, text String.Readable) { //gd:PopupMenu.set_item_text
	var frame = callframe.New()
	callframe.Arg(frame, index)
	callframe.Arg(frame, pointers.Get(gd.InternalString(text)))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PopupMenu.Bind_set_item_text, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Sets item's text base writing direction.
*/
//go:nosplit
func (self class) SetItemTextDirection(index int64, direction gdclass.ControlTextDirection) { //gd:PopupMenu.set_item_text_direction
	var frame = callframe.New()
	callframe.Arg(frame, index)
	callframe.Arg(frame, direction)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PopupMenu.Bind_set_item_text_direction, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Sets language code of item's text used for line-breaking and text shaping algorithms, if left empty current locale is used instead.
*/
//go:nosplit
func (self class) SetItemLanguage(index int64, language String.Readable) { //gd:PopupMenu.set_item_language
	var frame = callframe.New()
	callframe.Arg(frame, index)
	callframe.Arg(frame, pointers.Get(gd.InternalString(language)))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PopupMenu.Bind_set_item_language, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Replaces the [Texture2D] icon of the item at the given [param index].
*/
//go:nosplit
func (self class) SetItemIcon(index int64, icon [1]gdclass.Texture2D) { //gd:PopupMenu.set_item_icon
	var frame = callframe.New()
	callframe.Arg(frame, index)
	callframe.Arg(frame, pointers.Get(icon[0])[0])
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PopupMenu.Bind_set_item_icon, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Sets the maximum allowed width of the icon for the item at the given [param index]. This limit is applied on top of the default size of the icon and on top of [theme_item icon_max_width]. The height is adjusted according to the icon's ratio.
*/
//go:nosplit
func (self class) SetItemIconMaxWidth(index int64, width int64) { //gd:PopupMenu.set_item_icon_max_width
	var frame = callframe.New()
	callframe.Arg(frame, index)
	callframe.Arg(frame, width)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PopupMenu.Bind_set_item_icon_max_width, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Sets a modulating [Color] of the item's icon at the given [param index].
*/
//go:nosplit
func (self class) SetItemIconModulate(index int64, modulate Color.RGBA) { //gd:PopupMenu.set_item_icon_modulate
	var frame = callframe.New()
	callframe.Arg(frame, index)
	callframe.Arg(frame, modulate)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PopupMenu.Bind_set_item_icon_modulate, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Sets the checkstate status of the item at the given [param index].
*/
//go:nosplit
func (self class) SetItemChecked(index int64, checked bool) { //gd:PopupMenu.set_item_checked
	var frame = callframe.New()
	callframe.Arg(frame, index)
	callframe.Arg(frame, checked)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PopupMenu.Bind_set_item_checked, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Sets the [param id] of the item at the given [param index].
The [param id] is used in [signal id_pressed] and [signal id_focused] signals.
*/
//go:nosplit
func (self class) SetItemId(index int64, id int64) { //gd:PopupMenu.set_item_id
	var frame = callframe.New()
	callframe.Arg(frame, index)
	callframe.Arg(frame, id)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PopupMenu.Bind_set_item_id, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Sets the accelerator of the item at the given [param index]. An accelerator is a keyboard shortcut that can be pressed to trigger the menu button even if it's not currently open. [param accel] is generally a combination of [enum KeyModifierMask]s and [enum Key]s using bitwise OR such as [code]KEY_MASK_CTRL | KEY_A[/code] ([kbd]Ctrl + A[/kbd]).
*/
//go:nosplit
func (self class) SetItemAccelerator(index int64, accel Key) { //gd:PopupMenu.set_item_accelerator
	var frame = callframe.New()
	callframe.Arg(frame, index)
	callframe.Arg(frame, accel)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PopupMenu.Bind_set_item_accelerator, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Sets the metadata of an item, which may be of any type. You can later get it with [method get_item_metadata], which provides a simple way of assigning context data to items.
*/
//go:nosplit
func (self class) SetItemMetadata(index int64, metadata variant.Any) { //gd:PopupMenu.set_item_metadata
	var frame = callframe.New()
	callframe.Arg(frame, index)
	callframe.Arg(frame, pointers.Get(gd.InternalVariant(metadata)))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PopupMenu.Bind_set_item_metadata, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Enables/disables the item at the given [param index]. When it is disabled, it can't be selected and its action can't be invoked.
*/
//go:nosplit
func (self class) SetItemDisabled(index int64, disabled bool) { //gd:PopupMenu.set_item_disabled
	var frame = callframe.New()
	callframe.Arg(frame, index)
	callframe.Arg(frame, disabled)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PopupMenu.Bind_set_item_disabled, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Sets the submenu of the item at the given [param index]. The submenu is the name of a child [PopupMenu] node that would be shown when the item is clicked.
*/
//go:nosplit
func (self class) SetItemSubmenu(index int64, submenu String.Readable) { //gd:PopupMenu.set_item_submenu
	var frame = callframe.New()
	callframe.Arg(frame, index)
	callframe.Arg(frame, pointers.Get(gd.InternalString(submenu)))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PopupMenu.Bind_set_item_submenu, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Sets the submenu of the item at the given [param index]. The submenu is a [PopupMenu] node that would be shown when the item is clicked. It must either be a child of this [PopupMenu] or has no parent (in which case it will be automatically added as a child). If the [param submenu] popup has another parent, this method will fail.
*/
//go:nosplit
func (self class) SetItemSubmenuNode(index int64, submenu [1]gdclass.PopupMenu) { //gd:PopupMenu.set_item_submenu_node
	var frame = callframe.New()
	callframe.Arg(frame, index)
	callframe.Arg(frame, gd.PointerWithOwnershipTransferredToGodot(submenu[0].AsObject()[0]))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PopupMenu.Bind_set_item_submenu_node, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Mark the item at the given [param index] as a separator, which means that it would be displayed as a line. If [code]false[/code], sets the type of the item to plain text.
*/
//go:nosplit
func (self class) SetItemAsSeparator(index int64, enable bool) { //gd:PopupMenu.set_item_as_separator
	var frame = callframe.New()
	callframe.Arg(frame, index)
	callframe.Arg(frame, enable)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PopupMenu.Bind_set_item_as_separator, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Sets whether the item at the given [param index] has a checkbox. If [code]false[/code], sets the type of the item to plain text.
[b]Note:[/b] Checkable items just display a checkmark, but don't have any built-in checking behavior and must be checked/unchecked manually.
*/
//go:nosplit
func (self class) SetItemAsCheckable(index int64, enable bool) { //gd:PopupMenu.set_item_as_checkable
	var frame = callframe.New()
	callframe.Arg(frame, index)
	callframe.Arg(frame, enable)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PopupMenu.Bind_set_item_as_checkable, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Sets the type of the item at the given [param index] to radio button. If [code]false[/code], sets the type of the item to plain text.
*/
//go:nosplit
func (self class) SetItemAsRadioCheckable(index int64, enable bool) { //gd:PopupMenu.set_item_as_radio_checkable
	var frame = callframe.New()
	callframe.Arg(frame, index)
	callframe.Arg(frame, enable)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PopupMenu.Bind_set_item_as_radio_checkable, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Sets the [String] tooltip of the item at the given [param index].
*/
//go:nosplit
func (self class) SetItemTooltip(index int64, tooltip String.Readable) { //gd:PopupMenu.set_item_tooltip
	var frame = callframe.New()
	callframe.Arg(frame, index)
	callframe.Arg(frame, pointers.Get(gd.InternalString(tooltip)))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PopupMenu.Bind_set_item_tooltip, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Sets a [Shortcut] for the item at the given [param index].
*/
//go:nosplit
func (self class) SetItemShortcut(index int64, shortcut [1]gdclass.Shortcut, global bool) { //gd:PopupMenu.set_item_shortcut
	var frame = callframe.New()
	callframe.Arg(frame, index)
	callframe.Arg(frame, pointers.Get(shortcut[0])[0])
	callframe.Arg(frame, global)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PopupMenu.Bind_set_item_shortcut, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Sets the horizontal offset of the item at the given [param index].
*/
//go:nosplit
func (self class) SetItemIndent(index int64, indent int64) { //gd:PopupMenu.set_item_indent
	var frame = callframe.New()
	callframe.Arg(frame, index)
	callframe.Arg(frame, indent)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PopupMenu.Bind_set_item_indent, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Sets the state of a multistate item. See [method add_multistate_item] for details.
*/
//go:nosplit
func (self class) SetItemMultistate(index int64, state int64) { //gd:PopupMenu.set_item_multistate
	var frame = callframe.New()
	callframe.Arg(frame, index)
	callframe.Arg(frame, state)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PopupMenu.Bind_set_item_multistate, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Sets the max states of a multistate item. See [method add_multistate_item] for details.
*/
//go:nosplit
func (self class) SetItemMultistateMax(index int64, max_states int64) { //gd:PopupMenu.set_item_multistate_max
	var frame = callframe.New()
	callframe.Arg(frame, index)
	callframe.Arg(frame, max_states)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PopupMenu.Bind_set_item_multistate_max, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Disables the [Shortcut] of the item at the given [param index].
*/
//go:nosplit
func (self class) SetItemShortcutDisabled(index int64, disabled bool) { //gd:PopupMenu.set_item_shortcut_disabled
	var frame = callframe.New()
	callframe.Arg(frame, index)
	callframe.Arg(frame, disabled)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PopupMenu.Bind_set_item_shortcut_disabled, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Toggles the check state of the item at the given [param index].
*/
//go:nosplit
func (self class) ToggleItemChecked(index int64) { //gd:PopupMenu.toggle_item_checked
	var frame = callframe.New()
	callframe.Arg(frame, index)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PopupMenu.Bind_toggle_item_checked, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Cycle to the next state of a multistate item. See [method add_multistate_item] for details.
*/
//go:nosplit
func (self class) ToggleItemMultistate(index int64) { //gd:PopupMenu.toggle_item_multistate
	var frame = callframe.New()
	callframe.Arg(frame, index)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PopupMenu.Bind_toggle_item_multistate, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns the text of the item at the given [param index].
*/
//go:nosplit
func (self class) GetItemText(index int64) String.Readable { //gd:PopupMenu.get_item_text
	var frame = callframe.New()
	callframe.Arg(frame, index)
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PopupMenu.Bind_get_item_text, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = String.Via(gd.StringProxy{}, pointers.Pack(pointers.New[gd.String](r_ret.Get())))
	frame.Free()
	return ret
}

/*
Returns item's text base writing direction.
*/
//go:nosplit
func (self class) GetItemTextDirection(index int64) gdclass.ControlTextDirection { //gd:PopupMenu.get_item_text_direction
	var frame = callframe.New()
	callframe.Arg(frame, index)
	var r_ret = callframe.Ret[gdclass.ControlTextDirection](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PopupMenu.Bind_get_item_text_direction, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns item's text language code.
*/
//go:nosplit
func (self class) GetItemLanguage(index int64) String.Readable { //gd:PopupMenu.get_item_language
	var frame = callframe.New()
	callframe.Arg(frame, index)
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PopupMenu.Bind_get_item_language, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = String.Via(gd.StringProxy{}, pointers.Pack(pointers.New[gd.String](r_ret.Get())))
	frame.Free()
	return ret
}

/*
Returns the icon of the item at the given [param index].
*/
//go:nosplit
func (self class) GetItemIcon(index int64) [1]gdclass.Texture2D { //gd:PopupMenu.get_item_icon
	var frame = callframe.New()
	callframe.Arg(frame, index)
	var r_ret = callframe.Ret[gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PopupMenu.Bind_get_item_icon, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = [1]gdclass.Texture2D{gd.PointerWithOwnershipTransferredToGo[gdclass.Texture2D](r_ret.Get())}
	frame.Free()
	return ret
}

/*
Returns the maximum allowed width of the icon for the item at the given [param index].
*/
//go:nosplit
func (self class) GetItemIconMaxWidth(index int64) int64 { //gd:PopupMenu.get_item_icon_max_width
	var frame = callframe.New()
	callframe.Arg(frame, index)
	var r_ret = callframe.Ret[int64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PopupMenu.Bind_get_item_icon_max_width, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns a [Color] modulating the item's icon at the given [param index].
*/
//go:nosplit
func (self class) GetItemIconModulate(index int64) Color.RGBA { //gd:PopupMenu.get_item_icon_modulate
	var frame = callframe.New()
	callframe.Arg(frame, index)
	var r_ret = callframe.Ret[Color.RGBA](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PopupMenu.Bind_get_item_icon_modulate, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns [code]true[/code] if the item at the given [param index] is checked.
*/
//go:nosplit
func (self class) IsItemChecked(index int64) bool { //gd:PopupMenu.is_item_checked
	var frame = callframe.New()
	callframe.Arg(frame, index)
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PopupMenu.Bind_is_item_checked, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the ID of the item at the given [param index]. [code]id[/code] can be manually assigned, while index can not.
*/
//go:nosplit
func (self class) GetItemId(index int64) int64 { //gd:PopupMenu.get_item_id
	var frame = callframe.New()
	callframe.Arg(frame, index)
	var r_ret = callframe.Ret[int64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PopupMenu.Bind_get_item_id, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the index of the item containing the specified [param id]. Index is automatically assigned to each item by the engine and can not be set manually.
*/
//go:nosplit
func (self class) GetItemIndex(id int64) int64 { //gd:PopupMenu.get_item_index
	var frame = callframe.New()
	callframe.Arg(frame, id)
	var r_ret = callframe.Ret[int64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PopupMenu.Bind_get_item_index, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the accelerator of the item at the given [param index]. An accelerator is a keyboard shortcut that can be pressed to trigger the menu button even if it's not currently open. The return value is an integer which is generally a combination of [enum KeyModifierMask]s and [enum Key]s using bitwise OR such as [code]KEY_MASK_CTRL | KEY_A[/code] ([kbd]Ctrl + A[/kbd]). If no accelerator is defined for the specified [param index], [method get_item_accelerator] returns [code]0[/code] (corresponding to [constant @GlobalScope.KEY_NONE]).
*/
//go:nosplit
func (self class) GetItemAccelerator(index int64) Key { //gd:PopupMenu.get_item_accelerator
	var frame = callframe.New()
	callframe.Arg(frame, index)
	var r_ret = callframe.Ret[Key](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PopupMenu.Bind_get_item_accelerator, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the metadata of the specified item, which might be of any type. You can set it with [method set_item_metadata], which provides a simple way of assigning context data to items.
*/
//go:nosplit
func (self class) GetItemMetadata(index int64) variant.Any { //gd:PopupMenu.get_item_metadata
	var frame = callframe.New()
	callframe.Arg(frame, index)
	var r_ret = callframe.Ret[[3]uint64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PopupMenu.Bind_get_item_metadata, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = variant.Implementation(gd.VariantProxy{}, pointers.Pack(pointers.New[gd.Variant](r_ret.Get())))
	frame.Free()
	return ret
}

/*
Returns [code]true[/code] if the item at the given [param index] is disabled. When it is disabled it can't be selected, or its action invoked.
See [method set_item_disabled] for more info on how to disable an item.
*/
//go:nosplit
func (self class) IsItemDisabled(index int64) bool { //gd:PopupMenu.is_item_disabled
	var frame = callframe.New()
	callframe.Arg(frame, index)
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PopupMenu.Bind_is_item_disabled, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the submenu name of the item at the given [param index]. See [method add_submenu_item] for more info on how to add a submenu.
*/
//go:nosplit
func (self class) GetItemSubmenu(index int64) String.Readable { //gd:PopupMenu.get_item_submenu
	var frame = callframe.New()
	callframe.Arg(frame, index)
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PopupMenu.Bind_get_item_submenu, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = String.Via(gd.StringProxy{}, pointers.Pack(pointers.New[gd.String](r_ret.Get())))
	frame.Free()
	return ret
}

/*
Returns the submenu of the item at the given [param index], or [code]null[/code] if no submenu was added. See [method add_submenu_node_item] for more info on how to add a submenu.
*/
//go:nosplit
func (self class) GetItemSubmenuNode(index int64) [1]gdclass.PopupMenu { //gd:PopupMenu.get_item_submenu_node
	var frame = callframe.New()
	callframe.Arg(frame, index)
	var r_ret = callframe.Ret[gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PopupMenu.Bind_get_item_submenu_node, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = [1]gdclass.PopupMenu{gd.PointerLifetimeBoundTo[gdclass.PopupMenu](self.AsObject(), r_ret.Get())}
	frame.Free()
	return ret
}

/*
Returns [code]true[/code] if the item is a separator. If it is, it will be displayed as a line. See [method add_separator] for more info on how to add a separator.
*/
//go:nosplit
func (self class) IsItemSeparator(index int64) bool { //gd:PopupMenu.is_item_separator
	var frame = callframe.New()
	callframe.Arg(frame, index)
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PopupMenu.Bind_is_item_separator, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns [code]true[/code] if the item at the given [param index] is checkable in some way, i.e. if it has a checkbox or radio button.
[b]Note:[/b] Checkable items just display a checkmark or radio button, but don't have any built-in checking behavior and must be checked/unchecked manually.
*/
//go:nosplit
func (self class) IsItemCheckable(index int64) bool { //gd:PopupMenu.is_item_checkable
	var frame = callframe.New()
	callframe.Arg(frame, index)
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PopupMenu.Bind_is_item_checkable, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns [code]true[/code] if the item at the given [param index] has radio button-style checkability.
[b]Note:[/b] This is purely cosmetic; you must add the logic for checking/unchecking items in radio groups.
*/
//go:nosplit
func (self class) IsItemRadioCheckable(index int64) bool { //gd:PopupMenu.is_item_radio_checkable
	var frame = callframe.New()
	callframe.Arg(frame, index)
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PopupMenu.Bind_is_item_radio_checkable, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns [code]true[/code] if the specified item's shortcut is disabled.
*/
//go:nosplit
func (self class) IsItemShortcutDisabled(index int64) bool { //gd:PopupMenu.is_item_shortcut_disabled
	var frame = callframe.New()
	callframe.Arg(frame, index)
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PopupMenu.Bind_is_item_shortcut_disabled, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the tooltip associated with the item at the given [param index].
*/
//go:nosplit
func (self class) GetItemTooltip(index int64) String.Readable { //gd:PopupMenu.get_item_tooltip
	var frame = callframe.New()
	callframe.Arg(frame, index)
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PopupMenu.Bind_get_item_tooltip, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = String.Via(gd.StringProxy{}, pointers.Pack(pointers.New[gd.String](r_ret.Get())))
	frame.Free()
	return ret
}

/*
Returns the [Shortcut] associated with the item at the given [param index].
*/
//go:nosplit
func (self class) GetItemShortcut(index int64) [1]gdclass.Shortcut { //gd:PopupMenu.get_item_shortcut
	var frame = callframe.New()
	callframe.Arg(frame, index)
	var r_ret = callframe.Ret[gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PopupMenu.Bind_get_item_shortcut, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = [1]gdclass.Shortcut{gd.PointerWithOwnershipTransferredToGo[gdclass.Shortcut](r_ret.Get())}
	frame.Free()
	return ret
}

/*
Returns the horizontal offset of the item at the given [param index].
*/
//go:nosplit
func (self class) GetItemIndent(index int64) int64 { //gd:PopupMenu.get_item_indent
	var frame = callframe.New()
	callframe.Arg(frame, index)
	var r_ret = callframe.Ret[int64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PopupMenu.Bind_get_item_indent, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the max states of the item at the given [param index].
*/
//go:nosplit
func (self class) GetItemMultistateMax(index int64) int64 { //gd:PopupMenu.get_item_multistate_max
	var frame = callframe.New()
	callframe.Arg(frame, index)
	var r_ret = callframe.Ret[int64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PopupMenu.Bind_get_item_multistate_max, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the state of the item at the given [param index].
*/
//go:nosplit
func (self class) GetItemMultistate(index int64) int64 { //gd:PopupMenu.get_item_multistate
	var frame = callframe.New()
	callframe.Arg(frame, index)
	var r_ret = callframe.Ret[int64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PopupMenu.Bind_get_item_multistate, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Sets the currently focused item as the given [param index].
Passing [code]-1[/code] as the index makes so that no item is focused.
*/
//go:nosplit
func (self class) SetFocusedItem(index int64) { //gd:PopupMenu.set_focused_item
	var frame = callframe.New()
	callframe.Arg(frame, index)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PopupMenu.Bind_set_focused_item, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns the index of the currently focused item. Returns [code]-1[/code] if no item is focused.
*/
//go:nosplit
func (self class) GetFocusedItem() int64 { //gd:PopupMenu.get_focused_item
	var frame = callframe.New()
	var r_ret = callframe.Ret[int64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PopupMenu.Bind_get_focused_item, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetItemCount(count int64) { //gd:PopupMenu.set_item_count
	var frame = callframe.New()
	callframe.Arg(frame, count)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PopupMenu.Bind_set_item_count, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetItemCount() int64 { //gd:PopupMenu.get_item_count
	var frame = callframe.New()
	var r_ret = callframe.Ret[int64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PopupMenu.Bind_get_item_count, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Moves the scroll view to make the item at the given [param index] visible.
*/
//go:nosplit
func (self class) ScrollToItem(index int64) { //gd:PopupMenu.scroll_to_item
	var frame = callframe.New()
	callframe.Arg(frame, index)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PopupMenu.Bind_scroll_to_item, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Removes the item at the given [param index] from the menu.
[b]Note:[/b] The indices of items after the removed item will be shifted by one.
*/
//go:nosplit
func (self class) RemoveItem(index int64) { //gd:PopupMenu.remove_item
	var frame = callframe.New()
	callframe.Arg(frame, index)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PopupMenu.Bind_remove_item, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Adds a separator between items. Separators also occupy an index, which you can set by using the [param id] parameter.
A [param label] can optionally be provided, which will appear at the center of the separator.
*/
//go:nosplit
func (self class) AddSeparator(label String.Readable, id int64) { //gd:PopupMenu.add_separator
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalString(label)))
	callframe.Arg(frame, id)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PopupMenu.Bind_add_separator, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Removes all items from the [PopupMenu]. If [param free_submenus] is [code]true[/code], the submenu nodes are automatically freed.
*/
//go:nosplit
func (self class) Clear(free_submenus bool) { //gd:PopupMenu.clear
	var frame = callframe.New()
	callframe.Arg(frame, free_submenus)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PopupMenu.Bind_clear, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) SetHideOnItemSelection(enable bool) { //gd:PopupMenu.set_hide_on_item_selection
	var frame = callframe.New()
	callframe.Arg(frame, enable)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PopupMenu.Bind_set_hide_on_item_selection, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) IsHideOnItemSelection() bool { //gd:PopupMenu.is_hide_on_item_selection
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PopupMenu.Bind_is_hide_on_item_selection, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetHideOnCheckableItemSelection(enable bool) { //gd:PopupMenu.set_hide_on_checkable_item_selection
	var frame = callframe.New()
	callframe.Arg(frame, enable)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PopupMenu.Bind_set_hide_on_checkable_item_selection, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) IsHideOnCheckableItemSelection() bool { //gd:PopupMenu.is_hide_on_checkable_item_selection
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PopupMenu.Bind_is_hide_on_checkable_item_selection, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetHideOnStateItemSelection(enable bool) { //gd:PopupMenu.set_hide_on_state_item_selection
	var frame = callframe.New()
	callframe.Arg(frame, enable)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PopupMenu.Bind_set_hide_on_state_item_selection, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) IsHideOnStateItemSelection() bool { //gd:PopupMenu.is_hide_on_state_item_selection
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PopupMenu.Bind_is_hide_on_state_item_selection, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetSubmenuPopupDelay(seconds float64) { //gd:PopupMenu.set_submenu_popup_delay
	var frame = callframe.New()
	callframe.Arg(frame, seconds)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PopupMenu.Bind_set_submenu_popup_delay, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetSubmenuPopupDelay() float64 { //gd:PopupMenu.get_submenu_popup_delay
	var frame = callframe.New()
	var r_ret = callframe.Ret[float64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PopupMenu.Bind_get_submenu_popup_delay, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetAllowSearch(allow bool) { //gd:PopupMenu.set_allow_search
	var frame = callframe.New()
	callframe.Arg(frame, allow)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PopupMenu.Bind_set_allow_search, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetAllowSearch() bool { //gd:PopupMenu.get_allow_search
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PopupMenu.Bind_get_allow_search, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns [code]true[/code] if the menu is bound to the special system menu.
*/
//go:nosplit
func (self class) IsSystemMenu() bool { //gd:PopupMenu.is_system_menu
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PopupMenu.Bind_is_system_menu, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetSystemMenu(system_menu_id gdclass.NativeMenuSystemMenus) { //gd:PopupMenu.set_system_menu
	var frame = callframe.New()
	callframe.Arg(frame, system_menu_id)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PopupMenu.Bind_set_system_menu, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetSystemMenu() gdclass.NativeMenuSystemMenus { //gd:PopupMenu.get_system_menu
	var frame = callframe.New()
	var r_ret = callframe.Ret[gdclass.NativeMenuSystemMenus](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PopupMenu.Bind_get_system_menu, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self Instance) OnIdPressed(cb func(id int)) {
	self[0].AsObject()[0].Connect(gd.NewStringName("id_pressed"), gd.NewCallable(cb), 0)
}

func (self Instance) OnIdFocused(cb func(id int)) {
	self[0].AsObject()[0].Connect(gd.NewStringName("id_focused"), gd.NewCallable(cb), 0)
}

func (self Instance) OnIndexPressed(cb func(index int)) {
	self[0].AsObject()[0].Connect(gd.NewStringName("index_pressed"), gd.NewCallable(cb), 0)
}

func (self Instance) OnMenuChanged(cb func()) {
	self[0].AsObject()[0].Connect(gd.NewStringName("menu_changed"), gd.NewCallable(cb), 0)
}

func (self class) AsPopupMenu() Advanced        { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsPopupMenu() Instance     { return *((*Instance)(unsafe.Pointer(&self))) }
func (self class) AsPopup() Popup.Advanced      { return *((*Popup.Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsPopup() Popup.Instance   { return *((*Popup.Instance)(unsafe.Pointer(&self))) }
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
		return gd.VirtualByName(Popup.Advanced(self.AsPopup()), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(Popup.Instance(self.AsPopup()), name)
	}
}
func init() {
	gdclass.Register("PopupMenu", func(ptr gd.Object) any { return [1]gdclass.PopupMenu{*(*gdclass.PopupMenu)(unsafe.Pointer(&ptr))} })
}

type Key int

const (
	/*Enum value which doesn't correspond to any key. This is used to initialize [enum Key] properties with a generic state.*/
	KeyNone Key = 0
	/*Keycodes with this bit applied are non-printable.*/
	KeySpecial Key = 4194304
	/*Escape key.*/
	KeyEscape Key = 4194305
	/*Tab key.*/
	KeyTab Key = 4194306
	/*Shift + Tab key.*/
	KeyBacktab Key = 4194307
	/*Backspace key.*/
	KeyBackspace Key = 4194308
	/*Return key (on the main keyboard).*/
	KeyEnter Key = 4194309
	/*Enter key on the numeric keypad.*/
	KeyKpEnter Key = 4194310
	/*Insert key.*/
	KeyInsert Key = 4194311
	/*Delete key.*/
	KeyDelete Key = 4194312
	/*Pause key.*/
	KeyPause Key = 4194313
	/*Print Screen key.*/
	KeyPrint Key = 4194314
	/*System Request key.*/
	KeySysreq Key = 4194315
	/*Clear key.*/
	KeyClear Key = 4194316
	/*Home key.*/
	KeyHome Key = 4194317
	/*End key.*/
	KeyEnd Key = 4194318
	/*Left arrow key.*/
	KeyLeft Key = 4194319
	/*Up arrow key.*/
	KeyUp Key = 4194320
	/*Right arrow key.*/
	KeyRight Key = 4194321
	/*Down arrow key.*/
	KeyDown Key = 4194322
	/*Page Up key.*/
	KeyPageup Key = 4194323
	/*Page Down key.*/
	KeyPagedown Key = 4194324
	/*Shift key.*/
	KeyShift Key = 4194325
	/*Control key.*/
	KeyCtrl Key = 4194326
	/*Meta key.*/
	KeyMeta Key = 4194327
	/*Alt key.*/
	KeyAlt Key = 4194328
	/*Caps Lock key.*/
	KeyCapslock Key = 4194329
	/*Num Lock key.*/
	KeyNumlock Key = 4194330
	/*Scroll Lock key.*/
	KeyScrolllock Key = 4194331
	/*F1 key.*/
	KeyF1 Key = 4194332
	/*F2 key.*/
	KeyF2 Key = 4194333
	/*F3 key.*/
	KeyF3 Key = 4194334
	/*F4 key.*/
	KeyF4 Key = 4194335
	/*F5 key.*/
	KeyF5 Key = 4194336
	/*F6 key.*/
	KeyF6 Key = 4194337
	/*F7 key.*/
	KeyF7 Key = 4194338
	/*F8 key.*/
	KeyF8 Key = 4194339
	/*F9 key.*/
	KeyF9 Key = 4194340
	/*F10 key.*/
	KeyF10 Key = 4194341
	/*F11 key.*/
	KeyF11 Key = 4194342
	/*F12 key.*/
	KeyF12 Key = 4194343
	/*F13 key.*/
	KeyF13 Key = 4194344
	/*F14 key.*/
	KeyF14 Key = 4194345
	/*F15 key.*/
	KeyF15 Key = 4194346
	/*F16 key.*/
	KeyF16 Key = 4194347
	/*F17 key.*/
	KeyF17 Key = 4194348
	/*F18 key.*/
	KeyF18 Key = 4194349
	/*F19 key.*/
	KeyF19 Key = 4194350
	/*F20 key.*/
	KeyF20 Key = 4194351
	/*F21 key.*/
	KeyF21 Key = 4194352
	/*F22 key.*/
	KeyF22 Key = 4194353
	/*F23 key.*/
	KeyF23 Key = 4194354
	/*F24 key.*/
	KeyF24 Key = 4194355
	/*F25 key. Only supported on macOS and Linux due to a Windows limitation.*/
	KeyF25 Key = 4194356
	/*F26 key. Only supported on macOS and Linux due to a Windows limitation.*/
	KeyF26 Key = 4194357
	/*F27 key. Only supported on macOS and Linux due to a Windows limitation.*/
	KeyF27 Key = 4194358
	/*F28 key. Only supported on macOS and Linux due to a Windows limitation.*/
	KeyF28 Key = 4194359
	/*F29 key. Only supported on macOS and Linux due to a Windows limitation.*/
	KeyF29 Key = 4194360
	/*F30 key. Only supported on macOS and Linux due to a Windows limitation.*/
	KeyF30 Key = 4194361
	/*F31 key. Only supported on macOS and Linux due to a Windows limitation.*/
	KeyF31 Key = 4194362
	/*F32 key. Only supported on macOS and Linux due to a Windows limitation.*/
	KeyF32 Key = 4194363
	/*F33 key. Only supported on macOS and Linux due to a Windows limitation.*/
	KeyF33 Key = 4194364
	/*F34 key. Only supported on macOS and Linux due to a Windows limitation.*/
	KeyF34 Key = 4194365
	/*F35 key. Only supported on macOS and Linux due to a Windows limitation.*/
	KeyF35 Key = 4194366
	/*Multiply (*) key on the numeric keypad.*/
	KeyKpMultiply Key = 4194433
	/*Divide (/) key on the numeric keypad.*/
	KeyKpDivide Key = 4194434
	/*Subtract (-) key on the numeric keypad.*/
	KeyKpSubtract Key = 4194435
	/*Period (.) key on the numeric keypad.*/
	KeyKpPeriod Key = 4194436
	/*Add (+) key on the numeric keypad.*/
	KeyKpAdd Key = 4194437
	/*Number 0 on the numeric keypad.*/
	KeyKp0 Key = 4194438
	/*Number 1 on the numeric keypad.*/
	KeyKp1 Key = 4194439
	/*Number 2 on the numeric keypad.*/
	KeyKp2 Key = 4194440
	/*Number 3 on the numeric keypad.*/
	KeyKp3 Key = 4194441
	/*Number 4 on the numeric keypad.*/
	KeyKp4 Key = 4194442
	/*Number 5 on the numeric keypad.*/
	KeyKp5 Key = 4194443
	/*Number 6 on the numeric keypad.*/
	KeyKp6 Key = 4194444
	/*Number 7 on the numeric keypad.*/
	KeyKp7 Key = 4194445
	/*Number 8 on the numeric keypad.*/
	KeyKp8 Key = 4194446
	/*Number 9 on the numeric keypad.*/
	KeyKp9 Key = 4194447
	/*Context menu key.*/
	KeyMenu Key = 4194370
	/*Hyper key. (On Linux/X11 only).*/
	KeyHyper Key = 4194371
	/*Help key.*/
	KeyHelp Key = 4194373
	/*Back key.*/
	KeyBack Key = 4194376
	/*Forward key.*/
	KeyForward Key = 4194377
	/*Media stop key.*/
	KeyStop Key = 4194378
	/*Refresh key.*/
	KeyRefresh Key = 4194379
	/*Volume down key.*/
	KeyVolumedown Key = 4194380
	/*Mute volume key.*/
	KeyVolumemute Key = 4194381
	/*Volume up key.*/
	KeyVolumeup Key = 4194382
	/*Media play key.*/
	KeyMediaplay Key = 4194388
	/*Media stop key.*/
	KeyMediastop Key = 4194389
	/*Previous song key.*/
	KeyMediaprevious Key = 4194390
	/*Next song key.*/
	KeyMedianext Key = 4194391
	/*Media record key.*/
	KeyMediarecord Key = 4194392
	/*Home page key.*/
	KeyHomepage Key = 4194393
	/*Favorites key.*/
	KeyFavorites Key = 4194394
	/*Search key.*/
	KeySearch Key = 4194395
	/*Standby key.*/
	KeyStandby Key = 4194396
	/*Open URL / Launch Browser key.*/
	KeyOpenurl Key = 4194397
	/*Launch Mail key.*/
	KeyLaunchmail Key = 4194398
	/*Launch Media key.*/
	KeyLaunchmedia Key = 4194399
	/*Launch Shortcut 0 key.*/
	KeyLaunch0 Key = 4194400
	/*Launch Shortcut 1 key.*/
	KeyLaunch1 Key = 4194401
	/*Launch Shortcut 2 key.*/
	KeyLaunch2 Key = 4194402
	/*Launch Shortcut 3 key.*/
	KeyLaunch3 Key = 4194403
	/*Launch Shortcut 4 key.*/
	KeyLaunch4 Key = 4194404
	/*Launch Shortcut 5 key.*/
	KeyLaunch5 Key = 4194405
	/*Launch Shortcut 6 key.*/
	KeyLaunch6 Key = 4194406
	/*Launch Shortcut 7 key.*/
	KeyLaunch7 Key = 4194407
	/*Launch Shortcut 8 key.*/
	KeyLaunch8 Key = 4194408
	/*Launch Shortcut 9 key.*/
	KeyLaunch9 Key = 4194409
	/*Launch Shortcut A key.*/
	KeyLauncha Key = 4194410
	/*Launch Shortcut B key.*/
	KeyLaunchb Key = 4194411
	/*Launch Shortcut C key.*/
	KeyLaunchc Key = 4194412
	/*Launch Shortcut D key.*/
	KeyLaunchd Key = 4194413
	/*Launch Shortcut E key.*/
	KeyLaunche Key = 4194414
	/*Launch Shortcut F key.*/
	KeyLaunchf Key = 4194415
	/*"Globe" key on Mac / iPad keyboard.*/
	KeyGlobe Key = 4194416
	/*"On-screen keyboard" key on iPad keyboard.*/
	KeyKeyboard Key = 4194417
	/*英数 key on Mac keyboard.*/
	KeyJisEisu Key = 4194418
	/*かな key on Mac keyboard.*/
	KeyJisKana Key = 4194419
	/*Unknown key.*/
	KeyUnknown Key = 8388607
	/*Space key.*/
	KeySpace Key = 32
	/*Exclamation mark ([code]![/code]) key.*/
	KeyExclam Key = 33
	/*Double quotation mark ([code]"[/code]) key.*/
	KeyQuotedbl Key = 34
	/*Number sign or [i]hash[/i] ([code]#[/code]) key.*/
	KeyNumbersign Key = 35
	/*Dollar sign ([code]$[/code]) key.*/
	KeyDollar Key = 36
	/*Percent sign ([code]%[/code]) key.*/
	KeyPercent Key = 37
	/*Ampersand ([code]&[/code]) key.*/
	KeyAmpersand Key = 38
	/*Apostrophe ([code]'[/code]) key.*/
	KeyApostrophe Key = 39
	/*Left parenthesis ([code]([/code]) key.*/
	KeyParenleft Key = 40
	/*Right parenthesis ([code])[/code]) key.*/
	KeyParenright Key = 41
	/*Asterisk ([code]*[/code]) key.*/
	KeyAsterisk Key = 42
	/*Plus ([code]+[/code]) key.*/
	KeyPlus Key = 43
	/*Comma ([code],[/code]) key.*/
	KeyComma Key = 44
	/*Minus ([code]-[/code]) key.*/
	KeyMinus Key = 45
	/*Period ([code].[/code]) key.*/
	KeyPeriod Key = 46
	/*Slash ([code]/[/code]) key.*/
	KeySlash Key = 47
	/*Number 0 key.*/
	Key0 Key = 48
	/*Number 1 key.*/
	Key1 Key = 49
	/*Number 2 key.*/
	Key2 Key = 50
	/*Number 3 key.*/
	Key3 Key = 51
	/*Number 4 key.*/
	Key4 Key = 52
	/*Number 5 key.*/
	Key5 Key = 53
	/*Number 6 key.*/
	Key6 Key = 54
	/*Number 7 key.*/
	Key7 Key = 55
	/*Number 8 key.*/
	Key8 Key = 56
	/*Number 9 key.*/
	Key9 Key = 57
	/*Colon ([code]:[/code]) key.*/
	KeyColon Key = 58
	/*Semicolon ([code];[/code]) key.*/
	KeySemicolon Key = 59
	/*Less-than sign ([code]<[/code]) key.*/
	KeyLess Key = 60
	/*Equal sign ([code]=[/code]) key.*/
	KeyEqual Key = 61
	/*Greater-than sign ([code]>[/code]) key.*/
	KeyGreater Key = 62
	/*Question mark ([code]?[/code]) key.*/
	KeyQuestion Key = 63
	/*At sign ([code]@[/code]) key.*/
	KeyAt Key = 64
	/*A key.*/
	KeyA Key = 65
	/*B key.*/
	KeyB Key = 66
	/*C key.*/
	KeyC Key = 67
	/*D key.*/
	KeyD Key = 68
	/*E key.*/
	KeyE Key = 69
	/*F key.*/
	KeyF Key = 70
	/*G key.*/
	KeyG Key = 71
	/*H key.*/
	KeyH Key = 72
	/*I key.*/
	KeyI Key = 73
	/*J key.*/
	KeyJ Key = 74
	/*K key.*/
	KeyK Key = 75
	/*L key.*/
	KeyL Key = 76
	/*M key.*/
	KeyM Key = 77
	/*N key.*/
	KeyN Key = 78
	/*O key.*/
	KeyO Key = 79
	/*P key.*/
	KeyP Key = 80
	/*Q key.*/
	KeyQ Key = 81
	/*R key.*/
	KeyR Key = 82
	/*S key.*/
	KeyS Key = 83
	/*T key.*/
	KeyT Key = 84
	/*U key.*/
	KeyU Key = 85
	/*V key.*/
	KeyV Key = 86
	/*W key.*/
	KeyW Key = 87
	/*X key.*/
	KeyX Key = 88
	/*Y key.*/
	KeyY Key = 89
	/*Z key.*/
	KeyZ Key = 90
	/*Left bracket ([code][lb][/code]) key.*/
	KeyBracketleft Key = 91
	/*Backslash ([code]\[/code]) key.*/
	KeyBackslash Key = 92
	/*Right bracket ([code][rb][/code]) key.*/
	KeyBracketright Key = 93
	/*Caret ([code]^[/code]) key.*/
	KeyAsciicircum Key = 94
	/*Underscore ([code]_[/code]) key.*/
	KeyUnderscore Key = 95
	/*Backtick ([code]`[/code]) key.*/
	KeyQuoteleft Key = 96
	/*Left brace ([code]{[/code]) key.*/
	KeyBraceleft Key = 123
	/*Vertical bar or [i]pipe[/i] ([code]|[/code]) key.*/
	KeyBar Key = 124
	/*Right brace ([code]}[/code]) key.*/
	KeyBraceright Key = 125
	/*Tilde ([code]~[/code]) key.*/
	KeyAsciitilde Key = 126
	/*Yen symbol ([code]¥[/code]) key.*/
	KeyYen Key = 165
	/*Section sign ([code]§[/code]) key.*/
	KeySection Key = 167
)
