const crypto = require('crypto');
const randomString = crypto.randomUUID();

setInterval(() => {
  console.log(`${new Date().toString()}: ${randomString}`)
}, 5000);
