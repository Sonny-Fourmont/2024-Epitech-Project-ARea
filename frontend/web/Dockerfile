# Use an official Node.js runtime as a parent image
FROM node:22.11.0

# Set the working directory
WORKDIR /usr/src/app

# Copy package.json and package-lock.json
COPY package*.json ./

# Install dependencies
RUN npm install

# Copy the rest of the application code
COPY . .

EXPOSE 8081

# Define the command to run the app
CMD ["npm", "run", "dev"]