FROM concourse/dev

# install helm
ADD ./helm-release/version ./helm/version
RUN set -x && \
    curl -SL -o helm.tgz https://storage.googleapis.com/kubernetes-helm/helm-v$(cat ./helm/version)-linux-amd64.tar.gz && \
    tar xzf ./helm.tgz -C /usr/local/bin linux-amd64/helm --strip-components=1


# install kubectl
ADD ./kubernetes-release/version ./kubernetes/version
RUN set -x && \
    curl -SL -o /usr/local/bin/kubectl https://storage.googleapis.com/kubernetes-release/release/v$(cat ./kubernetes/version)/bin/linux/amd64/kubectl && \
    chmod +x /usr/local/bin/kubectl
