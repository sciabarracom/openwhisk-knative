FROM actionloop/golang-v1.12:knative-whisk
COPY src /action/src
RUN /bin/proxy -compile main </action/src >/action/exec 2>/action/err
ENV OW_AUTOINIT=/action/exec
ENV OW_AUTOINIT_MAIN=main