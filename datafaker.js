var faker = require('faker');
const fs = require('fs');

/**
 * This script takes 2 inputs - an input file and output file
 *
 * Example:
 * node ./scripts/datafaker.js ./openapi/kyc.yml ./openapi/kyc_faker.yml
 * */
var inputFile = process.argv[2];
var outputFile = process.argv[3];
console.log('Parsing ' + inputFile + ' for faker references');

let fakerRegex = /faker\.([A-Za-z0-9\.#\(\)\-\+]+)/;

fs.readFile(inputFile, 'utf8', (err, data) => {
  if (err) {
    console.log(err);
    return;
  } else {
    // Iterate through lines in file, find faker references then replace with
    // faked data
    let lines = data.split('\n');
    const seenFakeRequests = new Set();
    for (let line in lines) {
        let fakerContent = lines[line].match(fakerRegex);
        if (fakerContent) {
            let fakerGenerated = faker.fake('{{' + fakerContent[1]+ '}}');
            if (!seenFakeRequests.has(fakerContent[1])) {
              seenFakeRequests.add(fakerContent[1]);
              faker.seed(Date.now());
            }
            console.log('In line: ' + lines[line].trim() + '\n replacing '
                + fakerContent[1] + ' with ' + fakerGenerated);
            lines[line] = lines[line].replace(fakerRegex, fakerGenerated);
        }
    }

    // Merge lines back together and write to output file
    fs.writeFile (outputFile, lines.join('\n'), (err, data) => {
        if (err) {
            console.log(err);
        }
    });
  }
});
