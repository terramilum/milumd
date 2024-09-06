# Setup domain to peer

## Make the script executable

Make the script executable
```bash
chmod +x scripts/nginx_add_domain_port.sh 
```

## Set up a reverse proxy for the specified port

Run the script with the domain and port as parameters

```bash
./scripts/nginx_add_domain_port.sh subdomain.domain.io 26656
```

## This script will

- Install Nginx and Certbot if they are not installed.
- Create a new Nginx configuration file for your domain.
- Set up a reverse proxy for the specified port (127.0.0.1:<port>).
- Enable the site in Nginx.
- Use Certbot to obtain an SSL certificate for your domain.
- Test the SSL certificate renewal process.
- You can run this script for any other domain and port combination by changing the parameters when running the script.


# Define a Subdomain in Cloudflare

## Step 1: Log In to Cloudflare
1. Go to [Cloudflare](https://www.cloudflare.com/) and log in to your account.
2. If you don’t have an account, sign up and add your domain to Cloudflare by updating your domain's DNS nameservers to Cloudflare’s.

## Step 2: Select Your Domain
1. From the dashboard, select the domain where you want to add the subdomain (e.g., `domain.com`).

## Step 3: Add a Subdomain DNS Record
1. In the left-hand menu, click on the **DNS** tab.
2. Click **Add Record** at the top of the DNS records list.

3. Use the following settings to create your subdomain:

   - **Type**: Choose between an **A** or **CNAME** record:
     - **A Record**: If the subdomain (e.g., `node1.domain.com`) should point to an IP address (like your server's public IP).
     - **CNAME Record**: If the subdomain should point to another domain name.

   - **Name**: Enter your subdomain. For example, to create `node1.domain.com`, enter `node1`.

   - **IPv4 address** (for A Record): Enter the public IP address of your server. For a CNAME record, enter the domain name you are pointing to.

   - **TTL**: Set to **Auto**.

   - **Proxy Status**:
     - **Proxied (orange cloud)**: Cloudflare will act as a reverse proxy, offering DDoS protection, caching, etc.
     - **DNS Only (gray cloud)**: Cloudflare will manage only DNS without proxying traffic.

4. Click **Save** to add the DNS record.

### Example Setup:
- **A Record** example:
   - **Type**: A
   - **Name**: `node1`
   - **IPv4 address**: `113.232.1.121` (replace with your server's IP)
   - **Proxy Status**: Proxied (orange cloud) or DNS Only (gray cloud)
   
- **CNAME Record** example:
   - **Type**: CNAME
   - **Name**: `node1`
   - **Target**: `domain.com` or another domain
   - **Proxy Status**: Proxied or DNS Only

## Step 4: Verify DNS Propagation
Once you’ve added the subdomain, it may take up to 24 hours for DNS changes to propagate across the internet. To verify, use a tool like [DNS Checker](https://dnschecker.org/) to ensure the subdomain points to the correct IP.

## Step 5: Proceed with Nginx and Certbot Setup
Once your subdomain is set up in Cloudflare and points to your server's IP address, you can proceed with the Nginx and Certbot setup to configure SSL and proxy traffic.

## Additional Notes:
- If your site is **Proxied** by Cloudflare (orange cloud), Cloudflare handles SSL termination. In this case, you don’t need to manually configure SSL certificates on your server.
- If you're using **DNS Only** (gray cloud), you need to configure SSL on your server using Certbot and the provided scripts.
