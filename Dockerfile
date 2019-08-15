FROM scratch
EXPOSE 8080
ENTRYPOINT ["/objectstorage"]
COPY ./bin/ /