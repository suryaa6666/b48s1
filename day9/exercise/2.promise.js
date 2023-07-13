// status promise : pending, fulfill (resolve), rejected
let condition = true;

let janji = new Promise((resolve, reject) => {
    if (condition) {
        // menunggu 5 detik
        resolve("Janji ditepati!")
        resolve("Janji ditepati lagi!")
        // console.log("surya gans")
    }

    // if (condition) {
    //     resolve("Janji ditepati!")
    //     console.log("surya gans")
    // }

    if (!condition) {
        reject("Janji gugur!")
        // console.log("surya jelek")
    }
})


console.log(janji)
janji.then((value) => {
    console.log(value)
}).catch((err) => {
    console.log(err)
}).finally(() => {
    console.log("selesai")
})