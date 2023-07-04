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
        postAt: "04 Juli 2023",
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
                ${dataBlog[index].postAt} | ${dataBlog[index].author}
            </div>
            <p>
                ${dataBlog[index].content}
            </p>
            </div>
        </div>
        `
    }
}