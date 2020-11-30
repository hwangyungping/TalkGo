

# 参考书目

EffectiveModernCpp

https://github.com/kelthuzadx/EffectiveModernCppChinese



# [C++11 新特性学习](https://www.cnblogs.com/WindSun/p/11336631.html)



在Linux下编译C++11

[![复制代码](https://common.cnblogs.com/images/copycode.gif)](javascript:void(0);)

```
#include<typeinfo>
int main()
{
    auto a=10;
    cout<<typeid(a).name()<<endl;    //得到a的类型，只是一个字符串
    return 0;   
}
```

[![复制代码](https://common.cnblogs.com/images/copycode.gif)](javascript:void(0);)

编译需要加-std=c++11，如下例：

[![0.98332977099667](https://img2018.cnblogs.com/blog/1475571/201908/1475571-20190811204527160-1525165763.png)](https://img2018.cnblogs.com/blog/1475571/201908/1475571-20190811204518823-903747019.png)

## auto

C++11中引入auto第一种作用是为了自动类型推导

auto的自动类型推导，用于从初始化表达式中推断出变量的数据类型。通过auto的自动类型推导，可以大大简化我们的编程工作

auto实际上实在编译时对变量进行了类型推导，所以不会对程序的运行效率造成不良影响

另外，似乎auto并不会影响编译速度，因为编译时本来也要右侧推导然后判断与左侧是否匹配。

[![复制代码](https://common.cnblogs.com/images/copycode.gif)](javascript:void(0);)

```
auto a; // 错误，auto是通过初始化表达式进行类型推导，如果没有初始化表达式，就无法确定a的类型  
auto i = 1; 
auto d = 1.0; 
auto str = "Hello World"; 
auto ch = 'A'; 
auto func = less<int>(); 
vector<int> iv; 
auto ite = iv.begin(); 
auto p = new foo() // 对自定义类型进行类型推导
```

[![复制代码](https://common.cnblogs.com/images/copycode.gif)](javascript:void(0);)

auto不光有以上的应用，它在模板中也是大显身手，比如下例这个加工产品的例子中，如果不使用auto就必须声明Product这一模板参数：

```
template <typename Product, typename Creator>  
void processProduct(const Creator& creator) { 
    Product* val = creator.makeObject(); 
    // do somthing with val 
}
```

如果使用auto，则可以这样写：

```
template <typename Creator> 
void processProduct(const Creator& creator) {  
    auto val = creator.makeObject(); 
    // do somthing with val 
}
```

抛弃了麻烦的模板参数，整个代码变得更加正解了。

## decltype

decltype实际上有点像auto的反函数，auto可以让你声明一个变量，而decltype则可以从一个变量或表达式中得到类型，有实例如下：

```
int x = 3;  
decltype(x) y = x;
```

有人会问，decltype的实用之处在哪里呢，我们接着上边的例子继续说下去，如果上文中的加工产品的例子中我们想把产品作为返回值该怎么办呢？我们可以这样写：

```
template <typename Creator>  
auto processProduct(const Creator& creator) -> decltype(creator.makeObject()) {  
    auto val = creator.makeObject();  
    // do somthing with val  
}
```

**typeid().name()与decltype**

typeid().name得到的仅仅是一个字符串

而decltype可以提取出类型

## nullptr

nullptr是为了解决原来C++中NULL的二义性问题而引进的一种新的类型，因为NULL实际上代表的是0

[![复制代码](https://common.cnblogs.com/images/copycode.gif)](javascript:void(0);)

```
void F(int a){  
    cout<<a<<endl;  
}  
 
void F(int *p){  
    assert(p != NULL);  
 
    cout<< p <<endl;  
}  
 
int main(){  
 
    int *p = nullptr;  
    int *q = NULL;  
    bool equal = ( p == q ); // equal的值为true，说明p和q都是空指针  
    int a = nullptr; // 编译失败，nullptr不能转型为int  
    F(0); // 在C++98中编译失败，有二义性；在C++11中调用F（int）  
    F(nullptr);  
 
    return 0;  
}
```

[![复制代码](https://common.cnblogs.com/images/copycode.gif)](javascript:void(0);)

## 序列for循环

在C++中for循环可以使用类似java的简化的for循环，可以用于遍历数组，容器，string以及由begin和end函数定义的序列（即有Iterator），示例代码如下：

```
map<string, int> m{{"a", 1}, {"b", 2}, {"c", 3}};  
for (auto p : m){  
    cout<<p.first<<" : "<<p.second<<endl;  
}
```

## 更加优雅的初始化

在引入C++11之前，只有数组能使用初始化列表，其他容器想要使用初始化列表，只能用以下方法：

```
int arr[3] = {1, 2, 3};
vector<int> v(arr, arr + 3);
```

在C++11中，我们可以使用以下语法来进行替换：

```
int arr[3]{1, 2, 3};
vector<int> iv{1, 2, 3};
map<int, string>{{1, "a"}, {2, "b"}};
string str{"Hello World"};
```

## alignof和alignas 

C++11标准中，为了支持对齐，引入了两个关键字：操作符alignof和对齐描述符alignas。

操作符alignof：用于返回类的对齐方式，即对齐值

对齐描述符alignas：用于设置类使用哪种对齐方式

[![复制代码](https://common.cnblogs.com/images/copycode.gif)](javascript:void(0);)

```
struct alignas(32) A
{
    char a;
    int b;
};
 
int main()
{
    //由于用alignas设置了对齐方式，原本大小为8字节的类A变为32字节了
    cout << sizeof(A) << endl;       //输出：32
    cout << alignof(A) << endl;      //输出：32
 
    //alignas既可以接受常量表达式，也可以接受类型作为参数
    alignas(8) int a;
    alignas(alignof(double)) int b;
}
```

[![复制代码](https://common.cnblogs.com/images/copycode.gif)](javascript:void(0);)

  C++11标准中规定了一个“基本对齐值”。一般情况下其值等于平台上支持的最大标量类型数据的对齐值。我们可以通过alignof(std::max_align_t)来查询其值。而像之前我们把对齐值设置为32位的做法称为扩展对齐，这按照C++标准该程序是不规范的，可能会导致未知的编译错误或者运行时错误。

  其余对齐方式有STL库函数(std::align)，还有STL库模板类型(aligned_storage和aligned_union)。

## Lambda表达式

见另一篇：<https://www.cnblogs.com/WindSun/p/11182276.html>



# [C++的四种转换(const_cast、static_cast、dynamic_cast、reinterpreter_cast)](https://www.cnblogs.com/WindSun/p/11434419.html)



## static_cast

相当于C语言中的强制转换：(类型)表达式或类型(表达式)，用于各种隐式转换

非const转const、void*转指针、int和char相互转换

用于基类和子类之间的**指针和引用**转换,非指针直接报错

向上转化是安全的，如果向下转能(指针或引用)成功但是不安全，结果未知；

## dynamic_cast

用于动态类型转换。只能用于含有**虚函数**的类，必须用在多态体系种，用于类层次间的向上和向下转化。只能转指针或引用。向下转化时，如果是非法的对于指针返回NULL，对于引用抛异常。

在进行下行转换时，dynamic_cast具有类型检查的功能，比static_cast更安全。

如果没有virtual方法进行下行转换(指针或引用)会直接报错

## const_cast

常量指针被转化成非常量的指针，并且仍然指向原来的对象；

常量引用被转换成非常量的引用，并且仍然指向原来的对象；

const_cast一般用于修改底指针。如const char *p形式。

cpp

```cpp
const int a=10;
int *p=const_cast<int*>(&a);	//p和a指向同一块内存地址
*p = 100;	//修改*p，但a=10，*p=100
```

## reinterpret_cast

(重解释转换)几乎什么都可以转，比如将int转指针，可能会出问题，尽量少用；随意的转换编译都会通过，但是不安全的转换运行时会异常

错误的使用reinterpret_cast很容易导致程序的不安全，只有将转换后的类型值转换回到其原始类型，这样才是正确使用reinterpret_cast方式。

reinterpret_cast不能转换掉表达式的const

可以用在将void*转换为int类型

cpp

```cpp
unsigned short Hash( void *p ) {
   unsigned int val = reinterpret_cast<unsigned int>( p );
   return ( unsigned short )( val ^ (val >> 16));
}
```



# C++11 Lambda函数

[C++11 Lambda函数](https://www.cnblogs.com/WindSun/p/11182276.html)

## Lambda函数

C++11新增了lambda函数，其基本格式如下

```
1 [捕捉列表] (参数) mutable -> 返回值类型 {函数体}
```

说明

- []是lambda的引出符，捕捉列表能够捕捉上下文中的变量，来供lambda函数使用：

　　　　[var] 表示以值传递方式捕捉变量var

　　　　[=] 表示值传递捕捉所有父作用域变量

　　　　[&var] 表示以引用传递方式捕捉变量var

　　　　[&] 表示引用传递捕捉所有父作用域变量

　　　　[this] 表示值传递方式捕捉当前的this指针

​    　　　还有一些组合：

　　　　[=,&a] 表示以引用传递方式捕捉a,值传递方式捕捉其他变量

​    　　　注意：

​    　　　捕捉列表不允许变量重复传递，如：[=,a]、[&,&this]，会引起编译时期的错误

- 参数列表与普通函数的参数列表一致。如果不需要传递参数，可以联连同()一同【省略】。

- mutable 可以取消Lambda的常量属性，因为Lambda默认是const属性；multable仅仅是让Lamdba函数体修改值传递的变量，但是修改后并不会影响外部的变量。

- ->返回类型如果是void时，可以连->一起【省略】，如果返回类型很明确，可以省略，让编译器自动推倒类型。

- 函数体和普通函数一样，除了可以使用参数之外，还可以使用捕获的变量。

最简单的Lambda函数：

```
1 []{}
```

实例：

[![复制代码](https://common.cnblogs.com/images/copycode.gif)](javascript:void(0);)

```
 1 int main(int argc, char* argv[])
 2 {
 3     int a = 5, b = 7;
 4     auto total = [](int x, int y)->int {return x + y; };    //接受两个参数
 5     cout << total(a, b)<<endl;  //12
 6     auto fun1 = [=] {return a + b; };   //值传递捕捉父作用域变量
 7     cout << fun1() << endl; //12
 8     auto fun2 = [&](int c) {b = a + c; a = 1; };    //省略了返回值类型，引用捕获所有
 9     fun2(3);    //1 8
10     cout << a <<" "<< b << endl;
11     a = 5; b = 7;   //被修改后，重新赋值
12     auto fun3 = [=, &b](int c) mutable {b = a + c; a = 1; };    //以值传递捕捉的变量，在函数体里如果要修改，要加mutaple,因为默认const修饰
13     fun3(3);
14     cout << a << " " <<b<< endl;    //5,8
15     a = 5; b = 7;   //被修改后，重新赋值
16     auto fun4 = [=](int x, int y) mutable->int {a += x; b += y; return a + b; };
17     int t = fun4(10, 20);
18     cout << t << endl;  //42
19     cout << a <<" "<< b << endl;    //5 7
20     return 0;
21 }
```

[![复制代码](https://common.cnblogs.com/images/copycode.gif)](javascript:void(0);)

　　块作用域以外的Lambda函数捕捉列表必须为空，因此这样的函数除了语法上的不同，和普通函数区别不大。

　　块作用域以内的Lambda函数仅能捕捉块作用域以内的自动变量，捕捉任何非此作用域或非自动变量(静态变量)，都会引起编译器报错。

![img]()

![img](https://img2018.cnblogs.com/blog/1475571/201907/1475571-20190713215902895-421640182.png)

　　改为引用依旧会报错。

## Lambda函数与仿函数的关系

　　在C++11之前，STL中的一些算法需要使用一种函数对象---仿函数(functor)；其本质是重新定义和成员函数operator()，使其使用上很像普通函数，其实，细心的我们已经发现，Lambda函数与仿函数似乎有一些默契。

如下例子：折扣

[![复制代码](https://common.cnblogs.com/images/copycode.gif)](javascript:void(0);)

```
 1 class Price
 2 {
 3 private:
 4     float _rate;
 5 public:
 6     Price(float rate):_rate(rate){}
 7     float operator()(float price)
 8     {
 9         return price*(1 - _rate / 100);
10     }
11 };
12  
13 int main(int argc, char* argv[])
14 {
15     float rate=5.5f;
16  
17     Price c1(rate);
18     auto c2 = [rate](float price)->float {return price*(1 - rate / 100); };
19  
20     float p1 = c1(3699);    //仿函数
21     float p2 = c2(3699);    //Lambda函数
22  
23     return 0;
24 }
```

[![复制代码](https://common.cnblogs.com/images/copycode.gif)](javascript:void(0);)

　　仿函数以rate初始化,Lambda捕捉rate变量，参数传递上，两者一致。

事实上，仿函数就是实现Lambda函数一种方式，编译器通常会把Lambda函数转换为一个放函数对象，但是仿函数的语法却给我们带来了很大的便捷。

在C++11中，Lambda函数被广泛使用，很多仿函数被取代。

## Lambda与static inline函数

　　Lambda函数可以省略外部声明的static inline函数，其相当于一个局部函数。局部函数仅属于父作用域，

比起外部的static inline函数，或者是自定义的宏，Lambda函数并没有实际运行时的性能优势(但也不会差)，但是Lambda函数可读性更好。

父函数结束后，该Lambda函数就不再可用了，不会污染任何名字空间。

## 关于值传递捕捉和mutable

　　上面提到过mutable 可以取消Lambda的常量属性，如果值传递想要在函数域内修改就要加mutable

先看一个例子：

[![复制代码](https://common.cnblogs.com/images/copycode.gif)](javascript:void(0);)

```
 1 int main(int argc, char* argv[])
 2 {
 3     int j = 12;
 4     auto by_val = [=] {return j + 1; };
 5     auto by_ref = [&] {return j + 1; };
 6     cout << by_val() << endl;   //13
 7     cout << by_ref() << endl;   //13
 8     j++;
 9     cout << by_val() << endl;   //13
10     cout << by_ref() << endl;   //14
11     return 0;
12 }
```

[![复制代码](https://common.cnblogs.com/images/copycode.gif)](javascript:void(0);)

　　上面的例子，j++了之后调用值传递结果依旧是12，原因是，值传递j被视为一个常量，一旦初始化，就不会再修改（可以认为是一个和父作用域中j同名的常量），而再引用捕捉中，j仍然是父作用域中的值。

　　其实一个值传递的的Lambda转换为放函数后，会成为一个class的常量成员函数，

代码基本如下：

[![复制代码](https://common.cnblogs.com/images/copycode.gif)](javascript:void(0);)

```
1 class const_val_lambda
2 {
3 public:
4     const_val_lambda(int v):val(v){}
5 public:
6     void operator()()const { val = 3; } //报错
7 private:
8     int val;
9 };
```

[![复制代码](https://common.cnblogs.com/images/copycode.gif)](javascript:void(0);)

　　但是使用引用的方式不会报错，因为不会改变引用本身，只会改变引用的值

　　准确地讲，现有C++11标准中的lambda等价的是有常量operatorO的仿函数。因此在使用捕捉列表的时候必须注意，按值传递方式捕捉的变量是lambda函数中不可更改的常量。标准这么设计可能是源自早期STL算法一些设计上的缺陷（对仿函数没有做限制，从而导致一些设计不算特别良好的算法出错）。而更一般地讲，这样的设计有其合理性，改变从上下文中拷贝而来的临时变量通常不具有任何意义。绝大多数时候，临时变量只是用于lambda函数的输入，如果需要输出结果到上下文，我们可以使用引用，或者通过让lambda函数返回值来实现。此外，lambda函数的mutable修饰符可以消除其常量性，不过这实际上只是提供了一种语法上的可能性，现实中应该没有多少需要使用mutable的lambda函数的地方。大多数时候，我们使用默认版本的（非mutable）的lambda函数也就足够了。

## Lambda函数与函数指针

　　Lambda函数并不是简单的函数指针类型，或者自定义类型；每个Lambda函数会产生一个闭包类型的临时对象(右值)。但是C++11允许Lambda函数向函数指针的转换，前提是：

　　　　Lambda没有捕捉任何变量

　　　　函数指针所示的函数原型，必须和Lambda有相同的调用方式

[![复制代码](https://common.cnblogs.com/images/copycode.gif)](javascript:void(0);)

```
 1 int main(int argc, char* argv[])
 2 {
 3     int a = 3, b = 4;
 4  
 5     auto total = [](int x, int y)->int {return x + y; };
 6     typedef int(*all)(int x, int y);
 7     typedef int(*one)(int x);
 8  
 9     all p;
10     p = total;
11     one q;
12     q = total;  //报错,参数不一致
13  
14     decltype(total) all_1 = total;
15     decltype(total) all_2 = p;  //报错，指针无法转换为Lambda
16  
17     return 0;
18 }
```

[![复制代码](https://common.cnblogs.com/images/copycode.gif)](javascript:void(0);)

## Lambda与STL

　　从C++11开始，Lambda被广泛用在STL中，比如foreach。与函数指针比起来，函数指针有巨大的缺陷：1.函数定义在别处，阅读起来很困难；2.使用函数指针，很可能导致编译器不对其进行inline优化，循环次数太多时，函数指针和Lambda比起来性能差距太大。函数2指针不能应用在一些运行时才能决定的状态，在没有C++11时，只能用仿函数。使得学习STL算法的代价大大降低。

　　但是Lambda并不是仿函数的完全代替者。由Lambda的捕捉列表的限制造成的，仅能捕捉副作用域的变量。放函数具有天生跨作用域共享的特征。

## Lambda 悬挂引用（Dangling references）

lambda表达式的大致原理：每当你定义一个lambda表达式后，编译器会**自动生成一个匿名类**（这个类重载了()运算符），我们称为**闭包类型（closure type）**。那么在运行时，这个**lambda表达式就会返回一个匿名的闭包实例**，是一个右值。所以，我们上面的lambda表达式的结果就是一个个闭包。**对于复制传值捕捉方式，类中会相应添加对应类型的非静态数据成员**。在运行时，会用复制的值初始化这些成员变量，从而生成闭包。对于引用捕获方式，无论是否标记mutable，都可以在lambda表达式中修改捕获的值。至于闭包类中是否有对应成员，C++标准中给出的答案是：不清楚的，与具体实现有关。

lambda表达式是不能被赋值的：

```
auto a = [] { cout << "A" << endl; };
auto b = [] { cout << "B" << endl; };

a = b;   // 非法，lambda无法赋值
auto c = a;   // 合法，生成一个副本12345
```

闭包类型禁用了赋值操作符，但是没有禁用复制构造函数，所以你仍然可以用一个lambda表达式去初始化另外一个lambda表达式而产生副本。

在多种捕获方式中，**最好不要使用[=]和[&]默认捕获所有变量**。

**默认引用捕获所有变量，你有很大可能会出现悬挂引用（Dangling references），因为引用捕获不会延长引用的变量的生命周期**：

```
std::function<int(int)> add_x(int x)
{
    return [&](int a) { return x + a; };
}1234
```

上面函数返回了一个lambda表达式，参数x仅是一个临时变量，函数add_x调用后就被销毁了，但是返回的lambda表达式却引用了该变量，当调用这个表达式时，引用的是一个垃圾值，会产生没有意义的结果。上面这种情况，使用默认传值方式可以避免悬挂引用问题。

但是采用默认值捕获所有变量仍然有风险，看下面的例子：

```
class Filter
{
public:
    Filter(int divisorVal):
        divisor{divisorVal}
    {}

    std::function<bool(int)> getFilter() 
    {
        return [=](int value) {return value % divisor == 0; };
    }

private:
    int divisor;
};123456789101112131415
```

这个类中有一个成员方法，可以返回一个lambda表达式，这个表达式使用了类的数据成员divisor。而且采用默认值方式捕捉所有变量。你可能认为这个lambda表达式也捕捉了divisor的一份副本，但是实际上并没有。**因为数据成员divisor对lambda表达式并不可见**，你可以用下面的代码验证：

```
// 类的方法，下面无法编译，因为divisor并不在lambda捕捉的范围
std::function<bool(int)> getFilter() 
{
    return [divisor](int value) {return value % divisor == 0; };
}12345
```

原代码中，lambda表达式实际上捕捉的是this指针的副本，所以原来的代码等价于：

```
std::function<bool(int)> getFilter() 
{
    return [this](int value) {return value % this->divisor == 0; };
}1234
```

尽管还是以值方式捕获，但是捕获的是指针，其实相当于以引用的方式捕获了当前类对象，所以**lambda表达式的闭包与一个类对象绑定在一起了，这很危险，因为你仍然有可能在类对象析构后使用这个lambda表达式，那么类似“悬挂引用”的问题也会产生。所以，采用默认值捕捉所有变量仍然是不安全的，主要是由于指针变量的复制，实际上还是按引用传值。**

lambda表达式可以赋值给对应类型的函数指针。但是使用函数指针并不是那么方便。所以STL定义在< functional >头文件提供了一个多态的函数对象封装std::function，其类似于函数指针。它可以绑定任何类函数对象，只要参数与返回类型相同。如下面的返回一个bool且接收两个int的函数包装器：

```
std::function<bool(int, int)> wrapper = [](int x, int y) { return x < y; };1
```



## lambda表达式实现回调函数

**lambda表达式一个更重要的应用是其可以用于函数的参数，通过这种方式可以实现回调函数**。

最常用的是在STL算法中，比如你要统计一个数组中满足特定条件的元素数量，通过lambda表达式给出条件，传递给count_if函数：

```
int value = 3;
vector<int> v {1, 3, 5, 2, 6, 10};
int count = std::count_if(v.beigin(), v.end(), [value](int x) { return x > value; });123
```

再比如你想生成斐波那契数列，然后保存在数组中，此时你可以使用generate函数，并辅助lambda表达式：

```
vector<int> v(10);
int a = 0;
int b = 1;
std::generate(v.begin(), v.end(), [&a, &b] { int value = b; b = b + a; a = value; return value; });
// 此时v {1, 1, 2, 3, 5, 8, 13, 21, 34, 55}12345
```

当需要遍历容器并对每个元素进行操作时：

```
std::vector<int> v = { 1, 2, 3, 4, 5, 6 };
int even_count = 0;
for_each(v.begin(), v.end(), [&even_count](int val){
    if(!(val & 1)){
        ++ even_count;
    }
});
std::cout << "The number of even is " << even_count << std::endl;12345678
```

大部分STL算法，可以非常灵活地搭配lambda表达式来实现想要的效果。



# C++智能指针



[详解C++11智能指针](https://www.cnblogs.com/WindSun/p/11444429.html)



智能指针是一个`RAII`（`Resource Acquisition is initialization` 资源获取就是初始化）类模型，用来动态的分配内存。它提供所有普通指针提供的接口，却很少发生异常。在构造中，它分配内存，当离开作用域时，它会自动释放已分配的内存。这样的话，程序员就从手动管理动态内存的繁杂任务中解放出来了。

**C++98**提供了第一种智能指针：**auto_ptr**



完全取决于你想要如何拥有一个资源，如果需要共享资源使用`shared_ptr`,如果独占使用资源就使用`unique_ptr`.

除此之外，`shared_ptr`比`unique_ptr`更加重，因为他还需要分配空间做其它的事，比如存储强引用计数，弱引用计数。而`unique_ptr`不需要这些，它只需要独占着保存资源对象。



## 前言

C++里面的四个智能指针: auto_ptr, unique_ptr,shared_ptr, weak_ptr 其中后三个是C++11支持，并且第一个已经被C++11弃用。

## C++11智能指针介绍

智能指针主要用于管理在堆上分配的内存，它将普通的指针封装为一个栈对象。当栈对象的生存周期结束后，会在析构函数中释放掉申请的内存，从而防止内存泄漏。C++ 11中最常用的智能指针类型为shared_ptr,它采用引用计数的方法，记录当前内存资源被多少个智能指针引用。该引用计数的内存在堆上分配。当新增一个时引用计数加1，当过期时引用计数减一。只有引用计数为0时，智能指针才会自动释放引用的内存资源。对shared_ptr进行初始化时不能将一个普通指针直接赋值给智能指针，因为一个是指针，一个是类。可以通过make_shared函数或者通过构造函数传入普通指针。并可以通过get函数获得普通指针。

## 为什么要使用智能指针

智能指针的作用是管理一个指针，因为存在以下这种情况：申请的空间在函数结束时忘记释放，造成内存泄漏。使用智能指针可以很大程度上的避免这个问题，因为智能指针是一个类，当超出了类的实例对象的作用域时，会自动调用对象的析构函数，析构函数会自动释放资源。所以智能指针的作用原理就是在函数结束时自动释放内存空间，不需要手动释放内存空间。

## auto_ptr

（C++98的方案，C++11已经抛弃）采用所有权模式。

cpp

```cpp
auto_ptr<string> p1 (new string ("I reigned lonely as a cloud.")); 
auto_ptr<string> p2; 
p2 = p1; //auto_ptr不会报错.
```

此时不会报错，p2剥夺了p1的所有权，但是当程序运行时访问p1将会报错。所以auto_ptr的缺点是：存在潜在的内存崩溃问题！

## unique_ptr

（替换auto_ptr）unique_ptr实现独占式拥有或严格拥有概念，保证同一时间内只有一个智能指针可以指向该对象。它对于避免资源泄露(例如“以new创建对象后因为发生异常而忘记调用delete”)特别有用。

采用所有权模式，还是上面那个例子

cpp

```cpp
unique_ptr<string> p3 (new string ("auto"));   //#4
unique_ptr<string> p4；                       //#5
p4 = p3;//此时会报错！！
```

编译器认为p4=p3非法，避免了p3不再指向有效数据的问题。尝试复制p3时会编译期出错，而auto_ptr能通过编译期从而在运行期埋下出错的隐患。因此，unique_ptr比auto_ptr更安全。

另外unique_ptr还有更聪明的地方：当程序试图将一个 unique_ptr 赋值给另一个时，如果源 unique_ptr 是个临时右值，编译器允许这么做；如果源 unique_ptr 将存在一段时间，编译器将禁止这么做，比如：

cpp

```cpp
unique_ptr<string> pu1(new string ("hello world")); 
unique_ptr<string> pu2; 
pu2 = pu1;                                      // #1 不允许
unique_ptr<string> pu3; 
pu3 = unique_ptr<string>(new string ("You"));   // #2 允许
```

其中#1留下悬挂的unique_ptr(pu1)，这可能导致危害。而#2不会留下悬挂的unique_ptr，因为它调用 unique_ptr 的构造函数，该构造函数创建的临时对象在其所有权让给 pu3 后就会被销毁。这种随情况而已的行为表明，unique_ptr 优于允许两种赋值的auto_ptr 。

**注：**如果确实想执行类似与#1的操作，要安全的重用这种指针，可给它赋新值。C++有一个标准库函数std::move()，让你能够将一个unique_ptr赋给另一个。尽管转移所有权后 还是有可能出现原有指针调用（调用就崩溃）的情况。但是这个语法能强调你是在转移所有权，让你清晰的知道自己在做什么，从而不乱调用原有指针。

（**额外：**boost库的boost::scoped_ptr也是一个独占性智能指针，但是它不允许转移所有权，从始而终都只对一个资源负责，它更安全谨慎，但是应用的范围也更狭窄。）

例如：

cpp

```cpp
unique_ptr<string> ps1, ps2;
ps1 = demo("hello");
ps2 = move(ps1);
ps1 = demo("alexia");
cout << *ps2 << *ps1 << endl;
```

## shared_ptr

shared_ptr实现共享式拥有概念。多个智能指针可以指向相同对象，该对象和其相关资源会在“最后一个引用被销毁”时候释放。从名字share就可以看出了资源可以被多个指针共享，它使用计数机制来表明资源被几个指针共享。可以通过成员函数use_count()来查看资源的所有者个数。除了可以通过new来构造，还可以通过传入auto_ptr, unique_ptr,weak_ptr来构造。当我们调用release()时，当前指针会释放资源所有权，计数减一。当计数等于0时，资源会被释放。

shared_ptr 是为了解决 auto_ptr 在对象所有权上的局限性(auto_ptr 是独占的), 在使用引用计数的机制上提供了可以共享所有权的智能指针。

成员函数：

use_count 返回引用计数的个数

unique 返回是否是独占所有权( use_count 为 1)

swap 交换两个 shared_ptr 对象(即交换所拥有的对象)

reset 放弃内部对象的所有权或拥有对象的变更, 会引起原有对象的引用计数的减少

get 返回内部对象(指针), 由于已经重载了()方法, 因此和直接使用对象是一样的.如

cpp

```cpp
shared_ptr<int> sp(new int(1)); 
```

sp 与 sp.get()是等价的。

share_ptr的简单例子：

cpp

```cpp
int main()
{
	string *s1 = new string("s1");

	shared_ptr<string> ps1(s1);
	shared_ptr<string> ps2;
	ps2 = ps1;

	cout << ps1.use_count()<<endl;	//2
	cout<<ps2.use_count()<<endl;	//2
	cout << ps1.unique()<<endl;	//0

	string *s3 = new string("s3");
	shared_ptr<string> ps3(s3);

	cout << (ps1.get()) << endl;	//033AEB48
	cout << ps3.get() << endl;	//033B2C50
	swap(ps1, ps3);	//交换所拥有的对象
	cout << (ps1.get())<<endl;	//033B2C50
	cout << ps3.get() << endl;	//033AEB48

	cout << ps1.use_count()<<endl;	//1
	cout << ps2.use_count() << endl;	//2
	ps2 = ps1;
	cout << ps1.use_count()<<endl;	//2
	cout << ps2.use_count() << endl;	//2
	ps1.reset();	//放弃ps1的拥有权，引用计数的减少
	cout << ps1.use_count()<<endl;	//0
	cout << ps2.use_count()<<endl;	//1
}
```

## weak_ptr

share_ptr虽然已经很好用了，但是有一点share_ptr智能指针还是有内存泄露的情况，当两个对象相互使用一个shared_ptr成员变量指向对方，会造成循环引用，使引用计数失效，从而导致内存泄漏。

weak_ptr 是一种不控制对象生命周期的智能指针, 它指向一个 shared_ptr 管理的对象. 进行该对象的内存管理的是那个强引用的shared_ptr， weak_ptr只是提供了对管理对象的一个访问手段。weak_ptr 设计的目的是为配合 shared_ptr 而引入的一种智能指针来协助 shared_ptr 工作, 它只可以从一个 shared_ptr 或另一个 weak_ptr 对象构造, 它的构造和析构不会引起引用记数的增加或减少。weak_ptr是用来解决shared_ptr相互引用时的死锁问题,如果说两个shared_ptr相互引用,那么这两个指针的引用计数永远不可能下降为0,资源永远不会释放。它是对对象的一种弱引用，不会增加对象的引用计数，和shared_ptr之间可以相互转化，shared_ptr可以直接赋值给它，它可以通过调用lock函数来获得shared_ptr。

cpp

```cpp
class B;	//声明
class A
{
public:
	shared_ptr<B> pb_;
	~A()
	{
		cout << "A delete\n";
	}
};

class B
{
public:
	shared_ptr<A> pa_;
	~B()
	{
		cout << "B delete\n";
	}
};

void fun()
{
	shared_ptr<B> pb(new B());
	shared_ptr<A> pa(new A());
	cout << pb.use_count() << endl;	//1
	cout << pa.use_count() << endl;	//1
	pb->pa_ = pa;
	pa->pb_ = pb;
	cout << pb.use_count() << endl;	//2
	cout << pa.use_count() << endl;	//2
}

int main()
{
	fun();
	return 0;
}
```

可以看到fun函数中pa ，pb之间互相引用，两个资源的引用计数为2，当要跳出函数时，智能指针pa，pb析构时两个资源引用计数会减1，但是两者引用计数还是为1，导致跳出函数时资源没有被释放（A、B的析构函数没有被调用）运行结果没有输出析构函数的内容，造成内存泄露。如果把其中一个改为weak_ptr就可以了，我们把类A里面的shared_ptr pb_，改为weak_ptr pb_ ，运行结果如下：

cpp

```cpp
1
1
1
2
B delete
A delete
```

这样的话，资源B的引用开始就只有1，当pb析构时，B的计数变为0，B得到释放，B释放的同时也会使A的计数减1，同时pa析构时使A的计数减1，那么A的计数为0，A得到释放。

注意：我们不能通过weak_ptr直接访问对象的方法，比如B对象中有一个方法print()，我们不能这样访问，pa->pb_->print()，因为pb_是一个weak_ptr，应该先把它转化为shared_ptr，如：

cpp

```cpp
shared_ptr<B> p = pa->pb_.lock();
p->print();
```

weak_ptr 没有重载*和->但可以使用 lock 获得一个可用的 shared_ptr 对象. 注意, weak_ptr 在使用前需要检查合法性.

expired 用于检测所管理的对象是否已经释放, 如果已经释放, 返回 true; 否则返回 false.

lock 用于获取所管理的对象的强引用(shared_ptr). 如果 expired 为 true, 返回一个空的 shared_ptr; 否则返回一个 shared_ptr, 其内部对象指向与 weak_ptr 相同.

use_count 返回与 shared_ptr 共享的对象的引用计数.

reset 将 weak_ptr 置空.

weak_ptr 支持拷贝或赋值, 但不会影响对应的 shared_ptr 内部对象的计数.

## share_ptr和weak_ptr的核心实现

weakptr的作为弱引用指针，其实现依赖于counter的计数器类和share_ptr的赋值，构造，所以先把counter和share_ptr简单实现

### Counter简单实现

cpp

```cpp
class Counter
{
public:
    Counter() : s(0), w(0){};
    int s;	//share_ptr的引用计数
    int w;	//weak_ptr的引用计数
};
```

counter对象的目地就是用来申请一个块内存来存引用基数，s是share_ptr的引用计数，w是weak_ptr的引用计数，当w为0时，删除Counter对象。

### share_ptr的简单实现

cpp

```cpp
template <class T>
class WeakPtr; //为了用weak_ptr的lock()，来生成share_ptr用，需要拷贝构造用

template <class T>
class SharePtr
{
public:
    SharePtr(T *p = 0) : _ptr(p)
    {
        cnt = new Counter();
        if (p)
            cnt->s = 1;
        cout << "in construct " << cnt->s << endl;
    }
    ~SharePtr()
    {
        release();
    }

    SharePtr(SharePtr<T> const &s)
    {
        cout << "in copy con" << endl;
        _ptr = s._ptr;
        (s.cnt)->s++;
        cout << "copy construct" << (s.cnt)->s << endl;
        cnt = s.cnt;
    }
    SharePtr(WeakPtr<T> const &w) //为了用weak_ptr的lock()，来生成share_ptr用，需要拷贝构造用
    {
        cout << "in w copy con " << endl;
        _ptr = w._ptr;
        (w.cnt)->s++;
        cout << "copy w  construct" << (w.cnt)->s << endl;
        cnt = w.cnt;
    }
    SharePtr<T> &operator=(SharePtr<T> &s)
    {
        if (this != &s)
        {
            release();
            (s.cnt)->s++;
            cout << "assign construct " << (s.cnt)->s << endl;
            cnt = s.cnt;
            _ptr = s._ptr;
        }
        return *this;
    }
    T &operator*()
    {
        return *_ptr;
    }
    T *operator->()
    {
        return _ptr;
    }
    friend class WeakPtr<T>; //方便weak_ptr与share_ptr设置引用计数和赋值

protected:
    void release()
    {
        cnt->s--;
        cout << "release " << cnt->s << endl;
        if (cnt->s < 1)
        {
            delete _ptr;
            if (cnt->w < 1)
            {
                delete cnt;
                cnt = NULL;
            }
        }
    }

private:
    T *_ptr;
    Counter *cnt;
};
```

share_ptr的给出的函数接口为：构造，拷贝构造，赋值，解引用，通过release来在引用计数为0的时候删除_ptr和cnt的内存。

### weak_ptr简单实现

cpp

```cpp
template <class T>
class WeakPtr
{
public: //给出默认构造和拷贝构造，其中拷贝构造不能有从原始指针进行构造
    WeakPtr()
    {
        _ptr = 0;
        cnt = 0;
    }
    WeakPtr(SharePtr<T> &s) : _ptr(s._ptr), cnt(s.cnt)
    {
        cout << "w con s" << endl;
        cnt->w++;
    }
    WeakPtr(WeakPtr<T> &w) : _ptr(w._ptr), cnt(w.cnt)
    {
        cnt->w++;
    }
    ~WeakPtr()
    {
        release();
    }
    WeakPtr<T> &operator=(WeakPtr<T> &w)
    {
        if (this != &w)
        {
            release();
            cnt = w.cnt;
            cnt->w++;
            _ptr = w._ptr;
        }
        return *this;
    }
    WeakPtr<T> &operator=(SharePtr<T> &s)
    {
        cout << "w = s" << endl;
        release();
        cnt = s.cnt;
        cnt->w++;
        _ptr = s._ptr;
        return *this;
    }
    SharePtr<T> lock()
    {
        return SharePtr<T>(*this);
    }
    bool expired()
    {
        if (cnt)
        {
            if (cnt->s > 0)
            {
                cout << "empty" << cnt->s << endl;
                return false;
            }
        }
        return true;
    }
    friend class SharePtr<T>; //方便weak_ptr与share_ptr设置引用计数和赋值
    
protected:
    void release()
    {
        if (cnt)
        {
            cnt->w--;
            cout << "weakptr release" << cnt->w << endl;
            if (cnt->w < 1 && cnt->s < 1)
            {
                //delete cnt;
                cnt = NULL;
            }
        }
    }

private:
    T *_ptr;
    Counter *cnt;
};
```

weak_ptr一般通过share_ptr来构造，通过expired函数检查原始指针是否为空，lock来转化为share_ptr。



使用make_shared<T>的好处是可以一次性分配共享对象和智能指针自身的内存。而显示地使用shared_ptr构造函数来构造则至少需要两次内存分配



## std::make_unique和std::make_shared



让我们从**std::make_unique**和**std::make_shared**之间的比较开始讲起吧。**std::make_shared**是C++11的一部分，可惜的是，**std::make_unique**不是，它在C++14才纳入标准库。如果你使用的是C++11，不用忧伤，因为**std::make_unique**的简单版本很容易写出来，不信你看：



```cpp
template<typename T, typename... Ts>
std::unique_ptr<T> make_unique(Ts&&... params)
{
    return std::unique_ptr<T>(new T(std::forward<Ts>(params)...));
}
```

就像你看到的那样，**make_unique**只是把参数完美转发给要创建对象的构造函数，再从**new**出来的原生指针构造**std::unique_ptr**，最后返回创建的**std::unique_ptr**。这种形式的函数不支持数组和自定义删除器，但它说明了只要一点点工作，你就可以创造你需要的**make_unique**了。你要记住不要把你自己的版本放入命名空间std，因为当你提升到C++14标准库实现的时候，你不会想要它和标准库的版本冲突。

**std::make_unique和std::make_shared是三个make函数中的其中两个，而make函数是：把任意集合的参数完美转发给动态分配对象的构造函数，然后返回一个指向那对象的智能指针。第三个make函数是std::allocate_shared，它的行为与std::make_shared类似，除了它第一个参数是个分配器，指定动态分配对象的方式。**

通过琐碎比较使用make函数和不使用make函数创建智能指针，揭露了使用make函数更可取的第一个原因。考虑以下：



```cpp
auto upw1(std::make_unique<Widget>());     // 使用make函数
std::unique_ptr<Widget> upw2(new Widget);  // 不使用make函数
auto spw1(std::make_shared<Widget>());     // 使用make函数
std::shared_ptr<Widget> spw2(new Widget);  // 不使用make函数
```

它们本质上的不同是：使用**new**的版本重复着需要创建的类型（即出现了两次Widget），而使用make函数不需要。重复出现类型和软件工程的关键原则产生冲突：应该避免代码重复。源码中重复的代码会增加编译时间，导致对象代码膨胀，并且通常会让代码库更难运行——这经常引发不合逻辑的代码，而不合逻辑的代码库一般会出现bug。除非是写两次比写一个更有效果，不然谁不喜欢少些点代码吗？

------

更偏爱使用make函数的第二个原因异常安全。假如我们有个函数，根据优先级来处理Widget：



```cpp
void processWidget(std::shared_ptr<Widget> spw, int priority);
```

现在呢，假如我们有个计算优先级的函数，



```cpp
int computePriority();
```

然后我们用它和**new**创建的智能指针作为参数调用processWidget：



```cpp
processWidget(std::shared_ptr<Widget>(new Widget), 



              computePriority());    // 可能会资源泄漏
```

就如注释所说，这代码中**new**出来的Widget可能会泄漏，但是为什么？**std::shared_ptr**是为了防止资源泄漏而设计的，当最后一个指向资源的**std::shared_ptr**对象消失，它们指向的资源也会被销毁。如果每个人无论什么地方都使用**std::shared_ptr**，C++还有内存泄漏这回事吗？

答案是在编译期间，源代码转换为目标码时（*.o文件）。在运行时间，函数的参数在函数运行前必须被求值，所以调用processWidget时，下面的事请会在processWidget开始前执行：

- 表达式“new Widget”会被求值，即，一个Widget对象必须在堆上被创建。
- **std::shared_ptr**的接收原生指针的构造函数一定要执行。
- computePriority一定要运行。

编译器在生成代码时不会保证上面的执行顺序，“new Widget”一定会在**std::shared_ptr**构造函数之前执行，因为构造函数需要**new**的结果，但是computePriority可能在它们之前就被调用了，可能在它们之后，可能在它们之间。所以，编译器生成代码的执行顺序有可能是这样的：

1. 执行“new Widget”。
2. 执行computePriority。
3. 执行**std::shared_ptr**的构造函数。

如果生成的代码真的是这样，那么在运行时，computePriority产生了异常，步骤1中动态分配的Widget就泄漏了，因为它没有被步骤3中的**std::shared_ptr**保存。

使用**std::make_shared_ptr**可以避免这问题。这样调用代码：



```cpp
processWidget(std::make_shared<Widget>(), computePriority())
```

在运行期间，**std::make_shared**和**computePriority**都有可能先被调用，如果先调用的是**std::make_shared**，那么指向动态分配Widget对象的原生指针会安全地存储在要返回的**std::shared_ptr**中，然后再调用computePriority。如果computePriority产出异常，**std::shared_ptr**的析构函数就会销毁持有的Widget。而如果先调用的是computePriority，并且产生异常，**std::make_shared**就不会被执行，因此没有动态分配的**Widget**对象让你担心。

如果我们把**std::shared_ptr**和**std::make_shared**替换成**std::unique_ptr**和**std::make_unique**，效果一样。使用**std::make_unique**替代**new**的重要性就像使用**std::make_shared**那样：写异常安全的代码。

------

**std::make_shared**的一个特点（相比于直接使用**new**）是提高效率。使用**std::make_shared**允许编译器生成更小、更快的代码。考虑当我们直接使用**new**时：



```cpp
std::shared_ptr<Widget> spw(new Widget);
```

很明显这代码涉及一次内存分配，不过，它实际上分配两次。每个**std::shared_ptr**内都含有一个指向控制块的指针，这控制块的内存是由**std::shared_ptr**的构造函数分配的，那么直接使用**new**，需要为Widget分配一次内存，还需要为控制块分配一次内存。

如果用**std::make_shared**呢，



```cpp
auto spw = std::make_shared<Widget>();
```

一次分配就够了，因为**std::make_shared**会分配一大块内存来同时持有Widget对象和控制块。这种优化减少了程序的静态尺寸，因为代码只需要调用一次内存分配函数，然后它增加了代码执行的速度，因为只需要分配一次内存（说明是分配内存这个函数开销略大）。而且，使用**std::make_shared**能避免了一些控制块的簿记信息，潜在地减少了程序占用的内存空间。

**std::allocate_shared**的性能分析和**std::make_shared**一样，所以**std::make_shared**的性能优势也可以延伸到**std::allocate_shared**。

------

比起直接使用**new**，更偏爱使用make函数，这个争论是很热烈的。虽有软件工程、异常安全、性能优势，不过，本条款的指导方针是更偏爱使用make函数，而不是单独依赖它们，这是因为在某些状况下它们不适用。

例如，没有一个make函数可以指定自定义删除器，但是**std::unique_ptr**和**std::shared_ptr**都有这样的构造函数。给定一个Widget的自定义删除器，



```cpp
auto widgetDeleter = [](Widget* pw) {...}
```

我们可以直接使用**new**创建智能指针：



```cpp
std::unique_ptr<Widget, decltype(widgetDeleter)>



    upw(new Widget, widgetDeleter);



 



std::shared_ptr<Widget> spw(new Widget, widgetDeleter);
```

make函数就做不来这种事情。

------

make函数的第二个限制是来源于它们实现的句法细节。当创建一个对象时，如果该对象的重载构造函数带有**std::initializer_list**参数，那么使用大括号创建对象会偏向于使用带**std::initializer_list**构造，要使用圆括号创建对象才能使用到非**std::initializer_list**构造。make函数把它们的参数完美转发给对象的构造函数，那么它们用的是大括号还是圆括号呢？对于某些类型，这问题的答案的不同会导致结果有很大差异。例如，在这些调用中，



```cpp
auto upv = std::make_unique<std::vector<int>>(10, 20);



auto spv = std::make_shared<std::vector<int>>(10, 20);
```



指针指向的是带10个元素、每个值为20的**std::vector**呢，还是指向两个元素、一个10、一个20的**std::vector**呢？还是说结果不能确定吗？

好消息是结果是能确定的：上面两个都创建内含10个值为20的**std::vector**。那意味着在make函数内，完美转发使用的是圆括号，而不是大括号。坏消息是如果你想用大括号初始化来构造指向的对象，你只能直接使用**new**，如果你想使用make函数，就要求完美转发的能力支持大括号初始化，但是大括号初始化不能被完美转发。不过也有一种能工作的方法：用**auto**推断大括号，从而创建一个**std::initializer_list**对象，然后把**auto**变量传递给make函数：



```cpp
// 创建 std::initializer_list



auto initList = {10, 20};



 



// 使用std::initializer_list构造函数创建std::vector，容器中只有两个元素



auto spv = std::make_shared<std::vector<int>>(initList);
```

对于**std::unique_ptr**，只有两种情况（自定义删除器和大括号初始化）会让它的make函数出问题。对于**std::shared_ptr**和它的make函数，就多两种情况，这两种情况都是边缘情况，不过一些开发者就喜欢住在边缘，你可能就是他们中第一个。

一些类定义了自己的**operator new**和**operator delete**函数，这些函数的出现暗示着常规的全局内存分配和回收不适合这种类型的对象。通常情况下，设计这些函数只有为了精确分配和销毁对象，例如，Widget对象的**operator new**和**operator delete**只有为了精确分配和回收大小为**sizeof(Widget)**的内存块才会设计。这两个函数不适合**std::shared_ptr**的自定义分配（借助**std::allocate_shared**）和回收（借助自定义删除器），因为**std::allocate_shared**请求内存的大小不是对象的尺寸，而是对象尺寸加上控制块尺寸。结果就是，使用make函数为那些——定义自己版本的**operator new**和**operator delete**的——类创建对象是个糟糕的想法。

------

比起直接使用**new**，**std::make_shared**的占用内存大小和速度优势来源于：**std::shared_ptr**的控制块与它管理的对象放在同一块内存。当引用计数为0时，对象被销毁（即调用了析构函数），但是，它使用的内存不会释放，除非控制块也被销毁，因为对象和控制块在同一块动态分配的内存上。

就像我提起那样，控制块上除了引用计数还有别的薄记信息。引用计数记录的是有多少**std::shared_ptr**指向控制块，但是控制块还有第二种引用计数，记录有多少**std::weak_ptr**指向控制块。这种引用计数称为**weak count**。当**std::weak_ptr**检查它是否过期时（expired），它通过检查控制块中的引用计数（不是**weak count**）来实现。如果引用计数为0（即没有**std::shared_ptr**指向这个对象，因此被销毁），**std::weak_ptr**就过期，否则就没有过期。

但是，只要有**std::weak_ptr**指向控制块（weak count大于0），控制块就必须继续存在，而只要控制块存在，容纳它的内存块也依旧存在。那么，通过make函数创建对象分配的内存，要直到最后一个指向它的**std::shared_ptr**和**std::weak_ptr**对象销毁，才能被回收。

如果对象的类型非常大，并且最后一个**std::shared_ptr**销毁和最后一个**std::weak_ptr**销毁之间的时间间隔很大，那么是对象销毁和内存被回收之间的会有延迟：



```cpp
class ReallyBigType { ... };



 



auto pBigObj =                          // 借助std::make_shared



   std::make_shared<ReallyBigType>();   // 创建类型非常大的对象



...               // 创建std::shared_ptr和std::weak_ptr指向对象 



...               // 最后一个std::shared_ptr被销毁，那仍有std::weak_ptr存在



...               // 在这个期间，之前类型非常大的对象使用的内存仍然被占用



...               // 最后一个std::weak被销毁，控制块和对象共占的内存被释放
```

如果直接使用**new**，ReallyBigType对象的内存只要在最后一个**std::shared_ptr**被销毁就能被释放：



```cpp
class ReallyBigType { ... };



 



auto pBigObj =                          // 借助std::make_shared



   std::make_shared<ReallyBigType>();   // 创建类型非常大的对象



...               // 创建std::shared_ptr和std::weak_ptr指向对象 



...               // 最后一个std::shared_ptr被销毁，那仍有std::weak_ptr存在



...               // 在这个期间，之前类型非常大的对象使用的内存仍然被占用



...               // 最后一个std::weak被销毁，控制块和对象共占的内存被释放
```

当你发现某些情况不能使用或者不适合使用**std::make_shared**，却又想要防止容易发生的异常安全问题。最好的办法就是确保当你直接使用**new**时，用一条语句执行——把**new**的结果马上传递给智能指针的构造函数，并且该语句就做这一件事。这防止编译器生成**new**和**std::shared_ptr**构造之间发出异常。

作为例子，我们修改之前的异常不安全processWidget，并指定自定义删除器：



```cpp
void processWidget(std::shared_ptr<Widget> spw, int priority); // 如前



void cusDel(Widget *ptr);      //  自定义删除器
```

这里是异常不安全的调用：



```cpp
processWidget(           // 如前，可能资源泄漏



   std::shared_ptr<Widget>(new Widget, cusDel),



   computePriority()



);
```

回忆：如果computePriority调用在“new Widget”之前，**std::shared_ptr**构造之后，然后computePriority产生异常，那么动态分配的Widget就会泄漏。

这里要使用自定义删除器，不能使用**std::make_shared**，所以避免泄漏的方法就是把分配Widget和**std::shared_ptr**构造放在只属于它们的语句，然后再用**std::shared_ptr**的结果调用processWidget。这是这项技术的本质部分，等下我们可见到，我们可以修改它从而提高性能：



```cpp
std::shared_ptr<Widget> spw(new Widget, cusDel);



processWidget(spw, computeWidget);  // 正确，但没有优化，看下面
```

这代码是可行的，因为**std::shared_ptr**得到了原生指针的所有权，尽管构造函数可能发出异常。在这个例子中，如果spw的构造期间抛出异常（例如，由于不能为控制块动态分配内存），也能保证cusDel被调用（以“new Widget”的结果为参数）。

有个小小的性能问题，在异常不安全的调用中，我们传给processWidget的是一个右值，



```cpp
processWidget(



   std::shared_ptr<Widget>(new Widget, cusDel),  // 参数是右值



   computePriority()



);
```

但是在异常安全的调用中，我们传递的是个左值：



```cpp
processWidget(spw, computePriority());   // 参数是左值
```

因为processWidget的**std::shared_ptr**参数是值传递，从一个右值构造使用的是移动，从一个左值构造使用的是拷贝。对于**std::shared_ptr**，这差别挺大的，因为拷贝一个**std::shared_ptr**需要增加它的引用计数，这是原子操作，而移动操作完全不用操作引用计数。针对于异常安全代码想要达到异常不安全代码的性能水平，我们需要使用**std::move**来把spw转化为右值：



```cpp
processWidget(std::move(spw), computePriority());  // 现在也一样高效
```

这是有趣的而且值得知道，但是通常也是不相干的，因为你很少有理由不用make函数，除非你有迫不得已的理由，否则，你应该使用make函数。

### 总结

需要记住的3点：

- 相比于直接使用**new**，make函数可以消除代码重复，提高异常安全，而且**std::make_shared**和**std::allocate_shared**生成的代码更小更快。
- 不适合使用make函数的场合包括需要指定自定义删除器和想要传递大括号初始值。
- 对于**std::shared_ptr**，使用make函数可能是不明智的额外场合包括（1）自定义内存管理函数的类和（2）内存紧张的系统中，有非常大的对象，然后**std::weak_ptr**比**std::shared_ptr**长寿。

原文链接：<http://blog.csdn.net/big_yellow_duck/article/details/52347700>



# C++线程锁

**std::unique_lock<std::mutex> lk(mtx_sync_);**

**std::lock_guard**是`RAII模板类`的简单实现，功能简单。

> 1.std::lock_guard 在构造函数中进行加锁，析构函数中进行解锁。
> 2.锁在多线程编程中，使用较多，因此c++11提供了lock_guard模板类；在实际编程中，我们也可以根据自己的场景编写`resource_guard` RAII类，避免忘掉释放资源。

**std::unique_lock**

> 类 unique_lock 是通用互斥包装器，允许`延迟锁定、锁定的有时限尝试、递归锁定、所有权转移和与条件变量一同使用`。
> unique_lock比lock_guard使用更加灵活，功能更加强大。
> 使用unique_lock需要付出更多的时间、性能成本。

**Locking/unlocking**

- ##### [**lock**](http://www.cplusplus.com/reference/mutex/unique_lock/lock/)

  Lock mutex (public member function )

- [**try_lock**](http://www.cplusplus.com/reference/mutex/unique_lock/try_lock/)

  Lock mutex if not locked (public member function )

- [**try_lock_for**](http://www.cplusplus.com/reference/mutex/unique_lock/try_lock_for/)

  Try to lock mutex during time span (public member function )

- [**try_lock_until**](http://www.cplusplus.com/reference/mutex/unique_lock/try_lock_until/)

  Try to lock mutex until time point (public member function )

- [**unlock**](http://www.cplusplus.com/reference/mutex/unique_lock/unlock/)

  Unlock mutex (public member function )

  



线程之间的锁有：**互斥锁、条件锁、自旋锁、读写锁、递归锁**。一般而言，锁的功能与性能成反比。不过我们一般不使用递归锁（C++标准库提供了std::recursive_mutex），所以这里就不推荐了。

## 互斥锁（Mutex）

互斥锁用于控制多个线程对他们之间共享资源互斥访问的一个信号量。也就是说是为了避免多个线程在某一时刻同时操作一个共享资源。例如线程池中的有多个空闲线程和一个任务队列。任何是一个线程都要使用互斥锁互斥访问任务队列，以避免多个线程同时访问任务队列以发生错乱。

在某一时刻，只有一个线程可以获取互斥锁，在释放互斥锁之前其他线程都不能获取该互斥锁。如果其他线程想要获取这个互斥锁，那么这个线程只能以阻塞方式进行等待。

头文件：< mutex >
类型： std::mutex
用法：在C++中，通过构造std::mutex的实例创建互斥元，调用成员函数lock()来锁定它，调用unlock()来解锁，不过一般不推荐这种做法，标准C++库提供了std::lock_guard类模板，实现了互斥元的RAII惯用语法。std::mutex和std::lock _ guard。都声明在< mutex >头文件中。

参考代码：

```
//用互斥元保护列表
#include <list>
#include <mutex>

std::list<int> some_list;
std::mutex some_mutex;

void add_to_list(int new_value)
{
    std::lock_guard<std::mutex> guard(some_mutex);
    some_list.push_back(new_value);
}123456789101112
```

## 条件锁

条件锁就是所谓的条件变量，某一个线程因为某个条件为满足时可以使用条件变量使改程序处于阻塞状态。一旦条件满足以“信号量”的方式唤醒一个因为该条件而被阻塞的线程。最为常见就是在线程池中，起初没有任务时任务队列为空，此时线程池中的线程因为“任务队列为空”这个条件处于阻塞状态。一旦有任务进来，就会以信号量的方式唤醒一个线程来处理这个任务。

头文件：< condition_variable >
类型：std::condition_variable（只和std::mutex一起工作） 和 std::condition_variable_any（符合类似互斥元的最低标准的任何东西一起工作）。

```
//使用std::condition_variable等待数据
std::mutex mut;
std::queue<data_chunk> data_queue;
std::condition_variable data_cond;

void data_preparation_thread()
{
    while(more_data_to_prepare())
    {
        data_chunk const data=prepare_data();
        std::lock_guard<std::mutex> lk(mut);
        data_queue.push(data);
        data_cond.notify_one();
    }
}

void data_processing_thread()
{
    while(true)
    {
        std::unique_lock<std::mutex> lk(mut);   //这里使用unique_lock是为了后面方便解锁
        data_cond.wait(lk,{[]return !data_queue.empty();});
        data_chunk data=data_queue.front();
        data_queue.pop();
        lk.unlock();
        process(data);
        if(is_last_chunk(data))
            break;
    }
}123456789101112131415161718192021222324252627282930
```

- wait()的实现接下来检查条件，并在满足时返回。如果条件不满足，wait()解锁互斥元，并将该线程置于阻塞或等待状态。当来自数据准备线程中对notify_one()的调用通知条件变量时，线程从睡眠状态中苏醒（解除其阻塞），重新获得互斥元上的锁，并再次检查条件，如果条件已经满足，就从wait()返回值，互斥元仍被锁定。如果条件不满足，该线程解锁互斥元，并恢复等待。
- 如果等待线程只打算等待一次，那么当条件为true时它就不会再等待这个条件变量了，条件变量未必是同步机制的最佳选择。如果等待的条件是一个特定数据块的可用性时，这尤其正确。在这个场景中，使用期值（future）更合适。使用future等待一次性事件。

## 自旋锁

前面的两种锁是比较常见的锁，也比较容易理解。下面通过比较互斥锁和自旋锁原理的不同，这对于真正理解自旋锁有很大帮助。

假设我们有一个两个处理器core1和core2计算机，现在在这台计算机上运行的程序中有两个线程：T1和T2分别在处理器core1和core2上运行，两个线程之间共享着一个资源。

首先我们说明互斥锁的工作原理，**互斥锁是是一种sleep-waiting的锁**。假设线程T1获取互斥锁并且正在core1上运行时，此时线程T2也想要获取互斥锁（pthread_mutex_lock），但是由于T1正在使用互斥锁使得T2被阻塞。当T2处于阻塞状态时，T2被放入到等待队列中去，处理器core2会去处理其他任务而不必一直等待（忙等）。也就是说处理器不会因为线程阻塞而空闲着，它去处理其他事务去了。

而自旋锁就不同了，**自旋锁是一种busy-waiting的锁**。也就是说，如果T1正在使用自旋锁，而T2也去申请这个自旋锁，此时T2肯定得不到这个自旋锁。与互斥锁相反的是，此时运行T2的处理器core2会一直不断地循环检查锁是否可用（自旋锁请求），直到获取到这个自旋锁为止。

从“自旋锁”的名字也可以看出来，如果一个线程想要获取一个被使用的自旋锁，那么它会一致占用CPU请求这个自旋锁使得CPU不能去做其他的事情，直到获取这个锁为止，这就是“自旋”的含义。

当发生阻塞时，互斥锁可以让CPU去处理其他的任务；而自旋锁让CPU一直不断循环请求获取这个锁。通过两个含义的对比可以我们知道“自旋锁”是比较耗费CPU的。

```
//使用std::atomic_flag的自旋锁互斥实现
class spinlock_mutex
{
    std::atomic_flag flag;
public:
spinlock_mutex():flag(ATOMIC_FLAG_INIT) {}
void lock()
{
    while(flag.test_and_set(std::memory_order_acquire));
}
void unlock()
{
    flag.clear(std::memory_order_release);
}
}123456789101112131415
```

## 读写锁

说到读写锁我们可以借助于“读者-写者”问题进行理解。首先我们简单说下“读者-写者”问题。

计算机中某些数据被多个进程共享，对数据库的操作有两种：一种是读操作，就是从数据库中读取数据不会修改数据库中内容；另一种就是写操作，写操作会修改数据库中存放的数据。因此可以得到我们允许在数据库上同时执行多个“读”操作，但是某一时刻只能在数据库上有一个“写”操作来更新数据。这就是一个简单的读者-写者模型。

头文件：boost/thread/shared_mutex.cpp
类型：boost::shared_lock

用法：你可以使用boost::shared_ mutex的实例来实现同步，而不是使用std::mutex的实例。对于更新操作，std::lock_guard< boost::shared _mutex>和 std::unique _lock< boost::shared _mutex>可用于锁定，以取代相应的std::mutex特化。这确保了独占访问，就像std::mutex那样。那些不需要更新数据结构的线程能够转而使用 boost::shared _lock< boost::shared _mutex>来获得共享访问。这与std::unique _lock用起来正是相同的，除了多个线程在同一时间，同一boost::shared _mutex上可能会具有共享锁。唯一的限制是，如果任意一个线程拥有一个共享锁，试图获取独占锁的线程会被阻塞，知道其他线程全都撤回它们的锁。同样的，如果一个线程具有独占锁，其他线程都不能获取共享锁或独占锁，直到第一个线程撤回它的锁。





# C++拷贝（复制）构造函数和重载赋值操作符

**类如果没有显式的声明以下六种函数**，编译器会自动添加（需要的时候才添加），而且都不会被派生类继承：

- 构造函数
- 析构函数
- 拷贝构造函数
- 重载赋值操作符函数
- 取址运算符重载函数
- const 取址运算符重载函数

```cpp
class A{
public:
    A();
    ~A();
    A(const A &);
    A& operator=(const A &);
    A* operator&();
    const A * operator&() const;
}
```

## 拷贝构造函数和重载操作符调用时机

如果一个对象在实例化的时候（即在声明的时候将一个已经存在的对象赋值给他)，调用的是拷贝构造函数；如果对象已经存在，将另一个对象赋值给他，调用的就是重载赋值操作符函数。在传递参数或者是函数返回值时，如果不是引用，则会调用拷贝构造函数。

```cpp
#include <iostream>
using namespace std;

class CTest
{
public:
     CTest(){}
     ~CTest(){}

     CTest(const CTest &test)
     {
          cout<<"copy constructor."<<endl;
     }

     void operator=(const CTest &test)
     {
          cout<<"operator="<<endl;
     }

     void Test(CTest test)
     {}

     CTest Test2()
     {
          CTest a;
          return a;
     }

     void Test3(CTest &test)
     {}

     CTest &Test4()
     {
          CTest *pA = new CTest;
          return *pA;
     }
};

int main()
{
     CTest obj;

     CTest obj1(obj); // 调用复制构造函数

     obj1 = obj; // 调用重载赋值操作符

     /* 传参的过程中，要调用一次复制构造函数
     * obj1入栈时会调用复制构造函数创建一个临时对象，与函数内的局部变量具有相同的作用域
     */
     obj.Test(obj1);

     /* 函数返回值时，调用复制构造函数；将返回值赋值给obj2时，调用重载赋值操作符
     * 函数返回值时，也会构造一个临时对象；调用复制构造函数将返回值复制到临时对象上
     */
     CTest obj2;
     obj2 = obj.Test2();

     obj2.Test3(obj); // 参数是引用，没有调用复制构造函数

     CTest obj3;
     obj2.Test4(); // 返回值是引用，没有调用复制构造函数

     return 0;
}
```

## 重载操作符

- 带点的操作符不可重载： `.`， `.*`， `::`， `?:`
- 重载`-=`, `+=`, `=` 的时，返回值应该为引用。

## 函数调用顺序

- 调用基类构造函数；多继承的时候按照派生列表顺序。
- 成员变量声明的顺序初始化，而不是初始化列表
- 构造函数



# c++ 之 std::move 原理实现与用法总结

在C++11中，标准库在<utility>中提供了一个有用的函数std::move，std::move并不能移动任何东西，它唯一的功能是将一个左值强制转化为右值引用，继而可以通过右值引用使用该值，以用于移动语义。从实现上讲，std::move基本等同于一个类型转换：static_cast<T&&>(lvalue);

std::move函数可以以非常简单的方式将左值引用转换为右值引用。(左值 右值 引用 左值引用)概念 <https://blog.csdn.net/p942005405/article/details/84644101>

1. C++ 标准库使用比如vector::push_back 等这类函数时,会对参数的对象进行复制,连数据也会复制.这就会造成对象内存的额外创建, 本来原意是想把参数push_back进去就行了,通过std::move，可以避免不必要的拷贝操作。
2. std::move是将对象的状态或者所有权从一个对象转移到另一个对象，只是转移，没有内存的搬迁或者内存拷贝所以可以提高利用效率,改善性能.。
3. 对指针类型的标准库对象并不需要这么做.

## **用法:**

原lvalue值被moved from之后值被转移,所以为空字符串. 

```cpp
//摘自https://zh.cppreference.com/w/cpp/utility/move
#include <iostream>
#include <utility>
#include <vector>
#include <string>
int main()
{
    std::string str = "Hello";
	std::vector<std::string> v;
    //调用常规的拷贝构造函数，新建字符数组，拷贝数据
    v.push_back(str);
    std::cout << "After copy, str is \"" << str << "\"\n";
    //调用移动构造函数，掏空str，掏空后，最好不要使用str
    v.push_back(std::move(str));
    std::cout << "After move, str is \"" << str << "\"\n";
    std::cout << "The contents of the vector are \"" << v[0]
				<< "\", \"" << v[1] << "\"\n";
}
```

输出:

```cpp
After copy, str is "Hello"
After move, str is ""
The contents of the vector are "Hello", "Hello"
```

## [左值、左值引用、右值、右值引用](https://www.cnblogs.com/SZxiaochun/p/8017475.html)



### 1、左值和右值的概念

​         左值是可以放在赋值号左边可以被赋值的值；左值必须要在内存中有实体；
​         右值当在赋值号右边取出值赋给其他变量的值；右值可以在内存也可以在CPU寄存器。
​         一个对象被用作右值时，使用的是它的内容(值)，被当作左值时，使用的是它的地址**。**

### 2、引用

​        引用是C++语法做的优化，引用的本质还是靠指针来实现的。引用相当于变量的别名。

​        引用可以改变指针的指向，还可以改变指针所指向的值。

​        引用的基本规则：

1. 声明引用的时候必须初始化，且一旦绑定，不可把引用绑定到其他对象；即引用必须初始化，不能对引用重定义**；**
2. 对引用的一切操作，就相当于对原对象的操作。

### 3、左值引用和右值引用

​    3.1 左值引用
​         左值引用的基本语法：type &引用名 = 左值表达式；

​    3.2 右值引用

​        右值引用的基本语法type &&引用名 = 右值表达式；

​        右值引用在企业开发人员在代码优化方面会经常用到。

​        右值引用的“&&”中间不可以有空格。



# C++11 std::bind()，std::function，std::ref()，std::cref

### std::bind

std::bind可以预先把某个可调用对象的某些参数绑定到已有的变量，产生一个新的可调用对象。bind本身是一种延迟计算的思想，它本身可以绑定普通函数、全局函数、静态函数、类静态函数甚至是类成员函数。bind最终将会生成一个可调用对象，这个对象可以直接赋值给std::function对象，而std::bind绑定的可调用对象可以是Lambda表达式或者类成员函数等可调用对象。

使用方式：bind（&要调用的函数，&对象， 要调用函数的参数1，要调用函数的参数2…，_1(bind函数的参数1)，_2(bind函数的参数2)…)

```cpp
#include <iostream>
#include <functional>
using namespace std;
 
int TestFunc(int a, char c, float f)
{
    cout << a << endl;
    cout << c << endl;
    cout << f << endl;
 
    return a;
}
 
int main()
{
    auto bindFunc1 = bind(TestFunc, placeholders::_1, 'A', 10.1); //表示绑定函数TestFunc的第二、三个参数分别为'A'和10.1，第一个参数由调用bindFunc1的第一个参数来指定
    bindFunc1(10);　　　//输出为：10  A  10.1
 	bindFunc1(10,'C'); //输出为：10  A  10.1,已经绑定的值无法通过传参改变
    cout << "=================================\n";
 
    auto bindFunc2 = bind(TestFunc, std::placeholders::_2, placeholders::_1, 10.1);
    bindFunc2('B', 10);   //输出为：10  B  10.1
 
    cout << "=================================\n";
 
    auto bindFunc3 = bind(TestFunc, std::placeholders::_1, placeholders::_2, 10.1);
    bindFunc3(10, 'B');   //输出为：10  B  10.1
 
    return 0;
}
123456789101112131415161718192021222324252627282930
```

其中placeholders::_1是占位符，用于接收以后的传参。如第二个bindFunc2 函数来说，std::placeholders::_2, 表示现在的第一个入参传入原函数的第二个参数。

**注：bind对于不事先绑定的参数，通过std::placeholders传递的参数是通过引用传递的；对于事先绑定的函数参数是通过值传**，可以通过std::ref();std::cref来绑定函数参数通过引用传递

bind使用形式：

（1）bind（&f ，arg_1，…) 假设f是一个全局函数，绑定全局函数并调用；后面f后面跟参数事先绑定，不事先绑定的参数，通过std::placeholders传递。

（2）bind (&A::f, A()，arg_1，…) 假设A是一个构造函数为空的类，这个形式绑定了类的成员函数，故第二个参数需要传入一个成员
（成员静态函数除外）；

（3）bind (&A::f, _1，arg_1，…)(new A()) 同上，效果是一样的，但是使用了占位符，使得没有固定的的对象，推荐。

**注：使用的时候一定要注意指向的是没有this指针的函数（全局函数或静态成员函数），还是有this指针的函数。后面一种必须要用bind()函数，而且要多一个参数，因为静态成员函数与非静态成员函数的参 数表不一样，原型相同的非静态函数比静态成员函数多一个参数，即第一个参数this指针，指向所属的对象，任何非静态成员函数的第一个参数都是this指针。**

### std::function

可以理解为函数指针。C++中把函数指针封装成了一个类,这也正是C++中无处不类的思想的体现,即
std::function,且是模板类。
看实例就比较清楚的了解了

```cpp
#include <functional>
#include <iostream>
using namespace std;
 
std::function< int(int)> Functional;
 
// 普通函数
int TestFunc(int a)
{
    return a;
}
 
// Lambda表达式
auto lambda = [](int a)->int{ return a; };
 
// 仿函数(functor)
class Functor
{
public:
    int operator()(int a)
    {
        return a;
    }
};
 
// 1.类成员函数
// 2.类静态函数
class TestClass
{
public:
    int ClassMember(int a) { return a; }
    static int StaticMember(int a) { return a; }
};
 
int main()
{
    // 普通函数
    Functional = TestFunc;
    int result = Functional(10);
    cout << "普通函数："<< result << endl;
 
    // Lambda表达式
    Functional = lambda;
    result = Functional(20);
    cout << "Lambda表达式："<< result << endl;
 
    // 仿函数
    Functor testFunctor;
    Functional = testFunctor;
    result = Functional(30);
    cout << "仿函数："<< result << endl;
 
    // 类成员函数
    TestClass testObj;
    Functional = std::bind(&TestClass::ClassMember, testObj, std::placeholders::_1);
    result = Functional(40);
    cout << "类成员函数："<< result << endl;
 
    // 类静态函数
    Functional = TestClass::StaticMember;
    result = Functional(50);
    cout << "类静态函数："<< result << endl;
 
    return 0;
}
1234567891011121314151617181920212223242526272829303132333435363738394041424344454647484950515253545556575859606162636465
```

### std::ref();std::cref

std::ref 用于包装按引用传递的值。
std::cref 用于包装按const引用传递的值。

在bind中已经知道事先绑定的值是不使用引用而是传参的，因为bind()不知道生成的函数执行的时候，传递进来的参数是否还有效。std::ref和std::cref就派上用场了。

```cpp
#include <functional>
#include <iostream>

void f(int& n1, int& n2, const int& n3)
{
    std::cout << "In function: n1[" << n1 << "]    n2[" << n2 << "]    n3[" << n3 << "]" << std::endl;
    ++n1; // 增加存储于函数对象的 n1 副本
    ++n2; // 增加 main() 的 n2
    //++n3; // 编译错误
    std::cout << "In function end: n1[" << n1 << "]     n2[" << n2 << "]     n3[" << n3 << "]" << std::endl;
}

int main()
{
    int n1 = 1, n2 = 1, n3 = 1;
    std::cout << "Before function: n1[" << n1 << "]     n2[" << n2 << "]     n3[" << n3 << "]" << std::endl;
    std::function<void()> bound_f = std::bind(f, n1, std::ref(n2), std::cref(n3));
    bound_f();
    std::cout << "After function: n1[" << n1 << "]     n2[" << n2 << "]     n3[" << n3 << "]" << std::endl;
}
1234567891011121314151617181920
```

运行结果

```cpp
Before function: n1[1]     n2[1]     n3[1]
In function: n1[1]    n2[1]    n3[1]
In function end: n1[2]     n2[2]     n3[1]
After function: n1[1]     n2[2]     n3[1]
```



# C++11新特性之十：enable_shared_from_this



[灿哥哥](https://me.csdn.net/caoshangpa) 2018-02-27 20:43:06 ![img](https://csdnimg.cn/release/blogv2/dist/pc/img/articleReadEyes.png) 48197 ![img](https://csdnimg.cn/release/blogv2/dist/pc/img/tobarCollect.png) 收藏 96

分类专栏： [C++11](https://blog.csdn.net/caoshangpa/category_6465272.html) 文章标签： [C 11](https://so.csdn.net/so/search/s.do?q=C  11&t=blog&o=vip&s=&l=&f=&viparticle=) [shared_from_this](https://so.csdn.net/so/search/s.do?q=shared_from_this&t=blog&o=vip&s=&l=&f=&viparticle=)

版权

​       enable_shared_from_this是一个模板类，定义于头文件<memory>，其原型为：

```cpp
template< class T > class enable_shared_from_this;
```

​       std::enable_shared_from_this 能让一个对象（假设其名为 t ，且已被一个 std::shared_ptr 对象 pt 管理）安全地生成其他额外的 std::shared_ptr 实例（假设名为 pt1, pt2, ... ） ，它们与 pt 共享对象 t 的所有权。

​       若一个类 T 继承 std::enable_shared_from_this<T> ，则会为该类 T 提供成员函数： shared_from_this 。 当 T 类型对象 t 被一个为名为 pt 的 std::shared_ptr<T> 类对象管理时，调用 T::shared_from_this 成员函数，将会返回一个新的 std::shared_ptr<T> 对象，它与 pt 共享 t 的所有权。

**一.使用场合**

​       当类A被share_ptr管理，且在类A的成员函数里需要把当前类对象作为参数传给其他函数时，就需要传递一个指向自身的share_ptr。

1.为何不直接传递this指针

​       使用智能指针的初衷就是为了方便资源管理，如果在某些地方使用智能指针，某些地方使用原始指针，很容易破坏智能指针的语义，从而产生各种错误。

2.可以直接传递share_ptr<this>么？

​       答案是不能，因为这样会造成2个非共享的share_ptr指向同一个对象，未增加引用计数导对象被析构两次。例如：

```cpp
#include <memory>



#include <iostream>



 



class Bad



{



public:



	std::shared_ptr<Bad> getptr() {



		return std::shared_ptr<Bad>(this);



	}



	~Bad() { std::cout << "Bad::~Bad() called" << std::endl; }



};



 



int main()



{



	// 错误的示例，每个shared_ptr都认为自己是对象仅有的所有者



	std::shared_ptr<Bad> bp1(new Bad());



	std::shared_ptr<Bad> bp2 = bp1->getptr();



	// 打印bp1和bp2的引用计数



	std::cout << "bp1.use_count() = " << bp1.use_count() << std::endl;



	std::cout << "bp2.use_count() = " << bp2.use_count() << std::endl;



}  // Bad 对象将会被删除两次
```

输出结果如下：

![img](https://img-blog.csdn.net/20180228102531171?watermark/2/text/aHR0cDovL2Jsb2cuY3Nkbi5uZXQvY2Fvc2hhbmdwYQ==/font/5a6L5L2T/fontsize/400/fill/I0JBQkFCMA==/dissolve/70/gravity/SouthEast)
当然，一个对象被删除两次会导致崩溃。

![img](https://img-blog.csdn.net/20180228102640408?watermark/2/text/aHR0cDovL2Jsb2cuY3Nkbi5uZXQvY2Fvc2hhbmdwYQ==/font/5a6L5L2T/fontsize/400/fill/I0JBQkFCMA==/dissolve/70/gravity/SouthEast)

正确的实现如下：

```cpp
#include <memory>



#include <iostream>



 



struct Good : std::enable_shared_from_this<Good> // 注意：继承



{



public:



	std::shared_ptr<Good> getptr() {



		return shared_from_this();



	}



	~Good() { std::cout << "Good::~Good() called" << std::endl; }



};



 



int main()



{



	// 大括号用于限制作用域，这样智能指针就能在system("pause")之前析构



	{



		std::shared_ptr<Good> gp1(new Good());



		std::shared_ptr<Good> gp2 = gp1->getptr();



		// 打印gp1和gp2的引用计数



		std::cout << "gp1.use_count() = " << gp1.use_count() << std::endl;



		std::cout << "gp2.use_count() = " << gp2.use_count() << std::endl;



	}



	system("pause");



} 
```

输出结果如下：

![img](https://img-blog.csdn.net/20180228103637404?watermark/2/text/aHR0cDovL2Jsb2cuY3Nkbi5uZXQvY2Fvc2hhbmdwYQ==/font/5a6L5L2T/fontsize/400/fill/I0JBQkFCMA==/dissolve/70/gravity/SouthEast)

**二.为何会出现这种使用场合**

​       因为在异步调用中，存在一个保活机制，异步函数执行的时间点我们是无法确定的，然而异步函数可能会使用到异步调用之前就存在的变量。为了保证该变量在异步函数执期间一直有效，我们可以传递一个指向自身的share_ptr给异步函数，这样在异步函数执行期间share_ptr所管理的对象就不会析构，所使用的变量也会一直有效了（保活）。



# C++中Overload、Overwrite及Override的区别



*Overload(*重载*)*：在*C++*程序中，可以将语义、功能相似的几个函数用同一个名字表示，但参数或返回值不同（包括类型、顺序不同），即函数重载。（*1*）相同的范围（在同一个类中）；（*2*）函数名字相同；（*3*）参数不同；（*4*）*virtual* 关键字可有可无。

*Override(*覆盖*)*：是指派生类函数覆盖基类函数，特征是：（*1*）不同的范围（分别位于派生类与基类）；（*2*）函数名字相同；（*3*）参数相同；（*4*）基类函数必须有*virtual* 关键字。

*Overwrite(*重写*)*：是指派生类的函数屏蔽了与其同名的基类函数，规则如下：（*1*）如果派生类的函数与基类的函数同名，但是参数不同。此时，不论有无*virtual*关键字，基类的函数将被隐藏（注意别与重载混淆）。（*2*）如果派生类的函数与基类的函数同名，并且参数也相同，但是基类函数没有*virtual*关键字。此时，基类的函数被隐藏（注意别与覆盖混淆）。



# constexpr



constexpr 是 C++11 引入的，一方面是为了引入更多的编译时计算能力，另一方面也是解决 C++98 的 const 的双重语义问题。

在 C 里面，const 很明确只有**「只读」**一个语义，不会混淆。C++ 在此基础上增加了**「常量」**语义，也由 const 关键字来承担，引出来一些奇怪的问题。C++11 把「常量」语义拆出来，交给新引入的 constexpr 关键字。



在 C++11 以后，建议凡是「常量」语义的场景都使用 constexpr，只对「只读」语义使用 const。



constexpr是一种比const 更严格的束缚, 它修饰的**表达式本身**在编译期间可知, 并且编译器会尽可能的 evaluate at compile time. 在constexpr 出现之前, 可以在编译期初始化的const都是implicit constexpr. 直到c++ 11, constexpr才从const中细分出来成为一个关键字, 而 const从1983年 c++ 刚改名的时候就存在了... 如果你初学c++, 应当尽可能的, 合理的使用constexpr来帮助编译器优化代码.

加个cppreference上的简单例子

```cpp
// constexpr 声明factorial可以参与编译期的运算
constexpr int factorial(int n)
{
    return n <= 1? 1 : (n * factorial(n - 1));
}
int main()
{
    std::cout << "4! = " ;
    constN<factorial(4)> out1; // computed at compile time
 
    volatile int k = 8; // disallow optimization using volatile
    std::cout << k << "! = " << factorial(k) << '\n'; // computed at run time
}
```





# 支持 C++11/14/17 功能（现代 C++）



若要了解有关 Visual Studio 2017 RC 的最新文档，请参阅 [Visual Studio 2017 RC 文档](http://docs.microsoft.com/visualstudio)。

本文描述了 Visual C++ 中的 C++11/14/17 功能。

## 本文内容



- [C++11 功能列表](https://msdn.microsoft.com/zh-cn/library/hh567368.aspx#featurelist)
  - [C++11 核心语言功能表](https://msdn.microsoft.com/zh-cn/library/hh567368.aspx#corelanguagetable)
  - [C++11 核心语言功能表：并发](https://msdn.microsoft.com/zh-cn/library/hh567368.aspx#concurrencytable)
  - [C++11 核心语言功能：C99](https://msdn.microsoft.com/zh-cn/library/hh567368.aspx#c99table)
- [C++ 14 核心语言功能](https://msdn.microsoft.com/zh-cn/library/hh567368.aspx#cpp14table)
- [C++17 建议的核心语言功能](https://msdn.microsoft.com/zh-cn/library/hh567368.aspx#cpp17table)
- [功能表指南](https://msdn.microsoft.com/zh-cn/library/hh567368.aspx#tableguide)
  - [右值引用](https://msdn.microsoft.com/zh-cn/library/hh567368.aspx#rvref)
  - [Lambdas](https://msdn.microsoft.com/zh-cn/library/hh567368.aspx#lambdas)
  - [decltype](https://msdn.microsoft.com/zh-cn/library/hh567368.aspx#decltype)
  - [强类型/前向声明枚举](https://msdn.microsoft.com/zh-cn/library/hh567368.aspx#stronglytyped)
  - [对齐方式](https://msdn.microsoft.com/zh-cn/library/hh567368.aspx#alignment)
  - [标准布局和普通类型](https://msdn.microsoft.com/zh-cn/library/hh567368.aspx#standardlayout)
  - [默认函数和已删除的函数](https://msdn.microsoft.com/zh-cn/library/hh567368.aspx#defaultedanddeleted)
  - [override 和 final](https://msdn.microsoft.com/zh-cn/library/hh567368.aspx#overrideandfinal)
  - [原子化及更多信息](https://msdn.microsoft.com/zh-cn/library/hh567368.aspx#atomics)
  - [C99 **func** 和预处理器规则](https://msdn.microsoft.com/zh-cn/library/hh567368.aspx#c99)
- [标准库功能](https://msdn.microsoft.com/zh-cn/library/hh567368.aspx#stl)

## C++11 功能列表



Visual C++ 实现了 [C++11 核心语言规范](http://go.microsoft.com/fwlink/p/?LinkID=235092) 中的绝大多数功能、许多 C++14 库功能和某些为 C++17 建议的功能。 下表列出了 C++11/14/17 核心语言功能及其在 Visual Studio 2010、Visual Studio 2012 中的 Visual C++、Visual Studio 2013 中的 Visual C++ 和 Visual Studio 2015 中 Visual C++ 中的实现状态。

### C++11 核心语言功能表

| [C++11 核心语言功能](http://www.open-std.org/jtc1/sc22/wg21/docs/papers/2009/n2869.html) | Visual Studio 2010 | Visual Studio 2012 | Visual Studio 2013                                           | Visual Studio 2015 |
| :----------------------------------------------------------- | :----------------- | :----------------- | :----------------------------------------------------------- | :----------------- |
| 右值引用 [v0.1](http://www.open-std.org/jtc1/sc22/wg21/docs/papers/2004/n1610.html)、[v1.0](http://www.open-std.org/jtc1/sc22/wg21/docs/papers/2006/n2118.html)、[v2.0](http://www.open-std.org/jtc1/sc22/wg21/docs/papers/2009/n2844.html)、[v2.1](http://www.open-std.org/jtc1/sc22/wg21/docs/cwg_defects.html)、[v3.0](http://www.open-std.org/jtc1/sc22/wg21/docs/papers/2010/n3053.html) | 2.0 版             | 2.1* 版            | 2.1* 版                                                      | v3.0               |
| [引用限定符](http://www.open-std.org/jtc1/sc22/wg21/docs/papers/2007/n2439.htm) | 否                 | 否                 | 否                                                           | 是                 |
| [非静态数据成员初始值设定项](http://www.open-std.org/jtc1/sc22/wg21/docs/papers/2008/n2756.htm) | 否                 | 否                 | [是](https://msdn.microsoft.com/zh-cn/library/dn387583.aspx) | 是                 |
| 可变参数模板 [v0.9](http://www.open-std.org/jtc1/sc22/wg21/docs/papers/2007/n2242.pdf)、[v1.0](http://www.open-std.org/jtc1/sc22/wg21/docs/papers/2008/n2555.pdf) | 否                 | 否                 | [是](https://msdn.microsoft.com/zh-cn/library/dn439779.aspx) | 是                 |
| [初始值设定项列表](http://www.open-std.org/jtc1/sc22/wg21/docs/papers/2008/n2672.htm) | 否                 | 否                 | [是](https://msdn.microsoft.com/zh-cn/library/dn387583.aspx) | 是                 |
| [static_assert](http://www.open-std.org/jtc1/sc22/wg21/docs/papers/2004/n1720.html) | 是                 | 是                 | 是                                                           | 是                 |
| auto [v0.9](http://www.open-std.org/jtc1/sc22/wg21/docs/papers/2006/n1984.pdf)、[v1.0](http://www.open-std.org/jtc1/sc22/wg21/docs/papers/2008/n2546.htm) | v1.0               | v1.0               | v1.0                                                         | 是                 |
| [结尾返回类型](http://www.open-std.org/jtc1/sc22/wg21/docs/papers/2008/n2541.htm) | 是                 | 是                 | 是                                                           | 是                 |
| Lambdas [v0.9](http://www.open-std.org/jtc1/sc22/wg21/docs/papers/2008/n2550.pdf)、[v1.0](http://www.open-std.org/jtc1/sc22/wg21/docs/papers/2008/n2658.pdf)、[v1.1](http://www.open-std.org/jtc1/sc22/wg21/docs/papers/2009/n2927.pdf) | v1.0               | v1.1               | v1.1                                                         | 是                 |
| decltype [v1.0](http://www.open-std.org/jtc1/sc22/wg21/docs/papers/2007/n2343.pdf)、[v1.1](http://www.open-std.org/jtc1/sc22/wg21/docs/papers/2011/n3276.pdf) | v1.0               | v1.1**             | v1.1                                                         | 是                 |
| [右尖括号](http://www.open-std.org/jtc1/sc22/wg21/docs/papers/2005/n1757.html) | 是                 | 是                 | 是                                                           | 是                 |
| [函数模板的默认模板参数](http://www.open-std.org/jtc1/sc22/wg21/docs/cwg_defects.html) | 否                 | 否                 | 是                                                           | 是                 |
| [表达式 SFINAE](http://www.open-std.org/jtc1/sc22/wg21/docs/papers/2008/n2634.html) | 否                 | 否                 | 否                                                           | 否                 |
| [别名模板](http://www.open-std.org/jtc1/sc22/wg21/docs/papers/2007/n2258.pdf) | 否                 | 否                 | [是](https://msdn.microsoft.com/zh-cn/library/dn467695.aspx) | 是                 |
| [Extern 模板](http://www.open-std.org/jtc1/sc22/wg21/docs/papers/2006/n1987.htm) | 是                 | 是                 | 是                                                           | 是                 |
| [nullptr](http://www.open-std.org/jtc1/sc22/wg21/docs/papers/2007/n2431.pdf) | 是                 | 是                 | 是                                                           | 是                 |
| [强类型的枚举](http://www.open-std.org/jtc1/sc22/wg21/docs/papers/2007/n2347.pdf) | 部分               | 是                 | 是                                                           | 是                 |
| [前向声明枚举](http://www.open-std.org/jtc1/sc22/wg21/docs/papers/2008/n2764.pdf) | 否                 | 是                 | 是                                                           | 是                 |
| [特性](http://www.open-std.org/jtc1/sc22/wg21/docs/papers/2008/n2761.pdf) | 否                 | 否                 | 否                                                           | 是                 |
| [constexpr](http://www.open-std.org/jtc1/sc22/wg21/docs/papers/2007/n2235.pdf) | 否                 | 否                 | 否                                                           | 是                 |
| [对齐方式](http://www.open-std.org/jtc1/sc22/wg21/docs/papers/2007/n2341.pdf) | TR1                | 部分               | 部分                                                         | 是                 |
| [委托构造函数](http://www.open-std.org/jtc1/sc22/wg21/docs/papers/2006/n1986.pdf) | 否                 | 否                 | [是](https://msdn.microsoft.com/zh-cn/library/dn387583.aspx) | 是                 |
| [继承构造函数](http://www.open-std.org/jtc1/sc22/wg21/docs/papers/2008/n2540.htm) | 否                 | 否                 | 否                                                           | 是                 |
| [显式转换运算符](http://www.open-std.org/jtc1/sc22/wg21/docs/papers/2007/n2437.pdf) | 否                 | 否                 | 是                                                           | 是                 |
| [char16_t/char32_t](http://www.open-std.org/jtc1/sc22/wg21/docs/papers/2007/n2249.html) | 否                 | 否                 | 否                                                           | 是                 |
| [Unicode 字符串文本](http://www.open-std.org/jtc1/sc22/wg21/docs/papers/2007/n2442.htm) | 否                 | 否                 | 否                                                           | 是                 |
| [原始字符串文本](http://www.open-std.org/jtc1/sc22/wg21/docs/papers/2007/n2442.htm) | 否                 | 否                 | [是](https://msdn.microsoft.com/zh-cn/library/69ze775t.aspx) | 是                 |
| [文本中的通用字符名](http://www.open-std.org/jtc1/sc22/wg21/docs/papers/2007/n2170.html) | 否                 | 否                 | 否                                                           | 是                 |
| [用户定义的文本](http://www.open-std.org/jtc1/sc22/wg21/docs/papers/2008/n2765.pdf) | 否                 | 否                 | 否                                                           | 是                 |
| [标准布局和普通类型](http://www.open-std.org/jtc1/sc22/wg21/docs/papers/2007/n2342.htm) | 否                 | 是                 | 是                                                           | 是                 |
| [默认函数和已删除的函数](http://www.open-std.org/jtc1/sc22/wg21/docs/papers/2007/n2346.htm) | 否                 | 否                 | [是*](https://msdn.microsoft.com/zh-cn/library/dn457344.aspx) | 是                 |
| [扩展的友元声明](http://www.open-std.org/jtc1/sc22/wg21/docs/papers/2005/n1791.pdf) | 是                 | 是                 | 是                                                           | 是                 |
| [扩展的 sizeof](http://www.open-std.org/jtc1/sc22/wg21/docs/papers/2007/n2253.html) | 否                 | 否                 | 否                                                           | 是                 |
| [内联命名空间](http://www.open-std.org/jtc1/sc22/wg21/docs/papers/2008/n2535.htm) | 否                 | 否                 | 否                                                           | 是                 |
| [无限制的联合](http://www.open-std.org/jtc1/sc22/wg21/docs/papers/2008/n2544.pdf) | 否                 | 否                 | 否                                                           | 是                 |
| [作为模板参数的本地和未命名类型](http://www.open-std.org/jtc1/sc22/wg21/docs/papers/2008/n2657.htm) | 是                 | 是                 | 是                                                           | 是                 |
| [基于范围的 for 循环](http://www.open-std.org/jtc1/sc22/wg21/docs/papers/2009/n2930.html) | 否                 | 是                 | 是                                                           | 是                 |
| override 和 final [v0.8](http://www.open-std.org/jtc1/sc22/wg21/docs/papers/2009/n2928.htm)、[v0.9](http://www.open-std.org/jtc1/sc22/wg21/docs/papers/2010/n3206.htm)、[v1.0](http://www.open-std.org/jtc1/sc22/wg21/docs/papers/2011/n3272.htm) | 部分               | 是                 | 是                                                           | 是                 |
| [最低 GC 支持](http://www.open-std.org/jtc1/sc22/wg21/docs/papers/2008/n2670.htm) | 是                 | 是                 | 是                                                           | 是                 |
| [noexcept](http://www.open-std.org/jtc1/sc22/wg21/docs/papers/2010/n3050.html) | 否                 | 否                 | 否                                                           | 是                 |

[[本文内容](https://msdn.microsoft.com/zh-cn/library/hh567368.aspx#top)]

### C++11 核心语言功能表：并发

| C++11 核心语言功能：并发                                     | Visual Studio 2010 | Visual Studio 2012 | Visual Studio 2013 | Visual Studio 2015 |
| :----------------------------------------------------------- | :----------------- | :----------------- | :----------------- | :----------------- |
| [改写的序列点](http://www.open-std.org/jtc1/sc22/wg21/docs/papers/2007/n2239.html) | 不可用             | 不可用             | 不可用             | 是                 |
| [原子](http://www.open-std.org/jtc1/sc22/wg21/docs/papers/2007/n2427.html) | 否                 | 是                 | 是                 | 是                 |
| [强比较和交换](http://www.open-std.org/jtc1/sc22/wg21/docs/papers/2008/n2748.html) | 否                 | 是                 | 是                 | 是                 |
| [双向界定](http://www.open-std.org/jtc1/sc22/wg21/docs/papers/2008/n2752.htm) | 否                 | 是                 | 是                 | 是                 |
| [内存模型](http://www.open-std.org/jtc1/sc22/wg21/docs/papers/2007/n2429.htm) | 不可用             | 不可用             | 不可用             | 是                 |
| [数据依赖项排序](http://www.open-std.org/jtc1/sc22/wg21/docs/papers/2008/n2664.htm) | 否                 | 是                 | 是                 | 是                 |
| [数据依赖项排序：函数批注](http://www.open-std.org/jtc1/sc22/wg21/docs/papers/2008/n2782.htm) | 否                 | 否                 | 否                 | 是                 |
| [exception_ptr](http://www.open-std.org/jtc1/sc22/wg21/docs/papers/2007/n2179.html) | 是                 | 是                 | 是                 | 是                 |
| [quick_exit](http://www.open-std.org/jtc1/sc22/wg21/docs/papers/2007/n2440.htm) | 否                 | 否                 | 否                 | 是                 |
| [信号处理程序中的原子化](http://www.open-std.org/jtc1/sc22/wg21/docs/papers/2008/n2547.htm) | 否                 | 否                 | 否                 | 否                 |
| [线程本地存储区](http://www.open-std.org/jtc1/sc22/wg21/docs/papers/2008/n2659.htm) | 部分               | 部分               | 部分               | 是                 |
| [神奇的静态对象](http://www.open-std.org/jtc1/sc22/wg21/docs/papers/2008/n2660.htm) | 否                 | 否                 | 否                 | 是                 |

[[本文内容](https://msdn.microsoft.com/zh-cn/library/hh567368.aspx#top)]

### C++11 核心语言功能：C99

| C++11 核心语言功能：C99                                      | Visual Studio 2010 | Visual Studio 2012 | Visual Studio 2013 | Visual Studio 2015 |
| :----------------------------------------------------------- | :----------------- | :----------------- | :----------------- | :----------------- |
| [__func__](http://www.open-std.org/jtc1/sc22/wg21/docs/papers/2007/n2340.htm) | 部分               | 部分               | 部分               | 是                 |
| [C99 预处理器](http://www.open-std.org/jtc1/sc22/wg21/docs/papers/2004/n1653.htm) | 部分               | 部分               | 部分               | 部分               |
| [long long](http://www.open-std.org/jtc1/sc22/wg21/docs/papers/2005/n1811.pdf) | 是                 | 是                 | 是                 | 是                 |
| [扩展的整型](http://www.open-std.org/jtc1/sc22/wg21/docs/papers/2006/n1988.pdf) | 不可用             | 不可用             | 不可用             | 不可用             |

[[本文内容](https://msdn.microsoft.com/zh-cn/library/hh567368.aspx#top)]

### C++ 14 核心语言功能

|                                 |                    |                    |
| :------------------------------ | :----------------- | :----------------- |
| 功能                            | Visual Studio 2013 | Visual Studio 2015 |
| 上下文转换的已调整 workding     | 是                 | 是                 |
| 二进制文本                      | 否                 | 是                 |
| auto 和 decltype(auto) 返回类型 | 否                 | 是                 |
| init-capture                    | 否                 | 是                 |
| 泛型 lambda                     | 否                 | 是                 |
| 变量模板                        | 否                 | 否                 |
| 扩展的 constexpr                | 否                 | 否                 |
| 聚合的 NSDMI                    | 否                 | 否                 |
| 避免/合成分配                   | 否                 | 否                 |
| [已弃用] 特性                   | 否                 | 否                 |
| 大小经过调整的分配              | 否                 | 是                 |
| 数字分隔符                      | 否                 | 是                 |

### C++17 建议的核心语言功能

|                                                  |                    |                    |
| :----------------------------------------------- | :----------------- | :----------------- |
| 功能                                             | Visual Studio 2013 | Visual Studio 2015 |
| 针对自动使用大括号内的初始值设定项列表的新建规则 | 否                 | 否                 |
| 简要静态断言                                     | 否                 | 否                 |
| 模板-参数模板的类型名称                          | 否                 | 否                 |
| 删除三字符组                                     | 是                 | 是                 |
| 嵌套的命名空间定义                               | 否                 | 否                 |
| N4259 std::uncaught_exceptions()                 | 否                 | 否                 |
| N4261 修复限定转换                               | 否                 | 否                 |
| N4266 命名空间和枚举器的特性                     | 否                 | 否                 |
| N4267 u8 字符文本                                | 否                 | 否                 |
| N4268 允许更多非类型模板参数                     | 否                 | 否                 |
| N4295 Fold 折叠表达式                            | 否                 | 否                 |
| 等待/继续                                        | 否                 | 是                 |



# [C99标准的新特性](https://www.cnblogs.com/wuyudong/p/c99-new-feature.html)



## C语言标准的发展

C语言的发展历史大致上分为4个阶段：Old Style C、C89、C99和C11.

C89是最早的C语言规范，于1989年提出，1990年先由ANSI(美国国家标准委员会，American National Standards Institute)推出ANSI版本，后来被接纳为ISO国际标准(ISO/IEC9899:1990)，因而有时也称为C90，最经典的C语言教材[K&R]就是基于这个版本的，C89是目前最广泛采用的C语言标准,大多数编译器都完全支持C89，C99(ISO/IEC9899:1999）是在1999年推出的，加入了许多新的特性，但目前仍没有得到广泛支持，在C99推出之后相当长的一段时间里，连gcc也没有完全实现C99的所有特性。2011年12月8号，ISO 发布了新的 C 语言的新标准——C11，之前被称为C1X，官方名称 ISO/IEC 9899:2011。

本文地址：<http://www.cnblogs.com/archimedes/p/c99-new-feature.html>，转载请注明源地址。

现在介绍一下C99相对于C89或者ANSI C的新特性：

## 1.复数（complex）

`complex.h`是C标准函数库中的头文件，提供了复数算术所需要的宏定义与函数声明。

```
#define complex  _Complex
#define _Complex_I  ((const float _Complex)__I__)
#define I  _Complex_I
```

C99规定了关键字`_Complex`。因而有3种复数类型：

- `double _Complex`
- `float _Complex`
- `long double _Complex`

次序不是必须遵守的，比如float _Complex也可以写成_Complex float。`_Complex_I`扩展为类型为`const float _Complex`的常量值，其值为虚数单位。C99规定`complex`作为宏扩展为`_Complex`。但C++未定义`complex`宏。gcc仅支持complex type，不支持imaginary type。因此宏`I`扩展为`_Complex_I`。

<complex.h>里面还包含了不少支持复数的数学函数（c打头的就是）：

1、ccos,csin,ctan,cacos,casin,catan：复数域内的三角函数，有对应的f和l版本。

2、ccosh,csinh,ctanh,cacosh,casinh,catanh：复数域内的双曲函数，有对应的f和l版本。

3、cexp,clog,cabs,cpow,csqrt：复数域内的指数、对数、绝对值、幂函数，有对应的f和l版本。

4、carg,cimag,creal,conj,cproj：获取象限角、虚数部分、实数部分、a=x及b=-y、Riemann球上的投影，有对应的f和l版本。

代码：

```
#include<stdio.h>
#include<complex.h>
int main()  
{
    double complex cmp = 1.3 + 2.3*I;
    printf("%f + %fi\n", creal(cmp), cimag(cmp));
    return 0;  
}  
```

## 2.指定初始化（Designated Initializers）

在初始化结构体和数组时，可以通过指定具体成员名或数组下标来赋初值

要指定数组的索引对应的值，可以在相应的元素值前使用‘[index] =’，index必须是常量表达式例如：

```
int a[6] = { [4] = 29, [2] = 15 };
```

等价于：

```
 int a[6] = { 0, 0, 15, 0, 29, 0 };
```

还可以向下面这样初始化：

```
int a[10] = { [1] = 1, [8 ... 9] = 10 };
```

这样可以只初始化a[1], a[8], a[9]三个元素，其他元素的值为0，等价于：

```
int a[10] = {0, 1, 0, 0, 0, 0, 0, 0, 10, 10};
```

对于结构体，指定成员名初始化可以使用‘.fieldname=’，例如：

```
 struct point { int x, y; };
```

接下来初始化：

```
struct point p = { .y = yvalue, .x = xvalue };  // 等价于 struct point p = { xvalue, yvalue };
```

还可以使用冒号：

```
struct point p = { y: yvalue, x: xvalue };
```

当然也可以用在union中：

```
union foo { int i; double d; };
union foo f = { .d = 4 };
```

## 3.变长数组（Variable Length Arrays）

C99允许可以定义一个长度为变量的数组（这个数组的长度可以到运行时才决定）



```
FILE *
concat_fopen (char *s1, char *s2, char *mode)
{
       char str[strlen (s1) + strlen (s2) + 1];
       strcpy (str, s1);
       strcat (str, s2);
       return fopen (str, mode);
}
```

也可以在结构体或是联合中使用VLA：

```
void foo (int n)
{
      struct S { int x[n]; };
 }
```

你可以使用alloca函数实现类似的功能，但是alloca函数并不是都实现，从另一角度而言，VLA更加的优秀

也可以使用VLA作函数参数：

```
struct entry
tester (int len, char data[len][len])
{
       /* ... */
}
```

当然也可以后传len

```
struct entry
tester (int len; char data[len][len], int len)  //注意分号
{
       /* ... */
}
```

示例代码：

```
#include<stdio.h>
void func(int n)
{
    int vla[n];
    printf("%d\n", sizeof(vla));
}
int main()  
{
    func(4);
    return 0;  
}  
```

## 4.单行注释 

gcc支持像C++风格的注释，以‘//’开头直到一行的结束，很多其他支持C99的C编译器都支持，但是c99之前的版本有可能不支持

## 5.柔性数组成员（Flexible Array Members）

 参见《**C语言柔性数组**》一文

## 6.long long类型

C99支持64位整型，使用long long int 或使用unsigned long long int,将整型常量声明为long long int，在整数的后面加上‘LL’，若为unsigned long long int,则加上‘ULL’

## 7.inline函数

c/c++中的inline，使用在函数声明处，表示程序员请求编译器在此函数的被调用处将此函数实现插入，而不是像普通函数那样生成调用代码(申请是否有效取决于编译器)。一般地说，这样作的优点是省掉了调用函数的开销；缺点则是可能会增加代所生成目标代码的尺寸

实际上，即使没有手工指定inline函数，编译器一般还会选择一些代码量较小但使用频繁的函数作为inline函数，以此作为性能优化的途径之一。

和带参宏定义(Parameterized Macro)相比,具备以下优点：

- 参数类型检查：宏定义中所使用的参数仅仅是在宏定义中被替换，不进行任何的类型检查
- 返回值：宏定义中无法使用return返回
- 便于调试

示例代码：

```
static inline int
inc (int *a)
{
       return (*a)++;
}
```

## 8.bool类型

记得以前都是自己写#define TRUE 1, #define FALSE  0 或者 enum boolean之类的宏，现在可以使用<stdbool.h>的bool类型啦

## 9.复合常量（Compound Literals）

简单来说复合常量就是允许你定义一个匿名的结构体或数组变量。如：

```
struct foo {int a; char b[2];} structure;
structure = ((struct foo) {x + y, 'a', 0});
```

等价于：

```
{
     struct foo temp = {x + y, 'a', 0};
     structure = temp;
}
```

也可以创建一个数组：

```
char **foo = (char *[]) { "x", "y", "z" };
```

更多实例：

```
static struct foo x = (struct foo) {1, 'a', 'b'};
static int y[] = (int []) {1, 2, 3};
static int z[] = (int [3]) {1};
//等价于下面的代码：
static struct foo x = {1, 'a', 'b'};
static int y[] = {1, 2, 3};
static int z[] = {1, 0, 0};
```

## 10.for循环变量初始化（for loop intializers）

C99引入了C++中的for循环变量初始化方式：

```
for(int i = 0; i < 10; ++i) {
  ...;
}
```

除了写起来方便以外，循环变量的生存周期也被最小化了。这也顺便杜绝了那种把循环变量带到循环外头继续用的恶习

## 11.C99新增头文件

C89中标准的头文件：
  <assert.h>             定义宏assert()
  <ctype.h>              字符处理
  <errno.h>              错误报告
  <float.h>               定义与实现相关的浮点值勤
  <limits.h>              定义与实现相关的各种极限值
  <locale.h>              支持函数setlocale()
  <math.h>              数学函数库使用的各种定义
  <setjmp.h>           支持非局部跳转
  <signal.h>             定义信号值
  <stdarg.h>            支持可变长度的参数列表
  <stddef.h>            定义常用常数
  <stdio.h>              支持文件输入和输出
  <stdlib.h>             其他各种声明
  <string.h>             字符串函数
  <time.h>               支持系统时间函数

C99新增的头文件
  <complex.h>          支持复杂算法
  <fenv.h>               给出对浮点状态标记和浮点环境的其他方面的访问
  <inttypes.h>          定义标准的、可移植的整型类型集合,也支持处理最大宽度整数的函数
  <iso646.h>            首先在此1995年第一次修订时引进,用于定义对应各种运算符的宏
  <stdbool.h>           支持布尔数据类型类型,定义宏bool,以便兼容于C++
  <stdint.h>             定义标准的、可移植的整型类型集合,该文件包含在<inttypes.h>中
  <tgmath.h>           定义一般类型的浮点宏
  <wchar.h>             首先在1995年第一次修订时引进,用于支持多字节和宽字节函数
  <wctype.h>           首先在1995年第一次修订时引进,用于支持多字节和宽字节分类函数

注意：还有一些新特性未总结进来，待充分理解实践之后将陆续补充

### 参考资料

<https://gcc.gnu.org/onlinedocs/gcc/C-Extensions.html#C-Extensions>