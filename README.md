This repository was created with Thorsten Ball's book "Writing an interpreter with GO". It's an interpreter of the Monkey programming language with the following syntax:

### Variable definition

```
let i = 2 + 2;
let s = "str";
let s2 = "str" + "ing";
let b = false;
```

### Functions

```
let add = fn(a, b) { return a + b; };
let add = fn(a, b) { a + b; };

add(1, 2);
```

### Conditionals

```
if (x == 0) {
    0
}
else {
    if (x == 1) {
        1
    }
}
```

### Arrays and hashmaps

```
let arr = [1, 2, 3, 4, 5];
let hash = {"x": 1, "y": "str"};

let x = arr[0];
let y = hash["x"];
```

### Macros

```
let unless = macro(condition, consequence, alternative) {
     quote(if (!(unquote(condition))) { unquote(consequence);
}

unless(10 > 5, puts("not greater"), puts("greater"));
```
