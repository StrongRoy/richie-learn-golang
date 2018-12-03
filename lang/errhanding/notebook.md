defer 滴哦啊用

* 确保调用在函数结束时发生
* 参数在defer语句时计算
* defer列表为先进后出


# 何时使用defer调用
 * open/close
 * lock/unlock
 * PrintHead /PrintFooter
 
 
 # error vs panic 
 
 * 意料之中的： 使用error 如文件打不开
 *  意料之外的： 使用panic。如：数组越界
 