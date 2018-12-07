FROM postgres:11.1
RUN apt-get update #TODO: Squash this
RUN apt-get install -y openssh-server postgresql supervisor curl #TODO: Squash this
RUN curl https://dl.2ndquadrant.com/default/release/get/deb | bash \
  && apt-get install -y postgresql-11-repmgr
ADD ./config/supervisord.conf /etc/supervisor/conf.d/supervisord.conf
#  RUN apt-get clean TODO: Uncomment and squash in the end
  #  RUN rm -rf /var/cache/apt/archives/* /var/lib/apt/lists/*
RUN mkdir -p /run/sshd &&  chmod 0755 /run/sshd
ENTRYPOINT []
CMD ["/usr/bin/supervisord"]
