package imgui

/*
#include "includes/cimgui_android.h"
#include <stdlib.h>
*/
import "C"
import (
	"image"
	"unsafe"
)

// ==================== 纹理对象封装 ====================

// Texture 纹理对象
type Texture struct {
	ID     TextureRef
	Width  int
	Height int
}

// CreateTextureNrgba 从 NRGBA 图片创建纹理
func CreateTextureNrgba(img *image.NRGBA) *Texture {
	if img == nil {
		return nil
	}

	pix := img.Pix
	if len(pix) == 0 {
		return nil
	}

	// 自动从图片对象获取尺寸
	width := img.Bounds().Dx()
	height := img.Bounds().Dy()

	if width <= 0 || height <= 0 {
		return nil
	}

	// 创建OpenGL纹理
	texID := C.cimgui_android_create_texture(
		unsafe.Pointer(&pix[0]),
		C.int(width),
		C.int(height),
		C.int(C.CIMGUI_ANDROID_TEXTURE_RGBA),
	)

	return &Texture{
		ID:     *NewTextureRefTextureID(*NewTextureIDFromC(&texID)),
		Width:  width,
		Height: height,
	}
}

// Delete 删除纹理
func (t *Texture) Delete() {
	if t != nil {
		C.cimgui_android_delete_texture(C.uint(t.ID.TexID()))
	}
}
