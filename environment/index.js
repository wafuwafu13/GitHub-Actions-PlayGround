const fs = require('fs');

fs.copyFile('test.txt', 'copyed.txt', (err) => {
  if (err) {
      console.log(err.stack);
  }
  else {
      console.log('Copied!!!');
  }
});
