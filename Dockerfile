# syntax=docker/dockerfile:1
FROM kselnaag/gobuilder:1.19.1 AS algoscommit

# ENV GOPATH=/goBuilder
ARG  COMMITHASH=HEAD
WORKDIR $GOPATH
RUN git clone https://github.com/kselnaag/algos.git
WORKDIR $GOPATH/algos

RUN git checkout $COMMITHASH && chmod +x ./checks.sh
ENTRYPOINT ["/bin/bash", "./checks.sh"]

# docker build -t algoscommit:2639ae4fa --build-arg COMMITHASH=2639ae4fa .
# docker run -i -p 80:80 algoscommit:154637c