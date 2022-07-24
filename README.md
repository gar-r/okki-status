# ![logo](docs/logo.png "okki-status for dwm")

## What is okki-status?

okki-status is a simple status bar written in Go for tiling window managers, like [dwm](http://dwm.suckless.org/) or sway.

Here are some screenshots of it in action:

dwm status bar

![screenshot](docs/dwm.png "screen shot of dwm desktop with okki-status")

sway status bar

![screenshot](docs/sway.png "screen shot of sway desktop with okki-status")

## Installation

Make sure, that you have Go version 1.17 or above.

1. clone the source code: `git clone https://git.okki.hu/garric/okki-status`
2. switch to the source directory: `cd okki-status`
3. build&install: `sudo make clean install`

Install will add the following binary to your system:

```
/usr/local/bin/okki-status
```

## Configuring the status bar

`okki-status` configuration is actually go source code. But don't worry, if you are interested in tiling window managers and already reading this page, you can probably do it without much effort :)

You can find all config sources under the `config` package:
   * `config.go`: contains the bar itself
   * `block.go`: contains the block definitions
   * `render.go`: contains renderer configuration

After changing the configuration, you will need to rebuild the application as explained in the previous section.


## Window manager support

In order to support multiple window managers, `okki-status` uses the concept of _renderers_. Each renderer only works with the window manager it is designed for. For more information on how to configure or implement new renderers, see the relevant section.

The status bar contains a few pre-implemented renderers:
   * __stdout__: prints to the standard output
   * __xroot__: sample renderer for `dwm`
   * __swaybar__: sample renderer for `sway`

### dwm

The status bar uses the `xsetroot` program from the `xorg-xsetroot` package to interact with the Xorg root window, so this package must be installed on your system.

To automatically start the status bar, insert the following line into your `xinit` config right before you launch dwm:

```
exec /usr/local/bin/okki-status &
```

Example configuration for dwm:

```go
// rendering configuration
var R core.Renderer = &renderer.XRoot{
	Separator: "\t",
}
```

### sway

The status bar uses the swaybar json protocol, and depends on `swaybar` built into sway wm.

To enable `okki-status` under sway, add it as your `status_command` in the sway config file:

```
bar {
    status_command /usr/local/bin/okki-status
    # ...
}
```

The sway renderer supports a wider range of display options (text color, background color, alignment, etc). However, this means configuring the built-in sway renderer is a bit more involved.

Example sway renderer config:

```go
var R core.Renderer = &renderer.SwayBar{
	BlockCfg: []*renderer.SwayBarBlockConfig{
		{
			BlockName:      "volume",
			SeparatorWidth: 25,
		},
		{
			BlockName:      "ram",
			SeparatorWidth: 25,
		},
    // and so on...
  }
}
```

## Block configuration

As mentioned above, `block.go` contains the actual blocks that will be on the bar. Each block has the following attributes:
   * __Module__: the code that contains instructions on how to fetch the status text for the block
   * __Name__: the name of the block. You can use this handle to reference the block in `okki-refresh` (see later)
   * __Prefix__: text to render before the block
   * __Postfix__: text to render after the block

Prefix and postfix are useful to render icons, or other static text before or after the block's contents.

Example block:

```go
var battery = &core.Block{
	Name:   "battery",
	Prefix: " ",
	Module: core.NewCachingModule(
		&module.Battery{
			Device: "BAT0",
		},
		1 * time.Minute,
	),
}
```

Note: in the above example we can see another concept in action, a _caching_ module. A caching module can remember the last value, and does not always need to query the OS for its state. This can be useful for blocks where fetching the information is expensive or time-consuming. Expiration can be set independently, and forcing a refresh can be achieved with `okki-refresh` (see later).

## Modules

While it is relatively easy to implement custom modules, `okki-status` contains a set of built-in modules to use (and for reference):

| Module     | Description                                  | Special parameters                              |
| ---------- | -------------------------------------------- | ----------------------------------------------- |
| wifi       | connected network name and signal strength   | wifi device name                                |
| ram        | physical memory usage in percent             |                                                 |
| volume     | volume level or muted state                  |                                                 |
| brightness | display brightness percent                   |                                                 |
| battery    | remaining battery percent and charging state | battery device name                             |
| clock      | current date/time                            | layout                                          |
| updates    | displays the number of available updates     | command (must return one update per line), args |


## External Dependencies


Some built-in modules depend on external tools which need to be available on your system if you wish to enable them:

| Module         | Dependencies                               |
| -------------- | ------------------------------------------ |
| brightness     | brillo (aur)                               |
| volume         | pamixer                                    |
| wifi           | iw                                         |
| updates        | pacman-contrib                             |

Other modules work fine without external dependencies.

## Testing the bar

You can easily test the bar by setting up the `stdout` renderer:

```go
var R core.Renderer = &renderer.StdOut{
	Separator: "   ",
	Terminator: "\n",
}
```

Recompile and run the binary directly (instead of through your window manager of choice) to see the output on the standard output.

_Note_: only one instance of `okki-status` can run at the same time, so you will need to stop any running instances, _including_ the instances started by window managers.


## Advanced

### Implementing a custom module

This requires go programming skills. In order to implement a custom module:

1.  implement the `core.Module` interface
1.  configure a new block that refers to an instance of your new module
1.  recompile and test it

## Reacting to external events

In some cases it is not efficient for the module to continuously poll the system for status updates, but we still want to react promptly to external events.

Good examples for this are the **brightness** and **volume** modules. These values rarely change _by themselves_ so a relatively rare polling rate is sufficient. However when the user changes the volume or brightness manually, we want to update the status bar as promptly as possible.

For this specific scenario you can use the [okki-refresh](https://git.okki.hu/garric/okki-refresh) utility. Calling `okki-refresh` can send a signal to `okki-status` to immediately refresh a given module.

For more details, see the [okki-refresh readme](https://git.okki.hu/garric/okki-refresh).

### Binding multimedia keys

A typical example setup for immediately updating the status bar after pressing multimedia keys will involve using `SHCMD` to follow up the bound command with `okki-refresh module_name`.

Continuing with the example from the previous section using the standard dwm config:

```
## config.h (dwm source file)
static Key keys[] = {
   { 0, XF86XK_MonBrightnessUp, spawn, SHCMD("brillo -A 10; okki-refresh brightness") },
}
```

The same configuration in sway:

```
set $refresh "/usr/local/bin/okki-refresh"
bindsym XF86MonBrightnessUp exec $brightness up 1; exec $refresh brightness
```

All other multimedia keys can be configured in a similar fashion with the appropriate key-codes and module names.

### Pacman hook

If you are using the `updates` module with `pacman`, you can set up a hook to refresh the module after each package upgrade (on Arch: `/etc/pacman.d/hooks/okki-status.hook`):

```
[Trigger]
Operation=Upgrade
Type=Package
Target=*

[Action]
Exec=/usr/local/bin/okki-refresh updates
When=PostTransaction
```
