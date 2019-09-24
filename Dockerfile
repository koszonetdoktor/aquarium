FROM golang:1.13.0-alpine

WORKDIR /app/ui

COPY . /app

RUN apk add --update npm

RUN npm install

EXPOSE 8080

ENV NODE_ENV develpment

CMD [ "npm", "start" ]