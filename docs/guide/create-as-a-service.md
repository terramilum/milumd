# Automatically Restart Chain

To create a Linux service that starts automatically when the computer boots up or restarts and can be controlled with the mirumd commands, you'll typically need to create a systemd service. Here are the general steps to achieve this:

1. __Create a systemd Service File__


Create a new systemd service file. Typically, these are stored in the /etc/systemd/system/ directory and end with the .service extension.

```bash
sudo nano /etc/systemd/system/terramirum.service
```

Add the following content to the file:

Write "which mirumd" on terminal and replace with "which mirumd"

```bash
[Unit]
Description=Terramirum Chain Service
After=network.target

[Service]
Type=simple
ExecStart=<which mirumd> start
ExecStop=/usr/bin/killall mirumd
Restart=always
RestartSec=3

[Install]
WantedBy=default.target

```

2. __Reload systemd__

After creating the service file, reload the systemd configuration to recognize the new service:

```bash
sudo systemctl daemon-reload
```

3. __Enable and Start the Service__

Enable the service to start on boot:

```bash
sudo systemctl enable terramirum.service
```

Start the service:

```bash
sudo systemctl start terramirum.service
```

Verify that the service is running:

```bash
sudo systemctl status terramirum.service
```

This should show you the current status of your terramirum service.

Now, your terramirum service should start automatically on boot and restart if it crashes. You can also manually control it using the following commands:

- To stop the service: sudo systemctl stop terramirum.service
- To restart the service: sudo systemctl restart terramirum.service
