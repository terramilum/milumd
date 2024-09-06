#!/bin/bash

# Check if the user provided the necessary arguments
if [ "$#" -ne 2 ]; then
    echo "Usage: $0 <domain> <port>"
    echo "Example: $0 mirumnode.nodeks.com 26656"
    exit 1
fi

# Define variables
DOMAIN=$1
PORT=$2

# Update and install necessary packages
echo "Updating system and installing Nginx and Certbot..."
sudo apt update
sudo apt install -y nginx certbot python3-certbot-nginx

# Create Nginx configuration for the domain
echo "Creating Nginx configuration for $DOMAIN with port $PORT..."
sudo bash -c "cat > /etc/nginx/sites-available/$DOMAIN <<EOF
server {
    listen 80;
    server_name $DOMAIN;

    location / {
        proxy_pass http://127.0.0.1:$PORT;
        proxy_set_header Host \$host;
        proxy_set_header X-Real-IP \$remote_addr;
        proxy_set_header X-Forwarded-For \$proxy_add_x_forwarded_for;
        proxy_set_header X-Forwarded-Proto \$scheme;
    }
}
EOF"

# Enable the site by creating a symlink to sites-enabled
sudo ln -s /etc/nginx/sites-available/$DOMAIN /etc/nginx/sites-enabled/

# Test Nginx configuration
echo "Testing Nginx configuration..."
sudo nginx -t

# Reload Nginx to apply the new configuration
echo "Reloading Nginx..."
sudo systemctl reload nginx

# Obtain SSL certificate with Certbot for the domain
echo "Obtaining SSL certificate for $DOMAIN using Certbot..."
sudo certbot --nginx -d $DOMAIN

# Test automatic renewal of SSL certificates
echo "Testing SSL certificate renewal..."
sudo certbot renew --dry-run

echo "Setup completed for $DOMAIN with Nginx proxying to http://127.0.0.1:$PORT"
