FROM mongo:latest

ENV MONGO_DATA_DIR=/data/db

RUN mkdir -p $MONGO_DATA_DIR

EXPOSE 27017

CMD ["mongod"]