# Build a Docker image based on the cloud9/workspace Docker image.
FROM cloud9/workspace

# Use default answers for all setup prompts.
ENV DEBIAN_FRONTEND noninteractive

# Remove any default PHP debugging settings. 
# Also remove any default workspace files, as they might be replaced later. 
RUN rm -rf /etc/php5/mods-available/xdebug.ini /home/ubuntu/workspace/*

# Add Ruby Version Manager (rvm) to PATH.
ENV PATH="/usr/local/rvm/bin/:$PATH"

# Set ownership defaults.
RUN chown -R ubuntu:ubuntu /home/ubuntu

# Install Node.js.
RUN curl -sL https://deb.nodesource.com/setup_10.x | bash - && \
    apt-get install nodejs -y && \
    apt-get install golang-go -y

# Populate environment variables that are required by Cloud9 Core.
RUN echo "export USER=ubuntu\n\
export C9_PROJECT=c9-offline\n\
export C9_USER=ubuntu\n\
export C9_HOSTNAME=\$IP\n\
export C9_PORT=\$PORT\n\
export IDE_OFFLINE=1\n\
alias c9=/var/c9sdk/bin/c9" >/etc/profile.d/offline.sh

USER ubuntu

# Download Cloud9 Core into the Docker container.
WORKDIR /var
RUN sudo rm -rf c9sdk && \
    sudo mkdir c9sdk && \
    sudo chown ubuntu:ubuntu c9sdk && \
    git clone https://github.com/c9/core.git c9sdk

RUN sudo mkdir -p /home/ubuntu/workspace/
# Install Cloud9 Core within the Docker container.
WORKDIR c9sdk
RUN scripts/install-sdk.sh

# Set additional ownership defaults.
RUN sudo chown -R ubuntu:ubuntu /home/ubuntu/workspace/ && \
    sudo chown -R ubuntu:ubuntu /home/ubuntu/.c9/

# Have the Docker container listen to these ports at run time.
EXPOSE 5050 8080 8081 8082

# Run Cloud9 Core within the Docker container.
ENTRYPOINT ["node", "server.js", \
            "-w", "/home/ubuntu/workspace", \
            "--auth", ":", \
            "--listen", "0.0.0.0", \
            "--port", "5050"]