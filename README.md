# Ki project
Making a new language and its ecosystem

## ki language specification
#### Sum type
```
<...>
```
The sum type has **any** of its elements present at any given time. Classical computers can effectively work with only one element present at any given time.

#### Product type
```
(...)
```
The product type has **all** of its elements present at any given time. The sum and product types provides the only way of composing elements. 

---
#### Equality binding
```
a = b
```
The equality binding means that one element can be replaced with another with no consequences.

#### Type binding
```
a ~ b
```
The type binding means that one element can be replaced with another, but may have different behaviour.

#### Reference binding
```
a: b
```
The reference binding means that one element points to another.

---
#### nil 
The **nil** means nothing. It is the element of nothingness. It represents the zero from the set of natural number.

#### one
The **one** means something. It is the element of existence. It represents the one from the set of natural numbers. For a and b from above:
```
a ~ one
b ~ one
```

#### any 
The **any** means anything. It is the element of uncertainty. It represents the unknown element from the set of natural numbers.

#### all
The **all** means everything. It is the element of infinity. It represents the set of natural numbers

---
#### Playing with elements
```
0 = nil
1 = one
0, 1, 2, 3 = 0...4
all = nil...any
bool = <true, false>
byte = <0...256>
```

#### More about reference binding
```
b: bool
b = <bool>
b! ~ bool

s: (a, b)
s = <(a, b)>
s! ~ (a, b)

t: x, y
t = <x, y>
t ~ x, y

0, 1: a
0 = <a>
1 = <a>
```

---
#### Choice function
```
struct = (
    i: byte
    b: bool
)

struct i ~ byte
struct b ~ bool
```
```
array = (
    0...3: byte
)
array 0 ~ byte
array 1 ~ byte
array 2 ~ byte
```

#### Branching and conditions
```
a: any
a! b = a [any] b
```
```
b: bool
b = <bool>
b [bool] ~ bool
b! ~ bool
```
```
b: bool
b! {
    [true]  (...)
    [false] (...)
}
```

#### Functions
```
f = [int] -> bool
f = [int] -> . 0 ge

10 f = true
10 [int] -> . 0 ge = true
```
Functions are chains of linked elements. Elements at the start and the end make function signature. The f function here has two equality bindings. Compiler will use the one that can be decompose into base elements, which can be interpreted by the compiler.
```
f = [int] -> bool
f_var: f
f_var = <f>
f_var [int] -> . 0 ge
```
Here we create a function variable f_var. The ge function call here is the great or equal function. The type of function variable tells the start and the end of the fucntion chain. 
```
[int] -> . 0 ge ~ f
```
Any chain type-equal to another if they have the same boundaries.

---
