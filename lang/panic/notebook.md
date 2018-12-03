# panic


* 停止当前函数执行
* 一直向上返回 执行每一层的defer
* 如果没有遇见recover 程序退出


![](.notebook_images/1edd4ba3.png)


# recover 

* 仅在defer调用和使用
* 获取panic的值
* 如果无法处理 可️重新panic

