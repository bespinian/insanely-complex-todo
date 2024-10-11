FROM node:alpine AS build

WORKDIR /app

COPY package*.json /app/
RUN npm install

COPY . /app/
RUN npm run build

FROM nginx:alpine AS run

COPY --from=build /app/build /usr/share/nginx/html
