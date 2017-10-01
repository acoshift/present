const item = {
  id: 7,
  name: 'JavaScript eBook',
  price: 20,
  type: 'book',
  author: 'unknown'
}

// rest
const { name, price, ...other } = item
console.log(name)
// JavaScript eBook
console.log(price)
// 20

console.log(other)
// { id: 7, type: 'book', author: 'unknown' }

// spread
const x = { name, price, ...other }
console.log(x)
// { name: 'JavaScript eBook', price: 20, id: 7, type: 'book', author: 'unknown' }
