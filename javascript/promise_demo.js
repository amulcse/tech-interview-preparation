let p = new Promise((resolve,reject) => {
  let s = 'amul'
  if(s === 'amul'){
    resolve('success')
  }
  else{
    reject('error')
  }
})

p.then((message) => {
  console.log('string matched '+message)
})
.catch((message) => {
  console.log('string not matched '+message)
})