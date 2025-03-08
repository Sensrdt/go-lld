function test (record) {
    
    if ({} == {}) {
        console.log("test")
    }
    console.log(+true)
    console.log(!'true')
    
    if (record == {age : 28}) {
        console.log("match 1")
    } else if (record === {age : 28}) {
        console.log("match 2")
    } else {
        console.log("no match")
    }
}

test({age : 28})

let a = {};
let b = {age: 20};
let c = {age: 30, name: "Sridip"};

let {age, name} = c;
console.log(name)

// a[Object.keys(c)] = {age, name};
b[c] = 50;
b[b] = 500

console.log(a, b, c)

// console.log(a[c])
// console.log(a[b])
// console.log(b[b])
// console.log(b[c])
// console.log(b.age)
// console.log(c[c])

