#!/bin/bash

# This shell script compiles all database Data Definition Language (DDL) and Data Manipulation Language (DML) types into one file

# Variables DDL
ApplicationFolder="dbo"
SQLScriptType="SqlScripts.sql"

# Creating DLL SQL Script
echo "Creating SQL Scripts for $ApplicationFolder..."
echo "Creating DDL SQL Script for $ApplicationFolder..."

echo "/* Tables SQL Script */" >"${ApplicationFolder}_${SQLScriptType}"
cat "${ApplicationFolder}/Tables/"*.sql >>"${ApplicationFolder}_${SQLScriptType}"

echo "/* Stored Procedures SQL Script */" >>"${ApplicationFolder}_${SQLScriptType}"
cat "${ApplicationFolder}/Stored Procedures/"*.sql >>"${ApplicationFolder}_${SQLScriptType}"

# Adding DML SQL Script
echo "Creating DML SQL Script for $ApplicationFolder..."
echo "/* Data Insert SQL Script */" >>"${ApplicationFolder}_${SQLScriptType}"
cat "${ApplicationFolder}/Data Insert/"*.sql >>"${ApplicationFolder}_${SQLScriptType}"

exit 0