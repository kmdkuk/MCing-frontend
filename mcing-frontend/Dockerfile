FROM node:22.15.0-slim as build
LABEL org.opencontainers.image.source="https://github.com/kmdkuk/MCing-frontend"

WORKDIR /app

COPY package.json package-lock.json ./

RUN npm install --production

COPY . .

RUN npm run build

FROM nginx:1.25.2
COPY --from=build /app/dist /usr/share/nginx/html
EXPOSE 80
CMD ["nginx", "-g", "daemon off;"]
