package storages

// Get 从本地存储中取出键值为 key 的数据并返回。
// 参数:
//
//	table:表名称
//	key: 要查询的键。
//
// 返回:
//
//	string: 键对应的值，如果键不存在则返回空字符串。
func Get(table, key string) string {
	return ""
}

// Put 把值 value 保存到本地存储中。
// 参数:
//
//	table:表名称
//	key: 要保存的键。
//	value: 要保存的值。
//
// 返回:
//
//	无返回值。
func Put(table, key, value string) {

}

// Remove 移除键值为 key 的数据。
// 参数:
//
//	table:表名称
//	key: 要移除的键。
//
// 返回:
//
//	无返回值。
func Remove(table, key string) {

}

// Contains 返回该本地存储是否包含键值为 key 的数据。
// 参数:
//
//	table:表名称
//	key: 要检查的键。
//
// 返回:
//
//	bool: 如果键存在则返回 true，否则返回 false。
func Contains(table, key string) bool {
	return false
}

// GetAll 获取所有键值对。
// 参数:
//
//	table:表名称
//
// 返回:
//
//	map[string]string。
func GetAll(table string) map[string]string {
	return nil
}

// Clear 移除该本地存储的所有数据。
// 参数:
//
//	table:表名称
//
// 返回:
//
//	无返回值。
func Clear(table string) {

}
