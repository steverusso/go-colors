### MOVED TO [https://github.com/gio-tools/colors](https://github.com/gio-tools/colors)

---

# go-colors

A collection of established color values using `color.NRGBA` (non-alpha-premultiplied).

```go
import "github.com/steverusso/go-colors"

// Material colors such as...
colors.Pink200
colors.AmberA400
colors.BlueGrey500

// Apple foundation colors such as...
colors.IOSDarkBlue
colors.MacOSLightPurple
colors.WatchOSTeal

// CSS colors such as...
colors.AliceBlue
colors.DarkMagenta
colors.LimeGreen
colors.Red
```

## Why not alpha-premultiplied (`color.RGBA`)?

1. [Gio](https://gioui.org/) APIs use `color.NRGBA`, and the primary intent of this package is for developing with Gio.
2. As mentioned in [the commit description](https://github.com/gioui/gio/commit/21ef492cc9dfd9161e9fc57df25684865cd1c847) when Gio switched, most APIs use non-premultiplied for the sake of clarity.

## License

This is just a bunch of color values in Go. If a license needs to be mentioned
at all, it's unlicensed and in the public domain.
