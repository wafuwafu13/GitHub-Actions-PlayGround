const { exit } = require("shelljs");

if (1 < 2) {
  console.error("This error is critical");
  exit(1);
}

console.log('some process...')
