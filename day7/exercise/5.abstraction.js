class Car {
    #name = ""
    #model = ""
    #duit = 0
    
    constructor(name, model, duit){
        this.#name = name
        this.#model = model
        this.#duit = duit
    }

    // getter
    get name() {
        return this.#name
    }

    get model() {
        return this.#model
    }

    get duit() {
        // calculateDuit()
        // convertDuit()
        return this.#duit
    }
}

let myCar = new Car("Toyota", "X", 10000)
// myCar.#duit = 9999

console.log(myCar.duit)

