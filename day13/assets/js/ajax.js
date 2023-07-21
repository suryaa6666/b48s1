const promise = new Promise((resolve, reject) => {
    const xhr = new XMLHttpRequest()

    xhr.open("GET", "https://api.npoint.io/4f0ebf1a67ed565e898a", true)
    xhr.onload = function () {
        // http code : 200 -> OK
        if (xhr.status === 200) {
            resolve(JSON.parse(xhr.responseText))
        } else if (xhr.status >= 400) {
            reject("Error loading data")
        }
    }
    xhr.onerror = function () {
        reject("Network error")
    }
    xhr.send()
})

// promise-chaining
// promise.then((value) => {
//     console.log(value)
// }).catch((reason) => {
//     console.log(reason)
// })

// async-await

let testimonialData = []

async function getData(rating) {
    try {
        const response = await promise
        console.log(response)
        testimonialData = response
        allTestimonial()
    } catch (err) {
        console.log(er)
    }
}

getData()

function allTestimonial() {
    let testimonialHTML = ""

    testimonialData.forEach((card, index) => {
        testimonialHTML += `<div class="testimonial">
    <img src="${card.image}" class="profile-testimonial" />
    <p class="quote">"${card.quote}"</p>
    <p class="author">- ${card.user}</p>
    <p class="author">${card.rating} <i class="fa-solid fa-star"></i></p>
</div>
`
    })

    document.getElementById("testimonials").innerHTML = testimonialHTML
}

function filterTestimonial(rating) {
    let filteredTestimonialHTML = ""

    const filteredData = testimonialData.filter((card) => {
        return card.rating === rating
    }) 

    filteredData.forEach((card) => {
        filteredTestimonialHTML += `<div class="testimonial">
        <img src="${card.image}" class="profile-testimonial" />
        <p class="quote">"${card.quote}"</p>
        <p class="author">- ${card.user}</p>
        <p class="author">${card.rating} <i class="fa-solid fa-star"></i></p>
    </div>
    `
    })

    document.getElementById("testimonials").innerHTML = filteredTestimonialHTML
}