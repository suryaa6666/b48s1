
// FUNCTIONAL PROGRAMMING (FP)
let x = car1("red", 20000, "SUPRA")
let y = car2("blue", 6000000, "ASTREA")

function car1(color, price, surya) {
    return getInfo(color, price, surya)
}

function car2(motor, mobil, surya) {
    return getInfo(motor, mobil, surya)
}

function getInfo(color, price, motor) {
    return `I have a car with color ${color}, i buy it in ${price} jenis ${motor}`
}

console.log(x)
console.log(y)
