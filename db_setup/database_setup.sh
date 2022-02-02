#!/usr/bin/env bash

set -e

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" >/dev/null 2>&1 && pwd)"

SQL_POD=$(kubectl get pods | awk '{print $1}' | grep -e "sql")

echo "Copying setup.sql to sql pod $SQL_POD"
kubectl cp $SCRIPT_DIR/notchdb/setup.sql $SQL_POD:/root/setup.sql

echo "bootstrapping and setting up environment"
kubectl exec $SQL_POD -- /opt/mssql-tools/bin/sqlcmd -U sa -P notch@12345 -i /root/setup.sql > /dev/null

echo "DB has been set up with test data!"
echo ""
