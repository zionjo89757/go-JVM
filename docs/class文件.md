### 结构
```c
ClassFile {
    u4 magic;
    u2 minor_version;
    u2 major_version;
    u2 constant_pool_count;
    cp_info constant_pool[constant_pool_count-1];
    u2 access_flags;
    u2 this_class;
    u2 super_class;
    u2 interfaces_count;
    u2 interfaces[interfaces_count];
    u2 fields_count;
    field_info fields[fields_count];
    u2 methods_count;
    method_info methods[methods_count];
    u2 attributes_count;
    attribute_info attributes[attributes_count];
}
```
#### 魔数
`0xCAFEBABE`

#### 版本号
次版本号和主版本号:u2类型

#### 常量池
常量结构
```c
cp_info {
    u1 tag;
    u1 info[];
}
```
##### CONSTANT_Class              = 7
类或者接口的符号引用
##### CONSTANT_Fieldref           = 9
字段符号引用
##### CONSTANT_Methodref          = 10
普通（非接口）方法符号引用
##### CONSTANT_InterfaceMethodref = 11
接口方法符号引用
##### CONSTANT_String             = 8
java.lang.String字面量
##### CONSTANT_Integer            = 3
4字节存储整数常量
##### CONSTANT_Float              = 4
4字节存储IEEE754单精度浮点数
##### CONSTANT_Long               = 5
8字节存储整数常量
##### CONSTANT_Double             = 6
8字节存储IEEE754双精度浮点数
##### CONSTANT_NameAndType        = 12
字段或方法的名称和描述符
##### CONSTANT_Utf8               = 1
MUTF-8编码的字符串
##### CONSTANT_MethodHandle       = 15
支持新增的invokedynamic指令
##### CONSTANT_MethodType         = 16
支持新增的invokedynamic指令
##### CONSTANT_InvokeDynamic      = 18
支持新增的invokedynamic指令

#### 类访问标志
一个16位的“bitmask”，指出class文件定义的是类还是接口，访问级别是public还是private

#### 类和超类索引
两个u2类型的常量池索引，分别给出类名和超类名

#### 接口索引表
表中存放的也是常量池索引，给出该类实现的所有接口的名字

#### 字段和方法表
结构
```c
field_info {
    u2 access_flags;
    u2 name_index;
    u2 descriptor_index;
    u2 attributes_count;
    attribute_info attributes[attributes_count];
}
```
- 访问标志
- 常量池索引
- 描述符
  - 类型描述符
    - 基本类型byte、short、char、int、long、float和double的描述符是单个字母，分别对应B、S、C、I、J、F和D。
    - 引用类型的描述符是L＋类的完全限定名＋分号。
    - 数组类型的描述符是[＋数组元素类型描述符。
  - 字段描述符就是字段类型的描述符。
  - 方法描述符是（分号分隔的参数类型描述符）+返回值类型描述符，其中void返回值由单个字母V表示。
- 属性表
```结构
attribute_info {
    u2 attribute_name_index;
    u4 attribute_length;
    u1 info[attribute_length];
}
```
##### Deprecated和Synthetic
可以出现在ClassFile、field_info和method_info结构中
仅起标记作用，不包含任何数据
Deprecated属性用于指出类、接口、字段或方法已经不建议使用
Synthetic属性用来标记源文件中不存在、由编译器生成的类成员，引入Synthetic属性主要是为了支持嵌套类和嵌套接口。
##### SourceFile
可选定长属性，只会出现在ClassFile结构中，用于指出源文件名
##### ConstantValue
定长属性，只会出现在field_info结构中，用于表示常量表达式的值
##### Code
变长属性，只存在于method_info结构中，存放字节码等方法相关信息。
```java
Code_attribute {
    u2 attribute_name_index;
    u4 attribute_length;
    u2 max_stack;
    u2 max_locals;
    u4 code_length;
    u1 code[code_length];
    u2 exception_table_length;
    { u2 start_pc;
        u2 end_pc;
        u2 handler_pc;
        u2 catch_type;
    } exception_table[exception_table_length];
    u2 attributes_count;
    attribute_info attributes[attributes_count];
}
```
- max_stack给出操作数栈的最大深度
- max_locals给出局部变量表大小
- 字节码，存在u1表
- 异常处理表和属性表
##### Exceptions
变长属性,记录方法抛出的异常表
##### LineNumberTable和LocalVariableTable
LineNumberTable属性表存放方法的行号信息，
LocalVariableTable属性表中存放方法的局部变量信息。
#### 类属性

