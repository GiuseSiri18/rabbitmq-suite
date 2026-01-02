# Use Node for building the frontend
FROM node:18-alpine

WORKDIR /app

# Install a simple static server
RUN npm install -g serve

# In a real scenario, you'd copy your built Vue files here
# For this lab, we will use a single-file Vue approach for speed
COPY index.html .

CMD ["serve", "-s", ".", "-l", "8080"]