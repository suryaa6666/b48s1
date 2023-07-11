class Car {
    #name = ""
    #model = ""
    #duit = 0

    constructor(name, model, duit) {
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
        return this.#duit
    }

    // setter
    set duit(value) {
        if (value > 100) {
            return console.log("Isi duit gak boleh lebih dari 100")
        }
        this.#duit = this.#duit + value
    }
}

let myCar = new Car("Toyota", "X", 10000)
myCar.duit = 99

console.log(myCar.duit)

