// console.log("Zidane Ganteng")
// alert("Surya Ganteng")
// document.write("Hello Guys! 🔥")

// VAR -> VERSI DULU, BISA DEKLARASI ULANG -> FUNCTIONAL SCOPE
// LET -> TERBARU, GAK BISA DEKLARASI -> BLOCK SCOPE
// CONST (CONSTANT) -> VALUE GAK BISA DIUBAH -> BLOCK SCOPE

// function getNama() {
//     if(true) {
//         let x = 10;
//     }
//     console.log(x);
// }

// getNama()

// const x = 10;

// const x = 100

// console.log(x)

// LET
// let nama = "Zidane"
// let umur = 60

// alert(`Halo nama saya ${nama} dan umur saya ${umur}`)

/*
asdasd
asd
asd
sad
asdasd
*/

// tipe data : string (""), integer (123), object ({}), array ([])

// operator
// let bilanganPertama = 50
// let bilanganKedua = 100

// alert("hasilnya : " + bilanganPertama + bilanganKedua)

// condition (if-else, else if), switch-case
// perintah : bikin pengecekan nilai, kalau misal minimal 75 maka dia lulus, kalau dibawah itu maka dia nggak lulus, pake alert

// grade -> minimal 80 -> B, minimal 90 -> A, dibawah 80 dia C
// let grade = 79
// if (grade >= 90) {
//     alert('A')
// }else if (grade >= 80){
//     alert('B')
// }else{
//     alert('C')
// }


// function (named function, function with parameter)
// named function
// function hello() {
//     alert("Hello World 😁")
// }

// function penjumlahan() {
//     let bilanganPertama = 10
//     let bilanganKedua = 27
//     alert(`hasil penjumlahan : ${bilanganPertama + bilanganKedua}`)
// }

// hello()
// penjumlahan()

// function with parameter
// function perkalian(angka1, angka2) {
//     alert(`hasil perkalian : ${angka1 * angka2}`)
// }

// perkalian(2, 10) // 20 

// looping (for loop, while loop, do-while)

// for (let x = 0; x < 10; x++) {
//     console.log(x) // 0, 1, 2, 3, ... 9
// }

function submitData(surya) {
    surya.preventDefault()
    // DOM -> Document Object Model (untuk memanipulasi element HTML melalui javascript)

    let name = document.getElementById("input-name").value // Cintara Surya Elidanto
    let email = document.getElementById("input-email").value
    let phone = document.getElementById("input-phone").value
    let subject = document.getElementById("input-subject").value
    let message = document.getElementById("input-message").value

    let objectData = {
        name: name,
        email,
        phone,
        subject,
        message
    }

    let arrayData = [name, email, phone, subject, message]

    console.log(objectData)

    if (name === "") {
        return alert('Name harus diisi!')
    } else if (email === "") {
        return alert('Email harus diisi!')
    } else if (phone === "") {
        return alert('Phone harus diisi!')
    } else if (subject === "") {
        return alert('Subject harus diisi!')
    } else if (message === "") {
        return alert('Message harus diisi!')
    }

    const emailReceiver = "surya@gmail.com"

    let a = document.createElement('a')
    a.href = `mailto:${emailReceiver}?subject=${subject}&body=Halo nama saya ${name},\n${message}, silahkan kontak saya di nomor berikut : ${phone}`
    a.click()

    // alert(`name : ${name}\nemail : ${email}\nphone : ${phone}\nsubject : ${subject}\nmessage: ${message}`)
}

// PR : DARK MODE