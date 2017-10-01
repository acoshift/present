readFile('savedUrl.txt')
  .then(function (fileData) {
    var url = 'http://' + fileData
    return invokeHttp(url)
  })
  .then(function (resp) {
    return writeFile('savedData.txt', resp.data)
  })
  .then(function () {
    console.log('operator completed')
  })
  .catch(function (err) {
    console.log('something went wrong: ', err)
  })

