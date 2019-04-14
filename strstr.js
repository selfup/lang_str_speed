const fs = require('fs');

const lines = [];
const stream = fs.createReadStream('tmp/logs.log');

stream.on('data', data => {
  data
    .toString()
    .trim()
    .split('\n')
    .forEach(line => {
      if (line.indexOf('OK db=') > -1) {
        lines.push(line.split('OK ')[1].trim());
      }
    });
});

stream.on('close', () => {
  console.log(lines.slice(0, 40));
});
