readFile(function (fileData) {
  console.log('read file: ' + fileData)
})

invokeHttp(function (resp) {
  console.log(resp.data)
})

