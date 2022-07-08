# Ki project
Making a new language and its ecosystem

## ki language specification
#### nil 
The **nil** means nothing. It is the element of nothingness. It represents the zero from the set of natural number.

#### one
The **one** means something. It is the element of existence. It represents one from the set of natural numbers.

#### any 
The **any** means anything. It is the element of uncertainty. It represents the unknown element from the set of natural numbers.

#### all
The **all** means everything. It is the element of infinity. It represents the set of natural numbers

---
#### Equality binding
```
a = b

==
    a
    b
    c
```
The equality binding means that one element can be replaced with another with no consequences.

#### Type binding
```
a ~ b

~~
    a
    b
    c
```
The type binding means that one element can be replaced with another, but may have different behaviour.

---
#### State type
```
<...>
```
The state type has **any** of its elements present at any given time. Classical computers can effectively work with only one element present at any given time.

#### Switch type
```
[...]
```
The switch type has **all** of its elements present at any given time. The state and switch types provides the only way of composing elements. 

#### Series type
```
|...|
|a, b, c| = a -> b -> c
```
The series type represents series of actions/relations. It is the only way to create interfaces and functions. 

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

#### Making a product type
```
==
    <0, 1, 2> -> . [
        0 -> a
        1 -> b
        2 -> c
    ] -> <a, b, c>
    
    {
        0: a
        1: b
        2: c
    }
```
Struct example:
```
{
    key0: value0
    key1: value1
    ...
    keyn: valuen
}
```

Array example:
```
{
    0...3: type
}
```

---
#### Branching and conditions
```
bool [
    true  ->
    false ->
]
```

#### Choice function
```
struct = (
    i: #byte
    b: bool
)

i >> struct ~ byte
b >> struct = bool
```
```
array = (
    0...3: byte
)
0 >> array ~ byte
1 >> array ~ byte
2 >> array ~ byte
```

#### Functions
```
f = int -> bool
f = int -> . 0 >> ge

10 >> f = true
10 >> |int -> . 0 >> ge| = true
10 -> . 0 >> ge = true
10 0 >> ge = true
true = true
```
Functions are series of linked elements. Elements at the start and the end make function signature. The f function here has two equality bindings. Compiler will use the one that can be decompose into base elements, which can be interpreted by the compiler.
```
f = int -> bool
f_var = <f>
f_var int -> . 0 >> ge
```
Here we create a function variable f_var. The ge function call here is the great or equal function. The type of function variable tells the start and the end of the fucntion chain. 
```
|int -> . 0 >> ge| ~ f
```
Any chain type-equal to another if they have the same boundaries.

---
