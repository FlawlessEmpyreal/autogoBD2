package files

// IsFile 返回路径path是否是文件。
// 参数:
//
//	path: 路径。
//
// 返回:
//
//	bool: 路径是否是文件。
func IsFile(path string) bool {
	return false
}

// IsDir 返回路径path是否是文件夹。
// 参数:
//
//	path: 路径。
//
// 返回:
//
//	bool: 路径是否是文件夹。
func IsDir(path string) bool {
	return false
}

// IsEmptyDir 返回文件夹path是否为空文件夹。如果该路径并非文件夹，则直接返回false。
// 参数:
//
//	path: 文件夹路径。
//
// 返回:
//
//	bool: 文件夹是否为空。
func IsEmptyDir(path string) bool {
	return false
}

// Create 创建一个文件或文件夹并返回是否创建成功。如果文件已经存在，则直接返回true。
// 参数:
//
//	path: 路径。
//
// 返回:
//
//	bool: 是否创建成功。
func Create(path string) bool {
	return false
}

// Exists 返回在路径path处的文件是否存在。
// 参数:
//
//	path: 路径。
//
// 返回:
//
//	bool: 文件是否存在。
func Exists(path string) bool {
	return false
}

// EnsureDir 确保路径path所在的文件夹存在。如果该路径所在文件夹不存在，则创建该文件夹。
// 参数:
//
//	path: 路径。
//
// 返回:
//
//	bool: 是否确保成功。
func EnsureDir(path string) bool {
	return false
}

// Read 读取文本文件path的所有内容并返回。
// 参数:
//
//	path: 文件路径。
//
// 返回:
//
//	string: 文件内容。
func Read(path string) string {
	return ""
}

// ReadBytes 读取文件path的所有内容并返回。
// 参数:
//
//	path: 文件路径。
//
// 返回:
//
//	[]byte: 文件数据。
func ReadBytes(path string) []byte {
	return nil
}

// Write 把text写入到文件path中。如果文件存在则覆盖，不存在则创建。
// 参数:
//
//	path: 文件路径。
//	text: 要写入的文本。
func Write(path, text string) {

}

// WriteBytes 把bytes写入到文件path中。如果文件存在则覆盖，不存在则创建。
// 参数:
//
//	path: 文件路径。
//	bytes: 要写入的字节数组。
func WriteBytes(path string, bytes []byte) {

}

// Append 把text追加到文件path的末尾。如果文件不存在则创建。
// 参数:
//
//	path: 文件路径。
//	text: 要追加的文本。
func Append(path string, text string) {

}

// AppendBytes 把bytes追加到文件path的末尾。如果文件不存在则创建。
// 参数:
//
//	path: 文件路径。
//	bytes: 要追加的字节数组。
func AppendBytes(path string, bytes []byte) {

}

// Copy 复制文件，返回是否复制成功。
// 参数:
//
//	fromPath: 源文件路径。
//	toPath: 目标文件路径。
//
// 返回:
//
//	bool: 是否复制成功。
func Copy(fromPath, toPath string) bool {
	return false
}

// Move 移动文件，返回是否移动成功。
// 参数:
//
//	fromPath: 源文件路径。
//	toPath: 目标文件路径。
//
// 返回:
//
//	bool: 是否移动成功。
func Move(fromPath, toPath string) bool {
	return false
}

// Rename 重命名文件，并返回是否重命名成功。
// 参数:
//
//	path: 原文件路径。
//	newName: 新文件名。
//
// 返回:
//
//	bool: 是否重命名成功。
func Rename(path, newName string) bool {
	return false
}

// GetName 返回文件的文件名。例如files.GetName("/sdcard/1.txt")返回"1.txt"。
// 参数:
//
//	path: 文件路径。
//
// 返回:
//
//	string: 文件名。
func GetName(path string) string {
	return ""
}

// GetNameWithoutExtension 返回不含拓展名的文件的文件名。例如files.GetNameWithoutExtension("/sdcard/1.txt")返回"1"。
// 参数:
//
//	path: 文件路径。
//
// 返回:
//
//	string: 不含拓展名的文件名。
func GetNameWithoutExtension(path string) string {
	return ""
}

// GetExtension 返回文件的拓展名。例如files.GetExtension("/sdcard/1.txt")返回"txt"。
// 参数:
//
//	path: 文件路径。
//
// 返回:
//
//	string: 文件拓展名。
func GetExtension(path string) string {
	return ""
}

// GetMd5 返回文件的MD5。
// 参数:
//
//	path: 文件路径。
//
// 返回:
//
//	string: MD5值。
func GetMd5(path string) string {
	return ""
}

// Remove 删除文件或文件夹，如果是文件夹会删除整个文件夹包含里面的所有文件，返回是否删除成功。
// 参数:
//
//	path: 文件或文件夹路径。
//
// 返回:
//
//	bool: 是否删除成功。
func Remove(path string) bool {
	return false
}

// Path 返回相对路径对应的绝对路径。例如files.Path("./1.png")，如果运行这个语句的脚本位于文件夹"/sdcard/脚本/"中，则返回"/sdcard/脚本/1.png"。
// 参数:
//
//	relativePath: 相对路径。
//
// 返回:
//
//	string: 绝对路径。
func Path(relativePath string) string {
	return ""
}

// ListDir 列出文件夹path下的所有文件和文件夹。
// 参数:
//
//	path: 文件夹路径。
//
// 返回:
//
//	[]string: 文件和文件夹的路径列表。
func ListDir(path string) []string {
	return nil
}
