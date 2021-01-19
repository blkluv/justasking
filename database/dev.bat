START /wait concat.bat
node "%cd%\Environments\Interpolator.js" "dbo_SqlScripts.sql" "%cd%\Environments\dev.json" "scripts_dev.sql"
del "%cd%\dbo_SqlScripts.sql"
