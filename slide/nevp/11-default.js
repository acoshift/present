function invokeHttp (url, method = 'GET') {
  console.log(`invoke ${method}: ${url}`)
}

invokeHttp('https://google.com')
// invoke GET https://google.com
