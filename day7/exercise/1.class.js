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



