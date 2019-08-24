FROM maven as builder
RUN git clone https://github.com/briandealwis/cram &&\
   cd cram ; mvn package
FROM openjdk:11
COPY --from=builder /cram/target/cram-0.0.1-SNAPSHOT-jar-with-dependencies.jar /cram.jar
ADD docker-credential-simple /usr/bin/docker-credential-simple
ADD creds.json /root/.credhelper/creds.json
ADD ca.crt /etc/ca.crt
RUN echo "changeit" | keytool -import -trustcacerts -file /etc/ca.crt -alias registry -keystore $JAVA_HOME/lib/security/cacerts -noprompt
ENTRYPOINT ["java", "-jar", "/cram.jar"]
