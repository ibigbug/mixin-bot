FROM scratch

ADD ./main /main

ENV ImageDir /var/image

CMD /main
