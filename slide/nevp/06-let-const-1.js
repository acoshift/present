function f () {
  var x = 10

  if (x === 10) {
    var x = 5
    console.log(x)
  }
  console.log(x)
}

f()

// 5
// 5

