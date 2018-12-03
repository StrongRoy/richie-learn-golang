# map操作
* 创建： make(map[string]int)
* 获取元素m[key]
* key不存在是，获得value类型的初值
* 用value,ok := m[key]来判断是否存在key
* 用delete删除一个key


# map遍历

* 使用range遍历key，或者遍历key，value对
* 不保证遍历顺序，如需顺序，需手动对key进行排序
* 使用len获取长度

# map的key

* map使用哈希表，必须可以**比较相等**
* 除了slice，map，function的内建类型都可以作为key，因为这几个类型不可以判断相等
* Struct类型不包含上述字段，也可作为key


# 寻找最长不含重复字符串的字串
对于一个字母x

* lastOccurred[x]不存在，或者<start -> 无需操作
* lastOccurred[x] >= start -> 更新操作
* 更新lastOccurred[x] ,更新maxLength

# rune 相当于char