FREECHART_JAR=${JFREECHART_PATH}/jfreechart-1.0.19.jar:${JFREECHART_PATH}/jcommon-1.0.23.jar
MARIADB_JAR=~/bin/mariadb-java-client-1.1.10.jar

java -debug -cp ../../..:$FREECHART_JAR:$MARIADB_JAR org/phypo/SqlTools/SqlTools -Itest.ini $*
