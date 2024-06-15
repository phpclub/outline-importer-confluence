# outline-importer-confluence
Imports docs from Сonfluence to Outline



bee api [appname] [-tables=""] [-driver=mysql] [-conn="root:@tcp(127.0.0.1:3306)/test"]  [-gopath=false] [-beego=v1.12.3]
bee api confluence -driver=mysql -conn="root:@tcp(127.0.0.1:3306)/wikidb -tables="SPACES, LABEL, PAGETEMPLATES, BODYCONTENT, CONTENT, CONTENT_LABEL, CONTENTPROPERTIES, user_mapping"



=================================
bee help generate
USAGE
bee generate [command]

OPTIONS
-conn
Connection string used by the SQLDriver to connect to a database instance.

-ctrlDir
Controller directory. Bee scans this directory and its sub directory to generate routers

-ddl
Generate DDL Migration

-driver
Database SQLDriver. Either mysql, postgres or sqlite.

-fields
List of table Fields.

-level
Either 1, 2 or 3. i.e. 1=models; 2=models and controllers; 3=models, controllers and routers.

-routersFile
Routers file. If not found, Bee create a new one. Bee will truncates this file and output routers info into this file

-routersPkg
router's package. Default is routers, it means that "package routers" in the generated file

-tables
List of table names separated by a comma.

DESCRIPTION
▶ To scaffold out your entire application:

     $ bee generate scaffold [scaffoldname] [-fields="title:string,body:text"] [-driver=mysql] [-conn="root:@tcp(127.0.0.1:3306)/test"]

▶ To generate a Model based on fields:

     $ bee generate model [modelname] [-fields="name:type"]

▶ To generate a controller:

     $ bee generate controller [controllerfile]

▶ To generate a CRUD view:

     $ bee generate view [viewpath]

▶ To generate a migration file for making database schema updates:

     $ bee generate migration [migrationfile] [-fields="name:type"]

▶ To generate swagger doc file:

     $ bee generate docs

    ▶ To generate swagger doc file:

     $ bee generate routers [-ctrlDir=/path/to/controller/directory] [-routersFile=/path/to/routers/file.go] [-routersPkg=myPackage]

▶ To generate a test case:

     $ bee generate test [routerfile]

▶ To generate appcode based on an existing database:

     $ bee generate appcode [-tables=""] [-driver=mysql] [-conn="root:@tcp(127.0.0.1:3306)/test"] [-level=3]
