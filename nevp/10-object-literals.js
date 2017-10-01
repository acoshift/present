const name = 'JavaScript eBook'
const price = 20
const key = 'type'

const item = {
  name,
  price,
  f () {
    console.log(`Item ${this.name}: $${this.price}`)
  },
  [ key ]: 'ebook'
}

item.f()
// Item JavaScript eBook: $20

console.log(item.type)
// ebook

