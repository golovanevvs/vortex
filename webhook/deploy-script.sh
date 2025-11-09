#!/bin/sh
echo "========================================"
echo "   Starting Delayed Notifier Deployment"
echo "   Time: $(date)"
echo "========================================"
echo ""

echo "Changing to deploy directory..."
cd /etc/webhook/deploy

echo "Current directory: $(pwd)"
echo "Files in directory:"
ls -la

echo "Stopping existing containers..."
docker compose down

echo "Pulling latest image..."
docker compose pull

echo "Starting new containers..."
docker compose up -d

echo ""
echo "========================================"
echo "   Deployment completed successfully!"
echo "========================================"
echo ""

sleep 3