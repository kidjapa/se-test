#!/bin/bash

/usr/bin/mysqld_safe --skip-grant-tables &
sleep 5
mysql -u root -e "CREATE DATABASE se_test"
mysql -u root se_test < /tmp/make_database.sql