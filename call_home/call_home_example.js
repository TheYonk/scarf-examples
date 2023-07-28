const fs = require('fs');
const http = require('http');

const getUrlData = url => {
    return new Promise((resolve, reject) => {
        http.get(url, res => {
            let data = '';

            res.on('data', chunk => {
                data += chunk;
            });

            res.on('end', () => {
                resolve(JSON.parse(data).current_version);
            });

        }).on('error', err => {
            reject(err);
        });
    });
};

const readLocalVersionFile = file => {
    return new Promise((resolve, reject) => {
        fs.readFile(file, 'utf8', (err, data) => {
            if (err) {
                reject(err);
            } else {
                resolve(JSON.parse(data).current_version);
            }
        });
    });
};

const checkVersion = async () => {
    const url = 'http://127.0.0.1/data/version.json';
    const file = 'current_version.json';

    try {
        const versionFromUrl = await getUrlData(url);
        const versionFromFile = await readLocalVersionFile(file);

        if (versionFromUrl === versionFromFile) {
            console.log(`Version match. The version is: ${versionFromUrl}`);
        } else {
            console.log(`Version mismatch. URL version: ${versionFromUrl}. File version: ${versionFromFile}`);
        }
    } catch (err) {
        console.log(err);
    }
};

checkVersion();
