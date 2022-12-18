FROM node:latest AS build

ENV REACT_APP_API_BASE_URL="http://localhost:15001/"

WORKDIR /app
COPY . .

RUN npm ci
RUN npm run build
RUN ls


FROM nginx:alpine AS prod

WORKDIR /usr/share/nginx/html
RUN rm -rf ./*
COPY --from=build /app/build .

EXPOSE 80
CMD ["nginx", "-g", "daemon off;"]