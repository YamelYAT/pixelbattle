FROM node:lts-alpine as node-build
WORKDIR /app
COPY ./app/ .
RUN npm install
CMD ["npm","run","build"]

FROM nginx:latest
COPY nginx.conf /etc/nginx/nginx.conf
COPY --from=node-build /app/dist/ /etc/nginx/var/
EXPOSE 80





