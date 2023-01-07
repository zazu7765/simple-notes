FROM node:16 AS builder
WORKDIR /usr/src/web
COPY package*.json ./
RUN npm ci
COPY . .
EXPOSE 4173
CMD ["npm", "run","dev"]