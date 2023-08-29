# Use an official Node runtime as base image
FROM node:14

# Set working directory
WORKDIR /usr/src/app

# Copy package.json and install dependencies
COPY package.json ./
RUN npm install --omit=dev && npm ci --only=production

# Copy the current directory contents into the container
COPY . .

# Set environment to production
ENV NODE_ENV production

# Expose port and start the app
EXPOSE 3000
CMD ["npm", "start"]
