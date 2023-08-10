#!/bin/bash

# Run concat.bat (you'll need to convert this to a shell script or other executable as well)
./concat.sh
wait

# Execute the Node script
node "$(pwd)/Environments/Interpolator.js" "dbo_SqlScripts.sql" "$(pwd)/Environments/prod.json" "scripts_dev.sql"

# Remove the temporary file
rm "$(pwd)/dbo_SqlScripts.sql"