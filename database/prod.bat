START /wait concat.bat
node "%cd%\Environments\Interpolator.js" "dbo_SqlScripts.sql" "%cd%\Environments\prod.json" "scripts_prod.sql"
del "%cd%\dbo_SqlScripts.sql"
