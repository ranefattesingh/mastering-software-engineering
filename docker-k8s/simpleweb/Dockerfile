# Specify a base image
FROM node:24-alpine

WORKDIR /usr/app

# Install some dependencies
RUN npm install --global pnpm

# Install pnpm globally
COPY ./package.json ./
RUN pnpm install

COPY ./ ./

# Default command
CMD ["pnpm", "start"]