const {exec} = require("child_process");
const fs = require("fs");

// Read package.json to get the version
const packageJson = JSON.parse(fs.readFileSync("package.json", "utf-8"));
const version = packageJson.version;

// Construct the file path with the version
const filePath = `./tsp-output/@typespec/openapi3/openapi.${version}.yaml`;

// Run the command with the correct version
exec(`npx @redocly/cli build-docs ${filePath}`, (error, stdout, stderr) => {
    if (error) {
        console.error(`Error: ${error.message}`);
        return;
    }
    if (stderr) {
        console.error(`Stderr: ${stderr}`);
    }
    console.log(`Stdout: ${stdout}`);
});