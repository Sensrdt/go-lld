
// caching
const memoize = (fn) => {
    const res = {}
    return function (...args) {
        let resArgs = JSON.stringify(args)
        if (!res[resArgs]) {
            res[resArgs] = fn.call(this, ...args)
        }
        return res[resArgs]
    }
}


const MultiplyValues = (a, b) => {
    console.log(a, b)
    for (let i=0; i<100000;i++) {
    }
    return a * b;
}

const cacheFn = memoize(MultiplyValues);
console.log(cacheFn(10, 10))
// ----- 

function add(a) {
    return function (b) {
        if (b) return add(a + b);
        else 
        return a;
    }
}

add(4)(5)(6)(7)()
//-----

const calc = {
    total: 0,
    add(arg) {
        this.total += arg
        return this
    },
    multiply(arg) {
        this.total *= arg
        return this
    }
}

const result = calc.add(5).multiply(10)
console.log(
    result
)

// -- debounce

function debounce(cb, time) {
    let timer;

    return function (...args) {
        if (timer) clearTimeout(timer)
        timer = setTimeout(() => {
            cb(...args)
        }, time)
    }
}


// ---- polyfill promise.all

const myPromise  = (promies) => {

    let trackResolvedPromises = [];
    let countResolvedPromises = 0

    return new Promise((resolve, reject) => {

        promies.forEach((element, idx) => {
            Promise.resolve(element).then((resolvedValue) => {
                trackResolvedPromises[idx] = resolvedValue;
                countResolvedPromises++;

                if (countResolvedPromises == promies.length) {
                    let result = trackResolvedPromises;
                    resolve(trackResolvedPromises);
                }
            }).catch((err) => {
                reject(err);
            })
        });
    })
}