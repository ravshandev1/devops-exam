FROM node:18-alpine
WORKDIR /app
COPY package.json .
RUN npm install
COPY main.js .
EXPOSE 80
CMD ["node", "main.js"]