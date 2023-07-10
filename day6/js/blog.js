// bikin 3 variable untuk nama  3 orang, terus lakukan console log untuk ketiganya

// Variabel
// let name1 = "Andi"
// let name2 = "Budi"
// let name3 = "Caca"

// Output console.log
// console.log(name1)
// console.log(name2)
// console.log(name3)

// kita akan tampung data ketiga variable nama tadi dalam sebuah array dan lakukan print ke console dari array yang dibuat

// ARRAY

// let arrayName =[
// 'andi', 
// 'budi',
// 'caca' 
// ]

// console.log(arrayName) // simpen aja

// console.log(arrayName[0])
// console.log(arrayName[1])
// console.log(arrayName[2])

// // OBJECT
// let orang = {
//     name: "surya",
//     address: "bintaro"
// }

// console.log(orang.name);
// console.log(orang.address);

// let orang2 = {
//     name : "Rahmat",
//     address : "Depok"
// }

// console.log(orang2.name)
// console.log(orang2.address)

// const data = {
//     name: ["angga", "nur", "ardiansyah"],
//     address: "kalimantan"
// }

// halo namaku angga nur ardiansyah, asalku dari kalimantan

// console.log("Hallo nama ku " + data.name[0] + " " + data.name[1] + " " + data.name[2] + "," + " asalku dari " + data.address);

// ARRAY OF OBJECT
// const users = [
//     {
//         name: "surya",
//         address: "bintaro"
//     },
//     {
//         name: "angga",
//         address: "bandung"
//     },
//     {
//         name: "wahyu",
//         address: "bengkulu"
//     }
// ]

// console.log(users)

// // ambil data mas angga
// console.log(users[1].name)
// console.log(users[1].address)


// PUSHING OBJECT TO ARRAY
// let products = []

// function addData(event) {
//     event.preventDefault()

//     const product = {
//         name: "Popok Bayi",
//         price: 27000,
//         qty: 2,
//     }

//     products.push(product)
//     console.log(products)
// }

// let blogs = []

// function addData(event) {

//     event.preventDefault()

//     let title = document.getElementById("input-blog-title").value
//     let content = document.getElementById("input-blog-content").value

//     let blog = {
//         title,
//         content
//     }

//     blogs.push(blog)
//     console.log(blogs)
// }

let dataBlog = []

const addBlog = (event) => {
    event.preventDefault()

    let title = document.getElementById("input-blog-title").value
    let content = document.getElementById("input-blog-content").value
    let image = document.getElementById("input-blog-image").files

    // untuk membuat object file menjadi URL secara sementara, agar tampil
    image = URL.createObjectURL(image[0])

    // console.log(title)
    // console.log(content)
    // console.log(image)

    let blog = {
        title,
        content,
        image, // bentuknya blob url (sementara)
        postAt: new Date(),
        author: "Surya Gans"
    }

    dataBlog.push(blog)
    renderBlog()

    console.log(dataBlog)
}

function renderBlog() {
    document.getElementById("contents").innerHTML = ''
    // misal dataBlog = 3 object
    for (let index = 0; index < dataBlog.length; index++) {
        document.getElementById("contents").innerHTML += `
        <div class="blog-list-item">
            <div class="blog-image">
            <img src="${dataBlog[index].image}" alt="" />
            </div>
            <div class="blog-content">
            <div class="btn-group">
                <button class="btn-edit">Edit Post</button>
                <button class="btn-post">Delete Post</button>
            </div>
            <h1>
                <a href="blog-detail.html" target="_blank">${dataBlog[index].title}</a>
            </h1>
            <div class="detail-blog-content">
                ${convertDate(dataBlog[index].postAt)} | ${dataBlog[index].author}
            </div>
            <p>
                ${dataBlog[index].content}
            </p>
            <p>
                ${getDistanceTime(dataBlog[index].postAt)}
            </p>
            </div>
        </div>
        `
    }
}


// const date = new Date()
// console.log("cek tanggal : ", date.getDate())
// console.log("cek jam : ", date.getHours())
// console.log("cek jam : ", date.getUTCHours())
// console.log("cek Hari : ", date.getDay())
// console.log("cek Year : ", date.getFullYear())
// console.log("cek Bulan : ", date.getMonth())
// console.log("cek Second : ", date.getSeconds())

// challenge : 0 -> January, 1 -> February, dst...
// if else, switch case

function convertDate(surya) { // surya : new Date()
    let date = new Date(surya)

    // TANGGAL 
    const tanggal = date.getDate()

    // BULAN
    const listBulan = ["Jan", "Feb", "Mar", "Apr", "May", "Jun", "Jul", "Aug", "Sep", "Oct", "Nov", "Dec"]

    // bulan agustus
    // console.log(listBulan[7])

    // bulan sesuai dengan bulan ini
    // console.log(listBulan[date.getMonth()])
    const bulan = listBulan[date.getMonth()]

    const year = date.getFullYear()

    let hours = date.getHours()

    let minutes = date.getMinutes()

    if (hours < 10) {
        hours = "0" + hours // 0-9 -> 00, 01, 02, .. 09 -> 10, 11, 12
    }

    if (minutes < 10) {
        minutes = "0" + minutes // 0-9 00, 01, 02, .. 09 
    }

    // 5 Jul 2023 09:34 WIB
    return `${tanggal} ${bulan} ${year} ${hours}:${minutes} WIB`
}

function getDistanceTime(surya) { // surya : new Date()
    console.log("hai")
    let timeNow = new Date()
    let timePost = surya

    let timeDistance = timeNow - timePost

    // const miliseconds = 1000
    // const seconds = miliseconds * 60
    // const minutes = seconds * 60
    // const hours = minutes * 60
    // const day = hours * 24
    // const month = day * 30
    // const year = month * 12

    let distanceSeconds = Math.floor(timeDistance / 1000)
    let distanceMinutes = Math.floor(distanceSeconds / 60)
    let distanceHours = Math.floor(distanceMinutes / 60)
    let distanceDay = Math.floor(distanceHours / 24)
    let distanceMonth = Math.floor(distanceDay / 30)
    let distanceYear = Math.floor(distanceMonth / 12)

    // floor -> 1.8 -> 1
    // ceil -> 1.4 -> 2
    // round -> 1.3 -> 1

    // console.log(distanceSeconds)
    // console.log(distanceMinutes)

    if (distanceSeconds >= 60) {
        return `${distanceMinutes} menit yang lalu`
    } else if (distanceMinutes >= 60) {
        return `${distanceHours} jam yang lalu`
    } else if (distanceHours >= 24) {
        return `${distanceDay} hari yang lalu`
    } else if (distanceDay >= 30) {
        return `${distanceMonth} bulan yang lalu`
    } else if (distanceMonth >= 12) {
        return `${distanceYear} tahun yang lalu`
    }

    return distanceSeconds + " detik yang lalu"
}

setInterval(() => {
    renderBlog()
}, 1000) // 1000ms = 0.001 seconds
