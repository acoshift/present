FROM acoshift/go-scratch

ADD present /
ADD nevp /
EXPOSE 3999

ENTRYPOINT ["/present"]
