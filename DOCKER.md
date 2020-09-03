#Docker this project!

Now that we have a Go project with basic functionality, let's migrate this to a Docker container.

##Why Docker?
Docker packages the project and its contents in a virtual container to make it easier to create, deploy and run applications. Docker allows the application to be isolated from its environment, making it handy to work elsewhere other than the local machine. It simplifies the development workflow. For more information, see [Docker](https://www.docker.com/why-docker).

##Steps to migrating the project to Docker

1. Create a [Dockerfile](Dockerfile)  

2. Run the build `docker build -t fancybot .`

3. With the build finished, run Docker with `docker run -d -p 8080:8080 fancybot`

4. We should be able to see running containers now. Run `docker container ls`

5. Shell into the container e.g `docker run --rm -it --entrypoint /bin/bash fancybot`

6. The project is available in the container. Great! Run `go run main.go` and we'll realise it needs the Slack API token!

7. To allow the project to access secret in the container, we will need to set up Docker secrets. Quit the shell and return to the terminal.

8. Run `docker swarm init` to start up docker swarm. This is required prior to creating secrets.

9. Save secret by running ` echo "<token-here>" | docker secret create SLACK_TOKEN -`

10. By default, the secret is saved to `/run/secrets/`. Create a service to access the secrets: ` docker service create --name fancybot --secret SLACK_TOKEN -e FANCYBOT_PASSWORD_FILE=/run/secrets/SLACK_TOKEN fancybot`

11. Run `docker secret ls` and `docker service ls` to view the newly created secret and services. Note that when we run `docker secret inspect <secret-id>`, we won't be able to view the contents of the secret as this is encrypted.

12. Shell into the container again and export secret with `export SLACK_TOKEN=$(cat /run/secrets/SLACK_TOKEN)`

13. The project should be able to access the secret now. Run `go run main.go` and the project should work successfully! Congratulations, we now have a working Go project in a Docker container.

