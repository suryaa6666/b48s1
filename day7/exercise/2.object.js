// OBJECT ORIENTED PROGRAMMING (OOP)

class Car {
    // properties : warna, jumlahRoda, jumlahKursi, harga
    // method : autoDrive, gas, rem, belok
    color = "";
    price = 0;
    constructor(color, price) {
        this.color = color
        this.price = price
    }

    getInfo() {
        return `I have a car with color ${this.color}, i buy it in ${this.price}`
    }
}

const mobil1 = new Car("red", 20000)
const mobil2 = new Car("blue", 6000000)
console.log(mobil1.getInfo())
console.log(mobil2.getInfo())

