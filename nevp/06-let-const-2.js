function f () {
  let x = 10

  if (x === 10) {
    let x = 5
    console.log(x)
  }
  console.log(x)
}

f()

// 5
// 10

