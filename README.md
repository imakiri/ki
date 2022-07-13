# Ki project
Making a new language and its ecosystem

## ki language specification
#### Equality binding
```
a is b
```
The equality binding means that one element can be replaced with another with no consequences. Used to declare implementations.

#### Type binding
```
a as b
```
The type binding means that one element can be replaced with another, but may have different behaviour. Used to help infer or to impose type

---
#### any context
```
any a b c
```
The any context tells that **any** of its elements present at any given time.

#### all context
```
all a b c
```
The all context tells that **all** of its elements present at any given time.

#### link context
```
link a -> b -> c
```
The link context represents the series of actions/relations. It is the only way to create interfaces and functions. 

---
#### Different notations
##### In-line notation
```
any a b c
all a b c
link a -> b -> c
```
##### Go notation
```
any {
    a
    b
    c
}
all {
    a
    b
    c
}
link {
    a
    b
    c
}
```

##### Ruby/Python notation
```
any
    a
    b
    c
all
    a
    b
    c
link
    a
    b
    c
```

---
#### Playing with elements
```
0 1 2 3 is 0...4
byte is any 0...256
bool is any true false
bool is any {
    true
    false
}
c is all {
    link a -> b -> c
    link e -> f -> g
}
c is all
    link 
        a
        b
        c
    link 
        e
        f
        g
```

#### Structures
```
person is all {
    link age -> int
    link name -> string
}
age >> person is int
name >> person is string
```

#### Functions
```
f as link int -> bool
f is link int -> . 0 >> ge

10 >> f is true
10 >> link int -> . 0 >> ge is true
```
Functions are series of linked elements. Elements at the start and the end make function signature. The f function here has two bindings: `as` binding infers type `link int -> bool` to f element, `is` binding implements function. These two statements will be checked by the compiler. It will throw an error if they are not compatible.

```
f as link int -> . 10 >> le
```
Any link type-equal to another if they have the same boundaries.

---
#### Type classes
```
T is any a b c
+ as link (T T) -> T
+ is all {
    link (a b) c
    link (a c) b
    link (a a) a
    link (c c) c
    link (b b) a
    link (b a) b
    link (b c) c
    link (c a) c
    link (c b) b
}

bool is any true false
+ as link (bool bool) -> bool
+ is all {
    link (true false) true
    link (false false) false
    link (false true) true
    link (true true) false
}

+ as link {
    G: any {
        (bool bool)
        (T T)
    }
    G >> all {
        link (bool bool) bool
        link (T T) T
    }
}

+ as all {
    link bool bool -> bool
    link T T -> T
    // link (T bool) _
    // link (bool T) _
}

+ is all {
    link (a b) c
    link (a c) b
    link (a a) a
    link (c c) c
    link (b b) a
    link (b a) b
    link (b c) c
    link (c a) c
    link (c b) b
    link (true false) true
    link (false false) false
    link (false true) true
    link (true true) false
}
```

