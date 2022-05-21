REM script for running the application

call mvn clean
@echo on

call mvn package
@echo on

java -jar --add-opens java.base/java.lang=ALL-UNNAMED target\SciPaper-0.0.1b-SNAPSHOT.jar --PORT=50002