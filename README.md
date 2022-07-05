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

#### ? binding
```
a::b
```
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

#### Playing with types
```
0 = nil
1 = one
0, 1, 2, 3 = 0...4
all = nil...any
bool = <true, false>
byte = <0...256>
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
a = <any>
a! b = a {
    [any] b
}
```
```
b: bool
b = <bool>
```
```
b: bool
b! {
    [true]  (...)
    [false] (...)
}
```
---
#### Work in progress: functions, assignment, loop generic
```
condition: any -> any, bool
iteration: condition, bool, any -> any
cycle: condition, any -> any

iteration (cond, b, args) = b {
    [false] args
    [true]  iteration (cond, cond args)
}
cycle (cond, args) = iteration cond true args
```
```
new: any -> <any, nil>

condition: *any, *bool -> nil
iteration: condition, *bool, *any -> nil
cycle: condition, *any -> nil

iteration (cond, b, args) = b {
    [false] nil
    [true]  cond (args, b) {
        [nil] iteration (cond, b, args)
    }
}
cycle (cond, args) = b: new bool {
    [bool] b true {
        [nil] iteration (cond, b, args)
    }
    [nil]  nil
}
```
