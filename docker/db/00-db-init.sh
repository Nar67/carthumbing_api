DATABASE="carthumbing_db"
PG_SUPERUSER="NOSUPERUSER"
USER="carthumbing_user"
PASSWORD="ohphahRohfohZoh6"  # test password


echo "Creating db '${DATABASE}' for user '${USER}'"
psql -c "CREATE USER ${USER} CREATEDB ${PG_SUPERUSER} NOCREATEROLE INHERIT LOGIN UNENCRYPTED PASSWORD '${PASSWORD}';"
createdb --owner ${USER} --template template0 --encoding=UTF8 --lc-ctype=en_US.UTF-8 --lc-collate=en_US.UTF-8 ${DATABASE}