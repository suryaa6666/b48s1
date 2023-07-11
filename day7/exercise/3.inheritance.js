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
        if (this.color == "") {
            return `I have a car, i buy it in ${this.price}`
        }
        return `I have a car with color ${this.color}, i buy it in ${this.price}`
    }
}

class ElectricCar extends Car {
    constructor(color, price, batteryCapacity) {
        super(color, price)
        this.batteryCapacity = batteryCapacity 
    }

    // override / menimpa -> polymorphism
    getInfo() {
        return `I have an electric car with color ${this.color}, i buy it in ${this.price}, with batery capacity ${this.batteryCapacity}`
    }
}

const myElectricCar = new ElectricCar("red", 9000, 200)
console.log(myElectricCar.getInfo())
