# From project root
cd ./migrations/ || exit

# Get database environments
postgresPort=$(printenv POSTGRES_PORT)
postgresHost=$(printenv POSTGRES_HOST)
postgresPassword=$(printenv POSTGRES_PASSWORD)
postgresUser=$(printenv POSTGRES_USER)
postgresDbName=$(printenv POSTGRES_DBNAME)

initDatabaseScript=$(basename -- "init_database.sql")
echo "Start init database from $initDatabaseScript script"
# listOfDoneMigrationScript=("./history/*")

 errorlog=$(mktemp)
    trap 'rm -f "$errorlog"' EXIT
    pwcheck="$(psql postgresql://"$postgresUser":"$postgresPassword"@"$postgresHost":"$postgresPort"/"$postgresDbName" -v ON_ERROR_STOP=1 -f "$initDatabaseScript" 0 < "$errorlog")"

    if [[ 0 -ne $? ]]; then

      echo "Something went wrong in $initDatabaseScript"
      exit
    fi
echo "Successful initialized database"