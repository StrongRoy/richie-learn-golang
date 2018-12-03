# duck typing

大黄鸭是鸭子吗？
* 畅通类型： 脊椎动物，。。。
# duck typing 是鸭子
* 像鸭子走路，像鸭子叫（长得像鸭子），那么就是鸭子
* 描述食物的外部行为而非内部结构
* 严格来说属于结构化类型系统，类似duck typing

# python中的duck typing
```python
def download(retriever):
    return retriever.get("www.baidu.com")
```
   
* 运行时候才知道传入的retriever有没有get
* 需要注释说明接口


# c++中的duck typing
```c++
template  <class R>
string download(const R& retriever){
    return retriever.get("www.baidu.com");
}
```

* 编译时候才知道传入的retriever有没有get
* 需要注释说明接口

#java中类似代码

```java
<R extends   Retriever>
string download(R r){
    return r.get("www.baidu.com");
}
```

* 传入的参数必须实现Retriever接口
* 不是duck typing


# go语言中的duck typing
* 同时需要Readable，Appendable怎么办？（apache polygene）
* 同时具有python，c++的duck typing的灵活性
* 又具有java的类型检查



   