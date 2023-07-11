class Vehicle {
    drive() {
        return "The vehicle is driving"
    }
}

class Car extends Vehicle {
    drive() {
        return "The car is driving"
    }
}

class ElectricCar extends Vehicle {
    drive() {
        return "The electric car is driving"
    }
}

const myVehicle = new Vehicle()
const myCar = new Car()
const myElectricCar = new ElectricCar()

console.log(myVehicle.drive())
console.log(myCar.drive())
console.log(myElectricCar.drive())