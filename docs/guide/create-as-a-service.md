# Automatically Restart Chain

To create a Linux service that starts automatically when the computer boots up or restarts and can be controlled with the mirumd commands, you'll typically need to create a systemd service. Here are the general steps to achieve this:

1. __Create a systemd Service File__


Create a new systemd service file. Typically, these are stored in the /etc/systemd/system/ directory and end with the .service extension.

```bash
sudo nano /etc/systemd/system/mirumd.service
```

Add the following content to the file:

Write "which mirumd" on terminal and replace with "which mirumd"

```bash

[Unit]
Description=Terramirum Blockchain
After=network-online.target

[Service]
User=root
ExecStart=/home/code/bin/mirumd start --home /home/<user>/.mirumd
Restart=always
RestartSec=3
LimitNOFILE=4096

[Install]
WantedBy=multi-user.target

```

2. __Reload systemd__

After creating the service file, reload the systemd configuration to recognize the new service:

```bash
sudo systemctl daemon-reload
```

3. __Enable and Start the Service__

Enable the service to start on boot:

```bash
sudo systemctl enable mirumd.service
```

Start the service:

```bash
sudo systemctl start mirumd.service
```

Verify that the service is running:

```bash
sudo systemctl status mirumd.service
```

This should show you the current status of your terramirum service.

Now, your terramirum service should start automatically on boot and restart if it crashes. You can also manually control it using the following commands:

- To stop the service: sudo systemctl stop terramirum.service
- To restart the service: sudo systemctl restart terramirum.service

If you encounter an error while checking the status of the server, use the following command to find the exact error:

```bash
journalctl -xe | grep mirumd
```

The error may arise from different root paths between the root user and the current user. To address this issue, consider adding the --home flag to the service definition.