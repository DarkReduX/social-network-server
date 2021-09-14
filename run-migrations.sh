# From project root
cd ./migrations/ || exit

# Get database environments
postgresPort=$(printenv POSTGRES_PORT)
postgresHost=$(printenv POSTGRES_HOST)
postgresPassword=$(printenv POSTGRES_PASSWORD)
postgresUser=$(printenv POSTGRES_USER)
postgresDbName=$(printenv POSTGRES_DBNAME)

mkdir -p history

listOfMigrationScript=("*.sql")
# listOfDoneMigrationScript=("./history/*")

echo "Starting migrate process..."
for item in $listOfMigrationScript; do

  findItem=$(find ./history -name "${item}")
  findItem=${findItem##*/}

  if [ "$item" != "$findItem" ]; then
    errorlog=$(mktemp)
    trap 'rm -f "$errorlog"' EXIT
    pwcheck="$(psql postgresql://"$postgresUser":"$postgresPassword"@"$postgresHost":"$postgresPort"/"$postgresDbName" -v ON_ERROR_STOP=1 -f "$initDatabaseScript" 0 < "$errorlog")"

    if [[ 0 -ne $? ]]; then

      #echo "Something went wrong in migration $item"
      exit
    fi
    touch ./history/"$item"
    echo Success migration of "$item"
  else
    echo "$item" not need to be migrated
  fi


done

echo "All migration had been done successfully"