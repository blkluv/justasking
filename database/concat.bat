:: This batch file compiles all database Data Definition Language (DDL) and Data Manipulation Language (DML) types into one file  ::

:: Variables DDL ::
SET ApplicationFolder=dbo
SET SQLScripType=SqlScripts.sql

:: Creating DLL SQL Script::
Echo "Creating SQL Scripts for %ApplicationFolder%..."
Echo "Creating DDL SQL Script for %ApplicationFolder%..."

Echo /* Tables SQL Script */ >"%ApplicationFolder%_%SQLScripType%"
TYPE "%ApplicationFolder%/Tables\*.sql" >>"%ApplicationFolder%_%SQLScripType%"

Echo /* Views SQL Script */ >>"%ApplicationFolder%_%SQLScripType%"
TYPE "%ApplicationFolder%/Views\*.sql" >>"%ApplicationFolder%_%SQLScripType%"

Echo /* Functions SQL Script */ >>"%ApplicationFolder%_%SQLScripType%"
TYPE "%ApplicationFolder%/Functions\*.sql" >>"%ApplicationFolder%_%SQLScripType%"

Echo /* Stored Procedures SQL Script */ >>"%ApplicationFolder%_%SQLScripType%"
TYPE "%ApplicationFolder%/Stored Procedures\*.sql" >>"%ApplicationFolder%_%SQLScripType%"

Echo /* Triggers SQL Script */ >>"%ApplicationFolder%_%SQLScripType%"
TYPE "%ApplicationFolder%/Triggers\*.sql" >>"%ApplicationFolder%_%SQLScripType%"

Echo /* Sequences SQL Script */ >>"%ApplicationFolder%_%SQLScripType%"
TYPE "%ApplicationFolder%/Sequences\*.sql" >>"%ApplicationFolder%_%SQLScripType%"

:: Adding DML SQL Script ::
Echo "Creating DML SQL Script for %ApplicationFolder%..."
Echo /* Data Insert SQL Script */ >>"%ApplicationFolder%_%SQLScripType%"
TYPE "%ApplicationFolder%/Data Insert\*.sql" >>"%ApplicationFolder%_%SQLScripType%"

Exit 0