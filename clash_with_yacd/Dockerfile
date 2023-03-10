FROM dreamacro/clash:latest

RUN mv clash /usr/local/bin/clash && ln -s /root/.config/clash /clash && cd /clash &&\
    wget https://github.com/haishanh/yacd/releases/latest/download/yacd.tar.xz &&\
    tar xf yacd.tar.xz && rm -rf yacd.tar.xz

ADD ui.yaml /clash/ui.yaml

ENTRYPOINT [ "/usr/local/bin/clash", "-d", "/clash" ]