const { exit } = require("shelljs");
import { exit } from 'process';

if (1 < 2) {
  console.error("This error is critical");
  process.exit(1);
}

console.log('some process...')
