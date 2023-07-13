const xhr = new XMLHttpRequest()

// CRUD
xhr.open("GET", "https://your-url", true) // asynchronous
// param 1 : is the method
// param 2 : place of data by url
// param 3 : true -> asynchronous, false -> synchronous

xhr.onload = function () { } // mengecek status
xhr.onerror = function () { } // menampilkan error ketika request
xhr.send()

