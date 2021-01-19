# JustAsking: A real-time audience polling app

This repository contains everything needed to run JustAsking. From creating the database tables to signing up and billing users through Stripe, to accepting SMS messages from Twilio and broadcasting the results to all users currently viewing the poll.


## Folder Structure ##

* database - contains the DDL and DML scripts for creating the tables and stored procedures, as well as the config files which hold values for different environments.
    - ```dbo/Data Insert``` contains scripts which insert values into the database. Placeholders for the values are surrounded by double curly braces {{}}, and the values are stored in Environtments/dev.json and Environments/prod.json.
    - Running dev.bat will interpolate the dev.json file into the Data Insert files and output scripts_dev.sql. This file will be a complete database deployment with tables and associated values.

* emails - these are the HTML version of the emails which are sent out when users sign up, upgrade, or their plans expire.

* GO - this is where all the backend Go code is stored.    
    - ```api``` contains all the code for the REST API which is secured by JWT. It is a full standalone application.
    - ```realtimehub``` contains all the code for tracking and broadcasting messages to users depending on which poll they are currently viewing. Also a standalone application
    - ```common``` contains code which is used by both the API and Real Time Hub.
    - ```core``` contains the business logic for the application as well as models for interacting with the database.
    - ```syncservice``` contains the code to periodically check for expired subscriptions and pull in new phone numbers from Twilio.

* services_monitor - this is a bash script that periodically checks whether the syncservice and realtime hub are up.    

* web - contains all the frontend code. Written in TypeScript using Angular 5.

## Running the Application ##

At a minimum, the ```web``` app, ```GO/api``` and ```GO/realtimehub``` need to be running for the application to work. A MySQL instance will also need to be running, and the connectionstring in ```GO/api/config.json``` will need to be updated to point to it.

SMS and billing will not work until the appropriate configs are updated in ```database/environments```
 
## Authors ##
* Sebastian Chande (sebastian.chande@gmail.com) - API, realtimehub, database
* Ariel Diaz (arieldiaz92@gmail.com) - Angular frontend, database