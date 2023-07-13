let condition = true;

let janji = new Promise((resolve, reject) => {
    if (condition) {
        setTimeout(() => {
            resolve("Janji ditepati!")
        }, 3000)
    } else {
        reject("Janji gugur!")
    }
})


async function getData() {
    try {
        const response = await janji;
        console.log(response)
        // Swal.fire(
        //     'Good job!',
        //     response,
        //     'success'
        //   )
    } catch (err) {
        // alert -> registrasi user gagal
        // Swal.fire({
        //     icon: 'error',
        //     title: 'Oops...',
        //     text: 'Registrasi gagal!',
        // })
        console.log(err)
    }
}

// console.log(promise)
// janji.then((value) => {
//     console.log(value)
// }).catch((err) => {
//     console.log(err)
// }).finally(() => {
//     console.log("selesai")
// })