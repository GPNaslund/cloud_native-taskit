# Build
FROM node:18-alpine AS builder
WORKDIR /app
COPY client/package*.json ./
RUN npm install
COPY client/ .
ARG NEXT_PUBLIC_API_URL="https://www.gpnaslund.pro/a2/taskit/api"
ENV NEXT_PUBLIC_API_URL=${NEXT_PUBLIC_API_URL}
ARG NEXT_PUBLIC_BASE_PATH="/a2/taskit"
ENV NEXT_PUBLIC_BASE_URL=${NEXT_PUBLIC_BASE_PATH}
RUN mkdir -p public
RUN npm run build

# Production
FROM node:18-alpine AS runner
WORKDIR /app

ENV NODE_ENV=production
COPY --from=builder /app/package*.json ./
COPY --from=builder /app/.next ./.next
COPY --from=builder /app/public ./public
COPY --from=builder /app/node_modules ./node_modules

EXPOSE 3000

CMD ["npm", "start"]
