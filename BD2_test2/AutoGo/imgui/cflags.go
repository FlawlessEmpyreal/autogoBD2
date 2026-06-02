package imgui

// #cgo CPPFLAGS: -DCIMGUI_DEFINE_ENUMS_AND_STRUCTS -DIMGUI_USE_WCHAR32 -DIMGUI_IMPL_OPENGL_ES3
// #cgo arm64 LDFLAGS: -L../../resources/libs/arm64-v8a -limgui
// #cgo amd64 LDFLAGS: -L../../resources/libs/x86_64 -limgui
// #cgo 386 LDFLAGS: -L../../resources/libs/x86 -limgui
import "C"
