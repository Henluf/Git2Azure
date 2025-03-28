FROM ubuntu:latest
LABEL authors="Azureman"

ADD azureContainerTest ./azureContainerTest
RUN chmod a+x ./azureContainerTest

EXPOSE 31415/tcp

CMD "./azureContainerTest"
