package colors

import "image/color"

// The color foundations from the Apple Human Interface Guidelines.
// https://developer.apple.com/design/human-interface-guidelines/foundations/color/

// iOS system colors.
// https://developer.apple.com/design/human-interface-guidelines/foundations/color/#system-colors-ios
var (
	IOSLightRed    = color.NRGBA{255, 59, 48, 255}
	IOSLightOrange = color.NRGBA{255, 149, 0, 255}
	IOSLightYellow = color.NRGBA{255, 204, 0, 255}
	IOSLightGreen  = color.NRGBA{52, 199, 89, 255}
	IOSLightMint   = color.NRGBA{0, 199, 190, 255}
	IOSLightTeal   = color.NRGBA{48, 176, 199, 255}
	IOSLightCyan   = color.NRGBA{50, 173, 230, 255}
	IOSLightBlue   = color.NRGBA{0, 122, 255, 255}
	IOSLightIndigo = color.NRGBA{88, 86, 214, 255}
	IOSLightPurple = color.NRGBA{175, 82, 222, 255}
	IOSLightPink   = color.NRGBA{255, 45, 85, 255}
	IOSLightBrown  = color.NRGBA{162, 132, 94, 255}

	IOSDarkRed    = color.NRGBA{255, 69, 58, 255}
	IOSDarkOrange = color.NRGBA{255, 159, 10, 255}
	IOSDarkYellow = color.NRGBA{255, 214, 10, 255}
	IOSDarkGreen  = color.NRGBA{48, 209, 88, 255}
	IOSDarkMint   = color.NRGBA{102, 212, 207, 255}
	IOSDarkTeal   = color.NRGBA{64, 200, 224, 255}
	IOSDarkCyan   = color.NRGBA{100, 210, 255, 255}
	IOSDarkBlue   = color.NRGBA{10, 132, 255, 255}
	IOSDarkIndigo = color.NRGBA{94, 92, 230, 255}
	IOSDarkPurple = color.NRGBA{191, 90, 242, 255}
	IOSDarkPink   = color.NRGBA{255, 55, 95, 255}
	IOSDarkBrown  = color.NRGBA{172, 142, 104, 255}
)

// iOS system grays.
// https://developer.apple.com/design/human-interface-guidelines/foundations/color/#system-gray-colors-ios
var (
	IOSLightGray  = color.NRGBA{142, 142, 147, 255}
	IOSLightGray2 = color.NRGBA{174, 174, 178, 255}
	IOSLightGray3 = color.NRGBA{199, 199, 204, 255}
	IOSLightGray4 = color.NRGBA{209, 209, 214, 255}
	IOSLightGray5 = color.NRGBA{229, 229, 234, 255}
	IOSLightGray6 = color.NRGBA{242, 242, 247, 255}

	IOSDarkGray  = color.NRGBA{142, 142, 147, 255}
	IOSDarkGray2 = color.NRGBA{99, 99, 102, 255}
	IOSDarkGray3 = color.NRGBA{72, 72, 74, 255}
	IOSDarkGray4 = color.NRGBA{58, 58, 60, 255}
	IOSDarkGray5 = color.NRGBA{44, 44, 46, 255}
	IOSDarkGray6 = color.NRGBA{28, 28, 30, 255}
)

// MacOS system colors.
// https://developer.apple.com/design/human-interface-guidelines/foundations/color/#system-colors-macos
var (
	MacOSLightRed    = color.NRGBA{255, 59, 48, 255}
	MacOSLightOrange = color.NRGBA{255, 149, 0, 255}
	MacOSLightYellow = color.NRGBA{255, 204, 0, 255}
	MacOSLightGreen  = color.NRGBA{40, 205, 65, 255}
	MacOSLightMint   = color.NRGBA{0, 199, 190, 255}
	MacOSLightTeal   = color.NRGBA{89, 173, 196, 255}
	MacOSLightCyan   = color.NRGBA{85, 190, 240, 255}
	MacOSLightBlue   = color.NRGBA{0, 122, 255, 255}
	MacOSLightIndigo = color.NRGBA{88, 86, 214, 255}
	MacOSLightPurple = color.NRGBA{175, 82, 222, 255}
	MacOSLightPink   = color.NRGBA{255, 45, 85, 255}
	MacOSLightBrown  = color.NRGBA{162, 132, 94, 255}
	MacOSLightGray   = color.NRGBA{142, 142, 147, 255}

	MacOSDarkRed    = color.NRGBA{255, 69, 58, 255}
	MacOSDarkOrange = color.NRGBA{255, 159, 10, 255}
	MacOSDarkYellow = color.NRGBA{255, 214, 10, 255}
	MacOSDarkGreen  = color.NRGBA{50, 215, 75, 255}
	MacOSDarkMint   = color.NRGBA{102, 212, 207, 255}
	MacOSDarkTeal   = color.NRGBA{106, 196, 220, 255}
	MacOSDarkCyan   = color.NRGBA{90, 200, 245, 255}
	MacOSDarkBlue   = color.NRGBA{10, 132, 255, 255}
	MacOSDarkIndigo = color.NRGBA{94, 92, 230, 255}
	MacOSDarkPurple = color.NRGBA{191, 90, 242, 255}
	MacOSDarkPink   = color.NRGBA{255, 55, 95, 255}
	MacOSDarkBrown  = color.NRGBA{172, 142, 104, 255}
	MacOSDarkGray   = color.NRGBA{152, 152, 157, 255}
)

// watchOS system colors.
// https://developer.apple.com/design/human-interface-guidelines/foundations/color/#system-colors-watchos
var (
	WatchOSRed    = color.NRGBA{255, 59, 48, 255}
	WatchOSOrange = color.NRGBA{255, 149, 0, 255}
	WatchOSYellow = color.NRGBA{255, 230, 32, 255}
	WatchOSGreen  = color.NRGBA{4, 222, 113, 255}
	WatchOSMint   = color.NRGBA{102, 212, 207, 255}
	WatchOSTeal   = color.NRGBA{106, 196, 220, 255}
	WatchOSCyan   = color.NRGBA{90, 200, 250, 255}
	WatchOSBlue   = color.NRGBA{32, 148, 250, 255}
	WatchOSIndigo = color.NRGBA{120, 122, 255, 255}
	WatchOSPurple = color.NRGBA{191, 90, 242, 255}
	WatchOSPink   = color.NRGBA{250, 17, 79, 255}
	WatchOSBrown  = color.NRGBA{172, 142, 104, 255}
	WatchOSGray   = color.NRGBA{155, 160, 170, 255}
)
