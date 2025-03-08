const age = 10;
var person = {
    name: "Piyush",
  age: 20,
  getAge: function(){
    return this.age;
  }
}

var person2 = {age:  24};
person.getAge.call(person2); // show with apply and bind as well


// append arr to elements 
arr.push.apply(arr, elements); //[a,b,0,1,2]


const numbers = [5, 6, 2, 3, 7];

// using Math.min/Math.max apply

let max = Math.max.apply(null, numbers); // equal to Math.max

let min = Math.min.apply(null, numbers); // equal to Math.min 

// vs. simple loop based algorithm

max = -Infinity, min = +Infinity;

for (let i = 0; i < numbers.length; i++) {
  if (numbers[i] > max) {
    max = numbers[i];
  }
  if (numbers[i] < min) {
    min = numbers[i];
  }
}


let user = {
  name: "Piyush",
  age: 24,
    getDetails() { // function
        const nestedArrow = () => console.log(this.name); //Piyush
        nestedArrow();
    }
};


let user1 = {
  name: "Piyush",
  age: 24,
    getDetails: () => { // arrow function
        console.log(this.name); // window object
    }
};


const fs = require('fs');
setTimeout(() => {
    (() => {
        console.log("test")
    })()
}, 0)