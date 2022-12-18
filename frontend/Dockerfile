FROM node:latest AS build

WORKDIR /app
COPY . .

RUN npm ci
RUN npm run build

FROM nginx:alpine AS prod

WORKDIR /usr/share/nginx/html
RUN rm -rf ./*
COPY --from=build /app/build .

EXPOSE 80
CMD ["nginx", "-g", "daemon off;"]