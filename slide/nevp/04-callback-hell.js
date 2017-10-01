readFile('savedUrl.txt', function (err, fileData) {
  if (err) {
    console.log('read file error: ' + err)
    return
  }

  var url = 'http://' + fileData
  invokeHttp(url, function (err, resp) {
    if (err) {
      console.log('invoke http error: ' + err)
      return
    }

    writeFile('savedData.txt', resp.data, function (err) {
      if (err) {
       console.log('write file error: ' + err)
      }

      console.log('operator completed')
    })
  })
})
