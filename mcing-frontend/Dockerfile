FROM node:22.16.0-slim as build
LABEL org.opencontainers.image.source="https://github.com/kmdkuk/MCing-frontend"

WORKDIR /app

COPY package.json package-lock.json ./

RUN npm install --production

COPY . .

RUN npm run build

FROM nginx:1.28.0
COPY --from=build /app/dist /usr/share/nginx/html
EXPOSE 80
CMD ["nginx", "-g", "daemon off;"]
