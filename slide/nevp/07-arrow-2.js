const obj = {
  value: 10,
  f1: function () {
    console.log('f1: ' + this.value)
  },
  f2: () => {
    console.log('f2: ' + this.value)
  }
}

obj.f1()
obj.f2()

// f1: 10
// f2: undefined
