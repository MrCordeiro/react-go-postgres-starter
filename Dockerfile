FROM mcr.microsoft.com/devcontainers/go:1-1.22-bookworm

ARG USERNAME=dev-user
# Don't use UID 1000 to avoid conflicts with the default vscode user, from the
# Microsoft base image
ARG USER_UID=1001
ARG USER_GID=$USER_UID
ARG APP_HOME=/app


# [Optional] Uncomment this section to install additional OS packages.
# RUN apt-get update && export DEBIAN_FRONTEND=noninteractive \
#     && apt-get -y install --no-install-recommends <your-package-list-here>

# devcontainer dependencies and utils
RUN apt-get update && apt-get install --no-install-recommends -y \
  sudo git bash-completion nano ssh

# Create devcontainer user and add it to sudoers
RUN groupadd --gid $USER_GID $USERNAME \
  && useradd --uid $USER_UID --gid $USER_GID -m $USERNAME \
  # Sudo support
  && apt-get update \
  && apt-get install -y sudo \
  && echo $USERNAME ALL=\(root\) NOPASSWD:ALL > /etc/sudoers.d/$USERNAME \
  && chmod 0440 /etc/sudoers.d/$USERNAME

# [Optional] Uncomment the next lines to use go get to install anything else you need
# USER vscode
# COPY server/go.mod server/go.sum ./
# RUN go mod download
# USER root

# [Optional] Uncomment this line to install global node packages.
# RUN su vscode -c "source /usr/local/share/nvm/nvm.sh && npm install -g <your-package-here>" 2>&1

WORKDIR ${APP_HOME}
COPY . ${APP_HOME}

# Give the devcontainer user ownership of the app folder
RUN chown -R ${USERNAME}:${USERNAME} ${APP_HOME}

# Set the default user
USER $USERNAME
