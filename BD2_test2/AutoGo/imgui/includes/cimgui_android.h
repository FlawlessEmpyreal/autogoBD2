#ifndef CIMGUI_ANDROID_H
#define CIMGUI_ANDROID_H

#include <stdint.h>
#include <stdbool.h>

// 导出符号宏定义
#if defined(__GNUC__) || defined(__clang__)
    #define CIMGUI_ANDROID_API __attribute__((visibility("default")))
#else
    #define CIMGUI_ANDROID_API
#endif

#ifdef __cplusplus
extern "C" {
#endif

// ==================== 版本信息 ====================

#define CIMGUI_ANDROID_VERSION_MAJOR 1
#define CIMGUI_ANDROID_VERSION_MINOR 0
#define CIMGUI_ANDROID_VERSION_PATCH 0

CIMGUI_ANDROID_API const char* cimgui_android_get_version();

#ifndef CIMGUI_ANDROID_JAVAVM_DECLARED
#define CIMGUI_ANDROID_JAVAVM_DECLARED
#include <jni.h>
#endif

CIMGUI_ANDROID_API int cimgui_android_init(JavaVM* javavm);


/**
 * 高性能循环模式：C++ 接管主循环，Go 通过回调参与
 * @param callback Go 层的渲染回调函数指针
 * @param user_data 传递给回调的用户数据
 * 
 * 优势：每帧只有 1 次 CGO 调用（在回调处），而不是 7+ 次
 * C++ 负责：循环控制、NewFrame、Render、SwapBuffers
 * Go 负责：业务逻辑（在回调中）
 */
typedef void (*cimgui_android_render_callback)(void* user_data);
CIMGUI_ANDROID_API void cimgui_android_run_with_callback(cimgui_android_render_callback callback, void* user_data);

/**
 * 销毁后端并清理资源
 */
CIMGUI_ANDROID_API void cimgui_android_shutdown();

// ==================== 纹理管理 ====================

/**
 * 纹理格式
 */
typedef enum {
    CIMGUI_ANDROID_TEXTURE_RGBA = 1,
    CIMGUI_ANDROID_TEXTURE_RGB = 2,
    CIMGUI_ANDROID_TEXTURE_ALPHA = 3,
} CImGuiAndroidTextureFormat;

/**
 * 创建纹理（从像素数据）
 * @param pixels 像素数据指针
 * @param width 纹理宽度
 * @param height 纹理高度
 * @param format 纹理格式
 * @return OpenGL 纹理 ID，0表示失败
 */
CIMGUI_ANDROID_API unsigned int cimgui_android_create_texture(
    const void* pixels, 
    int width, 
    int height, 
    int format
);

/**
 * 更新纹理数据
 * @param texture_id OpenGL 纹理 ID
 * @param pixels 新的像素数据
 * @param width 纹理宽度
 * @param height 纹理高度
 * @param format 纹理格式
 * @return true=成功，false=失败
 */
CIMGUI_ANDROID_API bool cimgui_android_update_texture(
    unsigned int texture_id,
    const void* pixels,
    int width,
    int height,
    int format
);

/**
 * 删除纹理
 * @param texture_id OpenGL 纹理 ID
 */
CIMGUI_ANDROID_API void cimgui_android_delete_texture(unsigned int texture_id);

// ==================== 错误处理 ====================

/**
 * 获取最后的错误消息
 * @return 错误消息字符串，如果没有错误返回 NULL
 */
CIMGUI_ANDROID_API const char* cimgui_android_get_last_error();

#ifdef __cplusplus
}
#endif

#endif // CIMGUI_ANDROID_H

