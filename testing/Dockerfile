FROM ubuntu:24.04

#RUN apt update && apt install -y bash diffutils
RUN apt update \
    && apt install -y --no-install-recommends \
        bash \
        diffutils \
    && rm -rf /var/lib/apt/lists/*


# Maybe use some other name?
WORKDIR /workspace

