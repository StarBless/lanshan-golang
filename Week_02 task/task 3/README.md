（1）一个结构体所占空间大小，与下面哪些相关：ABC

​			A 成员本身大小；B 成员对齐系数；C 系统字长

（2）下面这个结构体的成员怎么排布，占用内存最小？（假设系统64位）

```go
type A struct{
  byte1 byte
  a struct{}
  num1 int32
  str string
}

type A struct{
  str string
  num1 int32
  byte1 byte
  a struct{}
}
```

（3）Go字符串中，每个字符占用多少字节：C

​			A 1；B 3；C 1-4

   (4)  Go的map使用的数据结构是：C

​			A B+树；B 开放寻址法的Hash表 ；C 拉链法Hash

（5）空结构体的地址在任何时候都是zerobase？A

​		   A 是；B 不是

（6）空接口就是nil接口？

空接口指定了0个方法，可以保存任何类型的值。

而即便接口内的具体值为 nil，方法仍然会被 nil 接收者调用。