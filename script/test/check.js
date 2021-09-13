const { exit } = require("shelljs");
const process = require("process")

if (1 < 2) {
  console.error("This error is critical");
  process.exit(1);
}

console.log('some process...')
