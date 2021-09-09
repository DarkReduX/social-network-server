# From project root
cd ./migrations/ || exit

initDatabaseScript=$(basename -- "init_database.sql")
echo "Start init database from $initDatabaseScript script"
# listOfDoneMigrationScript=("./history/*")

 errorlog=$(mktemp)
    trap 'rm -f "$errorlog"' EXIT
    pwcheck="$(psql postgresql://postgres:a!11111111@localhost:5432/social_network -v ON_ERROR_STOP=1 -f "$initDatabaseScript" 0 < "$errorlog")"

    if [[ 0 -ne $? ]]; then

      echo "Something went wrong in $initDatabaseScript"
      exit
    fi
echo "Successful initialized database"