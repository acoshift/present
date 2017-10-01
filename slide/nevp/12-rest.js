function sum (...v) {
  let s = 0
  for (let i = 0; i < v.length; ++i) {
    s += v[i]
  }
  return s
}

const result = sum(1, 2, 3, 4, 5)
console.log(result)
// 15
