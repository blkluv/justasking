//INVOKE with $ node environments\Interpolator.js "%WORKSPACE%\dbo_SqlScripts.sql" "%WORKSPACE%\environments\dev.json" "%WORKSPACE%\scripts.sql"

var sourceFilePath = process.argv[2];
var environmentFilePath = process.argv[3];
var outputFilePath = process.argv[4];
var environment = {};
var fs = require("fs");
var sourceText = "";

sourceText = fs.readFileSync(sourceFilePath,'utf8');
environment = require(environmentFilePath);

//process file content
for (var property in environment) {
    if (environment.hasOwnProperty(property)) {
        var pattern = "{{"+property+"}}",
        regex = new RegExp(pattern, "g");

    	sourceText = sourceText.replace(regex, environment[property]);
    }
}

//Write new file with interpolated text
fs.writeFile(outputFilePath, sourceText, function(err) {
    if(err) {
        return console.log(err);
    }
}); 

