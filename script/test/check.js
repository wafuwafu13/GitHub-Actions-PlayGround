const process = require('process');

if (1 < 2) {
  console.error("This error is critical");
  process.exitCode = 1;
}

console.log('some process...')
