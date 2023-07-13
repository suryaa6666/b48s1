// class✅
// object ✅
// inheritance ✅
// polymorphism ✅
// abstraction ✅
// encapsulation ✅

// class Testimonial {
//     #quote = ""
//     #image = ""

//     constructor(quote, image) {
//         this.#quote = quote
//         this.#image = image
//     }

//     get quote() {
//         return this.#quote
//     }

//     get image() {
//         return this.#image
//     }

//     get user() {
//         throw new Error('there is must be user to make testimonials')
//     }

//     get testimonialHTML() {
//         return `<div class="testimonial">
//             <img src="${this.image}" class="profile-testimonial" />
//             <p class="quote">"${this.quote}"</p>
//             <p class="author">- ${this.user}</p>
//         </div>
//         `
//     }
// }

// class UserTestimonial extends Testimonial {
//     #user = ""

//     constructor(user, quote, image) {
//         super(quote, image)
//         this.#user = user
//     }

//     get user() {
//         return "user : " + this.#user
//     }
// }

// class CompanyTestimonial extends Testimonial {
//     #company = ""

//     constructor(company, quote, image) {
//         super(quote, image)
//         this.#company = company
//     }

//     get user() {
//         return "company : " + this.#company
//     }
// }

// const testimonial1 = new UserTestimonial("Surya Elidanto", "GG gaming", "https://images.unsplash.com/photo-1580477667995-2b94f01c9516?ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D&auto=format&fit=crop&w=870&q=80")

// const testimonial2 = new UserTestimonial("Guswandi", "Keren kamu bang", "https://images.unsplash.com/photo-1541562232579-512a21360020?ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D&auto=format&fit=crop&w=387&q=80")

// const testimonial3 = new CompanyTestimonial("Dumbways", "Apasih ga jelas", "https://images.unsplash.com/photo-1578632767115-351597cf2477?ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D&auto=format&fit=crop&w=387&q=80")

// // const testimonial4 = new Testimonial("Apasih ga jelas", "https://images.unsplash.com/photo-1578632767115-351597cf2477?ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D&auto=format&fit=crop&w=387&q=80")

// // console.log(testimonial1, testimonial2, testimonial3)

// let testimonialData = [testimonial1, testimonial2, testimonial3]

// let testimonialHTML = ""

// for (let i = 0; i < testimonialData.length; i++) {
//     testimonialHTML += testimonialData[i].testimonialHTML
// }

// document.getElementById("testimonials").innerHTML = testimonialHTML

// fake data -> array of objects -> data boong2an
const testimonialData = [
    {
        user: "Surya Elidanto",
        quote: "GG gaming",
        image: "https://images.unsplash.com/photo-1580477667995-2b94f01c9516?ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D&auto=format&fit=crop&w=870&q=80",
        rating: 5
    },
    {
        user: "Chandra",
        quote: "OKLAH",
        image: "https://images.unsplash.com/photo-1580477667995-2b94f01c9516?ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D&auto=format&fit=crop&w=870&q=80",
        rating: 4
    },
    {
        user: "Guswandi",
        quote: "Keren kamu bang",
        image: "https://images.unsplash.com/photo-1541562232579-512a21360020?ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D&auto=format&fit=crop&w=387&q=80",
        rating: 4
    },
    {
        user: "Dumbways",
        quote: "Apasih ga jelas",
        image: "https://images.unsplash.com/photo-1578632767115-351597cf2477?ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D&auto=format&fit=crop&w=387&q=80",
        rating: 1
    }
]

// let testimonialHTML = ``

// const testimonialsMap = testimonialData.map((card) => {
//     return `<div class="testimonial">
//     <img src="${card.image}" class="profile-testimonial" />
//     <p class="quote">"${card.quote}"</p>
//     <p class="author">- ${card.user}</p>
// </div>
// `
// })

// console.log(testimonialsMap)
// console.log(testimonialsMap.join())

// document.getElementById("testimonials").innerHTML = testimonialsMap.join()
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

// eksekusi awal / default
allTestimonial()

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


