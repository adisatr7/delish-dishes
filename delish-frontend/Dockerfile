# Use the official Bun image as the base image
FROM oven/bun:alpine AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy package.json and bun.lockb to the working directory
COPY package.json bun.lockb ./

# Install dependencies using Bun
RUN bun install

# Copy the rest of the application code to the working directory
COPY . .

# Build the Next.js application
RUN bun run build

# Use a smaller base image for the final stage
FROM oven/bun:alpine

# Set the working directory
WORKDIR /app

# Copy only the necessary files from the builder stage
COPY --from=builder /app ./

# Install only production dependencies
RUN bun install --production

# Set the environment variable to use port 8080
ENV PORT 8080

# Expose port 8080
EXPOSE 8080

# Start the Next.js application
CMD ["bun", "run", "start"]