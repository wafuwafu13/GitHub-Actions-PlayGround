const process = require('process');

if (true) {
  console.error("This error is critical");
  process.exitCode = 1;
}

console.log('...some process...')
