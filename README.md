# tjsj.dev
### Yet another CS students personal web server

currently under development\
its very cool

## Setting up & Running

These instructions are an example of how you could setup this server.
We will be using LetsEncrypt to get SSL certificates, a systemd service for starting the server and a cron job to automatically renew the SSL certs.

The following steps assume the working directory is `/srv/tjsj.dev/`.
You will need to adapt [./init/tjsj.dev.service] appropriately if you plan to run the server in a different directory.

### Compiling the server

Clone the repository
```
$ git clone https://github.com/tedski999/tjsj.dev.git
$ cd tjsj.dev
```

Build the Go program
```
$ make
```

The program is built to `./bin/tjsj.dev`.

### Getting the SSL Certificates

Make sure the device you're hosting this server on is accessible on port 80 publicly

Running certbot
```
# certbot certonly -d "tjsj.dev,www.tjsj.dev"
```

### Auto-renewing SSL Certificates with Cron

Edit root cron jobs
```
# crontab -e
```

Add `7 7 7 * * "certbot renew -q --cert-name tjsj.dev --deploy-hook 'systemctl restart tjsj.dev.service'"`, then save and exit.
This will check if the certs need to be renewed on the 7th of every month at 07:07am. If they need to be renewed, the server will automatically be restarted after the new certs are retrieved. The provided systemd service file automatically refreshes the used SSL certs on startup.

### Installing and using the provided systemd service file

Add the appropriate unprivilaged user
```
# useradd tjsj_dev
```

Install the systemd service file
```
# cp ./init/tjsj.dev.service /etc/systemd/system/
```

Enabling the service so that it runs on startup
```
# systemctl enable tjsj.dev.service
```

Running the server
```
# systemctl start tjsj.dev.service
```

The website should now be hosted on your machine accessible via HTTPS: `https://localhost/`

